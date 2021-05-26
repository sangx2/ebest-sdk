package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// S3 KOSPI 체결
type S3 struct {
	InBlock res.S3InBlock

	OutBlock res.S3OutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewS3() *S3 {
	return &S3{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (s S3) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return s.ReceiveRealDataChan
}

func (s S3) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return s.ReceiveLinkDataChan
}

func (s *S3) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "S3_.res")

	if i, ok := inBlock.(res.S3InBlock); !ok {
		return errors.New("Invalid inBlock")
	} else {
		s.InBlock = i
	}

	e.SetFieldData("InBlock", "shcode", 0, s.InBlock.Shcode)

	return nil
}

func (s S3) GetOutBlock() interface{} {
	return s.OutBlock
}

func (s *S3) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	s.OutBlock.Chetime = e.GetFieldData("OutBlock", "chetime", 0)
	s.OutBlock.Sign = e.GetFieldData("OutBlock", "sign", 0)
	s.OutBlock.Change = e.GetFieldData("OutBlock", "change", 0)
	s.OutBlock.Drate = e.GetFieldData("OutBlock", "drate", 0)
	s.OutBlock.Price = e.GetFieldData("OutBlock", "price", 0)
	s.OutBlock.Opentime = e.GetFieldData("OutBlock", "opentime", 0)
	s.OutBlock.Open = e.GetFieldData("OutBlock", "open", 0)
	s.OutBlock.Hightime = e.GetFieldData("OutBlock", "hightime", 0)
	s.OutBlock.High = e.GetFieldData("OutBlock", "high", 0)
	s.OutBlock.Lowtime = e.GetFieldData("OutBlock", "lowtime", 0)
	s.OutBlock.Low = e.GetFieldData("OutBlock", "low", 0)
	s.OutBlock.Cgubun = e.GetFieldData("OutBlock", "cgubun", 0)
	s.OutBlock.Cvolume = e.GetFieldData("OutBlock", "cvolume", 0)
	s.OutBlock.Volume = e.GetFieldData("OutBlock", "volume", 0)
	s.OutBlock.Value = e.GetFieldData("OutBlock", "value", 0)
	s.OutBlock.Mdvolume = e.GetFieldData("OutBlock", "mdvolume", 0)
	s.OutBlock.Mdchecnt = e.GetFieldData("OutBlock", "mdchecnt", 0)
	s.OutBlock.Msvolume = e.GetFieldData("OutBlock", "msvolume", 0)
	s.OutBlock.Mschecnt = e.GetFieldData("OutBlock", "mschecnt", 0)
	s.OutBlock.Cpower = e.GetFieldData("OutBlock", "cpower", 0)
	s.OutBlock.Wavrg = e.GetFieldData("OutBlock", "w_avrg", 0)
	s.OutBlock.Offerho = e.GetFieldData("OutBlock", "offerho", 0)
	s.OutBlock.Bidho = e.GetFieldData("OutBlock", "bidho", 0)
	s.OutBlock.Status = e.GetFieldData("OutBlock", "status", 0)
	s.OutBlock.Jnilvolume = e.GetFieldData("OutBlock", "jnilvolume", 0)
	s.OutBlock.Shcode = e.GetFieldData("OutBlock", "shcode", 0)

	s.ReceiveRealDataChan <- x
}

func (s S3) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	s.ReceiveLinkDataChan <- x
}
