/* Aprofundamento no uso de padrões de gerenciamento de erros

OBS: É importante frisar que os padrões demonstrados são considerados como forma não idiomáticas
de lidar com erros em Go e que a única forma realmente idiomática é utilizando a estrutura:
> if err != nil {
>     // lida com o erro
> }
*/
package main

import (
	"errors"
	"fmt"
)

/* Unwrap

Unwrap é um padrão utilizado majoriamente em linguagens com suporte a paradigmas funcionais, como
Elixir e Rust, uma função de unwrap (desembrulhar) contém aridade variádica e trata o seu último
argumento como sendo um erro, ignorando-o se não for nulo. Unwrap retorna todos os elementos que
não sejam erros como uma slice de tipo []any
*/
func Unwrap(args ...any) []any {
	last, isErr := args[len(args)-1].(error)
	if !isErr && last != nil {
		panic("Expected last value to be an error")
	}

	return args[0 : len(args)-1]
}

/* Except

Except atua como unwrap, ditando que o último parâmetro da função seja um erro e disparando um pânico
caso esse erro seja não nulo, providenciando uma funcionalidade análoga as Exceções de outras
linguagens de progrmação. Assim com Unwrap, Except também retorna todos os elementos que não sejam
erros como uma slice de tipo []any
*/
func Except(args ...any) []any {
	last, isErr := args[len(args)-1].(error)
	if !isErr && last != nil {
		panic("Expected last value to be an error")
	}

	if last != nil {
		panic(last.Error())
	}

	return args[0 : len(args)-1]
}

func main() {
	f := func() (int, error) {
		return 1, errors.New("f err")
	}
	g := func(args ...any) {
		fmt.Println(args...)
	}
	h := func() (int, int, error) {
		return 2, 3, nil
	}

	n := Unwrap(f())

	// Um dos problemas da utilização de Unwrap em Go é a necessidade de acessar manualmente a slice
	// de resultados e utilizar type assertion ao passar seus retornos para outra função, tornando
	// o código mais "sujo"
	g(n[0].(int))

	o := Except(h())
	g(o[0].(int), o[1].(int))

	_ = Except(f())
}
