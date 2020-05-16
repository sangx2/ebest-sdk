package res

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// CSPAT00600 현물 정상주문
type CSPAT00600InBlock1 struct {
	AcntNo        string `json:"계좌번호"`
	InptPwd       string `json:"입력비밀번호"`
	IsuNo         string `json:"종목번호 {주식:종목코드orA+종목코드(모의투자는 A+종목코드),ELW:J+종목코드}"`
	OrdQty        string `json:"주문수량"`
	OrdPrc        string `json:"주문가"`
	BnsTpCode     string `json:"매매구분 {1:매도,2:매수}"`
	OrdprcPtnCode string `json:"호가유형코드 {00@지정가,03@시장가,05@조건부지정가,06@최유리지정가,07@최우선지정가,61@장개시전시간외종가,81@시간외종가,82@시간외단일가}"`
	MgntrnCode    string `json:"신용거래코드 {000:보통,003:유통/자기융자신규,005:유통대주신규,007:자기대주신규,101:유통융자상환,103:자기융자상환,105:유통대주상환,107:자기대주상환,180:예탁담보대출상환(신용)}"`
	LoanDt        string `json:"대출일"`
	OrdCndiTpCode string `json:"주문조건구분 {0:없음,1:IOC,2:FOK}"`
}

type CSPAT00600OutBlock1 struct {
	gorm.Model

	RecCnt            string `json:"레코드갯수"`
	AcntNo            string `json:"계좌번호"`
	InptPwd           string `json:"입력비밀번호"`
	IsuNo             string `json:"종목번호"`
	OrdQty            string `json:"주문수량"`
	OrdPrc            string `json:"주문가"`
	BnsTpCode         string `json:"매매구분"`
	OrdprcPtnCode     string `json:"호가유형코드"`
	PrgmOrdprcPtnCode string `json:"프로그램호가유형코드"`
	StslAbleYn        string `json:"공매도가능여부"`
	StslOrdprcTpCode  string `json:"공매도호가구분"`
	CommdaCode        string `json:"통신매체코드"`
	MgntrnCode        string `json:"신용거래코드"`
	LoanDt            string `json:"대출일"`
	MbrNo             string `json:"회원번호"`
	OrdCndiTpCode     string `json:"주문조건구분"`
	StrtgCode         string `json:"전략코드"`
	GrpID             string `json:"그룹ID"`
	OrdSeqNo          string `json:"주문회차"`
	PtflNo            string `json:"포트폴리오번호"`
	BskNo             string `json:"바스켓번호"`
	TrchNo            string `json:"트렌치번호"`
	ItemNo            string `json:"아이템번호"`
	OpDrtnNo          string `json:"운용지시번호"`
	LpYn              string `json:"유동성공급자여부"`
	CvrgTpCode        string `json:"반대매매구분"`
}

func (c CSPAT00600OutBlock1) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type CSPAT00600OutBlock2 struct {
	gorm.Model

	RecCnt      string `json:"레코드갯수"`
	OrdNo       string `json:"주문번호"`
	OrdTime     string `json:"주문시각"`
	OrdMktCode  string `json:"주문시장코드"`
	OrdPtnCode  string `json:"주문유형코드"`
	ShtnIsuNo   string `json:"단축종목번호"`
	MgempNo     string `json:"관리사원번호"`
	OrdAmt      string `json:"주문금액"`
	SpareOrdNo  string `json:"예비주문번호"`
	CvrgSeqno   string `json:"반대매매일련번호"`
	RsvOrdNo    string `json:"예약주문번호"`
	SpotOrdQty  string `json:"실물주문수량"`
	RuseOrdQty  string `json:"재사용주문수량"`
	MnyOrdAmt   string `json:"현금주문금액"`
	SubstOrdAmt string `json:"대용주문금액"`
	RuseOrdAmt  string `json:"재사용주문금액"`
	AcntNm      string `json:"계좌명"`
	IsuNm       string `json:"종목명"`
}

func (c CSPAT00600OutBlock2) ToJSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}
