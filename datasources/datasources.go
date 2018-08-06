package datasources

import (
	"fmt"
)

type Datasource interface {
	Id() uint64
	Name() string
	Description() string
	Value() (uint64, error)
	Interval() uint64
}

type Ticker struct {
	Name string
	Value uint64
}

var allSpot []Ticker
var allFuture []Ticker

func GetAllDatasources() []Datasource {
	var datasources []Datasource
	allSpot = getAllSpot()
	allFuture = getAllFuture()
	datasources = append(datasources, &USD{})
	datasources = append(datasources, &USDU18{})
	datasources = append(datasources, &USDZ18{})
	datasources = append(datasources, &ADA{})
	datasources = append(datasources, &ADAU18{})
	datasources = append(datasources, &BCH{})
	datasources = append(datasources, &BCHU18{})
	datasources = append(datasources, &EOS{})
	datasources = append(datasources, &EOSU18{})
	datasources = append(datasources, &ETH{})
	datasources = append(datasources, &ETHU18{})
	datasources = append(datasources, &LTC{})
	datasources = append(datasources, &LTCU18{})
	datasources = append(datasources, &TRX{})
	datasources = append(datasources, &TRXU18{})
	datasources = append(datasources, &XRP{})
	datasources = append(datasources, &XRPU18{})
  datasources = append(datasources, &XAU{})
	return datasources
}

func GetDatasource(id uint64) (Datasource, error) {
	switch id {
	case 1:
		return &USD{}, nil
	case 2:
		return &USDU18{}, nil
	case 3:
		return &USDZ18{}, nil
	case 4:
		return &ADA{}, nil
	case 5:
		return &ADAU18{}, nil
	case 6:
		return &BCH{}, nil
	case 7:
		return &BCHU18{}, nil
	case 8:
		return &EOS{}, nil
	case 9:
		return &EOSU18{}, nil
	case 10:
		return &ETH{}, nil
	case 11:
		return &ETHU18{}, nil
	case 12:
		return &LTC{}, nil
	case 13:
		return &LTCU18{}, nil
	case 14:
		return &TRX{}, nil
	case 15:
		return &TRXU18{}, nil
	case 16:
		return &XRP{}, nil
	case 17:
		return &XRPU18{}, nil
	case 18:
		return &XAU{}, nil
	default:
		return nil, fmt.Errorf("Data source with ID %d not known", id)
	}
}

func HasDatasource(id uint64) bool {
	return (id <= 18)
}

func getId(underlying string) (id uint64, err error) {
	switch underlying {
	case "USD":
		return 1, nil
	case "USDU18":
		return 2, nil
	case "USDZ18":
		return 3, nil
	case "ADA":
		return 4, nil
	case "ADAU18":
		return 5, nil
	case "BCH":
		return 6, nil
	case "BCHU18":
		return 7, nil
	case "EOS":
		return 8, nil
	case "EOSU18":
		return 9, nil
	case "ETH":
		return 10, nil
	case "ETHU18":
		return 11, nil
	case "LTC":
		return 12, nil
	case "LTCU18":
		return 13, nil
	case "TRX":
		return 14, nil
	case "TRXU18":
		return 15, nil
	case "XRP":
		return 16, nil
	case "XRPU18":
		return 17, nil
	case "XAU":
		return 18, nil
	default:
		return 0, fmt.Errorf("Data source with ID %d not known", id)
	}
}
