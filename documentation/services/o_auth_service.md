# OAuthService

A list of all methods in the `OAuthService` service. Click on the method name to view detailed information about that method.

| Methods                           | Description                       |
| :-------------------------------- | :-------------------------------- |
| [GetAccessToken](#getaccesstoken) | This endpoint was added by liblab |

## GetAccessToken

This endpoint was added by liblab

- HTTP Method: `POST`
- Endpoint: `/oauth2/token`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| getAccessTokenRequest | GetAccessTokenRequest | ✅       |                             |

**Return Type**

`GetAccessTokenOkResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitechconfig"
  "github.com/elmasrisaer/GoSDKTest/pkg/celitech"
  "github.com/elmasrisaer/GoSDKTest/pkg/util"
  "github.com/elmasrisaer/GoSDKTest/pkg/oauth"
)

config := celitechconfig.NewConfig()
client := celitech.NewCelitech(config)


request := oauth.GetAccessTokenRequest{

}

response, err := client.OAuth.GetAccessToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
