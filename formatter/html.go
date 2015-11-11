package formatter

import (
   "fmt"
   "strings"
	"github.com/Fepelus/goldhouse/entities"
)

func HtmlFormat(input entities.DatedPrices) string {
  data := make([]string, len(input))
  for i, price := range input {
    data[i] = fmt.Sprintf("[new Date('%s'), %d]", price.Datef(), price.Ratio())
  }
  return header + strings.Join(data, ",\n") + footer
}

const header = `<html>
<head>
  <script type="text/javascript" src="https://www.google.com/jsapi"></script>
  <script type="text/javascript">
    google.load('visualization', '1.1', {packages: ['line']});
    google.setOnLoadCallback(drawChart);
    function drawChart() {
      var data = new google.visualization.DataTable();
      data.addColumn('date', 'Date');
      data.addColumn('number', 'Oz Au / house');
      data.addRows([
`

const footer = `
      ]);
      var options = {
        chart: {
          title: 'Median price of a house in Brunswick',
          subtitle: 'in ounces of gold'
        },
        width: 900,
        height: 500
      };
      var chart = new google.charts.Line(document.getElementById('linechart_material'));
      chart.draw(data, options);
    }
  </script>
</head>
<body>
  <div id="linechart_material"></div>
</body>
</html>
`