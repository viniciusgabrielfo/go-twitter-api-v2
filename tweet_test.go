package twitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleTweet(t *testing.T) {
	twitterClient := NewClient(BEARER_TOKEN)
	tweet, err := twitterClient.GetSingleTweet("1400935254244966406")

	expectedTweet := &Tweet{
		ID:   "1400935254244966406",
		Text: "Publiquei meu primeiro artigo no Medium, quem puder dar aquela olhada e deixar aquele feedback humilde, agradeço demais: \n\n\"Otimizando tamanho em memória de Structs em Golang\" \nhttps://t.co/9dWmVzWbyi",
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedTweet, tweet)

}

func TestGetMultipleTweets(t *testing.T) {
	twitterClient := NewClient(BEARER_TOKEN)
	tweets, err := twitterClient.GetMultipleTweets([]string{"1400935254244966406", "1293595870563381249"})

	expectedTweets := &[]Tweet{
		{
			ID:   "1400935254244966406",
			Text: "Publiquei meu primeiro artigo no Medium, quem puder dar aquela olhada e deixar aquele feedback humilde, agradeço demais: \n\n\"Otimizando tamanho em memória de Structs em Golang\" \nhttps://t.co/9dWmVzWbyi",
		},
		{
			ID:   "1293595870563381249",
			Text: "Twitter API v2: Early Access released\n\nToday we announced Early Access to the first endpoints of the new Twitter API!\n\n#TwitterAPI #EarlyAccess #VersionBump https://t.co/g7v3aeIbtQ",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedTweets, tweets)

}
