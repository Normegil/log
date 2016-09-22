package log

import (
	"regexp"
	"strconv"
	"testing"
)

func TestStructureToString_Empty(t *testing.T) {
	structure := Structure{}

	toTest := structure.String()
	if "" != toTest {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", "", toTest)
	}
}

func TestStructure_With(t *testing.T) {
	structure := Structure{}

	key := "Test"
	value := "test"
	structure = structure.With(Structure{key: value})
	if value != structure[key] {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", value, structure[key])
	}
}

func TestStructureToString(t *testing.T) {
	structure := Structure{
		"String": "test",
		"bool":   true,
		"int":    2,
		"struct": struct {
			string
		}{"Test"},
	}

	key := `[a-zA-Z0-9]*`
	value := `[a-zA-Z0-9\{\}]*`
	keyValue := key + `:` + value
	extraProperties := `(;` + keyValue + `){` + strconv.Itoa(len(structure)-1) + `}`
	regex := ` \[` + keyValue + extraProperties + `\]$`

	toTest := structure.String()
	if regexp.MustCompile(regex).MatchString(toTest) {
		t.Errorf("Error (Mismatched strings) [Expected: '%s'; Received: '%s']", regex, toTest)
	}
}
