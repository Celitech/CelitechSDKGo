# Esim

A list of all methods in the `Esim` service. Click on the method name to view detailed information about that method.

| Methods             | Description |
| :------------------ | :---------- |
| [GetESim](#getesim) | Get eSIM    |

## GetESim

Get eSIM

- HTTP Method: `GET`
- Endpoint: `/esim`

**Parameters**

| Name   | Type                 | Required | Description                   |
| :----- | :------------------- | :------- | :---------------------------- |
| ctx    | Context              | ✅       | Default go language context   |
| params | GetESimRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/esim"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := esim.GetESimRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Esim.GetESim(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
