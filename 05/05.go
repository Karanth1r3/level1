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

	wg.Add(2)                    // in these conditions the number of goroutines is known already
	go Write(&wg, msgCh, doneCh) // reminds conveyour, one by one from wr to rd
	go Read(&wg, msgCh, doneCh)

	time.Sleep(delay)
	close(doneCh)
	wg.Wait()

	close(msgCh)
}

// Write
func Write(wg *sync.WaitGroup, msgChan chan string, doneCh chan struct{}) {
	for seqNum := 0; ; seqNum++ { //endless loop
		select {
		case <-doneCh: // if receiving an data from done channel, end goroutine & decrease wg counter
			wg.Done()
			return
		default:
			msgChan <- strconv.Itoa(seqNum) // otherwise allways will be sending data to msgChan
		}
	}
}

// Read
func Read(wg *sync.WaitGroup, msgChan chan string, doneCh chan struct{}) {
	for {
		select {
		case <-doneCh: // the same as with Write()
			wg.Done()
			return
		case msg := <-msgChan: // otherwise will be always waiting for input from msgChan
			fmt.Println("Message", msg, "processed")
		}
	}
}
