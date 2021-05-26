package impl

import (
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest-sdk/wrapper"
)

// SC0 주식주문접수
type SC0 struct {
	OutBlock res.SC0OutBlock

	ReceiveRealDataChan chan wrapper.XaRealReceiveRealData
	ReceiveLinkDataChan chan wrapper.XaRealReceiveLinkData
}

func NewSC0() *SC0 {
	return &SC0{
		ReceiveRealDataChan: make(chan wrapper.XaRealReceiveRealData, 1),
		ReceiveLinkDataChan: make(chan wrapper.XaRealReceiveLinkData, 1),
	}
}

func (s SC0) GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData {
	return s.ReceiveRealDataChan
}

func (s SC0) GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData {
	return s.ReceiveLinkDataChan
}

func (s *SC0) SetFieldData(e *wrapper.EBestWrapper, resPath string, inBlock interface{}) error {
	e.ResFileName(resPath + "SC0.res")

	return nil
}

func (s SC0) GetOutBlock() interface{} {
	return s.OutBlock
}

func (s *SC0) ReceivedRealData(e *wrapper.EBestWrapper, x wrapper.XaRealReceiveRealData) {
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
	s.OutBlock.Ordchegb = e.GetFieldData("OutBlock", "ordchegb", 0)
	s.OutBlock.Marketgb = e.GetFieldData("OutBlock", "marketgb", 0)
	s.OutBlock.Ordgb = e.GetFieldData("OutBlock", "ordgb", 0)
	s.OutBlock.Orgordno = e.GetFieldData("OutBlock", "orgordno", 0)
	s.OutBlock.Accno1 = e.GetFieldData("OutBlock", "accno1", 0)
	s.OutBlock.Accno2 = e.GetFieldData("OutBlock", "accno2", 0)
	s.OutBlock.Passwd = e.GetFieldData("OutBlock", "passwd", 0)
	s.OutBlock.Expcode = e.GetFieldData("OutBlock", "expcode", 0)
	s.OutBlock.Shtcode = e.GetFieldData("OutBlock", "shtcode", 0)
	s.OutBlock.Hname = e.GetFieldData("OutBlock", "hname", 0)
	s.OutBlock.Ordqty = e.GetFieldData("OutBlock", "ordqty", 0)
	s.OutBlock.Ordprice = e.GetFieldData("OutBlock", "ordprice", 0)
	s.OutBlock.Hogagb = e.GetFieldData("OutBlock", "hogagb", 0)
	s.OutBlock.Etfhogagb = e.GetFieldData("OutBlock", "etfhogagb", 0)
	s.OutBlock.Pgmtype = e.GetFieldData("OutBlock", "pgmtype", 0)
	s.OutBlock.Gmhogagb = e.GetFieldData("OutBlock", "gmhogagb", 0)
	s.OutBlock.Gmhogayn = e.GetFieldData("OutBlock", "gmhogayn", 0)
	s.OutBlock.Singb = e.GetFieldData("OutBlock", "singb", 0)
	s.OutBlock.Loandt = e.GetFieldData("OutBlock", "loandt", 0)
	s.OutBlock.Cvrgordtp = e.GetFieldData("OutBlock", "cvrgordtp", 0)
	s.OutBlock.Strtgcode = e.GetFieldData("OutBlock", "strtgcode", 0)
	s.OutBlock.Groupid = e.GetFieldData("OutBlock", "groupid", 0)
	s.OutBlock.Ordseqno = e.GetFieldData("OutBlock", "ordseqno", 0)
	s.OutBlock.Prtno = e.GetFieldData("OutBlock", "prtno", 0)
	s.OutBlock.Basketno = e.GetFieldData("OutBlock", "basketno", 0)
	s.OutBlock.Trchno = e.GetFieldData("OutBlock", "trchno", 0)
	s.OutBlock.Itemno = e.GetFieldData("OutBlock", "itemno", 0)
	s.OutBlock.Brwmgmyn = e.GetFieldData("OutBlock", "brwmgmyn", 0)
	s.OutBlock.Mbrno = e.GetFieldData("OutBlock", "mbrno", 0)
	s.OutBlock.Procgb = e.GetFieldData("OutBlock", "procgb", 0)
	s.OutBlock.Admbrchno = e.GetFieldData("OutBlock", "admbrchno", 0)
	s.OutBlock.Futaccno = e.GetFieldData("OutBlock", "futaccno", 0)
	s.OutBlock.Futmarketgb = e.GetFieldData("OutBlock", "futmarketgb", 0)
	s.OutBlock.Tongsingb = e.GetFieldData("OutBlock", "tongsingb", 0)
	s.OutBlock.Lpgb = e.GetFieldData("OutBlock", "lpgb", 0)
	s.OutBlock.Dummy = e.GetFieldData("OutBlock", "dummy", 0)
	s.OutBlock.Ordno = e.GetFieldData("OutBlock", "ordno", 0)
	s.OutBlock.Ordtm = e.GetFieldData("OutBlock", "ordtm", 0)
	s.OutBlock.Prntordno = e.GetFieldData("OutBlock", "prntordno", 0)
	s.OutBlock.Mgempno = e.GetFieldData("OutBlock", "mgempno", 0)
	s.OutBlock.Orgordundrqty = e.GetFieldData("OutBlock", "orgordundrqty", 0)
	s.OutBlock.Orgordmdfyqty = e.GetFieldData("OutBlock", "orgordmdfyqty", 0)
	s.OutBlock.Ordordcancelqty = e.GetFieldData("OutBlock", "ordordcancelqty", 0)
	s.OutBlock.Nmcpysndno = e.GetFieldData("OutBlock", "nmcpysndno", 0)
	s.OutBlock.Ordamt = e.GetFieldData("OutBlock", "ordamt", 0)
	s.OutBlock.Bnstp = e.GetFieldData("OutBlock", "bnstp", 0)
	s.OutBlock.Spareordno = e.GetFieldData("OutBlock", "spareordno", 0)
	s.OutBlock.Cvrgseqno = e.GetFieldData("OutBlock", "cvrgseqno", 0)
	s.OutBlock.Rsvordno = e.GetFieldData("OutBlock", "rsvordno", 0)
	s.OutBlock.Mtordseqno = e.GetFieldData("OutBlock", "mtordseqno", 0)
	s.OutBlock.Spareordqty = e.GetFieldData("OutBlock", "spareordqty", 0)
	s.OutBlock.Orduserid = e.GetFieldData("OutBlock", "orduserid", 0)
	s.OutBlock.Spotordqty = e.GetFieldData("OutBlock", "spotordqty", 0)
	s.OutBlock.Ordruseqty = e.GetFieldData("OutBlock", "ordruseqty", 0)
	s.OutBlock.Mnyordamt = e.GetFieldData("OutBlock", "mnyordamt", 0)
	s.OutBlock.Ordsubstamt = e.GetFieldData("OutBlock", "ordsubstamt", 0)
	s.OutBlock.Ruseordamt = e.GetFieldData("OutBlock", "ruseordamt", 0)
	s.OutBlock.Ordcmsnamt = e.GetFieldData("OutBlock", "ordcmsnamt", 0)
	s.OutBlock.Crdtuseamt = e.GetFieldData("OutBlock", "crdtuseamt", 0)
	s.OutBlock.Secbalqty = e.GetFieldData("OutBlock", "secbalqty", 0)
	s.OutBlock.Spotordableqty = e.GetFieldData("OutBlock", "spotordableqty", 0)
	s.OutBlock.Ordableruseqty = e.GetFieldData("OutBlock", "ordableruseqty", 0)
	s.OutBlock.Flctqty = e.GetFieldData("OutBlock", "flctqty", 0)
	s.OutBlock.Secbalqtyd2 = e.GetFieldData("OutBlock", "secbalqtyd2", 0)
	s.OutBlock.Sellableqty = e.GetFieldData("OutBlock", "sellableqty", 0)
	s.OutBlock.Unercsellordqty = e.GetFieldData("OutBlock", "unercsellordqty", 0)
	s.OutBlock.Avrpchsprc = e.GetFieldData("OutBlock", "avrpchsprc", 0)
	s.OutBlock.Pchsamt = e.GetFieldData("OutBlock", "pchsamt", 0)
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

func (s SC0) ReceivedLinkData(ew *wrapper.EBestWrapper, x wrapper.XaRealReceiveLinkData) {
	s.ReceiveLinkDataChan <- x
}
