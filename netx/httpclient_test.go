package netx

import (
	"testing"
)

func TestGet(t *testing.T) {
	HttpGet("http://localhost:8080/", nil)
}

func TestPost(t *testing.T) {
	HttpPostString("http://localhost:8080/", "", nil)
}
