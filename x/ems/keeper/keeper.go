// x/ems/keeper/keeper.go
package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	apiv1 "github.com/rollchains/hackmoschain/api/ems/v1"
	"github.com/rollchains/hackmoschain/x/ems/types"

	nft "cosmossdk.io/x/nft"
	nftKeeper "cosmossdk.io/x/nft/keeper"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	logger log.Logger

	// state management
	Schema collections.Schema
	Params collections.Item[types.Params]
	OrmDB  apiv1.StateStore

	authority string

	Events collections.Map[string, types.Event]

	Nftkeeper nftKeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	nftKeeper nftKeeper.Keeper,
) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	sb := collections.NewSchemaBuilder(storeService)

	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}

	db, err := ormdb.NewModuleDB(&types.ORMModuleSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		logger:       logger,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:  store,

		authority: authority,

		Events: collections.NewMap(sb, collections.NewPrefix(1), "events", collections.StringKey, codec.CollValue[types.Event](cdc)),

		Nftkeeper: nftKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) Logger() log.Logger {
	return k.logger
}

// InitGenesis initializes the module's state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return k.Params.Set(ctx, data.Params)
}

// ExportGenesis exports the module's state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params: params,
	}
}

func (k Keeper) CreateEvent(ctx context.Context, msg *types.MsgCreateEvent) error {
	has, err := k.Events.Has(ctx, msg.Id)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("event already exists: %s", msg.Id)
	}

	err = k.Events.Set(ctx, msg.Id, types.Event{
		Name:             msg.Name,
		Organizers:       []string{msg.Organizer},
		Description:      msg.Description,
		Location:         msg.Location,
		Date:             msg.Date,
		Time:             msg.Time,
		TicketPrice:      msg.TicketPrice,
		NumTickets:       msg.NumTickets,
		TicketCategories: msg.TicketCategories,
	})
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetEvent(ctx context.Context, id string) (*types.Event, error) {
	event, err := k.Events.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (k Keeper) RemoveEvent(ctx context.Context, id string) error {
	err := k.Events.Remove(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) MintEventNFT(ctx context.Context, senderAddr sdk.AccAddress, receiverAddr sdk.AccAddress, id string) error {
	nftId := receiverAddr.String() + "-" + id

	event, err := k.GetEvent(ctx, id)
	if err != nil {
		return err
	}

	// if !k.Nftkeeper.HasClass(ctx, id) {
	// 	k.Nftkeeper.SaveClass(ctx, nft.Class{Id: id, Name: event.Name, Symbol: event.Name, Description: event.Description, URI: "https://example.com/nft/" + id}) // Updated with event details
	// }

	if !k.Nftkeeper.HasClass(ctx, id) {
		k.Nftkeeper.SaveClass(ctx, nft.Class{Id: id, Name: event.Name, Symbol: event.Name, Description: event.Description})
	}

	if k.Nftkeeper.HasNFT(ctx, id, nftId) {
		return nil
	}

	// return k.Nftkeeper.Mint(ctx, nft.NFT{
	// 	ClassId: id,
	// 	Id:      nftId,
	// 	URI:     "https://example.com/nft/" + nftId, // URI for the NFT
	// }, receiverAddr)
	return k.Nftkeeper.Mint(ctx, nft.NFT{
		ClassId: id,
		Id:      nftId,
	}, receiverAddr)
}
