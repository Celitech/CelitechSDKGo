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
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)


request := purchases.CreatePurchaseV2Request{
  Destination: util.ToPointer("FRA"),
  DataLimitInGb: util.ToPointer(float64(1)),
  StartDate: util.ToPointer("2023-11-01"),
  EndDate: util.ToPointer("2023-11-20"),
  Quantity: util.ToPointer(float64(1)),
}

response, err := client.Purchases.CreatePurchaseV2(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
