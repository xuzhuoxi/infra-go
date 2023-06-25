// Package servicex
// Create on 2023/6/25
// @author xuzhuoxi
package servicex

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/eventx"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/logx"
)

var (
	DefaultManager = &ServiceManager{}
)

type ServiceManager struct {
	logx.LoggerSupport
	eventx.EventDispatcher
	config  *ServiceConfig
	endCall *lang.Callback
}

func (o *ServiceManager) Init(config *ServiceConfig, endCall lang.FuncCallback) {
	o.config, o.endCall = config, lang.NewCallback(endCall)
	o.GetLogger().Infoln("ServiceManager:Init ---------- Start!")
	o.addListeners()
	o.activeServices()
	o.startInitServices()
}

func (o *ServiceManager) addListeners() {
	o.AddEventListener(EventOnServiceAwaked, o.onEventLog)
	o.AddEventListener(EventOnServiceInitStart, o.onEventLog)
	o.AddEventListener(EventOnServiceInited, o.onEventLog)
	o.AddEventListener(EventOnServiceDataInitStart, o.onEventLog)
	o.AddEventListener(EventOnServiceDataInited, o.onEventLog)
	o.AddEventListener(EventOnServiceDataLoadStart, o.onEventLog)
	o.AddEventListener(EventOnServiceDataLoaded, o.onEventLog)

	o.AddEventListener(EventOnServiceAllAwaked, o.onEventLog)
	o.AddEventListener(EventOnServiceAllInited, o.onEventLog)
	o.AddEventListener(EventOnServiceDataAllInited, o.onEventLog)
	o.AddEventListener(EventOnServiceDataAllLoaded, o.onEventLog)
	o.AddEventListener(EventOnManagerInitFinish, o.onEventLog)
}

func (o *ServiceManager) onEventLog(evd *eventx.EventData) {
	evtType := evd.EventType
	logger := o.GetLogger()
	switch evtType {
	case EventOnServiceInitStart:
		r := evd.Data.(*ServiceResultData)
		logger.Infoln(fmt.Sprintf("Service[name=%s] Init Start: Ignore=%v! ", r.ServiceName, !r.Suc))
	case EventOnServiceDataInitStart:
		r := evd.Data.(*ServiceResultData)
		logger.Infoln(fmt.Sprintf("Service[name=%s] Data Init Start: Ignore=%v! ", r.ServiceName, !r.Suc))
	case EventOnServiceDataLoadStart:
		r := evd.Data.(*ServiceResultData)
		logger.Infoln(fmt.Sprintf("Service[name=%s] Data Load Start: Ignore=%v! ", r.ServiceName, !r.Suc))
	// ---
	case EventOnServiceAwaked:
		logger.Infoln(fmt.Sprintf("Service[name=%s] Awaked! ", evd.Data.(string)))
	case EventOnServiceInited:
		logger.Infoln(fmt.Sprintf("Service[name=%s] Inited! ", evd.Data.(string)))
	case EventOnServiceDataInited:
		logger.Infoln(fmt.Sprintf("Service[name=%s] Data Inited! ", evd.Data.(string)))
	case EventOnServiceDataLoaded:
		logger.Infoln(fmt.Sprintf("Service[name=%s] Data Loaded! ", evd.Data.(string)))
	// ---
	case EventOnServiceAllAwaked:
		logger.Infoln(fmt.Sprintf("Services[num=%d] All Awaked! ", evd.Data.(int)))
	case EventOnServiceAllInited:
		logger.Infoln(fmt.Sprintf("Services[num=%d] All Inited! ", evd.Data.(int)))
	case EventOnServiceDataAllInited:
		logger.Infoln(fmt.Sprintf("Services[num=%d] All Data Inited! ", evd.Data.(int)))
	case EventOnServiceDataAllLoaded:
		logger.Infoln(fmt.Sprintf("Services[num=%d] All Data Loaded! ", evd.Data.(int)))
	// ---
	case EventOnManagerInitFinish:
		logger.Infoln("ServiceManager:Init ---------- Finish!")
	}
}

func (o *ServiceManager) removeListeners() {
	o.RemoveEventListener(EventOnManagerInitFinish, o.onEventLog)

	o.RemoveEventListener(EventOnServiceDataAllLoaded, o.onEventLog)
	o.RemoveEventListener(EventOnServiceDataAllInited, o.onEventLog)
	o.RemoveEventListener(EventOnServiceAllInited, o.onEventLog)
	o.RemoveEventListener(EventOnServiceAllAwaked, o.onEventLog)

	o.RemoveEventListener(EventOnServiceDataLoaded, o.onEventLog)
	o.RemoveEventListener(EventOnServiceDataLoadStart, o.onEventLog)
	o.RemoveEventListener(EventOnServiceDataInited, o.onEventLog)
	o.RemoveEventListener(EventOnServiceDataInitStart, o.onEventLog)
	o.RemoveEventListener(EventOnServiceInited, o.onEventLog)
	o.RemoveEventListener(EventOnServiceInitStart, o.onEventLog)
	o.RemoveEventListener(EventOnServiceAwaked, o.onEventLog)
}

func (o *ServiceManager) activeServices() {
	count := 0
	o.config.foreachService(func(info *ServiceInfo) {
		info.GenImpl()
		suc := info.IsAwakableService()
		if suc {
			(info.impl.(IAwakableService)).Awake()
			count += 1
			o.DispatchEvent(EventOnServiceAwaked, o, info.name)
		}
	})
	o.DispatchEvent(EventOnServiceAllAwaked, o, count)
}

func (o *ServiceManager) startInitServices() {
	o.OnceEventListener(EventOnServiceAllInited, o.onServicesInited)
	handler := newInitHandler(o.config, &(o.EventDispatcher))
	handler.Start(nil)
}

func (o *ServiceManager) onServicesInited(evd *eventx.EventData) {
	o.startInitDataServices()
}

func (o *ServiceManager) startInitDataServices() {
	o.OnceEventListener(EventOnServiceDataAllInited, o.onDataServicesInited)
	handler := newInitDataHandler(o.config, &(o.EventDispatcher))
	handler.Start(nil)
}

func (o *ServiceManager) onDataServicesInited(evd *eventx.EventData) {
	o.startLoadDatas()
}

func (o *ServiceManager) startLoadDatas() {
	o.OnceEventListener(EventOnServiceDataAllLoaded, o.onServicesDataLoaded)
	handler := newLoadDataHandler(o.config, &(o.EventDispatcher))
	handler.Start(nil)
}

func (o *ServiceManager) onServicesDataLoaded(evd *eventx.EventData) {
	o.endCall.Invoke()
	o.DispatchEvent(EventOnManagerInitFinish, o, nil)
	o.removeListeners()
}
