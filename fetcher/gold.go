package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Fepelus/goldhouse/entities"
)

type goldFetcher struct {
	datedprices entities.DatedPrices
}

func NewGold(datedprices entities.DatedPrices) goldFetcher {
	return goldFetcher{datedprices}
}

func (this goldFetcher) Fetch() entities.DatedPrices {
	jsonString := this.call(this.url())
	golddata := this.parseJson(jsonString)
	return this.addGoldToPrices(golddata)
}

func (this goldFetcher) url() string {
	return "https://www.quandl.com/api/v3/datasets/WGC/GOLD_DAILY_AUD.json?start_date=2010-07-01"
}

func (this goldFetcher) call(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not fetch the Propertydata page: ", err)
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Thank you https://mholt.github.io/json-to-go/
type GoldData struct {
	Dataset struct {
		Data [][]interface{} `json:"data"`
	} `json:"dataset"`
}

func (this goldFetcher) parseJson(markup string) GoldData {
	jsonv := []byte(markup)
	var f GoldData
	_ = json.Unmarshal(jsonv, &f)
	return f
}
