package log

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestBasicLog_NoStructure_TraceDisabled(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, Level: DEBUG}
	msg := "Message"
	lvl := TRACE
	basic.Log(lvl, Structure{}, msg)

	logMsg := buffer.String()
	if "" != logMsg {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", "", logMsg)
	}
}

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

type test struct {
	Level     Level
	Structure Structure
	Message   string
}

func getBaseLogMessages() []test {
	return []test{
		{Level: TRACE, Message: "Message", Structure: Structure{}},
		{Level: DEBUG, Message: "Message", Structure: Structure{}},
		{Level: INFO, Message: "Message", Structure: Structure{}},
		{Level: WARN, Message: "Message", Structure: Structure{}},
		{Level: ERROR, Message: "Message", Structure: Structure{}},
		{Level: TRACE, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
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
		{Level: WARN, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
		{Level: ERROR, Message: "Message", Structure: Structure{
			"String": "test",
			"bool":   true,
			"int":    2,
			"struct": struct {
				string
			}{"Test"},
		}},
	}
}

func TestBasicLog(t *testing.T) {
	for _, test := range getBaseLogMessages() {
		logger := log.Logger{}
		buffer := &bytes.Buffer{}
		logger.SetOutput(buffer)

		basic := BasicLog{Logger: &logger, Level: TRACE}
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

		basic := BasicLog{Logger: &logger, Level: DEBUG}
		regex := loadRegex(test.Level, test.Message, test.Structure)

		func() {
			defer func() {
				if err := recover(); nil != err {
					if _, ok := err.(string); !ok {
						t.Error(err)
					}
				}
			}()
			defer func() {
				logMsg := buffer.String()
				if !regex.MatchString(logMsg) {
					t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", regex, logMsg)
				}
			}()
			basic.Log(test.Level, test.Structure, test.Message)
		}()
	}
}

func TestBasicLog_NoLogger(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{}

	basic.Log(PANIC, Structure{}, "Message")
	logMsg := buffer.String()
	if 0 != len(logMsg) {
		t.Errorf("Error (String not empty) [Received: '%s']", logMsg)
	}
}

func TestBasicLog_With(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	basic := BasicLog{Logger: &logger, Level: DEBUG}

	basic.With(Structure{"Test": "test"}).Log(INFO, Structure{}, "Message")
	logMsg := buffer.String()
	expect := "Test:test"
	if !strings.Contains(logMsg, expect) {
		t.Errorf("Error (Doesn't contains substring) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func TestBasicLog_With_Overwrite(t *testing.T) {
	logger := log.Logger{}
	buffer := &bytes.Buffer{}
	logger.SetOutput(buffer)

	var basic AgnosticLogger
	basic = BasicLog{Logger: &logger, Level: DEBUG}

	key := "Test"
	basic = basic.With(Structure{key: "test"})
	basic.With(Structure{key: "otherTest"}).Log(INFO, Structure{}, "Message")
	logMsg := buffer.String()
	expect := key + ":otherTest"
	if !strings.Contains(logMsg, expect) {
		t.Errorf("Error (Doesn't contains substring) [Expected: '%s'; Received: '%s']", expect, logMsg)
	}
}

func loadRegex(lvl Level, msg string, str Structure) *regexp.Regexp {
	regex := `^\[` + strings.ToUpper(lvl.String()) + `\]` + msg
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
