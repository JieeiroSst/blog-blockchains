package blog

import (
	"fmt"
	"github.com/JIeeiroSst/blog/x/blog/keeper"
	"github.com/JIeeiroSst/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgCreatePosts(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePosts) (*sdk.Result, error) {
	k.CreatePosts(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdatePosts(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdatePosts) (*sdk.Result, error) {
	var posts = types.Posts{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Checks that the element exists
	if !k.HasPosts(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetPostsOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPosts(ctx, posts)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeletePosts(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeletePosts) (*sdk.Result, error) {
	if !k.HasPosts(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetPostsOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.DeletePosts(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
