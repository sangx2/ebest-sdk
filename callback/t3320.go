package callback

import (
	"errors"
	"github.com/sangx2/ebest/res"
	"github.com/sangx2/ebest/wrapper"
)

// T3320 FNG 요약
type T3320 struct {
	InBlock res.T3320InBlock

	OutBlock  res.T3320OutBlock
	OutBlock1 res.T3320OutBlock1

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT3320() *T3320 {
	return &T3320{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T3320) GetTPS() int {
	return t.TPS
}

func (t T3320) GetLPP() int {
	return t.LPP
}

func (t T3320) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T3320) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T3320) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T3320) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T3320) SetFieldData(e *wrapper.Ebest, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t3320.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T3320InBlock); !ok {
		return errors.New("invalid inBlock1")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t3320InBlock", "gicode", 0, t.InBlock.Gicode)

	return nil
}

func (t T3320) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock, t.OutBlock1}
}

func (t *T3320) ReceivedData(e *wrapper.Ebest, x wrapper.XaQueryReceiveData) {
	t.OutBlock.Upgubunnm = e.GetFieldData("t3320OutBlock", "upgubunnm", 0)
	t.OutBlock.Sijangcd = e.GetFieldData("t3320OutBlock", "sijangcd", 0)
	t.OutBlock.Marketnm = e.GetFieldData("t3320OutBlock", "marketnm", 0)
	t.OutBlock.Company = e.GetFieldData("t3320OutBlock", "company", 0)
	t.OutBlock.Baddress = e.GetFieldData("t3320OutBlock", "baddress", 0)
	t.OutBlock.Btelno = e.GetFieldData("t3320OutBlock", "btelno", 0)
	t.OutBlock.Gsyyyy = e.GetFieldData("t3320OutBlock", "gsyyyy", 0)
	t.OutBlock.Gsmm = e.GetFieldData("t3320OutBlock", "gsmm", 0)
	t.OutBlock.Gsym = e.GetFieldData("t3320OutBlock", "gsym", 0)
	t.OutBlock.Lstprice = e.GetFieldData("t3320OutBlock", "lstprice", 0)
	t.OutBlock.Gstock = e.GetFieldData("t3320OutBlock", "gstock", 0)
	t.OutBlock.Homeurl = e.GetFieldData("t3320OutBlock", "homeurl", 0)
	t.OutBlock.Grdnm = e.GetFieldData("t3320OutBlock", "grdnm", 0)
	t.OutBlock.Foreignratio = e.GetFieldData("t3320OutBlock", "foreignratio", 0)
	t.OutBlock.Irtel = e.GetFieldData("t3320OutBlock", "irtel", 0)
	t.OutBlock.Capital = e.GetFieldData("t3320OutBlock", "capital", 0)
	t.OutBlock.Sigavalue = e.GetFieldData("t3320OutBlock", "sigavalue", 0)
	t.OutBlock.Cashsis = e.GetFieldData("t3320OutBlock", "cashsis", 0)
	t.OutBlock.Cashrate = e.GetFieldData("t3320OutBlock", "cashrate", 0)
	t.OutBlock.Price = e.GetFieldData("t3320OutBlock", "price", 0)
	t.OutBlock.Jnilclose = e.GetFieldData("t3320OutBlock", "jnilclose", 0)

	t.OutBlock1.Gicode = e.GetFieldData("t3320OutBlock1", "gicode", 0)
	t.OutBlock1.Gsym = e.GetFieldData("t3320OutBlock1", "gsym", 0)
	t.OutBlock1.Gsgb = e.GetFieldData("t3320OutBlock1", "gsgb", 0)
	t.OutBlock1.Per = e.GetFieldData("t3320OutBlock1", "per", 0)
	t.OutBlock1.Eps = e.GetFieldData("t3320OutBlock1", "eps", 0)
	t.OutBlock1.Pbr = e.GetFieldData("t3320OutBlock1", "pbr", 0)
	t.OutBlock1.Roa = e.GetFieldData("t3320OutBlock1", "roa", 0)
	t.OutBlock1.Roe = e.GetFieldData("t3320OutBlock1", "roe", 0)
	t.OutBlock1.Ebitda = e.GetFieldData("t3320OutBlock1", "ebitda", 0)
	t.OutBlock1.Evebitda = e.GetFieldData("t3320OutBlock1", "evebitda", 0)
	t.OutBlock1.Par = e.GetFieldData("t3320OutBlock1", "par", 0)
	t.OutBlock1.Sps = e.GetFieldData("t3320OutBlock1", "sps", 0)
	t.OutBlock1.Cps = e.GetFieldData("t3320OutBlock1", "cps", 0)
	t.OutBlock1.Bps = e.GetFieldData("t3320OutBlock1", "bps", 0)
	t.OutBlock1.TPer = e.GetFieldData("t3320OutBlock1", "t_per", 0)
	t.OutBlock1.TEps = e.GetFieldData("t3320OutBlock1", "t_eps", 0)
	t.OutBlock1.Peg = e.GetFieldData("t3320OutBlock1", "peg", 0)
	t.OutBlock1.TPeg = e.GetFieldData("t3320OutBlock1", "t_peg", 0)
	t.OutBlock1.TGsym = e.GetFieldData("t3320OutBlock1", "t_gsym", 0)

	t.ReceiveDataChan <- x
}

func (t T3320) ReceivedMessage(e *wrapper.Ebest, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T3320) ReceivedChartRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T3320) ReceivedSearchRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
