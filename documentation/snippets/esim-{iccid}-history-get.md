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
