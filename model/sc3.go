package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// SC3 주식주문취소
type SC3OutBlock struct {
	gorm.Model

	Lineseq          string `json:"라인일련번호"`
	Accno            string `json:"계좌번호"`
	User             string `json:"조작자ID"`
	Len              string `json:"헤더길이"`
	Gubun            string `json:"헤더구분"`
	Compress         string `json:"압축구분"`
	Encrypt          string `json:"암호구분"`
	Offset           string `json:"공통시작지점"`
	Trcode           string `json:"TRCODE"`
	Compid           string `json:"이용사번호"`
	Userid           string `json:"사용자ID"`
	Media            string `json:"접속매체"`
	Ifid             string `json:"I/F일련번호"`
	Seq              string `json:"전문일련번호"`
	Trid             string `json:"TR추적ID"`
	Pubip            string `json:"공인IP"`
	Prvip            string `json:"사설IP"`
	Pcbpno           string `json:"처리지점번호"`
	Bpno             string `json:"지점번호"`
	Termno           string `json:"단말번호"`
	Lang             string `json:"언어구분"`
	Proctm           string `json:"AP처리시간"`
	Msgcode          string `json:"메시지코드"`
	Outgu            string `json:"메시지출력구분"`
	Compreq          string `json:"압축요청구분"`
	Funckey          string `json:"기능키"`
	Reqcnt           string `json:"요청레코드개수"`
	Filler           string `json:"예비영역"`
	Cont             string `json:"연속구분"`
	Contkey          string `json:"연속키값"`
	Varlen           string `json:"가변시스템길이"`
	Varhdlen         string `json:"가변헤더길이"`
	Varmsglen        string `json:"가변메시지길이"`
	Trsrc            string `json:"조회발원지"`
	Eventid          string `json:"I/F이벤트ID"`
	Ifinfo           string `json:"I/F정보"`
	Filler1          string `json:"예비영역1"`
	Ordxctptncode    string `json:"주문체결유형코드"`
	Ordmktcode       string `json:"주문시장코드"`
	Ordptncode       string `json:"주문유형코드"`
	Mgmtbrnno        string `json:"관리지점번호"`
	Accno1           string `json:"계좌번호1"`
	Accno2           string `json:"계좌번호2"`
	Acntnm           string `json:"계좌명"`
	Isuno            string `json:"종목번호"`
	Isunm            string `json:"종목명"`
	Ordno            string `json:"주문번호"`
	Orgordno         string `json:"원주문번호"`
	Execno           string `json:"체결번호"`
	Ordqty           string `json:"주문수량"`
	Ordprc           string `json:"주문가격"`
	Execqty          string `json:"체결수량"`
	Execprc          string `json:"체결가격"`
	Mdfycnfqty       string `json:"정정확인수량"`
	Mdfycnfprc       string `json:"정정확인가격"`
	Canccnfqty       string `json:"취소확인수량"`
	Rjtqty           string `json:"거부수량"`
	Ordtrxptncode    string `json:"주문처리유형코드"`
	Mtiordseqno      string `json:"복수주문일련번호"`
	Ordcndi          string `json:"주문조건"`
	Ordprcptncode    string `json:"호가유형코드"`
	Nsavtrdqty       string `json:"비저축체결수량"`
	ShtnIsuno        string `json:"단축종목번호"`
	Opdrtnno         string `json:"운용지시번호"`
	Cvrgordtp        string `json:"반대매매주문구분"`
	Unercqty         string `json:"미체결수량(주문)"`
	Orgordunercqty   string `json:"원주문미체결수량"`
	Orgordmdfyqty    string `json:"원주문정정수량"`
	Orgordcancqty    string `json:"원주문취소수량"`
	Ordavrexecprc    string `json:"주문평균체결가격"`
	Ordamt           string `json:"주문금액"`
	StdIsuno         string `json:"표준종목번호"`
	BfstdIsuno       string `json:"전표준종목번호"`
	Bnstp            string `json:"매매구분"`
	Ordtrdptncode    string `json:"주문거래유형코드"`
	Mgntrncode       string `json:"신용거래코드"`
	Adduptp          string `json:"수수료합산코드"`
	Commdacode       string `json:"통신매체코드"`
	Loandt           string `json:"대출일"`
	Mbrnmbrno        string `json:"회원/비회원사번호"`
	Ordacntno        string `json:"주문계좌번호"`
	Agrgbrnno        string `json:"집계지점번호"`
	Mgempno          string `json:"관리사원번호"`
	FutsLnkbrnno     string `json:"선물연계지점번호"`
	FutsLnkacntno    string `json:"선물연계계좌번호"`
	Futsmkttp        string `json:"선물시장구분"`
	Regmktcode       string `json:"등록시장코드"`
	Mnymgnrat        string `json:"현금증거금률"`
	Substmgnrat      string `json:"대용증거금률"`
	Mnyexecamt       string `json:"현금체결금액"`
	Ubstexecamt      string `json:"대용체결금액"`
	Cmsnamtexecamt   string `json:"수수료체결금액"`
	Crdtpldgexecamt  string `json:"신용담보체결금액"`
	Crdtexecamt      string `json:"신용체결금액"`
	Prdayruseexecval string `json:"전일재사용체결금액"`
	Crdayruseexecval string `json:"금일재사용체결금액"`
	Spotexecqty      string `json:"실물체결수량"`
	Stslexecqty      string `json:"공매도체결수량"`
	Strtgcode        string `json:"전략코드"`
	GrpID            string `json:"그룹Id"`
	Ordseqno         string `json:"주문회차"`
	Ptflno           string `json:"포트폴리오번호"`
	Bskno            string `json:"바스켓번호"`
	Trchno           string `json:"트렌치번호"`
	Itemno           string `json:"아이템번호"`
	OrduserID        string `json:"주문자Id"`
	BrwmgmtYn        string `json:"차입관리여부"`
	Frgrunqno        string `json:"외국인고유번호"`
	TrtzxLevytp      string `json:"거래세징수구분"`
	Lptp             string `json:"유동성공급자구분"`
	Exectime         string `json:"체결시각"`
	Rcptexectime     string `json:"거래소수신체결시각"`
	RmndLoanamt      string `json:"잔여대출금액"`
	Secbalqty        string `json:"잔고수량"`
	Spotordableqty   string `json:"실물가능수량"`
	Ordableruseqty   string `json:"재사용가능수량(매도)"`
	Flctqty          string `json:"변동수량"`
	Secbalqtyd2      string `json:"잔고수량(d2)"`
	Sellableqty      string `json:"매도주문가능수량"`
	Unercsellordqty  string `json:"미체결매도주문수량"`
	Avrpchsprc       string `json:"평균매입가"`
	Pchsant          string `json:"매입금액"`
	Deposit          string `json:"예수금"`
	Substamt         string `json:"대용금"`
	Csgnmnymgn       string `json:"위탁증거금현금"`
	Csgnsubstmgn     string `json:"위탁증거금대용"`
	Crdtpldgruseamt  string `json:"신용담보재사용금"`
	Ordablemny       string `json:"주문가능현금"`
	Ordablesubstamt  string `json:"주문가능대용"`
	Ruseableamt      string `json:"재사용가능금액"`
}

func (s *SC3OutBlock) ToJSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}
