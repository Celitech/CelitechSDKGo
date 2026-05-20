# Topup

A list of all methods in the `Topup` service. Click on the method name to view detailed information about that method.

| Methods                 | Description                                                                                                                                                                                                                                           |
| :---------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [TopUpESim](#topupesim) | This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag. |

## TopUpESim

This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.

- HTTP Method: `POST`
- Endpoint: `/purchases/topup`

**Parameters**

| Name             | Type                   | Required | Description                   |
| :--------------- | :--------------------- | :------- | :---------------------------- |
| ctx              | Context                | ✅       | Default go language context   |
| topUpESimRequest | TopUpESimRequest       | ✅       |                               |
| params           | TopUpESimRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/topup"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := topup.TopUpESimRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}


request := topup.TopUpESimRequest{}

response, err := client.Topup.TopUpESim(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
