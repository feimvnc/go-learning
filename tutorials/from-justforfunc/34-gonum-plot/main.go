package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
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
		fmt.Println(s.Text())
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
	var x, c float64
	x = 1
	c = 1
	l, err := plotter.NewLine(plotter.XYs{
		{0, c}, {20, 20*x + c}, // from (0,1) to (20,21) points
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
