```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

)

config := celitechconfig.NewConfig()
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
client := celitech.NewCelitech(config)

response, err := client.Purchases.GetPurchaseConsumption(context.Background(), "purchaseId")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
