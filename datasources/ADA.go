package datasources

import (
	"fmt"
)

type ADA struct {
}

func (ds *ADA) Id() uint64 {
	id, err := getId("ADA")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *ADA) Name() string {
	return "ADA"
}

func (ds *ADA) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "ADA")
}

func (ds *ADA) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *ADA) Value() (uint64, error) {
	return spot("ADA")
}
