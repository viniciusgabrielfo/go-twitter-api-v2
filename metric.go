package twitter

type NonPublicMetrics struct {
	ImpressionCount   int
	URLLinkClicks     int
	UserProfileClicks int
}

type OrganicMetrics struct {
	ImpressionCount   int
	LikeCount         int
	ReplyCount        int
	RetweetCount      int
	URLLinkClicks     int
	UserProfileClicks int
}

type PromotedMetrics struct {
	ImpressionCount   int
	LikeCount         int
	ReplyCount        int
	RetweetCount      int
	URLLinkClicks     int
	UserProfileClicks int
}

type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_account"`
	RetweetCount   int
	ReplyCount     int
	LikeCount      int
	QuoteCount     int
}
