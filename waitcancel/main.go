package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func doAPICall(ctx context.Context, wg *sync.WaitGroup, logger *log.Logger) error {
	defer wg.Done()

	req, err := http.NewRequest("GET", "https://httpstat.us/200", nil)
	if err != nil {
		return err
	}

	// The httpstat.us API accepts a sleep parameter which sleeps the request for the
	// passed time in ms
	q := req.URL.Query()
	sleepMin := 5000
	sleepMax := 10000
	r := rand.Intn(sleepMax-sleepMin) + sleepMin
	q.Set("sleep", fmt.Sprintf("%d", r))
	req.URL.RawQuery = q.Encode()

	c := make(chan error, 1)
	go func() {
		// For the purposes of this example, we're not doing anything with the response.
		r, err := http.DefaultClient.Do(req)
		logger.Println("received:", r.StatusCode, q.Encode())
		logger.Println(err)
		c <- err
	}()

	// Block until either channel is populated or closed
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-c:
		if err != nil {
			logger.Println(err.Error())
		}
		return err
	}
}

func main() {
	var (
		closing     = make(chan struct{})
		ticker      = time.NewTicker(1 * time.Second)
		logger      = log.New(os.Stderr, "", log.LstdFlags)
		batchSize   = 6
		wg          sync.WaitGroup
		ctx, cancel = context.WithCancel(context.Background())
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
		<-signals
		logger.Println("Initiating shutdown of producer.")
		cancel()
		close(closing)
	}()
loop:
	for {
		select {
		case <-closing:
			break loop
		case <-ticker.C:
			for n := 0; n < batchSize; n++ {
				wg.Add(1)
				go doAPICall(ctx, &wg, logger)
			}
			wg.Wait()
			logger.Printf("Completed doing %d things.", batchSize)
		}
	}
}
