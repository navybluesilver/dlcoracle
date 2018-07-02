package datasources

import (
	"fmt"
)

type LTCU18 struct {
}

func (ds *LTCU18) Id() uint64 {
	id, err := getId("LTCU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *LTCU18) Name() string {
	return "LTCU18"
}

func (ds *LTCU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "LTCU18")
}

func (ds *LTCU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *LTCU18) Value() (uint64, error) {
	return future("LTC","U18")
}
