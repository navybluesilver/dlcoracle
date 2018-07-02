package datasources

import (
	"fmt"
)

type ETHU18 struct {
}

func (ds *ETHU18) Id() uint64 {
	id, err := getId("ETHU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *ETHU18) Name() string {
	return "ETHU18"
}

func (ds *ETHU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "ETHU18")
}

func (ds *ETHU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *ETHU18) Value() (uint64, error) {
	return future("ETH","U18")
}
