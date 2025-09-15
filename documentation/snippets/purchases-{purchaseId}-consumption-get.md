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

response, err := client.Purchases.GetPurchaseConsumption(context.Background(), "4973fa15-6979-4daa-9cf3-672620df819c")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
