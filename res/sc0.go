package res

import (
	"encoding/json"
)

// SC0 주식주문접수
type SC0OutBlock struct {
	Lineseq         string `json:"라인일련번호"`
	Accno           string `json:"계좌번호"`
	User            string `json:"조작자ID"`
	Len             string `json:"헤더길이"`
	Gubun           string `json:"헤더구분"`
	Compress        string `json:"압축구분"`
	Encrypt         string `json:"암호구분"`
	Offset          string `json:"공통시작지점"`
	Trcode          string `json:"TRCODE"`
	Compid          string `json:"이용사번호"`
	Userid          string `json:"사용자ID"`
	Media           string `json:"접속매체"`
	Ifid            string `json:"I/F일련번호"`
	Seq             string `json:"전문일련번호"`
	Trid            string `json:"TR추적ID"`
	Pubip           string `json:"공인IP"`
	Prvip           string `json:"사설IP"`
	Pcbpno          string `json:"처리지점번호"`
	Bpno            string `json:"지점번호"`
	Termno          string `json:"단말번호"`
	Lang            string `json:"언어구분"`
	Proctm          string `json:"AP처리시간"`
	Msgcode         string `json:"메시지코드"`
	Outgu           string `json:"메시지출력구분"`
	Compreq         string `json:"압축요청구분"`
	Funckey         string `json:"기능키"`
	Reqcnt          string `json:"요청레코드개수"`
	Filler          string `json:"예비영역"`
	Cont            string `json:"연속구분"`
	Contkey         string `json:"연속키값"`
	Varlen          string `json:"가변시스템길이"`
	Varhdlen        string `json:"가변헤더길이"`
	Varmsglen       string `json:"가변메시지길이"`
	Trsrc           string `json:"조회발원지"`
	Eventid         string `json:"I/F이벤트ID"`
	Ifinfo          string `json:"I/F정보"`
	Filler1         string `json:"예비영역1"`
	Ordchegb        string `json:"주문체결구분"`
	Marketgb        string `json:"마켓구분"`
	Ordgb           string `json:"주문구분"`
	Orgordno        string `json:"원주문번호"`
	Accno1          string `json:"계좌번호1"`
	Accno2          string `json:"계좌번호2"`
	Passwd          string `json:"비밀번호"`
	Expcode         string `json:"종목번호"`
	Shtcode         string `json:"단축종목번호"`
	Hname           string `json:"종목명"`
	Ordqty          string `json:"주문수량"`
	Ordprice        string `json:"주문가격"`
	Hogagb          string `json:"주문조건"`
	Etfhogagb       string `json:"호가유형코드"`
	Pgmtype         string `json:"프로그램호가구분"`
	Gmhogagb        string `json:"공매도호가구분"`
	Gmhogayn        string `json:"공매도가능여부"`
	Singb           string `json:"신용구분"`
	Loandt          string `json:"대출일"`
	Cvrgordtp       string `json:"반대매매주문구분"`
	Strtgcode       string `json:"전략코드"`
	Groupid         string `json:"그룹ID"`
	Ordseqno        string `json:"주문회차"`
	Prtno           string `json:"포트폴리오번호"`
	Basketno        string `json:"바스켓번호"`
	Trchno          string `json:"트렌치번호"`
	Itemno          string `json:"아이템번호"`
	Brwmgmyn        string `json:"차입구분"`
	Mbrno           string `json:"회원사구분"`
	Procgb          string `json:"처리구분"`
	Admbrchno       string `json:"관리지점번호"`
	Futaccno        string `json:"선물계좌번호"`
	Futmarketgb     string `json:"선물상품구분"`
	Tongsingb       string `json:"통신매체구분"`
	Lpgb            string `json:"유동성공급자구분"`
	Dummy           string `json:"DUMMY"`
	Ordno           string `json:"주문번호"`
	Ordtm           string `json:"주문시각"`
	Prntordno       string `json:"모주문번호"`
	Mgempno         string `json:"관리사원번호"`
	Orgordundrqty   string `json:"원주문미체결수량"`
	Orgordmdfyqty   string `json:"원주문정정수량"`
	Ordordcancelqty string `json:"원주문취소수량"`
	Nmcpysndno      string `json:"비회원사송신번호"`
	Ordamt          string `json:"주문금액"`
	Bnstp           string `json:"매매구분"`
	Spareordno      string `json:"예비주문번호"`
	Cvrgseqno       string `json:"반대매매일련번호"`
	Rsvordno        string `json:"예약주문번호"`
	Mtordseqno      string `json:"복수주문일련번호"`
	Spareordqty     string `json:"예비주문수량"`
	Orduserid       string `json:"주문사원번호"`
	Spotordqty      string `json:"실물주문수량"`
	Ordruseqty      string `json:"재사용주문수량"`
	Mnyordamt       string `json:"현금주문금액"`
	Ordsubstamt     string `json:"주문대용금액"`
	Ruseordamt      string `json:"재사용주문금액"`
	Ordcmsnamt      string `json:"수수료주문금액"`
	Crdtuseamt      string `json:"사용신용담보재사용금"`
	Secbalqty       string `json:"잔고수량"`
	Spotordableqty  string `json:"실물가능수량"`
	Ordableruseqty  string `json:"재사용가능수량(매도)"`
	Flctqty         string `json:"변동수량"`
	Secbalqtyd2     string `json:"잔고수량(D2)"`
	Sellableqty     string `json:"매도주문가능수량"`
	Unercsellordqty string `json:"미체결매도주문수량"`
	Avrpchsprc      string `json:"평균매입가"`
	Pchsamt         string `json:"매입금액"`
	Deposit         string `json:"예수금"`
	Substamt        string `json:"대용금"`
	Csgnmnymgn      string `json:"위탁증거금현금"`
	Csgnsubstmgn    string `json:"위탁증거금대용"`
	Crdtpldgruseamt string `json:"신용담보재사용금"`
	Ordablemny      string `json:"주문가능현금"`
	Ordablesubstamt string `json:"주문가능대용"`
	Ruseableamt     string `json:"재사용가능금액"`
}

func (s *SC0OutBlock) ToJSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}
