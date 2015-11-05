package main

import (
	"fmt"

	"github.com/Fepelus/goldhouse/fetcher"
)

func main() {
	propertydata := fetcher.NewProperty()
	houseprices := propertydata.Fetch()
	datedprices := fetcher.ConvertToDatedPrices(houseprices)
	goldfetcher := fetcher.NewGold(datedprices)
	matcheddata := goldfetcher.Fetch()
	fmt.Println(matcheddata)
}
