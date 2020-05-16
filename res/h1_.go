package res

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// H1 KOSPI호가잔량
type H1InBlock struct {
	Shcode string `json:"단축코드"`
}

type H1OutBlock struct {
	gorm.Model

	Hotime      string `json:"호가시간"`
	Offerho1    string `json:"매도호가1"`
	Offerrem1   string `json:"매도호가잔량1"`
	Bidho1      string `json:"매수호가1"`
	Bidrem1     string `json:"매수호가잔량1"`
	Offerho2    string `json:"매도호가2"`
	Offerrem2   string `json:"매도호가잔량2"`
	Bidho2      string `json:"매수호가2"`
	Bidrem2     string `json:"매수호가잔량2"`
	Offerho3    string `json:"매도호가3"`
	Offerrem3   string `json:"매도호가잔량3"`
	Bidho3      string `json:"매수호가3"`
	Bidrem3     string `json:"매수호가잔량3"`
	Offerho4    string `json:"매도호가4"`
	Offerrem4   string `json:"매도호가잔량4"`
	Bidho4      string `json:"매수호가4"`
	Bidrem4     string `json:"매수호가잔량4"`
	Offerho5    string `json:"매도호가5"`
	Offerrem5   string `json:"매도호가잔량5"`
	Bidho5      string `json:"매수호가5"`
	Bidrem5     string `json:"매수호가잔량5"`
	Offerho6    string `json:"매도호가6"`
	Offerrem6   string `json:"매도호가잔량6"`
	Bidho6      string `json:"매수호가6"`
	Bidrem6     string `json:"매수호가잔량6"`
	Offerho7    string `json:"매도호가7"`
	Offerrem7   string `json:"매도호가잔량7"`
	Bidho7      string `json:"매수호가7"`
	Bidrem7     string `json:"매수호가잔량7"`
	Offerho8    string `json:"매도호가8"`
	Offerrem8   string `json:"매도호가잔량8"`
	Bidho8      string `json:"매수호가8"`
	Bidrem8     string `json:"매수호가잔량8"`
	Offerho9    string `json:"매도호가9"`
	Offerrem9   string `json:"매도호가잔량9"`
	Bidho9      string `json:"매수호가9"`
	Bidrem9     string `json:"매수호가잔량9"`
	Offerho10   string `json:"매도호가10"`
	Offerrem10  string `json:"매도호가잔량10"`
	Bidho10     string `json:"매수호가10"`
	Bidrem10    string `json:"매수호가잔량10"`
	Totofferrem string `json:"총매도호가잔량"`
	Totbidrem   string `json:"총매수호가잔량"`
	Donsigubun  string `json:"동시호가구분"`
	Shcode      string `json:"단축코드"`
	Allocgubun  string `json:"배분적용구분"`
}

func (h H1OutBlock) ToJSON() string {
	b, _ := json.Marshal(h)
	return string(b)
}
