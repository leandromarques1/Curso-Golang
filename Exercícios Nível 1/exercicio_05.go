/* Exercício Nº 05
- Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". 
       O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.

    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, 
               utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"
*/

package main

import (
	"fmt"
)

type algumaCoisa int

var x algumaCoisa
var y int //o tipo SUBJACENTE é "int"

func main() {
	//Demonstrar o valor da variável "x"
	fmt.Println(x)

	//Demonstrar o tipo da variável "x"
	fmt.Printf("%T\n", x)

	//Atribuindo novo valor a "x"
	x = 42
	
	//Demonstrar o valor da variável "x"
	fmt.Println(x)

	//Utilizando conversão, transformar o tipo do valor da variável "x" em seu tipo subjacente
	//Atribuir o valor de "x" a "y"
	y = int(x)
	
	//Demonstrar o valor da variável "y"
	fmt.Println(y)

	//Demonstrar o tipo da variável "y"
	fmt.Printf("%T\n", y)
}