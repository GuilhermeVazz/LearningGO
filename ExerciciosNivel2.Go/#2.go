// - Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
// - Demonstre estes valores.

package main

import (
	"fmt"
)

var x int
var y int

func main(){

x := 10
y := 5

z := x == y
fmt.Println(z)
a := x <= y
fmt.Println(a)
b := x >= y
fmt.Println(b)
c := x != y
fmt.Println(c)
d := x > y
fmt.Println(d)
e := x < y
fmt.Println(e)

}