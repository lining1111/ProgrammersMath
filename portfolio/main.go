package main

import (
	"flag"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var p float64
	var c int
	flag.Float64Var(&p, "p", 0.7, "出事故的概率，默认0.7")
	flag.IntVar(&c, "c", 100, "实验次数，默认100")

	flag.Parse()
	pA := 0.7
	pB := 0.3
	//假设初始资金为1.0,计算100场的收益 0.7和0.3为随机变量的独立概率
	conditions := make([]float64, c)

	for k, _ := range conditions {
		conditions[k] = 1.0
	}
	//开启100场同样的赌局，看最后的收益
	var wg sync.WaitGroup
	for k, _ := range conditions {
		wg.Add(1)
		go func(val *float64) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			pa := rand.Float64()
			pb := rand.Float64()
			if pa <= pA {
				*val = *val * p * 2.0
			}

			if pb <= pB {
				*val += *val * (1 - p) * 2.0
			}
		}(&conditions[k])
	}
	wg.Wait()

	for k, v := range conditions {
		fmt.Printf("%v ", v)
		if k%10 == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")

	//计算统计结果
	df := dataframe.New(series.New(conditions, series.Float, "val"))
	summary := df.Describe()
	fmt.Println(summary)
}
