```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"
  "github.com/Celitech/CelitechSDKGo/pkg/util"
  "github.com/Celitech/CelitechSDKGo/pkg/purchases"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


request := purchases.TopUpEsimRequest{
  Iccid: util.ToPointer("Iccid"),
  DataLimitInGb: util.ToPointer(float64(123)),
  StartDate: util.ToPointer("StartDate"),
  EndDate: util.ToPointer("EndDate"),
}

response, err := client.Purchases.TopUpEsim(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
