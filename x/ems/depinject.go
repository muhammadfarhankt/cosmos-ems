// package module

// import (
// 	"os"

// 	"github.com/cosmos/cosmos-sdk/codec"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"

// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
// 	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

// 	"cosmossdk.io/core/address"
// 	"cosmossdk.io/core/appmodule"
// 	"cosmossdk.io/core/store"
// 	"cosmossdk.io/depinject"
// 	"cosmossdk.io/log"

// 	modulev1 "github.com/rollchains/hackmoschain/api/ems/module/v1"
// 	"github.com/rollchains/hackmoschain/x/ems/keeper"
// )

// var _ appmodule.AppModule = AppModule{}

// // IsOnePerModuleType implements the depinject.OnePerModuleType interface.
// func (am AppModule) IsOnePerModuleType() {}

// // IsAppModule implements the appmodule.AppModule interface.
// func (am AppModule) IsAppModule() {}

// func init() {
// 	appmodule.Register(
// 		&modulev1.Module{},
// 		appmodule.Provide(ProvideModule),
// 	)
// }

// type ModuleInputs struct {
// 	depinject.In

// 	Cdc          codec.Codec
// 	StoreService store.KVStoreService
// 	AddressCodec address.Codec

// 	StakingKeeper  stakingkeeper.Keeper
// 	SlashingKeeper slashingkeeper.Keeper
// }

// type ModuleOutputs struct {
// 	depinject.Out

// 	Module appmodule.AppModule
// 	Keeper keeper.Keeper
// }

// func ProvideModule(in ModuleInputs) ModuleOutputs {
// 	govAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

// 	k := keeper.NewKeeper(in.Cdc, in.StoreService, log.NewLogger(os.Stderr), govAddr)
// 	m := NewAppModule(in.Cdc, k)

// 	return ModuleOutputs{Module: m, Keeper: k, Out: depinject.Out{}}
// }

// // x/ems/depinject.go
// package module

// import (
// 	"os"

// 	"github.com/cosmos/cosmos-sdk/codec"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

// 	"cosmossdk.io/core/address"
// 	"cosmossdk.io/core/appmodule"
// 	"cosmossdk.io/core/store"
// 	"cosmossdk.io/depinject"
// 	"cosmossdk.io/log"

// 	modulev1 "github.com/rollchains/hackmoschain/api/ems/module/v1"
// 	"github.com/rollchains/hackmoschain/x/ems/keeper"

// 	nftKeeper "cosmossdk.io/x/nft/keeper" // Import nft keeper
// )

// var _ appmodule.AppModule = AppModule{}

// // IsOnePerModuleType implements the depinject.OnePerModuleType interface.
// func (am AppModule) IsOnePerModuleType() {}

// // IsAppModule implements the appmodule.AppModule interface.
// func (am AppModule) IsAppModule() {}

// func init() {
// 	appmodule.Register(
// 		&modulev1.Module{},
// 		appmodule.Provide(ProvideModule),
// 	)
// }

// type ModuleInputs struct {
// 	depinject.In

// 	Config       *modulev1.Config
// 	Cdc          codec.Codec
// 	StoreService store.KVStoreService
// 	AddressCodec address.Codec
// 	NftKeeper    nftKeeper.Keeper // Add nft keeper to inputs
// }

// type ModuleOutputs struct {
// 	depinject.Out

// 	Module appmodule.AppModule
// 	Keeper keeper.Keeper
// }

// func ProvideModule(in ModuleInputs) ModuleOutputs {
// 	govAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()
// 	// Pass address codec and nft keeper to the keeper
// 	k := keeper.NewKeeper(in.Cdc, in.AddressCodec, in.StoreService, log.NewLogger(os.Stderr), govAddr, in.NftKeeper)
// 	m := NewAppModule(in.Cdc, k)

// 	return ModuleOutputs{Module: m, Keeper: k}
// }

// x/ems/depinject.go
package module

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	modulev1 "github.com/rollchains/hackmoschain/api/ems/module/v1"
	"github.com/rollchains/hackmoschain/x/ems/keeper"

	nftKeeper "cosmossdk.io/x/nft/keeper"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	AddressCodec address.Codec
	NftKeeper    nftKeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	govAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	k := keeper.NewKeeper(in.Cdc, in.AddressCodec, in.StoreService, log.NewLogger(os.Stderr), govAddr, in.NftKeeper)
	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{Module: m, Keeper: k}
}
