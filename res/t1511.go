package res

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// T1511 업종 현재가
type T1511InBlock struct {
	Upcode string `json:"업종코드"`
}

type T1511OutBlock struct {
	gorm.Model `json:"-"`

	Gubun        string `json:"업종구분(1:KOSPI업종 2:KOSDAQ업종 3:섹터지수 9:기타)"`
	Hname        string `json:"업종명"`
	Pricejisu    string `json:"현재지수"`
	Jniljisu     string `json:"전일지수"`
	Sign         string `json:"전일대비구분(1:상한 2:상승 3:보합 4:하한 5:하락)"`
	Change       string `json:"전일대비"`
	Diffjisu     string `json:"지수등락율"`
	Jnilvolume   string `json:"전일거래량"`
	Volume       string `json:"당일거래량"`
	Volumechange string `json:"거래량전일대비"`
	Volumerate   string `json:"거래량비율"`
	Jnilvalue    string `json:"전일거래대금"`
	Value        string `json:"당일거래대금"`
	Valuechange  string `json:"거래대금전일대비"`
	Valuerate    string `json:"거래대금비율"`
	Openjisu     string `json:"시가지수"`
	Opendiff     string `json:"시가등락율"`
	Opentime     string `json:"시가시간"`
	Highjisu     string `json:"고가지수"`
	Highdiff     string `json:"고가등락율"`
	Hightime     string `json:"고가시간"`
	Lowjisu      string `json:"저가지수"`
	Lowdiff      string `json:"저가등락율"`
	Lowtime      string `json:"저가시간"`
	Whjisu       string `json:"52주최고지수"`
	Whchange     string `json:"52주최고현재가대비"`
	Whjday       string `json:"52주최고지수일차"`
	Wljisu       string `json:"52주최저지수"`
	Wlchange     string `json:"52주최저현재가대비"`
	Wljday       string `json:"52주최저지수일차"`
	Yhjisu       string `json:"연중최고지수"`
	Yhchange     string `json:"연중최고현재가대비"`
	Yhjday       string `json:"연중최고지수일차"`
	Yljisu       string `json:"연중최저지수"`
	Ylchange     string `json:"연중최저현재가대비"`
	Yljday       string `json:"연중최저지수일차"`
	Firstjcode   string `json:"첫번째지수코드"`
	Firstjname   string `json:"첫번째지수명"`
	Firstjisu    string `json:"첫번째지수"`
	Firsign      string `json:"첫번째대비구분(1:상한 2:상승 3:보합 4:하한 5:하락)"`
	Firchange    string `json:"첫번째전일대비"`
	Firdiff      string `json:"첫번째등락율"`
	Secondjcode  string `json:"두번째지수코드"`
	Secondjname  string `json:"두번째지수명"`
	Secondjisu   string `json:"두번째지수"`
	Secsign      string `json:"두번째대비구분(1:상한 2:상승 3:보합 4:하한 5:하락)"`
	Secchange    string `json:"두번째전일대비"`
	Secdiff      string `json:"두번째등락율"`
	Thirdjcode   string `json:"세번째지수코드"`
	Thirdjname   string `json:"세번째지수명"`
	Thirdjisu    string `json:"세번째지수"`
	Thrsign      string `json:"세번째대비구분(1:상한 2:상승 3:보합 4:하한 5:하락)"`
	Thrchange    string `json:"세번째전일대비"`
	Thrdiff      string `json:"세번째등락율"`
	Fourthjcode  string `json:"네번째지수코드"`
	Fourthjname  string `json:"네번째지수명"`
	Fourthjisu   string `json:"네번째지수"`
	Forsign      string `json:"네번째대비구분(1:상한,2:상승,3:보합,4:하한,5:하락)"`
	Forchange    string `json:"네번째전일대비"`
	Fordiff      string `json:"네번째등락율"`
	Highjo       string `json:"상승종목수"`
	Upjo         string `json:"상한종목수"`
	Unchgjo      string `json:"보합종목수"`
	Lowjo        string `json:"하락종목수"`
	Downjo       string `json:"하한종목수"`
}

func (t T1511OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
