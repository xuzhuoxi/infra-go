//
//Created by xuzhuoxi
//on 2019-02-17.
//@author xuzhuoxi
//
package extendx

type IExtension interface {
	//主键标识
	ExtensionName() string
}

type IInitExtension interface {
	//初始化
	InitExtension() error
	//反初始化
	DestroyExtension() error
}

type ISaveExtension interface {
	//保存数据
	SaveExtension() error
}

type IEnableExtension interface {
	//是否启用
	Enable() bool
	//启用
	EnableExtension() error
	//禁用
	DisableExtension() error
}
