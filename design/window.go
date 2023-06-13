package design

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"decomposition/analysis"
	"decomposition/core"
	"decomposition/maps"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func result_window(myApp fyne.App, s string, k []float64, keys []int64) {
	win2 := myApp.NewWindow("Result")
	win2.SetFixedSize(true)
	win2.Resize(fyne.NewSize(300, 400))
	green := color.NRGBA{R: 133, G: 133, B: 133, A: 255}
	label := canvas.NewText("Результаты исследования", green)
	label.Alignment = fyne.TextAlignCenter
	label.TextSize = 20
	list := strings.Split(s, "\n")
	label1 := canvas.NewText(list[0], green)
	label1.TextSize = 12
	label2 := canvas.NewText(list[1], green)
	label3 := canvas.NewText(list[2], green)
	label2.TextSize = 12
	label3.TextSize = 12
	label4 := canvas.NewText(list[3], green)
	label4.TextSize = 12
	fmt.Println(label4.Text)
	button_dist := widget.NewButtonWithIcon("Построить график распределения", theme.ContentAddIcon(), func() {
		MakeBar(k, keys)
	})
	win2.SetContent(container.NewGridWithRows(7, label, label1, label2, label3, label4, layout.NewSpacer(), button_dist))
	win2.Show()

}
func Show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Decomposition")
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))
	green := color.NRGBA{R: 133, G: 133, B: 133, A: 255}
	label := canvas.NewText("Примеры", green)
	label.TextSize = 16
	label.Alignment = fyne.TextAlignCenter
	combo := widget.NewSelect([]string{"Small", "last_fm", "git"}, func(s string) {})
	combo.SetSelectedIndex(0)
	option := widget.NewSelect([]string{"Все вершины", "Максимальный k-номер"}, func(s string) {})
	option.SetSelectedIndex(1)
	label_time := widget.NewLabel("Время выполнения")
	label_result_time := widget.NewLabel("")
	label_info_res := widget.NewLabel("Полученные результаты")
	label_result := widget.NewLabel("")
	label2 := canvas.NewText("Случайные графы", green)
	label2.TextSize = 16
	label2.Alignment = fyne.TextAlignCenter
	button_a_sample := widget.NewButtonWithIcon("Пуск", theme.ComputerIcon(), func() {
		num := combo.SelectedIndex()
		opt := option.SelectedIndex()

		t, result := core.Samples(num, opt)
		label_result_time.SetText(t)
		label_result.SetText(result)
	})
	first_abstract := container.NewGridWithRows(3, label, container.NewGridWithRows(4, container.NewGridWithColumns(2, combo, option), container.NewGridWithColumns(2, label_time, label_result_time), container.NewGridWithColumns(2, label_info_res, container.NewScroll(label_result))), container.NewPadded(button_a_sample))
	entry_prob := widget.NewEntry()
	entry_prob.SetPlaceHolder("Вероятность")
	entry_nodes_count := widget.NewEntry()
	entry_nodes_count.SetPlaceHolder("Количество вершин")
	option2 := widget.NewSelect([]string{"Все вершины", "Максимальный k-номер"}, func(s string) {})
	option2.SetSelectedIndex(1)
	// label_time2 := widget.NewLabel("Время выполнения")
	label_result_time2 := widget.NewLabel("")
	// label_info_res2 := widget.NewLabel("Полученные результаты")
	label_result2 := widget.NewLabel("")
	button_a_rand := widget.NewButtonWithIcon("Пуск", theme.ComputerIcon(), func() {
		num, _ := strconv.Atoi(entry_nodes_count.Text)
		prob, _ := strconv.ParseFloat(entry_prob.Text, 64)
		opt := option2.SelectedIndex()

		t, result, _, _ := core.Random_test(int64(num), prob, opt)
		label_result_time2.SetText(t)
		label_result2.SetText(result)
	})

	label3 := canvas.NewText("Случайные графы", green)
	label3.TextSize = 20
	label4 := canvas.NewText("Результаты", green)
	label4.TextSize = 15
	label3.Alignment = fyne.TextAlignCenter
	entry_nodes_count2 := widget.NewEntry()
	entry_nodes_count2.SetPlaceHolder("Количество вершин")
	entry_radius := widget.NewEntry()
	entry_radius.SetPlaceHolder("Радиус")

	option3 := widget.NewSelect([]string{"Все вершины", "Максимальный k-номер"}, func(s string) {})
	option3.SetSelectedIndex(1)
	label_time3 := widget.NewLabel("Время выполнения")
	label_result_time3 := widget.NewLabel("")
	cont := container.NewGridWithRows(5, entry_nodes_count2, entry_radius, option3, container.NewGridWithColumns(2, label_time3, label_result_time3))
	var (
		set      map[int64]int64
		res_time string
		size     int64
		results  string
		elements map[int64]*analysis.Element
	)
	button_g := widget.NewButtonWithIcon("Запуск алгоритма декомпозиции", theme.MediaPlayIcon(), func() {
		nodes, _ := strconv.Atoi(entry_nodes_count2.Text)
		strconv.Atoi(entry_radius.Text)
		res_time, _, set, size = core.Random_test(int64(nodes), 0.1, 0)
		label_result_time3.SetText(res_time)
	})
	button_analysis := widget.NewButtonWithIcon("Анализ", theme.SearchIcon(), func() {
		results, elements = analysis.Research(set, size)
	})
	button_result := widget.NewButtonWithIcon("Результаты", theme.InfoIcon(), func() {
		var res []float64
		keys := maps.Keys_ordered(elements)
		for _, e := range elements {
			res = append(res, e.Percent)
		}
		result_window(myApp, results, res, keys)
		// myWindow.ShowAndRun()
	})
	geometric_container := container.NewGridWithRows(3, label3, cont, container.NewAdaptiveGrid(3, container.NewPadded(button_g), container.NewPadded(button_analysis), container.NewPadded(button_result)))
	sec_a_container := container.NewGridWithRows(3, container.NewGridWithColumns(2, entry_nodes_count, entry_prob), option2)
	second_abstract := container.NewGridWithRows(3, label2, sec_a_container, container.NewPadded(button_a_rand))
	abstract := container.NewGridWithColumns(2, first_abstract, second_abstract)
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Abstract", theme.ColorChromaticIcon(), abstract),
		container.NewTabItemWithIcon("Geometric", theme.SearchIcon(), container.NewPadded(geometric_container)),
	)
	// label := widget.NewLabel("Decomposition")
	myWindow.SetContent(tabs)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
