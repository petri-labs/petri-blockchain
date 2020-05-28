package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace               sdk.CodespaceType = ModuleName
	CodeInvalidDistribution        sdk.CodeType      = 101
	CodeInvalidShare               sdk.CodeType      = 102
	CodeInvalidSubscriptionContent sdk.CodeType      = 103
	CodeInvalidFeeContractAction   sdk.CodeType      = 104
	CodeInvalidDiscount            sdk.CodeType      = 105
	CodeInvalidDiscountRequest     sdk.CodeType      = 106
	CodeInvalidFee                 sdk.CodeType      = 107
	CodeInvalidSubscriptionAction  sdk.CodeType      = 108
	CodeInvalidId                  sdk.CodeType      = 109
)

func ErrNegativeSharePercentage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("fee distribution share percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidShare, errMsg)
}

func ErrDistributionPercentagesNot100(codespace sdk.CodespaceType, total sdk.Dec) sdk.Error {
	errMsg := fmt.Sprintf("fee distribution percentages should add up to 100, not %s", total.String())
	return sdk.NewError(codespace, CodeInvalidDistribution, errMsg)
}

func ErrInvalidSubscriptionContent(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("subscription content is invalid: %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidSubscriptionContent, errMsg)
}

func ErrFeeContractCannotBeDeauthorised(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("fee contract cannot be deauthorised")
	return sdk.NewError(codespace, CodeInvalidFeeContractAction, errMsg)
}

func ErrDiscountIDsBeSequentialFrom1(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount IDs must be sequential starting with 1")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrNegativeDiscountPercentage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrDiscountPercentageGreaterThan100(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount percentage cannot exceed 100%%")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrDiscountIdIsNotInFee(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount ID specified is not one of the fee's discounts")
	return sdk.NewError(codespace, CodeInvalidDiscountRequest, errMsg)
}

func ErrInvalidFee(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("fee invalid; %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidFee, errMsg)
}

func ErrTriedToChargeSubscriptionFeeWhenShouldnt(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("tried to charge subscription fee when shouldn't")
	return sdk.NewError(codespace, CodeInvalidSubscriptionAction, errMsg)
}

func ErrInvalidId(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidId, errMsg)
}
