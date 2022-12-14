// SiGG-GoLang-On-the-Fly //
package websockets

import "github.com/hyperledger/firefly-common/pkg/config"

const (
	bufferSizeDefault = "16Kb"
)

const (
	// ReadBufferSize is the read buffer size for the socket
	ReadBufferSize = "readBufferSize"
	// WriteBufferSize is the write buffer size for the socket
	WriteBufferSize = "writeBufferSize"
)

func (ws *WebSockets) InitConfig(config config.Section) {
	config.AddKnownKey(ReadBufferSize, bufferSizeDefault)
	config.AddKnownKey(WriteBufferSize, bufferSizeDefault)

}
