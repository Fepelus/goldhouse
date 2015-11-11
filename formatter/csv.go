package formatter

import (
   "fmt"
  "strings"
	"github.com/Fepelus/goldhouse/entities"
)

func CsvFormat(input entities.DatedPrices) string {
  lines := make([]string, len(input) + 1)
  lines[0] = fmt.Sprint("Date,ozau/house")
  for i, price := range input {
    lines[i+1] = fmt.Sprintf("%s,%d", price.Datef(), price.Ratio())
  }
  return strings.Join(lines, "\n")
}

