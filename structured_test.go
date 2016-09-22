package log

import (
	"bytes"
	"testing"

	"strings"

	"fmt"

	"github.com/Sirupsen/logrus"
)

func TestStructuredLog(t *testing.T) {
	cases := []struct {
		Level     Level
		Structure Structure
		Message   string
	}{
		{Level: DEBUG, Message: "Message", Structure: Structure{}},
		{Level: INFO, Message: "Message", Structure: Structure{}},
		{Level: DEBUG, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
		{Level: INFO, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
	}

	for _, test := range cases {
		logger := logrus.New()
		buffer := &bytes.Buffer{}
		logger.Out = buffer
		logger.Level = logrus.DebugLevel
		logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

		structured := StructuredLog{Logger: logger}

		structured.Log(test.Level, test.Structure, test.Message)
		checkLogOutput(t, test.Level, test.Message, test.Structure, buffer.String())
	}
}

func TestStructuredLog_Panic(t *testing.T) {
	cases := []struct {
		Level     Level
		Structure Structure
		Message   string
	}{
		{Level: PANIC, Message: "Message", Structure: Structure{}},
		{Level: PANIC, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
		{Level: INFO, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
	}

	for _, test := range cases {
		logger := logrus.New()
		buffer := &bytes.Buffer{}
		logger.Out = buffer
		logger.Level = logrus.DebugLevel
		logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

		structured := StructuredLog{Logger: logger}

		func() {
			defer func() {
				recover()
				checkLogOutput(t, test.Level, test.Message, test.Structure, buffer.String())
			}()
			structured.Log(test.Level, test.Structure, test.Message)
		}()
	}
}

func TestStructuredLog_NoLogger(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer

	basic := StructuredLog{}

	basic.Log(PANIC, Structure{}, "Message")
	logMsg := buffer.String()
	if 0 != len(logMsg) {
		t.Errorf("Error (String not empty) [Received: '%s']", logMsg)
	}
}

func checkLogOutput(t *testing.T, lvl Level, msg string, str Structure, output string) {
	expect := "level=" + string(lvl)
	if strings.Contains(output, expect) {
		t.Errorf("Error (Doesn't contains substring) [Expected: '%s'; Received: '%s']", expect, output)
	}

	expect = "msg=" + msg + "\n"
	if strings.Contains(output, expect) {
		t.Errorf("Error (Doesn't contains substring) [Expected: '%s'; Received: '%s']", expect, output)
	}

	if 0 == len(str) {
		for key, value := range str {
			expect = key + "=" + fmt.Sprint(value)
			if strings.Contains(output, expect) {
				t.Errorf("Error (Doesn't contains substring) [Expected: '%s'; Received: '%s']", expect, output)
			}
		}
	}
}

func TestStructuredLog_StructureKept(t *testing.T) {
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.Formatter = &logrus.TextFormatter{DisableColors: true, DisableTimestamp: true}

	structured := StructuredLog{Logger: logger.WithField("Lib", "logrus")}
	msg := "Message"
	structure := Structure{}

	structured.Log(INFO, structure, msg)
	logMsg := buffer.String()

	if !strings.Contains(logMsg, "Lib=logrus") {
		t.Errorf("Error (substring not found) [Expected: '%s'; Received: '%s']", "Lib=logrus", logMsg)
	}
}
