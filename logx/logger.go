package logx

import (
	"github.com/xuzhuoxi/go-util/mathx"
	"github.com/xuzhuoxi/go-util/osxu"
	"github.com/xuzhuoxi/go-util/stringsx"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

type LogLevel uint8
type LogType uint8

const (
	LevelAll LogLevel = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

const (
	TypeConsole LogType = iota
	TypeRollingFile
	TypeDailyFile
	TypeDailyRollingFile
)

var level2prefix = make(map[LogLevel]string)

var defaultLogger ILogger = NewLogger()

func init() {
	level2prefix[LevelTrace] = "[Trace]"
	level2prefix[LevelDebug] = "[Debug]"
	level2prefix[LevelInfo] = "[Info]"
	level2prefix[LevelWarn] = "[Warn]"
	level2prefix[LevelError] = "[Error]"
	level2prefix[LevelFatal] = "[Fatal]"
}

func NewLogger() ILogger {
	instance := &logger{prefix: "", defaultFlag: log.LstdFlags, infoMap: make(map[LogType]*logInfo)}
	instance.SetConfig(TypeConsole, LevelAll, "", "", "", 0)
	return instance
}

type logInfo struct {
	level       LogLevel
	fileDir     string //以"/"结尾
	fileName    string
	fileExtName string
	file        *os.File
	logger      *log.Logger
	index       int
	maxSize     uint64
}

type ILogger interface {
	//设置日志前缀
	SetPrefix(prefix string)
	//设置日志等级，只有大于等级设置等级的日志才会记录
	//重置日志等级为level,t为空时重置全部
	SetLevel(level LogLevel, t ...LogType)
	// SetFlags sets the output flags for the logger.
	//重置日志flag,t为空时重置全部
	SetFlags(flag int, t ...LogType)
	//配置Log,要求fileDir以"/"结尾
	SetConfig(t LogType, level LogLevel, fileDir, fileName, fileExtName string, maxSize mathx.SizeUint)
	//移除配置
	RemoveConfig(t LogType)

	Log(level LogLevel, v ...interface{})
	Logf(level LogLevel, format string, v ...interface{})
	Logln(level LogLevel, v ...interface{})

	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Traceln(v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

type logger struct {
	prefix      string
	defaultFlag int

	mu      sync.RWMutex
	infoMap map[LogType]*logInfo
}

func (l *logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

func (l *logger) SetLevel(level LogLevel, t ...LogType) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(t) == 0 {
		for _, value := range l.infoMap {
			value.level = level
		}
	} else {
		for _, tp := range t {
			value, ok := l.infoMap[tp]
			if ok {
				value.level = level
			}
		}
	}
}

func (l *logger) SetFlags(flag int, t ...LogType) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.defaultFlag = flag
	if len(t) == 0 {
		for _, value := range l.infoMap {
			value.logger.SetFlags(flag)
		}
	} else {
		for _, tp := range t {
			value, ok := l.infoMap[tp]
			if ok {
				value.logger.SetFlags(flag)
			}
		}
	}
}

func (l *logger) SetConfig(t LogType, level LogLevel, fileDir, fileName, fileExtName string, maxSize mathx.SizeUint) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if "" == fileDir {
		return
	}
	newFileDir := osxu.GetUnitePath(fileDir)
	if stringsx.GetCharCount(newFileDir)-1 != stringsx.LastIndexOfChar(newFileDir, "/") { //保证最后一个为"/"
		newFileDir = newFileDir + "/"
	}
	val, ok := l.infoMap[t]
	if ok {
		val.level = level
		val.fileDir = newFileDir
		val.fileName = fileName
		val.fileExtName = fileExtName
		val.maxSize = uint64(maxSize)
	} else {
		l.infoMap[t] = &logInfo{level: level, fileDir: newFileDir, fileName: fileName, fileExtName: fileExtName, maxSize: uint64(maxSize), logger: genLogger(l.defaultFlag)}
	}
	switch t {
	case TypeRollingFile:
		l.infoMap[t].index = getRollingIndex(newFileDir, fileName, fileExtName)
	case TypeDailyRollingFile:
		l.infoMap[t].index = getRollingIndex(newFileDir, fileName+"_"+getTodayStr(), fileExtName)
	}
}

func (l *logger) RemoveConfig(t LogType) {
	l.mu.Lock()
	defer l.mu.Unlock()
	val, ok := l.infoMap[t]
	if ok && nil != val.file {
		closeFile(val.file)
	}
	delete(l.infoMap, t)
}

func (l *logger) Log(level LogLevel, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
		if level < info.level {
			continue
		}
		info.logger.SetPrefix(prefix)
		info.logger.Print(v...)
	}
}

func (l *logger) Logf(level LogLevel, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
		if level < info.level {
			continue
		}
		info.logger.SetPrefix(prefix)
		info.logger.Printf(format, v...)
	}
}

