// Package servicex
// Create on 2023/6/25
// @author xuzhuoxi
package servicex

import (
	"github.com/xuzhuoxi/infra-go/eventx"
)

type ServiceBase struct {
	eventx.EventDispatcher
	name       string
	inited     bool
	dataInited bool
	dataLoaded bool
}

func (o *ServiceBase) ServiceName() string {
	return o.name
}

func (o *ServiceBase) SetServiceName(name string) {
	o.name = name
}

func (o *ServiceBase) IsInited() bool {
	return o.inited
}

func (o *ServiceBase) IsDataInited() bool {
	return o.dataInited
}

func (o *ServiceBase) IsDataLoaded() bool {
	return o.dataLoaded
}

func (o *ServiceBase) Clear() {
	o.RemoveEventListeners()
	o.inited, o.dataInited, o.dataInited = false, false, false
}

func (o *ServiceBase) InvokeInitStart(ignore bool) {
	o.DispatchEvent(EventOnServiceInitStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

func (o *ServiceBase) InvokeInited(suc bool, err error) {
	o.DispatchEvent(EventOnServiceInited, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

func (o *ServiceBase) InvokeDataInitStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataInitStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

func (o *ServiceBase) InvokeDataInited(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataInited, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

func (o *ServiceBase) InvokeDataLoadStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataLoadStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

func (o *ServiceBase) InvokeDataLoaded(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataLoaded, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}

func (o *ServiceBase) InvokeDataSaveStart(ignore bool) {
	o.DispatchEvent(EventOnServiceDataSaveStart, o, ServiceStartData{ServiceName: o.name, Ignore: ignore})
}

func (o *ServiceBase) InvokeDataSaved(suc bool, err error) {
	o.DispatchEvent(EventOnServiceDataSaved, o, ServiceResultData{ServiceName: o.name, Suc: suc, Err: err})
}
