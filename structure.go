package log

import (
	"fmt"
	"strings"
)

// Structure define a map of "key:value" to log
type Structure map[string]interface{}

// With merge the given Structure with the current Structure and send back the result.
func (s Structure) With(str Structure) Structure {
	if nil == s {
		s = make(map[string]interface{})
	}

	toReturn := s
	for key, value := range str {
		toReturn[key] = value
	}
	return toReturn
}

// String representation of a Structure
func (s Structure) String() string {
	var toJoin []string
	if len(s) != 0 {
		first := true
		toJoin = append(toJoin, "[")
		for key, value := range s {
			if !first {
				toJoin = append(toJoin, ";")
			}
			toJoin = append(toJoin, key)
			toJoin = append(toJoin, ":")
			toJoin = append(toJoin, fmt.Sprint(value))
			first = false
		}
		toJoin = append(toJoin, "]")
	}
	return strings.Join(toJoin, "")
}
