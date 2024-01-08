package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitializeLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(file, "", log.LstdFlags)
}

func Log(message string) {
	if Logger != nil {
		Logger.Println(message)
	}
}

func Error(message string, err error) {
	if Logger != nil {
		Logger.Println(message, err)
	}
}
