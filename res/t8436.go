package res

import (
	"encoding/json"
)

// T8436 주식종목조회 API용
type T8436InBlock struct {
	Gubun string `json:"구분1"`
}

type T8436OutBlock struct {
	Hname      string `json:"종목명"`
	Shcode     string `json:"단축코드"`
	Expcode    string `json:"확장코드"`
	Etfgubun   string `json:"ETF구분"`
	Uplmtprice string `json:"상한가"`
	Dnlmtprice string `json:"하한가"`
	Jnilclose  string `json:"전일가"`
	Memedan    string `json:"주문수량단위"`
	Recprice   string `json:"기준가"`
	Gubun      string `json:"구분{1:코스피 2:코스닥}"`
	Bu12gubun  string `json:"증권그룹"`
	SpacGubun  string `json:"기업인수목적회사여부"`
	Filler     string `gorm:"-" json:"-"` // 미사용
}

func (t T8436OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
