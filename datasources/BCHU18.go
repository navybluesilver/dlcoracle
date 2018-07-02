package datasources

import (
	"fmt"
)

type BCHU18 struct {
}

func (ds *BCHU18) Id() uint64 {
	id, err := getId("BCHU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *BCHU18) Name() string {
	return "BCHU18"
}

func (ds *BCHU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "BCHU18")
}

func (ds *BCHU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *BCHU18) Value() (uint64, error) {
	return future("BCH","U18")
}
