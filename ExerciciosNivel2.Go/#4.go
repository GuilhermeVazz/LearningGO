// Crie um programa que:
//     - Atribua um valor int a uma variável
//     - Demonstre este valor em decimal, binário e hexadecimal
//     - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
//     // - Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import(
	"fmt"
)
var x int
func main(){

	x := 10
	fmt.Printf("%d\t%b\t%#x", x , x, x)
	
	b := x << 1
	fmt.Printf("\n%d\t%b\t%#x", b, b, b)
}