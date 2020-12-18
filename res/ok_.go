package res

import (
	"encoding/json"
)

// ok KOSDAQ 거래원
type OKInBlock struct {
	Shcode string `json:"단축코드"`
}

type OKOutBlock struct {
	Offerno1    string `json:"매도증권사코드1"`
	Bidno1      string `json:"매수증권사코드1"`
	Offertrad1  string `json:"매도회원사명1"`
	Bidtrad1    string `json:"매수회원사명1"`
	Tradmdvol1  string `json:"매도거래량1"`
	Tradmsvol1  string `json:"매수거래량1"`
	Tradmdrate1 string `json:"매도거래량비중1"`
	Tradmsrate1 string `json:"매수거개량비중1"`
	Tradmdcha1  string `json:"매도거래량직전대비1"`
	Tradmscha1  string `json:"매수거래량직전대비1"`

	Offerno2    string `json:"매도증권사코드2"`
	Bidno2      string `json:"매수증권사코드2"`
	Offertrad2  string `json:"매도회원사명2"`
	Bidtrad2    string `json:"매수회원사명2"`
	Tradmdvol2  string `json:"매도거래량2"`
	Tradmsvol2  string `json:"매수거래량2"`
	Tradmdrate2 string `json:"매도거래량비중2"`
	Tradmsrate2 string `json:"매수거개량비중2"`
	Tradmdcha2  string `json:"매도거래량직전대비2"`
	Tradmscha2  string `json:"매수거래량직전대비2"`

	Offerno3    string `json:"매도증권사코드3"`
	Bidno3      string `json:"매수증권사코드3"`
	Offertrad3  string `json:"매도회원사명3"`
	Bidtrad3    string `json:"매수회원사명3"`
	Tradmdvol3  string `json:"매도거래량3"`
	Tradmsvol3  string `json:"매수거래량3"`
	Tradmdrate3 string `json:"매도거래량비중3"`
	Tradmsrate3 string `json:"매수거개량비중3"`
	Tradmdcha3  string `json:"매도거래량직전대비3"`
	Tradmscha3  string `json:"매수거래량직전대비3"`

	Offerno4    string `json:"매도증권사코드4"`
	Bidno4      string `json:"매수증권사코드4"`
	Offertrad4  string `json:"매도회원사명4"`
	Bidtrad4    string `json:"매수회원사명4"`
	Tradmdvol4  string `json:"매도거래량4"`
	Tradmsvol4  string `json:"매수거래량4"`
	Tradmdrate4 string `json:"매도거래량비중4"`
	Tradmsrate4 string `json:"매수거개량비중4"`
	Tradmdcha4  string `json:"매도거래량직전대비4"`
	Tradmscha4  string `json:"매수거래량직전대비4"`

	Offerno5    string `json:"매도증권사코드5"`
	Bidno5      string `json:"매수증권사코드5"`
	Offertrad5  string `json:"매도회원사명5"`
	Bidtrad5    string `json:"매수회원사명5"`
	Tradmdvol5  string `json:"매도거래량5"`
	Tradmsvol5  string `json:"매수거래량5"`
	Tradmdrate5 string `json:"매도거래량비중5"`
	Tradmsrate5 string `json:"매수거개량비중5"`
	Tradmdcha5  string `json:"매도거래량직전대비5"`
	Tradmscha5  string `json:"매수거래량직전대비5"`

	Ftradmdvol  string `json:"외국계증권사매도합계"`
	Ftradmsvol  string `json:"외국계증권사매수합계"`
	Ftradmdrate string `json:"외국계증권사매도거개량비중"`
	Ftradmsrate string `json:"외국계증권사매수거래량비중"`
	Ftradmdcha  string `json:"외국계증권사매도거래량직전대비"`
	Ftradmscha  string `json:"외국계증권사매수거개량직전대비"`

	Shcode string `json:"단축코드"`

	Tradmdval1 string `json:"매도거래대금1"`
	Tradmsval1 string `json:"매수거래대금1"`
	Tradmdavg1 string `json:"매도평균단가1"`
	Tradmsavg1 string `json:"매수평균단가1"`

	Tradmdval2 string `json:"매도거래대금2"`
	Tradmsval2 string `json:"매수거래대금2"`
	Tradmdavg2 string `json:"매도평균단가2"`
	Tradmsavg2 string `json:"매수평균단가2"`

	Tradmdval3 string `json:"매도거래대금3"`
	Tradmsval3 string `json:"매수거래대금3"`
	Tradmdavg3 string `json:"매도평균단가3"`
	Tradmsavg3 string `json:"매수평균단가3"`

	Tradmdval4 string `json:"매도거래대금4"`
	Tradmsval4 string `json:"매수거래대금4"`
	Tradmdavg4 string `json:"매도평균단가4"`
	Tradmsavg4 string `json:"매수평균단가4"`

	Tradmdval5 string `json:"매도거래대금5"`
	Tradmsval5 string `json:"매수거래대금5"`
	Tradmdavg5 string `json:"매도평균단가5"`
	Tradmsavg5 string `json:"매수평균단가5"`

	Ftradmdval string `json:"외국계증권사매도거래대금"`
	Ftradmsval string `json:"외국계증권사매수거래대금"`
	Ftradmdavg string `json:"외국계증권사매도평균단가"`
	Ftradmsavg string `json:"외국계증권사매수평균단가"`
}

func (o OKOutBlock) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}
