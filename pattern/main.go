package main

import (
	"flag"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Str    string
	length int
}

func Pattern(m string) Result {
	str := ""
	rand.Seed(time.Now().UnixNano())
	isContain := false
	for !isContain {

		if rand.Float64() < 0.5 {
			str += "0"
		} else {
			str += "1"
		}

		if strings.Contains(str, m) {
			isContain = true
		}
	}

	return Result{str, len(str)}
}

func main() {
	var m string
	var c int
	flag.StringVar(&m, "m", "01", "匹配的模式，默认01")
	flag.IntVar(&c, "c", 20, "实验次数，默认20")

	flag.Parse()

	res := make([]Result, c)

	var wg sync.WaitGroup
	for i := 0; i < len(res); i++ {
		wg.Add(1)
		go func(r *Result) {
			defer wg.Done()
			*r = Pattern(m)
		}(&res[i])
	}
	wg.Wait()

	fmt.Printf("模式 %v 次数 %v\n", m, c)
	//计算统计结果
	count := make([]int, c)
	for k, v := range res {
		count[k] = v.length
	}

	df := dataframe.New(series.New(count, series.Int, "次数"))
	summary := df.Describe()
	fmt.Println(summary)

}
