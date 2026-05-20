# Token

A list of all methods in the `Token` service. Click on the method name to view detailed information about that method.

| Methods                         | Description                                   |
| :------------------------------ | :-------------------------------------------- |
| [GenerateToken](#generatetoken) | Generate a new token to be used in the iFrame |

## GenerateToken

Generate a new token to be used in the iFrame

- HTTP Method: `POST`
- Endpoint: `/iframe/token`

**Parameters**

| Name   | Type                       | Required | Description                   |
| :----- | :------------------------- | :------- | :---------------------------- |
| ctx    | Context                    | ✅       | Default go language context   |
| params | GenerateTokenRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/token"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := token.GenerateTokenRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Token.GenerateToken(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
