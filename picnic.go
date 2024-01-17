package picnic

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"strings"
)

// Client the picnic api client it is recommended to use the New method to create a new instance
type Client struct {
	http        *http.Client
	baseURL     string
	version     string
	country     string
	token       string
	parsedToken *token
	username    string
	secret      string
}

const appVersion = "1.15.243-18832"

type ClientOption func(client *Client)

// WithVersion specify the particular version of the picnic API you wish to use.
// The default is set to 17
func WithVersion(version string) ClientOption {
	return func(client *Client) {
		client.baseURL = makeUrl(client.country, version)
		client.version = version
	}
}

// WithToken Provide your own bearer token to be used as the authentication header
// if Authenticate or Logout is called this value with altered.
func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.token = token
	}
}

// WithBaseUrl replaces the full url for API calls. Cannot be used in combination with WithCountry
// And WithVersion
func WithBaseUrl(url string) ClientOption {
	return func(client *Client) {
		client.baseURL = url
	}
}

// WithCountry specify the particular country you wish to use.
// The default is set to nl
func WithCountry(countryCode string) ClientOption {
	return func(client *Client) {
		client.baseURL = makeUrl(strings.ToLower(countryCode), client.version)
		client.country = countryCode
	}
}

// WithUsername specify your username to be used for authentication.
func WithUsername(username string) ClientOption {
	return func(client *Client) {
		client.username = username
	}
}

// WithPassword specify your password to be used for authentication.
func WithPassword(password string) ClientOption {
	return func(client *Client) {
		client.secret = md5Hash(password)
	}
}

func md5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

// WithHashedPassword specify your md5 hashed password to be used for authentication.
func WithHashedPassword(hashedPassword string) ClientOption {
	return func(client *Client) {
		client.secret = hashedPassword
	}
}

// New Create a new instance of the picnic api client
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

// Authenticate using the configured username and secret the client will attempt ot authenticate
// Upon success the access token is stored in the Client.Token
// To leverage the token, calls to picnic endpoints expect the token to be added as a header named 'x-picnic-auth'
func (c *Client) Authenticate() error {
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
	if response.StatusCode != http.StatusOK {
		return c.parseError(response)
	}
	c.token = response.Header.Get("x-picnic-auth")
	return nil
}

// Logout assuming there is a valid session this endpoint will request the termination of the session and
// clear the token from the client.
func (c *Client) Logout() error {
	logoutUrl := c.baseURL + "/user/logout"
	err := c.post(logoutUrl, nil, nil)
	if err != nil {
		return err
	}
	c.token = ""
	return nil
}

// IsAuthenticated convince method to verify the client has an auth token.
func (c *Client) IsAuthenticated() bool {
	if c == nil {
		return false
	}
	return c.token != ""
}

func (c *Client) parseError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return fmt.Errorf("picnic-api: produced an error with code %d: %s empty error", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(body)

	var apiError struct {
		Error PicnicError `json:"error"`
	}

	jsonErr := json.NewDecoder(buf).Decode(&apiError)
	if jsonErr != nil {
		return fmt.Errorf("picnic-api: produced an [%d] response with an unprocessable error payload [%s]", resp.StatusCode, body)
	}
	if apiError.Error.IsCartError() {
		return apiError.Error.CreateCheckoutError()
	}
	return fmt.Errorf("picnic-api: produced an error with code [%s] and message [%s]", apiError.Error.Code, apiError.Error.Message)
}

func (c *Client) get(url string, result interface{}) error {
	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return requestErr
	}
	c.configureHeaders(request)
	if strings.Contains(url, "deliveries") {
		headerErr := c.includeAgentHeader(request)
		if headerErr != nil {
			return headerErr
		}
	}
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return httpErr
	}
	defer res.Body.Close()
	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return c.parseError(res)
	}
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
	c.configureHeaders(request)
	res, httpErr := c.http.Do(request)
	if httpErr != nil {
		return httpErr
	}
	defer res.Body.Close()
	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return c.parseError(res)
	}
	if result == nil {
		return nil
	}
	jsonErr := json.NewDecoder(res.Body).Decode(result)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

// GetArticleImageUrl generates the static url for a given article image.
//
// Open endpoint Authentication not required
func (c *Client) GetArticleImageUrl(articleImageId string, size ImageSize) (string, error) {
	if strings.TrimSpace(articleImageId) == "" {
		return "", createError("GetArticleImageUrl requires an articleImageId string")
	}
	imageBaseUrl := strings.Split(c.baseURL, "api")
	return fmt.Sprintf("%sstatic/images/%s/%s.png", imageBaseUrl[0], articleImageId, size.String()), nil
}

// GetArticleImage retrieves and decodes the png image of a given article.
//
// Open endpoint Authentication not required
func (c *Client) GetArticleImage(articleImageId string, size ImageSize) (*image.Image, error) {
	url, urlErr := c.GetArticleImageUrl(articleImageId, size)
	if urlErr != nil {
		return nil, urlErr
	}
	res, resErr := c.http.Get(url)
	if resErr != nil {
		return nil, resErr
	}
	if res.StatusCode != http.StatusOK {
		return nil, c.parseError(res)
	}
	articleImage, imageErr := png.Decode(res.Body)
	if imageErr != nil {
		return nil, imageErr
	}
	return &articleImage, nil
}

func makeUrl(countryCode string, version string) string {
	return fmt.Sprintf("https://storefront-prod.%s.picnicinternational.com/api/%s", countryCode, version)
}

func (c *Client) configureHeaders(request *http.Request) {
	request.Header.Set("x-picnic-auth", c.token)
	request.Header.Set("Content-Type", "application/json")
}

func (c *Client) includeAgentHeader(request *http.Request) error {
	if c.parsedToken == nil {
		err := c.parseJwt()
		if err != nil {
			return err
		}
	}
	request.Header.Set("x-picnic-agent", fmt.Sprintf("%d;%s;", c.parsedToken.PcClid, appVersion))
	request.Header.Set("x-picnic-did", c.parsedToken.PcDid)
	return nil
}
