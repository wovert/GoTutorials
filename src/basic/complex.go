package main

import ("fmt"
	"math/cmplx"
)

func main() {
	var x complex64 = 6 + 2i
	fmt.Print(x*x)

	// 3 + 4i, x3, y4, 斜线长度5
	// 3^2+4^2 的平方根是5
	// i^2=-1, i^3=-i, i^4=1, ...

	s := 3 + 4i;
	fmt.Println(cmplx.Abs(s))
}