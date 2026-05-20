# Device

A list of all methods in the `Device` service. Click on the method name to view detailed information about that method.

| Methods                         | Description     |
| :------------------------------ | :-------------- |
| [GetESimDevice](#getesimdevice) | Get eSIM Device |

## GetESimDevice

Get eSIM Device

- HTTP Method: `GET`
- Endpoint: `/esim/{iccid}/device`

**Parameters**

| Name   | Type                       | Required | Description                   |
| :----- | :------------------------- | :------- | :---------------------------- |
| ctx    | Context                    | ✅       | Default go language context   |
| iccid  | string                     | ✅       |                               |
| params | GetESimDeviceRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo"
  "github.com/Celitech/CelitechSDKGo/device"
)

config := celitech.NewConfig()
config.SetClientID("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


params := device.GetESimDeviceRequestParams{
  Accept: celitech.Nullable[string]("application/json"),
}

response, err := client.Device.GetESimDevice(context.Background(), "iccid", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
