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
	l.SetConfig(LogConfig{Type: TypeConsole, Level: LevelAll})
	l.SetConfig(LogConfig{Type: TypeDailyFile, Level: LevelAll, FileDir: dirDaily, FileName: "L", FileExtName: ".txt"})
	l.SetConfig(LogConfig{Type: TypeRollingFile, Level: LevelAll, FileDir: dirRolling, FileName: "L", FileExtName: ".txt", MaxSize: 2 * mathx.KB})
	l.SetConfig(LogConfig{Type: TypeDailyRollingFile, Level: LevelAll, FileDir: dirDailyRolling, FileName: "L", FileExtName: ".txt", MaxSize: 2 * mathx.KB})

	for i := 0; i < 300; i++ {
		l.Traceln("哈哈", "1111111111111111111111")
	}
}
