package keeper

import (
	"github.com/JIeeiroSst/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// GetPostsCount get the total number of posts
func (k Keeper) GetPostsCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsCountKey))
	byteKey := types.KeyPrefix(types.PostsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPostsCount set the total number of posts
func (k Keeper) SetPostsCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsCountKey))
	byteKey := types.KeyPrefix(types.PostsCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreatePosts creates a posts with a new id and update the count
func (k Keeper) CreatePosts(ctx sdk.Context, msg types.MsgCreatePosts) {
	// Create the posts
	count := k.GetPostsCount(ctx)
	var posts = types.Posts{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Title:   msg.Title,
		Body:    msg.Body,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	key := types.KeyPrefix(types.PostsKey + posts.Id)
	value := k.cdc.MustMarshalBinaryBare(&posts)
	store.Set(key, value)

	// Update posts count
	k.SetPostsCount(ctx, count+1)
}

// SetPosts set a specific posts in the store
func (k Keeper) SetPosts(ctx sdk.Context, posts types.Posts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	b := k.cdc.MustMarshalBinaryBare(&posts)
	store.Set(types.KeyPrefix(types.PostsKey+posts.Id), b)
}

// GetPosts returns a posts from its id
func (k Keeper) GetPosts(ctx sdk.Context, key string) types.Posts {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	var posts types.Posts
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostsKey+key)), &posts)
	return posts
}

// HasPosts checks if the posts exists
func (k Keeper) HasPosts(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	return store.Has(types.KeyPrefix(types.PostsKey + id))
}

// GetPostsOwner returns the creator of the posts
func (k Keeper) GetPostsOwner(ctx sdk.Context, key string) string {
	return k.GetPosts(ctx, key).Creator
}

// DeletePosts deletes a posts
func (k Keeper) DeletePosts(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	store.Delete(types.KeyPrefix(types.PostsKey + key))
}

// GetAllPosts returns all posts
func (k Keeper) GetAllPosts(ctx sdk.Context) (msgs []types.Posts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostsKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PostsKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Posts
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
