package tests

import (
	"crypto/sha256"

	. "github.com/cheqd/cheqd-node/x/resource/tests/setup"
	"github.com/google/uuid"

	cheqdsetup "github.com/cheqd/cheqd-node/x/cheqd/tests/setup"
	cheqdtypes "github.com/cheqd/cheqd-node/x/cheqd/types"
	resourcetypes "github.com/cheqd/cheqd-node/x/resource/types"

<<<<<<< HEAD
	"github.com/cheqd/cheqd-node/x/resource/types"

	cheqdutils "github.com/cheqd/cheqd-node/x/cheqd/utils"
	"github.com/stretchr/testify/require"
=======
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
>>>>>>> origin/develop
)

func ExpectPayloadToMatchResource(payload *resourcetypes.MsgCreateResourcePayload, resource *resourcetypes.Resource) {
	// Provided header
	Expect(payload.Id).To(Equal(resource.Header.Id))
	Expect(payload.CollectionId).To(Equal(resource.Header.CollectionId))
	Expect(payload.Name).To(Equal(resource.Header.Name))
	Expect(payload.ResourceType).To(Equal(resource.Header.ResourceType))

	// Generated header
	hash := sha256.Sum256(payload.Data)
	Expect(resource.Header.Checksum).To(Equal(hash[:]))

	// Provided data
	Expect(payload.Data).To(Equal(resource.Data))
}

var _ = Describe("Create Resource Tests", func() {
	var setup TestSetup
	var alice cheqdsetup.CreatedDidInfo

	BeforeEach(func() {
		setup = Setup()
		alice = setup.CreateSimpleDid()
	})

	Describe("Simple resource", func() {
		var msg *resourcetypes.MsgCreateResourcePayload

		BeforeEach(func() {
			msg = &resourcetypes.MsgCreateResourcePayload{
				CollectionId: alice.CollectionId,
				Id:           uuid.NewString(),
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				Data:         []byte(SchemaData),
<<<<<<< HEAD
			},
			mediaType: JsonResourceType,
			errMsg:    fmt.Sprintf("signer: %s: signature is required but not found", cheqdutils.NormalizeIdentifier(ExistingDID)),
		},
		{
			valid:      false,
			name:       "Not Valid: Invalid Resource Id",
			signerKeys: map[string]ed25519.PrivateKey{},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: ExistingDIDIdentifier,
				Id:           IncorrectResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				Data:         []byte(SchemaData),
			},
			mediaType: JsonResourceType,
			errMsg:    fmt.Sprintf("signer: %s: signature is required but not found", cheqdutils.NormalizeIdentifier(ExistingDID)),
		},
		{
			valid:      false,
			name:       "Not Valid: DID Doc is not found",
			signerKeys: map[string]ed25519.PrivateKey{},
			msg: &types.MsgCreateResourcePayload{
				CollectionId: NotFoundDIDIdentifier,
				Id:           IncorrectResourceId,
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				Data:         []byte(SchemaData),
			},
			mediaType: JsonResourceType,
			errMsg:    fmt.Sprintf("did:cheqd:test:%s: not found", NotFoundDIDIdentifier),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			msg := tc.msg
			resourceSetup := InitEnv(t, keys[ExistingDIDKey].PublicKey, keys[ExistingDIDKey].PrivateKey)

			resource, err := resourceSetup.SendCreateResource(msg, tc.signerKeys)
			if tc.valid {
				require.Nil(t, err)

				did := utils.JoinDID("cheqd", "test", resource.Header.CollectionId)
				didStateValue, err := resourceSetup.Keeper.GetDid(&resourceSetup.Ctx, did)
				require.Nil(t, err)
				require.Contains(t, didStateValue.Metadata.Resources, resource.Header.Id)

				require.Equal(t, cheqdutils.NormalizeIdentifier(tc.msg.CollectionId), resource.Header.CollectionId)
				require.Equal(t, cheqdutils.NormalizeIdentifier(tc.msg.Id), resource.Header.Id)
				require.Equal(t, tc.mediaType, resource.Header.MediaType)
				require.Equal(t, tc.msg.ResourceType, resource.Header.ResourceType)
				require.Equal(t, tc.msg.Data, resource.Data)
				require.Equal(t, tc.msg.Name, resource.Header.Name)
				require.Equal(t, sha256.New().Sum(resource.Data), resource.Header.Checksum)
				require.Equal(t, tc.previousVersionId, resource.Header.PreviousVersionId)
			} else {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
=======
			}
		})

		It("Can be created with DIDDoc controller signature", func() {
			_, err := setup.CreateResource(msg, []cheqdsetup.SignInput{alice.SignInput})
			Expect(err).To(BeNil())

			// check
			created, err := setup.QueryResource(alice.CollectionId, msg.Id)
			Expect(err).To(BeNil())

			ExpectPayloadToMatchResource(msg, created.Resource)
		})

		It("Can't be created without DIDDoc controller signatures", func() {
			_, err := setup.CreateResource(msg, []cheqdsetup.SignInput{})
			Expect(err.Error()).To(ContainSubstring("signature is required but not found"))
		})

		It("Can't be created with invalid collection id", func() {
			msg.CollectionId = cheqdsetup.GenerateDID(cheqdsetup.Base58_16chars)

			_, err := setup.CreateResource(msg, []cheqdsetup.SignInput{alice.SignInput})
			Expect(err.Error()).To(ContainSubstring("not found"))
		})
	})

	Describe("New version", func() {
		var existingResource *resourcetypes.MsgCreateResourceResponse

		BeforeEach(func() {
			existingResource = setup.CreateSimpleResource(alice.CollectionId, SchemaData, "Test Resource Name", CLSchemaType, []cheqdsetup.SignInput{alice.SignInput})
		})

		It("Is linked to the previous one when name matches", func() {
			msg := resourcetypes.MsgCreateResourcePayload{
				CollectionId: alice.CollectionId,
				Id:           uuid.NewString(),
				Name:         existingResource.Resource.Header.Name,
				ResourceType: CLSchemaType,
				Data:         []byte(SchemaData),
			}

			_, err := setup.CreateResource(&msg, []cheqdsetup.SignInput{alice.SignInput})
			Expect(err).To(BeNil())

			// check
			created, err := setup.QueryResource(alice.CollectionId, msg.Id)
			Expect(err).To(BeNil())

			ExpectPayloadToMatchResource(&msg, created.Resource)
			Expect(created.Resource.Header.PreviousVersionId).To(Equal(existingResource.Resource.Header.Id))
		})
	})

	Describe("Resource for deactivated DID", func() {
		var msg *resourcetypes.MsgCreateResourcePayload

		BeforeEach(func() {
			msg = &resourcetypes.MsgCreateResourcePayload{
				CollectionId: alice.CollectionId,
				Id:           uuid.NewString(),
				Name:         "Test Resource Name",
				ResourceType: CLSchemaType,
				Data:         []byte(SchemaData),
>>>>>>> origin/develop
			}
		})

		When("DIDDoc is deactivated", func() {
			It("Should fail with error", func() {
				// Deactivate DID
				DeactivateMsg := &cheqdtypes.MsgDeactivateDidPayload{
					Id: alice.Did,
				}

				signatures := []cheqdsetup.SignInput{alice.DidInfo.SignInput}

				res, err := setup.DeactivateDid(DeactivateMsg, signatures)
				Expect(err).To(BeNil())
				Expect(res.Metadata.Deactivated).To(BeTrue())

				// Create resource
				_, err = setup.CreateResource(msg, []cheqdsetup.SignInput{alice.SignInput})
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring(alice.Did + ": DID Doc already deactivated"))
			})
		})
	})
})
