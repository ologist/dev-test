package main 

import (
	"fmt"
)
// Named Result

func calculate(baseConst int)(x, y int){
	x = baseConst - 1
	y = baseConst + 1
	return
}

func main() {
	fmt.Println(calculate(5))
}

