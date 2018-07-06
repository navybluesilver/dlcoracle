package datasources

import (
	"encoding/json"
	"math"
	"net/http"
	"fmt"

	"github.com/gertjaap/dlcoracle/logging"
)

type MinApiCryptoCompareBTCResponse struct {
	ADA struct {
		BTC float64 `json:"BTC"`
	} `json:"ADA"`
	BCH struct {
		BTC float64 `json:"BTC"`
	} `json:"BCH"`
	EOS struct {
		BTC float64 `json:"BTC"`
	} `json:"EOS"`
	ETH struct {
		BTC float64 `json:"BTC"`
	} `json:"ETH"`
	LTC struct {
		BTC float64 `json:"BTC"`
	} `json:"LTC"`
	TRX struct {
		BTC float64 `json:"BTC"`
	} `json:"TRX"`
	USD struct {
		BTC float64 `json:"BTC"`
	} `json:"USD"`
	XRP struct {
		BTC float64 `json:"BTC"`
	} `json:"XRP"`
}

func spot(underlying string) (uint64, error) {
	for _, t := range allSpot {
		if t.Name == underlying {
			return t.Value, nil
		}
	}
	return 0, nil
}

func getAllSpot() ([]Ticker) {
	var t []Ticker
	req, err := http.NewRequest("GET", fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=ADA,BCH,EOS,ETH,LTC,TRX,USD,XRP&tsyms=BTC"), nil)
	if err != nil {
		logging.Error.Printf("SPOT.getAllSpot - NewRequest\n", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Printf("SPOT.getAllSpot - Do: \n",  err)
	}

	defer resp.Body.Close()

	var spots MinApiCryptoCompareBTCResponse

	if err := json.NewDecoder(resp.Body).Decode(&spots); err != nil {
		logging.Error.Printf("SPOT.getAllSpot - Json decode failed: \n", err)
	}

	satoshiValue := uint64(math.Floor((spots.ADA.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "ADA", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.BCH.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "BCH", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.EOS.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "EOS", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.ETH.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "ETH", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.LTC.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "LTC", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.BCH.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "BCH", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.TRX.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "TRX", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.USD.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "USD", Value: satoshiValue})
	satoshiValue = uint64(math.Floor((spots.XRP.BTC*100000000)+0.5)) * 1
	t = append(t, Ticker{Name: "XRP", Value: satoshiValue})
	return t
}
