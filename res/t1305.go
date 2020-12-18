package res

import (
	"encoding/json"
)

// T1511 기간별 주가
type T1305InBlock struct {
	Shcode  string `json:"업종코드"`
	Dwmcode string `json:"일주월구분"`
	Date    string `json:"날짜"`
	Idx     string `json:"IDX"`
	Cnt     string `json:"건수"`
}

type T1305OutBlock struct {
	Cnt  string `json:"CNT"`
	Date string `json:"날짜"`
	Idx  string `json:"IDX"`
}

func (t T1305OutBlock) ToJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}

type T1305OutBlock1 struct {
	Date       string `json:"날짜"`
	Open       string `json:"시가"`
	High       string `json:"고가"`
	Low        string `json:"저가"`
	Close      string `json:"종가"`
	Sign       string `json:"전일대비구분"`
	Change     string `json:"전일대비"`
	Diff       string `json:"등락율"`
	Volume     string `json:"누적거래량"`
	Diff_vol   string `json:"거래증가율"`
	Chdegree   string `json:"체결강도"`
	Sojinrate  string `json:"소진율"`
	Changerate string `json:"회전율"`
	Fpvolume   string `json:"외인순매수"`
	Covolume   string `json:"기관순매수"`
	Shcode     string `json:"종목코드"`
	Value      string `json:"누적거래대금(단위:백만)"`
	Ppvolume   string `json:"개인순매수"`
	O_sign     string `json:"시가대비구분"`
	O_change   string `json:"시가대비"`
	O_diff     string `json:"시가기준등락율"`
	H_sign     string `json:"고가대비구분"`
	H_change   string `json:"고가대비"`
	H_diff     string `json:"고기기준등락율"`
	L_sign     string `json:"저가대비구분"`
	L_change   string `json:"저가대비"`
	L_diff     string `json:"저가기준등락율"`
	Marketcap  string `json:"시가총액(단위:백만)"`
}

func (t T1305OutBlock1) ToJSON() []byte {
	b, _ := json.Marshal(t)
	return b
}
