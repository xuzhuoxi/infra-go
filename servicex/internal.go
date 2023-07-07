// Package servicex
// Create on 2023/6/25
// @author xuzhuoxi
package servicex

import (
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
)

func newHandler(config *ServiceConfig, dispatcher eventx.IEventDispatcher) handler {
	return handler{config: config, dispatcher: dispatcher}
}

type handler struct {
	config     *ServiceConfig
	dispatcher eventx.IEventDispatcher
	index      int
	count      int
	endCall    *lang.Callback
}

func (o *handler) currentServiceInfo() *ServiceInfo {
	return o.config.list[o.index]
}

func newInitHandler(config *ServiceConfig, dispatcher eventx.IEventDispatcher) *initHandler {
	return &initHandler{handler: newHandler(config, dispatcher)}
}

type initHandler struct {
	handler
}

func (o *initHandler) Start(endCall *lang.Callback) {
	o.endCall, o.index, o.count = endCall, 0, 0
	o.tryNext()
}

func (o *initHandler) tryNext() {
	if o.index >= o.config.Size() {
		if nil != o.endCall {
			o.endCall.Invoke()
			o.endCall = nil
		}
		o.dispatcher.DispatchEvent(EventOnServiceAllInited, o.dispatcher, o.count)
		return
	}
	info := o.currentServiceInfo()
	checkIs := info.IsInitService()
	o.dispatcher.DispatchEvent(EventOnServiceInitStart, o.dispatcher, ServiceStartData{ServiceName: info.name, Ignore: !checkIs})
	if checkIs {
		o.doInit(info)
	} else {
		o.index += 1
		o.tryNext()
	}
}

func (o *initHandler) doInit(info *ServiceInfo) {
	impl := info.impl.(IInitService)
	impl.OnceEventListener(EventOnServiceInited, o.onInited)
	impl.Init()
}

func (o *initHandler) onInited(evd *eventx.EventData) {
	o.dispatcher.DispatchEvent(EventOnServiceInited, o.dispatcher, evd.Data)
	o.index += 1
	o.count += 1
	o.tryNext()
}

func newInitDataHandler(config *ServiceConfig, dispatcher eventx.IEventDispatcher) *initDataHandler {
	return &initDataHandler{handler: newHandler(config, dispatcher)}
}

type initDataHandler struct {
	handler
}

func (o *initDataHandler) Start(endCall *lang.Callback) {
	o.endCall, o.index, o.count = endCall, 0, 0
	o.tryNext()
}

func (o *initDataHandler) tryNext() {
	if o.index >= o.config.Size() {
		if nil != o.endCall {
			o.endCall.Invoke()
			o.endCall = nil
		}
		o.dispatcher.DispatchEvent(EventOnServiceDataAllInited, o.dispatcher, o.count)
		return
	}
	info := o.currentServiceInfo()
	checkIs := info.IsInitDataService()
	o.dispatcher.DispatchEvent(EventOnServiceDataInitStart, o.dispatcher, ServiceStartData{ServiceName: info.name, Ignore: !checkIs})
	if checkIs {
		o.doInitData(info)
	} else {
		o.index += 1
		o.tryNext()
	}
}

func (o *initDataHandler) doInitData(info *ServiceInfo) {
	impl := info.impl.(IInitDataService)
	impl.OnceEventListener(EventOnServiceDataInited, o.onDataInited)
	impl.InitData()
}

func (o *initDataHandler) onDataInited(evd *eventx.EventData) {
	o.dispatcher.DispatchEvent(EventOnServiceDataInited, o.dispatcher, evd.Data)
	o.index += 1
	o.count += 1
	o.tryNext()
}

func newLoadDataHandler(config *ServiceConfig, dispatcher eventx.IEventDispatcher) *loadDataHandler {
	return &loadDataHandler{handler: newHandler(config, dispatcher)}
}

type loadDataHandler struct {
	handler
}

func (o *loadDataHandler) Start(endCall *lang.Callback) {
	o.endCall, o.index, o.count = endCall, 0, 0
	o.tryNext()
}

func (o *loadDataHandler) tryNext() {
	if o.index >= o.config.Size() {
		if nil != o.endCall {
			o.endCall.Invoke()
			o.endCall = nil
		}
		o.dispatcher.DispatchEvent(EventOnServiceDataAllLoaded, o.dispatcher, o.count)
		return
	}
	info := o.currentServiceInfo()
	checkIs := info.IsLoadDataService()
	o.dispatcher.DispatchEvent(EventOnServiceDataLoadStart, o.dispatcher, ServiceStartData{ServiceName: info.name, Ignore: !checkIs})
	if checkIs {
		o.doLoadData(info)
	} else {
		o.index += 1
		o.tryNext()
	}
}

func (o *loadDataHandler) doLoadData(info *ServiceInfo) {
	impl := info.impl.(ILoadDataService)
	impl.OnceEventListener(EventOnServiceDataLoaded, o.onDataLoaded)
	impl.LoadData()
}

func (o *loadDataHandler) onDataLoaded(evd *eventx.EventData) {
	o.dispatcher.DispatchEvent(EventOnServiceDataLoaded, o.dispatcher, evd.Data)
	o.index += 1
	o.count += 1
	o.tryNext()
}

func newSaveDataHandler(config *ServiceConfig, dispatcher eventx.IEventDispatcher) *saveDataHandler {
	return &saveDataHandler{handler: newHandler(config, dispatcher)}
}

type saveDataHandler struct {
	handler
}

func (o *saveDataHandler) Start(endCall *lang.Callback) {
	o.endCall, o.index, o.count = endCall, 0, 0
	o.tryNext()
}

func (o *saveDataHandler) tryNext() {
	if o.index >= o.config.Size() {
		if nil != o.endCall {
			o.endCall.Invoke()
			o.endCall = nil
		}
		o.dispatcher.DispatchEvent(EventOnServiceDataAllSaved, o.dispatcher, o.count)
		return
	}
	info := o.currentServiceInfo()
	checkIs := info.IsSaveDataService()
	o.dispatcher.DispatchEvent(EventOnServiceDataSaveStart, o.dispatcher, ServiceStartData{ServiceName: info.name, Ignore: !checkIs})
	if checkIs {
		o.doSaveData(info)
	} else {
		o.index += 1
		o.tryNext()
	}
}

func (o *saveDataHandler) doSaveData(info *ServiceInfo) {
	impl := info.impl.(ISaveDataService)
	impl.OnceEventListener(EventOnServiceDataSaved, o.onDataSaved)
	impl.SaveData()
}

func (o *saveDataHandler) onDataSaved(evd *eventx.EventData) {
	o.dispatcher.DispatchEvent(EventOnServiceDataSaved, o.dispatcher, evd.Data)
	o.index += 1
	o.count += 1
	o.tryNext()
}
