# History

A list of all methods in the `History` service. Click on the method name to view detailed information about that method.

| Methods                           | Description      |
| :-------------------------------- | :--------------- |
| [GetESimHistory](#getesimhistory) | Get eSIM History |

## GetESimHistory

Get eSIM History

- HTTP Method: `GET`
- Endpoint: `/esim/{iccid}/history`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| iccid  | string                      | ✅       |                               |
| params | GetESimHistoryRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/history"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := history.GetESimHistoryRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.History.GetESimHistory(context.Background(), "iccid", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
