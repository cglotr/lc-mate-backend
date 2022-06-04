package leetcode

//go:generate mockgen -package=leetcode -mock_names=LeetcodeApi=LeetcodeApiMock -source=./leetcode_api.go -destination=./leetcode_api_mock.go
type LeetcodeApi interface {
	GetUserInfo(username string) (*UserInfo, error)
}
