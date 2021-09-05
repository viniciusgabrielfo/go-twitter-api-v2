package twitter

import (
	"os"
)

var BEARER_TOKEN = os.Getenv("TWITTER_BEARER_TOKEN")

func NewTwitterTestApi(bearerToken string) *Client {
	return &Client{
		Tweet:  NewTweetService(),
		User:   NewUserService(),
		Stream: NewStreamService(),

		BasePath:    twitterAPIPath,
		BearerToken: "bearer " + bearerToken,
	}
}
