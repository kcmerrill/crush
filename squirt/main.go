package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	workers := flag.Int("workers", 10, "# of workers")
	consumer := flag.String("consumer", "echo ACK {{ .id }}", "Consumer script")
	backoff := flag.Duration("backoff", time.Second, "How long to backoff if queue is empty")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Please provide a producer")
		os.Exit(1)
	}

	// giddy up!
	squirt(*workers, *consumer, flag.Args(), *backoff)
}
