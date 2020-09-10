package impl

import (
	"errors"
	"github.com/sangx2/ebest/res"
	"github.com/sangx2/ebest/wrapper"
)

// CSPAT00800 현물 취소주문
type CSPAT00800 struct {
	InBlock1 res.CSPAT00800InBlock1

	OutBlock1 res.CSPAT00800OutBlock1
	OutBlock2 res.CSPAT00800OutBlock2

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewCSPAT00800() *CSPAT00800 {
	return &CSPAT00800{
		TPS: 30, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (c CSPAT00800) GetTPS() int {
	return c.TPS
}

func (c CSPAT00800) GetLPP() int {
	return c.LPP
}

func (c CSPAT00800) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return c.ReceiveDataChan
}

func (c CSPAT00800) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return c.ReceiveMessageChan
}

func (c CSPAT00800) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return c.ReceiveChartRealDataChan
}

func (c CSPAT00800) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return c.ReceiveChartSearchRealDataChan
}

func (c *CSPAT00800) SetFieldData(e *wrapper.Ebest, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "CSPAT00800.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.CSPAT00800InBlock1); !ok {
		return errors.New("invalid inBlock1")
	} else {
		c.InBlock1 = i
	}

	e.SetFieldData("CSPAT00800InBlock1", "OrgOrdNo", 0, c.InBlock1.OrgOrdNo)
	e.SetFieldData("CSPAT00800InBlock1", "AcntNo", 0, c.InBlock1.AcntNo)
	e.SetFieldData("CSPAT00800InBlock1", "InptPwd", 0, c.InBlock1.InptPwd)
	e.SetFieldData("CSPAT00800InBlock1", "IsuNo", 0, c.InBlock1.IsuNo)
	e.SetFieldData("CSPAT00800InBlock1", "OrdQty", 0, c.InBlock1.OrdQty)

	return nil
}

func (c CSPAT00800) GetOutBlocks() []interface{} {
	return []interface{}{c.OutBlock1, c.OutBlock2}
}

func (c *CSPAT00800) ReceivedData(e *wrapper.Ebest, x wrapper.XaQueryReceiveData) {
	c.OutBlock1.RecCnt = e.GetFieldData("CSPAT00800OutBlock1", "RecCnt", 0)
	c.OutBlock1.OrgOrdNo = e.GetFieldData("CSPAT00800OutBlock1", "OrgOrdNo", 0)
	c.OutBlock1.AcntNo = e.GetFieldData("CSPAT00800OutBlock1", "AcntNo", 0)
	c.OutBlock1.InptPwd = e.GetFieldData("CSPAT00800OutBlock1", "InptPwd", 0)
	c.OutBlock1.IsuNo = e.GetFieldData("CSPAT00800OutBlock1", "IsuNo", 0)
	c.OutBlock1.OrdQty = e.GetFieldData("CSPAT00800OutBlock1", "OrdQty", 0)
	c.OutBlock1.CommdaCode = e.GetFieldData("CSPAT00800OutBlock1", "CommdaCode", 0)
	c.OutBlock1.GrpID = e.GetFieldData("CSPAT00800OutBlock1", "GrpId", 0)
	c.OutBlock1.StrtgCode = e.GetFieldData("CSPAT00800OutBlock1", "StrtgCode", 0)
	c.OutBlock1.OrdSeqNo = e.GetFieldData("CSPAT00800OutBlock1", "OrdSeqNo", 0)
	c.OutBlock1.PtflNo = e.GetFieldData("CSPAT00800OutBlock1", "PtflNo", 0)
	c.OutBlock1.BskNo = e.GetFieldData("CSPAT00800OutBlock1", "BskNo", 0)
	c.OutBlock1.TrchNo = e.GetFieldData("CSPAT00800OutBlock1", "TrchNo", 0)
	c.OutBlock1.ItemNo = e.GetFieldData("CSPAT00800OutBlock1", "ItemNo", 0)

	c.OutBlock2.RecCnt = e.GetFieldData("CSPAT00800OutBlock2", "RecCnt", 0)
	c.OutBlock2.OrdNo = e.GetFieldData("CSPAT00800OutBlock2", "OrdNo", 0)
	c.OutBlock2.PrntOrdNo = e.GetFieldData("CSPAT00800OutBlock2", "PrntOrdNo", 0)
	c.OutBlock2.OrdTime = e.GetFieldData("CSPAT00800OutBlock2", "OrdTime", 0)
	c.OutBlock2.OrdMktCode = e.GetFieldData("CSPAT00800OutBlock2", "OrdMktCode", 0)
	c.OutBlock2.OrdPtnCode = e.GetFieldData("CSPAT00800OutBlock2", "OrdPtnCode", 0)
	c.OutBlock2.ShtnIsuNo = e.GetFieldData("CSPAT00800OutBlock2", "ShtnIsuNo", 0)
	c.OutBlock2.PrgmOrdprcPtnCode = e.GetFieldData("CSPAT00800OutBlock2", "PrgmOrdprcPtnCode", 0)
	c.OutBlock2.StslOrdprcTpCode = e.GetFieldData("CSPAT00800OutBlock2", "StslOrdprcTpCode", 0)
	c.OutBlock2.StslAbleYn = e.GetFieldData("CSPAT00800OutBlock2", "StslAbleYn", 0)
	c.OutBlock2.MgntrnCode = e.GetFieldData("CSPAT00800OutBlock2", "MgntrnCode", 0)
	c.OutBlock2.LoanDt = e.GetFieldData("CSPAT00800OutBlock2", "LoanDt", 0)
	c.OutBlock2.CvrgOrdTp = e.GetFieldData("CSPAT00800OutBlock2", "CvrgOrdTp", 0)
	c.OutBlock2.LpYn = e.GetFieldData("CSPAT00800OutBlock2", "LpYn", 0)
	c.OutBlock2.MgempNo = e.GetFieldData("CSPAT00800OutBlock2", "MgempNo", 0)
	c.OutBlock2.BnsTpCode = e.GetFieldData("CSPAT00800OutBlock2", "BnsTpCode", 0)
	c.OutBlock2.SpareOrdNo = e.GetFieldData("CSPAT00800OutBlock2", "SpareOrdNo", 0)
	c.OutBlock2.CvrgSeqno = e.GetFieldData("CSPAT00800OutBlock2", "CvrgSeqno", 0)
	c.OutBlock2.RsvOrdNo = e.GetFieldData("CSPAT00800OutBlock2", "RsvOrdNo", 0)
	c.OutBlock2.AcntNm = e.GetFieldData("CSPAT00800OutBlock2", "AcntNm", 0)
	c.OutBlock2.IsuNm = e.GetFieldData("CSPAT00800OutBlock2", "IsuNm", 0)

	c.ReceiveDataChan <- x
}

func (c CSPAT00800) ReceivedMessage(e *wrapper.Ebest, x wrapper.XaQueryReceiveMessage) {
	c.ReceiveMessageChan <- x
}

func (c CSPAT00800) ReceivedChartRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveChartRealData) {
	c.ReceiveChartRealDataChan <- x
}

func (c CSPAT00800) ReceivedSearchRealData(e *wrapper.Ebest, x wrapper.XaQueryReceiveSearchRealData) {
	c.ReceiveChartSearchRealDataChan <- x
}
