package impl

import (
	"errors"

	"github.com/sangx2/ebestsdk/res"
	"github.com/sangx2/ebestsdk/wrapper"
)

// CSPAQ12200 현물계좌 예수금/주문가능금액/총평가 조회
type CSPAQ12200 struct {
	InBlock1 res.CSPAQ12200InBlock1

	OutBlock1 res.CSPAQ12200OutBlock1
	OutBlock2 res.CSPAQ12200OutBlock2

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewCSPAQ12200() *CSPAQ12200 {
	return &CSPAQ12200{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (c CSPAQ12200) GetTPS() int {
	return c.TPS
}

func (c CSPAQ12200) GetLPP() int {
	return c.LPP
}

func (c CSPAQ12200) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return c.ReceiveDataChan
}

func (c CSPAQ12200) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return c.ReceiveMessageChan
}

func (c CSPAQ12200) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return c.ReceiveChartRealDataChan
}

func (c CSPAQ12200) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return c.ReceiveChartSearchRealDataChan
}

func (c *CSPAQ12200) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "CSPAQ12200.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.CSPAQ12200InBlock1); !ok {
		return errors.New("invalid inBlock")
	} else {
		c.InBlock1 = i
	}

	e.SetFieldData("CSPAQ12200InBlock1", "RecCnt", 0, c.InBlock1.RecCnt)
	e.SetFieldData("CSPAQ12200InBlock1", "MgmtBrnNo", 0, c.InBlock1.MgmtBrnNo)
	e.SetFieldData("CSPAQ12200InBlock1", "AcntNo", 0, c.InBlock1.AcntNo)
	e.SetFieldData("CSPAQ12200InBlock1", "Pwd", 0, c.InBlock1.Pwd)
	e.SetFieldData("CSPAQ12200InBlock1", "BalCreTp", 0, c.InBlock1.BalCreTp)

	return nil
}

func (c CSPAQ12200) GetOutBlocks() []interface{} {
	return []interface{}{c.OutBlock1, c.OutBlock2}
}

