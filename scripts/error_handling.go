/* Aprofundamento no uso de padrões de gerenciamento de erros

OBS 1: É importante frisar que os padrões demonstrados são considerados como forma não idiomáticas
de lidar com erros em Go e que a única forma realmente idiomática é utilizando a estrutura:
> if err != nil {
>     // lida com o erro
> }

OBS 2: Esses padrões não são formas comuns e aceitas pela comunidade Go, mas sim apenas formas
alternativas de gerenciamento de erros, trazidos muitas vezes de outras comunidades ou propostas
que não foram concretizadas
*/
package main

import (
	"errors"
	"fmt"
)

var ErrF = errors.New("err F")

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

/* Match

Match, ou MatchErr é uma forma de tratar erros a partir de uma função variádica que recebe o erro,
e uma sequência de argumentos na forma: Tipo, Função, etc...

Exemplo:
> func match_err(error, nil, func(), int_err, func_int_err(), ....)
*/
type H struct {
	errType    error
	errHandler func(error)
}

func MatchErr(err error, handlers ...H) {
	for _, handler := range handlers {
		if errors.Is(err, handler.errType) {
			handler.errHandler(err)
		}
	}
}

func main() {
	f := func() (int, error) {
		return 1, ErrF
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

	_, err := f()
	// Estruturas de Match são melhores usadas com handlers genéricos nomeados, pois tornam
	// a chamada de função mais limpa e permitem que o mesmo tipo de erro seja sempre tratado
	// pela mesma função de handler
	ignoreErr := func(err error) {}
	printErr := func(err error) { fmt.Println("Err:", err) }

	MatchErr(err, H{nil, ignoreErr}, H{ErrF, printErr})

	_ = Except(f())
}
