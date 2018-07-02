package datasources

import (
	"fmt"
)

type ETH struct {
}

func (ds *ETH) Id() uint64 {
	id, err := getId("ETH")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *ETH) Name() string {
	return "ETH"
}

func (ds *ETH) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "ETH")
}

func (ds *ETH) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *ETH) Value() (uint64, error) {
	return spot("ETH")
}
