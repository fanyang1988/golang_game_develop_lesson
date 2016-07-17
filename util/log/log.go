package log

import (
	"fmt"
	"github.com/cihub/seelog"
	"runtime/debug"
)

func Init(path string) error {
	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		return err
	}
	seelog.ReplaceLogger(logger)
	return nil
}

func Flush() {
	seelog.Flush()
}

func Trace(format string, params ...interface{}) {
	seelog.Tracef(format, params...)
}

func Debug(format string, params ...interface{}) {
	seelog.Debugf(format, params...)
}

func Info(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

func Error(format string, params ...interface{}) {
	seelog.Errorf(format, params...)
}

func Critical(format string, params ...interface{}) {
	seelog.Critical(
		fmt.Sprintf(format, params...),
		",",
		debug.Stack())
}

func PanicHandle() {
	if err := recover(); err != nil {
		seelog.Critical(
			"Panic Err by ",
			err,
			" , stack: ",
			debug.Stack())
		seelog.Flush()
	}
}
