package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const delay = 3 * time.Second

func main() {

	var (
		msgCh  = make(chan string)
		doneCh = make(chan struct{})
		wg     sync.WaitGroup
	)

	wg.Add(2)
	go Write(&wg, msgCh, doneCh)
	go Read(&wg, msgCh, doneCh)

	time.Sleep(delay)
	close(doneCh)
	wg.Wait()

	close(msgCh)
}

// Write
func Write(wg *sync.WaitGroup, msgChan chan string, doneCh chan struct{}) {
	for seqNum := 0; ; seqNum++ {
		select {
		case <-doneCh:
			wg.Done()
			return
		default:
			msgChan <- strconv.Itoa(seqNum)
		}
	}
}

// Read
func Read(wg *sync.WaitGroup, msgChan chan string, doneCh chan struct{}) {
	for seqNum := 0; ; seqNum++ {
		select {
		case <-doneCh:
			wg.Done()
			return
		case msg := <-msgChan:
			fmt.Println("Message", msg, "processed")
		}
	}
}
