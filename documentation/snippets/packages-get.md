```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"

  "github.com/elmasrisaer/GoSDKTest/pkg/packages"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


params := packages.ListPackagesRequestParams{

}

response, err := client.Packages.ListPackages(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```
