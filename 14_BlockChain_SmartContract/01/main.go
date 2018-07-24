package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
)

const key = `paste the contents of your *testnet* key json here`

func main() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding

	client, err := ethclient.Dial("https://mainnet.infura.io/e355610271914b8eb682258c2906ecd8")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ctx := context.Background()
	txHash := common.HexToHash("0x6316a1ccfe3b32b17f07cb7a92a1c71591fda66b2b2de43f90eb7d3cb30475d2")

	tx, pending, _ := client.TransactionByHash(ctx, txHash)
	if !pending {
		fmt.Println(tx)
	}
	blance, err := client.BalanceAt(ctx, common.HexToAddress("0x190a2409fc6434483d4c2cab804e75e3bc5ebfa6"), nil)
	if err != nil {
		log.Fatalf("Failed to BalanceAt the Ethereum client: %v", err)
	}
	_ = blance
	// fmt.Println(blance)
	fmt.Println(ctx)

	filter := ethereum.FilterQuery{}
	filter.Addresses = make([]common.Address, 0)
	filter.Addresses = append(filter.Addresses, common.HexToAddress("0x417bdc58ef9a3d7de04a66ab84ed13048d9a82bb"))
	// filter.FromBlock = big.NewInt(1000000)

	// typelog, err := client.FilterLogs(ctx, filter)
	// if err != nil {
	// 	log.Fatalf("Failed to FilterLogs the Ethereum client: %v", err)
	// }
	// fmt.Println("Println")
	// fmt.Println(len(typelog))

	/*
		ethereumLogsCh := make(chan types.Log)

		subscript, err := client.SubscribeFilterLogs(ctx, filter, ethereumLogsCh)
		if err != nil {
			log.Fatalf("Failed to SubscribeFilterLogs the Ethereum client: %v", err)
		}
		fmt.Println(subscript)
	*/
	var x bind.Lang
	_ = x

	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		b, err := json.Marshal(client)
		if err != nil {
			w.Write([]byte(`{"error":"Error roi"}`))
		}
		w.Write(b)
	})
	http.ListenAndServe(":8888", router)
}

//https://mainnet.infura.io/e355610271914b8eb682258c2906ecd8
