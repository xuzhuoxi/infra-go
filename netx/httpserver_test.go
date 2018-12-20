package netx

import "testing"

func TestStartServer(t *testing.T) {
	NewHttpServer().StartServer(":9999")
}
