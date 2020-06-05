package ebest

import (
	"fmt"
	"sync"
	"time"

	ole "github.com/go-ole/go-ole"
	"github.com/sangx2/ebest/callback"
	"github.com/sangx2/ebest/interfaces"
	"github.com/sangx2/ebest/wrapper"
)

const (
	// Query 객체 이름 목록
	CSPAQ12200 = "현물계좌 예수금/주문가능금액/총평가 조회"
	CSPAT00600 = "현물 정상주문"
	CSPAT00700 = "현물 정정주문"
	CSPAT00800 = "현물 취소주문"
	T0424      = "주식 잔고2"
	T1101      = "주식 현재가 호가 조회"
	T1511      = "업종 현재가"
	T3320      = "FNG 요약"
	T8424      = "업종전체조회"
	T8436      = "주식종목조회"
)

type Query struct {
	resPath string

	TPS, LPP   int
	queryTrade interfaces.QueryTrade

	ew *wrapper.Ebest

	create chan bool
	done   chan bool
	wg     sync.WaitGroup
}

// NewQuery Query 객체 생성
func NewQuery(resPath, resName string) *Query {
	var trade interfaces.QueryTrade

	switch resName {
	case CSPAQ12200:
		trade = callback.NewCSPAQ12200()
	case CSPAT00600:
		trade = callback.NewCSPAT00600()
	case CSPAT00700:
		trade = callback.NewCSPAT00700()
	case CSPAT00800:
		trade = callback.NewCSPAT00800()
	case T1101:
		trade = callback.NewT1101()
	case T1511:
		trade = callback.NewT1511()
	case T3320:
		trade = callback.NewT3320()
	case T0424:
		trade = callback.NewT0424()
	case T8424:
		trade = callback.NewT8424()
	case T8436:
		trade = callback.NewT8436()
	default:
		return nil
	}

	q := &Query{
		resPath: resPath,
		TPS:     trade.GetTPS(), LPP: trade.GetLPP(),
		queryTrade: trade, ew: new(wrapper.Ebest),
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
			time.Sleep(DELAY_CREATE_THREAD)
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
func (q Query) SetInBlock(inBlocks ...interface{}) error {
	return q.queryTrade.SetFieldData(q.ew, q.resPath, inBlocks[:]...)
}

// GetOutBlock : 블록의 필드 데이터(값)를 outblock으로 취득
func (q Query) GetOutBlocks() []interface{} {
	return q.queryTrade.GetOutBlocks()
}

// Request 조회 TR을 요청
func (q Query) Request() int {
	return int(q.ew.Request(false))
}

// GetReceiveDataChan 서버로부터 데이터를 수신 했을 때 발생
func (q Query) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return q.queryTrade.GetReceiveDataChan()
}

// GetReceiveMessage 서버로부터 메시지를 수신 했을 때 발생
func (q Query) GetReceiveMessage() (string, error) {
	msg := <-q.queryTrade.GetReceiveMessageChan()
	if msg.IsSystemError < 0 {
		return "", fmt.Errorf("%s:%s", msg.MessageCode, msg.Message)
	}
	return fmt.Sprintf("%s:%s", msg.MessageCode, msg.Message), nil
}

// GetReceiveChartRealData 차트 지표 실시간 데이터를 수신 했을 때 발생
//
// (차트 지표데이터 조회 시, 실시간 자동 등록을 “1”로 했을 경우)
func (q Query) GetReceiveChartRealData() chan wrapper.XaQueryReceiveChartRealData {
	return q.queryTrade.GetReceiveChartRealDataChan()
}

// GetReceiveSearchRealData 종목검색 실시간 데이터를 수신 했을 때 발생
//
// (종목검색(신버전API용) 조회 시, 실시간 등록을 “1”로 했을 경우)
func (q Query) GetReceiveSearchRealData() chan wrapper.XaQueryReceiveSearchRealData {
	return q.queryTrade.GetReceiveChartSearchRealDataChan()
}

func (q *Query) createObject(p uintptr) uintptr {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED)

	q.ew.Create("XAQuery")

	q.ew.BindEvent(q.queryTrade)

	q.create <- true
	close(q.create)

	for {
		pumpWaitingMessages()

		select {
		case <-q.done:
			// 쓰레드가 종료 되기 전에 메인 쓰레드가 먼처 종료되므로 defer로 처리 불가
			q.ew.UnBindEvent()

			q.ew.Release()

			ole.CoUninitialize()

			q.wg.Done()

			return 0
		default:
			time.Sleep(DELAY_PUMPWAITINGMESSAGES)
		}
	}
}
