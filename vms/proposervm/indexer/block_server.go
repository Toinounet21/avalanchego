// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"github.com/Toinounet21/swapalanchego/database/versiondb"
	"github.com/Toinounet21/swapalanchego/ids"
	"github.com/Toinounet21/swapalanchego/snow/consensus/snowman"
)

// BlockServer represents all requests heightIndexer can issue
// against ProposerVM. All methods must be thread-safe.
type BlockServer interface {
	versiondb.Commitable

	// Note: this is a contention heavy call that should be avoided
	// for frequent/repeated indexer ops
	GetFullPostForkBlock(blkID ids.ID) (snowman.Block, error)
}
