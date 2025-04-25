package logx

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

type logFormats struct {
	prefix string // 显示前缀
	format string // 显示颜色
}

func (o *logFormats) String() string {
	return fmt.Sprintf("logFromats{prefix:%v, format:%s}", o.prefix, o.format)
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

func (o *logInfo) String() string {
	return fmt.Sprintf("logInfo{level:%v,}", o.level)
}

type logger struct {
	prefix      string
	defaultFlag int

	mu      sync.RWMutex
	infoMap map[LogType]*logInfo
}

func (o *logger) String() string {
	return fmt.Sprintf("logger{prefix:%s, flag:%v, infoMap:%v}", o.prefix, o.defaultFlag, o.infoMap)
}

var (
	defaultLogger     = NewLogger()
	levelFormats      = make(map[LogLevel]logFormats)
	defaultFormat     = "\033[0m"
	enableColorFormat = false
)

func init() {
	defaultLogger.SetConfig(LogConfig{Type: TypeConsole, Level: LevelAll})
	levelFormats[LevelTrace] = logFormats{prefix: "[Trace] ", format: "\033[0m"}  // Default
	levelFormats[LevelDebug] = logFormats{prefix: "[Debug] ", format: "\033[95m"} // FG: Bright Cyan
	levelFormats[LevelInfo] = logFormats{prefix: "[Info] ", format: "\033[94m"}   // FG: Green
	levelFormats[LevelWarn] = logFormats{prefix: "[Warn] ", format: "\033[93m"}   // FG: Bright Yellow
	levelFormats[LevelError] = logFormats{prefix: "[Error] ", format: "\033[91m"} // FG: Bright Red
	levelFormats[LevelFatal] = logFormats{prefix: "[Fatal] ", format: "\033[31m"} // FG: Red
}

func (o *logger) SetPrefix(prefix string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.prefix = prefix
}

func (o *logger) SetLevel(level LogLevel, t ...LogType) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(t) == 0 {
		for _, value := range o.infoMap {
			value.level = level
		}
	} else {
		for _, tp := range t {
			value, ok := o.infoMap[tp]
			if ok {
				value.level = level
			}
		}
	}
}

func (o *logger) SetFlags(flag int, t ...LogType) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.defaultFlag = flag
	if len(t) == 0 {
		for _, value := range o.infoMap {
			value.logger.SetFlags(flag)
		}
	} else {
		for _, tp := range t {
			value, ok := o.infoMap[tp]
			if ok {
				value.logger.SetFlags(flag)
			}
		}
	}
}

func (o *logger) SetConfig(cfg LogConfig) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if 0 == cfg.Flag {
		cfg.Flag = o.defaultFlag
	}
	if cfg.Type == TypeConsole {
		val, ok := o.infoMap[cfg.Type]
		if ok {
			val.level = cfg.Level
			val.logger.SetFlags(cfg.Flag)
		} else {
			o.infoMap[cfg.Type] = &logInfo{level: cfg.Level, logger: genLogger(o.defaultFlag)}
		}
		return
	}
	fileDir, fileName, fileExt := cfg.GetFileInfos()
	if "" == fileDir {
		return
	}
	newFileDir := filex.FormatPath(fileDir)
	if !filex.IsExist(newFileDir) { //目标不存在，创建目录
		os.MkdirAll(newFileDir, os.ModePerm)
	}
	val, ok := o.infoMap[cfg.Type]
	if ok {
		val.level = cfg.Level
		val.fileDir = newFileDir
		val.fileName = fileName
		val.fileExtName = fileExt
		val.maxSize = uint64(cfg.MaxSize)
		val.logger.SetFlags(cfg.Flag)
	} else {
		o.infoMap[cfg.Type] = &logInfo{level: cfg.Level,
			fileDir: newFileDir, fileName: fileName, fileExtName: fileExt,
			maxSize: uint64(cfg.MaxSize), logger: genLogger(o.defaultFlag)}
	}
	switch cfg.Type {
	case TypeRollingFile:
		o.infoMap[cfg.Type].index = getRollingIndex(newFileDir, cfg.FileName, cfg.FileExtName)
	case TypeDailyRollingFile:
		o.infoMap[cfg.Type].index = getRollingIndex(newFileDir, cfg.FileName+"_"+getTodayStr(), cfg.FileExtName)
	}
}

func (o *logger) RemoveConfig(t LogType) {
	o.mu.Lock()
	defer o.mu.Unlock()
	val, ok := o.infoMap[t]
	if ok && nil != val.file {
		closeFile(val.file)
	}
	delete(o.infoMap, t)
}

func (o *logger) Print(v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	for _, info := range o.infoMap {
		flag := info.logger.Flags()
		info.logger.SetFlags(0)
		info.logger.SetPrefix("")
		info.logger.Print(v...)
		info.logger.SetFlags(flag)
	}
}

