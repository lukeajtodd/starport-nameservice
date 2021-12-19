package keeper

import (
	"github.com/lukeajtodd/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
