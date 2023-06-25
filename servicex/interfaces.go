// Package servicex
// Create on 2023/6/24
// @author xuzhuoxi
package servicex

import "github.com/xuzhuoxi/infra-go/eventx"

// IService
// Service base interface
// 服务的基础接口
type IService interface {
	// ServiceName
	// Service name.
	// 服务名称
	ServiceName() string
	// SetServiceName
	// set service name
	// 设置服务名称
	SetServiceName(name string)
}

// IAwakableService
// Service Awaka interface
// 服务激活接口
type IAwakableService interface {
	// Awake
	// Awake service
	// 激活Service
	// Async is not allowed
	// 不允许使用异步
	Awake()
}

// IClearService
// Reset service interface, providing reset service to initialized state
// 重置服务接口，提供把服务重置到初始化状态
type IClearService interface {
	// Clear
	// reset
	// 重置
	// Clear events, clear timers, etc.
	// 清除事件、清除计时器等
	Clear()
}

// IInitService
// Service initialization interface
// 服务初始化接口
// Only when the current interface is implemented and configured into ServiceConfig,
// the init method will be executed during the initialization process
// 只有实现了当前接口，并配置到ServiceConfig中时，在初始化过程中才会执行init方法
type IInitService interface {
	IService
	eventx.IEventDispatcher
	// IsInited
	// Whether the initialization has been completed
	// 是否已经完成初始化
	IsInited() bool
	// Init
	// Initialize base data
	// 初始化基础数据
	Init()
}

// IInitDataService
// Initialize data processing services
// 初始化数据处理服务
type IInitDataService interface {
	IService
	eventx.IEventDispatcher
	// IsDataInited
	// Whether data initialization has been completed
	// 是否已经完成数据初始化
	IsDataInited() bool
	// InitData
	// Initialization data
	// 初始化数据
	InitData()
}

// ILoadDataService
// Load data processing interface
// 加载数据处理接口
type ILoadDataService interface {
	eventx.IEventDispatcher
	// LoadData
	// Load Data
	// 加载数据
	LoadData()
}

// ISaveDataService
// Save data processing interface
// 保存数据处理接口
type ISaveDataService interface {
	eventx.IEventDispatcher
	// SaveData
	// Save data
	// 保存数据
	SaveData()
}
