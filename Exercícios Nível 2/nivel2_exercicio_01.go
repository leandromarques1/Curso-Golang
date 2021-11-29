/*Nivel 2
Ex. 01
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	num := 777
	fmt.Printf("Decimal: %d", num)
	fmt.Printf("\nBinário: %b", num)
	fmt.Printf("\nHexadecimal: %#x", num)

}