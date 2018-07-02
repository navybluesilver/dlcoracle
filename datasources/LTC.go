package datasources

import (
	"fmt"
)

type LTC struct {
}

func (ds *LTC) Id() uint64 {
	id, err := getId("LTC")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *LTC) Name() string {
	return "LTC"
}

func (ds *LTC) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "LTC")
}

func (ds *LTC) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *LTC) Value() (uint64, error) {
	return spot("LTC")
}
