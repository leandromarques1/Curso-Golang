/*Nivel 2
Ex. 06
Utilizando iota, crie 4 constantes cujos valores 
sejam os próximos 4 anos.
- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

const (
	_ = 2020 + iota
	b
	c
	d
	e  
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
