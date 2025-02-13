package httpx

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	HttpGet("http://localhost:9000/test", func(res *http.Response, body *[]byte) {
		fmt.Println(res.StatusCode, string(*body))
	})
}

func TestPost(t *testing.T) {
	HttpPostString("http://localhost:9000/test", "", func(res *http.Response, body *[]byte) {
		fmt.Println(res.StatusCode, string(*body))
	})
}
