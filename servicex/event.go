// Package servicex
// Create on 2023/6/24
// @author xuzhuoxi
package servicex

const (
	//// EventOnServiceProcessing
	//// A single service initializes the progress update event, which is dispatched by the service instance.
	//// The service instance must be an implementation class of the IProgressingService interface
	//// 单个服务初始化进度更新事件，由服务实例调度。服务实例必须为IProgressingService接口的实现类
	//EventOnServiceProcessing = "Service.OnProcessing"
	//
	//// EventOnInitProcessing
	//// Service initialization process progress update
	//// 服务初始化进程进度更新
	//EventOnManagerInitProcessing = "ServiceManager:OnProcessing"

	// EventOnManagerInitFinish
	// Service initialization process completed
	// 服务初始化进程完成
	EventOnManagerInitFinish = "ServiceManager:OnManagerInitFinish"

	//----------

	//// EventOnServiceInjected
	//// Service argument injection result event, dispatched by ServiceManager.
	//// Succ=true when the service implements the IArgumentService interface.
	//// 服务参数注入结果事件, 由ServiceManager调度。当服务实现IArgumentService接口时，Succ=true。
	//// Event data format: ServiceResultEventData
	//// 事件数据格式：ServiceResultData
	//EventOnServiceInjected = "Service:OnInjected"
	//
	//// EventOnServiceAllInjected
	//// All service argument injection completion event
	//// 全部服务参数注入完成事件
	//// Event data format: null
	//// 事件数据格式：null
	//EventOnServiceAllInjected = "Service:OnAllInjected"

	// EventOnServiceAwaked
	// Service activation result event, dispatched by ServiceManager.
	// Succ=true when the service implements the IWakableService interface.
	// 服务激活结果事件, 由ServiceManager调度。当服务实现IAwakableService接口时，Succ=true。
	// Event data format: {named:string}
	// 事件数据格式：{named:string}
	EventOnServiceAwaked = "Service:OnAwaked"

	// EventOnServiceAllAwaked
	// All service activation completion event
	// 全部服务激活完成事件
	// Event data format: null
	// 事件数据格式：len:int
	EventOnServiceAllAwaked = "Service:OnAllAwaked"

	//----------

	// EventOnServiceInitStart
	// Single service initialization start event, dispatched by ServiceManager
	// 单个服务初始化开始事件，由ServiceManager调度
	// Event data format: ServiceStartData
	// 事件数据格式：ServiceStartData
	EventOnServiceInitStart = "Service:OnInitStart"

	// EventOnServiceInited
	// A single service initialization complete event, which are dispatched by the service instance.
	// Re-dispatched after being captured by ServiceManager.
	// 单个服务初始化完成事件，由服务实例调度。被ServiceManager捕获后重新调度。
	// Event data format: ServiceResultData
	// 事件数据格式：ServiceResultData
	EventOnServiceInited = "Service:OnInited"

	// EventOnServiceAllInited
	// All service initialization complete event
	// 全部服务初始化完成事件
	// Event data format: len:int
	// 事件数据格式：len:int
	EventOnServiceAllInited = "Service:OnAllInited"

	// EventOnServiceDataInitStart
	// Single service data initialization start event, dispatched by ServiceManager
	// 单个服务数据初始化开始事件，由ServiceManager调度
	// Event data format: ServiceStartData
	// 事件数据格式：ServiceStartData
	EventOnServiceDataInitStart = "Service:OnDataInitStart"

	// EventOnServiceDataInited
	// A single service data initialization complete event, which are dispatched by the service instance.
	// Re-dispatched after being captured by ServiceManager.
	// 单个服务数据初始化完成事件，由服务实例调度。被ServiceManager捕获后重新调度。
	// Event data format: ServiceResultData
	// 事件数据格式：ServiceResultData
	EventOnServiceDataInited = "Service:OnDataInited"

	// EventOnServiceDataAllInited
	// All service data initialization complete event
	// 全部服务数据初始化完成事件
	// Event data format: len:int
	// 事件数据格式：len:int
	EventOnServiceDataAllInited = "Service:OnDataAllInited"

	//----------

	// EventOnServiceDataLoadStart
	// A single  data service load data start event, dispatched by ServiceManager
	// Succ=true when the service implements the ILoadDataService interface.
	// 单个服务数据加载数据开始事件，由ServiceManager调度。当服务实现ILoadDataService接口时，Succ=true。
	// Event data format: ServiceStartData
	// 事件数据格式：ServiceStartData
	EventOnServiceDataLoadStart = "Service:OnDataLoadStart"

	// EventOnServiceDataLoaded
	// A single data service load data completion event, which are dispatched by the service instance.
	// Re-dispatched after being captured by ServiceManager.
	// 单个数据服务加载数据完成事件，由服务实例调度。被ServiceManager捕获后重新调度。
	// Event data format: ServiceResultData
	// 事件数据格式：ServiceResultData
	EventOnServiceDataLoaded = "Service:OnDataLoaded"

	// EventOnServiceDataAllLoaded
	// All data service loading data completion event
	// 全部数据服务加载数据完成事件
	// Event data format: len:int
	// 事件数据格式：len:int
	EventOnServiceDataAllLoaded = "Service:OnDataAllLoaded"

	//----------

	// EventOnServiceDataSaveStart
	// A single data service save data start event, dispatched by ServiceManager
	// Succ=true when the service implements the ISaveDataService interface.
	// 单个服务数据保存数据开始事件，由ServiceManager调度。当服务实现ISaveDataService接口时，Succ=true。
	// Event data format: ServiceStartData
	// 事件数据格式：ServiceStartData
	EventOnServiceDataSaveStart = "Service:OnDataSaveStart"

	// EventOnServiceDataSaved
	// A single data service saves data completion events, which are dispatched by the service instance.
	// Re-dispatched after being captured by ServiceManager.
	// 单个数据服务保存数据完成事件，由服务实例调度。被ServiceManager捕获后重新调度。
	// Event data format: ServiceResultData
	// 事件数据格式：ServiceResultData
	EventOnServiceDataSaved = "Service:OnDataSaved"

	// EventOnServiceDataAllSaved
	// 全部数据服务保存数据完成事件
	// Event data format: len:int
	// 事件数据格式：len:int
	EventOnServiceDataAllSaved = "Service:OnDataAllSaved"
)

type ServiceStartData struct {
	ServiceName string
	Ignore      bool
}

type ServiceResultData struct {
	ServiceName string
	Suc         bool
	Err         error
}
