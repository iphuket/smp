package smp

import (
	"fmt"
	"testing"
)

func TestDisplayProportion(T *testing.T) {
	m, n := DisplayProportion(1920, 1080)
	if m == 16 && n == 9 {
		fmt.Println("smp test ok 1920/1080:m/n", m, "/", n)
		return
	}
	T.Errorf("smp test error")
}
