package wrapper

import (
	"reflect"
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
)

var (
	clsIDXASession = ole.NewGUID("{7FEF321C-6BFD-413C-AA80-541A275434A1}")
	clsIDXAQuery   = ole.NewGUID("{781520A9-4C8C-433B-AA6E-EE9E94108639}")
	clsIDXAReal    = ole.NewGUID("{4D654021-F9D9-49F7-B2F9-6529A19746F7}")

	iIDXASessionEvent = ole.NewGUID("{6D45238D-A5EB-4413-907A-9EA14D046FE5}")
	iIDXAQueryEvent   = ole.NewGUID("{AAF89E20-1F84-4B1F-B6EE-617B6F2C9CD4}")
	iIDXARealEvent    = ole.NewGUID("{16602768-2C96-4D93-984B-E36E7E35BFBE}")
)

const (
	eventXASessionLogin      = 1
	eventXASessionLogOut     = 2
	eventXASessionDisConnect = 3

	eventXAQueryReceiveData           = 1
	eventXAQueryReceiveMessage        = 2
	eventXAQueryReceiveChartRealData  = 3
	eventXAQueryReceiveSearchRealData = 4

	eventXARealReceiveRealData = 1
	eventXARealReceiveLinkData = 2
)

type callbackSession interface {
	ReceivedLogin(*Ebest, XaSessionLogin)
	ReceivedLogout(*Ebest)
	ReceivedDisConnect(*Ebest)
}

type callbackQuery interface {
	ReceivedData(*Ebest, XaQueryReceiveData)
	ReceivedMessage(*Ebest, XaQueryReceiveMessage)
	ReceivedChartRealData(*Ebest, XaQueryReceiveChartRealData)
	ReceivedSearchRealData(*Ebest, XaQueryReceiveSearchRealData)
}

type callbackReal interface {
	ReceivedRealData(*Ebest, XaRealReceiveRealData)
	ReceivedLinkData(*Ebest, XaRealRecieveLinkData)
}

type ebestEvent struct {
	lpVtbl *ole.IDispatchVtbl
	ref    int32
	host   *Ebest
}

type Ebest struct {
	objName string

	unk   *ole.IUnknown
	disp  *ole.IDispatch
	event *ebestEvent

	cbSession callbackSession
	cbQuery   callbackQuery
	cbReal    callbackReal

	point  *ole.IConnectionPoint
	cookie uint32
}

func (e *Ebest) Create(name string) {
	var clsid *ole.GUID
	switch name {
	case "XASession":
		clsid = clsIDXASession
	case "XAQuery":
		clsid = clsIDXAQuery
	case "XAReal":
		clsid = clsIDXAReal
	default:
	}
	e.objName = name

	unk, err := ole.CreateInstance(clsid, ole.IID_IUnknown)
	if err != nil {
		panic(err)
	}
	e.unk = unk

	disp, err := e.unk.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	e.disp = disp
}

func (e *Ebest) Release() {
	if e.unk != nil {
		e.unk.Release()
		e.unk = nil
	}
	if e.disp != nil {
		e.disp.Release()
		e.disp = nil
	}
}

func (e *Ebest) BindEvent(callback interface{}) {
	if e.event == nil {
		e.event = &ebestEvent{}
		e.event.lpVtbl = &ole.IDispatchVtbl{}
		e.event.lpVtbl.QueryInterface = syscall.NewCallback(queryInterface)
		e.event.lpVtbl.AddRef = syscall.NewCallback(addRef)
		e.event.lpVtbl.Release = syscall.NewCallback(release)
		e.event.lpVtbl.GetTypeInfoCount = syscall.NewCallback(getTypeInfoCount)
		e.event.lpVtbl.GetTypeInfo = syscall.NewCallback(getTypeInfo)
		e.event.lpVtbl.GetIDsOfNames = syscall.NewCallback(getIDsOfNames)
		e.event.lpVtbl.Invoke = syscall.NewCallback(invoke)
		e.event.host = e
	}

	unkConn, err := e.disp.QueryInterface(ole.IID_IConnectionPointContainer)
	if err != nil {
		panic(err)
	}

	container := (*ole.IConnectionPointContainer)(unsafe.Pointer(unkConn))

	var iid *ole.GUID
	switch e.objName {
	case "XASession":
		iid = iIDXASessionEvent
		e.cbSession = callback.(callbackSession)
	case "XAQuery":
		iid = iIDXAQueryEvent
		e.cbQuery = callback.(callbackQuery)
	case "XAReal":
		iid = iIDXARealEvent
		e.cbReal = callback.(callbackReal)
	default:
	}

	var point *ole.IConnectionPoint
	err = container.FindConnectionPoint(iid, &point)
	if err != nil {
		panic(err)
	}

	cookie, err := point.Advise((*ole.IUnknown)(unsafe.Pointer(e.event)))
	container.Release()
	if err != nil {
		point.Release()
		panic(err)
	}
	e.point = point
	e.cookie = cookie
}

func (e *Ebest) UnBindEvent() {
	if e.point != nil {
		e.point.Unadvise(e.cookie)
		e.point.Release()
		e.point = nil
		e.cookie = 0
	}
}

