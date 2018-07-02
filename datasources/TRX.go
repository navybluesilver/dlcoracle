package datasources

import (
	"fmt"
)

type TRX struct {
}

func (ds *TRX) Id() uint64 {
	id, err := getId("TRX")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *TRX) Name() string {
	return "TRX"
}

func (ds *TRX) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "TRX")
}

func (ds *TRX) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *TRX) Value() (uint64, error) {
	return spot("TRX")
}
