# e_sim

A list of all methods in the `e_sim` service. Click on the method name to view detailed information about that method.

| Methods                           | Description                                                                                                                                                    |
| :-------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetEsim](#getesim)               | Get eSIM                                                                                                                                                       |
| [GetEsimDevice](#getesimdevice)   | Get eSIM Device                                                                                                                                                |
| [GetEsimHistory](#getesimhistory) | Get eSIM History                                                                                                                                               |
| [SuspendEsim](#suspendesim)       | Suspends an eSIM to prevent abuse. All of the eSIM's packages are terminated and its connectivity is switched off. The eSIM is not refunded: the charge stays. |

## GetEsim

Get eSIM

- HTTP Method: `GET`
- Endpoint: `/esim`

**Parameters**

| Name  | Type    | Required | Description                 |
| :---- | :------ | :------- | :-------------------------- |
| ctx   | Context | ✅        | Default go language context |
| Iccid | string  | ✅        | ID of the eSIM              |

**Return Type**

[`GetEsimResponse`](../models/get_esim_response.md)

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

response, err := sdk.ESim.GetEsim(context.Background(), &api.GetEsimRequest{
  Iccid: "1111222233334444555000",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEsimDevice

Get eSIM Device

- HTTP Method: `GET`
- Endpoint: `/esim/{iccid}/device`

**Parameters**

| Name  | Type    | Required | Description                 |
| :---- | :------ | :------- | :-------------------------- |
| ctx   | Context | ✅        | Default go language context |
| Iccid | string  | ✅        | ID of the eSIM              |

**Return Type**

[`GetEsimDeviceResponse`](../models/get_esim_device_response.md)

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

response, err := sdk.ESim.GetEsimDevice(context.Background(), &api.GetEsimDeviceRequest{
  Iccid: "1111222233334444555000",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEsimHistory

Get eSIM History

- HTTP Method: `GET`
- Endpoint: `/esim/{iccid}/history`

**Parameters**

| Name  | Type    | Required | Description                 |
| :---- | :------ | :------- | :-------------------------- |
| ctx   | Context | ✅        | Default go language context |
| Iccid | string  | ✅        | ID of the eSIM              |

**Return Type**

[`GetEsimHistoryResponse`](../models/get_esim_history_response.md)

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

response, err := sdk.ESim.GetEsimHistory(context.Background(), &api.GetEsimHistoryRequest{
  Iccid: "1111222233334444555000",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SuspendEsim

Suspends an eSIM to prevent abuse. All of the eSIM's packages are terminated and its connectivity is switched off. The eSIM is not refunded: the charge stays.

- HTTP Method: `POST`
- Endpoint: `/esim/{iccid}/suspend`

**Parameters**

| Name           | Type    | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| :------------- | :------ | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ctx            | Context | ✅        | Default go language context                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Iccid          | string  | ✅        | ID of the eSIM                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| IdempotencyKey | string  | ❌        | Optional. A unique value (for example a UUID) that makes this request safe to retry. If a request with the same `Idempotency-Key` and the same body succeeds, repeating it returns the original response instead of performing the action again, and the response includes the header `Idempotent-Replayed: true`. Keys are scoped to your credentials and remembered for 3 days. Only successful (2xx) responses are remembered, so a request that failed can be retried with the same key. Reusing a key with a different request returns `422`; sending it while the original request is still being processed returns `409`. If the idempotency store is temporarily unavailable, the request is not processed and returns `503`; retry it. Must be 1 to 255 printable ASCII characters, without spaces. |

**Return Type**

[`SuspendEsimResponse`](../models/suspend_esim_response.md)

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

response, err := sdk.ESim.SuspendEsim(context.Background(), &api.SuspendEsimRequest{
  Iccid: "1111222233334444555000",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```
