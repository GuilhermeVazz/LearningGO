package main

import(
	"fmt"
)

func main(){

	for caixa := 1; caixa < 6; caixa++{
		fmt.Print("Caixa: ", caixa)
		fmt.Print("\n")
		for garrafa := 1; garrafa <= 12; garrafa++{
			fmt.Print("Garrafas: ", garrafa, ", ")
			
		}
		fmt.Print("\n")
	}
		
}