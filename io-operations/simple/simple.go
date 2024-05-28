package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	tStart := time.Now()

	content, err := io.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tRead := time.Now()

	words := strings.Fields(strings.ToLower(string(content)))
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	tProcess := time.Now()

	var ordered []Count
	for word, count := range counts {
		ordered = append(ordered, Count{word, count})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Count > ordered[j].Count
	})
	tSort := time.Now()

	topFive := ordered[:5]

	for _, count := range topFive {
		fmt.Println(string(count.Word), count.Count)
	}

	tEnd := time.Now()

	fmt.Fprintf(os.Stdout, "Reading   : %v seconds\n", tRead.Sub(tStart).Seconds())
	fmt.Fprintf(os.Stdout, "Processing:  %v seconds\n", tProcess.Sub(tRead).Seconds())
	fmt.Fprintf(os.Stdout, "Sorting   :  %v seconds\n", tSort.Sub(tProcess).Seconds())
	fmt.Fprintf(os.Stdout, "Outputting:  %v seconds\n", tEnd.Sub(tSort).Seconds())
	fmt.Fprintf(os.Stdout, "TOTAL     :  %v seconds\n", tEnd.Sub(tStart).Seconds())
}

type Count struct {
	Word  string
	Count int
}
