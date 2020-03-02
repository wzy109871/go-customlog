package customLog

import "fmt"

func printLog(le logLevel, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	funcName, fileName, line := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, str)
}

func (lg *Log) Debug(msg string, a ...interface{}) {
	if lg.level <= DEBUG {
		printLog(DEBUG, msg, a...)
	}
}
func (lg *Log) Info(msg string, a ...interface{}) {
	if lg.level <= INFO {
		printLog(INFO, msg, a...)
	}
}
func (lg *Log) Waring(msg string, a ...interface{}) {
	if lg.level <= WARING {
		printLog(WARING, msg, a...)
	}
}
func (lg *Log) Error(msg string, a ...interface{}) {
	if lg.level <= ERROR {
		printLog(ERROR, msg, a...)
	}
}
func (lg *Log) Fata(msg string, a ...interface{}) {
	if lg.level <= FATAL {
		printLog(FATAL, msg, a...)
	}
}
