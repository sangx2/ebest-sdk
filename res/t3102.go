package res

import (
	"encoding/json"
)

// T3102 뉴스본문
type T3102InBlock struct {
	SNewsno string `json:"뉴스번호"`
}

type T3102OutBlock struct {
	SJongcode string `json:"뉴스종목"`
}

func (t T3102OutBlock) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type T3102OutBlock1 struct {
	SBody string `json:"뉴스본문"`
}

func (t T3102OutBlock1) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type T3102OutBlock2 struct {
	STitle string `json:"뉴스타이틀"`
}

func (t T3102OutBlock2) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}
