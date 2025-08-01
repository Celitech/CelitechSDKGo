```go
import (
  "fmt"
  "encoding/json"
  "github.com/Celitech/CelitechSDKGo/pkg/celitechconfig"
  "github.com/Celitech/CelitechSDKGo/pkg/celitech"

  "github.com/Celitech/CelitechSDKGo/pkg/esim"
)

config := celitechconfig.NewConfig()
config.SetBaseUrl("BASE_URL")
config.SetClientId("CLIENT_ID")
config.SetClientSecret("CLIENT_SECRET")
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
