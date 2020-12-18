package res

import (
	"encoding/json"
)

// S3 KOSPI 체결
type S3InBlock struct {
	Shcode string `json:"단축코드"`
}

type S3OutBlock struct {
	Chetime    string `json:"체결시간"`
	Sign       string `json:"전일대비구분 {1:상한 2:상승 3:보합 4:하한 5:하락}"`
	Change     string `json:"전일대비"`
	Drate      string `json:"등락율"`
	Price      string `json:"현재가"`
	Opentime   string `json:"시가시간"`
	Open       string `json:"시가"`
	Hightime   string `json:"고가시간"`
	High       string `json:"고가"`
	Lowtime    string `json:"저가시간"`
	Low        string `json:"저가"`
	Cgubun     string `json:"체결구분{+:매수 -:매도}"`
	Cvolume    string `json:"체결량"`
	Volume     string `json:"누적거래량"`
	Value      string `json:"누적거래대금"`
	Mdvolume   string `json:"매도누적체결량"`
	Mdchecnt   string `json:"매도누적체결건수"`
	Msvolume   string `json:"매수누적체결량"`
	Mschecnt   string `json:"매수누적체결건수"`
	Cpower     string `json:"체결강도"`
	Wavrg      string `json:"가중평균가"`
	Offerho    string `json:"매도호가"`
	Bidho      string `json:"매수호가"`
	Status     string `json:"장정보 {' 0':장중 ' 4':장후시간외 '10':장전시간외}"`
	Jnilvolume string `json:"전일동시간대거래량"`
	Shcode     string `json:"단축코드"`
}

func (s S3OutBlock) ToJSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}
