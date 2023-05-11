package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Sugoroku 返回0成功 1失败
func Sugoroku(t int) int {
	rand.Seed(time.Now().UnixNano())
	length := 1000
	//置色子，每次走1-6中的随机数
	trip := make([]int, t)
	for k, _ := range trip {
		if k == 0 {
			trip[k] = rand.Intn(1000)
		} else {
			isFind := false
			var val int
			for !isFind {
				val = rand.Intn(1000)
				isSame := false
				//需要跟之前的都不一样
				for i := 0; i < k; i++ {
					if trip[i] == val {
						isSame = true
					}
				}
				isFind = !isSame
			}
			trip[k] = val
		}
	}
	ret := 0
	step := 0
	for step < length {
		for _, v := range trip {
			if step == v {
				ret = 1
				break
			}
		}
		if ret == 1 {
			break
		}
		step += rand.Intn(6) + 1
	}
	return ret
}

func main() {
	var t int
	var c int
	flag.IntVar(&t, "t", 1, "陷阱个数，默认1")
	flag.IntVar(&c, "c", 10000, "实验次数，默认10000")

	flag.Parse()
	//设置随机陷阱
	result := make([]int, c)
	var wg sync.WaitGroup
	for i := 0; i < len(result); i++ {
		wg.Add(1)
		go func(r *int) {
			defer wg.Done()
			*r = Sugoroku(t)
		}(&result[i])
	}
	wg.Wait()
	countWin := 0
	countLose := 0

	for _, v := range result {
		if v == 0 {
			countWin++
		} else if v == 1 {
			countLose++
		}
	}

	fmt.Printf("t %v c %v\n", t, c)
	fmt.Printf("win:%v %v\n", countWin, float64(countWin)/float64(c))
	fmt.Printf("lose:%v %v\n", countLose, float64(countLose)/float64(c))
}
