/*Nivel 2
Ex. 04
Crie um programa que:
- Atribua um valor int a uma variável
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa variável 1 para a esquerda,
  e atribua o resultado a outra variável
- Demonstre esta outra variável em decimal,
  binário e hexadecimal
*/

package main

import (
	"fmt"
)

func main() {
	a := 777
	fmt.Printf("\nDecimal: %d", a)
	fmt.Printf("\nDBinário: %b", a)
	fmt.Printf("\nHexadecimal: %#x", a)

	/*Desloque os bits dessa variável 1 para a esquerda*/
	b := a << 1
	
	/*atribua o resultado a outra variável*/
	fmt.Printf("\n\nNovo resultado: %v",b)
	
	/*Demonstre esta outra variável em decimal, binário e hexadecimal*/
	fmt.Printf("\nDecimal: %d", b)
	fmt.Printf("\nDBinário: %b", b)
	fmt.Printf("\nHexadecimal: %#x", b)
	
}
