package net

import (
	"net/http"
	"time"
)

type IHttpServer interface {
	StartServer(addr string)
	StopServer()
	MapHandle(pattern string, handler http.Handler)
	MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request))
}

type HttpServer struct {
	ServeMux *http.ServeMux
	Server   *http.Server
}

func (s *HttpServer) StartServer(addr string) {
	if nil != s.Server {
		return
	}
	if nil == s.ServeMux {
		s.ServeMux = http.NewServeMux()
	}
	s.Server = &http.Server{Addr: addr, Handler: s.ServeMux}
	s.Server.ListenAndServe()
}

func (s *HttpServer) StopServer() {
	s.Server.Close()
	s.Server = nil
}

func (s *HttpServer) MapHandle(pattern string, handler http.Handler) {
	s.ServeMux.Handle(pattern, handler)
}

func (s *HttpServer) MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	s.MapHandle(pattern, http.HandlerFunc(f))
}

func NewHttpServer() IHttpServer {
	rs := &HttpServer{http.NewServeMux(), nil}
	rs.MapFunc("/time", timeHandler)
	return rs
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
