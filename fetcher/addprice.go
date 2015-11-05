package fetcher

import (
	"fmt"

	"github.com/Fepelus/goldhouse/entities"
)

func (this goldFetcher) addGoldToPrices(golddata GoldData) entities.DatedPrices {
	fmt.Println(golddata)
	return this.datedprices
}
