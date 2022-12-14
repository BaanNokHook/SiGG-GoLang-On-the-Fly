// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/namespace"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

type coreRequest struct {
	mgr        namespace.Manager
	or         orchestrator.Orchestrator
	ctx        context.Context
	apiBaseURL string
}

type coreExtensions struct {
	EnabledIf             func(or orchestrator.Orchestrator) bool
	CoreJSONHandler       func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error)
	CoreFormUploadHandler func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error)
}

const (
	routeTagGlobal              = "Global"
	routeTagDefaultNamespace    = "Default Namespace"
	routeTagNonDefaultNamespace = "Non-Default Namespace"
)

var routes = append(
	globalRoutes([]*ffapi.Route{
		getNamespace,
		getNamespaces,
		getWebSockets,
	}),
	namespacedRoutes([]*ffapi.Route{
		deleteContractListener,
		deleteSubscription,
		getBatchByID,
		getBatches,
		getBlockchainEventByID,
		getBlockchainEvents,
		getChartHistogram,
		getContractAPIByName,
		getContractAPIInterface,
		getContractAPIs,
		getContractAPIListeners,
		getContractInterface,
		getContractInterfaceNameVersion,
		getContractInterfaces,
		getContractListenerByNameOrID,
		getContractListeners,
		getData,
		getDataBlob,
		getDataValue,
		getDataByID,
		getDataMsgs,
		getDatatypeByName,
		getDatatypes,
		getEventByID,
		getEvents,
		getGroupByHash,
		getGroups,
		getIdentities,
		getIdentityByDID,
		getIdentityByID,
		getIdentityDID,
		getIdentityVerifiers,
		getMsgByID,
		getMsgData,
		getMsgEvents,
		getMsgs,
		getMsgTxn,
		getNetworkDIDDocByDID,
		getNetworkIdentities,
		getNetworkIdentityByDID,
		getNetworkNode,
		getNetworkNodes,
		getNetworkOrg,
		getNetworkOrgs,
		getOpByID,
		getOps,
		getPins,
		getStatus,
		getStatusBatchManager,
		getSubscriptionByID,
		getSubscriptions,
		getTokenAccountPools,
		getTokenAccounts,
		getTokenApprovals,
		getTokenBalances,
		getTokenConnectors,
		getTokenPoolByNameOrID,
		getTokenPools,
		getTokenTransferByID,
		getTokenTransfers,
		getTxnBlockchainEvents,
		getTxnByID,
		getTxnOps,
		getTxns,
		getTxnStatus,
		getVerifierByID,
		getVerifiers,
		patchUpdateIdentity,
		postContractAPIInvoke,
		postContractAPIQuery,
		postContractAPIListeners,
		postContractInterfaceGenerate,
		postContractDeploy,
		postContractInvoke,
		postContractQuery,
		postData,
		postDataBlobPublish,
		postDataValuePublish,
		postNetworkAction,
		postNewContractAPI,
		postNewContractInterface,
		postNewContractListener,
		postNewDatatype,
		postNewIdentity,
		postNewMessageBroadcast,
		postNewMessagePrivate,
		postNewMessageRequestReply,
		postNewSubscription,
		postNewOrganization,
		postNewOrganizationSelf,
		postNodesSelf,
		postOpRetry,
		postPinsRewind,
		postTokenApproval,
		postTokenBurn,
		postTokenMint,
		postTokenPool,
		postTokenTransfer,
		putContractAPI,
		putSubscription,
		postVerifiersResolve,
	})...,
)

func globalRoutes(routes []*ffapi.Route) []*ffapi.Route {
	for _, route := range routes {
		route.Tag = routeTagGlobal
	}
	return routes
}

func namespacedRoutes(routes []*ffapi.Route) []*ffapi.Route {
	newRoutes := make([]*ffapi.Route, len(routes))
	for i, route := range routes {
		route.Tag = routeTagDefaultNamespace

		routeCopy := *route
		routeCopy.Name += "Namespace"
		routeCopy.Path = "namespaces/{ns}/" + route.Path
		routeCopy.PathParams = append(routeCopy.PathParams, &ffapi.PathParam{
			Name: "ns", ExampleFromConf: coreconfig.NamespacesDefault, Description: coremsgs.APIParamsNamespace,
		})
		routeCopy.Tag = routeTagNonDefaultNamespace
		newRoutes[i] = &routeCopy
	}
	return append(routes, newRoutes...)
}
