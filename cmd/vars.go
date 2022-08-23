package cmd

import "github.com/ethereum/go-ethereum/crypto"

const (
	baseDenom  = "atevmos"
	ropsten    = "wss://ropsten.infura.io/ws/v3/d0383d521441488fb754735af7fe0c59"
	localhost  = "ws://localhost:8546"
	evmos_test = "wss://eth.bd.evmos.dev:8546"
)

var (
	logDelegateSig   = []byte("Stake(address,address,uint256,uint256)")
	logUndelegateSig = []byte("Unstake(address,address,uint256,uint256)")

	logDelegateSigHash   = crypto.Keccak256Hash(logDelegateSig)
	logUndelegateSigHash = crypto.Keccak256Hash(logUndelegateSig)
)
