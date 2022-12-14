// SiGG-GoLang-On-the-Fly //
package es

import (
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"golang.org/x/text/language"
)

//revive:disable

/*
This file contains the field level descriptions that are used in
OpenAPI Spec generation. Each struct field that wants to use one of these
needs to have an ffstruct tag on it, indicating the name of the struct.
That will be combined with the JSON field name (note, it is not the GO
field name, but the JSON serialized name), separated by a "." This is the
key used to lookup the translation below. If it is not found, the description
is left blank in the OpenAPI spec

Example:
// message.go
type Message struct {
	Header    MessageHeader `ffstruct:"Message" json:"header"`

// en_translations_descriptions.go
MessageHeader    = ffm("Message.header", "The message header")

*/

var ffm = func(key, translation string) i18n.MessageKey {
	return i18n.FFM(language.Spanish, key, translation)
}

var (
	// MessageHeader field descriptions
	MessageHeaderID = ffm("MessageHeader.id", "El UUID del mensaje. Único para cada mensaje")
)
