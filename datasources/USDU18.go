package datasources

import (
	"fmt"
)

type USDU18 struct {
}

func (ds *USDU18) Id() uint64 {
	id, err := getId("USDU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *USDU18) Name() string {
	return "USDU18"
}

func (ds *USDU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "USDU18")
}

func (ds *USDU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *USDU18) Value() (uint64, error) {
	return future("XBT","U18")
}
