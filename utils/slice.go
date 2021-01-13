package utils

func IndexOf(inputSlice []string, element string) int {
	// TODO: generalize this function
	for i, e := range inputSlice {
		if e == element {
			return i
		}
	}

	return -1
}
