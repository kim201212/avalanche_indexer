package model

import (
	"avalanche_indexer/rpc"
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

func AllInsertTransactions(number string, collection *mongo.Collection) {

	blockNumberBigInt, _ := new(big.Int).SetString(number[2:], 16)

	i := big.NewInt(0)
	for i.Cmp(blockNumberBigInt) < 0 {
		hexblock := i.Text(16)
		hexblock = "0x" + hexblock
		block, _ := rpc.GetBlock(hexblock)
		for _, transaction := range block.Result.Transactions {
			// Insert block into MongoDB
			insertResult, err := collection.InsertOne(context.TODO(), transaction)
			if err != nil {
				if mongo.IsDuplicateKeyError(err) {
					fmt.Println("Transaction with the same hash already exists")
					continue
				}
				log.Fatal(err)
			}

			fmt.Println("Inserted a single document: ", insertResult.InsertedID)
			fmt.Println("From: ", transaction.From)
			if transaction.To == "" {
				fmt.Println("To: Contract Creation Tx")
			} else {
				fmt.Println("To: ", transaction.To)
			}

		}
		i.Add(i, big.NewInt(1))
	}
}

func InsertTransaction(number int64, collection *mongo.Collection) {
	hexblock := strconv.FormatInt(int64(number), 16)
	hexblock = "0x" + hexblock
	block, _ := rpc.GetBlock(hexblock)
	for _, transaction := range block.Result.Transactions {
		// Insert block into MongoDB
		insertResult, err := collection.InsertOne(context.TODO(), transaction)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				fmt.Println("Transaction with the same hash already exists")
				continue
			}
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		fmt.Println("From: ", transaction.From)
		if transaction.To == "" {
			fmt.Println("To: Contract Creation Tx")
		} else {
			fmt.Println("To: ", transaction.To)
		}

	}
}
