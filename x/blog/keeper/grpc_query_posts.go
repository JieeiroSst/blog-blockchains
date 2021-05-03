package keeper

import (
	"context"

	"github.com/JIeeiroSst/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PostsAll(c context.Context, req *types.QueryAllPostsRequest) (*types.QueryAllPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var postss []*types.Posts
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	postsStore := prefix.NewStore(store, types.KeyPrefix(types.PostsKey))

	pageRes, err := query.Paginate(postsStore, req.Pagination, func(key []byte, value []byte) error {
		var posts types.Posts
		if err := k.cdc.UnmarshalBinaryBare(value, &posts); err != nil {
			return err
		}

		postss = append(postss, &posts)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPostsResponse{Posts: postss, Pagination: pageRes}, nil
}

func (k Keeper) Posts(c context.Context, req *types.QueryGetPostsRequest) (*types.QueryGetPostsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var posts types.Posts
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostsKey+req.Id)), &posts)

	return &types.QueryGetPostsResponse{Posts: &posts}, nil
}
