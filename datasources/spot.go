package datasources

import (
	"encoding/json"
	"math"
	"net/http"
	"fmt"

	"github.com/gertjaap/dlcoracle/logging"
)


func spot(underlying string) (uint64, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=BTC",underlying), nil)
	if err != nil {
		logging.Error.Printf("%s.Value - NewRequest\n", underlying, err)
		return 0, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Printf("%s.Value - Do: \n", underlying, err)
		return 0, err
	}

	defer resp.Body.Close()

	var record MinApiCryptoCompareBTCResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logging.Error.Printf("%s.Value - Json decode failed: \n", underlying, err)
		return 0, err
	}

	satoshiValue := uint64(math.Floor((record.Value*100000000)+0.5)) * 1
	return satoshiValue, nil
}
