# i_frame

A list of all methods in the `i_frame` service. Click on the method name to view detailed information about that method.

| Methods         | Description                                   |
| :-------------- | :-------------------------------------------- |
| [Token](#token) | Generate a new token to be used in the iFrame |

## Token

Generate a new token to be used in the iFrame

- HTTP Method: `POST`
- Endpoint: `/iframe/token`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅        | Default go language context |

**Return Type**

[`TokenResponse`](../models/token_response.md)

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

response, err := sdk.IFrame.Token(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```
