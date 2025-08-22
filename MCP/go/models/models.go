package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UnfilteredCloudEnvironmentResult represents the UnfilteredCloudEnvironmentResult schema from the OpenAPI specification
type UnfilteredCloudEnvironmentResult struct {
	Organizationid string `json:"organizationId,omitempty"`
	Updated string `json:"updated,omitempty"`
	Billableownertype int `json:"billableOwnerType,omitempty"`
	Location string `json:"location,omitempty"`
	Planid string `json:"planId,omitempty"`
	Platform string `json:"platform,omitempty"`
	Displaystorageutilizationinkb bool `json:"displayStorageUtilizationInKb,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Recentfolders []string `json:"recentFolders,omitempty"`
	Exportedbloburl string `json:"exportedBlobUrl,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Id string `json:"id,omitempty"`
	Lastused string `json:"lastUsed,omitempty"`
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Portforwardingconnection ConnectionInfoBody `json:"portForwardingConnection,omitempty"`
	Created string `json:"created,omitempty"`
	Container ContainerInfoBody `json:"container,omitempty"`
	Active string `json:"active,omitempty"`
	Storageutilizationinkb int64 `json:"storageUtilizationInKb,omitempty"`
	TypeField string `json:"type,omitempty"`
	Subscriptiondata SubscriptionData `json:"subscriptionData,omitempty"`
	Skudisplayname string `json:"skuDisplayName,omitempty"`
	Accesstoken string `json:"accessToken,omitempty"`
	Gitstatus GitStatus `json:"gitStatus,omitempty"`
	Ownerid string `json:"ownerId,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	State string `json:"state,omitempty"`
	Createfromprebuild bool `json:"createFromPrebuild,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Clientusage ClientUsageSession `json:"clientUsage,omitempty"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Laststateupdatereason string `json:"lastStateUpdateReason,omitempty"`
	Prebuildtype string `json:"prebuildType,omitempty"`
	Resourcetier int `json:"resourceTier,omitempty"`
}

// HeartBeatBody represents the HeartBeatBody schema from the OpenAPI specification
type HeartBeatBody struct {
	Timestamp string `json:"timeStamp,omitempty"`
	Agentversion string `json:"agentVersion,omitempty"`
	Collecteddatalist []CollectedData `json:"collectedDataList,omitempty"`
	Environmentid string `json:"environmentId,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
}

// EnvironmentRegistrationCallbackPayloadBody represents the EnvironmentRegistrationCallbackPayloadBody schema from the OpenAPI specification
type EnvironmentRegistrationCallbackPayloadBody struct {
	Sessionid string `json:"sessionId,omitempty"`
	Sessionpath string `json:"sessionPath,omitempty"`
}

// BillableOwnerBody represents the BillableOwnerBody schema from the OpenAPI specification
type BillableOwnerBody struct {
	Id string `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	TypeField int `json:"type,omitempty"`
}

// ScopedCreateSecretBody represents the ScopedCreateSecretBody schema from the OpenAPI specification
type ScopedCreateSecretBody struct {
	Notes string `json:"notes,omitempty"`
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	TypeField int `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Filters []SecretFilterBody `json:"filters,omitempty"`
}

// GitStatus represents the GitStatus schema from the OpenAPI specification
type GitStatus struct {
	Behind int `json:"behind,omitempty"`
	Branch string `json:"branch,omitempty"`
	Commit string `json:"commit,omitempty"`
	Hasuncommittedchanges bool `json:"hasUncommittedChanges,omitempty"`
	Hasunpushedchanges bool `json:"hasUnpushedChanges,omitempty"`
	Nogitrepo bool `json:"noGitRepo,omitempty"`
	Ahead int `json:"ahead,omitempty"`
}

// RPSubscriptionNotification represents the RPSubscriptionNotification schema from the OpenAPI specification
type RPSubscriptionNotification struct {
	Registrationdate string `json:"registrationDate,omitempty"`
	State string `json:"state,omitempty"`
	Properties RPSubscriptionProperties `json:"properties,omitempty"`
}

// ClaimVMBody represents the ClaimVMBody schema from the OpenAPI specification
type ClaimVMBody struct {
	User UserIdentity `json:"user"`
}

// NetmonCorrelationDataBody represents the NetmonCorrelationDataBody schema from the OpenAPI specification
type NetmonCorrelationDataBody struct {
	Ownerdatabaseid string `json:"ownerDatabaseId,omitempty"`
	Ownerplan string `json:"ownerPlan,omitempty"`
	Repositorydatabaseid string `json:"repositoryDatabaseId,omitempty"`
	Repositoryglobalrelayid string `json:"repositoryGlobalRelayId,omitempty"`
	Ownercreatedat string `json:"ownerCreatedAt,omitempty"`
	Ownerglobalrelayid string `json:"ownerGlobalRelayId,omitempty"`
	Repositoryprivate bool `json:"repositoryPrivate,omitempty"`
	Billableownerglobalrelayid string `json:"billableOwnerGlobalRelayId,omitempty"`
	Repositorycreatedat string `json:"repositoryCreatedAt,omitempty"`
	Billableownercreatedat string `json:"billableOwnerCreatedAt,omitempty"`
	Billableownerplan string `json:"billableOwnerPlan,omitempty"`
	Billableownerdatabaseid string `json:"billableOwnerDatabaseId,omitempty"`
}

// VMResult represents the VMResult schema from the OpenAPI specification
type VMResult struct {
	Connection VMConnectionInfo `json:"connection,omitempty"`
	Provisioningstatus ProvisioningStatusResult `json:"provisioningStatus,omitempty"`
	Status int `json:"status,omitempty"`
}

// SkuInfoResult represents the SkuInfoResult schema from the OpenAPI specification
type SkuInfoResult struct {
	Availablesettings AvailableSettingsResult `json:"availableSettings,omitempty"`
	Displayname string `json:"displayName,omitempty"`
	Name string `json:"name,omitempty"`
	Os string `json:"os,omitempty"`
}

// ChangeResourceDeletionRequestBody represents the ChangeResourceDeletionRequestBody schema from the OpenAPI specification
type ChangeResourceDeletionRequestBody struct {
	Pooltype string `json:"poolType,omitempty"`
	Region string `json:"region,omitempty"`
	Comment string `json:"comment,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Poolcode string `json:"poolCode,omitempty"`
}

