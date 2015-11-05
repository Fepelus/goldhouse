package entities

import "time"

type DatedPrice struct {
	Date       time.Time
	HousePrice int // stored here in dollars
	GoldPrice  int // stored here in dollars
}

type DatedPrices []DatedPrice

func NewDatedPrice(indate time.Time, houseprice int, goldprice int) DatedPrice {
	return DatedPrice{indate, houseprice, goldprice}
}
