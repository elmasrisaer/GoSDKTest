```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"

  "github.com/elmasrisaer/GoSDKTest/pkg/esim"
)

config := celitechconfig.NewConfig()
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
