package mgox

import (
	"context"
	"io"
	"sync"

	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/storagex/mongox"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// NewMongoDriver creates a driver that manages multiple named MongoDB clients.
// NewMongoDriver 创建管理多个具名 MongoDB 客户端的驱动。
//
// maxClient enables a concurrency limit when greater than zero; zero disables the limit.
// maxClient 大于 0 时启用并发限制；为 0 时不限制。
func NewMongoDriver(maxClient int) IMongoDriver {
	rs := &MongoDriver{}
	if maxClient > 0 {
		rs.ChanLimit = &lang.ChannelLimit{}
	} else {
		rs.ChanLimit = &lang.ChannelLimitNone{}
	}
	rs.ChanLimit.StartLimit()
	return rs
}

// IMongoDriver manages the lifecycle of multiple named MongoDB clients.
// IMongoDriver 管理多个具名 MongoDB 客户端的生命周期。
type IMongoDriver interface {
	io.Closer
	// OpenClient opens a named client with the given options and context.
	// OpenClient 使用指定选项与 context 打开具名客户端。
	OpenClient(name string, options *options.ClientOptions, context context.Context) (client IMongoClient, err error)
	// CloseClient closes and removes a client by name.
	// CloseClient 按名称关闭并移除客户端。
	CloseClient(name string) error
	// NumClient returns the number of open clients.
	// NumClient 返回已打开客户端数量。
	NumClient() int

	// GetClient returns a client by name.
	// GetClient 按名称返回客户端。
	GetClient(name string) (client IMongoClient, ok bool)
	// SetMasterClient sets or clears the master client by name.
	// SetMasterClient 按名称设置或清除主客户端。
	SetMasterClient(name string) (client IMongoClient, change bool)
	// MasterClient returns the current master client, or nil if unset.
	// MasterClient 返回当前主客户端；未设置时返回 nil。
	MasterClient() IMongoClient
}

// MongoDriver holds named clients with optional concurrency limiting.
// MongoDriver 持有具名客户端，并支持可选的并发限制。
type MongoDriver struct {
	Clients     []IMongoClient
	MasterIndex int

	ChanLimit lang.IChannelLimit
	Lock      sync.RWMutex
}

// Close disconnects all clients and stops the concurrency limiter.
// Close 断开所有客户端并停止并发限制器。
func (d *MongoDriver) Close() error {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if 0 < len(d.Clients) {
		for index := len(d.Clients) - 1; index >= 0; index-- {
			d.Clients[index].Disconnect()
		}
	}
	d.ChanLimit.StopLimit()
	return nil
}

// OpenClient opens a named client with the given options and context.
// OpenClient 使用指定选项与 context 打开具名客户端。
func (d *MongoDriver) OpenClient(name string, options *options.ClientOptions, context context.Context) (client IMongoClient, err error) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	//funcName := "MongoDriver.OpenClient"
	d.ChanLimit.Add()
	if nil == context {
		context = DefaultContext
	}
	client = NewMongoClient(name, options, context)
	err = client.Connect()
	if nil != err {
		d.ChanLimit.Done()
		return nil, err
	}
	d.appendClient(client)
	return
}

// CloseClient closes and removes a client by name.
// CloseClient 按名称关闭并移除客户端。
func (d *MongoDriver) CloseClient(name string) error {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if "" == name {
		return mongox.ClientNilError("MongoDriver.CloseClient name=\"\"")
	}
	index := d.indexOfName(name)
	if -1 != index {
		defer func() {
			if index == d.MasterIndex {
				d.MasterIndex = -1
			}
			d.removeSession(d.Clients[index])
			d.ChanLimit.Done()
		}()
		return d.Clients[index].Disconnect()
	}
	return mongox.SessionNilError("MongoDriver.CloseClient name=\"" + name + "\"")
}

// NumClient returns the number of open clients.
// NumClient 返回已打开客户端数量。
func (d *MongoDriver) NumClient() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	return len(d.Clients)
}

// GetClient returns a client by name.
// GetClient 按名称返回客户端。
func (d *MongoDriver) GetClient(name string) (client IMongoClient, ok bool) {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	if "" == name {
		return nil, false
	}
	index := d.indexOfName(name)
	if -1 == index {
		return nil, false
	}
	return d.Clients[index], true
}

// SetMasterClient sets or clears the master client by name.
// SetMasterClient 按名称设置或清除主客户端。
func (d *MongoDriver) SetMasterClient(name string) (client IMongoClient, change bool) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if "" == name {
		if -1 == d.MasterIndex {
			return nil, false
		} else {
			d.MasterIndex = -1
			return nil, true
		}
	}
	index := d.indexOfName(name)
	change = index == d.MasterIndex
	d.MasterIndex = index
	return d.Clients[index], change
}

// MasterClient returns the current master client, or nil if unset.
// MasterClient 返回当前主客户端；未设置时返回 nil。
func (d *MongoDriver) MasterClient() IMongoClient {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	if -1 == d.MasterIndex {
		return nil
	}
	return d.Clients[d.MasterIndex]
}

//----------------------------------

func (d *MongoDriver) indexOfSession(client IMongoClient) int {
	if nil == client {
		return -1
	}
	return d.indexOfName(client.Name())
}

func (d *MongoDriver) indexOfName(name string) int {
	if "" == name {
		return -1
	}
	for index := range d.Clients {
		if name == d.Clients[index].Name() {
			return index
		}
	}
	return -1
}

func (d *MongoDriver) appendClient(client IMongoClient) {
	if nil != client {
		d.Clients = append(d.Clients, client)
	}
}

func (d *MongoDriver) removeSession(client IMongoClient) {
	index := d.indexOfSession(client)
	if -1 == index {
		return
	}
	d.Clients = append(d.Clients[:index], d.Clients[index+1:]...)
}
