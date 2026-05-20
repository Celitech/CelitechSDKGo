# Consumption

A list of all methods in the `Consumption` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                                                                                                                                                                      |
| :------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetPurchaseConsumption](#getpurchaseconsumption) | This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages. |

## GetPurchaseConsumption

This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.

- HTTP Method: `GET`
- Endpoint: `/purchases/{purchaseId}/consumption`

**Parameters**

| Name       | Type                                | Required | Description                   |
| :--------- | :---------------------------------- | :------- | :---------------------------- |
| ctx        | Context                             | ✅       | Default go language context   |
| purchaseID | string                              | ✅       |                               |
| params     | GetPurchaseConsumptionRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/consumption"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := consumption.GetPurchaseConsumptionRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Consumption.GetPurchaseConsumption(context.Background(), "purchaseId", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
