package main

import (
	db "avalanche_indexer/db"
	"avalanche_indexer/model"
	"avalanche_indexer/rpc"
	"fmt"
	"strconv"
	"time"
)

func main() {
	db.Connect("transactions")
	duration := 300 * time.Millisecond
	number, _ := rpc.GetBlockHeight()
	model.AllInsertTransactions(number, db.TransactionCollection)
	latestBlockNumber, err := strconv.ParseInt(number[2:], 16, 64)
	if err != nil {
		fmt.Println("Invalid hexadecimal number:", err)
		return
	}

	for {
		newNumber, _ := rpc.GetBlockHeight()
		newBlockNumber, err := strconv.ParseInt(newNumber[2:], 16, 64)
		if err != nil {
			fmt.Println("Invalid hexadecimal number:", err)
			return
		}

		if latestBlockNumber < newBlockNumber {
			model.InsertTransaction(newBlockNumber, db.TransactionCollection)
			latestBlockNumber = newBlockNumber
		}
		time.Sleep(duration)
	}
}
