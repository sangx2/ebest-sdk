package ebest

import (
	"errors"
	"fmt"
	"github.com/sangx2/ebest-sdk/interfaces"
	"sync"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/sangx2/ebest-sdk/impl"
	"github.com/sangx2/ebest-sdk/wrapper"
)

const (
	// ServerReal 실제 서버
	ServerReal = "hts.ebestsec.co.kr"
	// ServerVirtual 가상 서버
	ServerVirtual = "demo.ebestsec.co.kr"

	Port = 20001

	DelayCreateThread        = time.Millisecond * 10
	DelayPumpWaitingMessages = time.Millisecond * 10
)

const (
	// Query 객체 이름 목록
	CSPAQ12200 = "현물계좌 예수금/주문가능금액/총평가 조회"
	CSPAT00600 = "현물 정상주문"
	CSPAT00700 = "현물 정정주문"
	CSPAT00800 = "현물 취소주문"
	T0424      = "주식 잔고2"
	T1101      = "주식 현재가 호가 조회"
	T1305      = "기간별 추가"
	T1511      = "업종 현재가"
	T3320      = "FNG 요약"
	T8424      = "업종전체조회"
	T8436      = "주식종목조회"
)

const (
	// Real 객체 이름 목록
	H1  = "KOSPI 호가잔량"
	HA  = "KOSDAQ 호가잔량"
	S3  = "KOSPI 체결"
	K3  = "KOSDAQ 체결"
	K1  = "KOSPI 거래원"
	OK  = "KOSDAQ 거래원"
	NWS = "실시간 뉴스 제목 패킷"
	SC0 = "주식주문접수"
	SC1 = "주식주문체결"
	SC2 = "주식주문정정"
	SC3 = "주식주문취소"
	SC4 = "주식주문거부"
)

type EBest struct {
	id           string
	password     string
	certPassword string

	srvIP   string
	srvPort int

	resPath string

	cb *impl.Session
	ew *wrapper.EBestWrapper

	createChan chan error
	doneChan   chan bool
	wg         sync.WaitGroup
}

// NewEBest ebest 객체 생성
func NewEBest(id, password, certPassword, srvIP string, srvPort int, resPath string) *EBest {
	e := &EBest{id: id, password: password, certPassword: certPassword, srvIP: srvIP, srvPort: srvPort,
		cb: impl.NewSession(), ew: new(wrapper.EBestWrapper), resPath: resPath,
		createChan: make(chan error, 1), doneChan: make(chan bool, 1)}

	e.wg.Add(1)
	ret, _ := createThread(e.createObject, 0)
	if ret == 0 {
		return nil
	}

	return e
}

// NewQuery query 객체 생성
func (e *EBest) NewQuery(trade interfaces.QueryTrade) *Query {
	return NewQuery(e.resPath, trade)
}

// NewReal real 객체 생성
func (e *EBest) NewReal(trade interfaces.RealTrade) *Real {
	return NewReal(e.resPath, trade)
}

// Connect 서버에 연결
func (e *EBest) Connect() error {
	for {
		select {
		case e := <-e.createChan:
			return e
		default:
			time.Sleep(DelayCreateThread)
		}
	}
}

// Login 로그인 이벤트를 대기
func (e *EBest) Login() error {
	sessionLogin := <-e.cb.GetSessionLoginChan()
	if sessionLogin.Code != "0000" {
		return errors.New(sessionLogin.Msg)
	}
	return nil
}

// Disconnect 서버 연결 종료
func (e *EBest) Disconnect() {
	e.ew.DisconnectServer()

	e.doneChan <- true

	e.wg.Wait()
}

// GetAccountList 보유중인 계좌리스트를 취득
func (e *EBest) GetAccountList() []string {
	var accountList []string

	cnt := e.ew.GetAccountListCount()

	for i := 0; i < int(cnt); i++ {
		accountList = append(accountList, e.ew.GetAccountList(i))
	}

	return accountList
}

// GetAccountName 계좌명을 취득
func (e *EBest) GetAccountName(account string) string {
	return e.ew.GetAccountName(account)
}

// GetAccountDetailName 계좌 상세명을 취득
func (e *EBest) GetAccountDetailName(account string) string {
	return e.ew.GetAcctDetailName(account)
}

// GetAccountNickName 계좌 별명을 취득
func (e *EBest) GetAccountNickName(account string) string {
	return e.ew.GetAcctNickName(account)
}

// GetErrorMessage 에러코드값에 대한 에러메시지를 취득
func (e *EBest) GetErrorMessage(code int) string {
	return e.ew.GetErrorMessage(code)
}

// createObject EBest 객체 생성시 등록될 callback 함수
func (e *EBest) createObject(p uintptr) uintptr {
	defer e.wg.Done()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	defer ole.CoUninitialize()

	e.ew.Create("XASession")
	defer e.ew.Release()

	if !e.ew.ConnectServer(e.srvIP, e.srvPort) {
		e.createChan <- fmt.Errorf("%s:%s", "s.ew.ConnectServer", e.ew.GetErrorMessage(int(e.ew.GetLastError())))
		close(e.createChan)
		e.wg.Done()
		return 0
	}

	if !e.ew.Login(e.id, e.password, e.certPassword) {
		e.createChan <- fmt.Errorf("%s:%s", "s.ew.Login", e.ew.GetErrorMessage(int(e.ew.GetLastError())))
		close(e.createChan)
		e.wg.Done()
		return 0
	}

	e.ew.BindEvent(e.cb)
	defer e.ew.UnBindEvent()

	e.createChan <- nil
	close(e.createChan)

	for {
		pumpWaitingMessages()

		select {
		case <-e.doneChan:
			return 0
		default:
			time.Sleep(DelayPumpWaitingMessages)
		}
	}
}
