// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/Toinounet21/swapalanchego/codec"
	"github.com/Toinounet21/swapalanchego/codec/linearcodec"
	"github.com/Toinounet21/swapalanchego/codec/reflectcodec"
	"github.com/Toinounet21/swapalanchego/utils/units"
)

const (
	maxPackerSize  = 1 * units.GiB // max size, in bytes, of something being marshalled by Marshal()
	maxSliceLength = 256 * 1024

	codecVersion = 0
)

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLength)
	c = codec.NewManager(maxPackerSize)
	if err := c.RegisterCodec(codecVersion, lc); err != nil {
		panic(err)
	}
}
