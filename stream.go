package twitter

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const streamServicePath = "/tweets/search/stream"

type FilteredStreamTweet struct {
	Data          Tweet                `json:"data"`
	MatchingRules []FilteredStreamRule `json:"matching_rules"`
}

type FilteredStreamRule struct {
	ID  int
	Tag string
}

type streamService struct {
	path string
}

func NewStreamService() *streamService {
	return &streamService{
		path: streamServicePath,
	}
}

func (c *Client) GetFilteredStreamTweets(tweetsChan chan<- FilteredStreamTweet) {

	// TODO: ass backfill_minutes param to request
	req, err := http.NewRequest(http.MethodGet, c.Stream.path, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", c.BearerToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("filtered stream tweets request failed: ", err)
	}

	defer resp.Body.Close()

	log.Println("Filtered stream tweets started.")
	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("reading body failed: ", err)
		}

		line = bytes.TrimSpace(line)
		log.Println("identified new tweet from filtered stream: ", string(line))

		var tweetFiltered FilteredStreamTweet
		json.Unmarshal(line, &tweetFiltered)
		tweetsChan <- tweetFiltered
	}
}
