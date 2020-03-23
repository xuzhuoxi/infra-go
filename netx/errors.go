package netx

import "errors"

type EmptyAddrError string

func (e EmptyAddrError) Error() string   { return "EmptyAddrError Address At:" + string(e) }
func (e EmptyAddrError) Timeout() bool   { return false }
func (e EmptyAddrError) Temporary() bool { return false }

type NoAddrError string

func (e NoAddrError) Error() string   { return "NoAddrError At:" + string(e) }
func (e NoAddrError) Timeout() bool   { return false }
func (e NoAddrError) Temporary() bool { return false }

type ConnNilError string

func (e ConnNilError) Error() string   { return "Conn Is Nil At:" + string(e) }
func (e ConnNilError) Timeout() bool   { return false }
func (e ConnNilError) Temporary() bool { return false }

type ReaderUnSupportError string

func (e ReaderUnSupportError) Error() string   { return "Reader Is Not Support At:" + string(e) }
func (e ReaderUnSupportError) Timeout() bool   { return false }
func (e ReaderUnSupportError) Temporary() bool { return false }

var ErrHttpServerStarted = errors.New("http: Server started")
var ErrSockServerStarted = errors.New("socket: Server started")
