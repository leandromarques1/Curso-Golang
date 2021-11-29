/*Nivel 2
Ex. 05
Crie uma variável de tipo string utilizando uma 
raw string literal.
- Demonstre-a.
*/

package main

import (
	"fmt"
)

func main() {
	frase := `Say
		hello to
			my little friend!`
	fmt.Println(frase)
	fmt.Printf("\n%T", frase)
}
