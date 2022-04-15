package impl

import (
	"errors"
	"fmt"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// T0424 주식 잔고2
type T0424 struct {
	InBlock res.T0424InBlock

	OutBlock  res.T0424OutBlock
	OutBlock1 []res.T0424OutBlock1

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT0424() *T0424 {
	return &T0424{
		TPS: 1, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T0424) GetTPS() int {
	return t.TPS
}

func (t T0424) GetLPP() int {
	return t.LPP
}

func (t T0424) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T0424) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T0424) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T0424) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T0424) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t0424.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T0424InBlock); !ok {
		return errors.New(fmt.Sprintf("invalid inBlocks:%+v", inBlocks))
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t0424InBlock", "accno", 0, t.InBlock.Accno)
	e.SetFieldData("t0424InBlock", "passwd", 0, t.InBlock.Passwd)
	e.SetFieldData("t0424InBlock", "prcgb", 0, t.InBlock.Prcgb)
	e.SetFieldData("t0424InBlock", "chegb", 0, t.InBlock.Chegb)
	e.SetFieldData("t0424InBlock", "dangb", 0, t.InBlock.Dangb)
	e.SetFieldData("t0424InBlock", "charge", 0, t.InBlock.Charge)
	e.SetFieldData("t0424InBlock", "cts_expcode", 0, t.InBlock.CtsExpcode)

	return nil
}

func (t T0424) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock, t.OutBlock1}
}

func (t *T0424) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	t.OutBlock.Sunamt = e.GetFieldData("t0424OutBlock", "sunamt", 0)
	t.OutBlock.Dtsunik = e.GetFieldData("t0424OutBlock", "dtsunik", 0)
	t.OutBlock.Mamt = e.GetFieldData("t0424OutBlock", "mamt", 0)
	t.OutBlock.Sunamt1 = e.GetFieldData("t0424OutBlock", "sunamt1", 0)
	t.OutBlock.CtsExpcode = e.GetFieldData("t0424OutBlock", "cts_expcode", 0)
	t.OutBlock.Tappamt = e.GetFieldData("t0424OutBlock", "tappamt", 0)
	t.OutBlock.Tdtsunik = e.GetFieldData("t0424OutBlock", "tdtsunik", 0)

	for i := 0; i < int(e.GetBlockCount("t0424OutBlock1")); i++ {
		tr := res.T0424OutBlock1{
			Expcode:    e.GetFieldData("t0424OutBlock1", "expcode", i),
			Jangb:      e.GetFieldData("t0424OutBlock1", "jangb", i),
			Janqty:     e.GetFieldData("t0424OutBlock1", "janqty", i),
			Mdposqt:    e.GetFieldData("t0424OutBlock1", "mdposqt", i),
			Pamt:       e.GetFieldData("t0424OutBlock1", "pamt", i),
			Mamt:       e.GetFieldData("t0424OutBlock1", "mamt", i),
			Sinamt:     e.GetFieldData("t0424OutBlock1", "sinamt", i),
			Lastdt:     e.GetFieldData("t0424OutBlock1", "lastdt", i),
			Msat:       e.GetFieldData("t0424OutBlock1", "msat", i),
			Mpms:       e.GetFieldData("t0424OutBlock1", "mpms", i),
			Mdat:       e.GetFieldData("t0424OutBlock1", "mdat", i),
			Mpmd:       e.GetFieldData("t0424OutBlock1", "mpmd", i),
			Jsat:       e.GetFieldData("t0424OutBlock1", "jsat", i),
			Jpms:       e.GetFieldData("t0424OutBlock1", "jpms", i),
			Jdat:       e.GetFieldData("t0424OutBlock1", "jdat", i),
			Jpmd:       e.GetFieldData("t0424OutBlock1", "jpmd", i),
			Sysprocseq: e.GetFieldData("t0424OutBlock1", "sysprocseq", i),
			Loandt:     e.GetFieldData("t0424OutBlock1", "loandt", i),
			Hname:      e.GetFieldData("t0424OutBlock1", "hname", i),
			Marketgb:   e.GetFieldData("t0424OutBlock1", "marketgb", i),
			Jonggb:     e.GetFieldData("t0424OutBlock1", "jonggb", i),
			Janrt:      e.GetFieldData("t0424OutBlock1", "janrt", i),
			Price:      e.GetFieldData("t0424OutBlock1", "price", i),
			Appamt:     e.GetFieldData("t0424OutBlock1", "appamt", i),
			Dtsunik:    e.GetFieldData("t0424OutBlock1", "dtsunik", i),
			Sunikrt:    e.GetFieldData("t0424OutBlock1", "sunikrt", i),
			Fee:        e.GetFieldData("t0424OutBlock1", "fee", i),
			Tax:        e.GetFieldData("t0424OutBlock1", "tax", i),
			Sininter:   e.GetFieldData("t0424OutBlock1", "sininter", i)}

		t.OutBlock1 = append(t.OutBlock1, tr)
	}

	t.ReceiveDataChan <- x
}

func (t T0424) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T0424) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T0424) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
