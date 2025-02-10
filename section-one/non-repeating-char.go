package section_one

func FirstNonRepeating(value string) string {
	var check int
	var charOccur map[string]int = make(map[string]int)
	for i := 0; i < len(value); i++ {
		check = 0
		for k := i + 1; k < len(value) && check != 1; k++ {
			if value[i] == value[k] {
				check = 1
				charOccur[string(value[i])] = check
			}
		}
		if charOccur[string(value[i])] == 0 {
			return string(value[i])
		}
	}
	return "-1"
}
