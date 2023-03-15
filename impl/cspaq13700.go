package impl

import (
	"errors"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// CSPAQ13700
type CSPAQ13700 struct {
	InBlock1 res.CSPAQ13700InBlock1

	OutBlock1 res.CSPAQ13700OutBlock1
	OutBlock2 res.CSPAQ13700OutBlock2
	OutBlock3 []res.CSPAQ13700OutBlock3

	TPS, LPP int

	ReceiveDataChan                chan wrapper.XaQueryReceiveData
	ReceiveMessageChan             chan wrapper.XaQueryReceiveMessage
	ReceiveChartRealDataChan       chan wrapper.XaQueryReceiveChartRealData
	ReceiveChartSearchRealDataChan chan wrapper.XaQueryReceiveSearchRealData
}

func NewCSPAQ13700() *CSPAQ13700 {
	return &CSPAQ13700{
		TPS: 1, LPP: 200,
		ReceiveDataChan:                make(chan wrapper.XaQueryReceiveData, 1),
		ReceiveMessageChan:             make(chan wrapper.XaQueryReceiveMessage, 1),
		ReceiveChartRealDataChan:       make(chan wrapper.XaQueryReceiveChartRealData, 1),
		ReceiveChartSearchRealDataChan: make(chan wrapper.XaQueryReceiveSearchRealData, 1),
	}
}

func (c CSPAQ13700) GetTPS() int {
	return c.TPS
}

func (c CSPAQ13700) GetLPP() int {
	return c.LPP
}

func (c CSPAQ13700) GetReceiveDataChan() chan wrapper.XaQueryReceiveData {
	return c.ReceiveDataChan
}

func (c CSPAQ13700) GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage {
	return c.ReceiveMessageChan
}

func (c CSPAQ13700) GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData {
	return c.ReceiveChartRealDataChan
}

func (c CSPAQ13700) GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData {
	return c.ReceiveChartSearchRealDataChan
}

func (c *CSPAQ13700) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error {
	e.ResFileName(resPath + "CSPAQ13700.res")

	if len(inBlocks) != 1 {
		return errors.New("invalid inBlocks len")
	}

	if i, ok := inBlocks[0].(res.CSPAQ13700InBlock1); !ok {
		return errors.New("invalid inBlock")
	} else {
		c.InBlock1 = i
	}

	e.SetFieldData("CSPAQ13700InBlock1", "RecCnt", 0, c.InBlock1.RecCnt)
	e.SetFieldData("CSPAQ13700InBlock1", "AcntNo", 0, c.InBlock1.AcntNo)
	e.SetFieldData("CSPAQ13700InBlock1", "InptPwd", 0, c.InBlock1.InptPwd)
	e.SetFieldData("CSPAQ13700InBlock1", "OrdMktCode", 0, c.InBlock1.OrdMktCode)
	e.SetFieldData("CSPAQ13700InBlock1", "BnsTpCode", 0, c.InBlock1.BnsTpCode)
	e.SetFieldData("CSPAQ13700InBlock1", "IsuNo", 0, c.InBlock1.IsuNo)
	e.SetFieldData("CSPAQ13700InBlock1", "ExecYn", 0, c.InBlock1.ExecYn)
	e.SetFieldData("CSPAQ13700InBlock1", "OrdDt", 0, c.InBlock1.OrdDt)
	e.SetFieldData("CSPAQ13700InBlock1", "SrtOrdNo2", 0, c.InBlock1.SrtOrdNo2)
	e.SetFieldData("CSPAQ13700InBlock1", "BkseqTpCode", 0, c.InBlock1.BkseqTpCode)
	e.SetFieldData("CSPAQ13700InBlock1", "OrdPtnCode", 0, c.InBlock1.OrdPtnCode)

	return nil
}

func (c CSPAQ13700) GetOutBlocks() []interface{} {
	return []interface{}{c.OutBlock1, c.OutBlock2, c.OutBlock3}
}

func (c CSPAQ13700) GetBlockDate(e *wrapper.EBestWrapper, blockName string) string {
	return e.GetBlockData(blockName)
}

func (c *CSPAQ13700) ReceivedData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveData) {
	c.OutBlock1.RecCnt = e.GetFieldData("CSPAQ13700OutBlock1", "RecCnt", 0)
	c.OutBlock1.AcntNo = e.GetFieldData("CSPAQ13700OutBlock1", "AcntNo", 0)
	c.OutBlock1.InptPwd = e.GetFieldData("CSPAQ13700OutBlock1", "InptPwd", 0)
	c.OutBlock1.OrdMktCode = e.GetFieldData("CSPAQ13700OutBlock1", "OrdMktCode", 0)
	c.OutBlock1.BnsTpCode = e.GetFieldData("CSPAQ13700OutBlock1", "BnsTpCode", 0)
	c.OutBlock1.IsuNo = e.GetFieldData("CSPAQ13700OutBlock1", "BalCreTp", 0)
	c.OutBlock1.ExecYn = e.GetFieldData("CSPAQ13700OutBlock1", "ExecYn", 0)
	c.OutBlock1.OrdDt = e.GetFieldData("CSPAQ13700OutBlock1", "OrdDt", 0)
	c.OutBlock1.SrtOrdNo2 = e.GetFieldData("CSPAQ13700OutBlock1", "SrtOrdNo2", 0)
	c.OutBlock1.BkseqTpCode = e.GetFieldData("CSPAQ13700OutBlock1", "BkseqTpCode", 0)
	c.OutBlock1.OrdPtnCode = e.GetFieldData("CSPAQ13700OutBlock1", "OrdPtnCode", 0)

	c.OutBlock2.RecCnt = e.GetFieldData("CSPAQ13700OutBlock2", "RecCnt", 0)
	c.OutBlock2.SellExecAmt = e.GetFieldData("CSPAQ13700OutBlock2", "SellExecAmt", 0)
	c.OutBlock2.BuyExecAmt = e.GetFieldData("CSPAQ13700OutBlock2", "BuyExecAmt", 0)
	c.OutBlock2.SellExecQty = e.GetFieldData("CSPAQ13700OutBlock2", "SellExecQty", 0)
	c.OutBlock2.BuyExecQty = e.GetFieldData("CSPAQ13700OutBlock2", "BuyExecQty", 0)
	c.OutBlock2.SellOrdQty = e.GetFieldData("CSPAQ13700OutBlock2", "SellOrdQty", 0)
	c.OutBlock2.BuyOrdQty = e.GetFieldData("CSPAQ13700OutBlock2", "BuyOrdQty", 0)

	for i := 0; i < int(e.GetBlockCount("CSPAQ13700OutBlock3")); i++ {
		tr := res.CSPAQ13700OutBlock3{
			OrdDt:         e.GetFieldData("CSPAQ13700OutBlock3", "OrdDt", i),
			Mgmtbrnno:     e.GetFieldData("CSPAQ13700OutBlock3", "Mgmtbrnno", i),
			Ordmktcode:    e.GetFieldData("CSPAQ13700OutBlock3", "Ordmktcode", i),
			OrdNo:         e.GetFieldData("CSPAQ13700OutBlock3", "OrdNo", i),
			OrgOrdNo:      e.GetFieldData("CSPAQ13700OutBlock3", "OrgOrdNo", i),
			IsuNo:         e.GetFieldData("CSPAQ13700OutBlock3", "IsuNo", i),
			IsuNm:         e.GetFieldData("CSPAQ13700OutBlock3", "IsuNm", i),
			BnsTpCode:     e.GetFieldData("CSPAQ13700OutBlock3", "BnsTpCode", i),
			BnsTpNm:       e.GetFieldData("CSPAQ13700OutBlock3", "BnsTpNm", i),
			OrdPtnCode:    e.GetFieldData("CSPAQ13700OutBlock3", "OrdPtnCode", i),
			OrdPtnNm:      e.GetFieldData("CSPAQ13700OutBlock3", "OrdPtnNm", i),
			OrdTrxPtnCode: e.GetFieldData("CSPAQ13700OutBlock3", "OrdTrxPtnCode", i),
			OrdTrxPtnNm:   e.GetFieldData("CSPAQ13700OutBlock3", "OrdTrxPtnNm", i),
			MrcTpCode:     e.GetFieldData("CSPAQ13700OutBlock3", "MrcTpCode", i),
			MrcTpNm:       e.GetFieldData("CSPAQ13700OutBlock3", "MrcTpNm", i),
			MrcQty:        e.GetFieldData("CSPAQ13700OutBlock3", "MrcQty", i),
			MrcAbleQty:    e.GetFieldData("CSPAQ13700OutBlock3", "MrcAbleQty", i),
			OrdQty:        e.GetFieldData("CSPAQ13700OutBlock3", "OrdQty", i),
			OrdPrc:        e.GetFieldData("CSPAQ13700OutBlock3", "OrdPrc", i),
			ExecQty:       e.GetFieldData("CSPAQ13700OutBlock3", "ExecQty", i),
			ExecPrc:       e.GetFieldData("CSPAQ13700OutBlock3", "ExecPrc", i),
			ExecTrxTime:   e.GetFieldData("CSPAQ13700OutBlock3", "ExecTrxTime", i),
			LastExecTime:  e.GetFieldData("CSPAQ13700OutBlock3", "LastExecTime", i),
			OrdprcPtnCode: e.GetFieldData("CSPAQ13700OutBlock3", "OrdprcPtnCode", i),
			OrdprcPtnNm:   e.GetFieldData("CSPAQ13700OutBlock3", "OrdprcPtnNm", i),
			OrdCndiTpCode: e.GetFieldData("CSPAQ13700OutBlock3", "OrdCndiTpCode", i),
			AllExecQty:    e.GetFieldData("CSPAQ13700OutBlock3", "AllExecQty", i),
			RegCommdaCode: e.GetFieldData("CSPAQ13700OutBlock3", "RegCommdaCode", i),
			CommdaNm:      e.GetFieldData("CSPAQ13700OutBlock3", "CommdaNm", i),
			MbrNo:         e.GetFieldData("CSPAQ13700OutBlock3", "MbrNo", i),
			RsvOrdYn:      e.GetFieldData("CSPAQ13700OutBlock3", "RsvOrdYn", i),
			LoanDt:        e.GetFieldData("CSPAQ13700OutBlock3", "LoanDt", i),
			OrdTime:       e.GetFieldData("CSPAQ13700OutBlock3", "OrdTime", i),
			OpDrtnNo:      e.GetFieldData("CSPAQ13700OutBlock3", "OpDrtnNo", i),
			OdrrId:        e.GetFieldData("CSPAQ13700OutBlock3", "OdrrId", i),
		}
		c.OutBlock3 = append(c.OutBlock3, tr)
	}

	c.ReceiveDataChan <- x
}

func (c CSPAQ13700) ReceivedMessage(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveMessage) {
	c.ReceiveMessageChan <- x
}

func (c CSPAQ13700) ReceivedChartRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveChartRealData) {
	c.ReceiveChartRealDataChan <- x
}

func (c CSPAQ13700) ReceivedSearchRealData(e *wrapper.EBestWrapper, x wrapper.XaQueryReceiveSearchRealData) {
	c.ReceiveChartSearchRealDataChan <- x
}
