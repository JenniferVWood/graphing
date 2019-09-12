package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"time"
)

func main() {
	//scatterData := seriesXforOneR(initialX, r, n)
	//scatterData := xsForOneR(initialX, r, n, skip)

	// now plot multiple r's
	n := 1000
	initialX := 0.02
	skip := 200
	initialR := -2.51
	endR := 4.01
	step := 0.001

	numPoints := 0.0
	if initialR < 0 {
		numPoints = endR + -initialR
	} else {
		numPoints = endR - initialR
	}
	numPoints = (numPoints + float64(skip)) * float64(n) / step
	fmt.Printf("Calculating %d points...", int64(numPoints))
	start := time.Now()

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
	if err := p.Save(24*vg.Inch, 24*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

	fmt.Printf("in: %v", time.Since(start))
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
