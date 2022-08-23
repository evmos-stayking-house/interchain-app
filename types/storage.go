package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/tendermint/tendermint/libs/json"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	// The latest heights that are processed before
	DelegatorAddress string

	// track record of processing delegations & undelegations
	UndelegationProcessedHeight uint64
	DelegationProcessedHeight   uint64
	DelegationProcessed         []string
	UndelegationProcessed       []string

	// record of last undelegation
	LastUndelegationTime time.Time
}

func Init(cliCtx client.Context) error {
	homeDir := cliCtx.HomeDir

	initialStorage := Storage{
		DelegatorAddress:            "",
		UndelegationProcessedHeight: 0,
		DelegationProcessedHeight:   0,
		DelegationProcessed:         []string{},
		UndelegationProcessed:       []string{},
		LastUndelegationTime:        time.Time{},
	}
	initialStorageBz, err := json.Marshal(initialStorage)
	if err != nil {
		return err
	}

	filePath := filepath.Join(homeDir, "sched_worker_data.json")
	return os.WriteFile(filePath, initialStorageBz, 0644)
}

func Store(cliCtx client.Context, s Storage) error {
	homeDir := cliCtx.HomeDir

	sBz, err := json.Marshal(s)
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, "sched_worker_data.json")
	return os.WriteFile(filePath, sBz, 0644)
}

func ReadStore(cliCtx client.Context) error {

}
