package datasources

import (
	"fmt"
)

type EOSU18 struct {
}

func (ds *EOSU18) Id() uint64 {
	id, err := getId("EOSU18")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *EOSU18) Name() string {
	return "EOSU18"
}

func (ds *EOSU18) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "EOSU18")
}

func (ds *EOSU18) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *EOSU18) Value() (uint64, error) {
	return future("EOS","U18")
}
