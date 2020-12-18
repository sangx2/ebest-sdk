package wrapper

import (
	"github.com/go-ole/go-ole/oleutil"
)

// Event
type XaRealReceiveRealData = string

type XaRealReceiveLinkData struct {
	LinkName string
	Data     string
	Filler   string
}

// Property
/*
func (e *EBest) ResFileName(resFileName string) (interface{}) {
	r := oleutil.MustPutProperty(e.disp, "ResFileName", resFileName)
	return r.Value()
}
*/

// Method
func (e *EBestWrapper) AdviseRealData() {
	_ = oleutil.MustCallMethod(e.disp, "AdviseRealData")
}

func (e *EBestWrapper) UnAdviseRealData() {
	_ = oleutil.MustCallMethod(e.disp, "UnAdviseRealData")
}

func (e *EBestWrapper) UnAdviseRealDataWithKey(code string) {
	_ = oleutil.MustCallMethod(e.disp, "UnAdviseRealDataWithKey", code)
}

func (e *EBestWrapper) AdviseLinkFromHTS() {
	_ = oleutil.MustCallMethod(e.disp, "AdviseLinkFromHTS")
}

func (e *EBestWrapper) UnadviseLinkFromHTS() {
	_ = oleutil.MustCallMethod(e.disp, "UnadviseLinkFromHTS")
}

/*
func (e *EBest) GetFieldData(blockName string, fieldName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetFieldData", blockName, fieldName)
	return r.Value().(string)
}

func (e *EBest) SetFieldData(blockName string, fieldName string, data string) {
	_ = oleutil.MustCallMethod(e.disp, "SetFieldData", blockName, fieldName, data)
}

func (e *EBest) LoadFromResFile(fileName string) bool {
	r := oleutil.MustCallMethod(e.disp, "LoadFromResFile", fileName)
	return r.Value().(bool)
}

func (e *EBest) GetBlockData(blockName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetBlockdata", blockName)
	return r.Value().(string)
}
*/
