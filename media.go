package twitter

type Media struct {
	MediaKey             string           `json:"media_key"`
	Type                 string           `json:"type"`
	DurationMilliseconds int              `json:"duration_ms"`
	Height               int              `json:"height"`
	Width                int              `json:"width"`
	NonPublicMetrics     NonPublicMetrics `json:"non_public_metrics"`
	OrganicMetrics       OrganicMetrics   `json:"organic_metrics"`
	PromotedMetrics      PromotedMetrics  `json:"promoted_metrics"`
	PublicMetrics        PublicMetrics    `json:"public_metrics"`
	PreviewImageURL      string           `json:"preview_image_url"`
}
