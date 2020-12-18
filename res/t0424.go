package res

import (
	"encoding/json"
)

// T0424 주식 잔고2
type T0424InBlock struct {
	Accno      string `json:"계좌번호"`
	Passwd     string `json:"비밀번호"`
	Prcgb      string `json:"단가구분"`     // {1:평균단가,2:BEP단가}
	Chegb      string `json:"체결구분"`     // {0:결제기준잔고,2:체결기준(잔고가0이아닌종목만조회)}
	Dangb      string `json:"단일가구분"`    // {0:정규장,1:시간외단일가}
	Charge     string `json:"제비용포함여부"`  // {0:제비용미포함,1:제비용포함}
	CtsExpcode string `json:"CTS_종목번호"` // {처음조회시는Space,연속조회시에이전조회한OutBlock의cts_expcode값으로설정}
}

type T0424OutBlock struct {
	Sunamt     string `json:"순추정자산"`
	Dtsunik    string `json:"실현손익"`
	Mamt       string `json:"매입금액"`
	Sunamt1    string `json:"추정D2예수금"`
	CtsExpcode string `json:"CTS_종목번호"`
	Tappamt    string `json:"평가금액"`
	Tdtsunik   string `json:"평가손익"`
}

func (t T0424OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type T0424OutBlock1 struct {
	Expcode    string `json:"종목번호"`
	Jangb      string `json:"잔고구분"`
	Janqty     string `json:"잔고수량"`
	Mdposqt    string `json:"매도가능수량"`
	Pamt       string `json:"평균단가"`
	Mamt       string `json:"매입금액"`
	Sinamt     string `json:"대출금액"`
	Lastdt     string `json:"만기일자"`
	Msat       string `json:"당일매수금액"`
	Mpms       string `json:"당일매수단가"`
	Mdat       string `json:"당일매도금액"`
	Mpmd       string `json:"당일매도단가"`
	Jsat       string `json:"전일매수금액"`
	Jpms       string `json:"전일매수단가"`
	Jdat       string `json:"전일매도금액"`
	Jpmd       string `json:"전일매도단가"`
	Sysprocseq string `json:"처리순번"`
	Loandt     string `json:"대출일자"`
	Hname      string `json:"종목명"`
	Marketgb   string `json:"시장구분"` // {1:주식,2:채권}
	Jonggb     string `json:"종목구분"` // {1:프리보드,2:코스닥,3:거래소,Z:상장폐지,:비상장,8:코넥스,9:CMA&RP}
	Janrt      string `json:"보유비중"`
	Price      string `json:"현재가"`
	Appamt     string `json:"평가금액"`
	Dtsunik    string `json:"평가손익"`
	Sunikrt    string `json:"수익율"`
	Fee        string `json:"수수료"`
	Tax        string `json:"제세금"`
	Sininter   string `json:"신용이자"`
}

func (t T0424OutBlock1) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
