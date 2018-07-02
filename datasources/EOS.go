package datasources

import (
	"fmt"
)

type EOS struct {
}

func (ds *EOS) Id() uint64 {
	id, err := getId("EOS")
	if err != nil {
		panic(err)
	}
	return id
}

func (ds *EOS) Name() string {
	return "EOS"
}

func (ds *EOS) Description() string {
	return fmt.Sprintf("Publishes the value of %s denominated in 1/100000000th units of BTC (satoshi)", "EOS")
}

func (ds *EOS) Interval() uint64 {
	return 300 // every 5 minutes
}

func (ds *EOS) Value() (uint64, error) {
	return spot("EOS")
}
