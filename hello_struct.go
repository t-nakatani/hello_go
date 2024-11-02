package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Ethereum struct {
	client *ethclient.Client
}

func NewEthereum(rpcUrl string) *Ethereum {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &Ethereum{client: client}
}

func (ether *Ethereum) GetBlock() *types.Header {
	header, err := ether.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return header
}

func (ether *Ethereum) GetTransaction(txHash string) *types.Transaction {
	tx, _, err := ether.client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func main() {
	rpcUrl := "https://eth.llamarpc.com"
	ether := NewEthereum(rpcUrl)

	blockHeader := ether.GetBlock()
	fmt.Printf("Block Number: %d, Block Hash: %s\n", blockHeader.Number, blockHeader.Hash().Hex())

	txHash := "0x531726edeeb3fc7a520a4a4aac680b3a66e36e9d18eb9b3c79d96b1fb66d16a6"
	tx := ether.GetTransaction(txHash)
	fmt.Printf("tx hash: %s, chain id: %d, to: %s\n", tx.Hash().Hex(), tx.ChainId(), tx.To().Hex())
}
