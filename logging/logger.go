package logging

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
)

func InitLoggers() {
	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
