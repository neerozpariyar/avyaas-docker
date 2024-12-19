package utils

import (
	"database/sql"
	"regexp"
)

/*
Contains checks if a string value is in the given slice of string data.

It takes in two parameters:
  - elements: a slice of string data to check the value from
  - value: a string data to check
*/
func Contains(elements []string, value string) bool {
	for _, element := range elements {
		if value == element {
			return true
		}
	}
	return false
}

func ContainsInt(elements []int, value int) bool {
	for _, element := range elements {
		if value == element {
			return true
		}
	}
	return false
}

func ContainsUint(elements []uint, value uint) bool {
	for _, element := range elements {
		if value == element {
			return true
		}
	}
	return false
}

func ContainsNullSQL(elements []sql.NullInt32, value sql.NullInt32) bool {
	for _, element := range elements {
		if value == element {
			return true
		}
	}
	return false
}

func ContainsOnlyNumber(input string) bool {
	if ok := regexp.MustCompile(`^[0-9]+$`).MatchString(input); ok {
		return true
	} else {
		return false
	}
}
