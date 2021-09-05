package twitter

import "time"

type Poll struct {
	ID              string       `json:"id"`
	Options         []PollOption `json:"options"`
	DurationMinutes uint         `json:"duration_minutes"`
	EndDateTime     time.Time    `json:"end_datetime"`
	VotingStatus    string       `json:"voting_status"`
}

type PollOption struct {
	Position uint8  `json:"position"`
	Label    string `json:"label"`
	Votes    uint   `json:"votes"`
}
