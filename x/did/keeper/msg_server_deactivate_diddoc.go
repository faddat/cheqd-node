package keeper

import (
	"context"

	"github.com/cheqd/cheqd-node/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k MsgServer) DeactivateDidDoc(goCtx context.Context, msg *types.MsgDeactivateDidDoc) (*types.MsgDeactivateDidDocResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get sign bytes before modifying payload
	signBytes := msg.Payload.GetSignBytes()

	// Normalize UUID identifiers
	msg.Normalize()

	// Validate DID does exist
	if !k.HasDidDoc(&ctx, msg.Payload.Id) {
		return nil, types.ErrDidDocNotFound.Wrap(msg.Payload.Id)
	}

	// Validate namespaces
	namespace := k.GetDidNamespace(&ctx)
	err := msg.Validate([]string{namespace})
	if err != nil {
		return nil, types.ErrNamespaceValidation.Wrap(err.Error())
	}

	// Retrieve didDoc state value and did
	didDoc, err := k.GetLatestDidDoc(&ctx, msg.Payload.Id)
	if err != nil {
		return nil, err
	}

	// Validate DID is not deactivated
	if didDoc.Metadata.Deactivated {
		return nil, types.ErrDIDDocDeactivated.Wrap(msg.Payload.Id)
	}

	// Check version id
	if msg.Payload.PreviousVersionId != didDoc.Metadata.VersionId {
		return nil, types.ErrUnexpectedDidVersion.Wrapf("got: %s, must be: %s", msg.Payload.VersionId, didDoc.Metadata.VersionId)
	}

	// We neither create dids nor update
	inMemoryDids := map[string]types.DidDocWithMetadata{}

	// Verify signatures
	signers := GetSignerDIDsForDIDCreation(*didDoc.DidDoc)
	err = VerifyAllSignersHaveAllValidSignatures(&k.Keeper, &ctx, inMemoryDids, signBytes, signers, msg.Signatures)
	if err != nil {
		return nil, err
	}

	// Update metadata
	didDoc.Metadata.Deactivated = true
	didDoc.Metadata.Update(ctx, msg.Payload.VersionId)

	// Apply changes. We don't create new version on deactivation, we just update metadata
	err = k.AddNewDidDocVersion(&ctx, &didDoc)
	if err != nil {
		return nil, types.ErrInternal.Wrapf(err.Error())
	}

	// Build and return response
	return &types.MsgDeactivateDidDocResponse{
		Value: &didDoc,
	}, nil
}
