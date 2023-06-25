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

func (o *ServiceBase) InvokeInited() {
	o.DispatchEvent(EventOnServiceInited, o, o.name)
}

func (o *ServiceBase) InvokeDataInited() {
	o.DispatchEvent(EventOnServiceDataInited, o, o.name)
}

func (o *ServiceBase) InvokeDataLoaded() {
	o.DispatchEvent(EventOnServiceDataLoaded, o, o.name)
}

func (o *ServiceBase) InvokeDataSaved() {
	o.DispatchEvent(EventOnServiceDataSaved, o, o.name)
}
