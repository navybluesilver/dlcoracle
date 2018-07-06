package datasources

import (
	"encoding/json"
	"math"
	"net/http"
	"fmt"

	"github.com/gertjaap/dlcoracle/logging"
)


type BitmexResponse struct {
	Symbol												 string	         `json:"symbol"`
	LastPrice                      float64         `json:"lastPrice"`
}

type AllBitmexResponse []struct {
	Symbol												 string	         `json:"symbol"`
	LastPrice                      float64         `json:"lastPrice"`
}


func future(underlying string, future string) (uint64, error) {
	name := fmt.Sprintf("%s%s", underlying, future)
	for _, t := range allFuture {
		if t.Name == name {
			return t.Value, nil
		}
	}
	fmt.Println(name)
	return 0, nil
}

func getAllFuture() ([]Ticker) {
	var t []Ticker
	req, err := http.NewRequest("GET", "https://www.bitmex.com/api/v1/instrument/active", nil)
	if err != nil {
		logging.Error.Printf("Future.getAllFuture - NewRequest \n", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Printf("Future.getAllFuture - Do: \n", err)
	}

	defer resp.Body.Close()

	var futures AllBitmexResponse

	if err := json.NewDecoder(resp.Body).Decode(&futures); err != nil {
		logging.Error.Printf("Future.getAllFuture - Json decode failed: \n", err)
	}

	for _, symbol := range futures {
		f := convertSymbolTicker(symbol)
		t = append(t, f)
	}
	return t
}

func convertSymbolTicker(r BitmexResponse) (Ticker) {
	satoshiValue := uint64(math.Floor((r.LastPrice*100000000)+0.5)) * 1
	name := r.Symbol

	if name == "XBTU18" {
			name = "USDU18"
			satoshiValue = uint64(math.Floor(((1/r.LastPrice)*100000000)+0.5)) * 1
	}
	if name == "XBTZ18" {
			name = "USDZ18"
			satoshiValue = uint64(math.Floor(((1/r.LastPrice)*100000000)+0.5)) * 1
	}

	return Ticker{Name: name, Value: satoshiValue}

/*
*/
}