func (l *logger) Logln(level LogLevel, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
		if level < info.level {
			continue
		}
		info.logger.SetPrefix(prefix)
		info.logger.Println(v...)
	}
}

func (l *logger) Trace(v ...interface{}) {
	l.Log(LevelTrace, v...)
}

func (l *logger) Tracef(format string, v ...interface{}) {
	l.Logf(LevelTrace, format, v...)
}

func (l *logger) Traceln(v ...interface{}) {
	l.Logln(LevelTrace, v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.Log(LevelDebug, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.Logf(LevelDebug, format, v...)
}

func (l *logger) Debugln(v ...interface{}) {
	l.Logln(LevelDebug, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.Log(LevelInfo, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.Logf(LevelInfo, format, v...)
}

func (l *logger) Infoln(v ...interface{}) {
	l.Logln(LevelInfo, v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.Log(LevelWarn, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.Logf(LevelWarn, format, v...)
}

func (l *logger) Warnln(v ...interface{}) {
	l.Logln(LevelWarn, v...)
}

func (l *logger) Error(v ...interface{}) {
	l.Log(LevelError, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.Logf(LevelError, format, v...)
}

func (l *logger) Errorln(v ...interface{}) {
	l.Logln(LevelError, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	l.Log(LevelFatal, v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.Logf(LevelFatal, format, v...)
}

func (l *logger) Fatalln(v ...interface{}) {
	l.Logln(LevelFatal, v...)
}

//private--------------------------------------

func genLogger(flag int) *log.Logger {
	return log.New(os.Stderr, "", flag)
}

func getLevelPrefix(level LogLevel, prefix string) string {
	rs, ok := level2prefix[level]
	if ok {
		return rs + prefix
	}
	return "" + prefix
}

func checkFile(infoMap map[LogType]*logInfo) {
	for key, value := range infoMap {
		switch key {
		case TypeDailyFile:
			checkDailyFile(value)
		case TypeRollingFile:
			checkRollingFile(value, false)
		case TypeDailyRollingFile:
			checkRollingFile(value, true)
		}
	}
}

func checkDailyFile(info *logInfo) {
	todayStr := getTodayStr()
	newFileName := info.fileName + "_" + todayStr
	fileFullPath := getFullPath(info.fileDir, newFileName, info.fileExtName, "")
	if nil != info.file {
		if info.file.Name() == newFileName+info.fileExtName { //同一个文件
			return
		}
		closeFile(info.file)
		info.file = nil
	}
	updateLoggerOutput(info, fileFullPath)
}

func checkRollingFile(info *logInfo, daily bool) {
	newFileName := info.fileName
	if daily {
		newFileName = info.fileName + "_" + getTodayStr()
	}
	fileFullPath := getFullPath(info.fileDir, newFileName, info.fileExtName, "")
	isFull := checkCloseFullFile(info.file, fileFullPath, info.maxSize)
	if isFull {
		newName := getFullPath(info.fileDir, newFileName, info.fileExtName, "_"+strconv.Itoa(info.index))
		os.Rename(fileFullPath, newName)
		info.file = nil
		info.index++
	}
	updateLoggerOutput(info, fileFullPath)
}

func updateLoggerOutput(info *logInfo, fileFullPath string) {
	if nil != info.file {
		return
	}
	file, err := openFile(fileFullPath)
	if nil != err {
		return
	}
	info.file = file
	info.logger.SetOutput(file)
}

func checkCloseFullFile(file *os.File, fileFullPath string, maxSize uint64) bool {
	if nil != file {
		if checkFileSizeFull(file, maxSize) {
			closeFile(file)
			return true
		}
	} else {
		if osxu.IsExist(fileFullPath) {
			size, _ := osxu.GetFileSize(fileFullPath)
			return size >= maxSize
		}
	}
	return false
}

func getRollingIndex(fileDir, fileName, fileExtName string) int {
	for index := 0; index < math.MaxInt16; index++ {
		path := getFullPath(fileDir, fileName, fileExtName, "_"+strconv.Itoa(index))
		if !osxu.IsExist(path) {
			return index
		}
	}
	return -1
}

func checkFileSizeFull(file *os.File, maxSize uint64) bool {
	fileStat, _ := file.Stat()
	return maxSize <= uint64(fileStat.Size())
}

func openFile(fileFullPath string) (file *os.File, err error) {
	return os.OpenFile(fileFullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
}

func closeFile(file *os.File) {
	if nil == file {
		return
	}
	defer file.Close()
	file.Sync()
}

func getFullPath(fileDir string, fileName string, fileExtName string, other string) string {
	return fileDir + fileName + other + fileExtName
}

func getTodayStr() string {
	return time.Now().Format("20060102")
}
