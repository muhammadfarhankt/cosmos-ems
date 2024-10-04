// package keeper

// import (
// 	"context"

// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/rollchains/hackmoschain/x/ems/types"
// )

// var _ types.QueryServer = Querier{}

// type Querier struct {
// 	Keeper
// }

// func NewQuerier(keeper Keeper) Querier {
// 	return Querier{Keeper: keeper}
// }

// func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	p, err := k.Keeper.Params.Get(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &types.QueryParamsResponse{Params: &p}, nil
// }

// // GetEvent implements types.QueryServer.
// func (k Querier) GetEvent(goCtx context.Context, req *types.QueryGetEventRequest) (*types.QueryGetEventResponse, error) {
// 	// v, err := k.Keeper.EventMapping.Get(goCtx, req.Organizer)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return &types.QueryGetEventResponse{
// 	// 	Name: v,
// 	// }, nil
// 	event, err := k.Keeper.Events.Get(goCtx, req.Name) // Retrieve Event struct
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &types.QueryGetEventResponse{
// 		Organizer:        event.Organizer,
// 		Name:             event.Name,
// 		Description:      event.Description,
// 		Location:         event.Location,
// 		Date:             event.Date,
// 		Time:             event.Time,
// 		TicketPrice:      event.TicketPrice,
// 		NumTickets:       event.NumTickets,
// 		TicketCategories: event.TicketCategories,
// 	}, nil
// }

// x/ems/keeper/query_server.go
package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rollchains/hackmoschain/x/ems/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

// GetEvent implements types.QueryServer.
func (k Querier) GetEvent(goCtx context.Context, req *types.QueryGetEventRequest) (*types.QueryGetEventResponse, error) {
	event, err := k.Keeper.GetEvent(goCtx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetEventResponse{
		Event: event,
	}, nil

}
