package cli

import (
	"encoding/json"

	"github.com/cheqd/cheqd-node/tests/integration/helpers"
	"github.com/cheqd/cheqd-node/tests/integration/network"
	"github.com/cheqd/cheqd-node/x/did/client/cli"
	"github.com/cheqd/cheqd-node/x/did/types"
	resourcecli "github.com/cheqd/cheqd-node/x/resource/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var CLI_TX_PARAMS = []string{
	"--chain-id", network.CHAIN_ID,
	"--keyring-backend", KEYRING_BACKEND,
	"--output", OUTPUT_FORMAT,
	"--yes",
}

var CLI_GAS_PARAMS = []string{
	"--gas", GAS,
	"--gas-adjustment", GAS_ADJUSTMENT,
	"--gas-prices", GAS_PRICES,
}

func Tx(module, tx, from string, feeParams []string, txArgs ...string) (sdk.TxResponse, error) {
	args := []string{"tx", module, tx}

	// Common params
	args = append(args, CLI_TX_PARAMS...)

	// Fee params
	args = append(args, feeParams...)

	// Cosmos account
	args = append(args, "--from", from)

	// Other args
	args = append(args, txArgs...)

	output, err := Exec(args...)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// Skip 'gas estimate: xxx' string, trim 'Successfully migrated key' string
	output = helpers.TrimImportedStdout(output)

	var resp sdk.TxResponse

	err = helpers.Codec.UnmarshalJSON([]byte(output), &resp)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return resp, nil
}

func GrantFees(granter, grantee string, feeParams []string) (sdk.TxResponse, error) {
	return Tx("feegrant", "grant", granter, feeParams, granter, grantee)
}

func RevokeFeeGrant(granter, grantee string, feeParams []string) (sdk.TxResponse, error) {
	return Tx("feegrant", "revoke", granter, feeParams, granter, grantee)
}

func CreateDidDoc(tmpDir string, payload types.MsgCreateDidDocPayload, signInputs []cli.SignInput, from string, feeParams []string) (sdk.TxResponse, error) {
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDir, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "create-did", from, feeParams, payloadFile)
}

func UpdateDidDoc(tmpDir string, payload types.MsgUpdateDidDocPayload, signInputs []cli.SignInput, from string, feeParams []string) (sdk.TxResponse, error) {
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDir, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "update-did", from, feeParams, payloadFile)
}

func DeactivateDidDoc(tmpDir string, payload types.MsgDeactivateDidDocPayload, signInputs []cli.SignInput, from string, feeParams []string) (sdk.TxResponse, error) {
	payloadJson, err := helpers.Codec.MarshalJSON(&payload)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile(tmpDir, []byte(payloadWithSignInputsJson))

	return Tx("cheqd", "deactivate-did", from, feeParams, payloadFile)
}

func CreateResource(tmpDir string, options resourcecli.CreateResourceOptions, signInputs []cli.SignInput, from string, feeParams []string) (sdk.TxResponse, error) {
	payloadJson, err := json.Marshal(&options)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadWithSignInputs := cli.PayloadWithSignInputs{
		Payload:    payloadJson,
		SignInputs: signInputs,
	}

	payloadWithSignInputsJson, err := json.Marshal(&payloadWithSignInputs)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	payloadFile := helpers.MustWriteTmpFile("", []byte(payloadWithSignInputsJson))

	return Tx("resource", "create", from, feeParams, payloadFile)
}
