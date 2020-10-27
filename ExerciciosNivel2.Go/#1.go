// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import(
	"fmt"
)

var x int

func main(){

	x := 10

	fmt.Printf("Decimal: %d", x)
	fmt.Printf("\n Binário: %b", x)
	fmt.Printf("\n Binário: %x", x)

}