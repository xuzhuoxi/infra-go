package netx

import (
	"log"
	"net"
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
	serverMu sync.Mutex
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
	s.serverMu.Lock()
	defer s.serverMu.Unlock()
	return s.running
}

type SockClientBase struct {
	Name    string
	Network string
	baseMu  sync.RWMutex
	opening bool

	localAddress string
	conn         net.Conn
	messageProxy IMessageSendReceiver
}

func (c *SockClientBase) LocalAddress() string {
	return c.conn.LocalAddr().String()
}

func (c *SockClientBase) SetSplitHandler(handler func(buff []byte) ([]byte, []byte)) error {
	c.messageProxy.SetSplitHandler(handler)
	return nil
}

func (c *SockClientBase) SetMessageHandler(handler func(msgBytes []byte, info interface{})) error {
	c.messageProxy.SetMessageHandler(handler)
	return nil
}

func (c *SockClientBase) IsReceiving() bool {
	return c.messageProxy.IsReceiving()
}

func (c *SockClientBase) Opening() bool {
	c.baseMu.RLock()
	defer c.baseMu.RUnlock()
	return c.opening
}

func (c *SockClientBase) SendDataTo(msg []byte, rAddress ...string) error {
	_, err := c.messageProxy.SendMessage(msg, rAddress...)
	return err
}

func (c *SockClientBase) StartReceiving() error {
	log.Println(c.Name + ".StartReceiving()")
	err := c.messageProxy.StartReceiving()
	return err
}

func (c *SockClientBase) StopReceiving() error {
	log.Println(c.Name + ".StopReceiving()")
	err := c.messageProxy.StopReceiving()
	return err
}
