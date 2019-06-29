package callback

import (
	"errors"

	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/wrapper"
)

// K3 KOSDAQ 체결
type K3 struct {
	InBlock  model.K3InBlock
	OutBlock model.K3OutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealRecieveLinkData
}

func NewK3() *K3 {
	return &K3{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealRecieveLinkData, 1),
	}
}

func (k K3) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return k.ReceiveRealDataChan
}

func (k K3) GetReceivedLinkDataChan() chan wrapper.XaRealRecieveLinkData {
	return k.ReceiveLinkDataChan
}

func (k *K3) SetFieldData(e *wrapper.Ebest, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "K3_.res")

	i, ok := inBlock.(model.K3InBlock)
	if !ok {
		return errors.New("Invalid inBlock")
	}
	k.InBlock = i

	e.SetFieldData("InBlock", "shcode", 0, k.InBlock.Shcode)

	return nil
}

func (k K3) GetOutBlock() interface{} {
	return k.OutBlock
}

func (k *K3) ReceivedRealData(e *wrapper.Ebest, x wrapper.XaRealReceiveRealData) {
	k.OutBlock.Chetime = e.GetFieldData("OutBlock", "chetime", 0)
	k.OutBlock.Sign = e.GetFieldData("OutBlock", "sign", 0)
	k.OutBlock.Change = e.GetFieldData("OutBlock", "change", 0)
	k.OutBlock.Drate = e.GetFieldData("OutBlock", "drate", 0)
	k.OutBlock.Price = e.GetFieldData("OutBlock", "price", 0)
	k.OutBlock.Opentime = e.GetFieldData("OutBlock", "opentime", 0)
	k.OutBlock.Open = e.GetFieldData("OutBlock", "open", 0)
	k.OutBlock.Hightime = e.GetFieldData("OutBlock", "hightime", 0)
	k.OutBlock.High = e.GetFieldData("OutBlock", "high", 0)
	k.OutBlock.Lowtime = e.GetFieldData("OutBlock", "lowtime", 0)
	k.OutBlock.Low = e.GetFieldData("OutBlock", "low", 0)
	k.OutBlock.Cgubun = e.GetFieldData("OutBlock", "cgubun", 0)
	k.OutBlock.Cvolume = e.GetFieldData("OutBlock", "cvolume", 0)
	k.OutBlock.Volume = e.GetFieldData("OutBlock", "volume", 0)
	k.OutBlock.Value = e.GetFieldData("OutBlock", "value", 0)
	k.OutBlock.Mdvolume = e.GetFieldData("OutBlock", "mdvolume", 0)
	k.OutBlock.Mdchecnt = e.GetFieldData("OutBlock", "mdchecnt", 0)
	k.OutBlock.Msvolume = e.GetFieldData("OutBlock", "msvolume", 0)
	k.OutBlock.Mschecnt = e.GetFieldData("OutBlock", "mschecnt", 0)
	k.OutBlock.Cpower = e.GetFieldData("OutBlock", "cpower", 0)
	k.OutBlock.Wavrg = e.GetFieldData("OutBlock", "w_avrg", 0)
	k.OutBlock.Offerho = e.GetFieldData("OutBlock", "offerho", 0)
	k.OutBlock.Bidho = e.GetFieldData("OutBlock", "bidho", 0)
	k.OutBlock.Status = e.GetFieldData("OutBlock", "status", 0)
	k.OutBlock.Jnilvolume = e.GetFieldData("OutBlock", "jnilvolume", 0)
	k.OutBlock.Shcode = e.GetFieldData("OutBlock", "shcode", 0)

	k.ReceiveRealDataChan <- x
}

func (k K3) ReceivedLinkData(ew *wrapper.Ebest, x wrapper.XaRealRecieveLinkData) {
	k.ReceiveLinkDataChan <- x
}
