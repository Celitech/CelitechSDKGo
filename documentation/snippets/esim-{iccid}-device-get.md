```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BASE_URL")
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimDevice(context.Background(), "iccid")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
