package interfaces

import "github.com/sangx2/ebestsdk/wrapper"

type RealTrade interface {
	// event
	GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData
	GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData

	// in/out block
	SetFieldData(ebest *wrapper.EBestWrapper, resPath string, inBlock interface{}) error
	GetOutBlock() interface{}

	// callback func
	ReceivedRealData(*wrapper.EBestWrapper, wrapper.XaRealReceiveRealData)
	ReceivedLinkData(*wrapper.EBestWrapper, wrapper.XaRealReceiveLinkData)
}
