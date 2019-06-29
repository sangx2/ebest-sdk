package wrapper

import (
	"github.com/go-ole/go-ole/oleutil"
)

// Event
type XaSessionLogin struct {
	Code string
	Msg  string
}

// Property
func (e *Ebest) SendPacketSize(dataSize int) interface{} {
	r := oleutil.MustPutProperty(e.disp, "SendPacketSize", dataSize)
	return r.Value()
}

func (e *Ebest) ConnectTimeOut(timeout int) interface{} {
	r := oleutil.MustPutProperty(e.disp, "ConnectTimeOut", timeout)
	return r.Value()
}

// Method
func (e *Ebest) ConnectServer(serverIP string, serverPort int) bool {
	r := oleutil.MustCallMethod(e.disp, "ConnectServer", serverIP, serverPort)
	return r.Value().(bool)
}

func (e *Ebest) DisconnectServer() {
	_ = oleutil.MustCallMethod(e.disp, "DisconnectServer")
}

func (e *Ebest) IsConnected() bool {
	r := oleutil.MustCallMethod(e.disp, "IsConnected")
	return r.Value().(bool)
}

func (e *Ebest) Login(id string, pwd string, certPwd string) bool {
	r := oleutil.MustCallMethod(e.disp, "Login", id, pwd, certPwd, 0, false)
	return r.Value().(bool)
}

func (e *Ebest) GetAccountListCount() int32 {
	r := oleutil.MustCallMethod(e.disp, "GetAccountListCount")
	return r.Value().(int32)
}

func (e *Ebest) GetAccountList(index int) string {
	r := oleutil.MustCallMethod(e.disp, "GetAccountList", index)
	return r.Value().(string)
}

func (e *Ebest) GetAccountName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAccountName", acc)
	return r.Value().(string)
}

func (e *Ebest) GetAcctDetailName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAcctDetailName", acc)
	return r.Value().(string)
}

func (e *Ebest) GetAcctNickName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAcctNickName", acc)
	return r.Value().(string)
}

func (e *Ebest) GetLastError() int32 {
	r := oleutil.MustCallMethod(e.disp, "GetLastError")
	return r.Value().(int32)
}

func (e *Ebest) GetErrorMessage(errorCode int) string {
	r := oleutil.MustCallMethod(e.disp, "GetErrorMessage", errorCode)
	return r.Value().(string)
}

func (e *Ebest) IsLoadAPI() bool {
	r := oleutil.MustCallMethod(e.disp, "IsLoadAPI")
	return r.Value().(bool)
}

func (e *Ebest) GetServerName() bool {
	r := oleutil.MustCallMethod(e.disp, "GetServerName")
	return r.Value().(bool)
}
