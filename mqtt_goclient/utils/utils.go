package utils

func ArrayIndex(array interface{}, element interface{}) int {
	var index int
	switch dataType := element.(type) {
		case string:
			index = stringArrayIndex(array.([]string), element.(string))
			_ = dataType
	}
	return index
}

func stringArrayIndex(array []string, element string) int {
	for k,v := range array {
		if v == element {
			return k
		}
	}
	return -1
}
