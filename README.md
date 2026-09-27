# Celitech Go SDK 2.0.7

Welcome to the Celitech SDK documentation. This guide will help you get started with integrating and using the Celitech SDK in your project.

## Versions

- API version: `2.0.7`
- SDK version: `2.0.7`

## About the API

Welcome to the CELITECH API documentation!

Useful links: [Homepage](https://www.celitech.com) | [Support email](mailto:devops@celitech.com) | [Blog](https://www.celitech.com/blog/)

# Introduction

This guide is your go-to resource for the CELITECH API, with full documentation and schemas.

Need help? Email us at devops@celitech.com.

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

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
- [Authentication](#authentication)
  - [OAuth Authentication](#oauth-authentication)
  - [Environment Variables](#environment-variables)
- [Setting a Custom Timeout](#setting-a-custom-timeout)
- [Sample Usage](#sample-usage)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### OAuth Authentication

The Celitech API uses OAuth for authentication.

You need to provide the OAuth parameters when initializing the SDK.

```go
config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)
```

## Environment Variables

These are the environment variables for the SDK:

| Name          | Description             |
| :------------ | :---------------------- |
| CLIENT_ID     | Client ID parameter     |
| CLIENT_SECRET | Client Secret parameter |

Environment variables are a way to configure your application outside the code. You can set these environment variables on the command line or use your project's existing tooling for managing environment variables.

If you are using a `.env` file, a template with the variable names is provided in the `.env.example` file located in the same directory as this README.

## Setting a Custom Timeout

You can set a custom timeout for the SDK's HTTP requests as follows:

```go
import "time"

config := celitech.NewConfig()

sdk := celitech.NewCelitech(config)

sdk.SetTimeout(10 * time.Second)
```

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v2"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)

response, err := client.Destinations.ListDestinations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                   |
| :----------------------------------------------------- |
| [Destinations](documentation/services/destinations.md) |
| [Packages](documentation/services/packages.md)         |
| [Purchases](documentation/services/purchases.md)       |
| [ESim](documentation/services/e_sim.md)                |
| [IFrame](documentation/services/i_frame.md)            |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `CelitechResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                       | Description                                 |
| :------- | :------------------------- | :------------------------------------------ |
| Data     | `T`                        | The body of the API response                |
| Metadata | `CelitechResponseMetadata` | Status code and headers returned by the API |

#### `CelitechError[T]`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                    | Description                                                       |
| :------- | :---------------------- | :---------------------------------------------------------------- |
| Err      | `error`                 | The error that occurred                                           |
| Data     | `*T`                    | The deserialized error response data (nil if unmarshaling failed) |
| Body     | `[]byte`                | The raw body of the API response                                  |
| Metadata | `CelitechErrorMetadata` | Status code and headers returned by the API                       |

#### `CelitechResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details>
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                             | Description |
| :----------------------------------------------------------------------------------------------- | :---------- |
| [ListDestinationsOkResponse](documentation/models/list_destinations_ok_response.md)              |             |
| [ListPackagesOkResponse](documentation/models/list_packages_ok_response.md)                      |             |
| [CreatePurchaseV2OkResponse](documentation/models/create_purchase_v2_ok_response.md)             |             |
| [CreatePurchaseV2Request](documentation/models/create_purchase_v2_request.md)                    |             |
| [ListPurchasesOkResponse](documentation/models/list_purchases_ok_response.md)                    |             |
| [CreatePurchaseOkResponse](documentation/models/create_purchase_ok_response.md)                  |             |
| [CreatePurchaseRequest](documentation/models/create_purchase_request.md)                         |             |
| [TopUpEsimOkResponse](documentation/models/top_up_esim_ok_response.md)                           |             |
| [TopUpEsimRequest](documentation/models/top_up_esim_request.md)                                  |             |
| [EditPurchaseOkResponse](documentation/models/edit_purchase_ok_response.md)                      |             |
| [EditPurchaseRequest](documentation/models/edit_purchase_request.md)                             |             |
| [GetPurchaseConsumptionOkResponse](documentation/models/get_purchase_consumption_ok_response.md) |             |
| [GetEsimOkResponse](documentation/models/get_esim_ok_response.md)                                |             |
| [GetEsimDeviceOkResponse](documentation/models/get_esim_device_ok_response.md)                   |             |
| [GetEsimHistoryOkResponse](documentation/models/get_esim_history_ok_response.md)                 |             |
| [TokenOkResponse](documentation/models/token_ok_response.md)                                     |             |
| [GrantType](documentation/models/grant_type.md)                                                  |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
