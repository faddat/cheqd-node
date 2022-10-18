package types_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"

	. "github.com/cheqd/cheqd-node/x/cheqd/types"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/multiformats/go-multibase"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verification Method tests", func() {
	type TestCaseVerificationMethodStruct struct {
		vm VerificationMethod
		baseDid string
		allowedNamespaces []string
		isValid bool
		errorMsg string
	}

	DescribeTable("Verification Method Validation tests", func(testCase TestCaseVerificationMethodStruct) {
	
		err := testCase.vm.Validate(testCase.baseDid, testCase.allowedNamespaces)

		if testCase.isValid {
			Expect(err).To(BeNil())
		} else {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(testCase.errorMsg))
		}
	},

	Entry(
		"Verification method with expected multibase key", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#qwe",
				Type:               "Ed25519VerificationKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       nil,
				PublicKeyMultibase: ValidEd25519PubKey,
			},
			isValid: true,
			errorMsg: "",
		}),

	Entry(
		"Verification method with expected jwk key", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#rty",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			isValid: true,
			errorMsg: "",
		}),

	Entry(
		"Id has expected DID as a base", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#rty",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			baseDid: "did:cheqd:aaaaaaaaaaaaaaaa",
			isValid: true,
			errorMsg: "",
		}),

	Entry(
		"Id does not have expected DID as a base", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#rty",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			baseDid: "did:cheqd:bbbbbbbbbbbbbbbb",
			isValid: false,
			errorMsg: "id: must have prefix: did:cheqd:bbbbbbbbbbbbbbbb.",
		}),

	Entry(
		"Namespace is allowed", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:mainnet:aaaaaaaaaaaaaaaa#rty",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			allowedNamespaces: []string{"mainnet", ""},
			isValid: true,
		}),

	Entry(
		"Namespace is not allowed", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:mainnet:aaaaaaaaaaaaaaaa#rty",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			allowedNamespaces: []string{"testnet"},
			isValid: false,
			errorMsg: "controller: did namespace must be one of: testnet; id: did namespace must be one of: testnet.",
		}),
	Entry(
		"JWK key has expected format", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#qwe",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       ValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			isValid: true,
		}),
	Entry(
		"JWK key has unexpected format", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#qwe",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       NotValidPublicKeyJWK,
				PublicKeyMultibase: "",
			},
			isValid: false,
			errorMsg: "public_key_jwk: can't parse jwk: failed to parse key: invalid key type from JSON (SomeOtherKeyType).",
		}),
	Entry(
		"Not all keys and valuesin JWK have expected format", 
		TestCaseVerificationMethodStruct{
			vm: VerificationMethod{
				Id:                 "did:cheqd:aaaaaaaaaaaaaaaa#qwe",
				Type:               "JsonWebKey2020",
				Controller:         "did:cheqd:bbbbbbbbbbbbbbbb",
				PublicKeyJwk:       append(ValidPublicKeyJWK, &KeyValuePair{Key: "", Value: ""}),
				PublicKeyMultibase: "",
			},
			isValid: false,
			errorMsg: "public_key_jwk: (6: (key: cannot be blank; value: cannot be blank.).).",
		}),
	)
})

var _ = Describe("Validation ed25519 Signature in verification method", func() {
	var pubKey ed25519.PublicKey
	var privKey ed25519.PrivateKey
	var err error
	message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
		"tempor incididunt ut labore et dolore magna aliqua."
	msgBytes := []byte(message)
	var signature []byte

	pubKey, privKey, err = ed25519.GenerateKey(rand.Reader)
	Expect(err).To(BeNil())

	signature = ed25519.Sign(privKey, msgBytes)

	Context("when ed25519 key is placed", func() {
		It("is valid", func() {
			pubKeyStr, err := multibase.Encode(multibase.Base58BTC, pubKey)
			Expect(err).To(BeNil())

			vm := VerificationMethod{
				Id:                 "",
				Type:               "Ed25519VerificationKey2020",
				Controller:         "",
				PublicKeyJwk:       nil,
				PublicKeyMultibase: pubKeyStr,
			}

			err = VerifySignature(vm, msgBytes, signature)
			Expect(err).To(BeNil())
		})
	})

	Context("when with the same env but JWK is placed", func() {
		It("is valid", func() {
			jwk_, err := jwk.New(pubKey)
			Expect(err).To(BeNil())

			json_, err := json.MarshalIndent(jwk_, "", "  ")
			Expect(err).To(BeNil())

			pubKeyJwk := JSONToPubKeyJWK(string(json_))

			vm2 := VerificationMethod{
				Id:                 "",
				Type:               "JsonWebKey2020",
				Controller:         "",
				PublicKeyJwk:       pubKeyJwk,
				PublicKeyMultibase: "",
			}

			err = VerifySignature(vm2, msgBytes, signature)
			Expect(err).To(BeNil())
		})
	})
})

var _ = Describe("Validation ECDSA Signature in verification method", func() {
	Context("ECDSA signature preparations and verification", func() {
		It("is positive case", func() {
			message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
				"tempor incididunt ut labore et dolore magna aliqua."

			msgBytes := []byte(message)

			hasher := crypto.SHA256.New()
			hasher.Write(msgBytes)
			msgDigest := hasher.Sum(nil)

			privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			Expect(err).To(BeNil())

			pubKey := privKey.PublicKey

			signature, err := ecdsa.SignASN1(rand.Reader, privKey, msgDigest)
			Expect(err).To(BeNil())

			jwk_, err := jwk.New(pubKey)
			Expect(err).To(BeNil())

			json_, err := json.MarshalIndent(jwk_, "", "  ")
			Expect(err).To(BeNil())

			pubKeyJwk := JSONToPubKeyJWK(string(json_))

			vm := VerificationMethod{
				Id:                 "",
				Type:               "JsonWebKey2020",
				Controller:         "",
				PublicKeyJwk:       pubKeyJwk,
				PublicKeyMultibase: "",
			}

			err = VerifySignature(vm, msgBytes, signature)
			Expect(err).To(BeNil())
		})
	})
})

var _ = Describe("Validation RSA Signature in verification method", func() {
	Context("RSA signature preparations and verification", func() {
		It("is positive case", func() {
			message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
				"tempor incididunt ut labore et dolore magna aliqua."
			msgBytes := []byte(message)

			hasher := crypto.SHA256.New()
			hasher.Write(msgBytes)
			msgDigest := hasher.Sum(nil)

			privKey, err := rsa.GenerateKey(rand.Reader, 2048)
			Expect(err).To(BeNil())

			pubKey := privKey.PublicKey

			signature, err := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, msgDigest, nil)
			Expect(err).To(BeNil())

			jwk_, err := jwk.New(pubKey)
			Expect(err).To(BeNil())

			json_, err := json.MarshalIndent(jwk_, "", "  ")
			Expect(err).To(BeNil())

			pubKeyJwk := JSONToPubKeyJWK(string(json_))

			vm2 := VerificationMethod{
				Id:                 "",
				Type:               "JsonWebKey2020",
				Controller:         "",
				PublicKeyJwk:       pubKeyJwk,
				PublicKeyMultibase: "",
			}

			err = VerifySignature(vm2, msgBytes, signature)
			Expect(err).To(BeNil())
		})
	})
})
