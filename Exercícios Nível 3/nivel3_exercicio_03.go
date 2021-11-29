/*Nivel 3 - Exercício 3
  - Crie um loop utilizando a sintaxe: for condition {}
  - Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

import "fmt"

func main() {
	ano := 1995
	for ano <= 2021 {
		fmt.Println(ano)
		ano++
	}
}
