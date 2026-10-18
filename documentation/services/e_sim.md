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

| Name   | Type                 | Required | Description                   |
| :----- | :------------------- | :------- | :---------------------------- |
| ctx    | Context              | ✅       | Default go language context   |
| params | GetEsimRequestParams | ✅       | Additional request parameters |

**Return Type**

`GetEsimOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v2"
  "github.com/Celitech/CelitechSDKGo/v2/esim"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := esim.GetEsimRequestParams{
  Iccid: celitech.Ptr("1111222233334444555000"),
}

response, err := client.ESim.GetEsim(context.Background(), params)
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
| ctx   | Context | ✅       | Default go language context |
| iccid | string  | ✅       | ID of the eSIM              |

**Return Type**

`GetEsimDeviceOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v2"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimDevice(context.Background(), "1111222233334444555000")
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
| ctx   | Context | ✅       | Default go language context |
| iccid | string  | ✅       | ID of the eSIM              |

**Return Type**

`GetEsimHistoryOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v2"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimHistory(context.Background(), "1111222233334444555000")
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

| Name   | Type                     | Required | Description                   |
| :----- | :----------------------- | :------- | :---------------------------- |
| ctx    | Context                  | ✅       | Default go language context   |
| iccid  | string                   | ✅       | ID of the eSIM                |
| params | SuspendEsimRequestParams | ✅       | Additional request parameters |

**Return Type**

`SuspendEsimOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/v2"
  "github.com/Celitech/CelitechSDKGo/v2/esim"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := esim.SuspendEsimRequestParams{}

response, err := client.ESim.SuspendEsim(context.Background(), "1111222233334444555000", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
