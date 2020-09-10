package impl

import (
	"errors"
	"github.com/sangx2/ebest/res"
	"github.com/sangx2/ebest/wrapper"
)

// T8436 주식종목조회 API용
type T8436 struct {
	InBlock  res.T8436InBlock
	OutBlock []res.T8436OutBlock

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT8436() *T8436 {
	return &T8436{
		TPS: 2, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T8436) GetTPS() int {
	return t.TPS
}

func (t T8436) GetLPP() int {
	return t.LPP
}

func (t T8436) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T8436) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T8436) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T8436) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T8436) SetFieldData(e *wrapper.Ebest, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t8436.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T8436InBlock); !ok {
		return errors.New("invalid inBlock1")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t8436InBlock", "gubun", 0, t.InBlock.Gubun)

	return nil
}

func (t T8436) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock}
}

func (t *T8436) ReceivedData(e *wrapper.Ebest, x wrapper.XaQueryReceiveData) {
	TRcount := e.GetBlockCount("t8436OutBlock")

	t.OutBlock = nil

	for i := 0; i < int(TRcount); i++ {
		tr := res.T8436OutBlock{
			Hname:      e.GetFieldData("t8436OutBlock", "hname", i),
			Shcode:     e.GetFieldData("t8436OutBlock", "shcode", i),
			Expcode:    e.GetFieldData("t8436OutBlock", "expcode", i),
			Etfgubun:   e.GetFieldData("t8436OutBlock", "etfgubun", i),
			Uplmtprice: e.GetFieldData("t8436OutBlock", "uplmtprice", i),
			Dnlmtprice: e.GetFieldData("t8436OutBlock", "dnlmtprice", i),
			Jnilclose:  e.GetFieldData("t8436OutBlock", "jnilclose", i),
			Memedan:    e.GetFieldData("t8436OutBlock", "memedan", i),
			Recprice:   e.GetFieldData("t8436OutBlock", "recprice", i),
			Gubun:      e.GetFieldData("t8436OutBlock", "gubun", i),
			Bu12gubun:  e.GetFieldData("t8436OutBlock", "bu12gubun", i),
			SpacGubun:  e.GetFieldData("t8436OutBlock", "spac_gubun", i),
			Filler:     e.GetFieldData("t8436OutBlock", "filler", i)}

		t.OutBlock = append(t.OutBlock, tr)
	}

	t.ReceiveDataChan <- x
}

func (t T8436) ReceivedMessage(e *wrapper.Ebest, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T8436) ReceivedChartRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T8436) ReceivedSearchRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
