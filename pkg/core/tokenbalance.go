// SiGG-GoLang-On-the-Fly //
package core

import "github.com/hyperledger/firefly-common/pkg/fftypes"

type TokenBalance struct {
	Pool       *fftypes.UUID    `ffstruct:"TokenBalance" json:"pool,omitempty"`
	TokenIndex string           `ffstruct:"TokenBalance" json:"tokenIndex,omitempty"`
	URI        string           `ffstruct:"TokenBalance" json:"uri,omitempty"`
	Connector  string           `ffstruct:"TokenBalance" json:"connector,omitempty"`
	Namespace  string           `ffstruct:"TokenBalance" json:"namespace,omitempty"`
	Key        string           `ffstruct:"TokenBalance" json:"key,omitempty"`
	Balance    fftypes.FFBigInt `ffstruct:"TokenBalance" json:"balance"`
	Updated    *fftypes.FFTime  `ffstruct:"TokenBalance" json:"updated,omitempty"`
}

func TokenBalanceIdentifier(pool *fftypes.UUID, tokenIndex, identity string) string {
	return pool.String() + ":" + tokenIndex + ":" + identity
}

func (t *TokenBalance) Identifier() string {
	return TokenBalanceIdentifier(t.Pool, t.TokenIndex, t.Key)
}

// Currently these types are just filtered views of TokenBalance.
// If more fields/aggregation become needed, they might merit a new table in the database.
type TokenAccount struct {
	Key string `ffstruct:"TokenBalance" json:"key,omitempty"`
}
type TokenAccountPool struct {
	Pool *fftypes.UUID `ffstruct:"TokenBalance" json:"pool,omitempty"`
}
