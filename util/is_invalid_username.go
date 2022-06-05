package util

var invalidUsernames = map[string]bool{
	"support":   true,
	"jobs":      true,
	"bugbounty": true,
	"student":   true,
	"terms":     true,
	"privacy":   true,
	"region":    true,
	"explore":   true,
	"contest":   true,
	"discuss":   true,
	"interview": true,
	"store":     true,
	"profile":   true,
}

func IsInvalidUsername(username string) bool {
	_, in := invalidUsernames[username]
	return in
}
