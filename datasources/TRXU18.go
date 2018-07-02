package datasources

import (
	"fmt"
)

type TRXU18 struct {
}

func (ds *TRXU18) Id() uint64 {
	id, err := getId("TRXU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *TRXU18) Name() string {
	return "TRXU18"
}

func (ds *TRXU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "TRXU18")
}

func (ds *TRXU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *TRXU18) Value() (uint64, error) {
	return future("TRX","U18")
}
