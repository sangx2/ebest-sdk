package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// T1101 주식 현재가 호가 조회
type T1101InBlock struct {
	Shcode string `json:"단축코드"`
}

type T1101OutBlock struct {
	gorm.Model

	Hname         string `json:"한글명"`
	Price         string `json:"현재가"`
	Sign          string `json:"전일대비구분"`
	Change        string `json:"전일대비"`
	Diff          string `json:"등락율"`
	Volume        string `json:"누적거래량"`
	Jnilclose     string `json:"전일종가"`
	Offerho1      string `json:"매도호가1"`
	Bidho1        string `json:"매수호가1"`
	Offerrem1     string `json:"매도호가수량1"`
	Bidrem1       string `json:"매수호가수량1"`
	Preoffercha1  string `json:"직전매도대비수량1"`
	Prebidcha1    string `json:"직전매수대비수량1"`
	Offerho2      string `json:"매도호가2"`
	Bidho2        string `json:"매수호가2"`
	Offerrem2     string `json:"매도호가수량2"`
	Bidrem2       string `json:"매수호가수량2"`
	Preoffercha2  string `json:"직전매도대비수량2"`
	Prebidcha2    string `json:"직전매수대비수량2"`
	Offerho3      string `json:"매도호가3"`
	Bidho3        string `json:"매수호가3"`
	Offerrem3     string `json:"매도호가수량3"`
	Bidrem3       string `json:"매수호가수량3"`
	Preoffercha3  string `json:"직전매도대비수량3"`
	Prebidcha3    string `json:"직전매수대비수량3"`
	Offerho4      string `json:"매도호가4"`
	Bidho4        string `json:"매수호가4"`
	Offerrem4     string `json:"매도호가수량4"`
	Bidrem4       string `json:"매수호가수량4"`
	Preoffercha4  string `json:"직전매도대비수량4"`
	Prebidcha4    string `json:"직전매수대비수량4"`
	Offerho5      string `json:"매도호가5"`
	Bidho5        string `json:"매수호가5"`
	Offerrem5     string `json:"매도호가수량5"`
	Bidrem5       string `json:"매수호가수량5"`
	Preoffercha5  string `json:"직전매도대비수량5"`
	Prebidcha5    string `json:"직전매수대비수량5"`
	Offerho6      string `json:"매도호가6"`
	Bidho6        string `json:"매수호가6"`
	Offerrem6     string `json:"매도호가수량6"`
	Bidrem6       string `json:"매수호가수량6"`
	Preoffercha6  string `json:"직전매도대비수량6"`
	Prebidcha6    string `json:"직전매수대비수량6"`
	Offerho7      string `json:"매도호가7"`
	Bidho7        string `json:"매수호가7"`
	Offerrem7     string `json:"매도호가수량7"`
	Bidrem7       string `json:"매수호가수량7"`
	Preoffercha7  string `json:"직전매도대비수량7"`
	Prebidcha7    string `json:"직전매수대비수량7"`
	Offerho8      string `json:"매도호가8"`
	Bidho8        string `json:"매수호가8"`
	Offerrem8     string `json:"매도호가수량8"`
	Bidrem8       string `json:"매수호가수량8"`
	Preoffercha8  string `json:"직전매도대비수량8"`
	Prebidcha8    string `json:"직전매수대비수량8"`
	Offerho9      string `json:"매도호가9"`
	Bidho9        string `json:"매수호가9"`
	Offerrem9     string `json:"매도호가수량9"`
	Bidrem9       string `json:"매수호가수량9"`
	Preoffercha9  string `json:"직전매도대비수량9"`
	Prebidcha9    string `json:"직전매수대비수량9"`
	Offerho10     string `json:"매도호가10"`
	Bidho10       string `json:"매수호가10"`
	Offerrem10    string `json:"매도호가수량10"`
	Bidrem10      string `json:"매수호가수량10"`
	Preoffercha10 string `json:"직전매도대비수량10"`
	Prebidcha10   string `json:"직전매수대비수량10"`
	Offer         string `json:"매도호가수량합"`
	Bid           string `json:"매수호가수량합"`
	Preoffercha   string `json:"직전매도대비수량합"`
	Prebidcha     string `json:"직전매수대비수량합"`
	Hotime        string `json:"수신시간"`
	Yeprice       string `json:"예상체결가격"`
	Yevolume      string `json:"예상체결수량"`
	Yesign        string `json:"예상체결전일구분"`
	Yechange      string `json:"예상체결전일대비"`
	Yediff        string `json:"예상체결등락율"`
	Tmoffer       string `json:"시간외매도잔량"`
	Tmbid         string `json:"시간외매수잔량"`
	Ho_status     string `json:"동시구분"`
	Shcode        string `json:"단축코드"`
	Uplmtprice    string `json:"상한가"`
	Dnlmtprice    string `json:"하한가"`
	Open          string `json:"시가"`
	High          string `json:"고가"`
	Low           string `json:"저가"`
}

func (t T1101OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
