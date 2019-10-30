package main

import (
	"github.com/tetsuzawa/go-research/adflib"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math/rand"
)

func init() {
	rand.Seed(1)
}

func main() {
	//creation of data
	n := 512
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, 4)
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			xRow[j] = rand.NormFloat64()
		}
		x[i] = xRow
		v[i] = rand.NormFloat64() * 0.1
		d[i] = 1.5*x[i][0] + 0.8*x[i][1] + 2*x[i][2] + 0.4*x[i][3] + v[i]
	}

	//identification
	f, err := adflib.NewFiltLMS(n, 0.5, "zeros")
	if err != nil {
		log.Fatalln(err)
	}
	y, e, _, err := f.Run(d, x)
	if err != nil {
		log.Fatalln(err)
	}

	// show results
	p, err := plot.New()
	if err != nil {
		log.Fatalln(err)
	}
	//label
	p.Title.Text = "LMS Sample"
	p.X.Label.Text = "sample"
	p.Y.Label.Text = "y"

	p.Add(plotter.NewGrid())

	ptsD := make(plotter.XYs, n)
	ptsY := make(plotter.XYs, n)
	ptsE := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		ptsD[i].X = float64(i)
		ptsD[i].Y = d[i]
		ptsY[i].X = float64(i)
		ptsY[i].Y = y[i]
		ptsE[i].X = float64(i)
		ptsE[i].Y = e[i]
	}

	plotD, err := plotter.NewScatter(ptsD)
	if err != nil {
		log.Fatalln(err)
	}
	plotY, err := plotter.NewScatter(ptsY)
	if err != nil {
		log.Fatalln(err)
	}
	plotE, err := plotter.NewScatter(ptsE)
	if err != nil {
		log.Fatalln(err)
	}
	plotD.GlyphStyle.Color = color.RGBA{R: 87, G: 209, B: 201, A: 1}
	plotY.GlyphStyle.Color = color.RGBA{R: 237, G: 84, B: 133, A: 1}
	plotE.GlyphStyle.Color = color.RGBA{R: 255, G: 232, B: 105, A: 1}

	// \plot
	p.Add(plotD)
	p.Add(plotY)
	p.Add(plotE)

	//label
	p.Legend.Add("Desired", plotD)
	p.Legend.Add("Output", plotY)
	p.Legend.Add("Error", plotE)

	//座標範囲
	//p.X.Min = 0
	//p.X.Max = 10
	//p.Y.Min = 0
	//p.Y.Max = 10
	// plot.pngに保存
	if err := p.Save(20*vg.Centimeter, 20*vg.Centimeter, "FFT.png"); err != nil {
		log.Fatalln(err)
	}

}
