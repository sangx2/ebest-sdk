package interfaces

import "github.com/sangx2/ebest-sdk/wrapper"

type QueryTrade interface {
	GetTPS() int
	GetLPP() int

	// event
	GetReceiveDataChan() chan wrapper.XaQueryReceiveData
	GetReceiveMessageChan() chan wrapper.XaQueryReceiveMessage
	GetReceiveChartRealDataChan() chan wrapper.XaQueryReceiveChartRealData
	GetReceiveChartSearchRealDataChan() chan wrapper.XaQueryReceiveSearchRealData

	// in/out block
	SetFieldData(eBest *wrapper.EBestWrapper, resPath string, inBlocks ...interface{}) error
	GetOutBlocks() []interface{}
	GetBlockData(eBest *wrapper.EBestWrapper, blockName string) string

	// callback func
	ReceivedData(*wrapper.EBestWrapper, wrapper.XaQueryReceiveData)
	ReceivedMessage(*wrapper.EBestWrapper, wrapper.XaQueryReceiveMessage)
	ReceivedChartRealData(*wrapper.EBestWrapper, wrapper.XaQueryReceiveChartRealData)
	ReceivedSearchRealData(*wrapper.EBestWrapper, wrapper.XaQueryReceiveSearchRealData)
}
