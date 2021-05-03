package blog

import (
	"github.com/JIeeiroSst/blog/x/blog/keeper"
	"github.com/JIeeiroSst/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the posts
	for _, elem := range genState.PostsList {
		k.SetPosts(ctx, *elem)
	}

	// Set posts count
	k.SetPostsCount(ctx, int64(len(genState.PostsList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all posts
	postsList := k.GetAllPosts(ctx)
	for _, elem := range postsList {
		elem := elem
		genesis.PostsList = append(genesis.PostsList, &elem)
	}

	return genesis
}
