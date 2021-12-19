package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lukeajtodd/nameservice/x/nameservice/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// See if we can get the whois from the store
	whois, isFound := k.GetWhois(ctx, msg.Name)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Name does not exist")
	}

	if msg.Creator != whois.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not the owner")
	}

	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: msg.Value,
		Price: whois.Price,
		Owner: whois.Owner,
	}

	k.SetWhois(ctx, newWhois)

	return &types.MsgSetNameResponse{}, nil
}
