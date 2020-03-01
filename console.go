package customLog

import "fmt"

func printLog(le logLevel, msg string) {
	funcName, fileName, line := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, msg)
}

func (lg *Log) Debug(msg string) {
	if lg.level <= DEBUG {
		printLog(DEBUG, msg)
	}
}
func (lg *Log) Info(msg string) {
	if lg.level <= INFO {
		printLog(INFO, msg)
	}
}
func (lg *Log) Waring(msg string) {
	if lg.level <= WARING {
		printLog(WARING, msg)
	}
}
func (lg *Log) Error(msg string) {
	if lg.level <= ERROR {
		printLog(ERROR, msg)
	}
}
func (lg *Log) Fata(msg string) {
	if lg.level <= FATAL {
		printLog(FATAL, msg)
	}
}
