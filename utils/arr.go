package utils

func IndexOf[T comparable](vals []T, val T) int {
	for index, v := range vals {
		if v == val {
			return index
		}
	}
	return -1
}
