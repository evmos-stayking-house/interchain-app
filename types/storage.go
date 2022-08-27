package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/tendermint/tendermint/libs/json"
	"math/big"
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

	// record of undelegations
	LastUndelegationTime time.Time
	PendingUndelegations *big.Int
	PendingDelegations   *big.Int
}

func Init(cliCtx client.Context) error {
	initialStorage := Storage{
		DelegatorAddress:            "",
		UndelegationProcessedHeight: 0,
		DelegationProcessedHeight:   0,
		DelegationProcessed:         []string{},
		UndelegationProcessed:       []string{},
		LastUndelegationTime:        time.Time{},
		PendingUndelegations:        big.NewInt(0),
		PendingDelegations:          big.NewInt(0),
	}
	return WriteStore(cliCtx, initialStorage)
}

func WriteStore(cliCtx client.Context, s Storage) error {
	homeDir := cliCtx.HomeDir

	sBz, err := json.Marshal(s)
	if err != nil {
		return err
	}
	filePath := filepath.Join(homeDir, "sched_worker_data.json")
	return os.WriteFile(filePath, sBz, 0644)
}

func ReadStore(cliCtx client.Context) (Storage, error) {
	homeDir := cliCtx.HomeDir
	filePath := filepath.Join(homeDir, "sched_worker_data.json")
	bz, err := os.ReadFile(filePath)
	if err != nil {
		return Storage{}, err
	}

	var res Storage
	if err := json.Unmarshal(bz, &res); err != nil {
		return Storage{}, err
	}

	return res, nil
}
