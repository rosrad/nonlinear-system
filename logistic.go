package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

func logistic(a, x float64) float64 {
	return a * x * (1 - x)
}

func df_logistic(a, x float64) float64 {
	return math.Log(math.Abs(a * (1 - 2*x)))
}

func PhaseGraph(a float64) {
	para_f, _ := os.Create("logistic_phase.dat")
	defer para_f.Close()
	para_fb := bufio.NewWriter(para_f)
	n := 200 * 200
	x := 0.1
	for i := 0; i < n; i++ {
		x = logistic(a, x)
		para_fb.WriteString(fmt.Sprintf("%d %f\n", i, x))
	}
	para_fb.Flush()
}

func LyapunovExponents() {
	xf, _ := os.Create("logistic.dat")
	dxf, _ := os.Create("logistic_ly.dat")
	defer xf.Close()
	defer dxf.Close()
	xfb := bufio.NewWriter(xf)
	dxfb := bufio.NewWriter(dxf)

	zdxf, _ := os.Create("logistic_ly_zero.dat")
	defer zdxf.Close()
	zdxfb := bufio.NewWriter(zdxf)

	ze := 0.01
	x0 := 0.1
	var x, a float64
	n := 200

	for j := 0; j < 40000; j++ {
		a = a + 0.001
		x = x0
		sum := 0.0
		for i := 0; i < n; i++ {
			x = logistic(a, x)
		}
		for i := 0; i < n; i++ {
			sum += df_logistic(a, x)
			// fmt.Println("sum is ", sum)
			x = logistic(a, x)
			xfb.WriteString(fmt.Sprintf("%f %f\n", a, x))
		}
		ly := sum / float64(n*2)
		dxfb.WriteString(fmt.Sprintf("%f %f\n", a, ly))

		if math.Abs(ly) < ze {
			zdxfb.WriteString(fmt.Sprintf("%f %f\n", a, ly))
		}

	}
	zdxfb.Flush()
	xfb.Flush()
	dxfb.Flush()
}

func PhaseMain() {
	var a float64
	flag.Float64Var(&a, "a", 0.1, "para for logistic map")
	flag.Parse()
	fmt.Println("Para a:", a)
	PhaseGraph(a)
}

func main() {
	// LyapunovExponents()
	PhaseMain()
}
