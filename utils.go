package customLog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

func getNowtime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

func parseLoglevel(s string) (logLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("您输入的日志级别有误")
		return UKNOWN, err
	}
}

func unparseLoglevel(le logLevel) string {
	switch le {
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case WARING:
		return "Waring"
	case ERROR:
		return "Error"
	case FATAL:
		return "Fatal"
	default:
		return "Debug"
	}
}

func getInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file)
	return
}
