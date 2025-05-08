# TopUpEsimOkResponse

**Properties**

| Name     | Type                                  | Required | Description |
| :------- | :------------------------------------ | :------- | :---------- |
| Purchase | purchases.TopUpEsimOkResponsePurchase | ❌       |             |
| Profile  | purchases.TopUpEsimOkResponseProfile  | ❌       |             |

# TopUpEsimOkResponsePurchase

**Properties**

| Name        | Type    | Required | Description                                                                                       |
| :---------- | :------ | :------- | :------------------------------------------------------------------------------------------------ |
| Id          | string  | ❌       | ID of the purchase                                                                                |
| PackageId   | string  | ❌       | ID of the package                                                                                 |
| StartDate   | string  | ❌       | Start date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                        |
| EndDate     | string  | ❌       | End date of the package's validity in the format 'yyyy-MM-ddThh:mm:ssZZ'                          |
| Duration    | float64 | ❌       | It designates the number of days the eSIM is valid for within 90-day validity from issuance date. |
| CreatedDate | string  | ❌       | Creation date of the purchase in the format 'yyyy-MM-ddThh:mm:ssZZ'                               |
| StartTime   | float64 | ❌       | Epoch value representing the start time of the package's validity                                 |
| EndTime     | float64 | ❌       | Epoch value representing the end time of the package's validity                                   |

# TopUpEsimOkResponseProfile

**Properties**

| Name  | Type   | Required | Description    |
| :---- | :----- | :------- | :------------- |
| Iccid | string | ❌       | ID of the eSIM |
