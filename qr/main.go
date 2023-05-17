package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	//微分方程 du1/dt=-u1+2u2; du2/dt=u1-2u2;u[0]={1,0}
	//求解u(t)
	//矩阵[-1,2,1,-2],y[0]=[1,0];求y[t]
	A := mat.NewDense(2, 2, []float64{-1, 2, 1, -2})
	fmt.Printf("A\n%v\n", mat.Formatted(A))
	qr := &mat.QR{}
	qr.Factorize(A)
	q := &mat.Dense{}
	r := &mat.Dense{}
	qr.QTo(q)
	qr.RTo(r)
	fmt.Printf("Q\n%v\n", mat.Formatted(q))
	fmt.Printf("R\n%v\n", mat.Formatted(r))

	got := &mat.Dense{}
	got.Mul(q, r)
	if !mat.EqualApprox(got, A, 1e-12) {
		fmt.Errorf("QR does not equal original matrix. \nWant: %v\nGot: %v", A, got)
	}

}
