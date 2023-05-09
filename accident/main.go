package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var p float64
	var c int
	flag.Float64Var(&p, "p", 0.5, "出事故的概率")
	flag.IntVar(&c, "c", 1000, "实验次数，默认1000")

	flag.Parse()
	conds := []string{"o", "."}
	conditions := make([]string, c)

	//随机生产时间，即 o 或者.
	var wg sync.WaitGroup
	for k, _ := range conditions {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			if rand.Float64() >= p {
				conditions[index] = conds[1]
			} else {
				conditions[index] = conds[0]
			}

		}(k)
	}
	wg.Wait()
	//打印情形
	for k, v := range conditions {
		fmt.Printf("%v", v)
		if k%100 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
	//计算间隔数出现的次数 35 30 25 20 15 10 5 0
	interval35 := 0
	interval30 := 0
	interval25 := 0
	interval20 := 0
	interval15 := 0
	interval10 := 0
	interval5 := 0
	interval0 := 0
	countS := -1
	countE := -1
	countDot := 0
	countO := 0
	for k, v := range conditions {
		if v == "o" {
			countO++
		} else {
			countDot++
		}

		//先判断是否找到开始和结束
		if countS >= 0 && countE >= 0 {
			length := countE - countS - 1
			if length == 0 {
				interval0++
			} else if length > 0 && length <= 5 {
				interval5++
			} else if length > 5 && length <= 10 {
				interval10++
			} else if length > 10 && length <= 15 {
				interval15++
			} else if length > 15 && length <= 20 {
				interval20++
			} else if length > 20 && length <= 25 {
				interval25++
			} else if length > 25 && length <= 30 {
				interval30++
			} else if length > 30 && length <= 35 {
				interval35++
			}

			countS = countE
			countE = -1
		}

		if v == "o" {
			if countS == -1 {
				countS = k
			} else if countE == -1 {
				countE = k
			}
		}
	}

	fmt.Printf("30-35:%v,%v\n", interval35, float64(interval35)/float64(c))
	fmt.Printf("25-30:%v,%v\n", interval30, float64(interval30)/float64(c))
	fmt.Printf("20-25:%v,%v\n", interval25, float64(interval25)/float64(c))
	fmt.Printf("15-20:%v,%v\n", interval20, float64(interval20)/float64(c))
	fmt.Printf("10-15:%v,%v\n", interval15, float64(interval15)/float64(c))
	fmt.Printf("5-10:%v,%v\n", interval10, float64(interval10)/float64(c))
	fmt.Printf("0-5:%v,%v\n", interval5, float64(interval5)/float64(c))
	fmt.Printf("0:%v,%v\n", interval0, float64(interval0)/float64(c))

}
