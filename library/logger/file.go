package logger

import (
	"fmt"
	"framework/config"
	"framework/library/utli"
	"log"
	"os"
	"strings"
	"time"
)

var (
	LogSavePath   = config.GetRealPath("setting.log_path") //runtime/logs/
	LogSaveName   = "log"
	SqlSaveName   = "sql"
	LogFileExt    = "log"
	TimeFormat    = "20060102"
	DefaultPrefix = ""
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath) + utli.Date("Ymd") + "/"
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath() + LogSaveName + "/"
	utli.MkdirAll(prefixPath)
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func getSqlFileFullPath() string {
	prefixPath := getLogFilePath() + SqlSaveName + "/"
	utli.MkdirAll(prefixPath)
	suffixPath := fmt.Sprintf("%s%s.%s", SqlSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
	//	log.Fatalf("Dir not exist :%v", err)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	// 单个日志文件太大时，拆分
	fsize, _ := utli.FileSize(filePath)
	if fsize/1024/1024 >= 2 {
		t := utli.Date("His")
		newName := strings.Replace(filePath, ".log", "_"+t+".log", -1)
		os.Rename(filePath, newName)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}
