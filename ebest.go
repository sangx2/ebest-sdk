package ebest

import (
	"errors"
	"fmt"
	"sync"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/sangx2/ebest/impl"
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

	resPath string

	cb *impl.Session
	ew *wrapper.Ebest

	createChan chan error
	doneChan   chan bool
	wg         sync.WaitGroup
}

func NewEbest(id, passwd, certPasswd, srvIP string, srvPort int, resPath string) *Ebest {
	e := &Ebest{id: id, passwd: passwd, certPasswd: certPasswd, srvIP: srvIP, srvPort: srvPort,
		cb: impl.NewSession(), ew: new(wrapper.Ebest), resPath: resPath,
		createChan: make(chan error, 1), doneChan: make(chan bool, 1)}

	e.wg.Add(1)
	ret, _ := createThread(e.createObject, 0)
	if ret == 0 {
		return nil
	}

	return e
}

func (e *Ebest) CreateQuery(resName string) *Query {
	var trade interfaces.QueryTrade

	switch resName {
	case CSPAQ12200:
		trade = impl.NewCSPAQ12200()
	case CSPAT00600:
		trade = impl.NewCSPAT00600()
	case CSPAT00700:
		trade = impl.NewCSPAT00700()
	case CSPAT00800:
		trade = impl.NewCSPAT00800()
	case T1101:
		trade = impl.NewT1101()
	case T1511:
		trade = impl.NewT1511()
	case T3320:
		trade = impl.NewT3320()
	case T0424:
		trade = impl.NewT0424()
	case T8424:
		trade = impl.NewT8424()
	case T8436:
		trade = impl.NewT8436()
	default:
		return nil
	}

	return NewQuery(e.resPath, trade)
}

func (e *Ebest) CreateReal(resName string) *Real {
	var trade interfaces.RealTrade

	switch resName {
	case H1:
		trade = impl.NewH1()
	case HA:
		trade = impl.NewHA()
	case K3:
		trade = impl.NewK3()
	case S3:
		trade = impl.NewS3()
	case NWS:
		trade = impl.NewNWS()
	case SC0:
		trade = impl.NewSC0()
	case SC1:
		trade = impl.NewSC1()
	case SC2:
		trade = impl.NewSC2()
	case SC3:
		trade = impl.NewSC3()
	case SC4:
		trade = impl.NewSC4()
	default:
		return nil
	}

	return NewReal(e.resPath, trade)
}

// Connect 서버에 연결
func (e Ebest) Connect() error {
	for {
		select {
		case e := <-e.createChan:
			return e
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

	e.doneChan <- true
	close(e.doneChan)

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

// createObject Ebest 객체 생성시 등록될 callback 함수
func (e *Ebest) createObject(p uintptr) uintptr {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)

	e.ew.Create("XASession")

	if !e.ew.ConnectServer(e.srvIP, e.srvPort) {
		e.createChan <- fmt.Errorf("%s:%s", "s.ew.ConnectServer", e.ew.GetErrorMessage(int(e.ew.GetLastError())))
		close(e.createChan)
		e.wg.Done()
		return 0
	}

	if !e.ew.Login(e.id, e.passwd, e.certPasswd) {
		e.createChan <- fmt.Errorf("%s:%s", "s.ew.Login", e.ew.GetErrorMessage(int(e.ew.GetLastError())))
		close(e.createChan)
		e.wg.Done()
		return 0
	}

	e.ew.BindEvent(e.cb)

	e.createChan <- nil
	close(e.createChan)

	for {
		pumpWaitingMessages()

		select {
		case <-e.doneChan:
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
