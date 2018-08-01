package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal("Client: ", err)
	}

	contractAddress := common.HexToAddress("0x0b68107DBc64f0C7e3cA961c7C43756f1c2bec53")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("SubscribeFilterLogs: ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// fmt.Println(vLog) // pointer to event log
			contractAbi, err := abi.JSON(strings.NewReader(string(TokenABI)))
			if err != nil {
				log.Fatal("abi.JSON: ", err)
			}

			event := struct {
				Sender common.Address
				Value  *big.Int
			}{}

			err = contractAbi.Unpack(&event, "NewTransaction", vLog.Data)
			if err != nil {
				log.Fatal("Unpack: ", err)
			}

			fmt.Println(string(event.Sender.Hex()))   // foo
			fmt.Println(string(event.Value.Uint64())) // bar
			// fmt.Println(string(hex.Dump(vLog.Data)))

			// fmt.Println(vLog.Address.String())
			// fmt.Println(vLog.TxHash.String())

			txHash := common.HexToHash(vLog.TxHash.Hex())
			tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("String: ", string(tx.Data()))
			fmt.Println("BytesToHash: ", common.BytesToHash(tx.Data()).Hex())
			fmt.Println("BytesToAddress: ", common.BytesToAddress(tx.Data()).Hex())
			fmt.Println("EncodeToString: ", hex.EncodeToString(tx.Data()))

			// fmt.Println("To: ", tx.To().Hex())
			// fmt.Println(tx.Value())
			// fmt.Println(tx.Gas())
			// fmt.Println(tx.Hash().Hex())
			fmt.Println("Pending: ", isPending)
			// fmt.Println("-------------------------------------------------------------------")
		}
	}
}
