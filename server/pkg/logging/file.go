package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	FilePath   = "./runtime/logs/"
	FileName   = "pd_server"
	FileExt    = "log"
	TimeFormat = "20060102"
)

func getDirPath() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	return filepath.Join(dir, FilePath)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s_%s.%s", FilePath, FileName, time.Now().Format(TimeFormat), FileExt)
}

func InitLogger() {
	fmt.Println(getDirPath())
	fmt.Println(getLogFilePath())
}

//func openLogFile(filePath string) (file *os.File) {
//	stat, err := os.Stat(filePath)
//
//	if err != nil {
//		if os.IsNotExist(err) {
//			err := os.MkdirAll(filePath, os.ModePerm)
//		}
//	}
//
//	return
//}
