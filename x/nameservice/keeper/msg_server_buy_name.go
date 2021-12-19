package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lukeajtodd/nameservice/x/nameservice/types"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// See if we can get the whois from the store
	whois, isFound := k.GetWhois(ctx, msg.Name)

	// If the name doesn't have an owner, set a price to be bought at
	minPrice := sdk.Coins{
		sdk.NewInt64Coin("token", 10),
	}

	// Convert the price and bid into tokens
	price, _ := sdk.ParseCoinsNormalized(whois.Price)
	bid, _ := sdk.ParseCoinsNormalized(msg.Bid)

	// Convert the owner and buyer into sdk AccAddress
	owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	if isFound {
		if price.IsAllGT(bid) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough.")
		}

		k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
	} else {
		if minPrice.IsAllGT(bid) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is lower than minimum price.")
		}

		k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
	}

	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: whois.Value,
		Price: bid.String(),
		Owner: buyer.String(),
	}

	k.SetWhois(ctx, newWhois)

	return &types.MsgBuyNameResponse{}, nil
}
