package log

import "testing"

// Test if loggers implements the AgnosticLogger interface

func TestBasicLog_AgnosticInterface(t *testing.T) {
	var log AgnosticLogger
	log = BasicLog{}
	log.Log(DEBUG, Structure{}, "Test")
}

func TestStructuredLog_AgnosticInterface(t *testing.T) {
	var log AgnosticLogger
	log = StructuredLog{}
	log.Log(DEBUG, Structure{}, "Test")
}
