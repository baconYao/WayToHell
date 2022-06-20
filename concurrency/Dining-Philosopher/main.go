package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

// const
const hunger = 3 // 每位哲學家要吃三次飯

// variables - philosophers
var philosophers = []string{"1-Patrick", "2-Ling", "3-Peter", "4-sherry", "5-Judy"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Println(philosopher, "is seated")
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left\n", philosopher)
		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right\n", philosopher)

		// print a message
		fmt.Println(philosopher, "has both forks, and is eating.")
		time.Sleep(eatTime)

		// Give the philosopher some time to think
		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)

		// unlock the mutexes
		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left\n", philosopher)
		time.Sleep(sleepTime)
	}
	// print out done message
	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table")
}

func main() {
	// print intro
	fmt.Println("The Dining Philosophers problem")
	fmt.Println("-------------------------------")

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

	fmt.Println("The table is empty")
}