// VMSpecs represents the VMSpecs schema from the OpenAPI specification
type VMSpecs struct {
	Disktype int `json:"diskType"`
	Imageresourceid string `json:"imageResourceId"`
	Size string `json:"size"`
	Subnetresourceid string `json:"subnetResourceId"`
}

// CloudEnvironmentResult represents the CloudEnvironmentResult schema from the OpenAPI specification
type CloudEnvironmentResult struct {
	Accesstoken string `json:"accessToken,omitempty"`
	Skudisplayname string `json:"skuDisplayName,omitempty"`
	Portforwardingconnection ConnectionInfoBody `json:"portForwardingConnection,omitempty"`
	Displaystorageutilizationinkb bool `json:"displayStorageUtilizationInKb,omitempty"`
	Exportedbloburl string `json:"exportedBlobUrl,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Subscriptiondata SubscriptionData `json:"subscriptionData,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Id string `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	Clientusage ClientUsageSession `json:"clientUsage,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Resourcetier int `json:"resourceTier,omitempty"`
	Lastused string `json:"lastUsed,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Recentfolders []string `json:"recentFolders,omitempty"`
	Active string `json:"active,omitempty"`
	Ownerid string `json:"ownerId,omitempty"`
	State string `json:"state,omitempty"`
	Container ContainerInfoBody `json:"container,omitempty"`
	Organizationid string `json:"organizationId,omitempty"`
	Billableownertype int `json:"billableOwnerType,omitempty"`
	TypeField string `json:"type,omitempty"`
	Laststateupdatereason string `json:"lastStateUpdateReason,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Created string `json:"created,omitempty"`
	Storageutilizationinkb int64 `json:"storageUtilizationInKb,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Platform string `json:"platform,omitempty"`
	Updated string `json:"updated,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Createfromprebuild bool `json:"createFromPrebuild,omitempty"`
	Gitstatus GitStatus `json:"gitStatus,omitempty"`
	Planid string `json:"planId,omitempty"`
	Prebuildtype string `json:"prebuildType,omitempty"`
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
}

// ClientUsageData represents the ClientUsageData schema from the OpenAPI specification
type ClientUsageData struct {
	Activeminutes int `json:"activeMinutes,omitempty"`
	Lastactivity string `json:"lastActivity,omitempty"`
}

