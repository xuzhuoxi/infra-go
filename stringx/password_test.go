package stringx

import (
	"fmt"
	"testing"
)

func TestPasswordCheck(t *testing.T) {
	pwds := []string{"111", "ssss", "SSSSS", "?><>"}
	//flags := []PasswdFlag{N, U, L, S, DefaultPasswdFlag}
	for _, pwd := range pwds {
		fmt.Println(pwd, PasswordCheck(pwd, DefaultPasswdFlag, 1, 4))
		//for _, flag := range flags {
		//	fmt.Println(pwd, PasswordCheck(pwd, flag, 1, 4))
		//}
	}
}
