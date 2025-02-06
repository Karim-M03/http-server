package logger

import (
	"io"
	"log"
	"os"
)

var InfoLogger *log.Logger
var WarnLogger *log.Logger
var ErrorLogger *log.Logger

// init is called automatically and sets up the loggers to output
// to both the console and the server.log file.
func init() {
	// Open (or create) the log file for appending
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// If the file can’t be opened, log the error and exit.
		log.Fatalf("Failed to open server.log: %v", err)
	}

	// Create a multiwriter so that logs are sent to both stdout and the file.
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Initialize the loggers with their prefixes and standard flags.
	InfoLogger = log.New(multiWriter, "INFO: ", log.LstdFlags)
	WarnLogger = log.New(multiWriter, "WARN: ", log.LstdFlags)
	ErrorLogger = log.New(multiWriter, "ERROR: ", log.LstdFlags)
}

