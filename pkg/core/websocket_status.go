// SiGG-GoLang-On-the-Fly //
package core

type WSSubscriptionStatus struct {
	Ephemeral bool   `ffstruct:"WSSubscriptionStatus" json:"ephemeral"`
	Namespace string `ffstruct:"WSSubscriptionStatus" json:"namespace"`
	Name      string `ffstruct:"WSSubscriptionStatus" json:"name,omitempty"`
}

type WSConnectionStatus struct {
	ID            string                  `ffstruct:"WSConnectionStatus" json:"id"`
	RemoteAddress string                  `ffstruct:"WSConnectionStatus" json:"remoteAddress"`
	UserAgent     string                  `ffstruct:"WSConnectionStatus" json:"userAgent"`
	Subscriptions []*WSSubscriptionStatus `ffstruct:"WSConnectionStatus" json:"subscriptions"`
}

type WebSocketStatus struct {
	Enabled     bool                  `ffstruct:"WebSocketStatus" json:"enabled"`
	Connections []*WSConnectionStatus `ffstruct:"WebSocketStatus" json:"connections"`
}
