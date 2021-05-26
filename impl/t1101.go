package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// T1101 주식 현재가 호가 조회
type T1101 struct {
	InBlock res.T1101InBlock

	OutBlock res.T1101OutBlock

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewT1101() *T1101 {
	return &T1101{
		TPS: 5, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (t T1101) GetTPS() int {
	return t.TPS
}

func (t T1101) GetLPP() int {
	return t.LPP
}

func (t T1101) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return t.ReceiveDataChan
}

func (t T1101) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return t.ReceiveMessageChan
}

func (t T1101) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return t.ReceiveChartRealDataChan
}

func (t T1101) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return t.ReceiveChartSearchRealDataChan
}

func (t *T1101) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "t1101.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.T1101InBlock); !ok {
		return errors.New("invalid inBlock1")
	} else {
		t.InBlock = i
	}

	e.SetFieldData("t1101InBlock", "shcode", 0, t.InBlock.Shcode)

	return nil
}

func (t T1101) GetOutBlocks() []interface{} {
	return []interface{}{t.OutBlock}
}

func (t *T1101) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	t.OutBlock.Hname = e.GetFieldData("t1101OutBlock", "Hname", 0)
	t.OutBlock.Price = e.GetFieldData("t1101OutBlock", "price", 0)
	t.OutBlock.Sign = e.GetFieldData("t1101OutBlock", "sign", 0)
	t.OutBlock.Change = e.GetFieldData("t1101OutBlock", "change", 0)
	t.OutBlock.Diff = e.GetFieldData("t1101OutBlock", "diff", 0)
	t.OutBlock.Volume = e.GetFieldData("t1101OutBlock", "volume", 0)
	t.OutBlock.Jnilclose = e.GetFieldData("t1101OutBlock", "jnilclose", 0)
	t.OutBlock.Offerho1 = e.GetFieldData("t1101OutBlock", "offerho1", 0)
	t.OutBlock.Bidho1 = e.GetFieldData("t1101OutBlock", "bidho1", 0)
	t.OutBlock.Offerrem1 = e.GetFieldData("t1101OutBlock", "offerrem1", 0)
	t.OutBlock.Bidrem1 = e.GetFieldData("t1101OutBlock", "bidrem1", 0)
	t.OutBlock.Preoffercha1 = e.GetFieldData("t1101OutBlock", "preoffercha1", 0)
	t.OutBlock.Prebidcha1 = e.GetFieldData("t1101OutBlock", "prebidcha1", 0)
	t.OutBlock.Offerho2 = e.GetFieldData("t1101OutBlock", "offerho2", 0)
	t.OutBlock.Bidho2 = e.GetFieldData("t1101OutBlock", "bidho2", 0)
	t.OutBlock.Offerrem2 = e.GetFieldData("t1101OutBlock", "offerrem2", 0)
	t.OutBlock.Bidrem2 = e.GetFieldData("t1101OutBlock", "bidrem2", 0)
	t.OutBlock.Preoffercha2 = e.GetFieldData("t1101OutBlock", "preoffercha2", 0)
	t.OutBlock.Prebidcha2 = e.GetFieldData("t1101OutBlock", "prebidcha2", 0)
	t.OutBlock.Offerho3 = e.GetFieldData("t1101OutBlock", "offerho3", 0)
	t.OutBlock.Bidho3 = e.GetFieldData("t1101OutBlock", "bidho3", 0)
	t.OutBlock.Offerrem3 = e.GetFieldData("t1101OutBlock", "offerrem3", 0)
	t.OutBlock.Bidrem3 = e.GetFieldData("t1101OutBlock", "bidrem3", 0)
	t.OutBlock.Preoffercha3 = e.GetFieldData("t1101OutBlock", "preoffercha3", 0)
	t.OutBlock.Prebidcha3 = e.GetFieldData("t1101OutBlock", "prebidcha3", 0)
	t.OutBlock.Offerho4 = e.GetFieldData("t1101OutBlock", "offerho4", 0)
	t.OutBlock.Bidho4 = e.GetFieldData("t1101OutBlock", "bidho4", 0)
	t.OutBlock.Offerrem4 = e.GetFieldData("t1101OutBlock", "offerrem4", 0)
	t.OutBlock.Bidrem4 = e.GetFieldData("t1101OutBlock", "bidrem4", 0)
	t.OutBlock.Preoffercha4 = e.GetFieldData("t1101OutBlock", "preoffercha4", 0)
	t.OutBlock.Prebidcha4 = e.GetFieldData("t1101OutBlock", "prebidcha4", 0)
	t.OutBlock.Offerho5 = e.GetFieldData("t1101OutBlock", "offerho5", 0)
	t.OutBlock.Bidho5 = e.GetFieldData("t1101OutBlock", "bidho5", 0)
	t.OutBlock.Offerrem5 = e.GetFieldData("t1101OutBlock", "offerrem5", 0)
	t.OutBlock.Bidrem5 = e.GetFieldData("t1101OutBlock", "bidrem5", 0)
	t.OutBlock.Preoffercha5 = e.GetFieldData("t1101OutBlock", "preoffercha5", 0)
	t.OutBlock.Prebidcha5 = e.GetFieldData("t1101OutBlock", "prebidcha5", 0)
	t.OutBlock.Offerho6 = e.GetFieldData("t1101OutBlock", "offerho6", 0)
	t.OutBlock.Bidho6 = e.GetFieldData("t1101OutBlock", "bidho6", 0)
	t.OutBlock.Offerrem6 = e.GetFieldData("t1101OutBlock", "offerrem6", 0)
	t.OutBlock.Bidrem6 = e.GetFieldData("t1101OutBlock", "bidrem6", 0)
	t.OutBlock.Preoffercha6 = e.GetFieldData("t1101OutBlock", "preoffercha6", 0)
	t.OutBlock.Prebidcha6 = e.GetFieldData("t1101OutBlock", "prebidcha6", 0)
	t.OutBlock.Offerho7 = e.GetFieldData("t1101OutBlock", "offerho7", 0)
	t.OutBlock.Bidho7 = e.GetFieldData("t1101OutBlock", "bidho7", 0)
	t.OutBlock.Offerrem7 = e.GetFieldData("t1101OutBlock", "offerrem7", 0)
	t.OutBlock.Bidrem7 = e.GetFieldData("t1101OutBlock", "bidrem7", 0)
	t.OutBlock.Preoffercha7 = e.GetFieldData("t1101OutBlock", "preoffercha7", 0)
	t.OutBlock.Prebidcha7 = e.GetFieldData("t1101OutBlock", "prebidcha7", 0)
	t.OutBlock.Offerho8 = e.GetFieldData("t1101OutBlock", "offerho8", 0)
	t.OutBlock.Bidho8 = e.GetFieldData("t1101OutBlock", "bidho8", 0)
	t.OutBlock.Offerrem8 = e.GetFieldData("t1101OutBlock", "offerrem8", 0)
	t.OutBlock.Bidrem8 = e.GetFieldData("t1101OutBlock", "bidrem8", 0)
	t.OutBlock.Preoffercha8 = e.GetFieldData("t1101OutBlock", "preoffercha8", 0)
	t.OutBlock.Prebidcha8 = e.GetFieldData("t1101OutBlock", "prebidcha8", 0)
	t.OutBlock.Offerho9 = e.GetFieldData("t1101OutBlock", "offerho9", 0)
	t.OutBlock.Bidho9 = e.GetFieldData("t1101OutBlock", "bidho9", 0)
	t.OutBlock.Offerrem9 = e.GetFieldData("t1101OutBlock", "offerrem9", 0)
	t.OutBlock.Bidrem9 = e.GetFieldData("t1101OutBlock", "bidrem9", 0)
	t.OutBlock.Preoffercha9 = e.GetFieldData("t1101OutBlock", "preoffercha9", 0)
	t.OutBlock.Prebidcha9 = e.GetFieldData("t1101OutBlock", "prebidcha9", 0)
	t.OutBlock.Offerho10 = e.GetFieldData("t1101OutBlock", "offerho10", 0)
	t.OutBlock.Bidho10 = e.GetFieldData("t1101OutBlock", "bidho10", 0)
	t.OutBlock.Offerrem10 = e.GetFieldData("t1101OutBlock", "offerrem10", 0)
	t.OutBlock.Bidrem10 = e.GetFieldData("t1101OutBlock", "bidrem10", 0)
	t.OutBlock.Preoffercha10 = e.GetFieldData("t1101OutBlock", "preoffercha10", 0)
	t.OutBlock.Prebidcha10 = e.GetFieldData("t1101OutBlock", "prebidcha10", 0)
	t.OutBlock.Offer = e.GetFieldData("t1101OutBlock", "offer", 0)
	t.OutBlock.Bid = e.GetFieldData("t1101OutBlock", "bid", 0)
	t.OutBlock.Preoffercha = e.GetFieldData("t1101OutBlock", "preoffercha", 0)
	t.OutBlock.Prebidcha = e.GetFieldData("t1101OutBlock", "prebidcha", 0)
	t.OutBlock.Hotime = e.GetFieldData("t1101OutBlock", "hotime", 0)
	t.OutBlock.Yeprice = e.GetFieldData("t1101OutBlock", "yeprice", 0)
	t.OutBlock.Yevolume = e.GetFieldData("t1101OutBlock", "yevolume", 0)
	t.OutBlock.Yesign = e.GetFieldData("t1101OutBlock", "yesign", 0)
	t.OutBlock.Yechange = e.GetFieldData("t1101OutBlock", "yechange", 0)
	t.OutBlock.Yediff = e.GetFieldData("t1101OutBlock", "yediff", 0)
	t.OutBlock.Tmoffer = e.GetFieldData("t1101OutBlock", "tmoffer", 0)
	t.OutBlock.Tmbid = e.GetFieldData("t1101OutBlock", "tmbid", 0)
	t.OutBlock.Ho_status = e.GetFieldData("t1101OutBlock", "Ho_status", 0)
	t.OutBlock.Shcode = e.GetFieldData("t1101OutBlock", "shcode", 0)
	t.OutBlock.Uplmtprice = e.GetFieldData("t1101OutBlock", "uplmtprice", 0)
	t.OutBlock.Dnlmtprice = e.GetFieldData("t1101OutBlock", "dnlmtprice", 0)
	t.OutBlock.Open = e.GetFieldData("t1101OutBlock", "open", 0)
	t.OutBlock.High = e.GetFieldData("t1101OutBlock", "high", 0)
	t.OutBlock.Low = e.GetFieldData("t1101OutBlock", "low", 0)

	t.ReceiveDataChan <- x
}

func (t T1101) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	t.ReceiveMessageChan <- x
}

func (t T1101) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	t.ReceiveChartRealDataChan <- x
}

func (t T1101) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	t.ReceiveChartSearchRealDataChan <- x
}
