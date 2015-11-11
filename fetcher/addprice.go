package fetcher

import (
   "time"
	"github.com/Fepelus/goldhouse/entities"
)

func (this goldFetcher) addGoldToPrices(golddata GoldData) entities.DatedPrices {
  mappedGold := convertGolddataToMap(golddata)
  for i, _ := range this.datedprices {
    thisdate := this.datedprices[i].Date
    datestring := thisdate.Format("2006-01-02")
    for price := mappedGold[datestring]; price == 0; price = mappedGold[datestring] {
      thisdate = thisdate.Add(-24 * time.Hour)
      datestring = thisdate.Format("2006-01-02")
    } 
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
