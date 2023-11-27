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
	o.GetLogger().Infoln("[ServiceManager.Init]", "---------- Start!")
	o.addListeners()
	o.activeServices()
	o.startInitServices()
}

func (o *ServiceManager) SetLogStatus(open bool) {
	if open {
		o.addListeners()
	} else {
		o.removeListeners()
	}
}

func (o *ServiceManager) SaveServices() {
	handler := newSaveDataHandler(o.config, &(o.EventDispatcher))
	handler.Start(nil)
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

func (o *ServiceManager) onEventLog(evd *eventx.EventData) {
	funcName := "[ServiceManager.onEventLog]"
	evtType := evd.EventType
	logger := o.GetLogger()
	switch evtType {
	case EventOnServiceInitStart:
		r := evd.Data.(ServiceStartData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s] Init Start: Ignore=%v ", r.ServiceName, r.Ignore))
	case EventOnServiceDataInitStart:
		r := evd.Data.(ServiceStartData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s] Data Init Start: Ignore=%v ", r.ServiceName, r.Ignore))
	case EventOnServiceDataLoadStart:
		r := evd.Data.(ServiceStartData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s] Data Load Start: Ignore=%v ", r.ServiceName, r.Ignore))
	// ---
	case EventOnServiceInited:
		r := evd.Data.(ServiceResultData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s, suc=%v] Inited! ", r.ServiceName, r.Suc))
	case EventOnServiceDataInited:
		r := evd.Data.(ServiceResultData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s, suc=%v] Data Inited! ", r.ServiceName, r.Suc))
	case EventOnServiceDataLoaded:
		r := evd.Data.(ServiceResultData)
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s, suc=%v] Data Loaded! ", r.ServiceName, r.Suc))
	// ---
	case EventOnServiceAwaked:
		logger.Infoln(funcName, fmt.Sprintf("Service[name=%s] Awaked! ", evd.Data.(string)))
	case EventOnServiceAllAwaked:
		logger.Infoln(funcName, fmt.Sprintf("Services[num=%d] All Awaked! ", evd.Data.(int)))
	case EventOnServiceAllInited:
		logger.Infoln(funcName, fmt.Sprintf("Services[num=%d] All Inited! ", evd.Data.(int)))
	case EventOnServiceDataAllInited:
		logger.Infoln(funcName, fmt.Sprintf("Services[num=%d] All Data Inited! ", evd.Data.(int)))
	case EventOnServiceDataAllLoaded:
		logger.Infoln(funcName, fmt.Sprintf("Services[num=%d] All Data Loaded! ", evd.Data.(int)))
	// ---
	case EventOnManagerInitFinish:
		logger.Infoln(funcName, "---------- Finish!")
	}
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
	o.startLoadData()
}

func (o *ServiceManager) startLoadData() {
	o.OnceEventListener(EventOnServiceDataAllLoaded, o.onServicesDataLoaded)
	handler := newLoadDataHandler(o.config, &(o.EventDispatcher))
	handler.Start(nil)
}

func (o *ServiceManager) onServicesDataLoaded(evd *eventx.EventData) {
	o.endCall.Invoke()
	o.removeListeners()
	o.endInit()
}

func (o *ServiceManager) endInit() {
	funcName := "[ServiceManager.endInit]"
	logger := o.GetLogger()
	logger.Infoln(funcName)
	o.DispatchEvent(EventOnManagerInitFinish, o, nil)
}
