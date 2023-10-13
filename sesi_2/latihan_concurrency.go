// PINGPONG APPS
// 2 player => 2 goroutine

// kondisi kalah : saat flag/counter/random number, habis dibagi 11 (random % 11 == 0 THAN break)

// Contoh :
// Player A = Hit 1 // counter 8
// Player B = Hit 2 // counter 3
// Player A = Hit 3 // counter 24
// Player B = Hit 4 // counter 33

// Player B kalah, total hit : 4, kalah di nomor 33

// Player A = Hit 1 // counter 8
// Player B = Hit 2 // counter 9
// Player A = Hit 3 // counter 22

// Player A kalah, total hit : 3, kalah di nomor 22

// Requirement :
// - Struct Player {
// Name string
// Hit int
// }

// a := Player{Name: A, Hit:0}
// b := Player{Name: B, Hit:0}

// a.Hit++
// b.Hit++

package main

import (
	"fmt"
	"sync"
	"time"
)

// CLUE
// func play(numchan chan int){
// 	for {
// 	  // get valu
// 	  num := <-numchan
// 	  v := rand.Intn(100-1) + 1
// 	  log.Println(num)
// 	  // set value again
// 	  numchan <- v
// 	}
//   }

func main() {
	var wg sync.WaitGroup
	ball := make(chan string)

	wg.Add(2)

	// Ping goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			ball <- "Ping"
			time.Sleep(time.Second)
			msg := <-ball
			fmt.Println(msg)
		}
	}()

	// Pong goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			msg := <-ball
			fmt.Println(msg)
			ball <- "Pong"
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
