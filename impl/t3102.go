package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// T3102 뉴스본문
type T3102 struct {
	InBlock res.T3102InBlock

	OutBlock  []res.T3102OutBlock
	OutBlock1 []res.T3102OutBlock1
	OutBlock2 res.T3102OutBlock2

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT3102() *T3102 {
	return &T3102{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T3102) GetTPS() int {
	return t.TPS
}

func (t T3102) GetLPP() int {
	return t.LPP
}

func (t T3102) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T3102) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T3102) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T3102) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T3102) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t3102.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T3102InBlock); !ok {
		return errors.New("invalid inBlock")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t3102InBlock", "sNewsno", 0, t.InBlock.SNewsno)

	return nil
}

func (t T3102) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock, t.OutBlock1, t.OutBlock2}
}

func (t *T3102) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	for i := 0; i < int(e.GetBlockCount("t3102OutBlock")); i++ {
		tr := res.T3102OutBlock{
			SJongcode: e.GetFieldData("t3102OutBlock", "sJongcode", i),
		}
		t.OutBlock = append(t.OutBlock, tr)
	}

	for i := 0; i < int(e.GetBlockCount("t3102OutBlock1")); i++ {
		tr := res.T3102OutBlock1{
			SBody: e.GetFieldData("t3102OutBlock1", "sBody", i),
		}
		t.OutBlock1 = append(t.OutBlock1, tr)
	}

	t.OutBlock2.STitle = e.GetFieldData("t3102OutBlock2", "sTitle", 0)

	t.ReceiveDataChan <- x
}

func (t T3102) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T3102) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T3102) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
