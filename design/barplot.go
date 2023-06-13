package design

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func MakeBar(res []float64, keys []int64) {
	xaxiss := []string{}
	for _, key := range keys {
		xaxiss = append(xaxiss, fmt.Sprint(key))
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, Width: "1100px", Height: "700px"}),
		charts.WithTitleOpts(opts.Title{Title: "Распределение значений номеров k-ядер"}),
	)
	values := []opts.BarData{}
	for _, k := range res {
		values = append(values, opts.BarData{Value: k})
	}
	bar.SetXAxis(xaxiss).AddSeries("Значения k-ядер", values, charts.WithMarkPointNameCoordItemOpts(
		opts.MarkPointNameCoordItem{
			Value:      fmt.Sprint(keys[0]),
			Coordinate: []interface{}{"", res[0]},
		})).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "insideTopRight",
		}))

	bar.SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
		BarGap: "0%",
	}),

		charts.WithLabelOpts(opts.Label{Show: true, Position: "top"}))

	f, _ := os.Create("test/bar.html")
	bar.Render(f)
	cmd := exec.Command("xdg-open", "test/bar.html")
	cmd.Run()
}
