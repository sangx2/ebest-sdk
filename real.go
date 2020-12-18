package ebest

import (
	"sync"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/sangx2/ebest/interfaces"
	"github.com/sangx2/ebest/wrapper"
)

type Real struct {
	resPath string

	realTrade interfaces.RealTrade
	ew        *wrapper.EBestWrapper

	create chan bool
	done   chan bool
	wg     sync.WaitGroup
}

// NewReal Real 객체 생성
func NewReal(resPath string, trade interfaces.RealTrade) *Real {
	r := &Real{
		resPath: resPath,

		realTrade: trade, ew: new(wrapper.EBestWrapper),

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
			time.Sleep(DelayCreateThread)
		}
	}
}

// Close Real 객체 종료
func (r *Real) Close() {
	r.ew.UnAdviseRealData()

	r.done <- true
	close(r.done)

	r.wg.Wait()
}

// SetInBlock 블록의 데이터(값)를 inblock으로 설정
func (r *Real) SetInBlock(inBlock interface{}) error {
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
func (r Real) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
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
	defer r.wg.Done()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	defer ole.CoUninitialize()

	r.ew.Create("XAReal")
	defer r.ew.Release()

	r.ew.BindEvent(r.realTrade)
	defer r.ew.UnBindEvent()

	r.create <- true
	close(r.create)

	for {
		pumpWaitingMessages()

		select {
		case <-r.done:
			return 0
		default:
			time.Sleep(DelayPumpWaitingMessages)
		}
	}
}
