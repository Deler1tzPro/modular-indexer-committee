package main

import (
	"encoding/base64"
	"log"
	"testing"
	"time"

	"github.com/RiemaLabs/modular-indexer-committee/ord/stateless"
)

func Test_CatchupStage(t *testing.T) {
	log.Println("Test_CatchupStage")
	var catchupHeight uint = 780000
	ordGetterTest, arguments := loadMain(782000)
	startTime := time.Now()
	queue, _ := CatchupStage(ordGetterTest, &arguments, stateless.BRC20StartHeight-1, catchupHeight)
	if queue.Header.Height != catchupHeight {
		log.Println("Queue header not updated correctly")
	}
	ordGetterTest.SetLatestBlockHeight(catchupHeight)
	elapsed := time.Since(startTime)
	elapsedSeconds := float64(elapsed) / float64(time.Second)
	averageTime := elapsedSeconds / float64(catchupHeight-stateless.BRC20StartHeight)
	log.Printf("Successfully Updating From %d To %d", stateless.BRC20StartHeight, catchupHeight)
	log.Printf("Using Time %s, And %f Per Block on Average During CatchUp Stage", elapsed, averageTime)

	// Commitment logging
	bytes := queue.Header.Root.VerkleTree.Commit().Bytes()
	commitment := base64.StdEncoding.EncodeToString(bytes[:])
	log.Printf("Header's Commitment Is %s", commitment)
	stateless.CleanPath(stateless.VerkleDataPath)
}
