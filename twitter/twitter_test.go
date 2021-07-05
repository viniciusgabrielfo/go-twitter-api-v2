package twitter_test

import (
	"go-twitter-api-v2/twitter"
	"os"
)

const twitterAPIPath = "https://api.twitter.com/2"

var BEARER_TOKEN = os.Getenv("TWITTER_BEARER_TOKEN")

func NewTwitterTestApi(bearerToken string) *twitter.Client {
	return &twitter.Client{
		Tweet:  twitter.NewTweetService(),
		User:   twitter.NewUserService(),
		Stream: twitter.NewStreamService(),

		BasePath:    twitterAPIPath,
		BearerToken: "bearer " + bearerToken,
	}
}
