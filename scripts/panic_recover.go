package main

import "fmt"

// Aprofundamento no uso de panic e recover
func main() {
	defer func() {
		/* Recover

		recover é a função responsável por reassumir o controle do fluxo
		de uma função ou programa que sofreu um Panic, recovers são utilizados
		apenas dentro de blocos defer, pois a recuperação de um panic
		só pode ocorrer depois da parada ocorrer, logo disparando o bloco
		defer.

		Em um recover "r" representa a mensagem do panic que foi responsável
		por parar a execução da goroutine
		*/
		if r := recover(); r != nil {
			fmt.Println("Recovered from the crash state:", r)
		}
	}()

	/* Panic

	panic é a função responsável por lançar um erro capaz de parar o fluxo
	de execução de uma função, panics são de fato as exceções de Go, complementando
	os errors, que sozinhos não conseguem parar o fluxo de execução de um
	programa.

	Panics devem ser usados apenas quando é essencial que o programa pare seu
	fluxo de funcionamento quando um erro ocorrer. Exemplos:
	- Quando um arquivo de configuração não existe;
	- Quando o sistema está perto de uma sobrecarga de memória;
	- Quando uma das entradas de um CLI é errada;
	- Quando tentamos acessar um indíce que não existe em um array/slice.

	Panics exibem uma mensagem de erro dividida em duas partes>:
	1. A mensagem de erro em sim;
	2. O stack-trace do erro, usado para debug.
	*/
	panic("A exceptional error occurred, so the program must crash")
}
