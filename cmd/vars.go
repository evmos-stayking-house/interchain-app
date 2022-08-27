package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	baseDenom                   = "atevmos"
	ropsten                     = "wss://ropsten.infura.io/ws/v3/d0383d521441488fb754735af7fe0c59"
	localhost                   = "ws://localhost:8546"
	evmos_test                  = "http://bd-evmos-testnet-state-sync-node-01.bdnodes.net:26657"
	supplyUnbondedTokenFuncName = "supplyUnbondedToken"
	getAccruedValueFuncName     = "getAccruedValue"
	accrueFuncName              = "accrue"
)

var (
	logDelegateSig   = []byte("Stake(address,address,uint256,uint256)")
	logUndelegateSig = []byte("Unstake(address,address,uint256,uint256)")

	LogDelegateSigHash   = crypto.Keccak256Hash(logDelegateSig)
	LogUndelegateSigHash = crypto.Keccak256Hash(logUndelegateSig)

	reservedDec     = sdk.NewDecWithPrec(2, 1) // 20%
	distributionDec = sdk.NewDecWithPrec(5, 1) // 50%

)
