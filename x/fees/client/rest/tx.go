package rest

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"net/http"
	"strings"

	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/fees/createFee", createFeeHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/createFeeContract", createFeeContractHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/createSubscription", createSubscriptionHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/setFeeContractAuthorisation", setFeeContractAuthorisationHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/grantFeeDiscount", grantFeeDiscountHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/revokeFeeDiscount", revokeFeeDiscountHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/fees/chargeFee", chargeFeeHandler(cliCtx)).Methods("POST")
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func parseBool(boolStr, boolName string) (bool, error) {
	boolStr = strings.ToLower(strings.TrimSpace(boolStr))
	if boolStr == TRUE {
		return true, nil
	} else if boolStr == FALSE {
		return false, nil
	} else {
		return false, types.ErrInvalidArgument(types.DefaultCodespace, ""+
			fmt.Sprintf("%s is not a valid bool (true/false)", boolName))
	}
}

func IxoSignAndBroadcast(ctx context.CLIContext, msg sdk.Msg, sovrinDid sovrin.SovrinDid) ([]byte, error) {
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
	tx := ixo.NewIxoTxSingleMsg(msg, signature)

	bz, err := ctx.Codec.MarshalJSON(tx)
	if err != nil {
		return nil, fmt.Errorf("Could not marshall tx to binary. Error: %s", err.Error())
	}

	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return nil, fmt.Errorf("Could not broadcast tx. Error: %s", err.Error())
	}

	output, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, err
	}

	return output, nil
}

func createFeeHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeJsonParam := r.URL.Query().Get("feeJson")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		var fee types.Fee
		err := ctx.Codec.UnmarshalJSON([]byte(feeJsonParam), &fee)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgCreateFee(fee, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func createFeeContractHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeContractIdParam := r.URL.Query().Get("feeContractId")
		feeIdParam := r.URL.Query().Get("feeId")
		payerAddrParam := r.URL.Query().Get("payerAddr")
		canDeauthoriseParam := r.URL.Query().Get("canDeauthorise")
		discountIdParam := r.URL.Query().Get("discountId")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		payerAddr, err := sdk.AccAddressFromBech32(payerAddrParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		canDeauthorise, err := parseBool(canDeauthoriseParam, "canDeauthorise")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		discountId, err := sdk.ParseUint(discountIdParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgCreateFeeContract(feeIdParam, feeContractIdParam,
			payerAddr, canDeauthorise, discountId, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func createSubscriptionHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		subIdParam := r.URL.Query().Get("subId")
		feeContractIdParam := r.URL.Query().Get("feeContractId")
		maxPeriodsParam := r.URL.Query().Get("maxPeriods")
		periodParam := r.URL.Query().Get("period")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		maxPeriods, err := sdk.ParseUint(maxPeriodsParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var period types.Period
		err = ctx.Codec.UnmarshalJSON([]byte(periodParam), &period)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgCreateSubscription(subIdParam, feeContractIdParam,
			maxPeriods, period, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func setFeeContractAuthorisationHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeContractIdParam := r.URL.Query().Get("feeContractId")
		authorisedParam := r.URL.Query().Get("authorised")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		authorised, err := parseBool(authorisedParam, "authorised")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgSetFeeContractAuthorisation(feeContractIdParam,
			authorised, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func grantFeeDiscountHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeContractIdParam := r.URL.Query().Get("feeContractId")
		discountIdParam := r.URL.Query().Get("discountId")
		recipientAddrParam := r.URL.Query().Get("recipientAddr")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		discountId, err := sdk.ParseUint(discountIdParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		recipientAddr, err := sdk.AccAddressFromBech32(recipientAddrParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgGrantFeeDiscount(feeContractIdParam, discountId,
			recipientAddr, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func revokeFeeDiscountHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeContractIdParam := r.URL.Query().Get("feeContractId")
		holderAddrParam := r.URL.Query().Get("holderAddr")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		holderAddr, err := sdk.AccAddressFromBech32(holderAddrParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgRevokeFeeDiscount(feeContractIdParam, holderAddr,
			sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}

func chargeFeeHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		feeContractIdParam := r.URL.Query().Get("feeContractId")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		ctx = ctx.WithBroadcastMode(mode)

		sovrinDid := sovrin.SovrinDid{}
		err := json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		msg := types.NewMsgChargeFee(feeContractIdParam, sovrinDid)

		output, err := IxoSignAndBroadcast(ctx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(output)
	}
}
