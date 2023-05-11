package main

import (
	"flag"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Point struct {
	Pos []float64
}

func PointsInit(points *[]Point) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(*points); i++ {
		//生成的点要和前面的点坐标不一样
		if i == 0 {
			for j := 0; j < len((*points)[i].Pos); j++ {
				(*points)[i].Pos[j] = rand.Float64()
			}
		} else {
			//从第而个开始，后面的都要和前面的所有不一样(一样：坐标值完全相同)
			isFind := false
			for !isFind {
				//随机生成一个点
				point := Point{}
				point.Pos = make([]float64, len((*points)[i].Pos))
				for j := 0; j < len((*points)[i].Pos); j++ {
					point.Pos[j] = rand.Float64()
				}
				isSame := false
				//对比和前面所有的点
				for m := 0; m < i; m++ {
					for n := 0; n < len((*points)[i].Pos); n++ {
						if point.Pos[n] == (*points)[m].Pos[n] {
							isSame = true
						}
					}
				}
				if !isSame {
					(*points)[i] = point
					isFind = true
				}

			}
		}
	}
}

func main() {
	var d int
	var c int
	flag.IntVar(&d, "d", 2, "维数，默认2")
	flag.IntVar(&c, "c", 50, "实验次数，默认50")

	flag.Parse()

	//生成row=d col=100的数组，表示点
	points := make([]Point, 100)
	for i := 0; i < len(points); i++ {
		points[i].Pos = make([]float64, d)
	}

	PointsInit(&points)

	result := make([]float64, c)
	var wg sync.WaitGroup
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(val *float64) {
			defer wg.Done()
			PointsInit(&points)
			distance := 1.0
			for j := 1; j < len(points); j++ {
				dsqrt := 0.0
				for m := 0; m < len(points[j].Pos); m++ {
					dsqrt += math.Pow(math.Abs(points[0].Pos[m]-points[j].Pos[m]), 2)
				}
				d := math.Sqrt(dsqrt)
				if d < distance {
					distance = d
				}
			}
			*val = distance
		}(&result[i])
	}
	wg.Wait()
	//统计
	df := dataframe.New(series.New(result, series.Float, "distance_min"))
	summary := df.Describe()
	fmt.Println(summary)

}
