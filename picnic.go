package picnic

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	"strings"
)

type Client struct {
	http     *http.Client
	baseURL  string
	version  string
	country  string
	token    string
	username string
	secret   string
}

type ClientOption func(client *Client)

func WithVersion(version string) ClientOption {
	return func(client *Client) {
		client.baseURL = makeUrl(client.country, version)
		client.version = version
	}
}

func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.token = token
	}
}

func WithBaseUrl(url string) ClientOption {
	return func(client *Client) {
		client.baseURL = url
	}
}

func WithCountry(countryCode string) ClientOption {
	return func(client *Client) {
		client.baseURL = makeUrl(strings.ToLower(countryCode), client.version)
		client.country = countryCode
	}
}

func WithUserName(username string) ClientOption {
	return func(client *Client) {
		client.username = username
	}
}

func WithPassword(password string) ClientOption {
	return func(client *Client) {
		client.secret = md5Hash(password)
	}
}

func md5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func WithHashedPassword(hashedPassword string) ClientOption {
	return func(client *Client) {
		client.secret = hashedPassword
	}
}

func New(client *http.Client, opts ...ClientOption) *Client {
	c := &Client{
		http:    client,
		baseURL: "https://storefront-prod.nl.picnicinternational.com/api/17",
		version: "17",
		country: "nl",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Authenticate() (err error) {
	authenticationUrl := c.baseURL + "/user/login"
	requestBody, jsonErr := json.Marshal(LoginInput{
		Key:      c.username,
		Secret:   c.secret,
		ClientId: 1,
	})
	if jsonErr != nil {
		return jsonErr
	}
	request, requestErr := http.NewRequest("POST", authenticationUrl, bytes.NewReader(requestBody))
	if requestErr != nil {
		return requestErr
	}
	request.Header.Set("Content-Type", "application/json")
	response, responseErr := c.http.Do(request)
	if responseErr != nil {
		return responseErr
	}
	c.token = response.Header.Get("x-picnic-auth")
	return nil
}

func (c *Client) Logout() (err error) {
	logoutUrl := c.baseURL + "/user/login"
	request, requestErr := http.NewRequest("POST", logoutUrl, nil)
	if requestErr != nil {
		return requestErr
	}
	configureHeaders(request, c.token)
	res, httpErr := c.http.Do(request)
	c.post(logoutUrl, nil, nil)
	if httpErr != nil {
		return httpErr
	}
	defer res.Body.Close()
	c.token = ""
	return nil
}

func (c *Client) get(url string, result interface{}) error {
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return requestErr
	}
	configureHeaders(request, c.token)
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return httpErr
	}
	defer res.Body.Close()
	jsonErr := json.NewDecoder(res.Body).Decode(result)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (c *Client) post(url string, body any, result interface{}) error {
	jsonBody, jsonReqErr := json.Marshal(body)
	if jsonReqErr != nil {
		return jsonReqErr
	}
	request, requestErr := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if requestErr != nil {
		return requestErr
	}
	configureHeaders(request, c.token)
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return httpErr
	}
	defer res.Body.Close()
	jsonErr := json.NewDecoder(res.Body).Decode(result)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func (c *Client) GetArticleImage(articleImageId string, size ImageSize) (*image.Image, error) {
	imageUrl := "https://storefront-prod.nl.picnicinternational.com/static/images/" +
		articleImageId +
		"/" +
		size.String() +
		".png"
	res, resErr := http.Get(imageUrl)
	if resErr != nil {
		return nil, resErr
	}
	articleImage, imageErr := png.Decode(res.Body)
	if imageErr != nil {
		return nil, imageErr
	}
	return &articleImage, nil
}

func makeUrl(countryCode string, version string) string {
	return "https://storefront-prod." +
		countryCode +
		".picnicinternational.com/api/" +
		version
}

func configureHeaders(request *http.Request, token string) {
	request.Header.Set("x-picnic-auth", token)
	request.Header.Set("Content-Type", "application/json")
}
