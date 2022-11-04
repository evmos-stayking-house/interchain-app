package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

type DelegationChange struct {
	Delegator common.Address
	Amount    *big.Int
}

type SafeTotal struct {
	sync.Mutex
	Amount *big.Int
}

func (t *SafeTotal) Add(n *big.Int) {
	t.Lock()
	defer t.Unlock()
	t.Amount.Add(t.Amount, n)
}

func (t *SafeTotal) Clear() *big.Int {
	t.Lock()
	defer t.Unlock()
	res := t.Amount
	t.Amount = big.NewInt(0)
	return res
}
