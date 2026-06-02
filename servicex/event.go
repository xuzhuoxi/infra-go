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

	// EventOnManagerInitFinish is dispatched when the service initialization process completes.
	// EventOnManagerInitFinish 在服务初始化进程完成时分发。
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

	// EventOnServiceAwaked is dispatched when a service is awakened by ServiceManager.
	// EventOnServiceAwaked 是服务激活结果事件，由 ServiceManager 调度。
	//
	// Dispatched when the service implements IAwakableService.
	// 当服务实现 IAwakableService 接口时分发。
	//
	// Event data format: service name string.
	// 事件数据格式：服务名称 string。
	EventOnServiceAwaked = "Service:OnAwaked"

	// EventOnServiceAllAwaked is dispatched when all service awakening completes.
	// EventOnServiceAllAwaked 是全部服务激活完成事件。
	//
	// Event data format: int count.
	// 事件数据格式：int 数量。
	EventOnServiceAllAwaked = "Service:OnAllAwaked"

	//----------

	// EventOnServiceInitStart is dispatched when a single service initialization starts.
	// EventOnServiceInitStart 是单个服务初始化开始事件，由 ServiceManager 调度。
	//
	// Event data format: ServiceStartData.
	// 事件数据格式：ServiceStartData。
	EventOnServiceInitStart = "Service:OnInitStart"

	// EventOnServiceInited is dispatched when a single service finishes initialization.
	// EventOnServiceInited 是单个服务初始化完成事件，由服务实例调度。
	//
	// Re-dispatched after being captured by ServiceManager.
	// 被 ServiceManager 捕获后重新调度。
	//
	// Event data format: ServiceResultData.
	// 事件数据格式：ServiceResultData。
	EventOnServiceInited = "Service:OnInited"

	// EventOnServiceAllInited is dispatched when all service initialization completes.
	// EventOnServiceAllInited 是全部服务初始化完成事件。
	//
	// Event data format: int count.
	// 事件数据格式：int 数量。
	EventOnServiceAllInited = "Service:OnAllInited"

	// EventOnServiceDataInitStart is dispatched when a single service data initialization starts.
	// EventOnServiceDataInitStart 是单个服务数据初始化开始事件，由 ServiceManager 调度。
	//
	// Event data format: ServiceStartData.
	// 事件数据格式：ServiceStartData。
	EventOnServiceDataInitStart = "Service:OnDataInitStart"

	// EventOnServiceDataInited is dispatched when a single service finishes data initialization.
	// EventOnServiceDataInited 是单个服务数据初始化完成事件，由服务实例调度。
	//
	// Re-dispatched after being captured by ServiceManager.
	// 被 ServiceManager 捕获后重新调度。
	//
	// Event data format: ServiceResultData.
	// 事件数据格式：ServiceResultData。
	EventOnServiceDataInited = "Service:OnDataInited"

	// EventOnServiceDataAllInited is dispatched when all service data initialization completes.
	// EventOnServiceDataAllInited 是全部服务数据初始化完成事件。
	//
	// Event data format: int count.
	// 事件数据格式：int 数量。
	EventOnServiceDataAllInited = "Service:OnDataAllInited"

	//----------

	// EventOnServiceDataLoadStart is dispatched when a single service data load starts.
	// EventOnServiceDataLoadStart 是单个服务数据加载开始事件，由 ServiceManager 调度。
	//
	// Dispatched when the service implements ILoadDataService.
	// 当服务实现 ILoadDataService 接口时分发。
	//
	// Event data format: ServiceStartData.
	// 事件数据格式：ServiceStartData。
	EventOnServiceDataLoadStart = "Service:OnDataLoadStart"

	// EventOnServiceDataLoaded is dispatched when a single service finishes loading data.
	// EventOnServiceDataLoaded 是单个数据服务加载数据完成事件，由服务实例调度。
	//
	// Re-dispatched after being captured by ServiceManager.
	// 被 ServiceManager 捕获后重新调度。
	//
	// Event data format: ServiceResultData.
	// 事件数据格式：ServiceResultData。
	EventOnServiceDataLoaded = "Service:OnDataLoaded"

	// EventOnServiceDataAllLoaded is dispatched when all service data loading completes.
	// EventOnServiceDataAllLoaded 是全部数据服务加载数据完成事件。
	//
	// Event data format: int count.
	// 事件数据格式：int 数量。
	EventOnServiceDataAllLoaded = "Service:OnDataAllLoaded"

	//----------

	// EventOnServiceDataSaveStart is dispatched when a single service data save starts.
	// EventOnServiceDataSaveStart 是单个服务数据保存开始事件，由 ServiceManager 调度。
	//
	// Dispatched when the service implements ISaveDataService.
	// 当服务实现 ISaveDataService 接口时分发。
	//
	// Event data format: ServiceStartData.
	// 事件数据格式：ServiceStartData。
	EventOnServiceDataSaveStart = "Service:OnDataSaveStart"

	// EventOnServiceDataSaved is dispatched when a single service finishes saving data.
	// EventOnServiceDataSaved 是单个数据服务保存数据完成事件，由服务实例调度。
	//
	// Re-dispatched after being captured by ServiceManager.
	// 被 ServiceManager 捕获后重新调度。
	//
	// Event data format: ServiceResultData.
	// 事件数据格式：ServiceResultData。
	EventOnServiceDataSaved = "Service:OnDataSaved"

	// EventOnServiceDataAllSaved is dispatched when all service data saving completes.
	// EventOnServiceDataAllSaved 是全部数据服务保存数据完成事件。
	//
	// Event data format: int count.
	// 事件数据格式：int 数量。
	EventOnServiceDataAllSaved = "Service:OnDataAllSaved"
)

// ServiceStartData is the payload for service start events.
// ServiceStartData 是服务开始类事件的载荷数据。
type ServiceStartData struct {
	// ServiceName is the target service name.
	// ServiceName 是目标服务名称。
	ServiceName string
	// Ignore reports whether the service skipped the operation.
	// Ignore 表示该服务是否跳过了当前操作。
	Ignore bool
}

// ServiceResultData is the payload for service result events.
// ServiceResultData 是服务结果类事件的载荷数据。
type ServiceResultData struct {
	// ServiceName is the target service name.
	// ServiceName 是目标服务名称。
	ServiceName string
	// Suc reports whether the operation succeeded.
	// Suc 表示操作是否成功。
	Suc bool
	// Err carries the error when the operation failed.
	// Err 在操作失败时携带错误信息。
	Err error
}
