package datasources

import (
	"fmt"
)

type BCH struct {
}

func (ds *BCH) Id() uint64 {
	id, err := getId("BCH")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *BCH) Name() string {
	return "BCH"
}

func (ds *BCH) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "BCH")
}

func (ds *BCH) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *BCH) Value() (uint64, error) {
	return spot("BCH")
}
