// SiGG-GoLang-On-the-Fly //
package coremsgs

import (
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"golang.org/x/text/language"
)

var ffe = func(key, translation string, statusHint ...int) i18n.ErrorMessageKey {
	return i18n.FFE(language.AmericanEnglish, key, translation, statusHint...)
}

//revive:disable
var (
	MsgConfigFailed                       = ffe("FF10101", "Failed to read config")
	MsgJSONDecodeFailed                   = ffe("FF10103", "Failed to decode input JSON")
	MsgTLSConfigFailed                    = ffe("FF10105", "Failed to initialize TLS configuration")
	MsgWebsocketClientError               = ffe("FF10108", "Error received from WebSocket client: %s")
	Msg404NotFound                        = ffe("FF10109", "Not found", 404)
	MsgUnknownBlockchainPlugin            = ffe("FF10110", "Unknown blockchain plugin: %s")
	MsgEthconnectRESTErr                  = ffe("FF10111", "Error from ethconnect: %s")
	MsgDBInitFailed                       = ffe("FF10112", "Database initialization failed")
	MsgDBQueryBuildFailed                 = ffe("FF10113", "Database query builder failed")
	MsgDBBeginFailed                      = ffe("FF10114", "Database begin transaction failed")
	MsgDBQueryFailed                      = ffe("FF10115", "Database query failed")
	MsgDBInsertFailed                     = ffe("FF10116", "Database insert failed")
	MsgDBUpdateFailed                     = ffe("FF10117", "Database update failed")
	MsgDBDeleteFailed                     = ffe("FF10118", "Database delete failed")
	MsgDBCommitFailed                     = ffe("FF10119", "Database commit failed")
	MsgDBMissingJoin                      = ffe("FF10120", "Database missing expected join entry in table '%s' for id '%s'")
	MsgDBReadErr                          = ffe("FF10121", "Database resultset read error from table '%s'")
	MsgUnknownDatabasePlugin              = ffe("FF10122", "Unknown database plugin '%s'")
	MsgNullDataReferenceID                = ffe("FF10123", "Data id is null in message data reference %d")
	MsgDupDataReferenceID                 = ffe("FF10124", "Duplicate data ID in message '%s'", 409)
	MsgScanFailed                         = ffe("FF10125", "Failed to restore type '%T' into '%T'")
	MsgUnregisteredBatchType              = ffe("FF10126", "Unregistered batch type '%s'")
	MsgBatchDispatchTimeout               = ffe("FF10127", "Timed out dispatching work to batch")
	MsgInitializationNilDepError          = ffe("FF10128", "Initialization failed in %s due to unmet dependency")
	MsgNilResponseNon204                  = ffe("FF10129", "No output from API call")
	MsgDataNotFound                       = ffe("FF10133", "Data not found for message %s", 400)
	MsgUnknownSharedStoragePlugin         = ffe("FF10134", "Unknown Shared Storage plugin '%s'")
	MsgIPFSHashDecodeFailed               = ffe("FF10135", "Failed to decode IPFS hash into 32byte value '%s'")
	MsgIPFSRESTErr                        = ffe("FF10136", "Error from IPFS: %s")
	MsgSerializationFailed                = ffe("FF10137", "Serialization failed")
	MsgMissingPluginConfig                = ffe("FF10138", "Missing configuration '%s' for %s")
	MsgMissingDataHashIndex               = ffe("FF10139", "Missing data hash for index '%d' in message", 400)
	MsgInvalidEthAddress                  = ffe("FF10141", "Supplied ethereum address is invalid", 400)
	Msg404NoResult                        = ffe("FF10143", "No result found", 404)
	MsgUnsupportedSQLOpInFilter           = ffe("FF10150", "No SQL mapping implemented for filter operator '%s'", 400)
	MsgFilterSortDesc                     = ffe("FF10154", "Sort field. For multi-field sort use comma separated values (or multiple query values) with '-' prefix for descending")
	MsgContextCanceled                    = ffe("FF00154", "Context cancelled")
	MsgDBMigrationFailed                  = ffe("FF10163", "Database migration failed")
	MsgHashMismatch                       = ffe("FF10164", "Hash mismatch")
	MsgDefaultNamespaceNotFound           = ffe("FF10166", "namespaces.default '%s' must be included in the namespaces.predefined configuration")
	MsgEventTypesParseFail                = ffe("FF10168", "Unable to parse list of event types", 400)
	MsgUnknownEventType                   = ffe("FF10169", "Unknown event type '%s'", 400)
	MsgIDMismatch                         = ffe("FF10170", "ID mismatch")
	MsgRegexpCompileFailed                = ffe("FF10171", "Unable to compile '%s' regexp '%s'")
	MsgUnknownEventTransportPlugin        = ffe("FF10172", "Unknown event transport plugin: %s")
	MsgWSConnectionNotActive              = ffe("FF10173", "Websocket connection '%s' no longer active")
	MsgWSSubAlreadyInFlight               = ffe("FF10174", "Websocket subscription '%s' already has a message in flight")
	MsgWSMsgSubNotMatched                 = ffe("FF10175", "Acknowledgment does not match an inflight event + subscription")
	MsgWSClientSentInvalidData            = ffe("FF10176", "Invalid data")
	MsgWSClientUnknownAction              = ffe("FF10177", "Unknown action '%s'")
	MsgWSInvalidStartAction               = ffe("FF10178", "A start action must set namespace and either a name or ephemeral=true")
	MsgWSAutoAckChanged                   = ffe("FF10179", "The autoack option must be set consistently on all start requests")
	MsgWSAutoAckEnabled                   = ffe("FF10180", "The autoack option is enabled on this connection")
	MsgConnSubscriptionNotStarted         = ffe("FF10181", "Subscription %v is not started on connection")
	MsgDispatcherClosing                  = ffe("FF10182", "Event dispatcher closing")
	MsgMaxFilterSkip                      = ffe("FF10183", "You have reached the maximum pagination limit for this query (%d)", 400)
	MsgMaxFilterLimit                     = ffe("FF10184", "Your query exceeds the maximum filter limit (%d)", 400)
	MsgAPIServerStaticFail                = ffe("FF10185", "An error occurred loading static content", 500)
	MsgEventListenerClosing               = ffe("FF10186", "Event listener closing")
	MsgNamespaceDoesNotExist              = ffe("FF10187", "Namespace does not exist", 404)
	MsgInvalidSubscription                = ffe("FF10189", "Invalid subscription", 400)
	MsgMismatchedTransport                = ffe("FF10190", "Connection ID '%s' appears not to be unique between transport '%s' and '%s'", 400)
	MsgInvalidFirstEvent                  = ffe("FF10191", "Invalid firstEvent definition - must be 'newest','oldest' or a sequence number", 400)
	MsgNumberMustBeGreaterEqual           = ffe("FF10192", "Number must be greater than or equal to %d", 400)
	MsgAlreadyExists                      = ffe("FF10193", "A %s with name '%s:%s' already exists", 409)
	MsgJSONValidatorBadRef                = ffe("FF10194", "Cannot use JSON validator for data with type '%s' and validator reference '%v'", 400)
	MsgDatatypeNotFound                   = ffe("FF10195", "Datatype '%v' not found", 400)
	MsgSchemaLoadFailed                   = ffe("FF10196", "Datatype '%s' schema invalid", 400)
	MsgDataCannotBeValidated              = ffe("FF10197", "Data cannot be validated", 400)
	MsgJSONDataInvalidPerSchema           = ffe("FF10198", "Data does not conform to the JSON schema of datatype '%s': %s", 400)
	MsgDataValueIsNull                    = ffe("FF10199", "Data value is null", 400)
	MsgDataInvalidHash                    = ffe("FF10201", "Invalid data: hashes do not match Hash=%s Expected=%s", 400)
	MsgDataReferenceUnresolvable          = ffe("FF10204", "Data reference %d cannot be resolved", 400)
	MsgDataMissing                        = ffe("FF10205", "Data entry %d has neither 'id' to refer to existing data, or 'value' to include in-line JSON data", 400)
	MsgAuthorInvalid                      = ffe("FF10206", "Invalid author specified", 400)
	MsgNoTransaction                      = ffe("FF10207", "Message does not have a transaction", 404)
	MsgBatchNotSet                        = ffe("FF10208", "Message does not have an assigned batch", 404)
	MsgBatchNotFound                      = ffe("FF10209", "Batch '%s' not found for message", 500)
	MsgBatchTXNotSet                      = ffe("FF10210", "Batch '%s' does not have an assigned transaction", 404)
	MsgOwnerMissing                       = ffe("FF10211", "Owner missing", 400)
	MsgUnknownIdentityPlugin              = ffe("FF10212", "Unknown Identity plugin '%s'")
	MsgUnknownDataExchangePlugin          = ffe("FF10213", "Unknown Data Exchange plugin '%s'")
	MsgParentIdentityNotFound             = ffe("FF10214", "Identity '%s' not found in identity chain for %s '%s'")
	MsgInvalidSigningIdentity             = ffe("FF10215", "Invalid signing identity")
	MsgNodeAndOrgIDMustBeSet              = ffe("FF10216", "node.name, org.name and org.key must be configured first", 409)
	MsgBlobStreamingFailed                = ffe("FF10217", "Blob streaming terminated with error", 500)
	MsgNodeNotFound                       = ffe("FF10224", "Node with name or identity '%s' not found", 400)
	MsgLocalNodeResolveFailed             = ffe("FF10225", "Unable to find local node to add to group. Check the status API to confirm the node is registered", 500)
	MsgGroupNotFound                      = ffe("FF10226", "Group '%s' not found", 404)
	MsgDXRESTErr                          = ffe("FF10229", "Error from data exchange: %s")
	MsgInvalidHex                         = ffe("FF10231", "Invalid hex supplied", 400)
	MsgInvalidWrongLenB32                 = ffe("FF00107", "Byte length must be 32 (64 hex characters)", 400)
	MsgNodeNotFoundInOrg                  = ffe("FF10233", "Unable to find any nodes owned by org '%s', or parent orgs", 400)
	MsgDXBadResponse                      = ffe("FF10237", "Unexpected '%s' in data exchange response: %s")
	MsgDXBadHash                          = ffe("FF10238", "Unexpected hash returned from data exchange upload. Hash=%s Expected=%s")
	MsgBlobNotFound                       = ffe("FF10239", "No blob has been uploaded or confirmed received, with hash=%s", 404)
	MsgDownloadBlobFailed                 = ffe("FF10240", "Error download blob with reference '%s' from local data exchange")
	MsgDataDoesNotHaveBlob                = ffe("FF10241", "Data does not have a blob attachment", 404)
	MsgWebhookURLEmpty                    = ffe("FF10242", "Webhook subscription option 'url' cannot be empty", 400)
	MsgWebhookInvalidStringMap            = ffe("FF10243", "Webhook subscription option '%s' must be map of string values. %s=%T", 400)
	MsgWebsocketsNoData                   = ffe("FF10244", "Websockets subscriptions do not support streaming the full data payload, just the references (withData must be false)", 400)
	MsgWebhooksWithData                   = ffe("FF10245", "Webhook subscriptions require the full data payload (withData must be true)", 400)
	MsgWebhooksReplyBadJSON               = ffe("FF10257", "Failed to process reply from webhook as JSON")
	MsgRequestTimeout                     = ffe("FF10260", "The request with id '%s' timed out after %.2fms", 408)
	MsgRequestReplyTagRequired            = ffe("FF10261", "For request messages 'header.tag' must be set on the request message to route it to a suitable responder", 400)
	MsgRequestCannotHaveCID               = ffe("FF10262", "For request messages 'header.cid' must be unset", 400)
	MsgSystemTransportInternal            = ffe("FF10266", "You cannot create subscriptions on the system events transport")
	MsgFilterCountNotSupported            = ffe("FF10267", "This query does not support generating a count of all results")
	MsgRejected                           = ffe("FF10269", "Message with ID '%s' was rejected. Please check the FireFly logs for more information")
	MsgRequestMustBePrivate               = ffe("FF10271", "For request messages you must specify a group of private recipients", 400)
	MsgUnknownTokensPlugin                = ffe("FF10272", "Unknown tokens plugin '%s'", 400)
	MsgMissingTokensPluginConfig          = ffe("FF10273", "Invalid tokens configuration - name and plugin are required", 400)
	MsgTokensRESTErr                      = ffe("FF10274", "Error from tokens service: %s")
	MsgTokenPoolDuplicate                 = ffe("FF10275", "Duplicate token pool: %s", 409)
	MsgTokenPoolRejected                  = ffe("FF10276", "Token pool with ID '%s' was rejected. Please check the FireFly logs for more information")
	MsgIdentityNotFoundByString           = ffe("FF10277", "Identity could not be resolved via lookup string '%s'")
	MsgAuthorOrgSigningKeyMismatch        = ffe("FF10279", "Author organization '%s' is not associated with signing key '%s'")
	MsgCannotTransferToSelf               = ffe("FF10280", "From and to addresses must be different", 400)
	MsgLocalOrgLookupFailed               = ffe("FF10281", "Unable to resolve the local org '%s' by the configured signing key on the node. Please confirm the org is registered with key '%s'", 500)
	MsgFabconnectRESTErr                  = ffe("FF10284", "Error from fabconnect: %s")
	MsgInvalidIdentity                    = ffe("FF10285", "Supplied Fabric signer identity is invalid", 400)
	MsgFailedToDecodeCertificate          = ffe("FF10286", "Failed to decode certificate: %s", 500)
	MsgInvalidMessageType                 = ffe("FF10287", "Invalid message type - allowed types are %s", 400)
	MsgWSClosed                           = ffe("FF10290", "Websocket closed")
	MsgFieldNotSpecified                  = ffe("FF10292", "Field '%s' must be specified", 400)
	MsgTokenPoolNotConfirmed              = ffe("FF10293", "Token pool is not yet confirmed")
	MsgHistogramCollectionParam           = ffe("FF10297", "Collection to fetch")
	MsgInvalidNumberOfIntervals           = ffe("FF10298", "Number of time intervals must be between %d and %d", 400)
	MsgInvalidChartNumberParam            = ffe("FF10299", "Invalid %s. Must be a number.", 400)
	MsgHistogramInvalidTimes              = ffe("FF10300", "Start time must be before end time", 400)
	MsgUnsupportedCollection              = ffe("FF10301", "%s collection is not supported", 400)
	MsgContractInterfaceExists            = ffe("FF10302", "A contract interface already exists in the namespace: '%s' with name: '%s' and version: '%s'", 409)
	MsgContractInterfaceNotFound          = ffe("FF10303", "Contract interface %s not found", 404)
	MsgContractMissingInputArgument       = ffe("FF10304", "Missing required input argument '%s'", 400)
	MsgContractWrongInputType             = ffe("FF10305", "Input '%v' is of type '%v' not expected type of '%v'", 400)
	MsgContractMissingInputField          = ffe("FF10306", "Expected object of type '%v' to contain field named '%v' but it was missing", 400)
	MsgContractMapInputType               = ffe("FF10307", "Unable to map input type '%v' to known FireFly type - was expecting '%v'", 400)
	MsgContractByteDecode                 = ffe("FF10308", "Unable to decode field '%v' as bytes", 400)
	MsgContractInternalType               = ffe("FF10309", "Input '%v' of type '%v' is not compatible blockchain internalType of '%v'", 400)
	MsgContractLocationInvalid            = ffe("FF10310", "Failed to validate contract location: %v", 400)
	MsgContractParamInvalid               = ffe("FF10311", "Failed to validate contract param: %v", 400)
	MsgContractListenerNameExists         = ffe("FF10312", "A contract listener already exists in the namespace: '%s' with name: '%s'", 409)
	MsgContractMethodNotSet               = ffe("FF10313", "Either an interface reference and method path, or in-line method definition, must be supplied on invoke contract request", 400)
	MsgContractMethodResolveError         = ffe("FF10315", "Unable to resolve contract method: %s", 400)
	MsgContractLocationExists             = ffe("FF10316", "The contract location cannot be changed after it is created", 400)
	MsgListenerNoEvent                    = ffe("FF10317", "Either an interface reference and event path, or in-line event definition must be supplied when creating a contract listener", 400)
	MsgListenerEventNotFound              = ffe("FF10318", "No event was found in namespace '%s' with id '%s'", 400)
	MsgEventNameMustBeSet                 = ffe("FF10319", "Event name must be set", 400)
	MsgMethodNameMustBeSet                = ffe("FF10320", "Method name must be set", 400)
	MsgContractEventResolveError          = ffe("FF10321", "Unable to resolve contract event", 400)
	MsgQueryOpUnsupportedMod              = ffe("FF10322", "Operation '%s' on '%s' does not support modifiers", 400)
	MsgDXBadSize                          = ffe("FF10323", "Unexpected size returned from data exchange upload. Size=%d Expected=%d")
	MsgTooLargeBroadcast                  = ffe("FF10327", "Message size %.2fkb is too large for the max broadcast batch size of %.2fkb", 400)
	MsgTooLargePrivate                    = ffe("FF10328", "Message size %.2fkb is too large for the max private message size of %.2fkb", 400)
	MsgManifestMismatch                   = ffe("FF10329", "Manifest mismatch overriding '%s' status as failure: '%s'", 400)
	MsgFFIValidationFail                  = ffe("FF10331", "Field '%s' does not validate against the provided schema", 400)
	MsgFFISchemaParseFail                 = ffe("FF10332", "Failed to parse schema for param '%s'", 400)
	MsgFFISchemaCompileFail               = ffe("FF10333", "Failed compile schema for param '%s'", 400)
	MsgPluginInitializationFailed         = ffe("FF10334", "Plugin initialization error", 500)
	MsgUnknownTransactionType             = ffe("FF10336", "Unknown transaction type '%s'", 400)
	MsgGoTemplateCompileFailed            = ffe("FF10337", "Go template compilation for '%s' failed: %s", 500)
	MsgGoTemplateExecuteFailed            = ffe("FF10338", "Go template execution for '%s' failed: %s", 500)
	MsgAddressResolveFailed               = ffe("FF10339", "Failed to resolve signing key string '%s': %s", 500)
	MsgAddressResolveBadStatus            = ffe("FF10340", "Failed to resolve signing key string '%s' [%d]: %s", 500)
	MsgAddressResolveBadResData           = ffe("FF10341", "Failed to resolve signing key string '%s' - invalid address returned '%s': %s", 500)
	MsgDXNotInitialized                   = ffe("FF10342", "Data exchange is initializing")
	MsgGroupRequired                      = ffe("FF10344", "Group must be set", 400)
	MsgDBLockFailed                       = ffe("FF10345", "Database lock failed")
	MsgFFIGenerationFailed                = ffe("FF10346", "Error generating smart contract interface: %s", 400)
	MsgFFIGenerationUnsupported           = ffe("FF10347", "Smart contract interface generation is not supported by this blockchain plugin", 400)
	MsgBlobHashMismatch                   = ffe("FF10348", "Blob hash mismatch sent=%s received=%s", 400)
	MsgDIDResolverUnknown                 = ffe("FF10349", "DID resolver unknown for DID: %s", 400)
	MsgIdentityNotOrg                     = ffe("FF10350", "Identity '%s' with DID '%s' is not an organization", 400)
	MsgIdentityNotNode                    = ffe("FF10351", "Identity '%s' with DID '%s' is not a node", 400)
	MsgBlockchainKeyNotSet                = ffe("FF10352", "No blockchain key specified", 400)
	MsgNoVerifierForIdentity              = ffe("FF10353", "No %s verifier registered for identity %s", 400)
	MsgNodeMissingBlockchainKey           = ffe("FF10354", "No default signing key or organization signing key configured for this namespace", 400)
	MsgAuthorRegistrationMismatch         = ffe("FF10355", "Verifier '%s' cannot be used for signing with author '%s'. Verifier registered to '%s'", 400)
	MsgAuthorMissingForKey                = ffe("FF10356", "Key '%s' has not been registered by any identity, and a separate 'author' was not supplied", 404)
	MsgAuthorIncorrectForRootReg          = ffe("FF10357", "Author namespace '%s' and DID '%s' combination invalid for root organization registration", 400)
	MsgKeyIdentityMissing                 = ffe("FF10358", "Identity owner of key '%s' not found", 500)
	MsgIdentityChainLoop                  = ffe("FF10364", "Loop detected on identity %s in chain for %s (%s)", 400)
	MsgInvalidIdentityParentType          = ffe("FF10365", "Parent %s (%s) of type %s is invalid for child %s (%s) of type", 400)
	MsgParentIdentityMissingClaim         = ffe("FF10366", "Parent %s (%s) is invalid (missing claim)", 400)
	MsgDXInfoMissingID                    = ffe("FF10367", "Data exchange endpoint info missing 'id' field", 500)
	MsgEventNotFound                      = ffe("FF10370", "Event with name '%s' not found", 400)
	MsgOperationNotSupported              = ffe("FF10371", "Operation not supported: %s", 400)
	MsgFailedToRetrieve                   = ffe("FF10372", "Failed to retrieve %s %s", 500)
	MsgBlobMissingPublic                  = ffe("FF10373", "Blob for data %s missing public payload reference while flushing batch", 500)
	MsgDBMultiRowConfigError              = ffe("FF10374", "Database invalid configuration - using multi-row insert on DB plugin that does not support query syntax for input")
	MsgDBNoSequence                       = ffe("FF10375", "Failed to retrieve sequence for insert row %d (could mean duplicate insert)", 500)
	MsgDownloadSharedFailed               = ffe("FF10376", "Error downloading data with reference '%s' from shared storage")
	MsgDownloadBatchMaxBytes              = ffe("FF10377", "Error downloading batch with reference '%s' from shared storage - maximum size limit reached")
	MsgOperationDataIncorrect             = ffe("FF10378", "Operation data type incorrect: %T", 400)
	MsgDataMissingBlobHash                = ffe("FF10379", "Blob for data %s cannot be transferred as it is missing a hash", 500)
	MsgUnexpectedDXMessageType            = ffe("FF10380", "Unexpected websocket event type from DX plugin: %s", 500)
	MsgContractListenerExists             = ffe("FF10383", "A contract listener already exists for this combination of topic + location + event", 409)
	MsgInvalidOutputOption                = ffe("FF10385", "invalid output option '%s'")
	MsgInvalidPluginConfiguration         = ffe("FF10386", "Invalid %s plugin configuration - name and type are required")
	MsgReferenceMarkdownMissing           = ffe("FF10387", "Reference markdown file missing: '%s'")
	MsgFFSystemReservedName               = ffe("FF10388", "Invalid namespace configuration - %s is a reserved name")
	MsgInvalidNamespaceMode               = ffe("FF10389", "Invalid %s namespace configuration - unknown mode")
	MsgNamespaceUnknownPlugin             = ffe("FF10390", "Invalid %s namespace configuration - unknown plugin %s")
	MsgNamespaceWrongPluginsMultiparty    = ffe("FF10391", "Invalid %s namespace configuration - multiparty mode requires database, blockchain, shared storage, and data exchange plugins")
	MsgNamespaceNoDatabase                = ffe("FF10392", "Invalid %s namespace configuration - a database plugin is required")
	MsgNamespaceMultiplePluginType        = ffe("FF10394", "Invalid %s namespace configuration - multiple %s plugins provided")
	MsgDuplicatePluginName                = ffe("FF10395", "Invalid plugin configuration - plugin with name %s already exists", 409)
	MsgInvalidFireFlyContractIndex        = ffe("FF10396", "No configuration found for FireFly contract at %s")
	MsgUnrecognizedNetworkAction          = ffe("FF10397", "Unrecognized network action: %s", 400)
	MsgOverrideExistingFieldCustomOption  = ffe("FF10398", "Cannot override existing field with custom option named '%s'", 400)
	MsgTerminateNotSupported              = ffe("FF10399", "The 'terminate' operation to mark a switchover of smart contracts is not supported on namespace %s", 400)
	MsgDefRejectedBadPayload              = ffe("FF10400", "Rejected %s message '%s' - invalid payload")
	MsgDefRejectedAuthorBlank             = ffe("FF10401", "Rejected %s message '%s' - author is blank")
	MsgDefRejectedSignatureMismatch       = ffe("FF10402", "Rejected %s message '%s' - signature mismatch")
	MsgDefRejectedValidateFail            = ffe("FF10403", "Rejected %s '%s' - validate failed: %s")
	MsgDefRejectedIDMismatch              = ffe("FF10404", "Rejected %s '%s' - ID mismatch with existing record")
	MsgDefRejectedLocationMismatch        = ffe("FF10405", "Rejected %s '%s' - location mismatch with existing record")
	MsgDefRejectedSchemaFail              = ffe("FF10406", "Rejected %s '%s' - schema check: %s")
	MsgDefRejectedConflict                = ffe("FF10407", "Rejected %s '%s' - conflicts with existing: %s")
	MsgDefRejectedIdentityNotFound        = ffe("FF10408", "Rejected %s '%s' - identity not found: %s")
	MsgDefRejectedWrongAuthor             = ffe("FF10409", "Rejected %s '%s' - wrong author: %s")
	MsgDefRejectedHashMismatch            = ffe("FF10410", "Rejected %s '%s' - hash mismatch: %s != %s")
	MsgInvalidNamespaceUUID               = ffe("FF10411", "Expected 'namespace:' prefix on ID '%s'", 400)
	MsgBadNetworkVersion                  = ffe("FF10412", "Bad network version: %s")
	MsgDefinitionRejected                 = ffe("FF10413", "Definition rejected")
	MsgActionNotSupported                 = ffe("FF10414", "This action is not supported in this namespace", 400)
	MsgMessagesNotSupported               = ffe("FF10415", "Messages are not supported in this namespace", 400)
	MsgInvalidSubscriptionForNetwork      = ffe("FF10416", "Subscription name '%s' is invalid according to multiparty network rules in effect (network version=%d)")
	MsgBlockchainNotConfigured            = ffe("FF10417", "No blockchain plugin configured")
	MsgInvalidBatchPinEvent               = ffe("FF10418", "BatchPin event is not valid - %s (%s): %s")
	MsgDuplicatePluginBroadcastName       = ffe("FF10419", "Invalid %s plugin broadcast name: %s - broadcast names must be unique", 409)
	MsgInvalidConnectorName               = ffe("FF10420", "Could not find name %s for %s connector")
	MsgCannotInitLegacyNS                 = ffe("FF10421", "could not initialize legacy '%s' namespace - found conflicting V1 multi-party config in %s and %s")
	MsgInvalidGroupMember                 = ffe("FF10422", "invalid group member - node '%s' is not owned by '%s' or any of its ancestors")
	MsgContractListenerStatusInvalid      = ffe("FF10423", "Failed to validate contract listener status: %v", 400)
	MsgCacheMissSizeLimitKeyInternal      = ffe("FF10424", "could not initialize cache - size limit config key is not provided")
	MsgCacheMissTTLKeyInternal            = ffe("FF10425", "could not initialize cache - ttl config key is not provided")
	MsgCacheConfigKeyMismatchInternal     = ffe("FF10426", "could not initialize cache - '%s' and '%s' do not have identical prefix, mismatching prefixes are: '%s','%s'")
	MsgCacheUnexpectedSizeKeyNameInternal = ffe("FF10427", "could not initialize cache - '%s' is not an expected size configuration key suffix. Expected values are: 'size', 'limit'")
	MsgUnknownVerifierType                = ffe("FF10428", "Unknown verifier type", 400)
	MsgNotSupportedByBlockchainPlugin     = ffe("FF10429", "Not supported by blockchain plugin", 400)
	MsgIdempotencyKeyDuplicateMessage     = ffe("FF10430", "Idempotency key '%s' already used for message '%s'", 409)
	MsgIdempotencyKeyDuplicateTransaction = ffe("FF10431", "Idempotency key '%s' already used for transaction '%s'", 409)
	MsgNonIdempotencyKeyConflictTxInsert  = ffe("FF10432", "Conflict on insert of transaction '%s'. No existing transaction matching idempotency key '%s' found", 409)
)