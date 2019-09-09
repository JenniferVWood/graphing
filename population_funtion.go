package main

import (
	//"fmt"
	"gonum.org/v1/plot/plotter"
	"math"
)

func nextX(x float64, r float64) float64 {
	// xNext = r * x * (x -1)
	nextX := r * x * (1 - x)
	//fmt.Printf("%v * %v * (1 - %v) = %v\n", r, x, x, nextX )
	return nextX
}

func skipNPoints(x float64, r float64, n int) float64 {
	for i := 0; i < n; i++ {
		x = nextX(x, r)
	}
	return x
}

func seriesXforOneR(x float64, r float64, n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		x = nextX(x, r)
		pts[i].X = x
	}
	return pts
}

func xsForOneR(x float64, r float64, n int, skip int) plotter.XYs {
	//fmt.Printf("x: %v, r: %v, n: %v, skip: %v\n", x, r, n, skip)
	x = skipNPoints(x, r, skip)

	pts := make(plotter.XYs, n)

	for i := range pts {
		//fmt.Printf("x: %v\n", x)
		x = nextX(x, r)
		if !math.IsInf(x, 0) && math.Abs(x) <= 10000 {
			pts[i].Y = x
			pts[i].X = r
		} else {
			if x < 0 {
				pts[i].Y = -10000.0
			} else {
				pts[i].Y = 10000.0
			}
			pts[i].X = r
		}
	}
	return pts
}
