package impl

import (
	"errors"
	"github.com/sangx2/ebestsdk/res"
	"github.com/sangx2/ebestsdk/wrapper"
)

// HA KOSDAQ호가잔량
type HA struct {
	InBlock res.HAInBlock

	OutBlock res.HAOutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewHA() *HA {
	return &HA{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (h HA) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return h.ReceiveRealDataChan
}

func (h HA) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return h.ReceiveLinkDataChan
}

func (h *HA) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "HA_.res")

	if i, ok := inBlock.(res.HAInBlock); !ok {
		return errors.New("Invalid inBlock")
	} else {
		h.InBlock = i
	}

	e.SetFieldData("InBlock", "shcode", 0, h.InBlock.Shcode)

	return nil
}

func (h HA) GetOutBlock() interface{} {
	return h.OutBlock
}

func (h *HA) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	h.OutBlock.Hotime = e.GetFieldData("OutBlock", "hotime", 0)
	h.OutBlock.Offerho1 = e.GetFieldData("OutBlock", "offerho1", 0)
	h.OutBlock.Offerrem1 = e.GetFieldData("OutBlock", "offerrem1", 0)
	h.OutBlock.Bidho1 = e.GetFieldData("OutBlock", "bidho1", 0)
	h.OutBlock.Bidrem1 = e.GetFieldData("OutBlock", "bidrem1", 0)
	h.OutBlock.Offerho2 = e.GetFieldData("OutBlock", "offerho2", 0)
	h.OutBlock.Offerrem2 = e.GetFieldData("OutBlock", "offerrem2", 0)
	h.OutBlock.Bidho2 = e.GetFieldData("OutBlock", "bidho2", 0)
	h.OutBlock.Bidrem2 = e.GetFieldData("OutBlock", "bidrem2", 0)
	h.OutBlock.Offerho3 = e.GetFieldData("OutBlock", "offerho3", 0)
	h.OutBlock.Offerrem3 = e.GetFieldData("OutBlock", "offerrem3", 0)
	h.OutBlock.Bidho3 = e.GetFieldData("OutBlock", "bidho3", 0)
	h.OutBlock.Bidrem3 = e.GetFieldData("OutBlock", "bidrem3", 0)
	h.OutBlock.Offerho4 = e.GetFieldData("OutBlock", "offerho4", 0)
	h.OutBlock.Offerrem4 = e.GetFieldData("OutBlock", "offerrem4", 0)
	h.OutBlock.Bidho4 = e.GetFieldData("OutBlock", "bidho4", 0)
	h.OutBlock.Bidrem4 = e.GetFieldData("OutBlock", "bidrem4", 0)
	h.OutBlock.Offerho5 = e.GetFieldData("OutBlock", "offerho5", 0)
	h.OutBlock.Offerrem5 = e.GetFieldData("OutBlock", "offerrem5", 0)
	h.OutBlock.Bidho5 = e.GetFieldData("OutBlock", "bidho5", 0)
	h.OutBlock.Bidrem5 = e.GetFieldData("OutBlock", "bidrem5", 0)
	h.OutBlock.Offerho6 = e.GetFieldData("OutBlock", "offerho6", 0)
	h.OutBlock.Offerrem6 = e.GetFieldData("OutBlock", "offerrem6", 0)
	h.OutBlock.Bidho6 = e.GetFieldData("OutBlock", "bidho6", 0)
	h.OutBlock.Bidrem6 = e.GetFieldData("OutBlock", "bidrem6", 0)
	h.OutBlock.Offerho7 = e.GetFieldData("OutBlock", "offerho7", 0)
	h.OutBlock.Offerrem7 = e.GetFieldData("OutBlock", "offerrem7", 0)
	h.OutBlock.Bidho7 = e.GetFieldData("OutBlock", "bidho7", 0)
	h.OutBlock.Bidrem7 = e.GetFieldData("OutBlock", "bidrem7", 0)
	h.OutBlock.Offerho8 = e.GetFieldData("OutBlock", "offerho8", 0)
	h.OutBlock.Offerrem8 = e.GetFieldData("OutBlock", "offerrem8", 0)
	h.OutBlock.Bidho8 = e.GetFieldData("OutBlock", "bidho8", 0)
	h.OutBlock.Bidrem8 = e.GetFieldData("OutBlock", "bidrem8", 0)
	h.OutBlock.Offerho9 = e.GetFieldData("OutBlock", "offerho9", 0)
	h.OutBlock.Offerrem9 = e.GetFieldData("OutBlock", "offerrem9", 0)
	h.OutBlock.Bidho9 = e.GetFieldData("OutBlock", "bidho9", 0)
	h.OutBlock.Bidrem9 = e.GetFieldData("OutBlock", "bidrem9", 0)
	h.OutBlock.Offerho10 = e.GetFieldData("OutBlock", "offerho10", 0)
	h.OutBlock.Offerrem10 = e.GetFieldData("OutBlock", "offerrem10", 0)
	h.OutBlock.Bidho10 = e.GetFieldData("OutBlock", "bidho10", 0)
	h.OutBlock.Bidrem10 = e.GetFieldData("OutBlock", "bidrem10", 0)
	h.OutBlock.Totofferrem = e.GetFieldData("OutBlock", "totofferrem", 0)
	h.OutBlock.Totbidrem = e.GetFieldData("OutBlock", "totbidrem", 0)
	h.OutBlock.Donsigubun = e.GetFieldData("OutBlock", "donsigubun", 0)
	h.OutBlock.Shcode = e.GetFieldData("OutBlock", "shcode", 0)
	h.OutBlock.Allocgubun = e.GetFieldData("OutBlock", "alloc_gubun", 0)

	h.ReceiveRealDataChan <- x
}

func (h HA) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	h.ReceiveLinkDataChan <- x
}
