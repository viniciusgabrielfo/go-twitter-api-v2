package twitter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const tweetServicePath = "/tweets"

type singleTweetResponse struct {
	Data struct {
		Tweet
	}
}

type MultipleTweetsResponse struct {
	Data []Tweet
}

type Tweet struct {
	ID                 string             `json:"id"`
	Text               string             `json:"text"`
	Attachments        attachments        `json:"attachments"`
	AuthorID           string             `json:"author_id"`
	ContextAnnotations ContextAnnotations `json:"context_annotations"`
	ConversationID     string             `json:"conversation_id"`
	CreatedAt          time.Time          `json:"created_at"`
	Entities           struct {
		Annotations []Annotation `json:"annotations"`
		Hashtags    []Hashtag    `json:"hashtags"`
		Mentions    []Mention    `json:"mentions"`
		Cashtags    []Cashtag    `json:"cashtags"`
		URLs        []Url        `json:"urls"`
	}
	Geo               Geo               `json:"geo"`
	InReplyToUserID   string            `json:"in_reply_to_user_id"`
	Lang              string            `json:"lang"`
	NonPublicMetrics  NonPublicMetrics  `json:"non_public_metrics"`
	OrganicMetrics    OrganicMetrics    `json:"organic_metrics"`
	PossiblySensitive bool              `json:"possibly_metrics"`
	PromotedMetrics   PromotedMetrics   `json:"prometed_metrics"`
	PublicMetrics     PublicMetrics     `json:"public_metrics"`
	ReferencedTweets  []ReferencedTweet `json:"referenced_tweets"`
	ReplySettings     string            `json:"reply_settings"`
	Source            string            `json:"source"`
	// Withheld
}

type ContextAnnotations struct {
	Domain struct {
		ID          string
		Name        string
		Description string
		Entity      struct {
			ID   string
			Name string
		}
	}
}

type ReferencedTweet struct {
	ID   string
	Type string
}

type image struct {
	URL    string
	Width  int
	Height int
}

type attachments struct {
	PollIDs   []string
	MediaKeys []string
}

type tweetService struct {
	path string
}

func NewTweetService() *tweetService {
	return &tweetService{
		path: tweetServicePath,
	}
}

func (c *Client) GetSingleTweet(tweetID string) (*Tweet, error) {
	var singleTweetResponse singleTweetResponse

	if tweetID == "" {
		return &singleTweetResponse.Data.Tweet, errors.New("no tweet id identified on tweetID params in GetSingleTweet().")
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.Tweet.path+"/"+tweetID, nil)
	req.Header.Add("Authorization", c.BearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("request single tweet failed: ", err)
		return &singleTweetResponse.Data.Tweet, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("reading body single tweet failed: ", err)
		return &singleTweetResponse.Data.Tweet, err
	}

	json.Unmarshal(body, &singleTweetResponse)
	return &singleTweetResponse.Data.Tweet, nil
}

func (c *Client) GetMultipleTweets(tweetIDs []string) (*[]Tweet, error) {
	var tweets []Tweet

	if len(tweetIDs) <= 0 {
		return &tweets, errors.New("no tweet ids identified on tweetIDs params in GetMultipleTweets()")
	}

	var tweetIDsQuery = "?ids="
	for _, id := range tweetIDs {
		tweetIDsQuery += id

		if id != tweetIDs[len(tweetIDs)-1] {
			tweetIDsQuery += ","
		}
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.Tweet.path+tweetIDsQuery, nil)
	req.Header.Add("Authorization", c.BearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("request multiple tweets failed: ", err)
		return &tweets, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("reading body multiple tweets failed: ", err)
		return &tweets, err
	}

	var multipleTweetsResponse MultipleTweetsResponse
	json.Unmarshal(body, &multipleTweetsResponse)

	for _, tweet := range multipleTweetsResponse.Data {
		tweets = append(tweets, tweet)
	}
	return &tweets, nil
}
