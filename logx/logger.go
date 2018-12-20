package logx

import (
	"github.com/xuzhuoxi/go-util/mathx"
	"github.com/xuzhuoxi/go-util/osxu"
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
	instance := &logger{level: LevelAll, prefix: "", flag: log.LstdFlags, infoMap: make(map[LogType]*fileInfo)}
	instance.SetConfig(TypeConsole, "", "", "", 0)
	return instance
}

type fileInfo struct {
	fileDir     string //以"/"结尾
	fileName    string
	fileExtName string
	file        *os.File
	logger      *log.Logger
	index       int
	maxSize     uint64
}

type ILogger interface {
	//设置日志等级，只有大于等级设置等级的日志才会记录
	SetLevel(level LogLevel)
	//设置每一行log的时间格式
	SetPrefix(prefix string)
	// SetFlags sets the output flags for the logger.
	SetFlags(flag int)

	//配置Log,要求fileDir以"/"结尾
	SetConfig(t LogType, fileDir, fileName, fileExtName string, maxSize mathx.SizeUint)
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
	level  LogLevel
	prefix string
	flag   int

	mu      sync.RWMutex
	infoMap map[LogType]*fileInfo
}

func (l *logger) SetConfig(t LogType, fileDir, fileName, fileExtName string, maxSize mathx.SizeUint) {
	l.mu.Lock()
	defer l.mu.Unlock()
	newFileDir := osxu.GetUnitePath(fileDir)
	val, ok := l.infoMap[t]
	if ok {
		val.fileDir = newFileDir
		val.fileName = fileName
		val.fileExtName = fileExtName
		val.maxSize = uint64(maxSize)
	} else {
		l.infoMap[t] = &fileInfo{fileDir: newFileDir, fileName: fileName, fileExtName: fileExtName, maxSize: uint64(maxSize), logger: genLogger(l.flag)}
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

func (l *logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

func (l *logger) SetFlags(flag int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flag = flag
	for _, value := range l.infoMap {
		value.logger.SetFlags(flag)
	}
}

func (l *logger) Log(level LogLevel, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level || len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
		info.logger.SetPrefix(prefix)
		info.logger.Print(v...)
	}
}

func (l *logger) Logf(level LogLevel, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level || len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
		info.logger.SetPrefix(prefix)
		info.logger.Printf(format, v...)
	}
}

func (l *logger) Logln(level LogLevel, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level < l.level || len(l.infoMap) == 0 {
		return
	}
	checkFile(l.infoMap)
	prefix := getLevelPrefix(level, l.prefix)
	for _, info := range l.infoMap {
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
	newLog := log.New(os.Stderr, "", log.LstdFlags)
	newLog.SetFlags(flag)
	return newLog
}

func getLevelPrefix(level LogLevel, prefix string) string {
	rs, ok := level2prefix[level]
	if ok {
		return rs + prefix
	}
	return "" + prefix
}

func checkFile(infoMap map[LogType]*fileInfo) {
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

func checkDailyFile(fi *fileInfo) {
	todayStr := getTodayStr()
	newFileName := fi.fileName + "_" + todayStr
	fileFullPath := getFullPath(fi.fileDir, newFileName, fi.fileExtName, "")
	if nil != fi.file {
		if fi.file.Name() == newFileName+fi.fileExtName { //同一个文件
			return
		}
		closeFile(fi.file)
		fi.file = nil
	}
	updateLoggerOutput(fi, fileFullPath)
}

func checkRollingFile(fi *fileInfo, daily bool) {
	newFileName := fi.fileName
	if daily {
		newFileName = fi.fileName + "_" + getTodayStr()
	}
	fileFullPath := getFullPath(fi.fileDir, newFileName, fi.fileExtName, "")
	isFull := checkCloseFullFile(fi.file, fileFullPath, fi.maxSize)
	if isFull {
		newName := getFullPath(fi.fileDir, newFileName, fi.fileExtName, "_"+strconv.Itoa(fi.index))
		os.Rename(fileFullPath, newName)
		fi.file = nil
		fi.index++
	}
	updateLoggerOutput(fi, fileFullPath)
}

func updateLoggerOutput(fi *fileInfo, fileFullPath string) {
	if nil != fi.file {
		return
	}
	file, err := openFile(fileFullPath)
	if nil != err {
		return
	}
	fi.file = file
	fi.logger.SetOutput(file)
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
