package datasources

import (
	"fmt"
)

type USD struct {
}

func (ds *USD) Id() uint64 {
	id, err := getId("USD")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *USD) Name() string {
	return "USD"
}

func (ds *USD) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "USD")
}

func (ds *USD) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *USD) Value() (uint64, error) {
	return spot("USD")
}
