package main

import(
	"fmt"
)

var x bool

func main(){
	
	fmt.Println(x)
	x := true
	fmt.Println(x)

	x = 10 == 10
	y := (10 > 100)
	z := (20 >= 100)

	fmt.Println(y)
	fmt.Println(z)
}
