package logging

import (
	"fmt"
	"io"
	"news-api/internal/config"
	"os"
	"time"
)

type level int

const (
	info level = iota
	err
)

// Will be made better if more levels are added
func levelFromString(val string) level {
	if val == "info" {
		return info
	}

	return err
}

func (l level) String() string {
	switch l {
	case info:
		return "info"
	case err:
		return "err"
	}

	return ""
}

type logger struct {
	output io.Writer
	lvl    level
}

func newLogger(writer io.Writer, lvl level) *logger {
	return &logger{
		output: writer,
		lvl:    lvl,
	}
}

var _inst *logger = nil

func GetLogger() *logger {
	if _inst != nil {
		return _inst
	}

	lvl := levelFromString(config.Cfg.Logging.Level)
	_inst = newLogger(os.Stdout, lvl)

	return _inst
}

const timestampFmt = "2006/01/02 15:04:05"

func (l *logger) log(lvl level, msg string) {
	now := time.Now()

	log := fmt.Sprintf("%s %s %s\n", lvl.String(), now.Format(timestampFmt), msg)

	_, err := l.output.Write([]byte(log))

	if err != nil {
		panic(err)
	}
}

func (l *logger) Info(msg string) {
	// level > info means do not include
	if l.lvl > info {
		return
	}

	l.log(info, msg)
}

func (l *logger) Err(msg string) {
	if l.lvl > err {
		return
	}

	l.log(err, msg)
}
