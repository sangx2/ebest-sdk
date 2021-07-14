package ebest

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/sangx2/ebest-sdk/interfaces"
	"github.com/sangx2/ebest-sdk/wrapper"
)

type Query struct {
	resPath string

	TPS, LPP   int
	queryTrade interfaces.QueryTrade

	ew *wrapper.EBestWrapper

	create chan bool
	done   chan bool
	wg     sync.WaitGroup
}

// NewQuery Query 객체 생성
func NewQuery(resPath string, trade interfaces.QueryTrade) *Query {
	q := &Query{
		resPath: resPath,

		TPS: trade.GetTPS(), LPP: trade.GetLPP(),
		queryTrade: trade, ew: new(wrapper.EBestWrapper),

		create: make(chan bool, 1), done: make(chan bool, 1),
	}

	q.wg.Add(1)
	ret, _ := createThread(q.createObject, 0)
	if ret == 0 {
		return nil
	}

	for {
		select {
		case <-q.create:
			return q
		default:
			time.Sleep(DelayCreateThread)
		}
	}
}

// Close Query 객체 종료
func (q *Query) Close() {
	q.done <- true
	close(q.done)

	q.wg.Wait()
}

// SetInBlock 블록의 데이터(값)를 inblock으로 설정
func (q *Query) SetInBlock(inBlocks ...interface{}) error {
	err := q.queryTrade.SetFieldData(q.ew, q.resPath, inBlocks[:]...)

	if err != nil {
		return NewErrQuery("SetInBlock", "", err.Error())
	}
	return nil
}

// GetOutBlocks : 블록의 필드 데이터(값)를 outblock으로 취득
func (q *Query) GetOutBlocks() []interface{} {
	return q.queryTrade.GetOutBlocks()
}

// Request 조회 TR을 요청
func (q *Query) Request(isOccurs bool) int {
	return int(q.ew.Request(isOccurs))
}

// GetReceiveDataChan 서버로부터 데이터를 수신 했을 때 발생
func (q *Query) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return q.queryTrade.GetReceiveDataChan()
}

// GetReceiveMessage 서버로부터 메시지를 수신 했을 때 발생
func (q *Query) GetReceiveMessage() (string, error) {
	msg := <-q.queryTrade.GetReceiveMessageChan()
	if msg.IsSystemError < 0 {
		return "", NewErrQuery("GetReceiveMessage", msg.MessageCode, msg.Message)
	}
	return fmt.Sprintf("%s: %s", msg.MessageCode, msg.Message), nil
}

// GetReceiveChartRealData 차트 지표 실시간 데이터를 수신 했을 때 발생
//
// (차트 지표데이터 조회 시, 실시간 자동 등록을 “1”로 했을 경우)
func (q *Query) GetReceiveChartRealData() chan wrapper.XaQueryReceiveChartRealData {
	return q.queryTrade.GetReceiveChartRealDataChan()
}

// GetReceiveSearchRealData 종목검색 실시간 데이터를 수신 했을 때 발생
//
// (종목검색(신버전API용) 조회 시, 실시간 등록을 “1”로 했을 경우)
func (q *Query) GetReceiveSearchRealData() chan wrapper.XaQueryReceiveSearchRealData {
	return q.queryTrade.GetReceiveChartSearchRealDataChan()
}

func (q *Query) createObject(p uintptr) uintptr {
	defer q.wg.Done()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)
	defer ole.CoUninitialize()

	q.ew.Create("XAQuery")
	defer q.ew.Release()

	q.ew.BindEvent(q.queryTrade)
	defer q.ew.UnBindEvent()

	q.create <- true
	close(q.create)

	for {
		pumpWaitingMessages()

		select {
		case <-q.done:
			return 0
		default:
			time.Sleep(DelayPumpWaitingMessages)
		}
	}
}
