package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 Aprofundamento no uso de goroutines e channels
*/
func main() {
	/* Goroutines são o modelo de concorrência de mais alto nível em Go, sendo executadas como
	threads leves, gerenciadas pelo agendador go, que as executa de modo m:n, ou seja m goroutine
	agendadas para n threads de sistema.

	Goroutines atuam de modo puramente assíncrono, retornando imediamente após ser disparada
	*/
	go func() { // Disparando uma goroutine usando a keyword "go"
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
		}
	}()

	/* Vale notar que todas as execuções em go são goroutines, incluindo a função main(), então
	precisamos usar alguma estratégia para que o programa não feche sem que a goroutine seja
	executada, para isso podemos usar três estratégias:
	time.Sleep -> espera um determinado tempo antes de finalizar o programa
	select{} -> trava a execução do programa indefinidamente usando select
	sync.WaitGroup -> usa grupos de espera (WaitGroups) para esperar n goroutines terminarem
	*/
	var wg sync.WaitGroup   // Declarando um waitgroup
	fmt.Printf("%#v\n", wg) // Estrutura interna de um wait group
	for i := 0; i < 100; i++ {
		// Incrementa o contador de membros do grupo a cada execução, é importante que ele seja
		// usado fora de goroutines, para evitar condições de corrida
		wg.Add(1)
		go func(x int) {
			defer wg.Done() // Decrementa o contador de membros do grupo sempre que uma goroutine finaliza
			fmt.Printf("%d ", x)
		}(i)
	}
	wg.Wait() // Espera até que o contador de membros do grupo seja zero para finalizar a execução
	fmt.Println("")

	/* Channels atuam como canais que permitem que goroutines se comuniquem, ou seja, que troquem
	dados

	Uma de suas características mais importantes é podermos especificar a direção de um canal, ou
	seja, podemos indicar se um canal está enviando ou recebendo dados.

	Outra importante característica é que a comunicação por canais é bloqueante, ou seja, uma operação
	de leitura ou escrita de valores com canais espera até que o canal receba ou atribua um valor para
	continuar com a execução, parafraseado a documentação de Go (https://tour.golang.org/concurrency/2):

				"By default, sends and receives block until the other side is ready.
				This allows goroutines to synchronize without explicit locks or
				condition variables."
	*/
	c := make(chan int) // Cria um novo canal com a função make(chan), indicando seu tipo (int)
	go writeToChannel(c, 10)
	time.Sleep((1 * time.Second))
	k := <-c // Lendo o conteúdo do canal e atribuindo seu valor final a outra variável
	fmt.Printf("K: %d\n", k)

	// Exemplo de comunicação usando funções com canais direcionados
	in := make(chan int)
	out := make(chan int)
	go writeDirectedChannel(out, in) // Coloca a goroutine para esperar um valor no canal "in"
	in <- 20                         // Atribui um valor ao canal "in", liberando a gouroutine
	fmt.Printf("Out: %d\n", <-out)   // Exibe o valor do canal "out"

	/* As características dos canais permitem que criemos pipelines, ou seja, fluxos de dados
	continuos entre multiplas goroutines
	*/
}

func writeToChannel(c chan int, x int) {
	fmt.Printf("X: %d\n", x)
	c <- x   // Escreve no canal, a direção da flecha (<-) indica o fluxo dos dados
	close(c) // Fecha o canal

	// Essa linha nunca será executada, pois a ação c <- permanece bloqueando o contexto, pois
	// ninguém está lendo o que está sendo escrito no canal c, fazendo com que o tempo de sleep
	// acabe antes dessa função
	fmt.Printf("X: %d\n", x)
}

func writeDirectedChannel(out chan<- int, in <-chan int) {
	// Definindo restrições de direção de escrita, denotando que o canal out somente pode ser usado
	// para escrita e o canal in somente para leitura
	x := <-in // Espera até que um novo valor seja atribuido ao canal in para seguir com a execução
	out <- x
}
