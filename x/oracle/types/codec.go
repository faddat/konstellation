package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgSetExchangeRate{}, "oracle/SetExchangeRate", nil)
	cdc.RegisterConcrete(MsgSetExchangeRates{}, "oracle/SetExchangeRates", nil)
	cdc.RegisterConcrete(MsgDeleteExchangeRate{}, "oracle/DeleteExchangeRate", nil)
	cdc.RegisterConcrete(MsgDeleteExchangeRates{}, "oracle/DeleteExchangeRates", nil)
	cdc.RegisterConcrete(MsgSetAdminAddr{}, "oracle/MsgSetAdminAddr", nil)
}

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSetExchangeRate{},
		&MsgSetExchangeRates{},
		&MsgDeleteExchangeRate{},
		&MsgDeleteExchangeRates{},
		&MsgSetAdminAddr{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
