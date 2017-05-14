package main

import (
	"time"

	ui "github.com/gizak/termui"
)

func main() {

	intro()

	//By default You can call google Analytics Api 500K per day. GoLytics does two query's each time.
	// So we can do 250K visualizations per day.
	// That gives us 2,89 visualizations per day. If you add more widgets wiith new query's keep that in mind.
		var refreshRate uint64
		refreshRate = 2
	
	// String array of most viwed pages titles
	rtMostViewedTitles := getRtMostViewedPages()

	// Map of Traffic Type and its percentage
	var rtTrafficType map[string]int

	// Arrays to store big chart data(visitors) and labels (time)
	var bcData = make([]int, 90)
	var bcLabels = make([]string, 90)

	var bcLabelsMin = make([]string, 90)
	var bcDataMin = make([]int, 90)

	// Integer total active user counter
	rtTotalActiveUsers := 0

	updateRtMostViewedByTrafficType(&rtTotalActiveUsers, &rtTrafficType, &bcData, &bcLabels)

	err := ui.Init()intro
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	// We are goiung to name the different widgets by their side and position
	// l1 = LEFT 1
	l1 := setL1(rtMostViewedTitles)
	l1.ItemFgColor = ui.ColorWhite

	r1 := setR1(rtTotalActiveUsers)

	r2Bar := setBarChart(rtTrafficType, "CUSTOM", "blue")
	r3Bar := setBarChart(rtTrafficType, "DIRECT", "red")
	r4Bar := setBarChart(rtTrafficType, "ORGANIC", "yellow")
	r5Bar := setBarChart(rtTrafficType, "REFERRAL", "magenta")
	r6Bar := setBarChart(rtTrafficType, "SOCIAL", "green")

	// Big char
	bigChart := setBigBarChart(bcLabels, bcData)
	// Evolution chart small
	evoChart := setEvolutionChart(bcLabels, bcData)
	evoChart.BarColor = ui.ColorCyan

	par0 := ui.NewPar("Press Q to exit.")
	par0.Height = 1
	par0.Width = 20
	par0.Y = 1
	par0.Border = false

	// build
	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(8, 0, par0),
			ui.NewCol(4, 0, par0)),
		ui.NewRow(ui.NewCol(12, 0, evoChart)),
		ui.NewRow(
			ui.NewCol(6, 0, l1),
			ui.NewCol(6, 0, r1, r2Bar, r3Bar, r4Bar, r5Bar, r6Bar)),

		ui.NewRow(ui.NewCol(12, 0, bigChart)))
	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	// If Q key pressed, exit program
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	//For croneedbcLabelsMin tasks, every second we print the interface
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)

		// Each 4 seconds we update values by quering GA.
		if t.Count%refreshRate2  == 0 {
			rtMostViewedTitles = getRtMostViewedPages()
			updateRtMostViewedByTrafficType(&rtTotalActiveUsers, &rtTrafficType, &bcData, &bcLabels)
		}
		if t.Count%59 == 0 {
			t := time.Now().Format("03:04")
			bcDataMin = append([]int{rtTotalActiveUsers}, bcDataMin...)
			bcDataMin = bcDataMin[:len(bcDataMin)-1]
			bcLabelsMin = append([]string{t}, bcLabelsMin...)
			bcLabelsMin = bcLabelsMin[:len(bcLabelsMin)-1]
		}

		//Each second we print the UI again
		if t.Count%1 == 0 {
			l1.Items = rtMostViewedTitles
			r1.Items = bigText(rtTotalActiveUsers)
			r2Bar.Percent = rtTrafficType["CUSTOM"]
			r3Bar.Percent = rtTrafficType["DIRECT"]
			r4Bar.Percent = rtTrafficType["ORGANIC"]
			r5Bar.Percent = rtTrafficType["REFERRAL"]
			r6Bar.Percent = rtTrafficType["SOCIAL"]
			evoChart.Data = bcData
			bigChart.Data = bcDataMin
			bigChart.DataLabels = bcLabelsMin
			ui.Render(ui.Body)
		}
	})
	ui.Loop()

	return

}
