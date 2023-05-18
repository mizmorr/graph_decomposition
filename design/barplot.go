package design

import (
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func MakeBar(res []float64) {
	xaxiss := []string{}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros, Width: "1000px", Height: "800px"}),
		charts.WithTitleOpts(opts.Title{Title: "Распределение значений номеров k-ядер заданного геометрического графа"}),
	)
	values := []opts.BarData{}
	for _, k := range res {
		values = append(values, opts.BarData{Value: k})
	}
	bar.SetXAxis(xaxiss).AddSeries("Значения k-ядер", values).SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "top",
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
