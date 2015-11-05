package fetcher

type Fetcher interface {
	Fetch() HousePrices
}
