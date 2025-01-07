# TopUpEsimRequest

**Properties**

| Name          | Type    | Required | Description                                                                                                                                             |
| :------------ | :------ | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Iccid         | string  | ✅       | ID of the eSIM                                                                                                                                          |
| DataLimitInGb | float64 | ✅       | Size of the package in GB. The available options are 1, 2, 3, 5, 8, 20GB                                                                                |
| StartDate     | string  | ✅       | Start date of the package's validity in the format 'yyyy-MM-dd'. This date can be set to the current day or any day within the next 12 months.          |
| EndDate       | string  | ✅       | End date of the package's validity in the format 'yyyy-MM-dd'. End date can be maximum 90 days after Start date.                                        |
| Email         | string  | ❌       | Email address where the purchase confirmation email will be sent (excluding QR Code & activation steps)                                                 |
| ReferenceId   | string  | ❌       | An identifier provided by the partner to link this purchase to their booking or transaction for analytics and debugging purposes.                       |
| StartTime     | float64 | ❌       | Epoch value representing the start time of the package's validity. This timestamp can be set to the current time or any time within the next 12 months. |
| EndTime       | float64 | ❌       | Epoch value representing the end time of the package's validity. End time can be maximum 90 days after Start time.                                      |