func (o *logger) Printf(format string, v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	for _, info := range o.infoMap {
		flag := info.logger.Flags()
		info.logger.SetFlags(0)
		info.logger.SetPrefix("")
		info.logger.Printf(format, v...)
		info.logger.SetFlags(flag)
	}
}

func (o *logger) Println(v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	for _, info := range o.infoMap {
		flag := info.logger.Flags()
		info.logger.SetFlags(0)
		info.logger.SetPrefix("")
		info.logger.Println(v...)
		info.logger.SetFlags(flag)
	}
}

func (o *logger) Log(level LogLevel, v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	prefix := getLevelPrefix(level, o.prefix)
	logged := false
	for _, info := range o.infoMap {
		if level < info.level {
			continue
		}
		logged = true
		info.logger.SetPrefix(prefix)
		info.logger.Println(v...)
	}
	if logged && enableColorFormat {
		fmt.Printf(defaultFormat)
	}
}

func (o *logger) Logf(level LogLevel, format string, v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	prefix := getLevelPrefix(level, o.prefix)
	logged := false
	for _, info := range o.infoMap {
		if level < info.level {
			continue
		}
		logged = true
		info.logger.SetPrefix(prefix)
		info.logger.Printf(format, v...)
	}
	if logged && enableColorFormat {
		fmt.Printf(defaultFormat)
	}
}

func (o *logger) Logln(level LogLevel, v ...interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if len(o.infoMap) == 0 {
		return
	}
	checkFile(o.infoMap)
	prefix := getLevelPrefix(level, o.prefix)
	logged := false
	for _, info := range o.infoMap {
		if level < info.level {
			continue
		}
		logged = true
		info.logger.SetPrefix(prefix)
		info.logger.Println(v...)
	}
	if logged && enableColorFormat {
		fmt.Printf(defaultFormat)
	}
}

func (o *logger) Trace(v ...interface{}) {
	o.Log(LevelTrace, v...)
}

func (o *logger) Tracef(format string, v ...interface{}) {
	o.Logf(LevelTrace, format, v...)
}

func (o *logger) Traceln(v ...interface{}) {
	o.Logln(LevelTrace, v...)
}

func (o *logger) Debug(v ...interface{}) {
	o.Log(LevelDebug, v...)
}

func (o *logger) Debugf(format string, v ...interface{}) {
	o.Logf(LevelDebug, format, v...)
}

func (o *logger) Debugln(v ...interface{}) {
	o.Logln(LevelDebug, v...)
}

func (o *logger) Info(v ...interface{}) {
	o.Log(LevelInfo, v...)
}

func (o *logger) Infof(format string, v ...interface{}) {
	o.Logf(LevelInfo, format, v...)
}

func (o *logger) Infoln(v ...interface{}) {
	o.Logln(LevelInfo, v...)
}

func (o *logger) Warn(v ...interface{}) {
	o.Log(LevelWarn, v...)
}

func (o *logger) Warnf(format string, v ...interface{}) {
	o.Logf(LevelWarn, format, v...)
}

func (o *logger) Warnln(v ...interface{}) {
	o.Logln(LevelWarn, v...)
}

func (o *logger) Error(v ...interface{}) {
	o.Log(LevelError, v...)
}

func (o *logger) Errorf(format string, v ...interface{}) {
	o.Logf(LevelError, format, v...)
}

func (o *logger) Errorln(v ...interface{}) {
	o.Logln(LevelError, v...)
}

func (o *logger) Fatal(v ...interface{}) {
	o.Log(LevelFatal, v...)
}

func (o *logger) Fatalf(format string, v ...interface{}) {
	o.Logf(LevelFatal, format, v...)
}

func (o *logger) Fatalln(v ...interface{}) {
	o.Logln(LevelFatal, v...)
}

func genLogger(flag int) *log.Logger {
	return log.New(os.Stderr, "", flag)
}

func getLevelPrefix(level LogLevel, prefix string) (newPrefix string) {
	rs, ok := levelFormats[level]
	if ok {
		if enableColorFormat {
			return fmt.Sprintf("%s%s%s", rs.format, rs.prefix, prefix)
		} else {
			return fmt.Sprintf("%s%s", rs.prefix, prefix)
		}
	}
	return prefix
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
		if filex.IsExist(fileFullPath) {
			size, _ := filex.GetFileSize(fileFullPath)
			return size >= maxSize
		}
	}
	return false
}

func getRollingIndex(fileDir, fileName, fileExtName string) int {
	for index := 0; index < math.MaxInt16; index++ {
		path := getFullPath(fileDir, fileName, fileExtName, "_"+strconv.Itoa(index))
		if !filex.IsExist(path) {
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
	return fileDir + "/" + fileName + other + fileExtName
}

func getTodayStr() string {
	return time.Now().Format("20060102")
}
