package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// NWS 실시간 뉴스 제목 패킷
type NWSInBlock struct {
	NWcode string `json:"뉴스코드"`
}

type NWSOutBlock struct {
	gorm.Model

	Date      string `json:"날짜"`
	Time      string `json:"시간"`
	Publisher string `json:"뉴스구분자 {11:연합 뉴스 14:이데일리 15:공시 20:머니투데이 21:인포스탁 22:이트레이드 23:아이아경제 24:뉴스핌 25:매일경제 26:한국경제}"`
	Realkey   string `json:"키값"`
	Title     string `json:"제목"`
	Code      string `json:"단축종목코드"`
	Bodysize  string `json:"BODY길이" gorm:"-"`
}

func (n NWSOutBlock) ToJSON() string {
	b, _ := json.Marshal(n)
	return string(b)
}
