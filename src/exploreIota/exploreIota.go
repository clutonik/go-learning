package exploreIota

import (
	"fmt"
)

const (
	_ = iota            //Dumping the first iota = 0
	a = 1 << (iota * 2) // iota = 1, multiplying iota by 2 and using bitwise left shift
	b = 1 << (iota * 2) // iota = 2, multiplying iota by 2 and using bitwise left shift
	c = 1 << (iota * 2) // iota = 3, multiplying iota by 2 and using bitwise left shift
)

func PrintIota() {
	fmt.Println("First iota: ", a)
	fmt.Println("Second iota: ", b)
	fmt.Println("Third iota: ", c)
}
