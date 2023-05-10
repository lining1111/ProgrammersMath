package main

import (
	"flag"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var c int
	flag.IntVar(&c, "c", 10000, "实验次数，默认10000")

	flag.Parse()

	conditions := mat.NewDense(5, c, nil)
	for i := 0; i < 5; i++ {
		for j := 0; j < c; j++ {
			conditions.Set(i, j, 0.0)
		}
	}

	//同时开启切蛋糕
	var wg sync.WaitGroup
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			//第一个人得到的
			conditions.Set(0, row, rand.Float64())
			index := conditions.At(0, row)
			for j := 1; j < 4; j++ {
				conditions.Set(j, row, rand.Float64()*(1-index))
				index += conditions.At(j, row)
			}
			conditions.Set(4, row, 1-index)
		}(i)
	}

	wg.Wait()

	fmt.Println(conditions.At(0, 0), conditions.At(1, 0), conditions.At(2, 0), conditions.At(3, 0), conditions.At(4, 0),
		(conditions.At(0, 0) + conditions.At(1, 0) + conditions.At(2, 0) + conditions.At(3, 0) + conditions.At(4, 0)))

	//统计
	df := dataframe.New(series.New(conditions.RawRowView(0), series.Float, "A"),
		series.New(conditions.RawRowView(1), series.Float, "B"),
		series.New(conditions.RawRowView(2), series.Float, "C"),
		series.New(conditions.RawRowView(3), series.Float, "D"),
		series.New(conditions.RawRowView(4), series.Float, "E"))
	summary := df.Describe()
	fmt.Println(summary)

}
