package file

import "strings"

func Slugify(fileName string) string {
	return strings.Replace(fileName, " ", "-", -1)
}
