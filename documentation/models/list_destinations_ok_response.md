# ListDestinationsOkResponse

**Properties**

| Name         | Type                        | Required | Description |
| :----------- | :-------------------------- | :------- | :---------- |
| Destinations | []destinations.Destinations | ❌       |             |

# Destinations

**Properties**

| Name               | Type     | Required | Description                                                                                                                                                                                                                                                                                                               |
| :----------------- | :------- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Name               | string   | ❌       | Name of the destination                                                                                                                                                                                                                                                                                                   |
| Destination        | string   | ❌       | ISO representation of the destination                                                                                                                                                                                                                                                                                     |
| SupportedCountries | []string | ❌       | This array indicates the geographical area covered by a specific destination. If the destination represents a single country, the array will include that country. However, if the destination represents a broader regional scope, the array will be populated with the names of the countries belonging to that region. |
