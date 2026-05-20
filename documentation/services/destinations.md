# Destinations

A list of all methods in the `Destinations` service. Click on the method name to view detailed information about that method.

| Methods                               | Description       |
| :------------------------------------ | :---------------- |
| [ListDestinations](#listdestinations) | List Destinations |

## ListDestinations

List Destinations

- HTTP Method: `GET`
- Endpoint: `/destinations`

**Parameters**

| Name   | Type                          | Required | Description                   |
| :----- | :---------------------------- | :------- | :---------------------------- |
| ctx    | Context                       | ✅       | Default go language context   |
| params | ListDestinationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/destinations"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := destinations.ListDestinationsRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Destinations.ListDestinations(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
