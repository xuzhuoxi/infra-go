// Package servicex
// Create on 2023/6/24
// @author xuzhuoxi
package servicex

import "reflect"

type FuncNewService = func(...interface{}) IService

func NewServiceInfo(name string, funcNew FuncNewService, args ...interface{}) *ServiceInfo {
	return &ServiceInfo{
		name:    name,
		funcNew: funcNew,
		args:    args,
	}
}

type ServiceInfo struct {
	name    string
	funcNew FuncNewService
	args    []interface{}
	impl    IService
}

func (o *ServiceInfo) ServiceName() string {
	return o.name
}

func (o *ServiceInfo) CheckImplType(t reflect.Type) bool {
	return reflect.TypeOf(o) == t
}

func (o *ServiceInfo) IsAwakableService() bool {
	_, ok := o.impl.(IAwakableService)
	return ok
}

func (o *ServiceInfo) IsInitService() bool {
	_, ok := o.impl.(IInitService)
	return ok
}

func (o *ServiceInfo) IsInitDataService() bool {
	_, ok := o.impl.(IInitDataService)
	return ok
}

func (o *ServiceInfo) IsLoadDataService() bool {
	_, ok := o.impl.(ILoadDataService)
	return ok
}

func (o *ServiceInfo) IsSaveDataService() bool {
	_, ok := o.impl.(ISaveDataService)
	return ok
}

func (o *ServiceInfo) GenImpl() IService {
	if nil == o.funcNew {
		return nil
	}
	o.impl = o.funcNew(o.args...)
	o.impl.SetServiceName(o.name)
	return o.impl
}

func (o *ServiceInfo) Clone() *ServiceInfo {
	return &ServiceInfo{
		name:    o.name,
		funcNew: o.funcNew,
		args:    o.args,
	}
}
