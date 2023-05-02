package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
	"strconv"
)

// T1305 기간별 주가
type T1305 struct {
	InBlock res.T1305InBlock

	OutBlock   res.T1305OutBlock
	OutBlock1s []res.T1305OutBlock1

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT1305() *T1305 {
	return &T1305{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
		OutBlock1s:                     make([]res.T1305OutBlock1, 300),
	}
}

func (t T1305) GetTPS() int {
	return t.TPS
}

func (t T1305) GetLPP() int {
	return t.LPP
}

func (t T1305) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T1305) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T1305) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T1305) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T1305) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "T1305.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T1305InBlock); !ok {
		return errors.New("invalid inBlock1")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("T1305InBlock", "shcode", 0, t.InBlock.Shcode)
	e.SetFieldData("T1305InBlock", "dwmcode", 0, t.InBlock.Dwmcode)
	e.SetFieldData("T1305InBlock", "date", 0, t.InBlock.Date)
	e.SetFieldData("T1305InBlock", "idx", 0, t.InBlock.Idx)
	e.SetFieldData("T1305InBlock", "cnt", 0, t.InBlock.Cnt)

	return nil
}

func (t T1305) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock, t.OutBlock1s}
}

func (t T1305) GetBlockData(e *wrapper.EBestWrapper, blockName string) string {
	return e.GetBlockData(blockName)
}

func (t *T1305) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	t.OutBlock.Cnt = e.GetFieldData("T1305OutBlock", "cnt", 0)
	t.OutBlock.Date = e.GetFieldData("T1305OutBlock", "date", 0)
	t.OutBlock.Idx = e.GetFieldData("T1305OutBlock", "idx", 0)

	cnt, _ := strconv.Atoi(t.InBlock.Cnt)
	for index := 0; index < cnt; index++ {
		t.OutBlock1s[index].Date = e.GetFieldData("T1305OutBlock1", "date", index)
		t.OutBlock1s[index].Open = e.GetFieldData("T1305OutBlock1", "open", index)
		t.OutBlock1s[index].High = e.GetFieldData("T1305OutBlock1", "high", index)
		t.OutBlock1s[index].Low = e.GetFieldData("T1305OutBlock1", "low", index)
		t.OutBlock1s[index].Close = e.GetFieldData("T1305OutBlock1", "close", index)
		t.OutBlock1s[index].Sign = e.GetFieldData("T1305OutBlock1", "sign", index)
		t.OutBlock1s[index].Change = e.GetFieldData("T1305OutBlock1", "change", index)
		t.OutBlock1s[index].Diff = e.GetFieldData("T1305OutBlock1", "diff", index)
		t.OutBlock1s[index].Volume = e.GetFieldData("T1305OutBlock1", "volume", index)
		t.OutBlock1s[index].Diff_vol = e.GetFieldData("T1305OutBlock1", "diff_vol", index)
		t.OutBlock1s[index].Chdegree = e.GetFieldData("T1305OutBlock1", "chdegree", index)
		t.OutBlock1s[index].Sojinrate = e.GetFieldData("T1305OutBlock1", "sojinrate", index)
		t.OutBlock1s[index].Changerate = e.GetFieldData("T1305OutBlock1", "changerate", index)
		t.OutBlock1s[index].Fpvolume = e.GetFieldData("T1305OutBlock1", "fpvolume", index)
		t.OutBlock1s[index].Covolume = e.GetFieldData("T1305OutBlock1", "covolume", index)
		t.OutBlock1s[index].Shcode = e.GetFieldData("T1305OutBlock1", "shcode", index)
		t.OutBlock1s[index].Value = e.GetFieldData("T1305OutBlock1", "value", index)
		t.OutBlock1s[index].Ppvolume = e.GetFieldData("T1305OutBlock1", "ppvolume", index)
		t.OutBlock1s[index].O_sign = e.GetFieldData("T1305OutBlock1", "o_sign", index)
		t.OutBlock1s[index].O_change = e.GetFieldData("T1305OutBlock1", "o_change", index)
		t.OutBlock1s[index].O_diff = e.GetFieldData("T1305OutBlock1", "o_diff", index)
		t.OutBlock1s[index].H_sign = e.GetFieldData("T1305OutBlock1", "h_sign", index)
		t.OutBlock1s[index].H_change = e.GetFieldData("T1305OutBlock1", "h_change", index)
		t.OutBlock1s[index].H_diff = e.GetFieldData("T1305OutBlock1", "h_diff", index)
		t.OutBlock1s[index].L_sign = e.GetFieldData("T1305OutBlock1", "l_sign", index)
		t.OutBlock1s[index].L_change = e.GetFieldData("T1305OutBlock1", "l_change", index)
		t.OutBlock1s[index].L_diff = e.GetFieldData("T1305OutBlock1", "l_diff", index)
		t.OutBlock1s[index].Marketcap = e.GetFieldData("T1305OutBlock1", "marketcap", index)
	}

	t.ReceiveDataChan <- x
}

func (t T1305) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T1305) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T1305) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
