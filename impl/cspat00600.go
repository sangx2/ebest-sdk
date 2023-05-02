package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// CSPAT00600 현물 정상주문
type CSPAT00600 struct {
	InBlock1 res.CSPAT00600InBlock1

	OutBlock1 res.CSPAT00600OutBlock1
	OutBlock2 res.CSPAT00600OutBlock2

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewCSPAT00600() *CSPAT00600 {
	return &CSPAT00600{
		TPS: 30, LPP: -1,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (c CSPAT00600) GetTPS() int {
	return c.TPS
}

func (c CSPAT00600) GetLPP() int {
	return c.LPP
}

func (c CSPAT00600) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return c.ReceiveDataChan
}

func (c CSPAT00600) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return c.ReceiveMessageChan
}

func (c CSPAT00600) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return c.ReceiveChartRealDataChan
}

func (c CSPAT00600) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return c.ReceiveChartSearchRealDataChan
}

func (c *CSPAT00600) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "CSPAT00600.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.CSPAT00600InBlock1); !ok {
		return errors.New("invalid inBlock1")
	} else {
		c.InBlock1 = i
	}

	e.SetFieldData("CSPAT00600InBlock1", "AcntNo", 0, c.InBlock1.AcntNo)
	e.SetFieldData("CSPAT00600InBlock1", "InptPwd", 0, c.InBlock1.InptPwd)
	e.SetFieldData("CSPAT00600InBlock1", "IsuNo", 0, c.InBlock1.IsuNo)
	e.SetFieldData("CSPAT00600InBlock1", "OrdQty", 0, c.InBlock1.OrdQty)
	e.SetFieldData("CSPAT00600InBlock1", "OrdPrc", 0, c.InBlock1.OrdPrc)
	e.SetFieldData("CSPAT00600InBlock1", "BnsTpCode", 0, c.InBlock1.BnsTpCode)
	e.SetFieldData("CSPAT00600InBlock1", "OrdprcPtnCode", 0, c.InBlock1.OrdprcPtnCode)
	e.SetFieldData("CSPAT00600InBlock1", "MgntrnCode", 0, c.InBlock1.MgntrnCode)
	e.SetFieldData("CSPAT00600InBlock1", "LoanDt", 0, c.InBlock1.LoanDt)
	e.SetFieldData("CSPAT00600InBlock1", "OrdCndiTpCode", 0, c.InBlock1.OrdCndiTpCode)

	return nil
}

func (c CSPAT00600) GetOutBlocks() []interface{} {
	return []interface{}{c.OutBlock1, c.OutBlock2}
}

func (c CSPAT00600) GetBlockData(e *wrapper.EBestWrapper, blockName string) string {
	return e.GetBlockData(blockName)
}

