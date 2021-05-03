package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreatePosts{}, "blog/CreatePosts", nil)
	cdc.RegisterConcrete(&MsgUpdatePosts{}, "blog/UpdatePosts", nil)
	cdc.RegisterConcrete(&MsgDeletePosts{}, "blog/DeletePosts", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePosts{},
		&MsgUpdatePosts{},
		&MsgDeletePosts{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
