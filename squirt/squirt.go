package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func squirt(workers int, consumer string, producer []string, backoff time.Duration) {
	var wg sync.WaitGroup
	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func() {
			for {
				fmt.Println("starting worker ...")
				// start our worker
				work := exec.Command("sh", "-c", consumer)
				workerProcess, _ := work.StdinPipe()
				work.Stdout = os.Stdout
				workErr := work.Start()
				// error?
				if workErr != nil {
					// backoff ...
					fmt.Println("worker error ...")
					<-time.After(backoff)
				}
				for {
					<-time.After(time.Second)
					fmt.Println("executing command ...")
					cmd := exec.Command("sh", "-c", strings.Join(producer, " "))
					fmt.Println(strings.Join(producer, " "))

					msg, cmdErr := cmd.CombinedOutput()
					if cmdErr != nil {
						fmt.Println(cmdErr.Error())
						<-time.After(backoff)
						continue
					}

					fmt.Println("msg", string(msg))

					// write our message to the app
					_, writeErr := workerProcess.Write([]byte(msg))
					if writeErr != nil {
						fmt.Println("writeErr() :(")
						break
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
