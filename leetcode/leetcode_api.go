package leetcode

type LeetcodeApi interface {
	GetUserInfo(username string) (*UserInfo, error)
}
