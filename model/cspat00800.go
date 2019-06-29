package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// CSPAT00800 현물 취소주문
type CSPAT00800InBlock1 struct {
	OrgOrdNo string `json:"원주문번호"`
	AcntNo   string `json:"계좌번호"`
	InptPwd  string `json:"입력비밀번호"`
	IsuNo    string `json:"종목번호"`
	OrdQty   string `json:"주문수량"`
}

type CSPAT00800OutBlock1 struct {
	gorm.Model

	RecCnt     string `json:"레코드갯수"`
	OrgOrdNo   string `json:"원주문번호"`
	AcntNo     string `json:"계좌번호"`
	InptPwd    string `json:"입력비밀번호"`
	IsuNo      string `json:"종목번호"`
	OrdQty     string `json:"주문수량"`
	CommdaCode string `json:"통신매체코드"`
	GrpID      string `json:"그룹ID"`
	StrtgCode  string `json:"전략코드"`
	OrdSeqNo   string `json:"주문회차"`
	PtflNo     string `json:"포트폴리오번호"`
	BskNo      string `json:"바스켓번호"`
	TrchNo     string `json:"트렌치번호"`
	ItemNo     string `json:"아이템번호"`
}

func (c CSPAT00800OutBlock1) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAT00800OutBlock2 struct {
	gorm.Model

	RecCnt            string `json:"레코드갯수"`
	OrdNo             string `json:"주문번호"`
	PrntOrdNo         string `json:"모주문번호"`
	OrdTime           string `json:"주문시각"`
	OrdMktCode        string `json:"주문시장코드"`
	OrdPtnCode        string `json:"주문유형코드"`
	ShtnIsuNo         string `json:"단축종목코드"`
	PrgmOrdprcPtnCode string `json:"프로그램호가유형코드"`
	StslOrdprcTpCode  string `json:"공매도호가구분"`
	StslAbleYn        string `json:"공매도가능여부"`
	MgntrnCode        string `json:"신용거래코드"`
	LoanDt            string `json:"대출일"`
	CvrgOrdTp         string `json:"반대매매주문구분"`
	LpYn              string `json:"유동성공급자여부"`
	MgempNo           string `json:"관리사원번호"`
	BnsTpCode         string `json:"매매구분"`
	SpareOrdNo        string `json:"예비주문번호"`
	CvrgSeqno         string `json:"반대매매일련번호"`
	RsvOrdNo          string `json:"예약주문번호"`
	AcntNm            string `json:"계좌명"`
	IsuNm             string `json:"종목명"`
}

func (c CSPAT00800OutBlock2) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
