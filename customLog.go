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


