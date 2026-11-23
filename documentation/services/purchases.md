# purchases

A list of all methods in the `purchases` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| :------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreatePurchaseV2](#createpurchasev2)             | This endpoint is used to purchase a new eSIM by providing the package details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [ListPurchases](#listpurchases)                   | This endpoint can be used to list all the successful purchases made between a given interval.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| [CreatePurchase](#createpurchase)                 | This endpoint is used to purchase a new eSIM by providing the package details.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| [TopUpEsim](#topupesim)                           | This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.                                                                                                                                                                                                                                                                                                                                                                                                    |
| [EditPurchase](#editpurchase)                     | This endpoint allows you to modify the validity dates of an existing purchase. **Behavior:** - If the purchase has **not yet been activated**, both the start and end dates can be updated. - If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set). - Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category. The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits. |
| [GetPurchaseConsumption](#getpurchaseconsumption) | This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |

## CreatePurchaseV2

This endpoint is used to purchase a new eSIM by providing the package details.

- HTTP Method: `POST`
- Endpoint: `/purchases/v2`

**Parameters**

| Name           | Type                                                                                                       | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| :------------- | :--------------------------------------------------------------------------------------------------------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx            | Context                                                                                                    | ✅        | Default go language context                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Destination    | string                                                                                                     | ✅        | ISO representation of the package's destination. Supports both ISO2 (e.g., 'FR') and ISO3 (e.g., 'FRA') country codes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| DataLimitInGb  | float64                                                                                                    | ✅        | Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB. Use `-1` to purchase an unlimited (date-based) package — provide `startDate`/`endDate` spanning 3 to 30 days (`duration` is not supported for unlimited packages).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| StartDate      | time.Time                                                                                                  | ❌        | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| EndDate        | time.Time                                                                                                  | ❌        | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| Duration       | float64                                                                                                    | ❌        | Duration of the package in days. Available values are 1, 2, 7, 14, 30, or 90. Either provide startDate/endDate or duration. Not supported for unlimited packages (`dataLimitInGB` = -1), which are date-based — provide startDate/endDate instead.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| Quantity       | float64                                                                                                    | ✅        | Number of eSIMs to purchase.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| Email          | string                                                                                                     | ❌        | Email address where the purchase confirmation email will be sent (including QR Code & activation steps)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ReferenceID    | string                                                                                                     | ❌        | An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes. Letters, digits, `-`, `_`, `.`, `@` and `+` are allowed (so an email address works), up to 255 characters.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| NetworkBrand   | string                                                                                                     | ❌        | Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters, numbers, dots (.), ampersands (&), and spaces. This feature is available to platforms with Diamond tier only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| EmailBrand     | string                                                                                                     | ❌        | Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Language       | [CreatePurchaseV2RequestLanguage](../models/create_purchase_v2_request.md#createpurchasev2requestlanguage) | ❌        | Language of the confirmation email sent to the customer.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| IdempotencyKey | string                                                                                                     | ❌        | Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces. |

**Return Type**

[`[]*CreatePurchaseV2ResponseItem`](../models/create_purchase_v2_response_item.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.CreatePurchaseV2(context.Background(), &api.CreatePurchaseV2Request{
  Destination: "FRA",
  DataLimitInGb: 1,
  Quantity: 1,
})
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

| Name        | Type      | Required | Description                                                                                                                                                                                                         |
| :---------- | :-------- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| ctx         | Context   | ✅        | Default go language context                                                                                                                                                                                         |
| PurchaseID  | string    | ❌        | ID of the purchase                                                                                                                                                                                                  |
| Iccid       | string    | ❌        | ID of the eSIM                                                                                                                                                                                                      |
| AfterDate   | time.Time | ❌        | Start date of the interval for filtering purchases in the format 'yyyy-MM-dd'                                                                                                                                       |
| BeforeDate  | time.Time | ❌        | End date of the interval for filtering purchases in the format 'yyyy-MM-dd'                                                                                                                                         |
| Email       | string    | ❌        | Email associated to the purchase.                                                                                                                                                                                   |
| ReferenceID | string    | ❌        | The referenceId that was provided by the partner during the purchase or topup flow.                                                                                                                                 |
| AfterCursor | string    | ❌        | To get the next batch of results, use this parameter. It tells the API where to start fetching data after the last item you received. It helps you avoid repeats and efficiently browse through large sets of data. |
| Limit       | float64   | ❌        | Maximum number of purchases to be returned in the response. The value must be greater than 0 and less than or equal to 100. If not provided, the default value is 20                                                |
| After       | float64   | ❌        | Epoch value representing the start of the time interval for filtering purchases                                                                                                                                     |
| Before      | float64   | ❌        | Epoch value representing the end of the time interval for filtering purchases                                                                                                                                       |

**Return Type**

[`ListPurchasesResponse`](../models/list_purchases_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.ListPurchases(context.Background(), &api.ListPurchasesRequest{})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreatePurchase

This endpoint is used to purchase a new eSIM by providing the package details.

- HTTP Method: `POST`
- Endpoint: `/purchases`

**Parameters**

| Name           | Type                                                                                                | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| :------------- | :-------------------------------------------------------------------------------------------------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx            | Context                                                                                             | ✅        | Default go language context                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Destination    | string                                                                                              | ✅        | ISO representation of the package's destination. Supports both ISO2 (e.g., 'FR') and ISO3 (e.g., 'FRA') country codes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| DataLimitInGb  | float64                                                                                             | ✅        | Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB. Use `-1` to purchase an unlimited (date-based) package — provide `startDate`/`endDate` spanning 3 to 30 days.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| StartDate      | time.Time                                                                                           | ✅        | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| EndDate        | time.Time                                                                                           | ✅        | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| Email          | string                                                                                              | ❌        | Email address where the purchase confirmation email will be sent (including QR Code & activation steps)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ReferenceID    | string                                                                                              | ❌        | An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes. Letters, digits, `-`, `_`, `.`, `@` and `+` are allowed (so an email address works), up to 255 characters.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| NetworkBrand   | string                                                                                              | ❌        | Customize the network brand of the issued eSIM. The `networkBrand` parameter cannot exceed 15 characters in length and must contain only letters, numbers, dots (.), ampersands (&), and spaces. This feature is available to platforms with Diamond tier only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| EmailBrand     | string                                                                                              | ❌        | Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| Language       | [CreatePurchaseRequestLanguage](../models/create_purchase_request.md#createpurchaserequestlanguage) | ❌        | Language of the confirmation email sent to the customer.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| StartTime      | float64                                                                                             | ❌        | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| EndTime        | float64                                                                                             | ❌        | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| IdempotencyKey | string                                                                                              | ❌        | Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces. |

**Return Type**

[`CreatePurchaseResponse`](../models/create_purchase_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"
  "time"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.CreatePurchase(context.Background(), &api.CreatePurchaseRequest{
  Destination: "FRA",
  DataLimitInGb: 1,
  StartDate: time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
  EndDate: time.Date(2023, 11, 20, 0, 0, 0, 0, time.UTC),
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## TopUpEsim

This endpoint is used to top-up an existing eSIM with the previously associated destination by providing its ICCID and package details. To determine if an eSIM can be topped up, use the Get eSIM endpoint, which returns the `isTopUpAllowed` flag.

- HTTP Method: `POST`
- Endpoint: `/purchases/topup`

**Parameters**

| Name           | Type      | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| :------------- | :-------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx            | Context   | ✅        | Default go language context                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Iccid          | string    | ✅        | ID of the eSIM                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| DataLimitInGb  | float64   | ✅        | Size of the package in GB. The available options are 0.5, 1, 2, 3, 5, 8, 20, 50GB. Use `-1` to top up with an unlimited (date-based) package — provide `startDate`/`endDate` spanning 3 to 30 days (`duration` is not supported for unlimited packages).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| StartDate      | time.Time | ❌        | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| EndDate        | time.Time | ❌        | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| Duration       | float64   | ❌        | Duration of the package in days. Available values are 1, 2, 7, 14, 30, or 90. Either provide startDate/endDate or duration. Not supported for unlimited packages (`dataLimitInGB` = -1), which are date-based — provide startDate/endDate instead.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| Email          | string    | ❌        | Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| ReferenceID    | string    | ❌        | An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes. Letters, digits, `-`, `_`, `.`, `@` and `+` are allowed (so an email address works), up to 255 characters.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| EmailBrand     | string    | ❌        | Customize the email subject brand. The `emailBrand` parameter cannot exceed 25 characters in length and must contain only letters, numbers, and spaces. This feature is available to platforms with Diamond tier only.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| StartTime      | float64   | ❌        | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| EndTime        | float64   | ❌        | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| IdempotencyKey | string    | ❌        | Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces. |

**Return Type**

[`TopUpEsimResponse`](../models/top_up_esim_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.TopUpEsim(context.Background(), &api.TopUpEsimRequest{
  Iccid: "1111222233334444555000",
  DataLimitInGb: 1,
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## EditPurchase

This endpoint allows you to modify the validity dates of an existing purchase.  

**Behavior:**
- If the purchase has **not yet been activated**, both the start and end dates can be updated.  
- If the purchase is **already active**, only the **end date** can be updated, while the **start date must remain unchanged** (and should be passed as originally set).  
- Updates must comply with the same pricing structure; the modification cannot alter the package size or change its duration category.  

The end date can be extended or shortened as long as it adheres to the same pricing category and does not exceed the allowed duration limits.

- HTTP Method: `POST`
- Endpoint: `/purchases/edit`

**Parameters**

| Name           | Type      | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| :------------- | :-------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx            | Context   | ✅        | Default go language context                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| PurchaseID     | string    | ✅        | ID of the purchase                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| StartDate      | time.Time | ✅        | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| EndDate        | time.Time | ✅        | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| StartTime      | float64   | ❌        | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| EndTime        | float64   | ❌        | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| IdempotencyKey | string    | ❌        | Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces. |

**Return Type**

[`EditPurchaseResponse`](../models/edit_purchase_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"
  "time"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.EditPurchase(context.Background(), &api.EditPurchaseRequest{
  PurchaseID: "ae471106-c8b4-42cf-b83a-b061291f2922",
  StartDate: time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
  EndDate: time.Date(2023, 11, 20, 0, 0, 0, 0, time.UTC),
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetPurchaseConsumption

This endpoint can be called for consumption notifications (e.g. every 1 hour or when the user clicks a button). It returns the data balance (consumption) of purchased packages.

- HTTP Method: `GET`
- Endpoint: `/purchases/{purchaseId}/consumption`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅        | Default go language context |
| PurchaseID | string  | ✅        | ID of the purchase          |

**Return Type**

[`GetPurchaseConsumptionResponse`](../models/get_purchase_consumption_response.md)

**Example Usage Code Snippet**

```go
import (
  "context"
  "fmt"

  api "github.com/Celitech/CelitechSDKGo/v3"
  "github.com/Celitech/CelitechSDKGo/v3/client"
  "github.com/Celitech/CelitechSDKGo/v3/option"
)

sdk := client.NewCelitech(
  option.WithClientCredentials("CLIENT_ID", "CLIENT_SECRET"),
)

response, err := sdk.Purchases.GetPurchaseConsumption(context.Background(), &api.GetPurchaseConsumptionRequest{
  PurchaseID: "4973fa15-6979-4daa-9cf3-672620df819c",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```
