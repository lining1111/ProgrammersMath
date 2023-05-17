package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		-1, 2,
		1, -2,
	})
	fmt.Printf("A = %v\n\n", mat.Formatted(a, mat.Prefix("    ")))

	var eig mat.Eigen
	ok := eig.Factorize(a, mat.EigenRight)
	if !ok {
		log.Fatal("Eigendecomposition failed")
	}
	fmt.Printf("Eigenvalues of A:\n%v\n", eig.Values(nil))
	v := &mat.CDense{}
	eig.VectorsTo(v)
	fmt.Printf("vector:\n%v\n", v)

}
