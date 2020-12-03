package basic

import (
	"fmt"
)

func returnMulti() (int32, int32) {
	return 1, 2
}

func FunctionExperiment() {
	a, b := returnMulti()
	fmt.Println(a, b)

	c, _ := returnMulti()
	fmt.Println(c)

}
