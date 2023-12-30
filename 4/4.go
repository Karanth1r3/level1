package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

// Delay between sending message into channel.
const msgFrequency = time.Millisecond * 250

func main() {

	var (
		readerCount = getReaderCount()
		msgCh       = make(chan string)
		doneCh      = make(chan struct{})
		wg          sync.WaitGroup
	)

	// Run asyncRead gouroutines.
	wg.Add(readerCount)
	for i := 0; i < readerCount; i++ {
		go asyncRead(&wg, i, msgCh, doneCh)
	}
	defer func() {
		// Send done signal and wait for finish gouroutines.
		close(doneCh)
		wg.Wait()

		// Close message channel.
		close(msgCh)
	}()

	// Register interrupt signal channel.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// Start infinite loop of sending messages.
	for msgNum := 0; ; msgNum++ {

		select {

		// Quit program on receive interrupt signal.
		// All gourotines will be closed in defer func.
		case <-signalCh:
			fmt.Printf("\nReceive interrupt signal\n")
			return

		// Default action - send message to message channel
		// and sleep some time.
		default:
			msgCh <- fmt.Sprint(msgNum)
			time.Sleep(msgFrequency)
		}
	}
}

// getReaderCount return reader count.
// If we have additional os.Args - try to extract reader count from args,
// otherwise scan reader counter from console.
func getReaderCount() int {

	// If len(os.Args) > 1 - try to parse second arg as number.
	if len(os.Args) > 1 {
		n, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("failed to parse reader count from args")
		}
		if n < 1 {
			log.Fatal("reader count should be more than 0")
		}
		return int(n)
	}

	// Try to get reader count from console.
	var n int
	fmt.Printf("Please, enter reader count: ")
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatal("failed to get reader count from conlose input")
	}
	if n < 1 {
		log.Fatal("reader count should be more than 0")
	}
	return n
}

// asyncRead read messages from msgCh while not receive message from doneCh.
func asyncRead(wg *sync.WaitGroup, readerIndex int, msgCh <-chan string, doneCh chan struct{}) {

	fmt.Printf("Reader %d start listen channels\n", readerIndex)

	for {
		select {
		case msg := <-msgCh:
			fmt.Printf("Reader %d receive message %s\n", readerIndex, msg)
		case <-doneCh:
			fmt.Printf("Reader %d is shutdown\n", readerIndex)
			wg.Done()
			return
		}
	}
}
