package main

import (
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/evmos-stayking-house/scheduled-worker-golang/cmd"
	"os"
)

type Block struct {
	Number string
}

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
