package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Tomoe(ab float64, bc float64, ca float64) string {
	//获胜的记录，连续两次就结束
	history := make([]string, 2)
	historyIndex := 0
	player := []string{"A", "B", "C"}
	isWin := false
	rand.Seed(time.Now().UnixNano())
	//假设先上的是A、B
	fight := []string{"A", "B"}
	loser := ""
	for !isWin {

		if (fight[0] == "A" && fight[1] == "B") || (fight[0] == "B" && fight[1] == "A") {
			//进行对抗
			if rand.Float64() <= ab {
				//A胜
				history[historyIndex] = "A"
				loser = "B"
				historyIndex++
			} else {
				//B胜
				history[historyIndex] = "B"
				loser = "A"
				historyIndex++
			}
		} else if (fight[0] == "C" && fight[1] == "B") || (fight[0] == "B" && fight[1] == "C") {
			//进行对抗
			if rand.Float64() <= bc {
				//B胜
				history[historyIndex] = "B"
				loser = "C"
				historyIndex++
			} else {
				//C胜
				history[historyIndex] = "C"
				loser = "B"
				historyIndex++
			}
		} else if (fight[0] == "A" && fight[1] == "C") || (fight[0] == "C" && fight[1] == "A") {
			//进行对抗
			if rand.Float64() <= ca {
				//C胜
				history[historyIndex] = "C"
				loser = "A"
				historyIndex++
			} else {
				//A胜
				history[historyIndex] = "A"
				loser = "C"
				historyIndex++
			}
		}

		//对抗结果出来后，看历史记录
		if historyIndex == 2 {
			//有连续记录后，看是否一致，一致则得到优胜者,不一致的话，则删除前一个记录
			if history[0] == history[1] {
				isWin = true
			} else {
				history[0] = history[1]
				history[1] = ""
				historyIndex--
			}
		}

		//更换选手
		fightIndex := 0
		for _, v := range player {
			if v != loser {
				fight[fightIndex] = v
				fightIndex++
				if fightIndex == 2 {
					break
				}
			}
		}

	}
	return history[0]
}

func main() {
	var pab, pbc, pca float64
	var c int
	flag.Float64Var(&pab, "pab", 0.5, "A对B的胜率，默认0.5")
	flag.Float64Var(&pbc, "pbc", 0.5, "B对C的胜率，默认0.5")
	flag.Float64Var(&pca, "pca", 0.5, "C对A的胜率，默认0.5")
	flag.IntVar(&c, "c", 1000, "实验次数，默认1000")

	flag.Parse()

	//Tomoe(pab,pbc,pca)

	//优胜者结果记录，用A、B、C分别标记A、B、C为最终优胜者
	result := make([]string, c)
	var wg sync.WaitGroup
	for i := 0; i < len(result); i++ {
		wg.Add(1)
		go func(r *string) {
			defer wg.Done()
			*r = Tomoe(pab, pbc, pca)
		}(&result[i])
	}
	wg.Wait()
	countA := 0
	countB := 0
	countC := 0

	for i := 0; i < len(result); i++ {
		switch result[i] {
		case "A":
			countA++
		case "B":
			countB++
		case "C":
			countC++
		}
	}

	fmt.Printf("====%v %v %v\n", pab, pbc, pca)
	fmt.Printf("A:%v %v\n", countA, float64(countA)/float64(c))
	fmt.Printf("B:%v %v\n", countB, float64(countB)/float64(c))
	fmt.Printf("C:%v %v\n", countC, float64(countC)/float64(c))

}
