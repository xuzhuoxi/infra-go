package net

import (
	"testing"
)

func TestGet(t *testing.T) {
	Get("http://localhost:8080/", nil)
}

func TestPost(t *testing.T) {
	PostString("http://localhost:8080/", "", nil)
}
