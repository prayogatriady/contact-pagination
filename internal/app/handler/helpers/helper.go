package helpers

import "regexp"

func ValidateSort(sort string) bool {
	pattern := `^[a-zA-Z_]+ (asc|desc)$`

	return regexp.MustCompile(pattern).MatchString(sort)
}