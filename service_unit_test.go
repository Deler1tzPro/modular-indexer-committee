package main

import (
	"log"
	"encoding/base64"
	"time"
	"testing"

	"github.com/RiemaLabs/indexer-committee/ord/getter"
	"github.com/RiemaLabs/indexer-committee/ord/stateless"
)

func TestService(t *testing.T) {
	getter, _ := loadMain()
	queue := loadCatchUp()

	startTime := time.Now()
	loadService(getter, queue, 3) // partially update, some history still remain
	elapsed := time.Since(startTime)
	log.Printf("Using time %s\n", elapsed)

	startTime = time.Now()
	loadService(getter, queue, 10) // all update, no historical record stays
	elapsed = time.Since(startTime)
	log.Printf("Using time %s\n", elapsed)
}

func loadService(getter *getter.OPIOrdGetter, queue *stateless.Queue, upHeight uint) {
	curHeight := queue.LastestHeight()
	latestHeight := curHeight + upHeight

	if curHeight < latestHeight {
		queue.Lock()
		err := queue.Update(getter, latestHeight)
		queue.Unlock()
		if err != nil {
			log.Fatalf("Failed to update the queue: %v", err)
		}
	}

	// Hash and Height logging
	log.Printf("With header's height at %d, and header's hash to be %s", queue.Header.Height, queue.Header.Hash)

	// Commitment logging
	bytes := queue.Header.Root.Commit().Bytes()
	commitment := base64.StdEncoding.EncodeToString(bytes[:])
	log.Printf("Header's commitment is %s", commitment)
	queue.Println()
}