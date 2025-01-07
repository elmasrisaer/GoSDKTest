```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)

response, err := client.Destinations.ListDestinations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)

```
