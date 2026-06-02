package servicex

import "reflect"

// FuncNewService is a factory function that creates an IService from optional arguments.
// FuncNewService 是从可选参数创建 IService 的工厂函数类型。
type FuncNewService = func(...interface{}) IService

// NewServiceInfo creates a ServiceInfo with name, factory, and constructor arguments.
// NewServiceInfo 使用名称、工厂函数与构造参数创建 ServiceInfo。
func NewServiceInfo(name string, funcNew FuncNewService, args ...interface{}) *ServiceInfo {
	return &ServiceInfo{
		name:    name,
		funcNew: funcNew,
		args:    args,
	}
}

// ServiceInfo holds service metadata, factory, and generated implementation.
// ServiceInfo 保存服务元数据、工厂函数及已生成的实现实例。
type ServiceInfo struct {
	name    string
	funcNew FuncNewService
	args    []interface{}
	impl    IService
}

// ServiceName returns the configured service name.
// ServiceName 返回配置的服务名称。
func (o *ServiceInfo) ServiceName() string {
	return o.name
}

// CheckImplType reports whether the receiver's reflect.Type equals t.
// CheckImplType 判断接收者的 reflect.Type 是否等于 t。
func (o *ServiceInfo) CheckImplType(t reflect.Type) bool {
	return reflect.TypeOf(o) == t
}

// IsAwakableService reports whether the implementation satisfies IAwakableService.
// IsAwakableService 判断实现是否满足 IAwakableService 接口。
func (o *ServiceInfo) IsAwakableService() bool {
	_, ok := o.impl.(IAwakableService)
	return ok
}

// IsInitService reports whether the implementation satisfies IInitService.
// IsInitService 判断实现是否满足 IInitService 接口。
func (o *ServiceInfo) IsInitService() bool {
	_, ok := o.impl.(IInitService)
	return ok
}

// IsInitDataService reports whether the implementation satisfies IInitDataService.
// IsInitDataService 判断实现是否满足 IInitDataService 接口。
func (o *ServiceInfo) IsInitDataService() bool {
	_, ok := o.impl.(IInitDataService)
	return ok
}

// IsLoadDataService reports whether the implementation satisfies ILoadDataService.
// IsLoadDataService 判断实现是否满足 ILoadDataService 接口。
func (o *ServiceInfo) IsLoadDataService() bool {
	_, ok := o.impl.(ILoadDataService)
	return ok
}

// IsSaveDataService reports whether the implementation satisfies ISaveDataService.
// IsSaveDataService 判断实现是否满足 ISaveDataService 接口。
func (o *ServiceInfo) IsSaveDataService() bool {
	_, ok := o.impl.(ISaveDataService)
	return ok
}

// GenImpl creates the service implementation via the factory and sets its name.
// GenImpl 通过工厂函数创建服务实现并设置服务名称。
func (o *ServiceInfo) GenImpl() IService {
	if nil == o.funcNew {
		return nil
	}
	o.impl = o.funcNew(o.args...)
	o.impl.SetServiceName(o.name)
	return o.impl
}

// Clone returns a shallow copy without the generated implementation.
// Clone 返回不含已生成实现的浅拷贝。
func (o *ServiceInfo) Clone() *ServiceInfo {
	return &ServiceInfo{
		name:    o.name,
		funcNew: o.funcNew,
		args:    o.args,
	}
}
