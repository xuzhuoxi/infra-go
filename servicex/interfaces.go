package servicex

import "github.com/xuzhuoxi/infra-go/eventx"

// IService is the base service interface.
// IService 是服务的基础接口。
type IService interface {
	// ServiceName returns the service name.
	// ServiceName 返回服务名称。
	ServiceName() string
	// SetServiceName sets the service name.
	// SetServiceName 设置服务名称。
	SetServiceName(name string)
}

// IAwakableService defines synchronous service awakening.
// IAwakableService 是服务激活接口。
type IAwakableService interface {
	// Awake awakens the service synchronously.
	// Awake 同步激活 Service。
	//
	// Async is not allowed.
	// 不允许使用异步。
	Awake()
}

// IClearService resets a service to its initialized state.
// IClearService 是重置服务接口，提供把服务重置到初始化状态。
type IClearService interface {
	// Clear resets the service; clears events, timers, etc.
	// Clear 重置服务，清除事件、清除计时器等。
	Clear()
}

// IInitService defines service initialization.
// IInitService 是服务初始化接口。
//
// Only when this interface is implemented and configured in ServiceConfig,
// Init is executed during the initialization process.
// 只有实现了当前接口，并配置到 ServiceConfig 中时，在初始化过程中才会执行 Init 方法。
type IInitService interface {
	IService
	eventx.IEventDispatcher
	// IsInited reports whether initialization has completed.
	// IsInited 返回是否已经完成初始化。
	IsInited() bool
	// Init initializes base data.
	// Init 初始化基础数据。
	Init()
}

// IInitDataService defines service data initialization.
// IInitDataService 是初始化数据处理服务接口。
type IInitDataService interface {
	IService
	eventx.IEventDispatcher
	// IsDataInited reports whether data initialization has completed.
	// IsDataInited 返回是否已经完成数据初始化。
	IsDataInited() bool
	// InitData initializes service data.
	// InitData 初始化数据。
	InitData()
}

// ILoadDataService defines service data loading.
// ILoadDataService 是加载数据处理接口。
type ILoadDataService interface {
	eventx.IEventDispatcher
	// LoadData loads service data.
	// LoadData 加载数据。
	LoadData()
}

// ISaveDataService defines service data saving.
// ISaveDataService 是保存数据处理接口。
type ISaveDataService interface {
	eventx.IEventDispatcher
	// SaveData saves service data.
	// SaveData 保存数据。
	SaveData()
}
