package util

import "strings"

func GetUsernamesQuery(usernames []string) string {
	enclosed := []string{}
	for _, username := range usernames {
		enclosed = append(enclosed, `"`+username+`"`)
	}
	return strings.Join(enclosed, ",")
}
