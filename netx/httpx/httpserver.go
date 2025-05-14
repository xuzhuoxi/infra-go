package httpx

import (
	"github.com/xuzhuoxi/infra-go/netx"
	"net/http"
	"time"
)

func NewHttpServer() IHttpServer {
	rs := &HttpServer{Network: "http", ServeMux: http.NewServeMux(), Server: nil}
	//rs.MapFunc("/test", testHandler)
	return rs
}

type IHttpServer interface {
	// Running
	// 启动中
	Running() bool
	// StartServer
	// 启动Http服务
	StartServer(addr string) error
	// StopServer
	// 停止Http服务
	StopServer() error
	// MapHandle
	// 映射请求响应处理器
	MapHandle(pattern string, handler http.Handler)
	// MapLimitHandle
	// 映射请求响应处理器,
	// post=true, 强制只接受POST请求
	// post=false, 强制只接受GET请求
	MapLimitHandle(pattern string, post bool, handler http.Handler)
	// MapFunc
	// 映射请求响应函数
	MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request))
	// MapLimitFunc
	// 映射请求响应函数,
	// post=true, 强制只接受POST请求
	// post=false, 强制只接受GET请求
	MapLimitFunc(pattern string, post bool, f func(w http.ResponseWriter, r *http.Request))
	// ClearMaps
	// 清除映射
	ClearMaps()
}

type HttpServer struct {
	Network  string
	ServeMux *http.ServeMux
	Server   *http.Server
}

func (s *HttpServer) Running() bool {
	return s.Server != nil
}

func (s *HttpServer) StartServer(addr string) error {
	if nil == s.ServeMux {
		s.ServeMux = http.NewServeMux()
	}
	if nil != s.Server { //启动中
		return netx.ErrHttpServerStarted
	}
	s.Server = &http.Server{Addr: addr, Handler: s.ServeMux}
	err := s.Server.ListenAndServe()
	s.Server = nil
	return err
}

func (s *HttpServer) StopServer() error {
	if nil == s.Server {
		return http.ErrServerClosed
	}
	s.Server.Close()
	s.Server = nil
	return nil
}

func (s *HttpServer) MapHandle(pattern string, handler http.Handler) {
	if nil == handler {
		return
	}
	s.ServeMux.Handle(pattern, handler)
}

func (s *HttpServer) MapLimitHandle(pattern string, post bool, handler http.Handler) {
	if nil == handler {
		return
	}
	s.ServeMux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		if post == ("POST" == request.Method) {
			handler.ServeHTTP(writer, request)
		}
	})
}

func (s *HttpServer) MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	if nil == f {
		return
	}
	s.ServeMux.HandleFunc(pattern, f)
}

func (s *HttpServer) MapLimitFunc(pattern string, post bool, f func(w http.ResponseWriter, r *http.Request)) {
	if nil == f {
		return
	}
	s.ServeMux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		if post == ("POST" == request.Method) {
			f(writer, request)
		}
	})
}

func (s *HttpServer) ClearMaps() {
	s.ServeMux = http.NewServeMux()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
