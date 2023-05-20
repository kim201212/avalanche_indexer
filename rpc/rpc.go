package rpc

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func getBlock() {
	client, err := ethclient.Dial("http://aops-custom-202305-2crvsg-nlb-1d600174371701f9.elb.ap-northeast-2.amazonaws.com:9650/ext/bc/XpX1yGquejU5cma1qERzkHKDh4fsPKs4NttnS1tErigPzugx5/rpc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	blockNumber := big.NewInt(1) // replace with your block number
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("Failed to get block: %v", err)
	}
	fmt.Println(block)
	fmt.Println(block.Number().Uint64())     // block number
	fmt.Println(block.Time())                // block time
	fmt.Println(block.Difficulty().Uint64()) // block difficulty
	fmt.Println(block.Hash().Hex())          // block hash
	fmt.Println(len(block.Transactions()))   // number of transactions in the block
	fmt.Println(block.Transactions())        // number of transactions in the block
}