// PoolSettingsBody represents the PoolSettingsBody schema from the OpenAPI specification
type PoolSettingsBody struct {
	Pools []Pool `json:"pools,omitempty"`
	Repoid string `json:"repoId,omitempty"`
	Storagetype int `json:"storageType,omitempty"`
	Subscription string `json:"subscription,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
}

// ProblemDetails represents the ProblemDetails schema from the OpenAPI specification
type ProblemDetails struct {
	Detail string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	Status int `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// Sku represents the Sku schema from the OpenAPI specification
type Sku struct {
	Name string `json:"name,omitempty"`
	Tier string `json:"tier,omitempty"`
}

// NotificationDataBody represents the NotificationDataBody schema from the OpenAPI specification
type NotificationDataBody struct {
	Details string `json:"details,omitempty"`
	Displaymode string `json:"displayMode,omitempty"`
	Message string `json:"message,omitempty"`
	Modal bool `json:"modal,omitempty"`
}

// StorageUsageDetail represents the StorageUsageDetail schema from the OpenAPI specification
type StorageUsageDetail struct {
	Size int `json:"size,omitempty"`
	Sizeinkb int64 `json:"sizeInKB,omitempty"`
	Sku string `json:"sku,omitempty"`
	Usage float64 `json:"usage,omitempty"`
}

// SecretFilterBody represents the SecretFilterBody schema from the OpenAPI specification
type SecretFilterBody struct {
	TypeField int `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// PoolGroupResult represents the PoolGroupResult schema from the OpenAPI specification
type PoolGroupResult struct {
	Displayname string `json:"displayName,omitempty"`
	Region int `json:"region,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// SubscriptionAccountOwner represents the SubscriptionAccountOwner schema from the OpenAPI specification
type SubscriptionAccountOwner struct {
	Email string `json:"email,omitempty"`
	Puid string `json:"puid,omitempty"`
}

// RepositoryInfoBody represents the RepositoryInfoBody schema from the OpenAPI specification
type RepositoryInfoBody struct {
	Repoid int64 `json:"repoId,omitempty"`
	Diskusage string `json:"diskUsage,omitempty"`
	Name string `json:"name,omitempty"`
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Createtype string `json:"createType,omitempty"`
	Url string `json:"url,omitempty"`
	Commitid string `json:"commitId,omitempty"`
	Owner string `json:"owner,omitempty"`
}

// PrebuildReadinessResult represents the PrebuildReadinessResult schema from the OpenAPI specification
type PrebuildReadinessResult struct {
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Repoid string `json:"repoId,omitempty"`
	Supportedskus []string `json:"supportedSkus,omitempty"`
	Templateskus []string `json:"templateSkus,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Location int `json:"location,omitempty"`
	Poolskus []string `json:"poolSkus,omitempty"`
}

// FailoverDetails represents the FailoverDetails schema from the OpenAPI specification
type FailoverDetails struct {
	Failoverenabled bool `json:"failoverEnabled,omitempty"`
	Failoverregion int `json:"failoverRegion,omitempty"`
}

// PersonalizationInfoBody represents the PersonalizationInfoBody schema from the OpenAPI specification
type PersonalizationInfoBody struct {
	Dotfilesrepository string `json:"dotfilesRepository,omitempty"`
	Dotfilestargetpath string `json:"dotfilesTargetPath,omitempty"`
	Dotfilesinstallcommand string `json:"dotfilesInstallCommand,omitempty"`
}

// SubscriptionData represents the SubscriptionData schema from the OpenAPI specification
type SubscriptionData struct {
	Computequota int `json:"computeQuota,omitempty"`
	Computeusage int `json:"computeUsage,omitempty"`
	Subscriptionid string `json:"subscriptionId,omitempty"`
	Subscriptionstate string `json:"subscriptionState,omitempty"`
}

// NetworkSettingsResourceList represents the NetworkSettingsResourceList schema from the OpenAPI specification
type NetworkSettingsResourceList struct {
	Value []NetworkSettingsResource `json:"value,omitempty"`
}

// UserIdentity represents the UserIdentity schema from the OpenAPI specification
type UserIdentity struct {
	Userprincipalname string `json:"userPrincipalName"`
}

// EnvironmentRegistrationCallbackBody represents the EnvironmentRegistrationCallbackBody schema from the OpenAPI specification
type EnvironmentRegistrationCallbackBody struct {
	Payload EnvironmentRegistrationCallbackPayloadBody `json:"payload,omitempty"`
	TypeField string `json:"type"`
}

// SeedInfoBody represents the SeedInfoBody schema from the OpenAPI specification
type SeedInfoBody struct {
	Recurseclone bool `json:"recurseClone,omitempty"`
	Repository RepositoryInfoBody `json:"repository,omitempty"`
	Seedmoniker string `json:"seedMoniker,omitempty"`
	Seedtype string `json:"seedType,omitempty"`
	Cloneurl string `json:"cloneUrl,omitempty"`
	Gitconfig GitConfigOptionsBody `json:"gitConfig,omitempty"`
}

// ProfileSpecifier represents the ProfileSpecifier schema from the OpenAPI specification
type ProfileSpecifier struct {
	Oid string `json:"oid,omitempty"`
	Provider string `json:"provider,omitempty"`
	Tid string `json:"tid,omitempty"`
}

// PlanResource represents the PlanResource schema from the OpenAPI specification
type PlanResource struct {
	Properties PlanResourceProperties `json:"properties,omitempty"`
	Provisioningstate string `json:"provisioningState,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	TypeField string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Identity PlanResourceIdentity `json:"identity,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
}

// PoolStatusResponseBody represents the PoolStatusResponseBody schema from the OpenAPI specification
type PoolStatusResponseBody struct {
	Sku string `json:"sku,omitempty"`
	Poolcode string `json:"poolCode,omitempty"`
	Allwithlatestversion bool `json:"allWithLatestVersion,omitempty"`
	Readyunassignedcount int `json:"readyUnassignedCount,omitempty"`
	Readyunassignedlatestversioncount int `json:"readyUnassignedLatestVersionCount,omitempty"`
	Readyunassignednotlatestversioncount int `json:"readyUnassignedNotLatestVersionCount,omitempty"`
	Location string `json:"location,omitempty"`
	Readyunassignednotlatestversionandidlecount int `json:"readyUnassignedNotLatestVersionAndIdleCount,omitempty"`
	Isenvironmentpool bool `json:"isEnvironmentPool,omitempty"`
	Resourcetype string `json:"resourceType,omitempty"`
}

// UpdatePoolGroupBody represents the UpdatePoolGroupBody schema from the OpenAPI specification
type UpdatePoolGroupBody struct {
	Displayname string `json:"displayName"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// NetworkSettingsResource represents the NetworkSettingsResource schema from the OpenAPI specification
type NetworkSettingsResource struct {
	Properties NetworkSettingsResourceProperties `json:"properties,omitempty"`
	Provisioningstate string `json:"provisioningState,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	TypeField string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Location string `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
}

// RefreshProfileTelemetryPropertiesResponse represents the RefreshProfileTelemetryPropertiesResponse schema from the OpenAPI specification
type RefreshProfileTelemetryPropertiesResponse struct {
	Failed []ProfileSpecifier `json:"failed,omitempty"`
	Succeeded []ProfileSpecifier `json:"succeeded,omitempty"`
}

// ClientUsageSession represents the ClientUsageSession schema from the OpenAPI specification
type ClientUsageSession struct {
	Usagedata map[string]interface{} `json:"usageData,omitempty"`
	Sessionid string `json:"sessionId,omitempty"`
}

// ExperimentalFeaturesBody represents the ExperimentalFeaturesBody schema from the OpenAPI specification
type ExperimentalFeaturesBody struct {
	Usestoragev2 bool `json:"useStorageV2,omitempty"`
	Enabledynamichttpsdetection bool `json:"enableDynamicHttpsDetection,omitempty"`
	Queueresourceallocation bool `json:"queueResourceAllocation,omitempty"`
	Useprebuildfastpathifavailable bool `json:"usePrebuildFastPathIfAvailable,omitempty"`
	Useprebuiltimages bool `json:"usePrebuiltImages,omitempty"`
}

// IssueDelegatePlanAccessTokenBody represents the IssueDelegatePlanAccessTokenBody schema from the OpenAPI specification
type IssueDelegatePlanAccessTokenBody struct {
	Identity DelegateIdentity `json:"identity,omitempty"`
	Portnumbers []int `json:"portNumbers,omitempty"`
	Scope string `json:"scope,omitempty"`
	Environmentids []string `json:"environmentIds,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}

// UnderInvestigationResponseBody represents the UnderInvestigationResponseBody schema from the OpenAPI specification
type UnderInvestigationResponseBody struct {
	Underinvestigation bool `json:"underInvestigation,omitempty"`
	Updated bool `json:"updated,omitempty"`
	Investigationstarted string `json:"investigationStarted,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
}

// PlanResourceUpdateBody represents the PlanResourceUpdateBody schema from the OpenAPI specification
type PlanResourceUpdateBody struct {
	Identity PlanResourceIdentity `json:"identity,omitempty"`
	Properties PlanResourceProperties `json:"properties,omitempty"`
}

// ReplayBillRequestBody represents the ReplayBillRequestBody schema from the OpenAPI specification
type ReplayBillRequestBody struct {
	Endtime string `json:"endTime,omitempty"`
	Starttime string `json:"startTime,omitempty"`
}

// CloudEnvironmentFolderBody represents the CloudEnvironmentFolderBody schema from the OpenAPI specification
type CloudEnvironmentFolderBody struct {
	Recentfolderpaths []string `json:"recentFolderPaths,omitempty"`
}

// PrebuildTemplateInfoResult represents the PrebuildTemplateInfoResult schema from the OpenAPI specification
type PrebuildTemplateInfoResult struct {
	Prebuildhash string `json:"prebuildHash,omitempty"`
	Templatestatus string `json:"templateStatus,omitempty"`
	Lastusedtime string `json:"lastUsedTime,omitempty"`
	Branchname string `json:"branchName,omitempty"`
	Devcontainerpath string `json:"devcontainerPath,omitempty"`
	Isprebuild bool `json:"isPrebuild,omitempty"`
	Logicalskus []string `json:"logicalSkus,omitempty"`
	Repoid int64 `json:"repoId,omitempty"`
	Commitid string `json:"commitId,omitempty"`
	Id string `json:"id,omitempty"`
	Prebuildconfigurationid string `json:"prebuildConfigurationId,omitempty"`
}

// TelemetryData represents the TelemetryData schema from the OpenAPI specification
type TelemetryData struct {
	Level string `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
	Optionalvalues map[string]interface{} `json:"optionalValues,omitempty"`
	Time string `json:"time,omitempty"`
}

// HotPoolSettings represents the HotPoolSettings schema from the OpenAPI specification
type HotPoolSettings struct {
	Size int `json:"size,omitempty"`
}

// UpdateSystemConfigurationBody represents the UpdateSystemConfigurationBody schema from the OpenAPI specification
type UpdateSystemConfigurationBody struct {
	Comment string `json:"comment,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// UpdatePrebuildTemplateVersionsBody represents the UpdatePrebuildTemplateVersionsBody schema from the OpenAPI specification
type UpdatePrebuildTemplateVersionsBody struct {
	Repoid int64 `json:"repoId"`
	Branchname string `json:"branchName"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Maxprebuildtemplateversions int `json:"maxPrebuildTemplateVersions"`
}

// Pool represents the Pool schema from the OpenAPI specification
type Pool struct {
	Pooltype int `json:"poolType,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Targetcount int64 `json:"targetCount,omitempty"`
}

// ResourceUsageDetail represents the ResourceUsageDetail schema from the OpenAPI specification
type ResourceUsageDetail struct {
	Compute []ComputeUsageDetail `json:"compute,omitempty"`
	Storage []StorageUsageDetail `json:"storage,omitempty"`
}

// DeletePrebuildTemplatesBody represents the DeletePrebuildTemplatesBody schema from the OpenAPI specification
type DeletePrebuildTemplatesBody struct {
	Branchname string `json:"branchName"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Prebuildconfigurationid int64 `json:"prebuildConfigurationId,omitempty"`
	Repoid int64 `json:"repoId"`
}

// PrebuildTemplateInfo represents the PrebuildTemplateInfo schema from the OpenAPI specification
type PrebuildTemplateInfo struct {
	Workflowrunid string `json:"workFlowRunId,omitempty"`
	Container ContainerInfo `json:"container,omitempty"`
	Prebuildconfigurationid string `json:"prebuildConfigurationId,omitempty"`
	Templatesizeingb float64 `json:"templateSizeInGB,omitempty"`
	Totaltimesavingsinseconds string `json:"totalTimeSavingsInSeconds,omitempty"`
}

// DomainUserCredentials represents the DomainUserCredentials schema from the OpenAPI specification
type DomainUserCredentials struct {
	Domain string `json:"domain"`
	Organizationalunit string `json:"organizationalUnit,omitempty"`
	Passwordsecretidentifier string `json:"passwordSecretIdentifier"`
	Username string `json:"userName"`
}

// IdentityBody represents the IdentityBody schema from the OpenAPI specification
type IdentityBody struct {
	Displayname string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
	Username string `json:"userName,omitempty"`
}

// AgentResponse represents the AgentResponse schema from the OpenAPI specification
type AgentResponse struct {
	Asseturi string `json:"assetUri,omitempty"`
	Family string `json:"family,omitempty"`
	Name string `json:"name,omitempty"`
}

// CreateCloudEnvironmentBody represents the CreateCloudEnvironmentBody schema from the OpenAPI specification
type CreateCloudEnvironmentBody struct {
	Billableowner BillableOwnerBody `json:"billableOwner,omitempty"`
	Hasdevcontainerjson bool `json:"hasDevcontainerJson,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Githubappurl string `json:"gitHubAppUrl,omitempty"`
	Identity IdentityBody `json:"identity,omitempty"`
	Devcontainerjson string `json:"devContainerJson,omitempty"`
	Workingdirectory string `json:"workingDirectory,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Experimentalfeatures ExperimentalFeaturesBody `json:"experimentalFeatures,omitempty"`
	Runtimeconstraints RuntimeConstraintsBody `json:"runtimeConstraints,omitempty"`
	Githubenvironmentendpoint string `json:"githubEnvironmentEndpoint,omitempty"`
	Containerimage string `json:"containerImage,omitempty"`
	Planid string `json:"planId,omitempty"`
	Label string `json:"label,omitempty"`
	Netmoncorrelationdata NetmonCorrelationDataBody `json:"netmonCorrelationData,omitempty"`
	Secrets []SecretDataBody `json:"secrets,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	TypeField string `json:"type"`
	Connection ConnectionInfoBody `json:"connection,omitempty"`
	Platform string `json:"platform,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Analyticstrackingid string `json:"analyticsTrackingId,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
	Personalization PersonalizationInfoBody `json:"personalization,omitempty"`
	Testaccount bool `json:"testAccount,omitempty"`
	Createasprebuild bool `json:"createAsPrebuild,omitempty"`
	Githubpfsauthendpoint string `json:"gitHubPfsAuthEndpoint,omitempty"`
	Location string `json:"location,omitempty"`
	Usertier string `json:"userTier,omitempty"`
	Friendlyname string `json:"friendlyName"`
	Githubapiurl string `json:"gitHubApiUrl,omitempty"`
}

// ContainerInfo represents the ContainerInfo schema from the OpenAPI specification
type ContainerInfo struct {
	Id string `json:"id,omitempty"`
	Imagename string `json:"imageName,omitempty"`
	Schemaversion string `json:"schemaVersion,omitempty"`
}

// PlanResourceList represents the PlanResourceList schema from the OpenAPI specification
type PlanResourceList struct {
	Value []PlanResource `json:"value,omitempty"`
}

// UpdateCloudEnvironmentBody represents the UpdateCloudEnvironmentBody schema from the OpenAPI specification
type UpdateCloudEnvironmentBody struct {
	Failoverdetails FailoverDetails `json:"failoverDetails,omitempty"`
	Friendlyname string `json:"friendlyName,omitempty"`
	Planaccesstoken string `json:"planAccessToken,omitempty"`
	Planid string `json:"planId,omitempty"`
	Skuname string `json:"skuName,omitempty"`
	Autoshutdowndelayminutes int `json:"autoShutdownDelayMinutes,omitempty"`
}

// EnvironmentStateChange represents the EnvironmentStateChange schema from the OpenAPI specification
type EnvironmentStateChange struct {
	Partitionkey string `json:"partitionKey,omitempty"`
	Time string `json:"time,omitempty"`
	Environment EnvironmentBillingInfo `json:"environment,omitempty"`
	Id string `json:"id,omitempty"`
	Newvalue int `json:"newValue,omitempty"`
	Oldvalue int `json:"oldValue,omitempty"`
}

// SystemConfigurationResponse represents the SystemConfigurationResponse schema from the OpenAPI specification
type SystemConfigurationResponse struct {
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// SubscriptionAdditionalProperties represents the SubscriptionAdditionalProperties schema from the OpenAPI specification
type SubscriptionAdditionalProperties struct {
	Billingproperties BillingProperties `json:"billingProperties,omitempty"`
	Resourceproviderproperties string `json:"resourceProviderProperties,omitempty"`
}

// PlanResourceIdentity represents the PlanResourceIdentity schema from the OpenAPI specification
type PlanResourceIdentity struct {
	Principalid string `json:"principalId,omitempty"`
	Tenantid string `json:"tenantId,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// ScopedSecretResultBody represents the ScopedSecretResultBody schema from the OpenAPI specification
type ScopedSecretResultBody struct {
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	TypeField int `json:"type,omitempty"`
	Filters []SecretFilterBody `json:"filters,omitempty"`
	Id string `json:"id,omitempty"`
	Lastmodified string `json:"lastModified,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// CreateEnvironmentPoolResourceBody represents the CreateEnvironmentPoolResourceBody schema from the OpenAPI specification
type CreateEnvironmentPoolResourceBody struct {
	Environmentoptions PrebuildEnvironmentOptions `json:"environmentOptions,omitempty"`
	Secrets []SecretDataBody `json:"secrets,omitempty"`
}

// TenantInfoResult represents the TenantInfoResult schema from the OpenAPI specification
type TenantInfoResult struct {
	Id string `json:"id,omitempty"`
}

// PlanResourceEncryptionProperties represents the PlanResourceEncryptionProperties schema from the OpenAPI specification
type PlanResourceEncryptionProperties struct {
	Keysource string `json:"keySource,omitempty"`
	Keyvaultproperties PlanResourceKeyVaultProperties `json:"keyVaultProperties,omitempty"`
}

// ComputeUsageDetail represents the ComputeUsageDetail schema from the OpenAPI specification
type ComputeUsageDetail struct {
	Sku string `json:"sku,omitempty"`
	Usage float64 `json:"usage,omitempty"`
}

// LocationsResult represents the LocationsResult schema from the OpenAPI specification
type LocationsResult struct {
	Current int `json:"current,omitempty"`
	Hostnames map[string]interface{} `json:"hostnames,omitempty"`
	Available []int `json:"available,omitempty"`
}

// PrebuildEnvironmentOptions represents the PrebuildEnvironmentOptions schema from the OpenAPI specification
type PrebuildEnvironmentOptions struct {
	Correlationid string `json:"correlationId,omitempty"`
}

// PlanResourceKeyVaultProperties represents the PlanResourceKeyVaultProperties schema from the OpenAPI specification
type PlanResourceKeyVaultProperties struct {
	Keyversion string `json:"keyVersion,omitempty"`
	Keyname string `json:"keyName,omitempty"`
	Keyvaulturi string `json:"keyVaultUri,omitempty"`
}

// CreateTemplateResult represents the CreateTemplateResult schema from the OpenAPI specification
type CreateTemplateResult struct {
	Sasurl string `json:"sasUrl,omitempty"`
	Templateid string `json:"templateId,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// PoolDefinitionInput represents the PoolDefinitionInput schema from the OpenAPI specification
type PoolDefinitionInput struct {
	TypeField int `json:"type"`
	Dimensions map[string]interface{} `json:"dimensions"`
	Isenabled bool `json:"isEnabled"`
	Location int `json:"location"`
	Logicalskus []string `json:"logicalSkus,omitempty"`
	Subtype int `json:"subtype"`
	Targetcount int `json:"targetCount"`
}

// VmLogsUploadInfo represents the VmLogsUploadInfo schema from the OpenAPI specification
type VmLogsUploadInfo struct {
	Message string `json:"message,omitempty"`
	Pathincontainer string `json:"pathInContainer,omitempty"`
	Storageuri string `json:"storageUri,omitempty"`
	Vmresourceid string `json:"vmResourceId,omitempty"`
	Containername string `json:"containerName,omitempty"`
}

// ScopedUpdateSecretBody represents the ScopedUpdateSecretBody schema from the OpenAPI specification
type ScopedUpdateSecretBody struct {
	Notes string `json:"notes,omitempty"`
	Scope int `json:"scope,omitempty"`
	Secretname string `json:"secretName,omitempty"`
	Value string `json:"value,omitempty"`
	Filters []SecretFilterBody `json:"filters,omitempty"`
}

// PoolResult represents the PoolResult schema from the OpenAPI specification
type PoolResult struct {
	Domainusercredentials DomainUserCredentials `json:"domainUserCredentials,omitempty"`
	Hotpoolsettings HotPoolSettings `json:"hotPoolSettings,omitempty"`
	Poolgroupname string `json:"poolGroupName"`
	Provisioningstatus ProvisioningStatusResult `json:"provisioningStatus,omitempty"`
	Tags map[string]interface{} `json:"tags,omitempty"`
	Usergroupname string `json:"userGroupName,omitempty"`
	Vmspecs VMSpecs `json:"vmSpecs"`
}

// CollectedData represents the CollectedData schema from the OpenAPI specification
type CollectedData struct {
	Name string `json:"name,omitempty"`
	Parentactivityid string `json:"parentActivityId,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Environmentid string `json:"environmentId,omitempty"`
}

// CreateEnvironmentStateChangeBody represents the CreateEnvironmentStateChangeBody schema from the OpenAPI specification
type CreateEnvironmentStateChangeBody struct {
	Oldvalue int `json:"oldValue,omitempty"`
	Time string `json:"time,omitempty"`
	Newvalue int `json:"newValue,omitempty"`
}

// AvailableSettingsResult represents the AvailableSettingsResult schema from the OpenAPI specification
type AvailableSettingsResult struct {
	Sku []string `json:"sku,omitempty"`
}

// UpdatePrebuildTemplateBody represents the UpdatePrebuildTemplateBody schema from the OpenAPI specification
type UpdatePrebuildTemplateBody struct {
	Issuccess bool `json:"isSuccess"`
}

// PlanResourceHeaders represents the PlanResourceHeaders schema from the OpenAPI specification
type PlanResourceHeaders struct {
	Clienttenantid string `json:"clientTenantId,omitempty"`
	Hometenantid string `json:"homeTenantId,omitempty"`
	Identityprincipalid string `json:"identityPrincipalId,omitempty"`
	Identityurl string `json:"identityUrl,omitempty"`
}

// PoolConfigRequestBody represents the PoolConfigRequestBody schema from the OpenAPI specification
type PoolConfigRequestBody struct {
	Targetcount string `json:"targetCount,omitempty"`
	Comment string `json:"comment,omitempty"`
	Maxtargetcount string `json:"maxTargetCount,omitempty"`
	Mintargetcount string `json:"minTargetCount,omitempty"`
	Poolcode string `json:"poolCode,omitempty"`
	Pooltype string `json:"poolType,omitempty"`
	Region string `json:"region,omitempty"`
}

// CreatePrebuildTemplateBody represents the CreatePrebuildTemplateBody schema from the OpenAPI specification
type CreatePrebuildTemplateBody struct {
	Friendlyname string `json:"friendlyName"`
	Githubprebuildinstanceendpoint string `json:"gitHubPrebuildInstanceEndpoint,omitempty"`
	Devcontainerpath string `json:"devContainerPath,omitempty"`
	Planid string `json:"planId,omitempty"`
	Seed SeedInfoBody `json:"seed,omitempty"`
	Templateinfo PrebuildTemplateInfo `json:"templateInfo,omitempty"`
	Features map[string]interface{} `json:"features,omitempty"`
	Githubprebuildtemplateendpoint string `json:"gitHubPrebuildTemplateEndpoint,omitempty"`
	Storagetype int `json:"storageType,omitempty"`
	Experimentalfeatures ExperimentalFeaturesBody `json:"experimentalFeatures,omitempty"`
}

// TunnelProperties represents the TunnelProperties schema from the OpenAPI specification
type TunnelProperties struct {
	Tunnelname string `json:"tunnelName,omitempty"`
	Clusterid string `json:"clusterId,omitempty"`
	Connectaccesstoken string `json:"connectAccessToken,omitempty"`
	Domain string `json:"domain,omitempty"`
	Manageportsaccesstoken string `json:"managePortsAccessToken,omitempty"`
	Serviceuri string `json:"serviceUri,omitempty"`
	Tunnelid string `json:"tunnelId,omitempty"`
}

// SecretDataBody represents the SecretDataBody schema from the OpenAPI specification
type SecretDataBody struct {
	Value string `json:"value,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField int `json:"type,omitempty"`
}

// EnvironmentBillingInfo represents the EnvironmentBillingInfo schema from the OpenAPI specification
type EnvironmentBillingInfo struct {
	Userid string `json:"userId,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Sku Sku `json:"sku,omitempty"`
}

// VsoPlanInfo represents the VsoPlanInfo schema from the OpenAPI specification
type VsoPlanInfo struct {
	Location int `json:"location,omitempty"`
	Name string `json:"name,omitempty"`
	Providernamespace string `json:"providerNamespace,omitempty"`
	Resourcegroup string `json:"resourceGroup,omitempty"`
	Resourceid string `json:"resourceId,omitempty"`
	Subscription string `json:"subscription,omitempty"`
}

// RPSubscriptionProperties represents the RPSubscriptionProperties schema from the OpenAPI specification
type RPSubscriptionProperties struct {
	Tenantid string `json:"tenantId,omitempty"`
	Accountowner SubscriptionAccountOwner `json:"accountOwner,omitempty"`
	Additionalproperties SubscriptionAdditionalProperties `json:"additionalProperties,omitempty"`
	Locationplacementid string `json:"locationPlacementId,omitempty"`
	Managedbytenants []StringStringKeyValuePair `json:"managedByTenants,omitempty"`
	Quotaid string `json:"quotaId,omitempty"`
	Registeredfeatures []StringStringKeyValuePair `json:"registeredFeatures,omitempty"`
}

// RuntimeConstraintsBody represents the RuntimeConstraintsBody schema from the OpenAPI specification
type RuntimeConstraintsBody struct {
	Allowedportprivacysettings []int `json:"allowedPortPrivacySettings,omitempty"`
	Imageallowlist []string `json:"imageAllowList,omitempty"`
}

// UpdateUserSecretsRequestBody represents the UpdateUserSecretsRequestBody schema from the OpenAPI specification
type UpdateUserSecretsRequestBody struct {
	Secrets []SecretDataBody `json:"secrets,omitempty"`
}

// VMConnectionInfo represents the VMConnectionInfo schema from the OpenAPI specification
type VMConnectionInfo struct {
	Connectiontype int `json:"connectionType,omitempty"`
	Liveshareworkspaceid string `json:"liveShareWorkspaceId,omitempty"`
}

// CreateOrUpdatePoolBody represents the CreateOrUpdatePoolBody schema from the OpenAPI specification
type CreateOrUpdatePoolBody struct {
	Tags map[string]interface{} `json:"tags,omitempty"`
	Usergroupname string `json:"userGroupName,omitempty"`
	Vmspecs VMSpecs `json:"vmSpecs"`
	Domainusercredentials DomainUserCredentials `json:"domainUserCredentials,omitempty"`
	Hotpoolsettings HotPoolSettings `json:"hotPoolSettings,omitempty"`
	Poolgroupname string `json:"poolGroupName"`
}

// EnvironmentUsage represents the EnvironmentUsage schema from the OpenAPI specification
type EnvironmentUsage struct {
	Sku Sku `json:"sku,omitempty"`
	Endstate int `json:"endState,omitempty"`
	Id string `json:"id,omitempty"`
	Resourceusage ResourceUsageDetail `json:"resourceUsage,omitempty"`
}

// LocationInfoResult represents the LocationInfoResult schema from the OpenAPI specification
type LocationInfoResult struct {
	Skus []SkuInfoResult `json:"skus,omitempty"`
}

// StringStringKeyValuePair represents the StringStringKeyValuePair schema from the OpenAPI specification
type StringStringKeyValuePair struct {
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// ConnectionInfoBody represents the ConnectionInfoBody schema from the OpenAPI specification
type ConnectionInfoBody struct {
	Connectionsessionid string `json:"connectionSessionId,omitempty"`
	Connectionsessionpath string `json:"connectionSessionPath,omitempty"`
	Hostpublickeys []string `json:"hostPublicKeys,omitempty"`
	Relayendpoint string `json:"relayEndpoint,omitempty"`
	Relaysastoken string `json:"relaySasToken,omitempty"`
	Sessiontoken string `json:"sessionToken,omitempty"`
	Tunnelproperties TunnelProperties `json:"tunnelProperties,omitempty"`
	Connectionserviceuri string `json:"connectionServiceUri,omitempty"`
}

// ProvisioningStatusResult represents the ProvisioningStatusResult schema from the OpenAPI specification
type ProvisioningStatusResult struct {
	Totalsteps int `json:"totalSteps,omitempty"`
	Completedsteps int `json:"completedSteps,omitempty"`
	Currentstepdescription string `json:"currentStepDescription,omitempty"`
	Isready bool `json:"isReady,omitempty"`
	Operationstartedtimeutc string `json:"operationStartedTimeUtc,omitempty"`
}

// AddForwardedPortSettings represents the AddForwardedPortSettings schema from the OpenAPI specification
type AddForwardedPortSettings struct {
	Privacy int `json:"privacy,omitempty"`
	Tunneltype int `json:"tunnelType,omitempty"`
}

// CreatePoolGroupBody represents the CreatePoolGroupBody schema from the OpenAPI specification
type CreatePoolGroupBody struct {
	Displayname string `json:"displayName"`
	Region int `json:"region"`
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// TunnelPortInfoResult represents the TunnelPortInfoResult schema from the OpenAPI specification
type TunnelPortInfoResult struct {
	Portvisibility string `json:"portVisibility,omitempty"`
	Tunneltoken string `json:"tunnelToken,omitempty"`
}

// BillingProperties represents the BillingProperties schema from the OpenAPI specification
type BillingProperties struct {
	Paymenttype string `json:"paymentType,omitempty"`
	Tier string `json:"tier,omitempty"`
	Workloadtype string `json:"workloadType,omitempty"`
	Billingtype string `json:"billingType,omitempty"`
	Channeltype string `json:"channelType,omitempty"`
}

// DelegateIdentity represents the DelegateIdentity schema from the OpenAPI specification
type DelegateIdentity struct {
	Username string `json:"username,omitempty"`
	Displayname string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
}

// ContainerInfoBody represents the ContainerInfoBody schema from the OpenAPI specification
type ContainerInfoBody struct {
	Schemaversion string `json:"schemaVersion,omitempty"`
	Id string `json:"id,omitempty"`
}

// NetworkSettingsResourceProperties represents the NetworkSettingsResourceProperties schema from the OpenAPI specification
type NetworkSettingsResourceProperties struct {
	Subnetid string `json:"subnetId,omitempty"`
}

// RefreshProfileTelemetryPropertiesRequest represents the RefreshProfileTelemetryPropertiesRequest schema from the OpenAPI specification
type RefreshProfileTelemetryPropertiesRequest struct {
	Partner string `json:"partner,omitempty"`
	Tenantid string `json:"tenantId,omitempty"`
	Userids string `json:"userIds,omitempty"`
}

// VnetProperties represents the VnetProperties schema from the OpenAPI specification
type VnetProperties struct {
	Subnetid string `json:"subnetId,omitempty"`
}

// GitConfigOptionsBody represents the GitConfigOptionsBody schema from the OpenAPI specification
type GitConfigOptionsBody struct {
	Username string `json:"userName,omitempty"`
	Useremail string `json:"userEmail,omitempty"`
}

// PlanResourceProperties represents the PlanResourceProperties schema from the OpenAPI specification
type PlanResourceProperties struct {
	Userid string `json:"userId,omitempty"`
	Vnetproperties VnetProperties `json:"vnetProperties,omitempty"`
	Defaultcodespacesku string `json:"defaultCodespaceSku,omitempty"`
	Defaultenvironmentsku string `json:"defaultEnvironmentSku,omitempty"`
	Encryption PlanResourceEncryptionProperties `json:"encryption,omitempty"`
}

// BillSummary represents the BillSummary schema from the OpenAPI specification
type BillSummary struct {
	Environmentid string `json:"environmentId,omitempty"`
	Location int `json:"location,omitempty"`
	Usagedetail []EnvironmentUsage `json:"usageDetail,omitempty"`
	Billgenerationtime string `json:"billGenerationTime,omitempty"`
	Periodend string `json:"periodEnd,omitempty"`
	Periodstart string `json:"periodStart,omitempty"`
	Plan VsoPlanInfo `json:"plan,omitempty"`
	Partitionkey string `json:"partitionKey,omitempty"`
	Id string `json:"id,omitempty"`
	Usage map[string]interface{} `json:"usage,omitempty"`
}
