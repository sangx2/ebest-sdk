package ebest

import (
	"errors"
	"fmt"
	"sync"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/sangx2/ebest/callback"
	"github.com/sangx2/ebest/interfaces"
	"github.com/sangx2/ebest/wrapper"
)

const (
	// SERVER_REAL 실제 서버
	SERVER_REAL = "hts.ebestsec.co.kr"
	// SERVER_VIRTUAL 가상 서버
	SERVER_VIRTUAL = "demo.ebestsec.co.kr"

	PORT = 20001

	DELAY_CREATE_THREAD       = time.Millisecond * 10
	DELAY_PUMPWAITINGMESSAGES = time.Millisecond * 10
)

type Ebest struct {
	id         string
	passwd     string
	certPasswd string

	srvIP   string
	srvPort int

	log interfaces.Logger

	cb *callback.Session
	ew *wrapper.Ebest

	create chan bool
	done   chan bool
	wg     sync.WaitGroup
}

func NewEbest(id, passwd, certPasswd, srvIP string, srvPort int, logger interfaces.Logger) *Ebest {
	e := &Ebest{id: id, passwd: passwd, certPasswd: certPasswd, srvIP: srvIP, srvPort: srvPort, log: logger,
		cb: callback.NewSession(), ew: new(wrapper.Ebest),
		create: make(chan bool, 1), done: make(chan bool, 1)}

	e.wg.Add(1)
	ret, _ := createThread(e.createObject, 0)
	if ret == 0 {
		return nil
	}

	return e
}

// Connect 서버에 연결
func (e Ebest) Connect() {
	for {
		select {
		case <-e.create:
			return
		default:
			time.Sleep(DELAY_CREATE_THREAD)
		}
	}
}

// Login 로그인 이벤트를 대기
func (e Ebest) Login() error {
	sessionLogin := <-e.cb.GetSessionLoginChan()
	if sessionLogin.Code != "0000" {
		return errors.New(sessionLogin.Msg)
	}

	return nil
}

// Disconnect 서버 연결 종료
func (e *Ebest) Disconnect() {
	e.ew.DisconnectServer()

	e.done <- true

	e.wg.Wait()
}

// GetAccountList 보유중인 계좌리스트를 취득
func (e Ebest) GetAccountList() []string {
	var accountList []string

	cnt := e.ew.GetAccountListCount()

	for i := 0; i < int(cnt); i++ {
		accountList = append(accountList, e.ew.GetAccountList(i))
	}

	return accountList
}

// GetAccountName 계좌명을 취득
func (e Ebest) GetAccountName(account string) string {
	return e.ew.GetAccountName(account)
}

// GetAccountDetailName 계좌 상세명을 취득
func (e Ebest) GetAccountDetailName(account string) string {
	return e.ew.GetAcctDetailName(account)
}

// GetAccountNickName 계좌 별명을 취득
func (e Ebest) GetAccountNickName(account string) string {
	return e.ew.GetAcctNickName(account)
}

// GetErrorMessage 에러코드값에 대한 에러메시지를 취득
func (e Ebest) GetErrorMessage(code int) string {
	return e.ew.GetErrorMessage(code)
}

func (e *Ebest) createObject(p uintptr) uintptr {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)

	e.ew.Create("XASession")

	if !e.ew.ConnectServer(e.srvIP, e.srvPort) {
		if e.log != nil {
			e.log.Error(fmt.Sprintf("%s:%s", "s.ew.ConnectServer", e.ew.GetErrorMessage(int(e.ew.GetLastError()))))
		}
	} else {
		if e.log != nil {
			e.log.Info("ConnectServer ok")
		}
	}

	if !e.ew.Login(e.id, e.passwd, e.certPasswd) {
		if e.log != nil {
			e.log.Error(fmt.Sprintf("%s:%s", "s.ew.Login", e.ew.GetErrorMessage(int(e.ew.GetLastError()))))
		}
	} else {
		if e.log != nil {
			e.log.Info("Login ok")
		}
	}

	e.ew.BindEvent(e.cb)

	e.create <- true

	for {
		pumpWaitingMessages()

		select {
		case <-e.done:
			// 쓰레드가 종료 되기 전에 메인 쓰레드가 먼처 종료되므로 defer로 처리 불가
			e.ew.UnBindEvent()

			e.ew.Release()

			ole.CoUninitialize()

			e.wg.Done()
			return 0
		default:
			time.Sleep(DELAY_PUMPWAITINGMESSAGES)
		}

	}

}
