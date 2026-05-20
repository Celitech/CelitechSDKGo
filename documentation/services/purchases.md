# Purchases

A list of all methods in the `Purchases` service. Click on the method name to view detailed information about that method.

| Methods                           | Description                                                                                   |
| :-------------------------------- | :-------------------------------------------------------------------------------------------- |
| [CreatePurchase](#createpurchase) | This endpoint is used to purchase a new eSIM by providing the package details.                |
| [ListPurchases](#listpurchases)   | This endpoint can be used to list all the successful purchases made between a given interval. |

## CreatePurchase

This endpoint is used to purchase a new eSIM by providing the package details.

- HTTP Method: `POST`
- Endpoint: `/purchases`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| createPurchaseRequest | CreatePurchaseRequest       | ✅       |                               |
| params                | CreatePurchaseRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/purchases"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := purchases.CreatePurchaseRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}


request := purchases.CreatePurchaseRequest{}

response, err := client.Purchases.CreatePurchase(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListPurchases

This endpoint can be used to list all the successful purchases made between a given interval.

- HTTP Method: `GET`
- Endpoint: `/purchases`

**Parameters**

| Name   | Type                       | Required | Description                   |
| :----- | :------------------------- | :------- | :---------------------------- |
| ctx    | Context                    | ✅       | Default go language context   |
| params | ListPurchasesRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/purchases"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := purchases.ListPurchasesRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Purchases.ListPurchases(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
