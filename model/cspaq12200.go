package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// CSPAQ12200 현물계좌 예수금/주문가능금액/총평가 조회
type CSPAQ12200InBlock1 struct {
	RecCnt    string `json:"레코드갯수"`
	MgmtBrnNo string `json:"관리지점번호"`
	AcntNo    string `json:"계좌번호"`
	Pwd       string `json:"비밀번호"`
	BalCreTp  string `json:"잔고생성구분"`
}

type CSPAQ12200OutBlock1 struct {
	gorm.Model

	RecCnt    string `json:"레코드갯수"`
	MgmtBrnNo string `json:"관리지점번호"`
	AcntNo    string `json:"계좌번호"`
	Pwd       string `json:"비밀번호"`
	BalCreTp  string `json:"잔고생성구분"`
}

func (c CSPAQ12200OutBlock1) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAQ12200OutBlock2 struct {
	gorm.Model

	RecCnt                 string `json:"레코드갯수"`
	BrnNm                  string `json:"지점명"`
	AcntNm                 string `json:"계좌명"`
	MnyOrdAbleAmt          string `json:"현금주문가능금액"`
	MnyoutAbleAmt          string `json:"출금가능금액"`
	SeOrdAbleAmt           string `json:"거래소금액"`
	KdqOrdAbleAmt          string `json:"코스닥금액"`
	BalEvalAmt             string `json:"잔고평가금액"`
	RcvblAmt               string `json:"미수금액"`
	DpsastTotamt           string `json:"예탁자산총액"`
	PnlRat                 string `json:"손익율"`
	InvstOrgAmt            string `json:"투자원금"`
	InvstPlAmt             string `json:"투자손익금액"`
	CrdtPldgOrdAmt         string `json:"신용담보주문금액"`
	Dps                    string `json:"예수금"`
	SubstAmt               string `json:"대용금액"`
	D1Dps                  string `json:"D1예수금"`
	D2Dps                  string `json:"D2예수금"`
	MnyrclAmt              string `json:"현금미수금액"`
	MgnMny                 string `json:"증거금현금"`
	MgnSubst               string `json:"증거금대용"`
	ChckAmt                string `json:"수표금액"`
	SubstOrdAbleAmt        string `json:"대용주문가능금액"`
	MgnRat100pctOrdAbleAmt string `json:"증거금률100퍼센트주문가능금액"`
	MgnRat35ordAbleAmt     string `json:"증거금률35%주문가능금액"`
	MgnRat50ordAbleAmt     string `json:"증거금률50%주문가능금액"`
	PrdaySellAdjstAmt      string `json:"전일매도정산금액"`
	PrdayBuyAdjstAmt       string `json:"전일매수정산금액"`
	CrdaySellAdjstAmt      string `json:"금일매도정산금액"`
	CrdayBuyAdjstAmt       string `json:"금일매수정산금액"`
	D1ovdRepayRqrdAmt      string `json:"D1연체변제소요금액"`
	D2ovdRepayRqrdAmt      string `json:"D2연체변제소요금액"`
	D1PrsmptWthdwAbleAmt   string `json:"D1추정인출가능금액"`
	D2PrsmptWthdwAbleAmt   string `json:"D2추정인출가능금액"`
	DpspdgLoanAmt          string `json:"예탁담보대출금액"`
	Imreq                  string `json:"신용설정보증금"`
	MloanAmt               string `json:"융자금액"`
	ChgAfPldgRat           string `json:"변경후담보비율"`
	OrgPldgAmt             string `json:"원담보금액"`
	SubPldgAmt             string `json:"부담보금액"`
	RqrdPldgAmt            string `json:"소요담보금액"`
	OrgPdlckAmt            string `json:"원담보부족금액"`
	PdlckAmt               string `json:"담보부족금액"`
	AddPldgMny             string `json:"추가담보현금"`
	D1OrdAbleAmt           string `json:"D1주문가능금액"`
	CrdtIntdltAmt          string `json:"신용이자미납금액"`
	EtclndAmt              string `json:"기타대여금액"`
	NtdayPrsmptCvrgAmt     string `json:"익일추정반대매매금액"`
	OrgPldgSumAmt          string `json:"원담보합계금액"`
	CrdtOrdAbleAmt         string `json:"신용주문가능금액"`
	SubPldgSumAmt          string `json:"부담보합계금액"`
	CrdtPldgAmtMny         string `json:"신용담보금현금"`
	CrdtPldgSubstAmt       string `json:"신용담보대용금액"`
	AddCrdtPldgMny         string `json:"추가신용담보현금"`
	CrdtPldgRuseAmt        string `json:"신용담보재사용금액"`
	AddCrdtPldgSubst       string `json:"추가신용담보대용"`
	CslLoanAmtdt1          string `json:"매도대금담보대출금액"`
	DpslRestrcAmt          string `json:"처분제한금액"`
}

func (c CSPAQ12200OutBlock2) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
