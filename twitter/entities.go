package twitter

type Annotation struct {
	Start          uint16
	End            uint16
	Probability    uint8
	Type           string
	NormalizedText string
}

type Cashtag struct {
	Start uint16
	End   uint16
	Tag   string
}

type Hashtag struct {
	Start uint16
	End   uint16
	Tag   string
}

type Mention struct {
	Start uint16
	End   uint16
	Tag   string
}

type Url struct {
	Start       int
	End         int
	URL         string
	ExpandedURL string
	DisplayURL  string
	// Images      []image
	Status      int
	Title       string
	Description string
	UnwoundURL  string
}
