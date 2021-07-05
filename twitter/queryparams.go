package twitter

import (
	"errors"
	"go-twitter-api-v2/twitter/queryparams"
)

func GetTweetFieldsStringQuery(tweetFields []string) (string, error) {

	if len(tweetFields) <= 0 {
		return "", errors.New("no tweet fields identified com 'tweetFields []string' param.")
	}

	queryTweetFields := queryparams.TweetsFieldsParamKey + "="
	for _, field := range tweetFields {
		queryTweetFields += field

		if field != tweetFields[len(tweetFields)-1] {
			queryTweetFields += ","
		}
	}

	return queryTweetFields, nil
}

func GetQueryParamsString(tweetFields, expansions, mediaFields, placeFields, pollFields, userFields string) string {
	queryParams := ""

	if len(tweetFields) > 0 {
		queryParams += tweetFields
	}

	if len(expansions) > 0 {
		queryParams += expansions
	}

	if len(mediaFields) > 0 {
		queryParams += mediaFields
	}

	if len(placeFields) > 0 {
		queryParams += placeFields
	}

	if len(pollFields) > 0 {
		queryParams += placeFields
	}

	if len(userFields) > 0 {
		queryParams += placeFields
	}

	return queryParams
}
