package entities

import (
   "time"
   "fmt"
)

type DatedPrice struct {
	Date       time.Time
	HousePrice int // stored here in dollars
	GoldPrice  int // stored here in dollars
}

type DatedPrices []DatedPrice

func NewDatedPrice(indate time.Time, houseprice int, goldprice int) DatedPrice {
	return DatedPrice{indate, houseprice, goldprice}
}

func (this DatedPrice) String () string {
  return fmt.Sprintf("%s,%d,%d\n", this.Datef(), this.HousePrice, this.GoldPrice)
}

func (this DatedPrice) Datef() string {
  return this.Date.Format("2006-01-02")
}

func (this DatedPrice) Ratio() int {
  return this.HousePrice / this.GoldPrice
}