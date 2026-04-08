# ESim

A list of all methods in the `ESim` service. Click on the method name to view detailed information about that method.

| Methods                           | Description      |
| :-------------------------------- | :--------------- |
| [GetEsim](#getesim)               | Get eSIM         |
| [GetEsimDevice](#getesimdevice)   | Get eSIM Device  |
| [GetEsimHistory](#getesimhistory) | Get eSIM History |

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
  "example.com/celitech"
  "example.com/celitech/esim"
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
  "example.com/celitech"
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
  "example.com/celitech"
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
