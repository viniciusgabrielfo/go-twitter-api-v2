package twitter_test

import (
	"go-twitter-api-v2/twitter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleUserByID(t *testing.T) {
	twitterClient := twitter.NewTwitterApi(BEARER_TOKEN)
	user, err := twitterClient.GetSingleUserByID("1592725891")

	expectedUser := &twitter.User{
		ID:       "1592725891",
		Name:     "elfo.go",
		Username: "velfo",
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetMultipleUsersByID(t *testing.T) {
	twitterClient := twitter.NewTwitterApi(BEARER_TOKEN)
	users, err := twitterClient.GetMultipleUsersByID([]string{"1592725891", "783214"})

	expectedUsers := &[]twitter.User{
		{
			ID:       "1592725891",
			Name:     "elfo.go",
			Username: "velfo",
		},
		{
			ID:       "783214",
			Name:     "Twitter",
			Username: "Twitter",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestGetSingleUserByUsername(t *testing.T) {
	twitterClient := twitter.NewTwitterApi(BEARER_TOKEN)
	user, err := twitterClient.GetSingleUserByUsername("velfo")

	expectedUser := &twitter.User{
		ID:       "1592725891",
		Name:     "elfo.go",
		Username: "velfo",
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetMultipleUsersByUsername(t *testing.T) {
	twitterClient := twitter.NewTwitterApi(BEARER_TOKEN)
	users, err := twitterClient.GetMultipleUsersByUsername([]string{"velfo", "Twitter"})

	expectedUsers := &[]twitter.User{
		{
			ID:       "1592725891",
			Name:     "elfo.go",
			Username: "velfo",
		},
		{
			ID:       "783214",
			Name:     "Twitter",
			Username: "Twitter",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, users)
}
