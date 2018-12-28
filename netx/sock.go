package netx

import (
	"github.com/xuzhuoxi/util-go/logx"
	"sync"
)

type SockParams struct {
	Network       string
	LocalAddress  string
	RemoteAddress string
}

type SockServerBase struct {
	Name     string
	Network  string
	serverMu sync.RWMutex
	running  bool

	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(msgBytes []byte, info interface{})
}

func (s *SockServerBase) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	s.splitHandler = handler
	return nil
}

func (s *SockServerBase) SetMessageHandler(handler func(msgBytes []byte, info interface{})) error {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	s.messageHandler = handler
	return nil
}

func (s *SockServerBase) Running() bool {
	s.serverMu.RLock()
	defer s.serverMu.RUnlock()
	return s.running
}

type SockClientBase struct {
	Name     string
	Network  string
	clientMu sync.RWMutex
	opening  bool

	splitHandler   func(buff []byte) ([]byte, []byte)
	messageHandler func(msgBytes []byte, info interface{})

	localAddress string
	conn         ISockConn
	messageProxy IMessageSendReceiver
}

func (c *SockClientBase) LocalAddress() string {
	return c.conn.LocalAddr().String()
}

func (c *SockClientBase) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	c.splitHandler = handler
	if nil != c.messageProxy {
		c.messageProxy.SetSplitHandler(handler)
	}
	return nil
}

func (c *SockClientBase) SetMessageHandler(handler func(msgBytes []byte, info interface{})) error {
	c.messageHandler = handler
	if nil != c.messageProxy {
		c.messageProxy.SetMessageHandler(handler)
	}
	return nil
}

func (c *SockClientBase) IsReceiving() bool {
	return c.messageProxy.IsReceiving()
}

func (c *SockClientBase) Opening() bool {
	c.clientMu.RLock()
	defer c.clientMu.RUnlock()
	return c.opening
}

func (c *SockClientBase) SendDataTo(msg []byte, rAddress ...string) error {
	_, err := c.messageProxy.SendMessage(msg, rAddress...)
	return err
}

func (c *SockClientBase) StartReceiving() error {
	logx.Infoln(c.Name + ".StartReceiving()")
	err := c.messageProxy.StartReceiving()
	return err
}

func (c *SockClientBase) StopReceiving() error {
	logx.Infoln(c.Name + ".StopReceiving()")
	err := c.messageProxy.StopReceiving()
	return err
}

func (c *SockClientBase) setMessageProxy(messageProxy IMessageSendReceiver) {
	c.messageProxy = messageProxy
	if nil != c.splitHandler {
		messageProxy.SetSplitHandler(c.splitHandler)
	}
	if nil != c.messageHandler {
		messageProxy.SetMessageHandler(c.messageHandler)
	}
}
