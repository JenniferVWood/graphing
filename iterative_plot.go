package main

import (
	"gonum.org/v1/plot/vg/draw"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	//scatterData := seriesXforOneR(initialX, r, n)
	//scatterData := xsForOneR(initialX, r, n, skip)

	// now plot multiple r's
	n := 1000
	initialX := 0.02
	skip := 400
	initialR := 2.7
	endR := 3.9
	step := 0.001

	// Create a new plot, set its title and
	// axis labels.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Bifurcation Diagram"
	p.X.Label.Text = "r"
	p.Y.Label.Text = "all Xs"

	for r := initialR; r <= endR; r += step {
		go addPointsToPlotter(xsForOneR(initialX, r, n, skip), p)
		//addPointsToPlotter( xsForOneR(initialX, r, n, skip), p)
	}

	// Save the plot to a PNG file.
	if err := p.Save(16*vg.Inch, 16*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func addPointsToPlotter(scatterData plotter.XYs, p *plot.Plot) {
	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(0.5)
	s.GlyphStyle.Shape = draw.CircleGlyph{}

	// Add the plotters to the plot, with a legend
	// entry for each
	p.Add(s)
}
