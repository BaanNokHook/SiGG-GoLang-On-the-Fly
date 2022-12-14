// SiGG-GoLang-On-the-Fly //
package core

type WebhookSubOptions struct {
	Fastack  bool                `ffstruct:"WebhookSubOptions" json:"fastack,omitempty"`
	URL      string              `ffstruct:"WebhookSubOptions" json:"url,omitempty"`
	Method   string              `ffstruct:"WebhookSubOptions" json:"method,omitempty"`
	JSON     bool                `ffstruct:"WebhookSubOptions" json:"json,omitempty"`
	Reply    bool                `ffstruct:"WebhookSubOptions" json:"reply,omitempty"`
	ReplyTag string              `ffstruct:"WebhookSubOptions" json:"replytag,omitempty"`
	ReplyTX  string              `ffstruct:"WebhookSubOptions" json:"replytx,omitempty"`
	Headers  map[string]string   `ffstruct:"WebhookSubOptions" json:"headers,omitempty"`
	Query    map[string]string   `ffstruct:"WebhookSubOptions" json:"query,omitempty"`
	Input    WebhookInputOptions `ffstruct:"WebhookSubOptions" json:"input,omitempty"`
}

type WebhookInputOptions struct {
	Query   string `ffstruct:"WebhookInputOptions" json:"query,omitempty"`
	Headers string `ffstruct:"WebhookInputOptions" json:"headers,omitempty"`
	Body    string `ffstruct:"WebhookInputOptions" json:"body,omitempty"`
	Path    string `ffstruct:"WebhookInputOptions" json:"path,omitempty"`
	ReplyTX string `ffstruct:"WebhookInputOptions" json:"replytx,omitempty"`
}
