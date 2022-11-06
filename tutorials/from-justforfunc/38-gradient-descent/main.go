package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

var iterations int

func main() {
	flag.IntVar(&iterations, "n", 1000, "number of iterations")
	flag.Parse()

	xys, err := readData("data.txt")
	if err != nil {
		log.Printf("could not read data.txt: %v", err)
	}
	// _ = xys
	// for _, xy := range xys {
	// 	fmt.Println("--", xy.x, xy.y)
	// }

	err = plotData("out.png", xys)
	if err != nil {
		log.Printf("could not plot data: %v", err)
	}
}

// type xy struct{ x, y float64 }

// func readData(path string) ([]xy, error) {
func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// var xys []xy // slice of xy
	var xys plotter.XYs // change to plotter struct
	s := bufio.NewScanner(f)
	for s.Scan() {
		//fmt.Println(s.Text())
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
		}
		// xys = append(xys, xy{x, y})
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not reead: %v", err)
	}
	return xys, nil
}

// func plotData(path string, xys []xy) error {
func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: ", err)
	}

	p := plot.New()
	// if err != nil {
	// 	return fmt.Errorf("could not create plot: %v", err)
	// }

	// pxys := make(plotter.XYs, len(xys))

	// for i, xy := range xys {
	// 	pxys[i].X = xy.x
	// 	pxys[i].Y = xy.y
	// }
	// create scatter with all data points
	s, err := plotter.NewScatter(xys)

	// s, err := plotter.NewScatter(plotter.XYs{
	// 	{0, 0}, {0.5, 0.5}, {1, 1},
	// })
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}

	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	// create fake linear regression result
	// var x, c float64
	// x = 1.2
	// c = 3
	x, c := linearRegression(xys, 0.01)

	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c}, // from (0,1) to (20,21) points
	})
	p.Add(l)

	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %v", err)
	}
	return nil
}

// x * m + c
// The equation y = mx + c is called the slope-intercept
// form of the equation of the line.
// It requires the slope value 'm' and
// the y-intercept c of the line.
// alpha = learning rate
func linearRegression(xys plotter.XYs, alpha float64) (m, c float64) {
	// m = 0
	// c = 0
	// cost(0.00, 0.00) = 64.15
	// const (
	// 	min   = -100.0
	// 	max   = 100.0
	// 	delta = 0.1 // leaning rate
	// )

	// minCost := math.MaxFloat64
	// for im := min; im < max; im += delta {
	// 	for ic := min; ic < max; ic += delta {
	// 		cost := computeCost(xys, im, ic)
	// 		if cost < minCost {
	// 			minCost = cost
	// 			m, c = im, ic
	// 			dm, dc := computeGradient(xys, m, c)
	// 			fmt.Printf("grad(%.2f, %.2f) = (%.2f, %.2f)\n", m, c, dm, dc)
	// 		}
	// 	}
	// }

	for i := 0; i < iterations; i++ {
		// for i := 0; i < 100000; i++ {
		// for i := 0; i < 30000; i++ {  // same result as 10000
		dm, dc := computeGradient(xys, m, c)
		m += -dm * alpha // move to opposition direction
		c += -dc * alpha
		// fmt.Printf("cost(%.2f, %.2f) = (%.2f, %.2f)\n", m, c, dm, dc)
		fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(xys, m, c))

	}
	fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(xys, m, c))
	return m, c // explicit return better than implicit return
}

// show how good or bad it performs or does
// cost show how well we are doing
// cost = 0 means we can match every single point
// goal is trying to find the points with lowest cost
func computeCost(xys plotter.XYs, m, c float64) float64 {
	// cost = 1 / N * sum((y - (m*x+c))^2)
	s := 0.0 // call it sum
	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c) // distance
		s += d * d               // above fomula
	}
	return s / float64(len(xys))

}

// which directions we are doing good or bad
// based on cost information, move toward to solution 
func computeGradient(xys plotter.XYs, m, c float64) (dm, dc float64) {
	// cost = 1/N * sum((y - (m*x+c))^2)
	// cost/dm = 2/N * sum(-x * (y - (m*x+c)))
	// cost/dc = 2/N * sum(-(y - (m*x+c)))
	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		dm += -xy.X * d
		dc += -d
	}
	n := float64(len(xys))
	return 2 / n * dm, 2 / n * dc
}
