package main

import "sync"

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

// variables - philosophers
var philosophers = []string{"Patrick", "Ling", "Peter", "sherry", "Judy"}
var wg sync.WaitGroup

func diningProblem(philosopher string, dominantHand, otherHand *sync.Mutex) {
	defer wg.Done()

	// print a message

	// lock both forks

	// print a message

	// unlock the mutexes
}

func main() {
	// print intro

	forkLeft := &sync.Mutex{}
	wg.Add(len(philosophers))
	// spawn one goroutine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forkRight := &sync.Mutex{}

		// call a goroutine
		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight
	}
	wg.Wait()
}
