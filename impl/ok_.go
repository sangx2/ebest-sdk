package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// OK KOSDAQ거래원
type OK struct {
	InBlock res.OKInBlock

	OutBlock res.OKOutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewOK() *OK {
	return &OK{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (o OK) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return o.ReceiveRealDataChan
}

func (o OK) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return o.ReceiveLinkDataChan
}

func (o *OK) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "OK_.res")

	if i, ok := inBlock.(res.OKInBlock); !ok {
		return errors.New("Invalid inBlock")
	} else {
		o.InBlock = i
	}

	e.SetFieldData("InBlock", "shcode", 0, o.InBlock.Shcode)

	return nil
}

func (o OK) GetOutBlock() interface{} {
	return o.OutBlock
}

func (o OK) GetBlockData(e *wrapper.EBestWrapper, blockName string) string {
	return e.GetBlockData(blockName)
}

func (o *OK) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	o.OutBlock.Offerno1 = e.GetFieldData("OutBlock", "offerno1", 0)
	o.OutBlock.Bidno1 = e.GetFieldData("OutBlock", "bidno1", 0)
	o.OutBlock.Offertrad1 = e.GetFieldData("OutBlock", "offertrad1", 0)
	o.OutBlock.Bidtrad1 = e.GetFieldData("OutBlock", "bidtrad1", 0)
	o.OutBlock.Tradmdvol1 = e.GetFieldData("OutBlock", "tradmdvol1", 0)
	o.OutBlock.Tradmsvol1 = e.GetFieldData("OutBlock", "tradmsvol1", 0)
	o.OutBlock.Tradmdrate1 = e.GetFieldData("OutBlock", "tradmdrate1", 0)
	o.OutBlock.Tradmsrate1 = e.GetFieldData("OutBlock", "tradmsrate1", 0)
	o.OutBlock.Tradmdcha1 = e.GetFieldData("OutBlock", "tradmdcha1", 0)
	o.OutBlock.Tradmscha1 = e.GetFieldData("OutBlock", "tradmscha1", 0)

	o.OutBlock.Offerno2 = e.GetFieldData("OutBlock", "offerno2", 0)
	o.OutBlock.Bidno2 = e.GetFieldData("OutBlock", "bidno2", 0)
	o.OutBlock.Offertrad2 = e.GetFieldData("OutBlock", "offertrad2", 0)
	o.OutBlock.Bidtrad2 = e.GetFieldData("OutBlock", "bidtrad2", 0)
	o.OutBlock.Tradmdvol2 = e.GetFieldData("OutBlock", "tradmdvol2", 0)
	o.OutBlock.Tradmsvol2 = e.GetFieldData("OutBlock", "tradmsvol2", 0)
	o.OutBlock.Tradmdrate2 = e.GetFieldData("OutBlock", "tradmdrate2", 0)
	o.OutBlock.Tradmsrate2 = e.GetFieldData("OutBlock", "tradmsrate2", 0)
	o.OutBlock.Tradmdcha2 = e.GetFieldData("OutBlock", "tradmdcha2", 0)
	o.OutBlock.Tradmscha2 = e.GetFieldData("OutBlock", "tradmscha2", 0)

	o.OutBlock.Offerno3 = e.GetFieldData("OutBlock", "offerno3", 0)
	o.OutBlock.Bidno3 = e.GetFieldData("OutBlock", "bidno3", 0)
	o.OutBlock.Offertrad3 = e.GetFieldData("OutBlock", "offertrad3", 0)
	o.OutBlock.Bidtrad3 = e.GetFieldData("OutBlock", "bidtrad3", 0)
	o.OutBlock.Tradmdvol3 = e.GetFieldData("OutBlock", "tradmdvol3", 0)
	o.OutBlock.Tradmsvol3 = e.GetFieldData("OutBlock", "tradmsvol3", 0)
	o.OutBlock.Tradmdrate3 = e.GetFieldData("OutBlock", "tradmdrate3", 0)
	o.OutBlock.Tradmsrate3 = e.GetFieldData("OutBlock", "tradmsrate3", 0)
	o.OutBlock.Tradmdcha3 = e.GetFieldData("OutBlock", "tradmdcha3", 0)
	o.OutBlock.Tradmscha3 = e.GetFieldData("OutBlock", "tradmscha3", 0)

	o.OutBlock.Offerno4 = e.GetFieldData("OutBlock", "offerno4", 0)
	o.OutBlock.Bidno4 = e.GetFieldData("OutBlock", "bidno4", 0)
	o.OutBlock.Offertrad4 = e.GetFieldData("OutBlock", "offertrad4", 0)
	o.OutBlock.Bidtrad4 = e.GetFieldData("OutBlock", "bidtrad4", 0)
	o.OutBlock.Tradmdvol4 = e.GetFieldData("OutBlock", "tradmdvol4", 0)
	o.OutBlock.Tradmsvol4 = e.GetFieldData("OutBlock", "tradmsvol4", 0)
	o.OutBlock.Tradmdrate4 = e.GetFieldData("OutBlock", "tradmdrate4", 0)
	o.OutBlock.Tradmsrate4 = e.GetFieldData("OutBlock", "tradmsrate4", 0)
	o.OutBlock.Tradmdcha4 = e.GetFieldData("OutBlock", "tradmdcha4", 0)
	o.OutBlock.Tradmscha4 = e.GetFieldData("OutBlock", "tradmscha4", 0)

	o.OutBlock.Offerno5 = e.GetFieldData("OutBlock", "offerno5", 0)
	o.OutBlock.Bidno5 = e.GetFieldData("OutBlock", "bidno5", 0)
	o.OutBlock.Offertrad5 = e.GetFieldData("OutBlock", "offertrad5", 0)
	o.OutBlock.Bidtrad5 = e.GetFieldData("OutBlock", "bidtrad5", 0)
	o.OutBlock.Tradmdvol5 = e.GetFieldData("OutBlock", "tradmdvol5", 0)
	o.OutBlock.Tradmsvol5 = e.GetFieldData("OutBlock", "tradmsvol5", 0)
	o.OutBlock.Tradmdrate5 = e.GetFieldData("OutBlock", "tradmdrate5", 0)
	o.OutBlock.Tradmsrate5 = e.GetFieldData("OutBlock", "tradmsrate5", 0)
	o.OutBlock.Tradmdcha5 = e.GetFieldData("OutBlock", "tradmdcha5", 0)
	o.OutBlock.Tradmscha5 = e.GetFieldData("OutBlock", "tradmscha5", 0)

	o.OutBlock.Ftradmdvol = e.GetFieldData("OutBlock", "ftradmdvol", 0)
	o.OutBlock.Ftradmsvol = e.GetFieldData("OutBlock", "ftradmsvol", 0)
	o.OutBlock.Ftradmdrate = e.GetFieldData("OutBlock", "ftradmdrate", 0)
	o.OutBlock.Ftradmsrate = e.GetFieldData("OutBlock", "ftradmsrate", 0)
	o.OutBlock.Ftradmdcha = e.GetFieldData("OutBlock", "ftradmdcha", 0)
	o.OutBlock.Ftradmscha = e.GetFieldData("OutBlock", "ftradmscha", 0)

	o.OutBlock.Shcode = e.GetFieldData("OutBlock", "shcode", 0)

	o.OutBlock.Tradmdval1 = e.GetFieldData("OutBlock", "tradmdval1", 0)
	o.OutBlock.Tradmsval1 = e.GetFieldData("OutBlock", "tradmsval1", 0)
	o.OutBlock.Tradmdavg1 = e.GetFieldData("OutBlock", "tradmdavg1", 0)
	o.OutBlock.Tradmsavg1 = e.GetFieldData("OutBlock", "tradmsavg1", 0)

	o.OutBlock.Tradmdval2 = e.GetFieldData("OutBlock", "tradmdval2", 0)
	o.OutBlock.Tradmsval2 = e.GetFieldData("OutBlock", "tradmsval2", 0)
	o.OutBlock.Tradmdavg2 = e.GetFieldData("OutBlock", "tradmdavg2", 0)
	o.OutBlock.Tradmsavg2 = e.GetFieldData("OutBlock", "tradmsavg2", 0)

	o.OutBlock.Tradmdval3 = e.GetFieldData("OutBlock", "tradmdval3", 0)
	o.OutBlock.Tradmsval3 = e.GetFieldData("OutBlock", "tradmsval3", 0)
	o.OutBlock.Tradmdavg3 = e.GetFieldData("OutBlock", "tradmdavg3", 0)
	o.OutBlock.Tradmsavg3 = e.GetFieldData("OutBlock", "tradmsavg3", 0)

	o.OutBlock.Tradmdval4 = e.GetFieldData("OutBlock", "tradmdval4", 0)
	o.OutBlock.Tradmsval4 = e.GetFieldData("OutBlock", "tradmsval4", 0)
	o.OutBlock.Tradmdavg4 = e.GetFieldData("OutBlock", "tradmdavg4", 0)
	o.OutBlock.Tradmsavg4 = e.GetFieldData("OutBlock", "tradmsavg4", 0)

	o.OutBlock.Tradmdval5 = e.GetFieldData("OutBlock", "tradmdval5", 0)
	o.OutBlock.Tradmsval5 = e.GetFieldData("OutBlock", "tradmsval5", 0)
	o.OutBlock.Tradmdavg5 = e.GetFieldData("OutBlock", "tradmdavg5", 0)
	o.OutBlock.Tradmsavg5 = e.GetFieldData("OutBlock", "tradmsavg5", 0)

	o.OutBlock.Ftradmdval = e.GetFieldData("OutBlock", "ftradmdval", 0)
	o.OutBlock.Ftradmsval = e.GetFieldData("OutBlock", "ftradmsval", 0)
	o.OutBlock.Ftradmdavg = e.GetFieldData("OutBlock", "ftradmdavg", 0)
	o.OutBlock.Ftradmsavg = e.GetFieldData("OutBlock", "ftradmsavg", 0)

	o.ReceiveRealDataChan <- x
}

func (o OK) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	o.ReceiveLinkDataChan <- x
}
