package logx

import (
	"github.com/xuzhuoxi/util-go/mathx"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger()
	dirDaily := "D:\\log\\Daily\\"
	dirRolling := "D:\\log\\Rolling\\"
	dirDailyRolling := "D:\\log\\DailyRolling\\"
	l.SetConfig(TypeDailyFile, LevelAll, dirDaily, "L", ".txt", 0)
	l.SetConfig(TypeRollingFile, LevelAll, dirRolling, "L", ".txt", 2*mathx.KB)
	l.SetConfig(TypeDailyRollingFile, LevelAll, dirDailyRolling, "L", ".txt", 2*mathx.KB)
	//for {
	for i := 0; i < 300; i++ {
		l.Traceln("哈哈", "1111111111111111111111")
	}
}
