package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/yield module sentinel errors
var (
	ErrSenderAddressNotAuthorized = errorsmod.Register(ModuleName, 1000, "sender address not authorized")
	ErrInvalidHeader              = errorsmod.Register(ModuleName, 1100, "invalid block header")
	ErrReorgFailed                = errorsmod.Register(ModuleName, 1101, "failed to reorg chain")
	ErrForkedBlockHeader          = errorsmod.Register(ModuleName, 1102, "Invalid forked block header")

	ErrInvalidSenders = errorsmod.Register(ModuleName, 2100, "invalid allowed senders")

	ErrInvalidBtcTransaction    = errorsmod.Register(ModuleName, 3100, "invalid bitcoin transaction")
	ErrBlockNotFound            = errorsmod.Register(ModuleName, 3101, "block not found")
	ErrTransactionNotIncluded   = errorsmod.Register(ModuleName, 3101, "transaction not included in block")
	ErrNotConfirmed             = errorsmod.Register(ModuleName, 3200, "transaction not confirmed")
	ErrExceedMaxAcceptanceDepth = errorsmod.Register(ModuleName, 3201, "exceed max acceptance block depth")
)
