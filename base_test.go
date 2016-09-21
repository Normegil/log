package log

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestBasicLog_NoStructure_DebugDisabled(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger}
	msg := "Message"
	basic.Debug(Structure{}, msg)

	logMsg := buffer.String()
	if "" != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", "", logMsg)
	}
}

func TestBasicLog_NoStructure_DebugEnabled(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, PrintDebug: true}
	msg := "Message"
	basic.Debug(Structure{}, msg)

	logMsg := buffer.String()
	expect := "[DEBUG]Message\n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestBasicLog_NoStructure_Info(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger}
	msg := "Message"
	basic.Info(Structure{}, msg)

	logMsg := buffer.String()
	expect := "[INFO]Message\n"
	if expect != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestBasicLog_NoStructure_Panic(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger}
	msg := "Message"
	defer func() {
		recover()
		logMsg := buffer.String()
		expect := "[PANIC]Message\n"
		if expect != logMsg {
			t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
		}
	}()

	basic.Panic(Structure{}, msg)
}

func TestBasicLog_Structure_Debug(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, PrintDebug: true}
	msg := "Message"
	structure := Structure{"Test": "test", "Test1": 1, "Test2": struct{ test int }{2}}
	basic.Debug(structure, msg)

	expect := loadRegex("Debug", msg, structure)
	logMsg := buffer.String()
	if !expect.MatchString(logMsg) {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestBasicLog_Structure_Info(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, PrintDebug: true}
	msg := "Message"
	structure := Structure{"Test": "test", "Test1": 1, "Test2": struct{ test int }{2}}
	basic.Info(structure, msg)

	expect := loadRegex("Info", msg, structure)
	logMsg := buffer.String()
	if !expect.MatchString(logMsg) {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestBasicLog_Structure_Panic(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, PrintDebug: true}

	msg := "Message"
	structure := Structure{"Test": "test", "Test1": 1, "Test2": struct{ test int }{2}}
	defer func() {
		recover()
		logMsg := buffer.String()
		expect := loadRegex("Panic", msg, structure)
		if !expect.MatchString(logMsg) {
			t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", expect, logMsg)
		}
	}()
	basic.Panic(structure, msg)
}

func loadRegex(lvl string, msg string, str Structure) *regexp.Regexp {
	regex := `^\[` + strings.ToUpper(lvl) + `\]` + msg
	size := len(str)
	if 0 != size {
		key := `[a-zA-Z0-9]*`
		value := `[a-zA-Z0-9\{\}]*`
		keyValue := key + `:` + value
		extraProperties := `(;` + keyValue + `){` + strconv.Itoa(size-1) + `}`
		regex += ` \[` + keyValue + extraProperties + `\]`
	}
	regex += `\n$`
	return regexp.MustCompile(regex)
}
