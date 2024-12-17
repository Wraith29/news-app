package internal

import (
	"log"
	"os"
)

var _logger *log.Logger = nil

func GetLogger() *log.Logger {
	if _logger != nil {
		return _logger
	}

	logger := log.Default()
	logger.SetOutput(os.Stdout)

	_logger = logger

	return _logger
}
