package picnic

import (
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const mockToken = `fakeheader.eyJzdWIiOiI4MDgtODEwLTAzMzkiLCJwYzpjbGlkIjoyMDEwMCwicGM6cHY6ZW5hYmxlZCI6ZmFsc2UsInBjOmxvZ2ludHMiOjE1NjEwNDUyNTYxMDcsImlzcyI6InBpY25pYy1kZXYiLCJwYzpwdjp2ZXJpZmllZCI6dHJ1ZSwicGM6MmZhIjoiTk9UX1JFUVVJUkVEIiwicGM6cm9sZSI6IlNUQU5EQVJEX1VTRVIiLCJwYzpkaWQiOiJCQTAyNjAxOC1GMEVFLTRDOTAtQTFENC00Q0MzNjg3RUE4ODEiLCJleHAiOjE3MjAyNjM5NzgsImlhdCI6MTcwNDcxMTk3OCwianRpIjoiMlJJN0c0UFUifQ.FakeSiganture`

func testClient(code int, body io.Reader, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.WriteHeader(code)
		_, _ = io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := &Client{
		http:    http.DefaultClient,
		baseURL: server.URL + "/",
		token:   mockToken,
	}
	return client, server
}

func testImageClient(code int, body io.Reader, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(code)
		_, _ = io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := &Client{
		http:    http.DefaultClient,
		baseURL: server.URL + "/",
		token:   mockToken,
	}
	return client, server
}

func testClientFile(code int, filename string, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testClient(code, f, validators...)
}

func testClientImage(code int, filename string, validators ...func(*http.Request)) (*Client, *httptest.Server) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testImageClient(code, f, validators...)
}

func TestAuthenticate(t *testing.T) {
	tokenVal := "tokenVal"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-picnic-auth", tokenVal)
		w.WriteHeader(http.StatusOK)
	}))
	client := &Client{
		http:    http.DefaultClient,
		baseURL: server.URL + "/",
	}
	defer server.Close()
	err := client.Authenticate()
	if err != nil {
		t.Fatal(err)
	}
	if client.token != tokenVal {
		t.Error("Token not set correctly")
	}
}

func TestGetArticleImageUrl(t *testing.T) {
	articleId := "s1005863"
	expectedUrl := "https://storefront-prod.nl.picnicinternational.com/static/images/s1005863/medium.png"
	c := New(&http.Client{})

	pictureUrl, err := c.GetArticleImageUrl(articleId, Medium)
	if err != nil {
		t.Fatal(err)
	}
	if pictureUrl != expectedUrl {
		t.Error("Invalid url")
	}
}

func TestGetArticleImage(t *testing.T) {
	articleId := "s1005863"
	c, s := testClientImage(http.StatusOK, "test/augurken.png")
	defer s.Close()
	res, err := c.GetArticleImage(articleId, Medium)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Error("Invalid Image")
	}
}

func TestGetArticleImage_NotFound(t *testing.T) {
	articleId := "s1005863"
	c, s := testClientFile(http.StatusNotFound, "test/error.json")
	defer s.Close()
	res, err := c.GetArticleImage(articleId, Medium)
	if res != nil {
		t.Error("Unexpected result")
	}
	if err == nil {
		t.Error("No error produced")
	}
}

func TestNew(t *testing.T) {
	c := New(&http.Client{})
	if c.baseURL != "https://storefront-prod.nl.picnicinternational.com/api/17" {
		t.Error("Invalid baseurl")
	}
	if c.version != "17" {
		t.Error("Invalid version")
	}
	if c.country != "nl" {
		t.Error("Invalid version")
	}
}

func TestNew_With_Version(t *testing.T) {
	version := "19"
	c := New(&http.Client{}, WithVersion(version))
	if c.baseURL != "https://storefront-prod.nl.picnicinternational.com/api/19" {
		t.Error("Invalid baseurl")
	}
	if c.version != version {
		t.Error("Invalid version")
	}
}

func TestNew_With_Country(t *testing.T) {
	country := "be"
	c := New(&http.Client{}, WithCountry(country))
	if c.baseURL != "https://storefront-prod.be.picnicinternational.com/api/17" {
		t.Error("Invalid baseurl")
	}
	if c.country != country {
		t.Error("Invalid version")
	}
}

func TestNew_With_BaseUrl(t *testing.T) {
	bespokeUrl := "https://whateveryourwant/api/17"
	c := New(&http.Client{}, WithBaseUrl(bespokeUrl))
	if c.baseURL != bespokeUrl {
		t.Error("Invalid baseurl")
	}
}

func TestNew_With_Country_And_Version(t *testing.T) {
	country := "be"
	version := "19"
	c := New(&http.Client{}, WithCountry(country), WithVersion(version))
	if c.baseURL != "https://storefront-prod.be.picnicinternational.com/api/19" {
		t.Error("Invalid baseurl")
	}
	if c.version != version {
		t.Error("Invalid version")
	}
	if c.country != country {
		t.Error("Invalid version")
	}
}

func TestNew_With_Username(t *testing.T) {
	email := "johnny@jojo.com"
	c := New(&http.Client{}, WithUsername(email))
	if c.username != email {
		t.Error("Invalid email")
	}
}

func TestNew_With_Password(t *testing.T) {
	password := "examplePass"
	expected := "8ba9442c2dc0e616d26b5ab84162ed48"
	c := New(&http.Client{}, WithPassword(password))
	if c.secret != expected {
		t.Error("Invalid secret")
	}
}

func TestNew_With_Hashed_Password(t *testing.T) {
	hashed := "8ba9442c2dc0e616d26b5ab84162ed48"
	c := New(&http.Client{}, WithHashedPassword(hashed))
	if c.secret != hashed {
		t.Error("Invalid secret")
	}
}

func TestNew_With_Token(t *testing.T) {
	token := "tokenValue"
	c := New(&http.Client{}, WithToken(token))
	if c.token != token {
		t.Error("Invalid token")
	}
}

func Test_Integration(t *testing.T) {
	godotenv.Load()
	c := New(&http.Client{},
		WithUsername(os.Getenv("USERNAME")),
		WithHashedPassword(os.Getenv("SECRET")),
	)
	authErr := c.Authenticate()
	if authErr != nil {
		t.Error("auth failed")
	}
	user, userErr := c.GetUser()
	if userErr != nil {
		t.Fatal(userErr)
	}
	if user == nil {
		t.Error("invalid user")
	}
}

func Test_Logout(t *testing.T) {
	c, s := testClientFile(http.StatusOK, "test/empty.json")
	defer s.Close()
	err := c.Logout()
	if err != nil {
		t.Fatal(err.Error())
	}
	if c.token != "" {
		t.Error("Invalid token value")
	}
}

func TestClient_IsAuthenticated_False(t *testing.T) {
	c := New(&http.Client{})
	if c.IsAuthenticated() {
		t.Error("Invalid authenticated result")
	}
}

func TestClient_IsAuthenticated_True(t *testing.T) {
	c := New(&http.Client{}, WithToken("example"))
	if !c.IsAuthenticated() {
		t.Error("Invalid authenticated result")
	}
}

func TestClient_IsAuthenticated_False_Nil(t *testing.T) {
	var c *Client
	if c.IsAuthenticated() {
		t.Error("Invalid authenticated result")
	}
}
