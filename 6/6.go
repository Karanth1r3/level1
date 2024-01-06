package main

/*
Task : implement all possible ways to stop goroutines
*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// goroutine can't be stopped by another goroutine
func main() {
	launcFirst()
	launchSecond()
	launchThird()
	fmt.Println("Main Goroutine executed")
}

// launch & test all 3 goroutines with 3 funcs
func launcFirst() {
	stopChan := make(chan struct{})
	go firstWay(stopChan)
	time.Sleep(3 * time.Second)
	stopChan <- struct{}{} // can use empty structs, lightweight way for cases like this
	// close(stopChan) // works similar i guess but can be used once
	fmt.Println("First launch ended")
}

func launchSecond() {
	var (
		stop bool
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go secondWay(&stop, &wg)
	time.Sleep(3 * time.Second)
	stop = true
	wg.Wait()
	fmt.Println("Second launch ended")
}/

func launchThird() {
	ctx, cancel := context.WithCancel(context.Background()) // todo - read about different contexts types
	go thirdWay(ctx)
	time.Sleep(3 * time.Second)
	cancel() // this sends smthng to signal and tells goroutine to select case with exit
	time.Sleep(time.Second)
	fmt.Println("Third launch ended")
}

// first way to stop g-tine, using channel with select case, with ch <- or channel close()
func firstWay(stCh chan struct{}) {
	for {
		select {
		case <-stCh: // if something is coming from stCh, this case will be executed
			fmt.Println("First goroutine stopped")
		default:
			fmt.Println("Running first goroutine")
			time.Sleep(time.Second)
		}

	}
}

// second way - using shared variable with its condition check
func secondWay(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if *stop {
			fmt.Println("Second goroutine stopped")
			return
		}

		fmt.Println("Running second goroutine")
		time.Sleep(time.Second)
	}
}

// third way - using context
func thirdWay(ctx context.Context) {
	for { // infinite loop
		select {
		case <-ctx.Done(): // if signal channel sent something, this case will be selected
			fmt.Println("Third goroutine stopped")
			return
		default: // if no cancellation signals (in this case) -> just sleep and move on
			fmt.Println("Third goroutine running")
			time.Sleep(time.Second)
		}
	}
}
