package res

import (
	"encoding/json"
)

// T3320 FNG 요약
type T3320InBlock struct {
	Gicode string `json:"종목코드"`
}

type T3320OutBlock struct {
	Upgubunnm    string `json:"업종구분명"`
	Sijangcd     string `json:"시장구분"`
	Marketnm     string `json:"시장구분명"`
	Company      string `json:"한글기업명"`
	Baddress     string `json:"본사주소"`
	Btelno       string `json:"본사전화번호"`
	Gsyyyy       string `json:"최근결산년도"`
	Gsmm         string `json:"결산월"`
	Gsym         string `json:"최근결산년월"`
	Lstprice     string `json:"주당액면가"`
	Gstock       string `json:"주식수"`
	Homeurl      string `json:"Homepage"`
	Grdnm        string `json:"그룹명"`
	Foreignratio string `json:"외국인"`
	Irtel        string `json:"주담전화"`
	Capital      string `json:"자본금"`
	Sigavalue    string `json:"시가총액"`
	Cashsis      string `json:"배당금"`
	Cashrate     string `json:"배당수익율"`
	Price        string `json:"현재가"`
	Jnilclose    string `json:"전일종가"`
}

func (t T3320OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type T3320OutBlock1 struct {
	Gicode   string `json:"기업코드"`
	Gsym     string `json:"결산년월"`
	Gsgb     string `json:"결산구분"`
	Per      string `json:"PER"`
	Eps      string `json:"EPS"`
	Pbr      string `json:"PBR"`
	Roa      string `json:"ROA"`
	Roe      string `json:"ROE"`
	Ebitda   string `json:"EBITDA"`
	Evebitda string `json:"EVEBITDA"`
	Par      string `json:"액면가"`
	Sps      string `json:"SPS"`
	Cps      string `json:"CPS"`
	Bps      string `json:"BPS"`
	TPer     string `json:"T.PER"`
	TEps     string `json:"T.EPS"`
	Peg      string `json:"PEG"`
	TPeg     string `json:"T.PEG"`
	TGsym    string `json:"최근분기년도"`
}

func (t T3320OutBlock1) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
