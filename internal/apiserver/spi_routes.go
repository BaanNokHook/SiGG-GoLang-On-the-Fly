// SiGG-GoLang-On-the-Fly //
package apiserver

import "github.com/hyperledger/firefly-common/pkg/ffapi"

// The Service Provider Interface (SPI) allows external microservices (such as the FireFly Transaction Manager)
// to act as augmented components to the core.
var spiRoutes = []*ffapi.Route{
	spiGetNamespaceByName,
	spiGetNamespaces,
	spiGetOpByID,
	spiGetOps,
	spiPatchOpByID,
	spiPostReset,
}
