package leetcode

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)
		if string(bytes) == `{"query":"query userContestRankingInfo($username: String!) {userContestRanking(username: $username) {rating badge {name}}}","variables":{"username":"awice"}}` {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w,
				`
				{
					"data": {
						"userContestRanking": {
							"rating": 2944.761,
							"badge": {
								"name": "Guardian"
							}
						}
					}
				}
				`,
			)
		}
	}))
	defer ts.Close()
	leetcodeApiImpl := NewLeetcodeApiImpl(ts.URL)
	userInfo, _ := leetcodeApiImpl.GetUserInfo("awice")
	assert.Equal(t, "awice", userInfo.Username)
	assert.Equal(t, 2944, userInfo.Rating)
	assert.Equal(t, "Guardian", userInfo.Rank)
}
