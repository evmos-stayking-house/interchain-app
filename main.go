package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Block struct {
	Number string
}

func main() {
	wsclient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/d0383d521441488fb754735af7fe0c59")
	if err != nil {
		log.Fatal(err)
	}

	//contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{}

	logs := make(chan types.Log)

	sub, err := wsclient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
