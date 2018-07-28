package datasources

import (
	"fmt"
	"math"
	"net/http"
  "io/ioutil"
	"strconv"
	"github.com/gertjaap/dlcoracle/logging"
)

type XAU struct {
}

func (ds *XAU) Id() uint64 {
	id, err := getId("XAU")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *XAU) Name() string {
	return "XAU"
}

func (ds *XAU) Description() string {
	return fmt.Sprintf("Publishes the value of %s (per gram) denominated in 1/100000000th units of BTC (satoshi)", "XAU")
}

func (ds *XAU) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *XAU) Value() (uint64, error) {
	req, err := http.NewRequest("GET", "https://api.vaultoro.com/latest/", nil)
	if err != nil {
		logging.Error.Printf("XAU.Value - NewRequest\n", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Printf("XAU.Value - Do: \n",  err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	satoshiValue, _ := strconv.ParseFloat(string(body), 64)
	satoshiValue = math.Floor((satoshiValue*100000000)+0.5)

	return uint64(satoshiValue), nil
}
