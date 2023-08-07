# WEATHER API SDK
## https://www.weatherapi.com/
### Example usage:
```go
package main

import (
	"encoding/json"
	"log"
	"os"

	weather "github.com/zumosik/weatherapi-sdk-go"
)

func main() {
	client, err := weather.New("173e99c92c3b4be9ba870813230708") // you can get token here https://www.weatherapi.com/my/
	if err != nil {
		log.Fatal(err)
	}

	q := "Paris"
	/*
		q can be:

		Latitude and Longitude (Decimal degree) e.g: q=48.8567,2.3508
		city name e.g.: q=Paris
		US zip e.g.: q=10001
		UK postcode e.g: q=SW1
		Canada postal code e.g: q=G2J
		metar:<metar code> e.g: q=metar:EGLL
		iata:<3 digit airport code> e.g: q=iata:DXB
		auto:ip IP lookup e.g: q=auto:ip
		IP address (IPv4 and IPv6 supported) e.g: q=100.0.0.1
		bulk New
	*/
	days := 3 // from 1 to 9

	res, err := client.GetForecast(q, days)
	if err != nil {
		log.Fatal(err)
	}
	file, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	_ = os.WriteFile("result.json", file, 0644)
	// you will get result formatted to json in result.json
}

```