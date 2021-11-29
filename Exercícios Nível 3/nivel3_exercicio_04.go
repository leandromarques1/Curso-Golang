/*Nivel 3 - Exercício 4
  - Crie um loop utilizando a sintaxe: for {}
  - Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import "fmt"

func main() {
	anoNasc := 1995
	anoAtual := 2021
	for {
		if anoNasc == anoAtual+1 {
			break
		}
		fmt.Println(anoNasc)
		anoNasc++
	}
}
