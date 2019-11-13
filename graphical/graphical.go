package graphical

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg"
)

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Relabeling tick marks example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	plotutil.AddLinePoints(
		p,
		"First", plotter.XYs { 
			{X: 5, Y: 15}, 
			{X: 6, Y: 1}, 
			{X: 7, Y: -5}, 
			{X: 8, Y: 18}, 
			{X: 10, Y: 2}, 
		},
		"Second", plotter.XYs { 
			{X: 2, Y: 13}, 
			{X: 3, Y: 13}, 
			{X: 4, Y: -2}, 
			{X: 5, Y: 18}, 
			{X: 8, Y: 4}, 
		},
	)

	draw.New(vg.CanvasSizer{})
	p.Draw()
}