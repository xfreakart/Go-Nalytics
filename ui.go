package main

import (
	"strconv"

	figure "github.com/common-nighthawk/go-figure"
	ui "github.com/gizak/termui"
)

type uiPageList struct {
	data   []string
	title  string
	height int
	Width  int
	Y      int
}

func (l *uiPageList) construct() *ui.List {

	ls := ui.NewList()
	ls.Items = l.data
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabelFg = ui.ColorCyan
	ls.BorderLabel = l.title
	ls.Height = l.height
	ls.Width = l.Width
	ls.Y = l.Y
	return ls

}

//L1 widget: List of most viewed pages by visits + Title
func setL1(data []string) *ui.List {

	list := uiPageList{
		data, "Most Viewed", maxResults + 2, 25, 0,
	} //Height = Max results + two lines of inline (top  bottom)

	uilst := list.construct()

	return uilst
}

type uiBarchart struct {
	borderLabel string
	percentage  int
	width       int
	height      int
	theme       string //Background color
}

func (bar *uiBarchart) construct() *ui.Gauge {

	bch := ui.NewGauge() //New Bar Chart type
	bch.Percent = bar.percentage
	bch.Width = bar.width
	bch.Height = bar.height
	bch.BorderLabel = bar.borderLabel

	switch bar.theme {
	case "blue":
		bch.BarColor = ui.ColorBlue
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "red":
		bch.BarColor = ui.ColorRed
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "black":
		bch.BarColor = ui.ColorBlack
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "green":
		bch.BarColor = ui.ColorGreen
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "cyan":
		bch.BarColor = ui.ColorCyan
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "magenta":
		bch.BarColor = ui.ColorMagenta
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "yellow":
		bch.BarColor = ui.ColorYellow
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	case "white":
		bch.BarColor = ui.ColorWhite
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	default:
		bch.BarColor = ui.ColorDefault
		bch.BorderLabelFg = ui.ColorDefault
		bch.BorderFg = ui.ColorRed
	}

	return bch

}

// R1 widget real time visitors big number
func setR1(n int) *ui.List {

	list := uiPageList{
		bigText(n), "Actual visitors", 8, maxResults + 5, 0,
	}
	uilst := list.construct()

	return uilst
}

func setBarChart(data map[string]int, tt string, color string) *ui.Gauge {
	barchart := uiBarchart{tt, data[tt], 50, 3, color}
	uiChart := barchart.construct()

	return uiChart
}

func bigText(n int) []string {
	figure := figure.NewFigure(strconv.Itoa(n), "basic", true)
	return figure.Slicify()
}

// Big Bar Chart

type uiBigBarChart struct {
	data     []int
	labels   []string
	title    string
	height   int
	width    int
	barwidth int
}

func (bc *uiBigBarChart) construct() *ui.BarChart {

	uibc := ui.NewBarChart()

	uibc.BorderLabel = bc.title
	uibc.Data = bc.data
	uibc.Width = bc.width
	uibc.Height = bc.height
	uibc.BarWidth = bc.barwidth
	uibc.DataLabels = bc.labels
	uibc.TextColor = ui.ColorGreen
	uibc.BarColor = ui.ColorBlue
	uibc.NumColor = ui.ColorYellow

	return uibc
}

func setBigBarChart(labels []string, data []int) *ui.BarChart {

	chart := uiBigBarChart{
		data, labels, "Evolution per minute", 15, 50, 5,
	}

	uiChart := chart.construct()

	return uiChart
}

func setEvolutionChart(labels []string, data []int) *ui.BarChart {

	chart := uiBigBarChart{
		data, labels, "recent evolution", 10, 50, 1,
	}

	uiChart := chart.construct()

	return uiChart

}

func intro() {
	myFigure := figure.NewFigure("Go-Nalytics", "doom", true)
	myFigure.Print()

}
