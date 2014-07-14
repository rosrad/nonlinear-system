package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gonum/blas/goblas"
	mat "github.com/gonum/matrix/mat64"
	"math"
	"os"
)

func init() {
	mat.Register(goblas.Blas{})
}

type fm struct {
	mat.Matrix
	margin int
}

func (m fm) Format(fs fmt.State, c rune) {
	if c == 'v' && fs.Flag('#') {
		fmt.Fprintf(fs, "%#v", m.Matrix)
		return
	}
	mat.Format(m.Matrix, m.margin, '.', fs, c)
}

func henon(a, b, x, y float64) (xn, yn float64) {
	xn = 1 - a*x*x + y
	yn = b * x
	return xn, yn
}

func Jcb(a, b, x float64) *mat.Dense {
	jcb := mat.NewDense(2, 2, []float64{-2 * a * x, 1, b, 0})
	return jcb
}

func MatrixLength(m mat.Matrix) float64 {
	tmp := &mat.Dense{}
	tmp.MulElem(m, m)
	res := math.Sqrt(tmp.Sum())
	return res
}

func PhaseGraph(a, b float64) {
	para_f, _ := os.Create("henon_phase.dat")
	defer para_f.Close()
	para_fb := bufio.NewWriter(para_f)

	num := 10000 * 20 * 10
	x, y := 0.1, 0.3

	for i := 0; i < 1000; i++ {
		x, y = henon(a, b, x, y)
	}
	for i := 0; i < num; i++ {
		x, y = henon(a, b, x, y)
		para_fb.WriteString(fmt.Sprintf("%f %f\n", x, y))
	}
	para_fb.Flush()
}

func LyapunovExponents() {
	xf, _ := os.Create("henon.dat")
	defer xf.Close()
	xfb := bufio.NewWriter(xf)

	dxf, _ := os.Create("henon_ly.dat")
	defer dxf.Close()
	dxfb := bufio.NewWriter(dxf)

	zdxf, _ := os.Create("henon_ly_zero.dat")
	defer zdxf.Close()
	zdxfb := bufio.NewWriter(zdxf)

	// for initial condition
	a, b := 0.0, 0.3
	scale := 10000

	// for the Lyapunov

	eb := 1.0 / math.Sqrt(2)
	e0 := mat.NewDense(2, 1, []float64{-eb, eb})
	f0 := mat.NewDense(2, 1, []float64{eb, eb})

	ze := 0.005
	x0, y0 := 0.1, 0.3
	fmt.Printf("x0: %f, y0: %f \n", x0, y0)
	num := 200

	for idx := 0; a < 1.5; idx++ {
		a = a + 1.0/float64(scale)
		e0_sum := 0.0
		x, y := x0, y0

		for i := 0; i < num; i++ {
			x, y = henon(a, b, x, y)
		}
		for i := 0; i < num; i++ {
			// for the ly
			// Step 2
			tmp := &mat.Dense{}
			jcb := Jcb(a, b, x)
			// calculate the error
			e1 := &mat.Dense{}
			e1.Mul(jcb, e0)
			f1 := &mat.Dense{}
			f1.Mul(jcb, f0)

			// Step 3
			e0_sum += math.Log(MatrixLength(e1))

			// Step 4
			tmp.Reset()
			tmp.TCopy(f1)
			tmp.Mul(tmp, e1)
			fe := tmp.Sum()

			tmp.Reset()
			tmp.MulElem(e1, e1)
			ee := tmp.Sum()

			tmp.Reset()
			tmp.Scale(fe/ee, e1)

			f1.Sub(f1, tmp)
			// normorlize
			f0.Scale(1.0/MatrixLength(f1), f1)
			e0.Scale(1.0/MatrixLength(e1), e1)
			x, y = henon(a, b, x, y)
			xfb.WriteString(fmt.Sprintf("%f %f\n", a, x))
		}
		ly := e0_sum / float64(num)
		dxfb.WriteString(fmt.Sprintf("%f %f\n", a, ly))

		if math.Abs(ly) < ze {
			zdxfb.WriteString(fmt.Sprintf("%f %f\n", a, ly))
		}
	}
	zdxfb.Flush()
	dxfb.Flush()
	xfb.Flush()
}

func PhaseMain() {
	var a, b float64
	flag.Float64Var(&a, "a", 0.1, "para for logistic map")
	flag.Float64Var(&b, "b", 0.3, "para for logistic map")
	flag.Parse()
	fmt.Printf("Para a: %f, b: %f\n", a, b)
	PhaseGraph(a, b)
}

func main() {
	// LyapunovExponents()
	PhaseMain()
}
