// Package servicex
// Create on 2023/6/24
// @author xuzhuoxi
package servicex

import (
	"errors"
	"fmt"
)

var (
	DefaultConfig = &ServiceConfig{}
)

type ServiceConfig struct {
	list []*ServiceInfo
}

// Size
// Total number of configured services
// 已配置的服务的总个数
func (o *ServiceConfig) Size() int {
	return len(o.list)
}

// FindServiceInfo
// Get a specified service configuration information
// 取指定一个的服务配置信息
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

// ContainsService
// Check service existence
// 检查服务存在性
func (o *ServiceConfig) ContainsService(name string) bool {
	return o.FindServiceInfo(name) != nil
}

// GetServiceImpl
// Take the specified service implementation object
// 取指定服务器实现对象，
func (o *ServiceConfig) GetServiceImpl(name string) IService {
	return o.FindServiceInfo(name).impl
}

// AddConfig
// Add a service to the end of the configuration list
// 添加服务到配置列表尾部
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
