package keeper

import (
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ty "github.com/ixofoundation/ixo-blockchain/x/project/internal/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/abci/types"
)

const (
	QueryProjectDoc      = "queryProjectDoc"
	QueryProjectAccounts = "queryProjectAccounts"
	QueryProjectTx       = "queryProjectTx"
	QueryParams          = "queryParams"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req types.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryProjectDoc:
			return queryProjectDoc(ctx, path[1:], k)
		case QueryProjectAccounts:
			return queryProjectAccounts(ctx, path[1:], k)
		case QueryProjectTx:
			return queryProjectTx(ctx, path[1:], k)
		case QueryParams:
			return queryParams(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown project query endpoint")
		}
	}
}

func queryProjectDoc(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	storedDoc, err := k.GetProjectDoc(ctx, path[0])
	if err != nil {
		return nil, err
	}

	res, errRes := codec.MarshalJSONIndent(k.cdc, storedDoc)
	if errRes != nil {
		return nil, sdkerrors.Wrap(ty.ErrInternal, "failed to marshal data")
	}

	return res, nil
}

func queryProjectAccounts(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {

	resp := k.GetAccountMap(ctx, path[0])
	res, err := json.Marshal(resp)
	if err != nil {
		return nil, sdkerrors.Wrap(ty.ErrInternal, "failed to marshal data")
	}

	return res, nil
}

func queryProjectTx(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	info, err := k.GetProjectWithdrawalTransactions(ctx, path[0])
	if err != nil {
		return nil, err
	}

	res, err2 := codec.MarshalJSONIndent(k.cdc, info)
	if err2 != nil {
		return nil, sdkerrors.Wrap(ty.ErrInternal, "failed to marshal data")
	}

	return res, nil
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(ty.ErrInternal, "failed to marshal data")
	}

	return res, nil
}
