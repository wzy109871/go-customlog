package customLog

type logLevel uint16

const (
	UKNOWN logLevel = iota
	DEBUG
	INFO
	WARING
	ERROR
	FATAL
)

type Log struct {
	level logLevel
}

func NewLog(levelStr string) *Log {
	level, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &Log{
		level: level,
	}
}
