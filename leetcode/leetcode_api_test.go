package leetcode_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	ts := CreateTestServer()
	defer ts.Close()
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl(ts.URL)
	userInfo, _ := leetcodeApiImpl.GetUserInfo("awice")
	assert.Equal(t, "awice", userInfo.Username)
	assert.Equal(t, 2944, userInfo.Rating)
	assert.Equal(t, "Guardian", userInfo.Rank)
}

func TestServerError(t *testing.T) {
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl("")
	_, err := leetcodeApiImpl.GetUserInfo("awice")
	assert.NotNil(t, err)
}

func TestUserNotFound(t *testing.T) {
	ts := CreateTestServer()
	defer ts.Close()
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl(ts.URL)
	_, err := leetcodeApiImpl.GetUserInfo("")
	assert.Equal(t, "User matching query does not exist.", err.Error())
}

func TestInvalidJson(t *testing.T) {
	ts := CreateTestServer()
	defer ts.Close()
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl(ts.URL)
	_, err := leetcodeApiImpl.GetUserInfo("!")
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestNullContestRanking(t *testing.T) {
	ts := CreateTestServer()
	defer ts.Close()
	leetcodeApiImpl := leetcode.NewLeetcodeApiImpl(ts.URL)
	userInfo, _ := leetcodeApiImpl.GetUserInfo("fabrizio3")
	assert.Equal(t, "fabrizio3", userInfo.Username)
	assert.Equal(t, 0, userInfo.Rating)
	assert.Equal(t, "", userInfo.Rank)
}

func CreateTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)
		switch string(bytes) {
		case `{"query":"query userContestRankingInfo($username: String!) {userContestRanking(username: $username) {rating badge {name}}}","variables":{"username":"awice"}}`:
			{
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
				break
			}
		case `{"query":"query userContestRankingInfo($username: String!) {userContestRanking(username: $username) {rating badge {name}}}","variables":{"username":"!"}}`:
			{
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				io.WriteString(w, "")
				break
			}
		case `{"query":"query userContestRankingInfo($username: String!) {userContestRanking(username: $username) {rating badge {name}}}","variables":{"username":"fabrizio3"}}`:
			{
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				io.WriteString(w,
					`
					{
						"data": {
							"userContestRanking": null
						}
					}
					`,
				)
				break
			}
		default:
			{
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				io.WriteString(w,
					`
					{
						"errors": [
							{
								"message": "User matching query does not exist.",
								"locations": [
									{
										"line": 1,
										"column": 51
									}
								],
								"path": [
									"userContestRanking"
								]
							}
						],
						"data": {
							"userContestRanking": null
						}
					}
					`,
				)
				break
			}
		}
	}))
	return ts
}
