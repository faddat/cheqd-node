package setup

import (
	"crypto/ed25519"
	"encoding/base64"

	"github.com/cheqd/cheqd-node/x/cheqd/types"
	"github.com/cheqd/cheqd-node/x/cheqd/utils"
)

func (s *TestSetup) CreateDid(payload *types.MsgCreateDidPayload, signInputs []SignInput) (*types.MsgCreateDidResponse, error) {
	signBytes := payload.GetSignBytes()
	var signatures []*types.SignInfo

	for _, input := range signInputs {
		signature := ed25519.Sign(input.Key, signBytes)

		signatures = append(signatures, &types.SignInfo{
			VerificationMethodId: input.VerificationMethodId,
			Signature:            base64.StdEncoding.EncodeToString(signature),
		})
	}

	msg := &types.MsgCreateDid{
		Payload:    payload,
		Signatures: signatures,
	}

	return s.MsgServer.CreateDid(s.StdCtx, msg)
}

func (s *TestSetup) BuildSimpleDid() DidInfo {
	did := GenerateDID(Base58_16chars)
	_, _, collectionId := utils.MustSplitDID(did)

	keyPair := GenerateKeyPair()
	keyId := did + "#key-1"

	msg := &types.MsgCreateDidPayload{
		Id: did,
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                 keyId,
				Type:               types.Ed25519VerificationKey2020,
				Controller:         did,
				PublicKeyMultibase: MustEncodeBase58(keyPair.Public),
			},
		},
		Authentication: []string{keyId},
	}

	signInput := SignInput{
		VerificationMethodId: keyId,
		Key:                  keyPair.Private,
	}

	return DidInfo{
		Did:          did,
		CollectionId: collectionId,
		KeyPair:      keyPair,
		KeyId:        keyId,
		Msg:          msg,
		SignInput:    signInput,
	}
}

func (s *TestSetup) BuildUUIDDid(uuid string) DidInfo {
	did := "did:cheqd:" + DID_NAMESPACE + ":" + uuid
	_, _, collectionId := utils.MustSplitDID(did)

	keyPair := GenerateKeyPair()
	keyId := did + "#key-1"

	msg := &types.MsgCreateDidPayload{
		Id: did,
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                 keyId,
				Type:               types.Ed25519VerificationKey2020,
				Controller:         did,
				PublicKeyMultibase: MustEncodeBase58(keyPair.Public),
			},
		},
		Authentication: []string{keyId},
	}

	signInput := SignInput{
		VerificationMethodId: keyId,
		Key:                  keyPair.Private,
	}

	return DidInfo{
		Did:          did,
		CollectionId: collectionId,
		KeyPair:      keyPair,
		KeyId:        keyId,
		Msg:          msg,
		SignInput:    signInput,
	}
}

func (s *TestSetup) CreateSimpleDid() CreatedDidInfo {
	did := s.BuildSimpleDid()

	_, err := s.CreateDid(did.Msg, []SignInput{did.SignInput})
	if err != nil {
		panic(err)
	}

	created, err := s.QueryDid(did.Did)
	if err != nil {
		panic(err)
	}

	return CreatedDidInfo{
		DidInfo:   did,
		VersionId: created.Metadata.VersionId,
	}
}

func (s *TestSetup) CreateUUIDDid(uuid string) CreatedDidInfo {
	did := s.BuildUUIDDid(uuid)

	_, err := s.CreateDid(did.Msg, []SignInput{did.SignInput})
	if err != nil {
		panic(err)
	}

	created, err := s.QueryDid(did.Did)
	if err != nil {
		panic(err)
	}

	return CreatedDidInfo{
		DidInfo:   did,
		VersionId: created.Metadata.VersionId,
	}
}

func (s *TestSetup) CreateDidWithExternalConterllers(controllers []string, signInputs []SignInput) CreatedDidInfo {
	did := s.BuildSimpleDid()
	did.Msg.Controller = append(did.Msg.Controller, controllers...)

	_, err := s.CreateDid(did.Msg, append(signInputs, did.SignInput))
	if err != nil {
		panic(err)
	}

	created, err := s.QueryDid(did.Did)
	if err != nil {
		panic(err)
	}

	return CreatedDidInfo{
		DidInfo:   did,
		VersionId: created.Metadata.VersionId,
	}
}
