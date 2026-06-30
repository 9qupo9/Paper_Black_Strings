package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	G         = 6.67430e-11 // m^3 kg^-1 s^-2
	C         = 299792458.0 // m/s
	SolarMass = 1.989e30    // kg
	L_AdS     = 1e-4        // m
)

func main() {
	p := plot.New()

	p.Title.Text = "Black String Instability in AdS Space (Gregory-Laflamme)"
	p.X.Label.Text = "Black Hole Mass [kg] (Log Scale)"
	p.Y.Label.Text = "Decay Time [s] (Log Scale)"

	p.X.Scale = plot.LogScale{}
	p.Y.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{}
	p.Y.Tick.Marker = plot.LogTicks{}

	ptsDecay := make(plotter.XYs, 0)
	ptsHawking := make(plotter.XYs, 0)

	// Mass range from 10^5 to 10^30 kg
	for i := 0; i < 200; i++ {
		mass := math.Pow(10, 5.0+float64(i)*(25.0/199.0))
		rs := 2 * G * mass / (C * C)

		// Hawking evaporation time (approximate 4D)
		tHawking := (5120 * math.Pi * G * G * mass * mass * mass) / (1.054e-34 * math.Pow(C, 4))
		ptsHawking = append(ptsHawking, plotter.XY{X: mass, Y: tHawking})

		// Gregory-Laflamme Decay time in AdS
		if rs < L_AdS {
			tDecay := math.Pow(rs, 3) / (L_AdS * L_AdS * C)
			ptsDecay = append(ptsDecay, plotter.XY{X: mass, Y: tDecay})
		}
	}

	err := plotutil.AddLinePoints(p,
		"Hawking Evaporation (4D)", ptsHawking,
		"AdS Decay Time (5D)", ptsDecay,
	)
	if err != nil {
		log.Fatal(err)
	}

	p.X.Min = 1e5
	p.X.Max = 1e30
	p.Y.Min = 1e-45
	p.Y.Max = 1e80

	filename := "bh_decay_time.pdf"
	if err := p.Save(8*vg.Inch, 6*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Plot saved as %s\n", filename)
}
