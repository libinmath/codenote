package charts

import (
	"io"
	"os"
	"testing"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 建议先在bash中使用grep -ir "xxxx" ./* | wc -l 命令进行词频统计，然后再生成词云。

var gs_db_key_words = map[string]interface{}{
	"ustore":      10000,
	"asotre":      6181,
	"存储":          6666,
	"openGauss":   4386,
	"高斯":          4386,
	"oracle":      4055,
	"甲骨文":         4055,
	"index":       2467,
	"索引":          2467,
	"undo":        2244,
	"历史版本":        2244,
	"uheap":       1898,
	"ubtree":      1898,
	"fsm":         1484,
	"空闲空间管理":      1689,
	"可见性":         1112,
	"buffer":      985,
	"缓存":          985,
	"cache":       985,
	"事务":          847,
	"transaction": 847,
	"锁":           847,
	"lock":        847,
	"并发":          582,
	"元组":          555,
	"tuple":       550,
	"回收":          462,
	"recycle":     366,
	"回滚":          282,
	"rollback":    273,
	"redo":        265,
	"xlog":        265,
	"回放":          265,
}

var gs_team = map[string]interface{}{}

var wcData = map[string]interface{}{
	"database":                 10000,
	"libin":                    6181,
	"Amy Schumer":              4386,
	"Jurassic World":           4055,
	"Charter Communications":   2467,
	"Chick Fil A":              2244,
	"Planet Fitness":           1898,
	"Pitch Perfect":            1484,
	"Express":                  1689,
	"Home":                     1112,
	"Johnny Depp":              985,
	"Lena Dunham":              847,
	"Lewis Hamilton":           582,
	"KXAN":                     555,
	"Mary Ellen Mark":          550,
	"Farrah Abraham":           462,
	"Rita Ora":                 366,
	"Serena Williams":          282,
	"NCAA baseball tournament": 273,
	"Point Break":              265,
}

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic WordCloud example",
		}))

	wc.AddSeries("wordcloud", generateWCData(gs_db_key_words)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func wcRect() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "cardioid shape"}),
	)

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "rect",
				}),
		)
	return wc
}

func wcStar() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "star shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "star",
				}),
		)
	return wc
}

func TestWordCloud(t *testing.T) {
	page := components.NewPage()
	page.AddCharts(
		wcBase(),
		wcRect(),
		wcStar(),
	)

	f, err := os.Create("html/wordcloud.html")
	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		return
	}
}
