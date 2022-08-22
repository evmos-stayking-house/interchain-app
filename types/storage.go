package types

import "math/big"

type Storage struct {
	// The latest heights that are processed before
	UndelegationProcessedHeight *big.Int
	DelegationProcessedHeight   *big.Int
	DelegatorAddress            string
}

func Store(s Storage) error {
	return nil
}
