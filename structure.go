package log

import (
	"bytes"
	"fmt"
)

type Structure map[string]interface{}

func (s Structure) String() string {
	buffer := &bytes.Buffer{}
	if len(s) != 0 {
		first := true
		buffer.WriteString("[")
		for key, value := range s {
			if !first {
				buffer.WriteString(";")
			}
			buffer.WriteString(key)
			buffer.WriteString(":")
			fmt.Fprint(buffer, value)
			first = false
		}
		buffer.WriteString("]")
	}
	return buffer.String()
}
