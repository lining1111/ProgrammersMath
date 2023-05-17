package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	//2x+3y-z=1
	//x+y+z=4
	//4x+5y+6z=10
	//求解方程的解，这里可以看作是[x,y,z]的三维的点，通过方阵
	//2 3 -1
	//1 1 1
	//4 5 6
	//得到1 4 10的点，求原点
	y := mat.NewDense(3, 1, []float64{1, 4, 10})
	A := mat.NewDense(3, 3, []float64{2, 3, -1, 1, 1, 1, 4, 5, 6})
	fmt.Println(mat.Det(A))
	fmt.Println(A.Trace())
	//求A的逆矩阵
	At := &mat.Dense{}
	err := At.Inverse(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	//求自变量
	x := mat.NewDense(3, 1, []float64{0, 0, 0})
	x.Mul(At, y)
	fmt.Println(mat.Formatted(x))
	//将自变量代入矩阵，乘法，求出因变量，做验证
	y1 := mat.NewDense(3, 1, []float64{0, 0, 0})
	y1.Mul(A, x)
	fmt.Println(mat.Formatted(y1))

}
