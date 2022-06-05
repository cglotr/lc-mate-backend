package leetcode

import "fmt"

type UserInfo struct {
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	Rank     string `json:"rank"`
}

func (u *UserInfo) String() string {
	return fmt.Sprintf(`%v, %v, %v`, u.Username, u.Rank, u.Rating)
}
