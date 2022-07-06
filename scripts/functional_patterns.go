/* Aprofundamento no uso de padrões funcionais em Go

Como Go possui suporte de primeira classe a funções alguns padrões de linguagens funcionais
se tornam aplicáveis

OBS: Alguns padrões funcionais utilizam além de funções, objetos (structs) e interfaces
*/
package main

import "fmt"

// Função placeholder, para ser usada em exemplos e aplicações
func Sum(a, b int) int {
	return a + b
}

/* Currying

É o processo de converter uma função que recebe multiplos argumentos em uma função
e ordem maior que aceita apenas um argumento e retorna funções que aceitam os demais argumentos
até que todos sejam passados

Função sem currying:
> func sum(a, b) -> a + b
> sum(1, 2) -> 3

Função com currying:
> func curried(a) -> func (b) -> a + b
> curried(1)(2) -> 3

Essas função são úteis quando se precisa aplicar um valor padrão a uma operação em que os outros
valores variam, resutlando em um código mais limpo

Exemplo:
> func curried(a) -> func (b) -> a + b
> sum10 <- curried(10)
> sum10(1) -> 11
> sum10(2) -> 12
*/
func CurriedSum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

/* Aplicações parciais

É o processo de criar uma função com uma aridade (número de argumentos) menor em relação a função
original, complementando o número de argumentos com valores nulos ou vazios

Função sem aplicação parcial:
> func sum(a, b) -> a + b

Função com aplicação parcial de sum
> func partial(a) -> sum(a, 0)
*/
func PartialSum(b int) int {
	return Sum(0, b)
}

/* Predicado

É uma função que retona um resultado booleano a partir de uma ou mais verificações, são utilizadas
para validar sequencias de operações em filtros e callbacks.

Função predicado:
> predicate(v) -> value > 25
> predicate(20) -> false
> predicate(30) -> true

Predicados podem ser implementados comparando diretamente os seus argumentos ou operando-os e
comparando o resultado da operação

Função predicado que opera com uma operação interna
> predicate_sum(a, b) -> (a + b) > 25
> predicate_sum(10, 10) -> false
> predicate_sum(10, 20) -> true
*/
func SumIsLowerThan25(a, b int) bool {
	return Sum(a, b) < 25
}

/* Setóide

É qualquer objeto que implementa um método de comparação que pode er usado para comparar outros
objetos do mesmo tipo

Objeto setoide:
> obj Foo(attr) ->
>     this.attr = attr
>     this.equals -> func(b) -> this.attr == b.attr
>
> f1 = Foo('a')
> f2 = Foo('b')
>
> f1.equals(f2) -> false
*/
type SetoideSlice []int

func (s SetoideSlice) Equals(slice SetoideSlice) bool {
	length := len(s)

	if length != len(slice) {
		return false
	}

	for i := 0; i < length; i++ {
		if s[i] != slice[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Printf("Resultado Curry: %d\n", CurriedSum(1)(2))
	fmt.Printf("Resultado Parcial: %d\n", PartialSum(2))
	fmt.Printf("Resultado Predicado: %v\n", SumIsLowerThan25(10, 5))
	fmt.Printf("Resultado Setoide: %v\n", SetoideSlice{1, 2, 3}.Equals(SetoideSlice{1, 2, 3}))
}
