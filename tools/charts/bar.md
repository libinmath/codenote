// plot
package plot

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"os"
)

var wikiCommitStat = []struct {
	date    string
	commits int
}{
	{"8/21-8/27", 20},
	{"8/28-9/3", 35},
	{"9/4-9/10", 21},
	{"9/11-9/17", 13},
	{"9/18-9/24", 6},
	{"9/25-10/1", 4},
	{"10/2-10/8", 4},
	{"10/9-10/15", 3},
	{"10/16-10/22", 0},
	{"10/23-10/29", 6},
	{"10/30-11/5", 19},
	{"11/6-11/12", 7},
	{"11/13-11/19", 13},
	{"11/20-11/26", 28},
	{"11/27-12/3", 32},
	{"12/4-12/10", 41},
	{"12/11-12/17", 14},
	{"12/18-12/24", 4},
	{"12/25-12/31", 0},
	{"1/1-1/7", 12},
	{"1/8-1/14", 3},
	{"1/15-1/21", 2},
	{"1/22-1/28", 4},
	{"1/29-2/4", 18},
	{"2/5-2/11", 20},
	{"2/12-2/18", 19},
	{"2/19-2/25", 16},
	{"2/26-3/4", 13},
	{"3/5-3/11", 15},
	{"3/12-3/18", 18},
	{"3/19-3/25", 32},
}

func generateCommitItems() ([]opts.BarData, []string) {
	items := make([]opts.BarData, 0)
	xaxis := make([]string, 0)

	for _, v := range wikiCommitStat {
		xaxis = append(xaxis, v.date)
		items = append(items, opts.BarData{Value: v.commits})
	}
	return items, xaxis
}

func BarExample() {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "团队知识贡献",
			Subtitle: "wiki仓提交数统计",
			Left:     "middle",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "400px",
			Theme:  types.ThemeWonderland,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:         true,
				Interval:     "0",
				Rotate:       35,
				ShowMinLabel: true,
				ShowMaxLabel: true,
				FontSize:     "10",
			},
			Name: "周",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "提交数",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{XAxisIndex: []int{0}, Start: 25, End: 100}),
	)

	// Put data into instance
	items, xaxis := generateCommitItems()
	bar.SetXAxis(xaxis).
		AddSeries("Commit", items)
	bar.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "top",
		}),
		charts.WithBarChartOpts(opts.BarChart{
			BarCategoryGap: "40%",
		}),
		charts.WithMarkLineNameCoordItemOpts(opts.MarkLineNameCoordItem{
			Name:        "2023年",
			Coordinate0: []interface{}{"1/1-1/7", 0},
			Coordinate1: []interface{}{"1/1-1/7", 40},
		}),
		charts.WithMarkLineStyleOpts(opts.MarkLineStyle{
			Symbol:     []string{"pin"},
			SymbolSize: 10,
			Label:      &opts.Label{Show: true},
		}),
	)
	f, _ := os.Create("./tools/plot/bar.html")
	err := bar.Render(f)
	if err != nil {
		fmt.Println(err)
		return
	}
}
