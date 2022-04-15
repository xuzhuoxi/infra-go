package regexpx

import (
	"fmt"
	"regexp"
	"testing"
)

var (
	zeroStrArr = []string{"0", "1", "00", "01", "0A"}
)

var (
	digStrArr = []string{"0", "a"}
)

func TestZero(t *testing.T) {
	testRegexp(t, Zero, zeroStrArr)
}

func TestDigital(t *testing.T) {
	testRegexp(t, Digital, digStrArr)
}

func TestConst(t *testing.T) {
	html := `<link rel="stylesheet" href="https://static.studygolang.com/cosmo_bootstrap.min.css">`
	html2 := `<ul class="nav navbar-nav navbar-right" id="userbar"><li class="first"><a href="/account/register">注册</a></li><li class="last"><a href="/account/login">登录</a></li></ul>`
	fmt.Println(regexp.MatchString(HTML, html))
	fmt.Println(regexp.MustCompile(HTML).FindStringSubmatch(html2))
	fmt.Println("---")
	chi := "你好"
	fmt.Println(regexp.MatchString(Chinese, chi))
}

func testRegexp(t *testing.T, regStr string, data []string) {
	reg, err := regexp.Compile(regStr)
	if nil != err {
		t.Fatal(err)
		return
	}
	for _, str := range data {
		fmt.Println(str, "\t:", reg.MatchString(str))
	}
}
