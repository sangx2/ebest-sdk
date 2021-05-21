package impl

import (
	"errors"
	"github.com/sangx2/ebestsdk/res"
	"github.com/sangx2/ebestsdk/wrapper"
)

// NWS 실시간 뉴스 제목 패킷
type NWS struct {
	InBlock res.NWSInBlock

	OutBlock res.NWSOutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewNWS() *NWS {
	return &NWS{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (n NWS) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return n.ReceiveRealDataChan
}

func (n NWS) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return n.ReceiveLinkDataChan
}

func (n *NWS) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
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

func (n *NWS) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	n.OutBlock.Date = e.GetFieldData("OutBlock", "date", 0)
	n.OutBlock.Time = e.GetFieldData("OutBlock", "time", 0)
	n.OutBlock.Publisher = e.GetFieldData("OutBlock", "id", 0)
	n.OutBlock.Realkey = e.GetFieldData("OutBlock", "realkey", 0)
	n.OutBlock.Title = e.GetFieldData("OutBlock", "title", 0)
	n.OutBlock.Code = e.GetFieldData("OutBlock", "code", 0)
	n.OutBlock.Bodysize = e.GetFieldData("OutBlock", "bodysize", 0)

	n.ReceiveRealDataChan <- x
}

func (n NWS) ReceivedLinkData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	n.ReceiveLinkDataChan <- x
}
