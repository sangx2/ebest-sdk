package interfaces

import "github.com/sangx2/ebest-sdk/wrapper"

type RealTrade interface {
	// event
	GetReceivedRealDataChan() chan wrapper.XaRealReceiveRealData
	GetReceivedLinkDataChan() chan wrapper.XaRealReceiveLinkData

	// in/out block
	SetFieldData(ebest *wrapper.EBestWrapper, resPath string, inBlock interface{}) error
	GetOutBlock() interface{}
	GetBlockData(eBest *wrapper.EBestWrapper, blockName string) string

	// callback func
	ReceivedRealData(*wrapper.EBestWrapper, wrapper.XaRealReceiveRealData)
	ReceivedLinkData(*wrapper.EBestWrapper, wrapper.XaRealReceiveLinkData)
}
