package datasources

import (
	"fmt"
)

type XRP struct {
}

func (ds *XRP) Id() uint64 {
	id, err := getId("XRP")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *XRP) Name() string {
	return "XRP"
}

func (ds *XRP) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "XRP")
}

func (ds *XRP) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *XRP) Value() (uint64, error) {
	return spot("XRP")
}
