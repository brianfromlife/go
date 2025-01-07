package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	words := []string{"foo", "bar", "baz"}

	for _, word := range words {
		wg.Add(1)
		go func(word string) {
			log.Println("adding", word)
			time.Sleep(1 * time.Second)
			defer wg.Done()
			fmt.Println(word)
		}(word)
	}
	// do concurrent things here

	// blocks/waits for waitgroup
	wg.Wait()
}
