// - For: inicialização, condição, pós
// - For: condição ("while")
// - For: ...ever? (http servers)
// - For: break
// - golang.org/ref/spec#For_statements, Effective Go
// - (Range vem mais pra frente.) 

package main

import(
	"fmt"
)

func main(){
	a := 2

	for {
		fmt.Println("Oi")
	}

	for a < 10{
		a +=2
		fmt.Println(a)
	}

	for a:=2 ;a < 10; a++{
		fmt.Println(a)
	}

}