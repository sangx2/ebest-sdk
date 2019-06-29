package ebest

import (
	"sync"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/sangx2/ebest/callback"
	"github.com/sangx2/ebest/interfaces"
	"github.com/sangx2/ebest/wrapper"
)

const (
	// Real 객체 이름 목록
	H1  = "KOSPI호가잔량"
	HA  = "KOSDAQ호가잔량"
	K3  = "KOSDAQ 체결"
	S3  = "KOSPI 체결"
	NWS = "실시간 뉴스 제목 패킷"
	SC0 = "주식주문접수"
	SC1 = "주식주문체결"
	SC2 = "주식주문정정"
	SC3 = "주식주문취소"
	SC4 = "주식주문거부"
)

type Real struct {
	resPath string

	realTrade interfaces.RealTrade
	ew        *wrapper.Ebest

	create chan bool
	done   chan bool
	wg     sync.WaitGroup
}

// NewReal Real 객체 생성
func NewReal(resPath, resName string) *Real {
	var trade interfaces.RealTrade

	switch resName {
	case H1:
		trade = callback.NewH1()
	case HA:
		trade = callback.NewHA()
	case K3:
		trade = callback.NewK3()
	case S3:
		trade = callback.NewS3()
	case NWS:
		trade = callback.NewNWS()
	case SC0:
		trade = callback.NewSC0()
	case SC1:
		trade = callback.NewSC1()
	case SC2:
		trade = callback.NewSC2()
	case SC3:
		trade = callback.NewSC3()
	case SC4:
		trade = callback.NewSC4()
	default:
		return nil
	}

	r := &Real{
		resPath:   resPath,
		realTrade: trade, ew: new(wrapper.Ebest),
		create: make(chan bool, 1), done: make(chan bool, 1),
	}

	r.wg.Add(1)
	ret, _ := createThread(r.createObject, 0)
	if ret == 0 {
		return nil
	}

	for {
		select {
		case <-r.create:
			return r
		default:
			time.Sleep(DELAY_CREATE_THREAD)
		}
	}
}

// Close Real 객체 종료
func (r *Real) Close() {
	r.ew.UnAdviseRealData()

	r.done <- true

	r.wg.Wait()
}

// SetInBlock 블록의 데이터(값)를 inblock으로 설정
func (r Real) SetInBlock(inBlock interface{}) error {
	err := r.realTrade.SetFieldData(r.ew, r.resPath, inBlock)
	if err != nil {
		return err
	}
	return nil
}

// GetOutBlock : 블록의 필드 데이터(값)를 outblock으로 취득
func (r Real) GetOutBlock() interface{} {
	return r.realTrade.GetOutBlock()
}

// GetReceivedRealDataChan 서버로부터 데이터를 수신 했을 때 발생
func (r Real) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return r.realTrade.GetReceivedRealDataChan()
}

// GetReceivedLinkDataChan HTS로부터 연동 정보를 수신 했을 때 발생
func (r Real) GetReceivedLinkDataChan() chan wrapper.XaRealRecieveLinkData {
	return r.realTrade.GetReceivedLinkDataChan()
}

// Start 실시간TR을 등록
func (r Real) Start() {
	r.ew.AdviseRealData()
}

// Stop 한 종목의 실시간TR을 해제
func (r Real) Stop(key string) {
	r.ew.UnAdviseRealDataWithKey(key)
}

func (r *Real) createObject(p uintptr) uintptr {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)

	r.ew.Create("XAReal")

	r.ew.BindEvent(r.realTrade)

	r.create <- true

	for {
		pumpWaitingMessages()

		select {
		case <-r.done:
			// 쓰레드가 종료 되기 전에 메인 쓰레드가 먼처 종료되므로 defer로 처리 불가
			r.ew.UnBindEvent()

			r.ew.Release()

			ole.CoUninitialize()

			r.wg.Done()

			return 0
		default:
			time.Sleep(DELAY_PUMPWAITINGMESSAGES)
		}
	}
}
