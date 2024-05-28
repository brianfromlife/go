// Run with:
// go build optimized.go
// sudo sysctl vm.drop_caches=3; ./optimized <kjvbible_x100.txt >/dev/null

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	tStart := time.Now()
	dRead := time.Duration(0)

	var word []byte
	buf := make([]byte, 64*1024)
	counts := make(map[string]*int)
	for {
		// Read input in 64KB blocks till EOF.
		tRead1 := time.Now()
		n, err := os.Stdin.Read(buf)
		tRead2 := time.Now()
		dRead += tRead2.Sub(tRead1)
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}

		// Count words in the buffer.
		for i := 0; i < n; i++ {
			c := buf[i]

			// Found a whitespace char, count last word.
			if c <= ' ' {
				if len(word) > 0 {
					increment(counts, word)
					word = word[:0] // reset word buffer
				}
				continue
			}

			// Convert to ASCII lowercase as we go.
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
			}

			// Add non-space char to word buffer.
			word = append(word, c)
		}
	}

	// Count last word, if any.
	if len(word) > 0 {
		increment(counts, word)
	}

	tProcess := time.Now()

	// Convert to slice of Count, sort by count descending, print.
	ordered := make([]Count, 0, len(counts))
	for word, count := range counts {
		ordered = append(ordered, Count{word, *count})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Count > ordered[j].Count
	})
	tSort := time.Now()

	topFive := ordered[:5]

	for _, count := range topFive {
		fmt.Println(count.Word, count.Count)
	}

	tEnd := time.Now()

	fmt.Fprintf(os.Stderr, "Reading   : %v\n", dRead.Seconds())
	fmt.Fprintf(os.Stderr, "Processing: %v\n", (tProcess.Sub(tStart) - dRead).Seconds())
	fmt.Fprintf(os.Stderr, "Sorting   : %v\n", tSort.Sub(tProcess).Seconds())
	fmt.Fprintf(os.Stderr, "Outputting: %v\n", tEnd.Sub(tSort).Seconds())
	fmt.Fprintf(os.Stderr, "TOTAL     : %v\n", tEnd.Sub(tStart).Seconds())
}

func increment(counts map[string]*int, word []byte) {
	if p, ok := counts[string(word)]; ok {
		// Word already in map, increment existing int via pointer.
		*p++
		return
	}
	// Word not in map, insert new int.
	n := 1
	counts[string(word)] = &n
}

type Count struct {
	Word  string
	Count int
}
