package customLog

import (
	"fmt"
	"os"
	"path"
)

type FileLog struct {
	level      logLevel
	filepath   string
	filename   string
	fileobj    *os.File
	errfileobj *os.File
}

func NewFileLog(levelStr, filepath, filename string) *FileLog {
	level, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	os.MkdirAll(filepath, 0644)
	var fileLog FileLog
	fileLog.level = level
	fileLog.filepath = filepath
	fileLog.filename = filename
	fileLog.initFileobj()
	return &fileLog

}

func (lg *FileLog) initFileobj() error {
	fileFullName := path.Join(lg.filepath, lg.filename)
	fileobj, err := os.OpenFile(fileFullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errfileobj, err := os.OpenFile(fileFullName+"error", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	lg.fileobj = fileobj
	lg.errfileobj = errfileobj
	return nil
}

func (lg *FileLog) FileClose() {
	lg.fileobj.Close()
	lg.errfileobj.Close()
}

func (lg *FileLog) fileprintLog(le logLevel, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	funcName, fileName, line := getInfo(3)
	fmt.Fprintf(lg.fileobj, "[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, str)
	if le >= ERROR {
		fmt.Fprintf(lg.errfileobj, "[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, str)
	}

}

func (lg *FileLog) FileDebug(msg string, a ...interface{}) {
	if lg.level <= DEBUG {
		lg.fileprintLog(DEBUG, msg, a...)
	}
}
func (lg *FileLog) FileInfo(msg string, a ...interface{}) {
	if lg.level <= INFO {
		lg.fileprintLog(INFO, msg, a...)
	}
}
func (lg *FileLog) FileWaring(msg string, a ...interface{}) {
	if lg.level <= WARING {
		lg.fileprintLog(WARING, msg, a...)
	}
}
func (lg *FileLog) FileError(msg string, a ...interface{}) {
	if lg.level <= ERROR {
		lg.fileprintLog(ERROR, msg, a...)
	}
}
func (lg *FileLog) FileFata(msg string, a ...interface{}) {
	if lg.level <= FATAL {
		lg.fileprintLog(FATAL, msg, a...)
	}
}
