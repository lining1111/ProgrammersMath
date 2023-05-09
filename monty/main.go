package main

import (
	"flag"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
	"sync"
	"time"

	"gonum.org/v1/plot/plotter"
)

type ChangeNot struct {
	Result int //结果 0 错误， 1 正确
}

type Change struct {
	Result int   //结果 0 错误， 1 正确
	Choice []int //可选结果
	Choose int   //选择结果
}

var doors = []int{0, 1, 2}

type Condition struct {
	Awarded   int       //有奖的门
	Player    int       //玩家选择的门
	Emcee     int       //主持人打开的门
	ChangeNot ChangeNot //不改变的情景
	Change    Change    //改变的情景
}

func (c *Condition) Init() {
	//有将门和玩家选择的门，随机生成
	rand.Seed(time.Now().UnixNano())
	c.Awarded = doors[rand.Intn(len(doors))]
	c.Player = doors[rand.Intn(len(doors))]
	//主持人打开的必须和有奖的门还有玩家的门不一样
	var emceeChoice []int
	for _, v := range doors {
		if v != c.Player && v != c.Awarded {
			emceeChoice = append(emceeChoice, v)
		}
	}
	c.Emcee = emceeChoice[rand.Intn(len(emceeChoice))]

	//不改变的话，可以在初始化的时候就确定结果
	if c.Awarded == c.Player {
		c.ChangeNot.Result = 1
	} else {
		c.ChangeNot.Result = 0
	}

	//默认改变的是0
	for _, v := range doors {
		if v != c.Player && v != c.Emcee {
			c.Change.Choice = append(c.Change.Choice, v)
		}
	}
	c.Change.Result = 0
}

func (c *Condition) MakeChange() {
	//作出改变
	rand.Seed(time.Now().UnixNano())
	c.Change.Choose = c.Change.Choice[rand.Intn(len(c.Change.Choice))]
	if c.Change.Choose == c.Awarded {
		c.Change.Result = 1
	} else {
		c.Change.Result = 0
	}
}

type MontyResult struct {
	Count               int     //总场数
	WinChangeNot        int     //不改变赢的场数
	WinChangeNotPercent float64 //不改变赢的场数百分比
	WinChange           int     //改变赢的场数
	WinChangePercent    float64 //改变赢的场数百分比
}

func monty(c int) MontyResult {
	conditions := make([]Condition, c)
	var wg sync.WaitGroup
	for k, _ := range conditions {
		wg.Add(1)
		go func(condition *Condition) {
			defer wg.Done()
			condition.Init()
			condition.MakeChange()
		}(&conditions[k])
	}
	wg.Wait()

	//跑完后，统计结果
	winChangeNot := 0
	winChange := 0

	for _, v := range conditions {
		if v.ChangeNot.Result == 1 {
			winChangeNot++
		}
		if v.Change.Result == 1 {
			winChange++
		}
	}

	fmt.Printf("一共%v场,不改变赢的有%v(%v),改变赢的有%v(%v)\n", c, winChangeNot, float64(winChangeNot)/float64(c),
		winChange, float64(winChange)/float64(c))
	return MontyResult{c, winChangeNot, float64(winChangeNot) / float64(c),
		winChange, float64(winChange) / float64(c)}
}

func main() {
	c := 1000
	flag.IntVar(&c, "c", 1000, "实验次数，尽量的大")
	flag.Parse()

	//每一次的场景内容，包括有奖门，玩家选择门，主持人开启门，不选择改变的结果，选择改变的结果(其中选择改变时，可选的结果，选择的结果)
	//monty(c)

	count := 0
	montyResult := make([]MontyResult, c)
	for count < c {
		count++
		mr := monty(count)
		montyResult = append(montyResult, mr)
	}
	//画函数曲线来看趋势
	pointsChangeNot := plotter.XYs{}
	pointsChange := plotter.XYs{}
	for _, v := range montyResult {
		pointsChangeNot = append(pointsChangeNot, plotter.XY{
			X: float64(v.Count),
			Y: v.WinChangeNotPercent,
		})

		pointsChange = append(pointsChange, plotter.XY{
			X: float64(v.Count),
			Y: v.WinChangePercent,
		})
	}

	plt := plot.New()

	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 1,float64(c)

	if err := plotutil.AddLines(plt,
		"line1", pointsChangeNot,
		"line2", pointsChange,
	); err != nil {
		panic(err)
	}

	if err := plt.Save(5*vg.Inch, 5*vg.Inch, "monty.png"); err != nil {
		panic(err)
	}
}
