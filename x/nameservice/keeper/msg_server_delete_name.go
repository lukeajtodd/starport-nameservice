package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lukeajtodd/nameservice/x/nameservice/types"
)

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, isFound := k.GetWhois(ctx, msg.Name)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Name does not exist")
	}

	if msg.Creator != whois.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.RemoveWhois(ctx, msg.Name)

	return &types.MsgDeleteNameResponse{}, nil
}
