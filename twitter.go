package twitter

const twitterAPIPath = "https://api.twitter.com/2"

type Client struct {
	Tweet    *tweetService
	User     *userService
	Stream   *streamService
	BasePath string

	BearerToken string
}

func NewClient(bearerToken string) *Client {
	return &Client{
		Tweet:  NewTweetService(),
		User:   NewUserService(),
		Stream: NewStreamService(),

		BasePath:    twitterAPIPath,
		BearerToken: "bearer " + bearerToken,
	}
}
