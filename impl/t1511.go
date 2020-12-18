package impl

import (
	"errors"
	"github.com/sangx2/ebest/res"
	"github.com/sangx2/ebest/wrapper"
)

// T1511 업종 현재가
type T1511 struct {
	InBlock res.T1511InBlock

	OutBlock res.T1511OutBlock

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT1511() *T1511 {
	return &T1511{
		TPS: 5, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T1511) GetTPS() int {
	return t.TPS
}

func (t T1511) GetLPP() int {
	return t.LPP
}

func (t T1511) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T1511) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T1511) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T1511) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T1511) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t1511.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T1511InBlock); !ok {
		return errors.New("invalid inBlock1")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t1511InBlock", "upcode", 0, t.InBlock.Upcode)

	return nil
}

func (t T1511) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock}
}

func (t *T1511) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	t.OutBlock.Gubun = e.GetFieldData("t1511OutBlock", "gubun", 0)
	t.OutBlock.Hname = e.GetFieldData("t1511OutBlock", "hname", 0)
	t.OutBlock.Pricejisu = e.GetFieldData("t1511OutBlock", "pricejisu", 0)
	t.OutBlock.Jniljisu = e.GetFieldData("t1511OutBlock", "jniljisu", 0)
	t.OutBlock.Sign = e.GetFieldData("t1511OutBlock", "sign", 0)
	t.OutBlock.Change = e.GetFieldData("t1511OutBlock", "change", 0)
	t.OutBlock.Diffjisu = e.GetFieldData("t1511OutBlock", "diffjisu", 0)
	t.OutBlock.Jnilvolume = e.GetFieldData("t1511OutBlock", "jnilvolume", 0)
	t.OutBlock.Volume = e.GetFieldData("t1511OutBlock", "volume", 0)
	t.OutBlock.Volumechange = e.GetFieldData("t1511OutBlock", "volumechange", 0)
	t.OutBlock.Volumerate = e.GetFieldData("t1511OutBlock", "volumerate", 0)
	t.OutBlock.Jnilvalue = e.GetFieldData("t1511OutBlock", "jnilvalue", 0)
	t.OutBlock.Value = e.GetFieldData("t1511OutBlock", "value", 0)
	t.OutBlock.Valuechange = e.GetFieldData("t1511OutBlock", "valuechange", 0)
	t.OutBlock.Valuerate = e.GetFieldData("t1511OutBlock", "valuerate", 0)
	t.OutBlock.Openjisu = e.GetFieldData("t1511OutBlock", "openjisu", 0)
	t.OutBlock.Opendiff = e.GetFieldData("t1511OutBlock", "opendiff", 0)
	t.OutBlock.Opentime = e.GetFieldData("t1511OutBlock", "opentime", 0)
	t.OutBlock.Highjisu = e.GetFieldData("t1511OutBlock", "highjisu", 0)
	t.OutBlock.Highdiff = e.GetFieldData("t1511OutBlock", "highdiff", 0)
	t.OutBlock.Hightime = e.GetFieldData("t1511OutBlock", "hightime", 0)
	t.OutBlock.Lowjisu = e.GetFieldData("t1511OutBlock", "lowjisu", 0)
	t.OutBlock.Lowdiff = e.GetFieldData("t1511OutBlock", "lowdiff", 0)
	t.OutBlock.Lowtime = e.GetFieldData("t1511OutBlock", "lowtime", 0)
	t.OutBlock.Whjisu = e.GetFieldData("t1511OutBlock", "whjisu", 0)
	t.OutBlock.Whchange = e.GetFieldData("t1511OutBlock", "whchange", 0)
	t.OutBlock.Whjday = e.GetFieldData("t1511OutBlock", "whjday", 0)
	t.OutBlock.Wljisu = e.GetFieldData("t1511OutBlock", "wljisu", 0)
	t.OutBlock.Wlchange = e.GetFieldData("t1511OutBlock", "wlchange", 0)
	t.OutBlock.Wljday = e.GetFieldData("t1511OutBlock", "wljday", 0)
	t.OutBlock.Yhjisu = e.GetFieldData("t1511OutBlock", "yhjisu", 0)
	t.OutBlock.Yhchange = e.GetFieldData("t1511OutBlock", "yhchange", 0)
	t.OutBlock.Yhjday = e.GetFieldData("t1511OutBlock", "yhjday", 0)
	t.OutBlock.Yljisu = e.GetFieldData("t1511OutBlock", "yljisu", 0)
	t.OutBlock.Ylchange = e.GetFieldData("t1511OutBlock", "ylchange", 0)
	t.OutBlock.Yljday = e.GetFieldData("t1511OutBlock", "yljday", 0)
	t.OutBlock.Firstjcode = e.GetFieldData("t1511OutBlock", "firstjcode", 0)
	t.OutBlock.Firstjname = e.GetFieldData("t1511OutBlock", "firstjname", 0)
	t.OutBlock.Firstjisu = e.GetFieldData("t1511OutBlock", "firstjisu", 0)
	t.OutBlock.Firsign = e.GetFieldData("t1511OutBlock", "firsign", 0)
	t.OutBlock.Firchange = e.GetFieldData("t1511OutBlock", "firchange", 0)
	t.OutBlock.Firdiff = e.GetFieldData("t1511OutBlock", "firdiff", 0)
	t.OutBlock.Secondjcode = e.GetFieldData("t1511OutBlock", "secondjcode", 0)
	t.OutBlock.Secondjname = e.GetFieldData("t1511OutBlock", "secondjname", 0)
	t.OutBlock.Secondjisu = e.GetFieldData("t1511OutBlock", "secondjisu", 0)
	t.OutBlock.Secsign = e.GetFieldData("t1511OutBlock", "secsign", 0)
	t.OutBlock.Secchange = e.GetFieldData("t1511OutBlock", "secchange", 0)
	t.OutBlock.Secdiff = e.GetFieldData("t1511OutBlock", "secdiff", 0)
	t.OutBlock.Thirdjcode = e.GetFieldData("t1511OutBlock", "thirdjcode", 0)
	t.OutBlock.Thirdjname = e.GetFieldData("t1511OutBlock", "thirdjname", 0)
	t.OutBlock.Thirdjisu = e.GetFieldData("t1511OutBlock", "thirdjisu", 0)
	t.OutBlock.Thrsign = e.GetFieldData("t1511OutBlock", "thrsign", 0)
	t.OutBlock.Thrchange = e.GetFieldData("t1511OutBlock", "thrchange", 0)
	t.OutBlock.Thrdiff = e.GetFieldData("t1511OutBlock", "thrdiff", 0)
	t.OutBlock.Fourthjcode = e.GetFieldData("t1511OutBlock", "fourthjcode", 0)
	t.OutBlock.Fourthjname = e.GetFieldData("t1511OutBlock", "fourthjname", 0)
	t.OutBlock.Fourthjisu = e.GetFieldData("t1511OutBlock", "fourthjisu", 0)
	t.OutBlock.Forsign = e.GetFieldData("t1511OutBlock", "forsign", 0)
	t.OutBlock.Forchange = e.GetFieldData("t1511OutBlock", "forchange", 0)
	t.OutBlock.Fordiff = e.GetFieldData("t1511OutBlock", "fordiff", 0)
	t.OutBlock.Highjo = e.GetFieldData("t1511OutBlock", "highjo", 0)
	t.OutBlock.Upjo = e.GetFieldData("t1511OutBlock", "upjo", 0)
	t.OutBlock.Unchgjo = e.GetFieldData("t1511OutBlock", "unchgjo", 0)
	t.OutBlock.Lowjo = e.GetFieldData("t1511OutBlock", "lowjo", 0)
	t.OutBlock.Downjo = e.GetFieldData("t1511OutBlock", "downjo", 0)

	t.ReceiveDataChan <- x
}

func (t T1511) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T1511) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T1511) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
