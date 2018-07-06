package datasources

import (
	"fmt"
)

type USDZ18 struct {
}

func (ds *USDZ18) Id() uint64 {
	id, err := getId("USDZ18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *USDZ18) Name() string {
	return "USDZ18"
}

func (ds *USDZ18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "USDZ18")
}

func (ds *USDZ18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *USDZ18) Value() (uint64, error) {
	return future("USD","Z18")
}
