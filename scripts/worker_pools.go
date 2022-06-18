package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jobs := make(chan int, 5)    // Canal usado para receber o "trabalho"
	results := make(chan int, 5) // Canal usado para repassar os resultados

	// Inicializa a worker pool
	for i := 1; i <= 3; i++ {
		// Inicializa o worker, originalmente travado, esperando por algum novo job entrar no canal
		// para ser processado
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i // Enfilerando os jobs a serem consumidos pelos workers
	}
	close(jobs) // Fecha o canal de jobs após disparar todos

	for i := 1; i <= 5; i++ {
		// Coleta os resultados, também garante que todas as goroutines finalizem
		fmt.Printf("Result: %d\n", <-results)
	}
}

// Declara o worker que processará as informações usando um canal de leitura (jobs) e um de
// saída (results)
func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, i)
		time.Sleep(randSeconds()) // Simula o tempo de processamento de um job longo
		results <- i * 2
		fmt.Printf("Worker %d finished job %d\n", id, i)
	}
}

func randSeconds() time.Duration {
	rand.Seed(time.Now().UnixMicro())

	n := rand.Intn(10)
	seconds := n * int(time.Second)

	return time.Duration(seconds)
}
