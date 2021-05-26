package impl

import (
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// SC3 주식주문취소
type SC3 struct {
	OutBlock res.SC3OutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewSC3() *SC3 {
	return &SC3{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (s SC3) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return s.ReceiveRealDataChan
}

func (s SC3) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return s.ReceiveLinkDataChan
}

func (s *SC3) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "SC3.res")

	return nil
}

func (s SC3) GetOutBlock() interface{} {
	return s.OutBlock
}

func (s *SC3) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
	s.OutBlock.Lineseq = e.GetFieldData("OutBlock", "lineseq", 0)
	s.OutBlock.Accno = e.GetFieldData("OutBlock", "accno", 0)
	s.OutBlock.User = e.GetFieldData("OutBlock", "user", 0)
	s.OutBlock.Len = e.GetFieldData("OutBlock", "len", 0)
	s.OutBlock.Gubun = e.GetFieldData("OutBlock", "gubun", 0)
	s.OutBlock.Compress = e.GetFieldData("OutBlock", "compress", 0)
	s.OutBlock.Encrypt = e.GetFieldData("OutBlock", "encrypt", 0)
	s.OutBlock.Offset = e.GetFieldData("OutBlock", "offset", 0)
	s.OutBlock.Trcode = e.GetFieldData("OutBlock", "trcode", 0)
	s.OutBlock.Compid = e.GetFieldData("OutBlock", "compid", 0)
	s.OutBlock.Userid = e.GetFieldData("OutBlock", "userid", 0)
	s.OutBlock.Media = e.GetFieldData("OutBlock", "media", 0)
	s.OutBlock.Ifid = e.GetFieldData("OutBlock", "ifid", 0)
	s.OutBlock.Seq = e.GetFieldData("OutBlock", "seq", 0)
	s.OutBlock.Trid = e.GetFieldData("OutBlock", "trid", 0)
	s.OutBlock.Pubip = e.GetFieldData("OutBlock", "pubip", 0)
	s.OutBlock.Prvip = e.GetFieldData("OutBlock", "prvip", 0)
	s.OutBlock.Pcbpno = e.GetFieldData("OutBlock", "pcbpno", 0)
	s.OutBlock.Bpno = e.GetFieldData("OutBlock", "bpno", 0)
	s.OutBlock.Termno = e.GetFieldData("OutBlock", "termno", 0)
	s.OutBlock.Lang = e.GetFieldData("OutBlock", "lang", 0)
	s.OutBlock.Proctm = e.GetFieldData("OutBlock", "proctm", 0)
	s.OutBlock.Msgcode = e.GetFieldData("OutBlock", "msgcode", 0)
	s.OutBlock.Outgu = e.GetFieldData("OutBlock", "outgu", 0)
	s.OutBlock.Compreq = e.GetFieldData("OutBlock", "compreq", 0)
	s.OutBlock.Funckey = e.GetFieldData("OutBlock", "funckey", 0)
	s.OutBlock.Reqcnt = e.GetFieldData("OutBlock", "reqcnt", 0)
	s.OutBlock.Filler = e.GetFieldData("OutBlock", "filler", 0)
	s.OutBlock.Cont = e.GetFieldData("OutBlock", "cont", 0)
	s.OutBlock.Contkey = e.GetFieldData("OutBlock", "contkey", 0)
	s.OutBlock.Varlen = e.GetFieldData("OutBlock", "varlen", 0)
	s.OutBlock.Varhdlen = e.GetFieldData("OutBlock", "varhdlen", 0)
	s.OutBlock.Varmsglen = e.GetFieldData("OutBlock", "varmsglen", 0)
	s.OutBlock.Trsrc = e.GetFieldData("OutBlock", "trsrc", 0)
	s.OutBlock.Eventid = e.GetFieldData("OutBlock", "eventid", 0)
	s.OutBlock.Ifinfo = e.GetFieldData("OutBlock", "ifinfo", 0)
	s.OutBlock.Filler1 = e.GetFieldData("OutBlock", "filler1", 0)
	s.OutBlock.Ordxctptncode = e.GetFieldData("OutBlock", "ordxctptncode", 0)
	s.OutBlock.Ordmktcode = e.GetFieldData("OutBlock", "ordmktcode", 0)
	s.OutBlock.Ordptncode = e.GetFieldData("OutBlock", "ordptncode", 0)
	s.OutBlock.Mgmtbrnno = e.GetFieldData("OutBlock", "mgmtbrnno", 0)
	s.OutBlock.Accno1 = e.GetFieldData("OutBlock", "accno1", 0)
	s.OutBlock.Accno2 = e.GetFieldData("OutBlock", "accno2", 0)
	s.OutBlock.Acntnm = e.GetFieldData("OutBlock", "acntnm", 0)
	s.OutBlock.Isuno = e.GetFieldData("OutBlock", "Isuno", 0)
	s.OutBlock.Isunm = e.GetFieldData("OutBlock", "Isunm", 0)
	s.OutBlock.Ordno = e.GetFieldData("OutBlock", "ordno", 0)
	s.OutBlock.Orgordno = e.GetFieldData("OutBlock", "orgordno", 0)
	s.OutBlock.Execno = e.GetFieldData("OutBlock", "execno", 0)
	s.OutBlock.Ordqty = e.GetFieldData("OutBlock", "ordqty", 0)
	s.OutBlock.Ordprc = e.GetFieldData("OutBlock", "ordprc", 0)
	s.OutBlock.Execqty = e.GetFieldData("OutBlock", "execqty", 0)
	s.OutBlock.Execprc = e.GetFieldData("OutBlock", "execprc", 0)
	s.OutBlock.Mdfycnfqty = e.GetFieldData("OutBlock", "mdfycnfqty", 0)
	s.OutBlock.Mdfycnfprc = e.GetFieldData("OutBlock", "mdfycnfprc", 0)
	s.OutBlock.Canccnfqty = e.GetFieldData("OutBlock", "canccnfqty", 0)
	s.OutBlock.Rjtqty = e.GetFieldData("OutBlock", "rjtqty", 0)
	s.OutBlock.Ordtrxptncode = e.GetFieldData("OutBlock", "ordtrxptncode", 0)
	s.OutBlock.Mtiordseqno = e.GetFieldData("OutBlock", "mtiordseqno", 0)
	s.OutBlock.Ordcndi = e.GetFieldData("OutBlock", "ordcndi", 0)
	s.OutBlock.Ordprcptncode = e.GetFieldData("OutBlock", "ordprcptncode", 0)
	s.OutBlock.Nsavtrdqty = e.GetFieldData("OutBlock", "nsavtrdqty", 0)
	s.OutBlock.ShtnIsuno = e.GetFieldData("OutBlock", "shtnIsuno", 0)
	s.OutBlock.Opdrtnno = e.GetFieldData("OutBlock", "opdrtnno", 0)
	s.OutBlock.Cvrgordtp = e.GetFieldData("OutBlock", "cvrgordtp", 0)
	s.OutBlock.Unercqty = e.GetFieldData("OutBlock", "unercqty", 0)
	s.OutBlock.Orgordunercqty = e.GetFieldData("OutBlock", "orgordunercqty", 0)
	s.OutBlock.Orgordmdfyqty = e.GetFieldData("OutBlock", "orgordmdfyqty", 0)
	s.OutBlock.Orgordcancqty = e.GetFieldData("OutBlock", "orgordcancqty", 0)
	s.OutBlock.Ordavrexecprc = e.GetFieldData("OutBlock", "ordavrexecprc", 0)
	s.OutBlock.Ordamt = e.GetFieldData("OutBlock", "ordamt", 0)
	s.OutBlock.StdIsuno = e.GetFieldData("OutBlock", "stdIsuno", 0)
	s.OutBlock.BfstdIsuno = e.GetFieldData("OutBlock", "bfstdIsuno", 0)
	s.OutBlock.Bnstp = e.GetFieldData("OutBlock", "bnstp", 0)
	s.OutBlock.Ordtrdptncode = e.GetFieldData("OutBlock", "ordtrdptncode", 0)
	s.OutBlock.Mgntrncode = e.GetFieldData("OutBlock", "mgntrncode", 0)
	s.OutBlock.Adduptp = e.GetFieldData("OutBlock", "adduptp", 0)
	s.OutBlock.Commdacode = e.GetFieldData("OutBlock", "commdacode", 0)
	s.OutBlock.Loandt = e.GetFieldData("OutBlock", "Loandt", 0)
	s.OutBlock.Mbrnmbrno = e.GetFieldData("OutBlock", "mbrnmbrno", 0)
	s.OutBlock.Ordacntno = e.GetFieldData("OutBlock", "ordacntno", 0)
	s.OutBlock.Agrgbrnno = e.GetFieldData("OutBlock", "agrgbrnno", 0)
	s.OutBlock.Mgempno = e.GetFieldData("OutBlock", "mgempno", 0)
	s.OutBlock.FutsLnkbrnno = e.GetFieldData("OutBlock", "futsLnkbrnno", 0)
	s.OutBlock.FutsLnkacntno = e.GetFieldData("OutBlock", "futsLnkacntno", 0)
	s.OutBlock.Futsmkttp = e.GetFieldData("OutBlock", "futsmkttp", 0)
	s.OutBlock.Regmktcode = e.GetFieldData("OutBlock", "regmktcode", 0)
	s.OutBlock.Mnymgnrat = e.GetFieldData("OutBlock", "mnymgnrat", 0)
	s.OutBlock.Substmgnrat = e.GetFieldData("OutBlock", "substmgnrat", 0)
	s.OutBlock.Mnyexecamt = e.GetFieldData("OutBlock", "mnyexecamt", 0)
	s.OutBlock.Ubstexecamt = e.GetFieldData("OutBlock", "ubstexecamt", 0)
	s.OutBlock.Cmsnamtexecamt = e.GetFieldData("OutBlock", "cmsnamtexecamt", 0)
	s.OutBlock.Crdtpldgexecamt = e.GetFieldData("OutBlock", "crdtpldgexecamt", 0)
	s.OutBlock.Crdtexecamt = e.GetFieldData("OutBlock", "crdtexecamt", 0)
	s.OutBlock.Prdayruseexecval = e.GetFieldData("OutBlock", "prdayruseexecval", 0)
	s.OutBlock.Crdayruseexecval = e.GetFieldData("OutBlock", "crdayruseexecval", 0)
	s.OutBlock.Spotexecqty = e.GetFieldData("OutBlock", "spotexecqty", 0)
	s.OutBlock.Stslexecqty = e.GetFieldData("OutBlock", "stslexecqty", 0)
	s.OutBlock.Strtgcode = e.GetFieldData("OutBlock", "strtgcode", 0)
	s.OutBlock.GrpID = e.GetFieldData("OutBlock", "grpId", 0)
	s.OutBlock.Ordseqno = e.GetFieldData("OutBlock", "ordseqno", 0)
	s.OutBlock.Ptflno = e.GetFieldData("OutBlock", "ptflno", 0)
	s.OutBlock.Bskno = e.GetFieldData("OutBlock", "bskno", 0)
	s.OutBlock.Trchno = e.GetFieldData("OutBlock", "trchno", 0)
	s.OutBlock.Itemno = e.GetFieldData("OutBlock", "itemno", 0)
	s.OutBlock.OrduserID = e.GetFieldData("OutBlock", "orduserId", 0)
	s.OutBlock.BrwmgmtYn = e.GetFieldData("OutBlock", "brwmgmtYn", 0)
	s.OutBlock.Frgrunqno = e.GetFieldData("OutBlock", "frgrunqno", 0)
	s.OutBlock.TrtzxLevytp = e.GetFieldData("OutBlock", "trtzxLevytp", 0)
	s.OutBlock.Lptp = e.GetFieldData("OutBlock", "lptp", 0)
	s.OutBlock.Exectime = e.GetFieldData("OutBlock", "exectime", 0)
	s.OutBlock.Rcptexectime = e.GetFieldData("OutBlock", "rcptexectime", 0)
	s.OutBlock.RmndLoanamt = e.GetFieldData("OutBlock", "rmndLoanamt", 0)
	s.OutBlock.Secbalqty = e.GetFieldData("OutBlock", "secbalqty", 0)
	s.OutBlock.Spotordableqty = e.GetFieldData("OutBlock", "spotordableqty", 0)
	s.OutBlock.Ordableruseqty = e.GetFieldData("OutBlock", "ordableruseqty", 0)
	s.OutBlock.Flctqty = e.GetFieldData("OutBlock", "flctqty", 0)
	s.OutBlock.Secbalqtyd2 = e.GetFieldData("OutBlock", "secbalqtyd2", 0)
	s.OutBlock.Sellableqty = e.GetFieldData("OutBlock", "sellableqty", 0)
	s.OutBlock.Unercsellordqty = e.GetFieldData("OutBlock", "unercsellordqty", 0)
	s.OutBlock.Avrpchsprc = e.GetFieldData("OutBlock", "avrpchsprc", 0)
	s.OutBlock.Pchsant = e.GetFieldData("OutBlock", "pchsant", 0)
	s.OutBlock.Deposit = e.GetFieldData("OutBlock", "deposit", 0)
	s.OutBlock.Substamt = e.GetFieldData("OutBlock", "substamt", 0)
	s.OutBlock.Csgnmnymgn = e.GetFieldData("OutBlock", "csgnmnymgn", 0)
	s.OutBlock.Csgnsubstmgn = e.GetFieldData("OutBlock", "csgnsubstmgn", 0)
	s.OutBlock.Crdtpldgruseamt = e.GetFieldData("OutBlock", "crdtpldgruseamt", 0)
	s.OutBlock.Ordablemny = e.GetFieldData("OutBlock", "ordablemny", 0)
	s.OutBlock.Ordablesubstamt = e.GetFieldData("OutBlock", "ordablesubstamt", 0)
	s.OutBlock.Ruseableamt = e.GetFieldData("OutBlock", "ruseableamt", 0)

	s.ReceiveRealDataChan <- x
}

func (s SC3) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	s.ReceiveLinkDataChan <- x
}
