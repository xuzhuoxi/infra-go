package netx

import (
	"net/http"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	server := NewHttpServer()
	server.MapFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(time.RFC1123)
		w.Write([]byte("The time is: " + tm))
	})
	server.StartServer(":9000")
}
