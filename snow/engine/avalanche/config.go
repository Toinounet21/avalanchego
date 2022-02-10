// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/Toinounet21/crabalanchego/snow"
	"github.com/Toinounet21/crabalanchego/snow/consensus/avalanche"
	"github.com/Toinounet21/crabalanchego/snow/engine/avalanche/vertex"
	"github.com/Toinounet21/crabalanchego/snow/engine/common"
	"github.com/Toinounet21/crabalanchego/snow/validators"
)

// Config wraps all the parameters needed for an avalanche engine
type Config struct {
	Ctx *snow.ConsensusContext
	common.AllGetsServer
	VM         vertex.DAGVM
	Manager    vertex.Manager
	Sender     common.Sender
	Validators validators.Set

	Params    avalanche.Parameters
	Consensus avalanche.Consensus
}
