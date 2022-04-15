package res

import (
	"encoding/json"
)

// CSPAQ13700 현물계좌 예수금/주문가능금액/총평가 조회
type CSPAQ13700InBlock1 struct {
	RecCnt      string `json:"레코드갯수"`
	AcntNo      string `json:"계좌번호"`
	InptPwd     string `json:"입력비밀번호"`
	OrdMktCode  string `json:"주문시장코드"`
	BnsTpCode   string `json:"매매구분"`
	IsuNo       string `json:"종목번호"`
	ExecYn      string `json:"체결여부"`
	OrdDt       string `json:"주문일"`
	SrtOrdNo2   string `json:"시작주문번호2"`
	BkseqTpCode string `json:"역순구분"`
	OrdPtnCode  string `json:"주문유형코드"`
}

type CSPAQ13700OutBlock1 struct {
	RecCnt      string `json:"레코드갯수"`
	AcntNo      string `json:"계좌번호"`
	InptPwd     string `json:"입력비밀번호"`
	OrdMktCode  string `json:"주문시장코드"`
	BnsTpCode   string `json:"매매구분"`
	IsuNo       string `json:"종목번호"`
	ExecYn      string `json:"체결여부"`
	OrdDt       string `json:"주문일"`
	SrtOrdNo2   string `json:"시작주문번호2"`
	BkseqTpCode string `json:"역순구분"`
	OrdPtnCode  string `json:"주문유형코드"`
}

func (c CSPAQ13700OutBlock1) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAQ13700OutBlock2 struct {
	RecCnt      string `json:"레코드갯수"`
	SellExecAmt string `json:"매도체결금액"`
	BuyExecAmt  string `json:"매수체결금액"`
	SellExecQty string `json:"매도체결수량"`
	BuyExecQty  string `json:"매수체결수량"`
	SellOrdQty  string `json:"매도주문수량"`
	BuyOrdQty   string `json:"매수주문수량"`
}

func (c CSPAQ13700OutBlock2) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAQ13700OutBlock3 struct {
	OrdDt         string `json:"주문일"`
	Mgmtbrnno     string `json:"관리지점번호"`
	Ordmktcode    string `json:"주문시장코드"`
	OrdNo         string `json:"주문번호"`
	OrgOrdNo      string `json:"원주문번호"`
	IsuNo         string `json:"종목번호"`
	IsuNm         string `json:"종목명"`
	BnsTpCode     string `json:"매매구분코드"`
	BnsTpNm       string `json:"매매구분"`
	OrdPtnCode    string `json:"주문유형코드"`
	OrdPtnNm      string `json:"주문유형명"`
	OrdTrxPtnCode string `json:"주문처리유형코드"`
	OrdTrxPtnNm   string `json:"주문처리유형명"`
	MrcTpCode     string `json:"정정취소구분"`
	MrcTpNm       string `json:"정정취소구분명"`
	MrcQty        string `json:"정정취소수량"`
	MrcAbleQty    string `json:"정정취소가능수량"`
	OrdQty        string `json:"주문수량"`
	OrdPrc        string `json:"주문가격"`
	ExecQty       string `json:"체결수량"`
	ExecPrc       string `json:"체결가"`
	ExecTrxTime   string `json:"체결처리시각"`
	LastExecTime  string `json:"최종체결시각"`
	OrdprcPtnCode string `json:"호가유형코드"`
	OrdprcPtnNm   string `json:"호가유형명"`
	OrdCndiTpCode string `json:""`
	AllExecQty    string `json:"전체체결수량"`
	RegCommdaCode string `json:"통신매체코드"`
	CommdaNm      string `json:"통신매체명"`
	MbrNo         string `json:"회원번호"`
	RsvOrdYn      string `json:"예약주문여부"`
	LoanDt        string `json:"대출일"`
	OrdTime       string `json:"주문시각"`
	OpDrtnNo      string `json:"운용지시번호"`
	OdrrId        string `json:"주문자ID"`
}

func (c CSPAQ13700OutBlock3) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
