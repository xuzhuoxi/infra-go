package netx

import (
	"net/http"
	"time"
)

func NewHttpServer() IHttpServer {
	rs := &HttpServer{Network: "http", ServeMux: http.NewServeMux(), Server: nil}
	rs.MapFunc("/test", testHandler)
	return rs
}

type IHttpServer interface {
	// 启动Http服务
	StartServer(addr string)
	// 停止Http服务
	StopServer()
	// 映射请求响应处理器
	MapHandle(pattern string, handler http.Handler)
	// 映射请求响应函数
	MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request))
}

type HttpServer struct {
	Network  string
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

func testHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
