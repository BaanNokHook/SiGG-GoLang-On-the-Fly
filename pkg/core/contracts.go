// SiGG-GoLang-On-the-Fly //
package core

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
)

type ContractCallType = fftypes.FFEnum

var (
	// CallTypeInvoke is an invocation that submits a transaction for inclusion in the chain
	CallTypeInvoke = fftypes.FFEnumValue("contractcalltype", "invoke")
	// CallTypeQuery is a query that returns data from the chain
	CallTypeQuery = fftypes.FFEnumValue("contractcalltype", "query")
)

type ContractCallRequest struct {
	Type           ContractCallType       `ffstruct:"ContractCallRequest" json:"type,omitempty" ffenum:"contractcalltype" ffexcludeinput:"true"`
	Interface      *fftypes.UUID          `ffstruct:"ContractCallRequest" json:"interface,omitempty" ffexcludeinput:"postContractAPIInvoke,postContractAPIQuery"`
	Location       *fftypes.JSONAny       `ffstruct:"ContractCallRequest" json:"location,omitempty"`
	Key            string                 `ffstruct:"ContractCallRequest" json:"key,omitempty"`
	Method         *fftypes.FFIMethod     `ffstruct:"ContractCallRequest" json:"method,omitempty" ffexcludeinput:"postContractAPIInvoke,postContractAPIQuery"`
	MethodPath     string                 `ffstruct:"ContractCallRequest" json:"methodPath,omitempty" ffexcludeinput:"postContractAPIInvoke,postContractAPIQuery"`
	Input          map[string]interface{} `ffstruct:"ContractCallRequest" json:"input"`
	Options        map[string]interface{} `ffstruct:"ContractCallRequest" json:"options"`
	IdempotencyKey IdempotencyKey         `ffstruct:"ContractCallRequest" json:"idempotencyKey,omitempty" ffexcludeoutput:"true"`
}

type ContractDeployRequest struct {
	Key            string                 `ffstruct:"ContractDeployRequest" json:"key,omitempty"`
	Input          []interface{}          `ffstruct:"ContractDeployRequest" json:"input"`
	Definition     *fftypes.JSONAny       `ffstruct:"ContractDeployRequest" json:"definition"`
	Contract       *fftypes.JSONAny       `ffstruct:"ContractDeployRequest" json:"contract"`
	Options        map[string]interface{} `ffstruct:"ContractDeployRequest" json:"options"`
	IdempotencyKey IdempotencyKey         `ffstruct:"ContractDeployRequest" json:"idempotencyKey,omitempty" ffexcludeoutput:"true"`
}

type ContractURLs struct {
	OpenAPI string `ffstruct:"ContractURLs" json:"openapi"`
	UI      string `ffstruct:"ContractURLs" json:"ui"`
}

type ContractAPI struct {
	ID        *fftypes.UUID         `ffstruct:"ContractAPI" json:"id,omitempty" ffexcludeinput:"true"`
	Namespace string                `ffstruct:"ContractAPI" json:"namespace,omitempty" ffexcludeinput:"true"`
	Interface *fftypes.FFIReference `ffstruct:"ContractAPI" json:"interface"`
	Location  *fftypes.JSONAny      `ffstruct:"ContractAPI" json:"location,omitempty"`
	Name      string                `ffstruct:"ContractAPI" json:"name"`
	Message   *fftypes.UUID         `ffstruct:"ContractAPI" json:"message,omitempty" ffexcludeinput:"true"`
	URLs      ContractURLs          `ffstruct:"ContractAPI" json:"urls" ffexcludeinput:"true"`
}

func (c *ContractAPI) Validate(ctx context.Context, existing bool) (err error) {
	if err = fftypes.ValidateFFNameField(ctx, c.Namespace, "namespace"); err != nil {
		return err
	}
	if err = fftypes.ValidateFFNameField(ctx, c.Name, "name"); err != nil {
		return err
	}
	return nil
}

func (c *ContractAPI) Topic() string {
	return fftypes.TypeNamespaceNameTopicHash("contractapi", c.Namespace, c.Name)
}

func (c *ContractAPI) SetBroadcastMessage(msgID *fftypes.UUID) {
	c.Message = msgID
}

func (c *ContractAPI) LocationAndLedgerEquals(a *ContractAPI) bool {
	if c == nil || a == nil {
		return false
	}
	return c.Location.Hash().Equals(a.Location.Hash())
}
