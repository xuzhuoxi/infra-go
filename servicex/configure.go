package servicex

import (
	"errors"
	"fmt"
)

var (
	// DefaultConfig is the default empty ServiceConfig instance.
	// DefaultConfig 是默认的空 ServiceConfig 实例。
	DefaultConfig = &ServiceConfig{}
)

// ServiceConfig holds an ordered list of registered service definitions.
// ServiceConfig 保存已注册服务的有序配置列表。
type ServiceConfig struct {
	list []*ServiceInfo
}

// Size returns the total number of configured services.
// Size 返回已配置的服务总个数。
func (o *ServiceConfig) Size() int {
	return len(o.list)
}

// FindServiceInfo returns configuration for the named service, or nil if not found.
// FindServiceInfo 取指定服务名的配置信息，未找到时返回 nil。
func (o *ServiceConfig) FindServiceInfo(name string) *ServiceInfo {
	if len(name) == 0 {
		return nil
	}
	for index := range o.list {
		if o.list[index].name == name {
			return o.list[index]
		}
	}
	return nil
}

// ContainsService reports whether a service with the given name is configured.
// ContainsService 检查指定名称的服务是否已配置。
func (o *ServiceConfig) ContainsService(name string) bool {
	return o.FindServiceInfo(name) != nil
}

// GetServiceImpl returns the implementation object for the named service.
// GetServiceImpl 取指定服务的实现对象。
func (o *ServiceConfig) GetServiceImpl(name string) IService {
	return o.FindServiceInfo(name).impl
}

// AddConfig appends a service to the configuration list.
// AddConfig 添加服务到配置列表尾部。
func (o *ServiceConfig) AddConfig(info *ServiceInfo, ignoreSame bool) error {
	if nil == info {
		return errors.New("Nil ServiceInfo! ")
	}
	if !ignoreSame && o.ContainsService(info.name) {
		return errors.New(fmt.Sprintf("Duplicate name[%s]!", info.name))
	}
	o.list = append(o.list, info)
	return nil
}

func (o *ServiceConfig) foreachService(funcEach func(info *ServiceInfo)) {
	if len(o.list) == 0 {
		return
	}
	for index := range o.list {
		funcEach(o.list[index])
	}
}
