package mathx

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(EB)
}

func TestSystem26To10(t *testing.T) {
	check := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ",
		"AR", "AS", "AT", "AU", "AV", "AW"}
	for _, s := range check {
		fmt.Println(System26To10(s))
	}
}

func TestSystem10To26(t *testing.T) {
	start := 0
	for start < 50 {
		fmt.Println("\"" + System10To26(start) + "\",")
		start++
	}
}

func TestSystemTo(t *testing.T) {
	start := 0
	system := []int{8, 8, 8}
	for start < 50 {
		fmt.Println(SystemTo(start, system))
		start++
	}
}
