package twitter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const userServicePath = "/users"

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Entities    struct {
		URL struct {
			URLs []Url `json:"urls"`
		}
		Description struct {
			URLs     []Url     `json:"urls"`
			Hashtags []Hashtag `json:"hashtags"`
			Mentions []Mention `json:"mentions"`
			Cashtags []Cashtag `json:"cashtags"`
		}
	}
	Location        string        `json:"location"`
	PinnedTweetID   string        `json:"pinned_tweet_id"`
	ProfileImageURL string        `json:"profile_image_url"`
	Protected       bool          `json:"protected"`
	PublicMetrics   PublicMetrics `json:"public_metrics"`
	URL             string        `json:"url"`
	Verified        bool          `json:"verified"`
	// Withheld
}

type singleUserResponse struct {
	Data User `json:"data"`
}

type multipleUsersResponse struct {
	Data []User
}

type userService struct {
	path string
}

func NewUserService() *userService {
	return &userService{
		path: userServicePath,
	}
}

func (c *Client) GetSingleUserByID(userID string, params QueryParams) (*User, error) {
	if userID == "" {
		return nil, errors.New("no userID identified on in GetSingleUserByID()")
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.User.path+"/"+userID, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.BearerToken)

	c.getQueryParamsRawQuery(req, params)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userResponse singleUserResponse
	json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse.Data, nil
}

func (c *Client) GetMultipleUsersByID(usersID []string) (*[]User, error) {
	var users []User

	if len(usersID) <= 0 {
		return &users, errors.New("no usersID identified on in GetMultipleUserByID()")
	}

	var userIDsQuery = "?ids="
	for _, id := range usersID {
		userIDsQuery += id

		if id != usersID[len(usersID)-1] {
			userIDsQuery += ","
		}
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.User.path+userIDsQuery, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.BearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &users, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &users, err
	}

	var userResponse multipleUsersResponse
	json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	users = append(users, userResponse.Data...)

	return &users, nil
}

func (c *Client) GetSingleUserByUsername(username string, params QueryParams) (*User, error) {
	if username == "" {
		return nil, errors.New("no username identified on in GetSingleUserByUsername()")
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.User.path+"/by/username/"+username, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.BearerToken)

	c.getQueryParamsRawQuery(req, params)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userResponse singleUserResponse
	json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	return &userResponse.Data, nil
}

func (c *Client) GetMultipleUsersByUsername(usernames []string) (*[]User, error) {
	var users []User

	if len(usernames) <= 0 {
		return &users, errors.New("no usersID identified on in GetMultipleUserByID()")
	}

	var usernamesQuery = "?usernames="
	for _, username := range usernames {
		usernamesQuery += username

		if username != usernames[len(usernames)-1] {
			usernamesQuery += ","
		}
	}

	req, err := http.NewRequest(http.MethodGet, c.BasePath+c.User.path+"/by"+usernamesQuery, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.BearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &users, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &users, err
	}

	var userResponse multipleUsersResponse
	json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	users = append(users, userResponse.Data...)

	return &users, nil
}
