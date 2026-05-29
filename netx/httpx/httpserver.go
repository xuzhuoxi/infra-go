package httpx

import (
	"net/http"
	"time"

	"github.com/xuzhuoxi/infra-go/netx"
)

// NewHttpServer creates a new HttpServer with an empty ServeMux.
// NewHttpServer 创建带空 ServeMux 的 HttpServer。
func NewHttpServer() IHttpServer {
	rs := &HttpServer{Network: "http", ServeMux: http.NewServeMux(), Server: nil}
	//rs.MapFunc("/test", testHandler)
	return rs
}

// IHttpServer defines HTTP server lifecycle and route mapping.
// IHttpServer 定义 HTTP 服务端生命周期与路由映射。
type IHttpServer interface {
	// Running reports whether the server is running.
	// Running 是否启动中。
	Running() bool
	// StartServer starts the HTTP server on addr.
	// StartServer 启动 Http 服务。
	StartServer(addr string) error
	// StopServer stops the HTTP server.
	// StopServer 停止 Http 服务。
	StopServer() error
	// MapHandle maps pattern to an http.Handler.
	// MapHandle 映射请求响应处理器。
	MapHandle(pattern string, handler http.Handler)
	// MapLimitHandle maps pattern to a handler restricted by HTTP method.
	// MapLimitHandle 映射请求响应处理器。
	//   post: true accepts POST only; false accepts GET only
	//   post: true 强制只接受 POST 请求；false 强制只接受 GET 请求
	MapLimitHandle(pattern string, post bool, handler http.Handler)
	// MapFunc maps pattern to a handler function.
	// MapFunc 映射请求响应函数。
	MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request))
	// MapLimitFunc maps pattern to a function restricted by HTTP method.
	// MapLimitFunc 映射请求响应函数。
	//   post: true accepts POST only; false accepts GET only
	//   post: true 强制只接受 POST 请求；false 强制只接受 GET 请求
	MapLimitFunc(pattern string, post bool, f func(w http.ResponseWriter, r *http.Request))
	// ClearMaps resets all route mappings.
	// ClearMaps 清除映射。
	ClearMaps()
}

// HttpServer wraps http.Server and http.ServeMux for route management.
// HttpServer 封装 http.Server 与 http.ServeMux 以管理路由。
type HttpServer struct {
	Network  string
	ServeMux *http.ServeMux
	Server   *http.Server
}

// Running reports whether the server is running.
// Running 是否启动中。
func (s *HttpServer) Running() bool {
	return s.Server != nil
}

// StartServer starts listening on addr and blocks until the server stops.
// StartServer 在 addr 上启动监听并阻塞直至服务停止。
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

// StopServer closes the HTTP server.
// StopServer 停止 Http 服务。
func (s *HttpServer) StopServer() error {
	if nil == s.Server {
		return http.ErrServerClosed
	}
	s.Server.Close()
	s.Server = nil
	return nil
}

// MapHandle maps pattern to an http.Handler.
// MapHandle 映射请求响应处理器。
func (s *HttpServer) MapHandle(pattern string, handler http.Handler) {
	if nil == handler {
		return
	}
	s.ServeMux.Handle(pattern, handler)
}

// MapLimitHandle maps pattern to a method-restricted handler.
// MapLimitHandle 映射请求响应处理器。
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

// MapFunc maps pattern to a handler function.
// MapFunc 映射请求响应函数。
func (s *HttpServer) MapFunc(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	if nil == f {
		return
	}
	s.ServeMux.HandleFunc(pattern, f)
}

// MapLimitFunc maps pattern to a method-restricted handler function.
// MapLimitFunc 映射请求响应函数。
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

// ClearMaps resets all route mappings.
// ClearMaps 清除映射。
func (s *HttpServer) ClearMaps() {
	s.ServeMux = http.NewServeMux()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
