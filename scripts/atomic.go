package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/* Além de canais e mutexes golang possui outro primitivo para concorrência: operações atômicas, essas
 * são operações implementadas em nível de hardware e são utilizadas prinicipalmente para sincronização
 * de goroutines
 *
 * Por conta de serem implementadas em nível de hardware operações atômicas funcionam apenas com
 * números inteiros, por isso são primariamente utilizadas para a criação de contadores em worker
 * pools e outras pipelines concorrentes
 */
func main() {
	var counter uint64 // Instancia uma varável que irá atuar como contador
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go add(&wg, &counter)
	}

	wg.Wait()

	fmt.Printf("Counter: %d", counter)
}

func add(wg *sync.WaitGroup, counter *uint64) {
	for i := 1; i <= 1000; i++ {
		// Utiliza uma operação atômica para incrementar o valor do contador
		atomic.AddUint64(counter, 1)
	}

	wg.Done()
}
