// SiGG-GoLang-On-the-Fly //
package core

import (
	"crypto/sha256"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
)

// VerifierType is the type of an identity verifier. Where possible we use established DID verifier type strings
type VerifierType = fftypes.FFEnum

var (
	// VerifierTypeEthAddress is an Ethereum (secp256k1) address string
	VerifierTypeEthAddress = fftypes.FFEnumValue("verifiertype", "ethereum_address")
	// VerifierTypeMSPIdentity is the MSP id (X509 distinguished name) of an issued signing certificate / keypair
	VerifierTypeMSPIdentity = fftypes.FFEnumValue("verifiertype", "fabric_msp_id")
	// VerifierTypeFFDXPeerID is the peer identifier that FireFly Data Exchange verifies (using plugin specific tech) when receiving data
	VerifierTypeFFDXPeerID = fftypes.FFEnumValue("verifiertype", "dx_peer_id")
)

// VerifierRef is just the type + value (public key identifier etc.) from the verifier
type VerifierRef struct {
	Type  VerifierType `ffstruct:"Verifier" json:"type" ffenum:"verifiertype"`
	Value string       `ffstruct:"Verifier" json:"value"`
}

// Verifier is an identity verification system that has been established for this identity, such as a blockchain signing key identifier
type Verifier struct {
	Hash      *fftypes.Bytes32 `ffstruct:"Verifier" json:"hash"` // Used to ensure the same ID is generated on each node, but not critical for verification. In v0.13 migration was set to the ID of the parent.
	Identity  *fftypes.UUID    `ffstruct:"Verifier" json:"identity,omitempty"`
	Namespace string           `ffstruct:"Verifier" json:"namespace,omitempty"`
	VerifierRef
	Created *fftypes.FFTime `ffstruct:"Verifier" json:"created,omitempty"`
}

// Seal updates the hash to be deterministically generated from the namespace+type+value, such that
// it will be the same on every node, and unique.
func (v *Verifier) Seal() *Verifier {
	h := sha256.New()
	h.Write([]byte(v.Namespace))
	h.Write([]byte(v.Type))
	h.Write([]byte(v.Value))
	v.Hash = fftypes.HashResult(h)
	return v
}
