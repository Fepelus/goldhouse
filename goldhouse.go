package main

import (
	"fmt"

	"github.com/Fepelus/goldhouse/fetcher"
	"github.com/Fepelus/goldhouse/formatter"
)

func main() {
	propertydata := fetcher.NewProperty()
	houseprices := propertydata.Fetch()
	datedprices := fetcher.ConvertToDatedPrices(houseprices)
	goldfetcher := fetcher.NewGold(datedprices)
	matcheddata := goldfetcher.Fetch()
//  fmt.Println(formatter.CsvFormat(matcheddata))
  fmt.Println(formatter.HtmlFormat(matcheddata))
}
