# Celitech Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=Celitech%2FGo)

Welcome to the CELITECH API documentation!

Useful links: [Homepage](https://www.celitech.com) | [Support email](mailto:devops@celitech.com) | [Blog](https://www.celitech.com/blog/)

# Introduction

This is the reference documentation and schemas for the CELITECH API and webhooks. For tutorials and other documentation please refer to the [CELITECH Documentation](https://docs.celitech.com).

"Partners" refers to online service providers that use our eSIM API. Access levels include Gold, Platinum, and Diamond.

## API

The CELITECH API is designed for use by partner platforms, including both web and mobile applications. It's assumed all endpoint calls are initiated from the backend of an integrated platform.

API URL: `https://api.celitech.net/v1`

## Authentication & Authorization

CELITECH API uses the OAuth 2.0 protocol for authentication and authorization.
The endpoints are protected using client credentials flow which is based on a token exchange. The token has a defined life span (typically 1 hour), after which a new token must be obtained.

To begin, obtain OAuth 2.0 client credentials ( **CLIENT_ID** & **CLIENT_SECRET** ) from the [CELITECH Dashboard](https://www.dashboard.celitech.com/). Then your client application requests an access token from the CELITECH Authorization Server, extracts a token from the response, and sends the token to the CELITECH API that you want to access.

Security Scheme Type: `OAuth2`

Flow type: `clientCredentials`

Token URL: `https://auth.celitech.net/oauth2/token`

## Table of Contents

- [Reference](#reference)
- [Installation](#installation)
- [Usage](#usage)
- [Environments](#environments)
- [Errors](#errors)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Custom Client](#custom-client)
  - [Timeouts](#timeouts)
- [Contributing](#contributing)

## Reference

A full reference for this library is available [here](./reference.md).

## Installation

This SDK requires Go 1.22 or later.

```sh
go get github.com/Celitech/CelitechSDKGo/v3
```

## Usage

```go
import (
	apiclient "github.com/Celitech/CelitechSDKGo/v3/client"
	context "context"
	option "github.com/Celitech/CelitechSDKGo/v3/option"
)

client := apiclient.NewCelitech(
	option.WithClientID(
		"YOUR_CLIENT_ID",
	),
	option.WithClientSecret(
		"YOUR_CLIENT_SECRET",
	),
)
client.Destinations.ListDestinations(
	context.TODO(),
)
```

## Environments

This SDK allows you to configure a different environment for API requests by passing the
`WithBaseURL` request option to the client.

```go
import (
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
  "github.com/Celitech/CelitechSDKGo/v3"
)

client := client.NewCelitech(
  option.WithBaseURL(api.Environments.Default),
)
```

Available environments:

- `api.Environments.Default`

`WithBaseURL` also accepts any URL string, such as a staging server or a local mock.

## Errors

A non-2xx response returns an `error` wrapping a `*core.APIError`, which carries the status code
and raw response body:

```go
var apiErr *core.APIError
if errors.As(err, &apiErr) {
  fmt.Println(apiErr.StatusCode)
}
```

The `core` package also exposes typed errors for endpoints that declare them, each
embedding `*core.APIError` so the status code is still reachable through it.

## Request Options

Every method accepts variadic `opts ...option.RequestOption` to override behavior for a single
call — a different base URL, additional headers, or a longer timeout:

```go
response, err := client.SomeResource.SomeMethod(
  ctx,
  request,
  option.WithHTTPHeader(httpHeader),
  option.WithMaxAttempts(1),
)
```

## Advanced

### Custom Client

This SDK is built to work with any instance of `*http.Client`. By default, if no client is
provided, the SDK constructs one. You can pass your own through `option.WithHTTPClient`:

```go
customClient := &http.Client{}

client := client.NewClient(
  option.WithHTTPClient(customClient),
)
```

### Timeouts

```go
import (
  "net/http"
  "time"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

client := client.NewCelitech(
  option.WithHTTPClient(
    &http.Client{
      Timeout: 10 * time.Second,
    },
  ),
)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
