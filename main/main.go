package main

import (
	"fmt"

	"github.com/iphuket/smp"
)

func main() {
	// fmt.Println(smp.EuclideanAlgorithm(2560, 1440))
	m, n := smp.DisplayProportion(2560, 1440)
	fmt.Println("显示器比例：", m, n)
}
