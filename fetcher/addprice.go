package fetcher

import (
	"github.com/Fepelus/goldhouse/entities"
)

func (this goldFetcher) addGoldToPrices(golddata GoldData) entities.DatedPrices {
  mappedGold := convertGolddataToMap(golddata)
  for i, _ := range this.datedprices {
    datestring := this.datedprices[i].Date.Format("2006-01-02")
    this.datedprices[i].GoldPrice = mappedGold[datestring]
  }
	return this.datedprices
}

func convertGolddataToMap (golddata GoldData) map[string]int {
  output := make (map[string]int)
  for _, current := range golddata.Dataset.Data {
    price := int(current[1].(float64))
    output[current[0].(string)] = price
    
  }
  return output
}
