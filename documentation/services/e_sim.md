# ESim

A list of all methods in the `ESim` service. Click on the method name to view detailed information about that method.

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

| Name   | Type           | Required | Description                   |
| :----- | :------------- | :------- | :---------------------------- |
| ctx    | Context        | ✅       | Default go language context   |
| params | GetEsimRequest | ✅       | Additional request parameters |

**Return Type**

`api.GetEsimResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v3"
)

config := api.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := api.NewClient(config)

response, err := client.ESim.GetEsim(context.Background(), &api.GetEsimRequest{
  Iccid: api.Ptr("1111222233334444555000"),
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

| Name   | Type                 | Required | Description                   |
| :----- | :------------------- | :------- | :---------------------------- |
| ctx    | Context              | ✅       | Default go language context   |
| iccid  | string               | ✅       | ID of the eSIM                |
| params | GetEsimDeviceRequest | ✅       | Additional request parameters |

**Return Type**

`api.GetEsimDeviceResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v3"
)

config := api.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := api.NewClient(config)

response, err := client.ESim.GetEsimDevice(context.Background(), &api.GetEsimDeviceRequest{
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

| Name   | Type                  | Required | Description                   |
| :----- | :-------------------- | :------- | :---------------------------- |
| ctx    | Context               | ✅       | Default go language context   |
| iccid  | string                | ✅       | ID of the eSIM                |
| params | GetEsimHistoryRequest | ✅       | Additional request parameters |

**Return Type**

`api.GetEsimHistoryResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v3"
)

config := api.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := api.NewClient(config)

response, err := client.ESim.GetEsimHistory(context.Background(), &api.GetEsimHistoryRequest{
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

| Name   | Type               | Required | Description                   |
| :----- | :----------------- | :------- | :---------------------------- |
| ctx    | Context            | ✅       | Default go language context   |
| iccid  | string             | ✅       | ID of the eSIM                |
| params | SuspendEsimRequest | ✅       | Additional request parameters |

**Return Type**

`api.SuspendEsimResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v3"
)

config := api.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := api.NewClient(config)

response, err := client.ESim.SuspendEsim(context.Background(), &api.SuspendEsimRequest{
  Iccid: "1111222233334444555000",
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```
