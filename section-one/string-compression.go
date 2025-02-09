package section_one

import (
	"strconv"
	"strings"
)

func StringCompress(value string) string {
	valueSlice := []rune(value)
	if len(valueSlice) == 1 {
		return value
	}
	state, statenum := ' ', 0
	var output strings.Builder
	for k, v := range valueSlice {
		if v != state {
			if state != ' ' {
				output.WriteRune([]rune(strconv.Itoa(statenum))[0])
				statenum = 0
			}
			output.WriteRune(v)
			state = v
			statenum++
		} else if v == state {
			statenum++
			if len(valueSlice)-1 == k {
				output.WriteRune([]rune(strconv.Itoa(statenum))[0])
			}
		}
	}
	return output.String()
}
