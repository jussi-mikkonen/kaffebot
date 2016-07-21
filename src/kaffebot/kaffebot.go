package main

import (
	//"fmt"
	"log"
	"os"
	//"net/http"
	"sync"
	"kaffebot/kahvikanava"
)



func main() {
	var logger = log.New(os.Stdout, "MAIN: ", 0)

	logger.Println("Starting bots")

	var wg sync.WaitGroup

	wg.Add(1)
	go kahvikanava.KaffeKanavaBot(&wg) // See kahvikanava.go

	// To add different flows, implement something similar to kaffeKanavaBot and call:
	// wg.Add(1)
	// go yourFlowBot(&wg)

	logger.Println("All bots launched. Waiting for their termination.")
	wg.Wait()
}




