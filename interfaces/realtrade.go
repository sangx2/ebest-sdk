package interfaces

import "github.com/sangx2/ebest/wrapper"

type RealTrade interface {
	// event
	GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData
	GetReceivedLinkDataChan() chan wrapper.XaRealRecieveLinkData

	// in/out block
	SetFieldData(ebest *wrapper.Ebest, resPath string, inBlock interface{}) error
	GetOutBlock() interface{}

	// callback func
	ReceivedRealData(*wrapper.Ebest, wrapper.XaRealReceiveRealData)
	ReceivedLinkData(*wrapper.Ebest, wrapper.XaRealRecieveLinkData)
}
