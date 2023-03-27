package speed

import (
	"io"
	"log"
	"os"
	"runtime"
)

var errlog *log.Logger
var infolog *log.Logger

func init() {
	if _, err := os.Stat("./store"); os.IsNotExist(err) {
		os.Mkdir("./storage", 0777)
	}
	file, err := os.OpenFile("./storage/logInfo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("无法创建logInfo.log文件", err)
	}
	errlog = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime)
	infolog = log.New(io.MultiWriter(file, os.Stderr), "_Info: ", log.Ldate|log.Ltime)
}
func LogErr(err ...interface{}) {
	writeData(errlog, err...)
}

func Log(data ...interface{}) {
	writeData(infolog, data...)
}

func writeData(logger *log.Logger, data ...interface{}) {
	projectPath, _ := os.Getwd()
	_, file, line, ok := runtime.Caller(2)
	if data[0] != nil && ok {
		var slice []interface{}
		slice = append(slice, file[len(projectPath)+1:], line)
		slice = append(slice, data...)
		logger.Println(slice...)
	}
}