func (c *CSPAT00600) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	c.OutBlock1.RecCnt = e.GetFieldData("CSPAT00600OutBlock1", "RecCnt", 0)
	c.OutBlock1.AcntNo = e.GetFieldData("CSPAT00600OutBlock1", "AcntNo", 0)
	c.OutBlock1.InptPwd = e.GetFieldData("CSPAT00600OutBlock1", "InptPwd", 0)
	c.OutBlock1.IsuNo = e.GetFieldData("CSPAT00600OutBlock1", "IsuNo", 0)
	c.OutBlock1.OrdQty = e.GetFieldData("CSPAT00600OutBlock1", "OrdQty", 0)
	c.OutBlock1.OrdPrc = e.GetFieldData("CSPAT00600OutBlock1", "OrdPrc", 0)
	c.OutBlock1.BnsTpCode = e.GetFieldData("CSPAT00600OutBlock1", "BnsTpCode", 0)
	c.OutBlock1.OrdprcPtnCode = e.GetFieldData("CSPAT00600OutBlock1", "OrdprcPtnCode", 0)
	c.OutBlock1.PrgmOrdprcPtnCode = e.GetFieldData("CSPAT00600OutBlock1", "PrgmOrdprcPtnCode", 0)
	c.OutBlock1.StslAbleYn = e.GetFieldData("CSPAT00600OutBlock1", "StslAbleYn", 0)
	c.OutBlock1.StslOrdprcTpCode = e.GetFieldData("CSPAT00600OutBlock1", "StslOrdprcTpCode", 0)
	c.OutBlock1.CommdaCode = e.GetFieldData("CSPAT00600OutBlock1", "CommdaCode", 0)
	c.OutBlock1.MgntrnCode = e.GetFieldData("CSPAT00600OutBlock1", "MgntrnCode", 0)
	c.OutBlock1.LoanDt = e.GetFieldData("CSPAT00600OutBlock1", "LoanDt", 0)
	c.OutBlock1.MbrNo = e.GetFieldData("CSPAT00600OutBlock1", "MbrNo", 0)
	c.OutBlock1.OrdCndiTpCode = e.GetFieldData("CSPAT00600OutBlock1", "OrdCndiTpCode", 0)
	c.OutBlock1.StrtgCode = e.GetFieldData("CSPAT00600OutBlock1", "StrtgCode", 0)
	c.OutBlock1.GrpID = e.GetFieldData("CSPAT00600OutBlock1", "GrpId", 0)
	c.OutBlock1.OrdSeqNo = e.GetFieldData("CSPAT00600OutBlock1", "OrdSeqNo", 0)
	c.OutBlock1.PtflNo = e.GetFieldData("CSPAT00600OutBlock1", "PtflNo", 0)
	c.OutBlock1.BskNo = e.GetFieldData("CSPAT00600OutBlock1", "BskNo", 0)
	c.OutBlock1.TrchNo = e.GetFieldData("CSPAT00600OutBlock1", "TrchNo", 0)
	c.OutBlock1.ItemNo = e.GetFieldData("CSPAT00600OutBlock1", "ItemNo", 0)
	c.OutBlock1.OpDrtnNo = e.GetFieldData("CSPAT00600OutBlock1", "OpDrtnNo", 0)
	c.OutBlock1.LpYn = e.GetFieldData("CSPAT00600OutBlock1", "LpYn", 0)
	c.OutBlock1.CvrgTpCode = e.GetFieldData("CSPAT00600OutBlock1", "CvrgTpCode", 0)

	c.OutBlock2.RecCnt = e.GetFieldData("CSPAT00600OutBlock2", "RecCnt", 0)
	c.OutBlock2.OrdNo = e.GetFieldData("CSPAT00600OutBlock2", "OrdNo", 0)
	c.OutBlock2.OrdTime = e.GetFieldData("CSPAT00600OutBlock2", "OrdTime", 0)
	c.OutBlock2.OrdMktCode = e.GetFieldData("CSPAT00600OutBlock2", "OrdMktCode", 0)
	c.OutBlock2.OrdPtnCode = e.GetFieldData("CSPAT00600OutBlock2", "OrdPtnCode", 0)
	c.OutBlock2.ShtnIsuNo = e.GetFieldData("CSPAT00600OutBlock2", "ShtnIsuNo", 0)
	c.OutBlock2.MgempNo = e.GetFieldData("CSPAT00600OutBlock2", "MgempNo", 0)
	c.OutBlock2.OrdAmt = e.GetFieldData("CSPAT00600OutBlock2", "OrdAmt", 0)
	c.OutBlock2.SpareOrdNo = e.GetFieldData("CSPAT00600OutBlock2", "SpareOrdNo", 0)
	c.OutBlock2.CvrgSeqno = e.GetFieldData("CSPAT00600OutBlock2", "CvrgSeqno", 0)
	c.OutBlock2.RsvOrdNo = e.GetFieldData("CSPAT00600OutBlock2", "RsvOrdNo", 0)
	c.OutBlock2.SpotOrdQty = e.GetFieldData("CSPAT00600OutBlock2", "SpotOrdQty", 0)
	c.OutBlock2.RuseOrdQty = e.GetFieldData("CSPAT00600OutBlock2", "RuseOrdQty", 0)
	c.OutBlock2.MnyOrdAmt = e.GetFieldData("CSPAT00600OutBlock2", "MnyOrdAmt", 0)
	c.OutBlock2.SubstOrdAmt = e.GetFieldData("CSPAT00600OutBlock2", "SubstOrdAmt", 0)
	c.OutBlock2.RuseOrdAmt = e.GetFieldData("CSPAT00600OutBlock2", "RuseOrdAmt", 0)
	c.OutBlock2.AcntNm = e.GetFieldData("CSPAT00600OutBlock2", "AcntNm", 0)
	c.OutBlock2.IsuNm = e.GetFieldData("CSPAT00600OutBlock2", "IsuNm", 0)

	c.ReceiveDataChan <- x
}

func (c CSPAT00600) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	c.ReceiveMessageChan <- x
}

func (c CSPAT00600) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	c.ReceiveChartRealDataChan <- x
}

func (c CSPAT00600) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	c.ReceiveChartSearchRealDataChan <- x
}
