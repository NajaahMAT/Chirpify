package helper

import "strconv"

func DerefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func DerefStringArray(ptr *[]string) []string {
	if ptr != nil {
		return *ptr
	}
	return nil
}

func DerefInt64(ptr *int64) int64 {
	if ptr != nil {
		return *ptr
	}
	return 0
}

func DerefBool(ptr *bool) bool {
	if ptr != nil {
		return *ptr
	}
	return false
}

// IntToString converts an integer to a string.
func IntToString(value int) string {
	return strconv.Itoa(value)
}
