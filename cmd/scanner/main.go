package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"PortScanner/internal/scanner"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "Host to scan")
	ports := flag.String("ports", "75-85", "Port range to scan(e.g. 20-1000)")
	flag.Parse()

	start, end, err := scanner.ParsePortRange(*ports)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	results := make(chan string)
	var wg sync.WaitGroup

	// Scan ports
	for port := start; port <= end; port++ {
		wg.Add(1)
		go scanner.ScanPort(*host, port, &wg, results)
	}

	// Close results once done
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("Scanning %s on port %s...\n\n", *host, *ports)

	for r := range results {
		if strings.Contains(r, "open") {
			fmt.Println(r)
		}
	}
}
