/*Nivel 3 - Exercício 5
  - Demonstre o resto da divisão por 4
    de todos os números entre 10 e 100
*/

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		resto := i % 4
		fmt.Println("Nº", i, "dividido por 4 é igual a", resto)
	}
}
