package stringx

import (
	"fmt"
	"testing"
)

func TestPasswordCheck(t *testing.T) {
	pwds := []string{"1116666", "ssss", "SSSSSSSS", "?><>?><>", "☺☺☺"}
	for index, pwd := range pwds {
		rs1 := CheckPassword(pwd, N, 6, 8)
		fmt.Println(index, pwd, rs1)
		rs2 := CheckPassword(pwd, LOrU, 6, 8)
		fmt.Println(index, pwd, rs2)
		rs3 := CheckPassword(pwd, DefaultPasswdFlag, 6, 8)
		fmt.Println(index, pwd, rs3)
	}

	// Output:
	// 0 1116666 true
	// 0 1116666 false
	// 0 1116666 true
	// 1 ssss false
	// 1 ssss false
	// 1 ssss false
	// 2 SSSSSSSS false
	// 2 SSSSSSSS true
	// 2 SSSSSSSS true
	// 3 ?><>?><> false
	// 3 ?><>?><> false
	// 3 ?><>?><> false
	// 4 ☺☺☺ false
	// 4 ☺☺☺ false
	// 4 ☺☺☺ false
}
