```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"
  "github.com/Celitech/CelitechSDKGo/pkg/util"
  "github.com/Celitech/CelitechSDKGo/pkg/purchases"
)

config := celitechconfig.NewConfig()
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


request := purchases.TopUpEsimRequest{
  Iccid: util.ToPointer("1111222233334444555000"),
  DataLimitInGb: util.ToPointer(float64(1)),
}

response, err := client.Purchases.TopUpEsim(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
