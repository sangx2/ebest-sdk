package wrapper

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// Event
type XaQueryReceiveData = string

type XaQueryReceiveMessage struct {
	IsSystemError int
	MessageCode   string
	Message       string
}

type XaQueryReceiveChartRealData = string

type XaQueryReceiveSearchRealData = string

// Property
func (e *EBestWrapper) ResFileName(resFileName string) interface{} {
	r := oleutil.MustPutProperty(e.disp, "ResFileName", resFileName)
	return r.Value()
}

func (e *EBestWrapper) IsNext() interface{} {
	r := oleutil.MustPutProperty(e.disp, "IsNext")
	return r.Value()
}

// Method
func (e *EBestWrapper) Request(next bool) int32 {
	r := oleutil.MustCallMethod(e.disp, "Request", next)
	return r.Value().(int32)
}

func (e *EBestWrapper) GetFieldData(blockName string, fieldName string, occursIndex int) string {
	var r *ole.VARIANT

	switch e.objName {
	case "XAQuery":
		r = oleutil.MustCallMethod(e.disp, "GetFieldData", blockName, fieldName, occursIndex)
	case "XAReal":
		r = oleutil.MustCallMethod(e.disp, "GetFieldData", blockName, fieldName)
	}
	return r.Value().(string)
}

func (e *EBestWrapper) SetFieldData(blockName string, fieldName string, occursIndex int, data string) {
	switch e.objName {
	case "XAQuery":
		_ = oleutil.MustCallMethod(e.disp, "SetFieldData", blockName, fieldName, occursIndex, data)
	case "XAReal":
		_ = oleutil.MustCallMethod(e.disp, "SetFieldData", blockName, fieldName, data)
	}
}

func (e *EBestWrapper) GetBlockCount(blockName string) int32 {
	r := oleutil.MustCallMethod(e.disp, "GetBlockCount", blockName)
	return r.Value().(int32)
}

func (e *EBestWrapper) SetBlockCount(blockName string, count int) {
	_ = oleutil.MustCallMethod(e.disp, "SetBlockCount", blockName, count)
}

func (e *EBestWrapper) LoadFromResFile(fileName string) bool {
	r := oleutil.MustCallMethod(e.disp, "LoadFromResFile", fileName)
	return r.Value().(bool)
}

func (e *EBestWrapper) ClearBlockdata(blockName string) {
	_ = oleutil.MustCallMethod(e.disp, "ClearBlockdata", blockName)
}

func (e *EBestWrapper) GetBlockData(blockName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetBlockdata", blockName)
	return r.Value().(string)
}

func (e *EBestWrapper) GetTRCountPerSec(code string) int {
	r := oleutil.MustCallMethod(e.disp, "GetTRCountPerSec", code)
	return r.Value().(int)
}

func (e *EBestWrapper) RequestService(code string, data string) int32 {
	r := oleutil.MustCallMethod(e.disp, "RequestService", code, data)
	return r.Value().(int32)
}

func (e *EBestWrapper) RemoveService(code string, data string) int32 {
	r := oleutil.MustCallMethod(e.disp, "RemoveService", code, data)
	return r.Value().(int32)
}

func (e *EBestWrapper) RequestLinkToHTS(linkName string, data string, filler string) bool {
	r := oleutil.MustCallMethod(e.disp, "RequestLinkToHTS", linkName, data, filler)
	return r.Value().(bool)
}

func (e *EBestWrapper) Decompress(blockName string) int32 {
	r := oleutil.MustCallMethod(e.disp, "Decompress", blockName)
	return r.Value().(int32)
}

func (e *EBestWrapper) GetFieldChartRealData(blockName string, fieldName string) string {
	r := oleutil.MustCallMethod(e.disp, "GetFieldChartRealData", blockName, fieldName)
	return r.Value().(string)
}

func (e *EBestWrapper) GetAttribute(blockName string, fieldName string, attribute string, occursIndex int) string {
	r := oleutil.MustCallMethod(e.disp, "GetAttribute", blockName, fieldName, attribute, occursIndex)
	return r.Value().(string)
}

func (e *EBestWrapper) GetTRCountBaseSec(code string) int {
	r := oleutil.MustCallMethod(e.disp, "GetTRCountBaseSec", code)
	return r.Value().(int)
}

func (e *EBestWrapper) GetTRCountRequest(code string) int32 {
	r := oleutil.MustCallMethod(e.disp, "GetTRCountRequest", code)
	return r.Value().(int32)
}

func (e *EBestWrapper) GetTRCountLimit(code string) int {
	r := oleutil.MustCallMethod(e.disp, "GetTRCountLimit", code)
	return r.Value().(int)
}
