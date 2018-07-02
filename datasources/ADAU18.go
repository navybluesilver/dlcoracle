package datasources

import (
	"fmt"
)

type ADAU18 struct {
}

func (ds *ADAU18) Id() uint64 {
	id, err := getId("ADAU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *ADAU18) Name() string {
	return "ADAU18"
}

func (ds *ADAU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "ADAU18")
}

func (ds *ADAU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *ADAU18) Value() (uint64, error) {
	return future("ADA","U18")
}
