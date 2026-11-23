# destinations

A list of all methods in the `destinations` service. Click on the method name to view detailed information about that method.

| Methods                               | Description       |
| :------------------------------------ | :---------------- |
| [ListDestinations](#listdestinations) | List Destinations |

## ListDestinations

List Destinations

- HTTP Method: `GET`
- Endpoint: `/destinations`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅        | Default go language context |

**Return Type**

[`ListDestinationsResponse`](../models/list_destinations_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Destinations.ListDestinations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```
