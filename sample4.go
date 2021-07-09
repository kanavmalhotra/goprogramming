package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sample(flag chan bool, wg *sync.WaitGroup) {
// 	for {
// 		select {
// 		case msg := <-flag:
// 			fmt.Printf("exiting %v", msg)

// 			wg.Done()
// 		default:
// 			fmt.Println("Hello World")
// 		}

// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	flag := make(chan bool, 1)

// 	wg.Add(1)
// 	go sample(flag, &wg)

// 	time.Sleep(3 * time.Second)
// 	flag <- true

// 	wg.Wait()
// 	fmt.Println("quit")
// }
