package goToolLog

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	LevelDebug = uint32(0)
	LevelInfo  = uint32(1)
	LevelWarn  = uint32(2)
	LevelError = uint32(3)
)

const (
	LevelHeaderDebug = "[DEBUG]"
	LevelHeaderInfo  = "[INFO]"
	LevelHeaderWarn  = "[WARN]"
	LevelHeaderError = "[Error]"
)

var Prefix = ""
var Suffix = ""

var StdOut bool

var Path string

var Level uint32
var fileLock *sync.RWMutex

func init() {
	fileLock = new(sync.RWMutex)
	Level = LevelWarn
	StdOut = false
}

func Debug(msg string) {
	if Level <= LevelDebug {
		go log(msg, LevelHeaderDebug)
	}
}

func Info(msg string) {
	if Level <= LevelInfo {
		log(msg, LevelHeaderInfo)
	}
}

func Warn(msg string) {
	if Level <= LevelWarn {
		log(msg, LevelHeaderWarn)
	}
}

func Error(msg string) {
	if Level <= LevelError {
		log(msg, LevelHeaderError)
	}
}

func log(msg string, header string) {
	msg = header + "" + goToolCommon.GetDateTimeStr(time.Now()) + " " + msg + goToolCommon.GetWrapStr()
	if StdOut {
		fmt.Print(msg)
	}
	path := getLogPath()
	err := goToolCommon.CheckAndCreateFolder(path)
	if err != nil {
		fmt.Println(err)
	}
	fileName := getLogFileName()

	fileLock.Lock()
	defer fileLock.Unlock()
	f, err := os.OpenFile(path+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Println(err)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func getLogPath() string {
	if strings.Trim(Path, " ") == "" {
		path, err := goToolCommon.GetCurrPath()
		if err != nil {
			return ""
		}
		return path + "\\" + "log" + "\\"
	}
	if !strings.HasSuffix(Path, "\\") {
		Path = Path + "\\"
	}
	return Path
}

func getLogFileName() string {
	return Prefix + goToolCommon.GetDateStr(time.Now()) + Suffix + ".log"
}
