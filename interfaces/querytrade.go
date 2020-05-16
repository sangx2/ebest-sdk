package interfaces

import "github.com/sangx2/ebest/wrapper"

type QueryTrade interface {
	GetTPS() int
	GetLPP() int

	// event
	GetReceiveDataChan() chan wrapper.XaQueryReceiveData
	GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage
	GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData
	GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData

	// in/out block
	SetFieldData(eBest *wrapper.Ebest, resPath string, inBlocks ...interface{}) error
	GetOutBlocks() []interface{}

	// callback func
	ReceivedData(*wrapper.Ebest, wrapper.XaQueryReceiveData)
	ReceivedMessage(*wrapper.Ebest, wrapper.XaQueryReceiveMessage)
	ReceivedChartRealData(*wrapper.Ebest, wrapper.XaQueryReceiveChartRealData)
	ReceivedSearchRealData(*wrapper.Ebest, wrapper.XaQueryReceiveSearchRealData)
}
