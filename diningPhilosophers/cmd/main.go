package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Philosophers names
var philosophers = []string{"Mark", "Rusell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var fmt = log.New(os.Stdout, "", 0)

var dinning sync.WaitGroup

// dinningProblem simulates in a goroutine the behaviour of each philosopher
func dinningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(phName, "Seated")

	h := fnv.New64a()
	h.Write([]byte(phName))

	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}

	for h := hunger; h > 0; h-- {
		fmt.Println(phName, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}

	fmt.Println(phName, "Satisfied")
	dinning.Done()
	fmt.Println(phName, "Left the table")
}

func main() {
	fmt.Println("Table empty")
	dinning.Add(5) // wait group of 5 philosophers
	fork0 := &sync.Mutex{}
	forkLeft := fork0

	for i := 1; i < len(philosophers); i++ {
		forkRight := &sync.Mutex{}
		go dinningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}

	go dinningProblem(philosophers[0], fork0, forkLeft)
	dinning.Wait() // wait for philosophers to finish
	fmt.Println("Table empty")
}
