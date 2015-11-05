package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Fepelus/goldhouse/entities"
)

type mapped struct {
	datatype string
	index    int
	content  string
}

type property struct {
	entries   int
	collector chan mapped
	completed chan bool
}

type Datapoint struct {
	Month string
	Year  string
	Price string
}

type housePrices map[int]Datapoint
type HousePrices []Datapoint

func NewProperty() Fetcher {
	return property{0, make(chan mapped), make(chan bool)}
}

func (this property) Fetch() HousePrices {
	markup := this.call(this.url())
	go this.makePrices(markup)
	return toHousePrices(this.collectResults())
}

func toHousePrices(input housePrices) HousePrices {
	output := make(HousePrices, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i]
	}
	return output
}

func (this property) url() string {
	return "http://propertydata.realestateview.com.au/propertydata/median-prices/victoria/brunswick/"
}

func (this property) call(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not fetch the Propertydata page: ", err)
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (this property) makePrices(markup string) {
	for _, line := range strings.Split(markup, "\n") {
		if strings.Index(line, "median_graph") > -1 {
			this.splitthis(line)
		}
	}
}

func (this property) splitthis(linetosplit string) {
	for _, line := range strings.Split(linetosplit, "&") {
		if strings.Index(line, "chxl=") == 0 {
			this.splitCalendar(line)
		}
		if strings.Index(line, "chd=") == 0 {
			this.splitPrices(line)
		}
	}

}

func (this property) splitCalendar(param string) {
	indices := strings.Split(param, ":")
	this.splitMonths(indices[1])
	this.splitYears(indices[3])
}

func (this property) splitMonths(monthline string) {
	idx := 0
	for _, month := range strings.Split(monthline, "|") {
		if len(month) == 3 {
			mpd := mapped{"month", idx, month}
			this.collector <- mpd
			idx += 1
		}
	}
	this.completed <- true
}

func (this property) splitYears(yearline string) {
	idx := 0
	for _, year := range strings.Split(yearline, "|") {
		if len(year) == 2 {
			mpd := mapped{"year", idx, fmt.Sprintf("20%s", year)}
			this.collector <- mpd
			idx += 1
		}
	}
	this.completed <- true
}

func (this property) splitPrices(line string) {
	indices := strings.Split(line, "|")
	this.splitPrice(indices[3])
}

func (this property) splitPrice(line string) {
	idx := 0
	for _, price := range strings.Split(line, ",") {
		mpd := mapped{"price", idx, price}
		this.collector <- mpd
		idx += 1
	}
	this.completed <- true
}

func (this property) collectResults() housePrices {
	expectedCompletions := 3
	priceCollection := make(housePrices)
	for {
		select {
		case <-this.completed:
			expectedCompletions -= 1
			if expectedCompletions == 0 {
				return priceCollection
			}
		case data := <-this.collector:
			_, exists := priceCollection[data.index]
			if !exists {
				priceCollection[data.index] = Datapoint{"blank", "blank", "blank"}
			}
			point := priceCollection[data.index]
			if data.datatype == "month" {
				point.Month = data.content
			} else if data.datatype == "year" {
				point.Year = data.content
			} else if data.datatype == "price" {
				point.Price = data.content
			}
			priceCollection[data.index] = point
		}
	}
}

func convertToDatedPrice(point Datapoint) entities.DatedPrice {
	thedate, _ := time.Parse("Jan 2006", fmt.Sprintf("%s %s", point.Month, point.Year))
	thehouseprice, _ := strconv.Atoi(point.Price)
	return entities.NewDatedPrice(thedate, thehouseprice, 0)
}

func ConvertToDatedPrices(input HousePrices) entities.DatedPrices {
	output := make(entities.DatedPrices, len(input))
	for i, point := range input {
		output[i] = convertToDatedPrice(point)
	}
	return output
}
