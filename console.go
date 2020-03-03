package customLog

import "fmt"

type consoleLog struct {
	level logLevel
}

func NewconsoleLog(levelStr string) *consoleLog {
	level, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &consoleLog{
		level: level,
	}
}

func printLog(le logLevel, msg string, a ...interface{}) {
	str := fmt.Sprintf(msg, a...)
	funcName, fileName, line := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", getNowtime(), unparseLoglevel(le), fileName, funcName, line, str)
}

func (lg *consoleLog) ConsoleDebug(msg string, a ...interface{}) {
	if lg.level <= DEBUG {
		printLog(DEBUG, msg, a...)
	}
}
func (lg *consoleLog) ConsoleInfo(msg string, a ...interface{}) {
	if lg.level <= INFO {
		printLog(INFO, msg, a...)
	}
}
func (lg *consoleLog) ConsoleWaring(msg string, a ...interface{}) {
	if lg.level <= WARING {
		printLog(WARING, msg, a...)
	}
}
func (lg *consoleLog) ConsoleError(msg string, a ...interface{}) {
	if lg.level <= ERROR {
		printLog(ERROR, msg, a...)
	}
}
func (lg *consoleLog) ConsoleFata(msg string, a ...interface{}) {
	if lg.level <= FATAL {
		printLog(FATAL, msg, a...)
	}
}
