package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func checkArgs() bool {
	if len(os.Args) != 2 {
		return false
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return false
	}
	if !(n >= 1) {
		return false
	}
	return true
}

func startChanReader(dataChan <-chan int) {
	for range dataChan {
	}
}

func main() {
	if !checkArgs() {
		fmt.Printf("Invalid input arguments\nUsage:\nprogram N\nwhere N - natural number.\n")
		return
	}
	nInt, _ := strconv.Atoi(os.Args[1])
	n := time.Duration(nInt)

	dataChan := make(chan int)
	end := false
	timer := time.NewTimer(n * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-timer.C
		end = true
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		startChanReader(dataChan)
	}()

	for !end {
		dataChan <- 123
	}
	close(dataChan)
	wg.Wait()
}
