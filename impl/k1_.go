package impl

import (
	"errors"
	"github.com/sangx2/ebestsdk/res"
	"github.com/sangx2/ebestsdk/wrapper"
)

// K1 KOSPI거래원
type K1 struct {
	InBlock res.K1InBlock

	OutBlock res.K1OutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewK1() *K1 {
	return &K1{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (k1 K1) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return k1.ReceiveRealDataChan
}

func (k1 K1) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return k1.ReceiveLinkDataChan
}

func (k1 *K1) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "K1_.res")

	if i, ok := inBlock.(res.K1InBlock); !ok {
		return errors.New("Invalid inBlock")
	} else {
		k1.InBlock = i
	}

	e.SetFieldData("InBlock", "shcode", 0, k1.InBlock.Shcode)

	return nil
}

func (k1 K1) GetOutBlock() interface{} {
	return k1.OutBlock
}

func (k1 *K1) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	k1.OutBlock.Offerno1 = e.GetFieldData("OutBlock", "offerno1", 0)
	k1.OutBlock.Bidno1 = e.GetFieldData("OutBlock", "bidno1", 0)
	k1.OutBlock.Offertrad1 = e.GetFieldData("OutBlock", "offertrad1", 0)
	k1.OutBlock.Bidtrad1 = e.GetFieldData("OutBlock", "bidtrad1", 0)
	k1.OutBlock.Tradmdvol1 = e.GetFieldData("OutBlock", "tradmdvol1", 0)
	k1.OutBlock.Tradmsvol1 = e.GetFieldData("OutBlock", "tradmsvol1", 0)
	k1.OutBlock.Tradmdrate1 = e.GetFieldData("OutBlock", "tradmdrate1", 0)
	k1.OutBlock.Tradmsrate1 = e.GetFieldData("OutBlock", "tradmsrate1", 0)
	k1.OutBlock.Tradmdcha1 = e.GetFieldData("OutBlock", "tradmdcha1", 0)
	k1.OutBlock.Tradmscha1 = e.GetFieldData("OutBlock", "tradmscha1", 0)

	k1.OutBlock.Offerno2 = e.GetFieldData("OutBlock", "offerno2", 0)
	k1.OutBlock.Bidno2 = e.GetFieldData("OutBlock", "bidno2", 0)
	k1.OutBlock.Offertrad2 = e.GetFieldData("OutBlock", "offertrad2", 0)
	k1.OutBlock.Bidtrad2 = e.GetFieldData("OutBlock", "bidtrad2", 0)
	k1.OutBlock.Tradmdvol2 = e.GetFieldData("OutBlock", "tradmdvol2", 0)
	k1.OutBlock.Tradmsvol2 = e.GetFieldData("OutBlock", "tradmsvol2", 0)
	k1.OutBlock.Tradmdrate2 = e.GetFieldData("OutBlock", "tradmdrate2", 0)
	k1.OutBlock.Tradmsrate2 = e.GetFieldData("OutBlock", "tradmsrate2", 0)
	k1.OutBlock.Tradmdcha2 = e.GetFieldData("OutBlock", "tradmdcha2", 0)
	k1.OutBlock.Tradmscha2 = e.GetFieldData("OutBlock", "tradmscha2", 0)

	k1.OutBlock.Offerno3 = e.GetFieldData("OutBlock", "offerno3", 0)
	k1.OutBlock.Bidno3 = e.GetFieldData("OutBlock", "bidno3", 0)
	k1.OutBlock.Offertrad3 = e.GetFieldData("OutBlock", "offertrad3", 0)
	k1.OutBlock.Bidtrad3 = e.GetFieldData("OutBlock", "bidtrad3", 0)
	k1.OutBlock.Tradmdvol3 = e.GetFieldData("OutBlock", "tradmdvol3", 0)
	k1.OutBlock.Tradmsvol3 = e.GetFieldData("OutBlock", "tradmsvol3", 0)
	k1.OutBlock.Tradmdrate3 = e.GetFieldData("OutBlock", "tradmdrate3", 0)
	k1.OutBlock.Tradmsrate3 = e.GetFieldData("OutBlock", "tradmsrate3", 0)
	k1.OutBlock.Tradmdcha3 = e.GetFieldData("OutBlock", "tradmdcha3", 0)
	k1.OutBlock.Tradmscha3 = e.GetFieldData("OutBlock", "tradmscha3", 0)

	k1.OutBlock.Offerno4 = e.GetFieldData("OutBlock", "offerno4", 0)
	k1.OutBlock.Bidno4 = e.GetFieldData("OutBlock", "bidno4", 0)
	k1.OutBlock.Offertrad4 = e.GetFieldData("OutBlock", "offertrad4", 0)
	k1.OutBlock.Bidtrad4 = e.GetFieldData("OutBlock", "bidtrad4", 0)
	k1.OutBlock.Tradmdvol4 = e.GetFieldData("OutBlock", "tradmdvol4", 0)
	k1.OutBlock.Tradmsvol4 = e.GetFieldData("OutBlock", "tradmsvol4", 0)
	k1.OutBlock.Tradmdrate4 = e.GetFieldData("OutBlock", "tradmdrate4", 0)
	k1.OutBlock.Tradmsrate4 = e.GetFieldData("OutBlock", "tradmsrate4", 0)
	k1.OutBlock.Tradmdcha4 = e.GetFieldData("OutBlock", "tradmdcha4", 0)
	k1.OutBlock.Tradmscha4 = e.GetFieldData("OutBlock", "tradmscha4", 0)

	k1.OutBlock.Offerno5 = e.GetFieldData("OutBlock", "offerno5", 0)
	k1.OutBlock.Bidno5 = e.GetFieldData("OutBlock", "bidno5", 0)
	k1.OutBlock.Offertrad5 = e.GetFieldData("OutBlock", "offertrad5", 0)
	k1.OutBlock.Bidtrad5 = e.GetFieldData("OutBlock", "bidtrad5", 0)
	k1.OutBlock.Tradmdvol5 = e.GetFieldData("OutBlock", "tradmdvol5", 0)
	k1.OutBlock.Tradmsvol5 = e.GetFieldData("OutBlock", "tradmsvol5", 0)
	k1.OutBlock.Tradmdrate5 = e.GetFieldData("OutBlock", "tradmdrate5", 0)
	k1.OutBlock.Tradmsrate5 = e.GetFieldData("OutBlock", "tradmsrate5", 0)
	k1.OutBlock.Tradmdcha5 = e.GetFieldData("OutBlock", "tradmdcha5", 0)
	k1.OutBlock.Tradmscha5 = e.GetFieldData("OutBlock", "tradmscha5", 0)

	k1.OutBlock.Ftradmdvol = e.GetFieldData("OutBlock", "ftradmdvol", 0)
	k1.OutBlock.Ftradmsvol = e.GetFieldData("OutBlock", "ftradmsvol", 0)
	k1.OutBlock.Ftradmdrate = e.GetFieldData("OutBlock", "ftradmdrate", 0)
	k1.OutBlock.Ftradmsrate = e.GetFieldData("OutBlock", "ftradmsrate", 0)
	k1.OutBlock.Ftradmdcha = e.GetFieldData("OutBlock", "ftradmdcha", 0)
	k1.OutBlock.Ftradmscha = e.GetFieldData("OutBlock", "ftradmscha", 0)

	k1.OutBlock.Shcode = e.GetFieldData("OutBlock", "shcode", 0)

	k1.OutBlock.Tradmdval1 = e.GetFieldData("OutBlock", "tradmdval1", 0)
	k1.OutBlock.Tradmsval1 = e.GetFieldData("OutBlock", "tradmsval1", 0)
	k1.OutBlock.Tradmdavg1 = e.GetFieldData("OutBlock", "tradmdavg1", 0)
	k1.OutBlock.Tradmsavg1 = e.GetFieldData("OutBlock", "tradmsavg1", 0)

	k1.OutBlock.Tradmdval2 = e.GetFieldData("OutBlock", "tradmdval2", 0)
	k1.OutBlock.Tradmsval2 = e.GetFieldData("OutBlock", "tradmsval2", 0)
	k1.OutBlock.Tradmdavg2 = e.GetFieldData("OutBlock", "tradmdavg2", 0)
	k1.OutBlock.Tradmsavg2 = e.GetFieldData("OutBlock", "tradmsavg2", 0)

	k1.OutBlock.Tradmdval3 = e.GetFieldData("OutBlock", "tradmdval3", 0)
	k1.OutBlock.Tradmsval3 = e.GetFieldData("OutBlock", "tradmsval3", 0)
	k1.OutBlock.Tradmdavg3 = e.GetFieldData("OutBlock", "tradmdavg3", 0)
	k1.OutBlock.Tradmsavg3 = e.GetFieldData("OutBlock", "tradmsavg3", 0)

	k1.OutBlock.Tradmdval4 = e.GetFieldData("OutBlock", "tradmdval4", 0)
	k1.OutBlock.Tradmsval4 = e.GetFieldData("OutBlock", "tradmsval4", 0)
	k1.OutBlock.Tradmdavg4 = e.GetFieldData("OutBlock", "tradmdavg4", 0)
	k1.OutBlock.Tradmsavg4 = e.GetFieldData("OutBlock", "tradmsavg4", 0)

	k1.OutBlock.Tradmdval5 = e.GetFieldData("OutBlock", "tradmdval5", 0)
	k1.OutBlock.Tradmsval5 = e.GetFieldData("OutBlock", "tradmsval5", 0)
	k1.OutBlock.Tradmdavg5 = e.GetFieldData("OutBlock", "tradmdavg5", 0)
	k1.OutBlock.Tradmsavg5 = e.GetFieldData("OutBlock", "tradmsavg5", 0)

	k1.OutBlock.Ftradmdval = e.GetFieldData("OutBlock", "ftradmdval", 0)
	k1.OutBlock.Ftradmsval = e.GetFieldData("OutBlock", "ftradmsval", 0)
	k1.OutBlock.Ftradmdavg = e.GetFieldData("OutBlock", "ftradmdavg", 0)
	k1.OutBlock.Ftradmsavg = e.GetFieldData("OutBlock", "ftradmsavg", 0)

	k1.ReceiveRealDataChan <- x
}

func (k1 K1) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	k1.ReceiveLinkDataChan <- x
}
