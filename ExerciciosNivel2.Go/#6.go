// - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

package main

import(
	"fmt"
)

const (
	a = 2021 + iota
	b
	c
	d
)

func main(){

	fmt.Println("Proximos anos: ", a, b, c ,d)
}