func queryInterface(this *ole.IUnknown, iid *ole.GUID, punk **ole.IUnknown) uint32 {
	*punk = nil
	if ole.IsEqualGUID(iid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(iid, ole.IID_IDispatch) ||
		ole.IsEqualGUID(iid, iIDXASessionEvent) ||
		ole.IsEqualGUID(iid, iIDXAQueryEvent) ||
		ole.IsEqualGUID(iid, iIDXARealEvent) {
		addRef(this)
		*punk = this
		return ole.S_OK
	}
	return ole.E_NOINTERFACE
}

func addRef(this *ole.IUnknown) int32 {
	pthis := (*ebestEvent)(unsafe.Pointer(this))
	pthis.ref++
	return pthis.ref
}

func release(this *ole.IUnknown) int32 {
	pthis := (*ebestEvent)(unsafe.Pointer(this))
	pthis.ref--
	return pthis.ref
}

func getTypeInfoCount(pcount *int) uintptr {
	if pcount != nil {
		*pcount = 0
	}
	return uintptr(ole.S_OK)
}

func getTypeInfo(ptypeif *uintptr) uintptr {
	return uintptr(ole.E_NOTIMPL)
}

func getIDsOfNames(args *uintptr) uint32 {
	p := (*[6]int32)(unsafe.Pointer(args))
	//this := (*ole.IDispatch)(unsafe.Pointer(uintptr(p[0])))
	//iid := (*ole.GUID)(unsafe.Pointer(uintptr(p[1])))
	//wnames := *(*[]*uint16)(unsafe.Pointer(uintptr(p[2])))
	namelen := int(uintptr(p[3]))
	//lcid := int(uintptr(p[4]))
	pdisp := *(*[]int32)(unsafe.Pointer(uintptr(p[5])))
	for n := 0; n < namelen; n++ {
		pdisp[n] = int32(n)
	}
	return ole.S_OK
}

type dispParams struct {
	rgvarg            uintptr
	rgdispidNamedArgs uintptr
	cArgs             uint32
	cNamedArgs        uint32
}

func invoke(this *ole.IDispatch, dispid int, riid *ole.GUID, lcid int, flags int16, dispparams *ole.DISPPARAMS, result *ole.VARIANT, pexcepinfo *ole.EXCEPINFO, nerr *uint) uintptr {
	dispp := (*dispParams)(unsafe.Pointer(dispparams))
	vv := reflect.SliceHeader{Data: dispp.rgvarg, Len: int(dispp.cArgs), Cap: int(dispp.cArgs)}
	v := *(*[]ole.VARIANT)(unsafe.Pointer(&vv))

	pthis := (*ebestEvent)(unsafe.Pointer(this))

	switch pthis.host.objName {
	case "XASession":
		switch dispid {
		case eventXASessionLogin:
			x := XaSessionLogin{Code: v[1].Value().(string), Msg: v[0].Value().(string)}

			if pthis.host.cbSession != nil {
				pthis.host.cbSession.ReceivedLogin(pthis.host, x)
			}
		case eventXASessionLogOut:
			if pthis.host.cbSession != nil {
				pthis.host.cbSession.ReceivedLogout(pthis.host)
			}
		case eventXASessionDisConnect:
			if pthis.host.cbSession != nil {
				pthis.host.cbSession.ReceivedDisConnect(pthis.host)
			}
		}
	case "XAQuery":
		switch dispid {
		case eventXAQueryReceiveData:
			if pthis.host.cbQuery != nil {
				pthis.host.cbQuery.ReceivedData(pthis.host, XaQueryReceiveData(v[0].ToString()))
			}
		case eventXAQueryReceiveMessage:
			x := XaQueryReceiveMessage{IsSystemError: int(v[2].Value().(int16)), MessageCode: v[1].Value().(string), Message: v[0].Value().(string)}

			if pthis.host.cbQuery != nil {
				pthis.host.cbQuery.ReceivedMessage(pthis.host, x)
			}
		case eventXAQueryReceiveChartRealData:
			if pthis.host.cbQuery != nil {
				pthis.host.cbQuery.ReceivedChartRealData(pthis.host, XaQueryReceiveChartRealData(v[0].ToString()))
			}
		case eventXAQueryReceiveSearchRealData:
			if pthis.host.cbQuery != nil {
				pthis.host.cbQuery.ReceivedSearchRealData(pthis.host, XaQueryReceiveSearchRealData(v[0].ToString()))
			}
		}
	case "XAReal":
		switch dispid {
		case eventXARealReceiveRealData:
			if pthis.host.cbReal != nil {
				pthis.host.cbReal.ReceivedRealData(pthis.host, XaRealReceiveRealData(v[0].ToString()))
			}

		case eventXARealReceiveLinkData:
			x := XaRealRecieveLinkData{LinkName: v[2].Value().(string), Data: v[1].Value().(string), Filler: v[0].Value().(string)}
			if pthis.host.cbReal != nil {
				pthis.host.cbReal.ReceivedLinkData(pthis.host, x)
			}
		}
	default:
	}
	return ole.E_NOTIMPL
}
