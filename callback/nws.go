package callback

import (
	"errors"
	"github.com/sangx2/ebest/res"
	"github.com/sangx2/ebest/wrapper"
)

// NWS 실시간 뉴스 제목 패킷
type NWS struct {
	InBlock res.NWSInBlock

	OutBlock res.NWSOutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealRecieveLinkData
}

func NewNWS() *NWS {
	return &NWS{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealRecieveLinkData, 1),
	}
}

func (n NWS) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return n.ReceiveRealDataChan
}

func (n NWS) GetReceivedLinkDataChan() chan wrapper.XaRealRecieveLinkData {
	return n.ReceiveLinkDataChan
}

func (n *NWS) SetFieldData(e *wrapper.Ebest, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "NWS.res")

	if i, ok := inBlock.(res.NWSInBlock); !ok {
		return errors.New("Invalid inBlock")
	} else {
		n.InBlock = i
	}

	e.SetFieldData("InBlock", "nwcode", 0, n.InBlock.NWcode)

	return nil
}

func (n NWS) GetOutBlock() interface{} {
	return n.OutBlock
}

func (n *NWS) ReceivedRealData(e *wrapper.Ebest, x wrapper.XaRealReceiveRealData) {
	n.OutBlock.Date = e.GetFieldData("OutBlock", "date", 0)
	n.OutBlock.Time = e.GetFieldData("OutBlock", "time", 0)
	n.OutBlock.Publisher = e.GetFieldData("OutBlock", "id", 0)
	n.OutBlock.Realkey = e.GetFieldData("OutBlock", "realkey", 0)
	n.OutBlock.Title = e.GetFieldData("OutBlock", "title", 0)
	n.OutBlock.Code = e.GetFieldData("OutBlock", "code", 0)
	n.OutBlock.Bodysize = e.GetFieldData("OutBlock", "bodysize", 0)

	n.ReceiveRealDataChan <- x
}

func (n NWS) ReceivedLinkData(e *wrapper.Ebest, x wrapper.XaRealRecieveLinkData) {
	n.ReceiveLinkDataChan <- x
}
