package datasources

import (
	"encoding/json"
	"math"
	"net/http"
	"fmt"

	"github.com/gertjaap/dlcoracle/logging"
)


func future(underlying string, future string) (uint64, error) {
	ticker := fmt.Sprintf("%s%s", underlying, future)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.bitmex.com/api/v1/instrument?symbol=%s&count=100&reverse=false",ticker), nil)
	if err != nil {
		logging.Error.Printf("%s.Value - NewRequest \n", ticker, err)
		return 0, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Printf("%s.Value - Do: \n", ticker, err)
		return 0, err
	}

	defer resp.Body.Close()

	var record BitmexLastPriceResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logging.Error.Printf("%s.Value - Json decode failed: \n", ticker, err)
		return 0, err
	}

	value := record[0].LastPrice

	if underlying == "XBT" {
			value = 1/value
	}


	satoshiValue := uint64(math.Floor((value*100000000)+0.5)) * 1
	return satoshiValue, nil
}
