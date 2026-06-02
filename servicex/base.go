package servicex

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

// ServiceBase embeds EventDispatcher and tracks common service lifecycle flags.
// ServiceBase 内嵌 EventDispatcher 并跟踪通用的服务生命周期状态。
type ServiceBase struct {
	eventx.EventDispatcher
	name       string
	inited     bool
	dataInited bool
	dataLoaded bool
}

// ServiceName returns the service name.
// ServiceName 返回服务名称。
func (o *ServiceBase) ServiceName() string {
	return o.name
}

// SetServiceName sets the service name.
// SetServiceName 设置服务名称。
func (o *ServiceBase) SetServiceName(name string) {
	o.name = name
}

// IsInited reports whether initialization has completed.
// IsInited 返回是否已经完成初始化。
func (o *ServiceBase) IsInited() bool {
	return o.inited
}

// IsDataInited reports whether data initialization has completed.
// IsDataInited 返回是否已经完成数据初始化。
func (o *ServiceBase) IsDataInited() bool {
	return o.dataInited
}

// IsDataLoaded reports whether data loading has completed.
// IsDataLoaded 返回是否已经完成数据加载。
func (o *ServiceBase) IsDataLoaded() bool {
	return o.dataLoaded
}

// Clear resets lifecycle flags and removes all event listeners.
// Clear 重置生命周期标志并清除全部事件监听器。
func (o *ServiceBase) Clear() {
	o.RemoveEventListeners()
	o.inited, o.dataInited, o.dataInited = false, false, false
}

// InvokeInitStart dispatches EventOnServiceInitStart with ignore flag.
// InvokeInitStart 分发 EventOnServiceInitStart 事件并携带 ignore 标志。
func (o *ServiceBase) InvokeInitStart(ignore bool) {
	o.DispatchEvent(EventOnServiceInitStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

// InvokeInited dispatches EventOnServiceInited with success and error result.
// InvokeInited 分发 EventOnServiceInited 事件并携带成功与错误结果。
func (o *ServiceBase) InvokeInited(suc bool, err error) {
	o.DispatchEvent(EventOnServiceInited, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

// InvokeDataInitStart dispatches EventOnServiceDataInitStart with ignore flag.
// InvokeDataInitStart 分发 EventOnServiceDataInitStart 事件并携带 ignore 标志。
func (o *ServiceBase) InvokeDataInitStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataInitStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

// InvokeDataInited dispatches EventOnServiceDataInited with success and error result.
// InvokeDataInited 分发 EventOnServiceDataInited 事件并携带成功与错误结果。
func (o *ServiceBase) InvokeDataInited(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataInited, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

// InvokeDataLoadStart dispatches EventOnServiceDataLoadStart with ignore flag.
// InvokeDataLoadStart 分发 EventOnServiceDataLoadStart 事件并携带 ignore 标志。
func (o *ServiceBase) InvokeDataLoadStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataLoadStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

// InvokeDataLoaded dispatches EventOnServiceDataLoaded with success and error result.
// InvokeDataLoaded 分发 EventOnServiceDataLoaded 事件并携带成功与错误结果。
func (o *ServiceBase) InvokeDataLoaded(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataLoaded, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

// InvokeDataSaveStart dispatches EventOnServiceDataSaveStart with ignore flag.
// InvokeDataSaveStart 分发 EventOnServiceDataSaveStart 事件并携带 ignore 标志。
func (o *ServiceBase) InvokeDataSaveStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataSaveStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

// InvokeDataSaved dispatches EventOnServiceDataSaved with success and error result.
// InvokeDataSaved 分发 EventOnServiceDataSaved 事件并携带成功与错误结果。
func (o *ServiceBase) InvokeDataSaved(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataSaved, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}
