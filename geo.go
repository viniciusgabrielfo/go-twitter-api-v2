package twitter

type Geo struct {
	PlaceID     string
	Type        string    `json:"type"`
	Bbox        []float64 `json:"bbox"`
	Coordinates struct {
		Type        string
		Coordinates []float32
	}
	Properties interface{} `json:"properties"`
}
