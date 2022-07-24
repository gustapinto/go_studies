package main

import (
	"fmt"
	"sync"
)

// Aprofundamento no uso de mutexes em Golang
func main() {
	counter := 0

	/* Mutex

	Mutexes são operadores sincronização (assim como atomics e channels) que
	implementam o mecanismo de exclusão mútua, que atua como um semáforo,
	permitindo que uma goroutine altere o valor de um elemento apenas se
	nenhuma outra estiver atuando sobre esse elemente, evitando que condições
	de corrida aconteçam.

	Mutexes podem ser implementados em funções, atuando como semáforos ou
	acoplados diretamente em structs.
	*/
	var m sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go add(&m, &wg, &counter)
	}

	wg.Wait()

	fmt.Println(counter)
}

func add(m *sync.Mutex, wg *sync.WaitGroup, counter *int) {
	m.Lock() // Fecha o semáforo do mutex

	for i := 0; i < 1000; i++ {
		*counter += i
	}

	m.Unlock() // Abre o semáforo do mutexes
	wg.Done()
}
