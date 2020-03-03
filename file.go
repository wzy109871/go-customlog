package customLog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLog struct {
	level       logLevel
	filepath    string
	filename    string
	fileobj     *os.File
	errfileobj  *os.File
	maxFileSize int64
}

func NewFileLog(levelStr, filepath, filename string, maxSize int64) *FileLog {
	level, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	os.MkdirAll(filepath, 0644)
	var fileLog FileLog
	fileLog.level = level
	fileLog.filepath = filepath
	fileLog.filename = filename
	fileLog.maxFileSize = maxSize
	err = fileLog.initFileobj()
	if err != nil {
		panic(err)
	}
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

func (lg *FileLog) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= lg.maxFileSize
}

func (lg *FileLog) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102-150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(lg.filepath, fileInfo.Name())
	bakLogName := fmt.Sprintf("%s.bak.%s", logName, nowStr)
	file.Close()
	os.Rename(logName, bakLogName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (lg *FileLog) fileprintLog(le logLevel, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	funcName, fileName, line := getInfo(3)
	if lg.checkFileSize(lg.fileobj) {
		newFile, err := lg.splitFile(lg.fileobj)
		if err != nil {
			return
		}
		lg.fileobj = newFile
	}
	fmt.Fprintf(lg.fileobj, "[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, str)
	if le >= ERROR {
		if lg.checkFileSize(lg.errfileobj) {
			newFile, err := lg.splitFile(lg.errfileobj)
			if err != nil {
				return
			}
			lg.errfileobj = newFile
		}
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
