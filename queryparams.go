package twitter

import (
	"net/http"
)

const (
	userFieldsKey  = "user.fields"
	expansionsKey  = "expansions"
	tweetFieldsKey = "tweet.fields"
)

type QueryParamList []string

func (q *QueryParamList) getString() string {
	var queryParamRawQuery string

	if len(*q) == 0 {
		return queryParamRawQuery
	}

	for i, param := range *q {
		queryParamRawQuery += param

		if i != len(*q)-1 {
			queryParamRawQuery += ","
		}
	}

	return queryParamRawQuery
}

type QueryParams struct {
	UserFields  QueryParamList
	Expansions  QueryParamList
	TweetFields QueryParamList
}

func (c *Client) getQueryParamsRawQuery(req *http.Request, params QueryParams) {
	q := req.URL.Query()

	if userFieldsStr := params.UserFields.getString(); userFieldsStr != "" {
		q.Add(userFieldsKey, userFieldsStr)
	}

	if expansionsStr := params.Expansions.getString(); expansionsStr != "" {
		q.Add(expansionsKey, expansionsStr)
	}

	if userTweetsStr := params.TweetFields.getString(); userTweetsStr != "" {
		q.Add(tweetFieldsKey, userTweetsStr)
	}

	req.URL.RawQuery = q.Encode()
}
