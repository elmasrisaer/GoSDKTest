```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)

response, err := client.ESim.GetEsimDevice(context.Background(), "iccid")
if err != nil {
  panic(err)
}

fmt.Println(response)

```
