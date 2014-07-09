package main

import (
	"bufio"
	"fmt"
	"github.com/gonum/blas/goblas"
	mat "github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
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

func draw_henon(a, b float64) (x, y float64) {
	x, y = rand.Float64(), rand.Float64()
	for i := 0; i < 1000*10; i++ {
		x, y = henon(a, b, x, y)
	}
	return x, y
}

func Jcb(a, b, x float64) *mat.Dense {
	jcb := mat.NewDense(2, 2, []float64{-2 * a * x, 1, b, 0})
	return jcb
}

func MatrixLength(m mat.Matrix) float64 {
	tmp := &mat.Dense{}
	tmp.MulElem(m, m)
	return math.Sqrt(tmp.Sum())
}

func main() {
	xf, _ := os.Create("henon.dat")
	defer xf.Close()
	xfb := bufio.NewWriter(xf)

	dxf, _ := os.Create("henon_ly.dat")
	defer dxf.Close()
	dxfb := bufio.NewWriter(dxf)

	// for initial condition
	a, b := 0.0, 0.3
	scale := 10000
	steps := 2 * scale

	// for the Lyapunov

	eb := 1.0 / math.Sqrt(2)
	e0 := mat.NewDense(2, 1, []float64{-eb, eb})
	f0 := mat.NewDense(2, 1, []float64{eb, eb})

	for idx := 0; idx < steps; idx++ {
		a = a + 1.0/float64(scale)
		x, y := rand.Float64(), rand.Float64()
		e0_sum := 0.0
		iterates := 1000
		for i := 0; i < iterates; i++ {
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
		}

		dxfb.WriteString(fmt.Sprintf("%f %f\n", a, e0_sum/float64(iterates)))
		xfb.WriteString(fmt.Sprintf("%f %f\n", a, x))
	}
	dxfb.Flush()
	xfb.Flush()
}