func (c *CSPAQ12200) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	c.OutBlock1.RecCnt = e.GetFieldData("CSPAQ12200OutBlock1", "RecCnt", 0)
	c.OutBlock1.MgmtBrnNo = e.GetFieldData("CSPAQ12200OutBlock1", "MgmtBrnNo", 0)
	c.OutBlock1.AcntNo = e.GetFieldData("CSPAQ12200OutBlock1", "AcntNo", 0)
	c.OutBlock1.Pwd = e.GetFieldData("CSPAQ12200OutBlock1", "Pwd", 0)
	c.OutBlock1.BalCreTp = e.GetFieldData("CSPAQ12200OutBlock1", "BalCreTp", 0)

	c.OutBlock2.RecCnt = e.GetFieldData("CSPAQ12200OutBlock2", "RecCnt", 0)
	c.OutBlock2.BrnNm = e.GetFieldData("CSPAQ12200OutBlock2", "BrnNm", 0)
	c.OutBlock2.AcntNm = e.GetFieldData("CSPAQ12200OutBlock2", "AcntNm", 0)
	c.OutBlock2.MnyOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MnyOrdAbleAmt", 0)
	c.OutBlock2.MnyoutAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MnyoutAbleAmt", 0)
	c.OutBlock2.SeOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "SeOrdAbleAmt", 0)
	c.OutBlock2.KdqOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "KdqOrdAbleAmt", 0)
	c.OutBlock2.BalEvalAmt = e.GetFieldData("CSPAQ12200OutBlock2", "BalEvalAmt", 0)
	c.OutBlock2.RcvblAmt = e.GetFieldData("CSPAQ12200OutBlock2", "RcvblAmt", 0)
	c.OutBlock2.DpsastTotamt = e.GetFieldData("CSPAQ12200OutBlock2", "DpsastTotamt", 0)
	c.OutBlock2.PnlRat = e.GetFieldData("CSPAQ12200OutBlock2", "PnlRat", 0)
	c.OutBlock2.InvstOrgAmt = e.GetFieldData("CSPAQ12200OutBlock2", "InvstOrgAmt", 0)
	c.OutBlock2.InvstPlAmt = e.GetFieldData("CSPAQ12200OutBlock2", "InvstPlAmt", 0)
	c.OutBlock2.CrdtPldgOrdAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtPldgOrdAmt", 0)
	c.OutBlock2.Dps = e.GetFieldData("CSPAQ12200OutBlock2", "Dps", 0)
	c.OutBlock2.SubstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "SubstAmt", 0)
	c.OutBlock2.D1Dps = e.GetFieldData("CSPAQ12200OutBlock2", "D1Dps", 0)
	c.OutBlock2.D2Dps = e.GetFieldData("CSPAQ12200OutBlock2", "D2Dps", 0)
	c.OutBlock2.MnyrclAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MnyrclAmt", 0)
	c.OutBlock2.MgnMny = e.GetFieldData("CSPAQ12200OutBlock2", "MgnMny", 0)
	c.OutBlock2.MgnSubst = e.GetFieldData("CSPAQ12200OutBlock2", "MgnSubst", 0)
	c.OutBlock2.ChckAmt = e.GetFieldData("CSPAQ12200OutBlock2", "ChckAmt", 0)
	c.OutBlock2.SubstOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "SubstOrdAbleAmt", 0)
	c.OutBlock2.MgnRat100pctOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MgnRat100pctOrdAbleAmt", 0)
	c.OutBlock2.MgnRat35ordAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MgnRat35ordAbleAmt", 0)
	c.OutBlock2.MgnRat50ordAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MgnRat50ordAbleAmt", 0)
	c.OutBlock2.PrdaySellAdjstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "PrdaySellAdjstAmt", 0)
	c.OutBlock2.PrdayBuyAdjstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "PrdayBuyAdjstAmt", 0)
	c.OutBlock2.CrdaySellAdjstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdaySellAdjstAmt", 0)
	c.OutBlock2.CrdayBuyAdjstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdayBuyAdjstAmt", 0)
	c.OutBlock2.D1ovdRepayRqrdAmt = e.GetFieldData("CSPAQ12200OutBlock2", "D1ovdRepayRqrdAmt", 0)
	c.OutBlock2.D2ovdRepayRqrdAmt = e.GetFieldData("CSPAQ12200OutBlock2", "D2ovdRepayRqrdAmt", 0)
	c.OutBlock2.D1PrsmptWthdwAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "D1PrsmptWthdwAbleAmt", 0)
	c.OutBlock2.D2PrsmptWthdwAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "D2PrsmptWthdwAbleAmt", 0)
	c.OutBlock2.DpspdgLoanAmt = e.GetFieldData("CSPAQ12200OutBlock2", "DpspdgLoanAmt", 0)
	c.OutBlock2.Imreq = e.GetFieldData("CSPAQ12200OutBlock2", "Imreq", 0)
	c.OutBlock2.MloanAmt = e.GetFieldData("CSPAQ12200OutBlock2", "MloanAmt", 0)
	c.OutBlock2.ChgAfPldgRat = e.GetFieldData("CSPAQ12200OutBlock2", "ChgAfPldgRat", 0)
	c.OutBlock2.OrgPldgAmt = e.GetFieldData("CSPAQ12200OutBlock2", "OrgPldgAmt", 0)
	c.OutBlock2.SubPldgAmt = e.GetFieldData("CSPAQ12200OutBlock2", "SubPldgAmt", 0)
	c.OutBlock2.RqrdPldgAmt = e.GetFieldData("CSPAQ12200OutBlock2", "RqrdPldgAmt", 0)
	c.OutBlock2.OrgPdlckAmt = e.GetFieldData("CSPAQ12200OutBlock2", "OrgPdlckAmt", 0)
	c.OutBlock2.PdlckAmt = e.GetFieldData("CSPAQ12200OutBlock2", "PdlckAmt", 0)
	c.OutBlock2.AddPldgMny = e.GetFieldData("CSPAQ12200OutBlock2", "AddPldgMny", 0)
	c.OutBlock2.D1OrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "D1OrdAbleAmt", 0)
	c.OutBlock2.CrdtIntdltAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtIntdltAmt", 0)
	c.OutBlock2.EtclndAmt = e.GetFieldData("CSPAQ12200OutBlock2", "EtclndAmt", 0)
	c.OutBlock2.NtdayPrsmptCvrgAmt = e.GetFieldData("CSPAQ12200OutBlock2", "NtdayPrsmptCvrgAmt", 0)
	c.OutBlock2.OrgPldgSumAmt = e.GetFieldData("CSPAQ12200OutBlock2", "OrgPldgSumAmt", 0)
	c.OutBlock2.CrdtOrdAbleAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtOrdAbleAmt", 0)
	c.OutBlock2.SubPldgSumAmt = e.GetFieldData("CSPAQ12200OutBlock2", "SubPldgSumAmt", 0)
	c.OutBlock2.CrdtPldgAmtMny = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtPldgAmtMny", 0)
	c.OutBlock2.CrdtPldgSubstAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtPldgSubstAmt", 0)
	c.OutBlock2.AddCrdtPldgMny = e.GetFieldData("CSPAQ12200OutBlock2", "AddCrdtPldgMny", 0)
	c.OutBlock2.CrdtPldgRuseAmt = e.GetFieldData("CSPAQ12200OutBlock2", "CrdtPldgRuseAmt", 0)
	c.OutBlock2.AddCrdtPldgSubst = e.GetFieldData("CSPAQ12200OutBlock2", "AddCrdtPldgSubst", 0)
	c.OutBlock2.CslLoanAmtdt1 = e.GetFieldData("CSPAQ12200OutBlock2", "CslLoanAmtdt1", 0)
	c.OutBlock2.DpslRestrcAmt = e.GetFieldData("CSPAQ12200OutBlock2", "DpslRestrcAmt", 0)

	c.ReceiveDataChan <- x
}

func (c CSPAQ12200) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	c.ReceiveMessageChan <- x
}

func (c CSPAQ12200) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	c.ReceiveChartRealDataChan <- x
}

func (c CSPAQ12200) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	c.ReceiveChartSearchRealDataChan <- x
}
