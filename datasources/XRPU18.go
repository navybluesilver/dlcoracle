package datasources

import (
	"fmt"
)

type XRPU18 struct {
}

func (ds *XRPU18) Id() uint64 {
	id, err := getId("XRPU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *XRPU18) Name() string {
	return "XRPU18"
}

func (ds *XRPU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "XRPU18")
}

func (ds *XRPU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *XRPU18) Value() (uint64, error) {
	return future("XRP","U18")
}
