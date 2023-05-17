package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	A := mat.NewDense(3, 3, []float64{2, 6, 4, 5, 7, 9, 12, 11, 5})
	fmt.Printf("A\n%v\n", mat.Formatted(A))
	lu := mat.LU{}
	lu.Factorize(A)
	L := &mat.TriDense{}
	U := &mat.TriDense{}
	lu.LTo(L)
	lu.UTo(U)
	P := &mat.Dense{}
	pivot := lu.Pivot(nil)
	P.Permutation(3, pivot)
	fmt.Printf("L\n%v\n", mat.Formatted(L))
	fmt.Printf("U\n%v\n", mat.Formatted(U))
	fmt.Printf("P\n%v\n", mat.Formatted(P))
	A1 := &mat.Dense{}
	A1.Product(P, L, U)
	fmt.Printf("A1\n%v\n", mat.Formatted(A1))
	Lt := &mat.TriDense{}
	Ut := &mat.TriDense{}
	err := Lt.InverseTri(L)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Lt\n%v\n", mat.Formatted(Lt))
	Il := &mat.Dense{}
	Il.Mul(L, Lt)
	fmt.Printf("Il\n%v\n", mat.Formatted(Il))
	err = Ut.InverseTri(U)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ut\n%v\n", mat.Formatted(Ut))
	Iu := &mat.Dense{}
	Iu.Mul(U, Ut)
	fmt.Printf("Iu\n%v\n", mat.Formatted(Iu))

	Pt := &mat.Dense{}
	Pt.Inverse(P)
	fmt.Printf("Pt\n%v\n", mat.Formatted(Pt))
	At := &mat.Dense{}
	At.Mul(Lt, Pt)
	At.Mul(Ut, At)
	fmt.Printf("At\n%v\n", mat.Formatted(At))
	At1 := &mat.Dense{}
	err = At1.Inverse(A)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("At1\n%v\n", mat.Formatted(At1))

	I := &mat.Dense{}
	I.Mul(A, At)
	fmt.Printf("I\n%v\n", mat.Formatted(I))

}
