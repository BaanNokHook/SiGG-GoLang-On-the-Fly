// SiGG-GoLang-On-the-Fly //
package core

type PublishInput struct {
	IdempotencyKey IdempotencyKey `ffstruct:"PublishInput" json:"idempotencyKey,omitempty" ffexcludeoutput:"true"`
}
