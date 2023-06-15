// Package mgox
// Created by xuzhuoxi
// on 2019-06-13.
// @author xuzhuoxi
//
package mgox

import (
	"context"
	"github.com/xuzhuoxi/infra-go/lang"
	"github.com/xuzhuoxi/infra-go/storagex/mongox"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"sync"
)

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

type IMongoDriver interface {
	io.Closer
	OpenClient(name string, options *options.ClientOptions, context context.Context) (client IMongoClient, err error)
	CloseClient(name string) error
	NumClient() int

	GetClient(name string) (client IMongoClient, ok bool)
	SetMasterClient(name string) (client IMongoClient, change bool)
	MasterClient() IMongoClient
}

type MongoDriver struct {
	Clients     []IMongoClient
	MasterIndex int

	ChanLimit lang.IChannelLimit
	Lock      sync.RWMutex
}

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

func (d *MongoDriver) NumClient() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()
	return len(d.Clients)
}

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
