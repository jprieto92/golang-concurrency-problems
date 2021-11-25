package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	partList   = []string{"A", "B", "C", "D"}
	nAssembles = 3
	wGroup     sync.WaitGroup
)

// worker runs in a goroutine simulating the behavior of every worker per cycle
func worker(part string) {
	log.Println(part, "Worker begins part")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "Worker completes part")
	wGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssembles; c++ {
		log.Println("Begin assemble cycle", c)
		wGroup.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}
		wGroup.Wait()
		log.Println("Assemble. Cycle", c, "complete")
	}
}
