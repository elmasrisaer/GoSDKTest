# ListPurchasesOkResponse

**Properties**

| Name        | Type                  | Required | Description                                                                                                                                                                                                                                                                                     |
| :---------- | :-------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Purchases   | []purchases.Purchases | ❌       |                                                                                                                                                                                                                                                                                                 |
| AfterCursor | string                | ❌       | The cursor value representing the end of the current page of results. Use this cursor value as the "afterCursor" parameter in your next request to retrieve the subsequent page of results. It ensures that you continue fetching data from where you left off, facilitating smooth pagination. |

# Purchases

**Properties**

| Name        | Type                    | Required | Description                                                                                                                                                                                                       |
| :---------- | :---------------------- | :------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Id          | string                  | ❌       | ID of the purchase                                                                                                                                                                                                |
| StartDate   | string                  | ❌       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                        |
| EndDate     | string                  | ❌       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                          |
| CreatedDate | string                  | ❌       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'                                                                                                                                               |
| StartTime   | float64                 | ❌       | Epoch value representing the start time of the package's validity                                                                                                                                                 |
| EndTime     | float64                 | ❌       | Epoch value representing the end time of the package's validity                                                                                                                                                   |
| CreatedAt   | float64                 | ❌       | Epoch value representing the date of creation of the purchase                                                                                                                                                     |
| Package\_   | purchases.Package\_     | ❌       |                                                                                                                                                                                                                   |
| Esim        | purchases.PurchasesEsim | ❌       |                                                                                                                                                                                                                   |
| Source      | string                  | ❌       | The source indicates where the eSIM was purchased, which can be from the API, dashboard, landing-page or promo-page. For purchases made before September 8, 2023, the value will be displayed as 'Not available'. |
| ReferenceId | string                  | ❌       | The referenceId that was provided by the partner during the purchase or topup flow. This identifier can be used for analytics and debugging purposes.                                                             |

# Package\_

**Properties**

| Name             | Type    | Required | Description                                     |
| :--------------- | :------ | :------- | :---------------------------------------------- |
| Id               | string  | ❌       | ID of the package                               |
| DataLimitInBytes | float64 | ❌       | Size of the package in Bytes                    |
| Destination      | string  | ❌       | ISO representation of the package's destination |
| DestinationName  | string  | ❌       | Name of the package's destination               |
| PriceInCents     | float64 | ❌       | Price of the package in cents                   |

# PurchasesEsim

**Properties**

| Name  | Type   | Required | Description    |
| :---- | :----- | :------- | :------------- |
| Iccid | string | ❌       | ID of the eSIM |
