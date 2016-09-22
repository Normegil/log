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
	lvl := DEBUG
	basic.Log(lvl, Structure{}, msg)

	logMsg := buffer.String()
	if "" != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", "", logMsg)
	}
}

func TestBasicLog(t *testing.T) {
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
		logger := log.Logger{}
		buffer := &bytes.Buffer{}
		logger.SetOutput(buffer)

		basic := BasicLog{Logger: &logger, PrintDebug: true}
		regex := loadRegex(test.Level, test.Message, test.Structure)

		basic.Log(test.Level, test.Structure, test.Message)
		logMsg := buffer.String()
		if !regex.MatchString(logMsg) {
			t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", regex, logMsg)
		}
	}
}

func TestBasicLog_Panic(t *testing.T) {
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
	}

	for _, test := range cases {
		logger := log.Logger{}
		buffer := &bytes.Buffer{}
		logger.SetOutput(buffer)

		basic := BasicLog{Logger: &logger, PrintDebug: true}
		regex := loadRegex(test.Level, test.Message, test.Structure)

		func() {
			defer func() {
				recover()
				logMsg := buffer.String()
				if !regex.MatchString(logMsg) {
					t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", regex, logMsg)
				}
			}()
			basic.Log(test.Level, test.Structure, test.Message)
		}()
	}
}

func loadRegex(lvl Level, msg string, str Structure) *regexp.Regexp {
	regex := `^\[` + strings.ToUpper(string(lvl)) + `\]` + msg
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
