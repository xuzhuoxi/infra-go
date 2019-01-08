package mathx

import (
	"fmt"
	"testing"
)

func TestBitString(t *testing.T) {
	start := 0
	for start < 50 {
		fmt.Println(BitString(start))
		start++
	}
}

func TestBitValid(t *testing.T) {
	start := 0
	for start < 50 {
		var index uint = 0
		for index < 8 {
			fmt.Println(BitString(start), index, ":", BitValid(start, index))
			index++
		}
		start++
	}
}

func TestBitFit(t *testing.T) {
	value := 50
	fmt.Println(BitString(value))
	start := 0
	for start < 50 {
		fmt.Println(BitString(start), BitFit(value, start, true), BitFit(value, start, false))
		start++
	}
}
