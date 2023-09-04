package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/gmm module sentinel errors
var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrAlreadyCreatedPool       = sdkerrors.Register(ModuleName, 1101, "already exist pool in the system")
	ErrNotEnoughBalance         = sdkerrors.Register(ModuleName, 1102, "not enough balance")
	ErrInvalidPoolType          = sdkerrors.Register(ModuleName, 1103, "invalid pool type")
	ErrInvalidNumOfAssets       = sdkerrors.Register(ModuleName, 1104, "invalid number of assets")
	ErrNotFoundAssetInPool      = sdkerrors.Register(ModuleName, 1105, "did not find asset")
	ErrNotFoundAPool            = sdkerrors.Register(ModuleName, 1106, "did not find pool")
	ErrOverflowShareAmount      = sdkerrors.Register(ModuleName, 1107, "share amount is over than pool total supply")
	ErrInvalidAddress           = sdkerrors.Register(ModuleName, 1109, "invalid betch32 address")
	ErrInvalidPoolParams        = sdkerrors.Register(ModuleName, 1110, "invalid pool params")
	ErrEmptyLiquidity           = sdkerrors.Register(ModuleName, 1111, "empty pool liquidity")
	ErrInvalidLiquidityInLength = sdkerrors.Register(ModuleName, 1112, "invalid pool liquidity in length")
	ErrInvalidPoolID            = sdkerrors.Register(ModuleName, 1113, "invalid poolID")
	ErrInvalidLiquidityAmount   = sdkerrors.Register(ModuleName, 1114, "invalid liquidity amount")
	ErrInvalidTokenAmount       = sdkerrors.Register(ModuleName, 1115, "invalid token amount")
	ErrEmptyDenom               = sdkerrors.Register(ModuleName, 1116, "empty denom")
)