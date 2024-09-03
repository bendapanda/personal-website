package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogging(loggingFile string) error {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	defaultLogger := "console"
	if loggingFile != defaultLogger {
		file, err := os.OpenFile(loggingFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		log.SetOutput(file)

		return nil
	}
	return nil
}
