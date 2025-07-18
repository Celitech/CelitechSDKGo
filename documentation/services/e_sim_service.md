# ESimService

A list of all methods in the `ESimService` service. Click on the method name to view detailed information about that method.

| Methods                           | Description      |
| :-------------------------------- | :--------------- |
| [GetEsim](#getesim)               | Get eSIM Status  |
| [GetEsimDevice](#getesimdevice)   | Get eSIM Device  |
| [GetEsimHistory](#getesimhistory) | Get eSIM History |
| [GetEsimMac](#getesimmac)         | Get eSIM MAC     |

## GetEsim

Get eSIM Status

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
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

  "github.com/Celitech/CelitechSDKGo/pkg/esim"
)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)


params := esim.GetEsimRequestParams{
  Iccid: util.toPointer(util.ToPointer("Iccid")),
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
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimDevice(context.Background(), "iccid")
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
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimHistory(context.Background(), "iccid")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## GetEsimMac

Get eSIM MAC

- HTTP Method: `GET`
- Endpoint: `/esim/{iccid}/mac`

**Parameters**

| Name  | Type    | Required | Description                 |
| :---- | :------ | :------- | :-------------------------- |
| ctx   | Context | ✅       | Default go language context |
| iccid | string  | ✅       | ID of the eSIM              |

**Return Type**

`GetEsimMacOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BaseUrl")
config.SetTimeout("Timeout")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimMac(context.Background(), "iccid")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
