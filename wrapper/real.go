package wrapper

import (
	"github.com/go-ole/go-ole/oleutil"
)

// Event
type XaRealReceiveRealData = string

type XaRealRecieveLinkData struct {
	LinkName string
	Data     string
	Filler   string
}

// Property
/*
func (e *Ebest) ResFileName(resFileName string) (interface{}) {
	r := oleutil.MustPutProperty(e.disp, "ResFileName", resFileName)
	return r.Value()
}
*/

// Method
func (e *Ebest) AdviseRealData() {
	_ = oleutil.MustCallMethod(e.disp, "AdviseRealData")
}

func (e *Ebest) UnAdviseRealData() {
	_ = oleutil.MustCallMethod(e.disp, "UnAdviseRealData")
}

func (e *Ebest) UnAdviseRealDataWithKey(code string) {
	_ = oleutil.MustCallMethod(e.disp, "UnAdviseRealDataWithKey", code)
}

func (e *Ebest) AdviseLinkFromHTS() {
	_ = oleutil.MustCallMethod(e.disp, "AdviseLinkFromHTS")
}

func (e *Ebest) UnadviseLinkFromHTS() {
	_ = oleutil.MustCallMethod(e.disp, "UnadviseLinkFromHTS")
}

/*
func (e *Ebest) GetFieldData(blockName string, fieldName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetFieldData", blockName, fieldName)
	return r.Value().(string)
}

func (e *Ebest) SetFieldData(blockName string, fieldName string, data string) {
	_ = oleutil.MustCallMethod(e.disp, "SetFieldData", blockName, fieldName, data)
}

func (e *Ebest) LoadFromResFile(fileName string) bool {
	r := oleutil.MustCallMethod(e.disp, "LoadFromResFile", fileName)
	return r.Value().(bool)
}

func (e *Ebest) GetBlockData(blockName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetBlockdata", blockName)
	return r.Value().(string)
}
*/
