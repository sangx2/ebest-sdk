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
func (e *EBestWrapper) SendPacketSize(dataSize int) interface{} {
	r := oleutil.MustPutProperty(e.disp, "SendPacketSize", dataSize)
	return r.Value()
}

func (e *EBestWrapper) ConnectTimeOut(timeout int) interface{} {
	r := oleutil.MustPutProperty(e.disp, "ConnectTimeOut", timeout)
	return r.Value()
}

// Method
func (e *EBestWrapper) ConnectServer(serverIP string, serverPort int) bool {
	r := oleutil.MustCallMethod(e.disp, "ConnectServer", serverIP, serverPort)
	return r.Value().(bool)
}

func (e *EBestWrapper) DisconnectServer() {
	_ = oleutil.MustCallMethod(e.disp, "DisconnectServer")
}

func (e *EBestWrapper) IsConnected() bool {
	r := oleutil.MustCallMethod(e.disp, "IsConnected")
	return r.Value().(bool)
}

func (e *EBestWrapper) Login(id string, pwd string, certPwd string) bool {
	r := oleutil.MustCallMethod(e.disp, "Login", id, pwd, certPwd, 0, false)
	return r.Value().(bool)
}

func (e *EBestWrapper) GetAccountListCount() int32 {
	r := oleutil.MustCallMethod(e.disp, "GetAccountListCount")
	return r.Value().(int32)
}

func (e *EBestWrapper) GetAccountList(index int) string {
	r := oleutil.MustCallMethod(e.disp, "GetAccountList", index)
	return r.Value().(string)
}

func (e *EBestWrapper) GetAccountName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAccountName", acc)
	return r.Value().(string)
}

func (e *EBestWrapper) GetAcctDetailName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAcctDetailName", acc)
	return r.Value().(string)
}

func (e *EBestWrapper) GetAcctNickName(acc string) string {
	r := oleutil.MustCallMethod(e.disp, "GetAcctNickName", acc)
	return r.Value().(string)
}

func (e *EBestWrapper) GetLastError() int32 {
	r := oleutil.MustCallMethod(e.disp, "GetLastError")
	return r.Value().(int32)
}

func (e *EBestWrapper) GetErrorMessage(errorCode int) string {
	r := oleutil.MustCallMethod(e.disp, "GetErrorMessage", errorCode)
	return r.Value().(string)
}

func (e *EBestWrapper) IsLoadAPI() bool {
	r := oleutil.MustCallMethod(e.disp, "IsLoadAPI")
	return r.Value().(bool)
}

func (e *EBestWrapper) GetServerName() bool {
	r := oleutil.MustCallMethod(e.disp, "GetServerName")
	return r.Value().(bool)
}
