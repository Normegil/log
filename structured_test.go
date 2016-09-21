package log

import (
	"bytes"
	"testing"

	"strings"

	"github.com/Sirupsen/logrus"
)

func TestStructuredLog_NoStructure_Debug(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"
	structured.Debug(Structure{}, msg)

	logMsg := buffer.String()
	expect := "level=debug msg=" + msg + " \n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestStructuredLog_NoStructure_Info(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"
	structured.Info(Structure{}, msg)

	logMsg := buffer.String()
	expect := "level=info msg=" + msg + " \n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestStructuredLog_NoStructure_Panic(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"

	defer func() {
		recover()
		logMsg := buffer.String()
		expect := "level=panic msg=" + msg + " \n"
		if expect != logMsg {
			t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
		}
	}()
	structured.Panic(Structure{}, msg)
}

func TestStructuredLog_Structure_Debug(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"
	structure := Structure{
		"String": "test",
		"bool":   true,
		"int":    2,
		"struct": struct {
			string
		}{"Test"},
	}
	structured.Debug(structure, msg)

	logMsg := buffer.String()
	expect := "level=debug msg=" + msg + " String=test bool=true int=2 struct={Test} \n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestStructuredLog_Structure_Info(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"
	structure := Structure{
		"String": "test",
		"bool":   true,
		"int":    2,
		"struct": struct {
			string
		}{"Test"},
	}
	structured.Info(structure, msg)

	logMsg := buffer.String()
	expect := "level=info msg=" + msg + " String=test bool=true int=2 struct={Test} \n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestStructuredLog_Structure_Panic(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger}
	msg := "Message"
	structure := Structure{
		"String": "test",
		"bool":   true,
		"int":    2,
		"struct": struct {
			string
		}{"Test"},
	}

	defer func() {
		recover()
		logMsg := buffer.String()
		expect := "level=panic msg=" + msg + " String=test bool=true int=2 struct={Test} \n"
		if expect != logMsg {
			t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
		}
	}()
	structured.Panic(structure, msg)
}

func TestStructuredLog_StructureKept(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger.WithField("Lib", "logrus")}
	msg := "Message"
	structure := Structure{}

	structured.Info(structure, msg)
	logMsg := buffer.String()

	if !strings.Contains(logMsg, "Lib=logrus") {
		t.Errorf("Error (substring not found) [Expected: '%s'; Received: '%s']", "Lib=logrus", logMsg)
	}
}
