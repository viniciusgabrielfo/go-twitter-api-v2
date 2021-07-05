package twitter

type Place struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	ContainedWithin []string `json:"contained_within"`
	Country         string   `json:"contry"`
	CountryCode     string   `json:"country_code"`
	Geo             Geo      `json:"geo"`
	PlaceType       string   `json:"place_type"`
}
