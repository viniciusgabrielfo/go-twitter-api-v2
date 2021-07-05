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
	RetweetCount int
	ReplyCount   int
	LikeCount    int
	QuoteCount   int
}
