package picnic

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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
