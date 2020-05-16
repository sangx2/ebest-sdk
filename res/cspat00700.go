package res

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// CSPAT00700 현물 정정주문
type CSPAT00700InBlock1 struct {
	OrgOrdNo      string `json:"원주문번호"`
	AcntNo        string `json:"계좌번호"`
	InptPwd       string `json:"입력비밀번호"`
	IsuNo         string `json:"종목번호 {주식:종목코드orA+종목코드(모의투자는 A+종목코드),ELW:J+종목코드}"`
	OrdQty        string `json:"주문수량"`
	OrdprcPtnCode string `json:"호가유형코드 {00@지정가,03@시장가,05@조건부지정가,06@최유리지정가,07@최우선지정가,61@장개시전시간외종가,81@시간외종가,82@시간외단일가}"`
	OrdCndiTpCode string `json:"주문조건구분 {0:없음,1:IOC,2:FOK}"`
	OrdPrc        string `json:"주문가"`
}

type CSPAT00700OutBlock1 struct {
	gorm.Model

	RecCnt        string `json:"레코드갯수"`
	OrgOrdNo      string `json:"원주문번호"`
	AcntNo        string `json:"계좌번호"`
	InptPwd       string `json:"입력비밀번호"`
	IsuNo         string `json:"종목번호"`
	OrdQty        string `json:"주문수량"`
	OrdprcPtnCode string `json:"호가유형코드"`
	OrdCndiTpCode string `json:"주문조건구분"`
	OrdPrc        string `json:"주문가"`
	CommdaCode    string `json:"통신매체코드"`
	StrtgCode     string `json:"전략코드"`
	GrpID         string `json:"그룹ID"`
	OrdSeqNo      string `json:"주문회차"`
	PtflNo        string `json:"포트폴리오번호"`
	BskNo         string `json:"바스켓번호"`
	TrchNo        string `json:"트렌치번호"`
	ItemNo        string `json:"아이템번호"`
}

func (c CSPAT00700OutBlock1) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAT00700OutBlock2 struct {
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
	OrdAmt            string `json:"주문금액"`
	BnsTpCode         string `json:"매매구분"`
	SpareOrdNo        string `json:"예비주문번호"`
	CvrgSeqno         string `json:"반대매매일련번호"`
	RsvOrdNo          string `json:"예약주문번호"`
	MnyOrdAmt         string `json:"현금주문금액"`
	SubstOrdAmt       string `json:"대용주문금액"`
	RuseOrdAmt        string `json:"재사용주문금액"`
	AcntNm            string `json:"계좌명"`
	IsuNm             string `json:"종목명"`
}

func (c CSPAT00700OutBlock2) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
