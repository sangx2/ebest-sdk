package ebest

import "fmt"

// ErrSession Session 에러
type ErrSession struct {
	FuncName string
	Code     string
	Msg      string
}

func NewErrSession(funcName, code, msg string) *ErrSession {
	return &ErrSession{
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}

func (e *ErrSession) Error() string {
	return fmt.Sprintf("Session Function: %s, Code: %s, Msg: %s", e.FuncName, e.Code, e.Msg)
}

// ErrQuery Query 에러
type ErrQuery struct {
	FuncName string
	Code     string
	Msg      string
}

func NewErrQuery(funcName, code, msg string) *ErrQuery {
	return &ErrQuery{
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}

func (e *ErrQuery) Error() string {
	return fmt.Sprintf("Query Function: %s, Code: %s, Msg: %s", e.FuncName, e.Code, e.Msg)
}

// ErrReal Real 에러
type ErrReal struct {
	FuncName string
	Msg      string
}

func NewErrReal(funcName, msg string) *ErrReal {
	return &ErrReal{
		FuncName: funcName,
		Msg:      msg,
	}
}

func (e *ErrReal) Error() string {
	return fmt.Sprintf("Real Function: %s, Msg: %s", e.FuncName, e.Msg)
}
