package logx

import (
	"github.com/xuzhuoxi/go-util/mathx"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger()
	dirDaily := "D:\\log\\Daily\\"
	dirRolling := "D:\\log\\Rolling\\"
	dirDailyRolling := "D:\\log\\DailyRolling\\"
	l.RemoveConfig(TypeConsole)
	l.SetConfig(TypeDailyFile, dirDaily, "L", ".txt", 0)
	l.SetConfig(TypeRollingFile, dirRolling, "L", ".txt", 2*mathx.KB)
	l.SetConfig(TypeDailyRollingFile, dirDailyRolling, "L", ".txt", 2*mathx.KB)
	//for {
	for i := 0; i < 300; i++ {
		l.Traceln("哈哈", "1111111111111111111111")
	}
}
