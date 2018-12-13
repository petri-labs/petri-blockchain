package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"
)

type AccDetails struct {
	Did     string  `json:"did"`
	Account string  `json:"account"`
	Balance sdk.Int `json:"balance"`
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/project/{did}",
		queryProjectDocRequestHandler(cdc, project.GetProjectDocDecoder(cdc)),
	).Methods("GET")
	r.HandleFunc(
		"/projectAccounts/{projectDid}",
		queryProjectAccountsRequestHandler(cdc, project.GetProjectDocDecoder(cdc)),
	).Methods("GET")
	r.HandleFunc(
		"/projectTxs/{projectDid}",
		queryProjectTxsRequestHandler(cdc),
	).Methods("GET")
}

func queryProjectDocRequestHandler(cdc *wire.Codec, decoder project.ProjectDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := ixo.Did(didAddr)

		res, err := ctx.QueryStore([]byte(key), storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this did
		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// decode the value
		projectDoc, err := decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole projectDoc
		output, err := json.MarshalIndent(projectDoc, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}

func queryProjectAccountsRequestHandler(cdc *wire.Codec, decoder project.ProjectDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout).
			WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		var buffer bytes.Buffer
		buffer.WriteString("ACC-")
		buffer.WriteString(projectDid)
		key := buffer.Bytes()

		res, err := ctx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this did
		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// decode the value
		var f interface{}
		err = json.Unmarshal(res, &f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}
		accMap := f.(map[string]interface{})

		accDetails := make([]AccDetails, len(accMap))
		i := 0
		for k, v := range accMap {
			addr := v.(string)
			fmt.Println("Address:", addr)
			accAddr := sdk.AccAddress([]byte(addr))

			account, err := ctx.GetAccount(accAddr)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Could't find account. Error: %s", err.Error())))
				return
			}
			baseAcc := account.(*types.AppAccount)
			balance := baseAcc.Coins.AmountOf(ixo.IxoNativeToken)
			accDetails[i] = AccDetails{Did: k, Account: addr, Balance: balance}
			i = i + 1
		}

		// print out whole didDoc
		output, err := json.MarshalIndent(accDetails, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}

}

func queryProjectTxsRequestHandler(cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		var buffer bytes.Buffer
		buffer.WriteString("TX-")
		buffer.WriteString(projectDid)
		key := buffer.Bytes()

		fmt.Println("Lookup key:", string(key))

		res, err := ctx.QueryStore(key, storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		fmt.Println("Len(res):", len(res))
		txs := []project.WithdrawalInfo{}
		// the query will return empty if there is no data for this did
		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			err = json.Unmarshal(res, &txs)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
				return
			}
		}

		// decode the value
		fmt.Println("Txs:", txs)
		// print out whole didDoc
		output, err := json.MarshalIndent(txs, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshal query result. Error: %s", err.Error())))
			return
		}

		fmt.Println("Output:", output)

		w.Write(output)
	}

}