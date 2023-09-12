# Picnic-API
![workflow](https://github.com/simonmartyr/picnic-api/actions/workflows/tests.yaml/badge.svg)

This is an (unofficial) Go wrapper for working with the Picnic WebAPI.

## Installation

To install the library:

`go get github.com/simonmartyr/picnic-api`

## Getting started

To create a client & authenticate you and choose to provide a valid access token
or credentials to manually authenticate with.

```go
//authentication with a token
package main
import picnic "github.com/simonmartyr/picnic-api"

func main() {
    client := picnic.New(&http.Client{},
		picnic.WithToken("your accessToken"),
    )
}
```
or

```go
//authentication using auth credentials
package main
import picnic "github.com/simonmartyr/picnic-api"

func main() {
    client := picnic.New(&http.Client{},
        picnic.WithUserName("user@emailaddress.com"),
        picnic.WithHashedPassword("hashedPassword"),
    )
	err := client.Authenticate()
	if err != nil {
		panic(err.Error())
    }
}
```

## Demo Project

For a complete example of the API being used, checkout [picnic-tui](https://github.com/simonmartyr/picnic-tui/)

## References & Credits

Thanks to the efforts of [MRVDH](https://github.com/MRVDH/picnic-api/) which documented the API.