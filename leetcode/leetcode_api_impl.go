package leetcode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://leetcode.com"

type LeetcodeApiImpl struct {
	baseUrl string
}

func NewLeetcodeApiImpl(baseUrl string) *LeetcodeApiImpl {
	return &LeetcodeApiImpl{
		baseUrl: baseUrl,
	}
}

func (l *LeetcodeApiImpl) GetUserInfo(username string) (*UserInfo, error) {
	body, err := json.Marshal(map[string]interface{}{
		"query": "query userContestRankingInfo($username: String!) {userContestRanking(username: $username) {rating badge {name}}}",
		"variables": map[string]interface{}{
			"username": username,
		},
	})
	if err != nil {
		return nil, err
	}
	url := l.baseUrl + "/graphql/"
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	type Badge struct {
		Name string
	}
	type UserContestRanking struct {
		Rating float64
		Badge  Badge
	}
	type Data struct {
		UserContestRanking UserContestRanking
	}
	type JsonUnmarshal struct {
		Data Data
	}
	var jsonUnmarshal JsonUnmarshal
	err = json.Unmarshal(bytes, &jsonUnmarshal)
	if err != nil {
		return nil, err
	}
	userInfo := &UserInfo{
		Username: username,
		Rating:   int(jsonUnmarshal.Data.UserContestRanking.Rating),
		Rank:     jsonUnmarshal.Data.UserContestRanking.Badge.Name,
	}
	return userInfo, nil
}
