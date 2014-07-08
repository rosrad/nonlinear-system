package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// func logistic(a, x float64) float64 {
// 	return a * x * (1 - x)
// }

// func df_logistic(a, x float64) float64 {
// 	return math.Log(math.Abs(a * (1 - 2*x)))
// }

func logistic(a, b, x, y float64) (xn, yn float64) {
	xn = 1 - a*x*x + y
	yn = b * x
	return xn, yn
}

// func df_logistic(a, b, x, y float64) (xl, yl float64) {
// 	xl = math.Log(math.Abs(-2 * a * x))
// }
func main() {

	xf, _ := os.Create("logistic.dat")
	dxf, _ := os.Create("logistic_dx.dat")
	defer xf.Close()
	defer dxf.Close()
	xfb := bufio.NewWriter(xf)
	dxfb := bufio.NewWriter(dxf)

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
		dxfb.WriteString(fmt.Sprintf("%f %f\n", a, sum/float64(n*2)))
	}
	xfb.Flush()
	dxfb.Flush()
}
