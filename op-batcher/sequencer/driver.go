package sequencer

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/db"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type Config struct {
	Log  log.Logger
	Name string

	// API to submit txs to
	L1Client *ethclient.Client

	// API to hit for batch data
	L2Client *ethclient.Client

	// Limit the size of txs
	MinL1TxSize uint64
	MaxL1TxSize uint64

	// Where to send the batch txs to.
	BatchInboxAddress common.Address

	// Persists progress of submitting block data, to avoid redoing any work
	HistoryDB db.HistoryDatabase

	// The batcher can decide to set it shorter than the actual timeout,
	//  since submitting continued channel data to L1 is not instantaneous.
	//  It's not worth it to work with nearly timed-out channels.
	ChannelTimeout uint64

	// Chain ID of the L1 chain to submit txs to.
	ChainID *big.Int

	// Private key to sign batch txs with
	PrivKey *ecdsa.PrivateKey

	PollInterval time.Duration
}
