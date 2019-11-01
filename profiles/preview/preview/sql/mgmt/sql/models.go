// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package sql

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2018-06-01-preview/sql"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CatalogCollationType = original.CatalogCollationType

const (
	DATABASEDEFAULT         CatalogCollationType = original.DATABASEDEFAULT
	SQLLatin1GeneralCP1CIAS CatalogCollationType = original.SQLLatin1GeneralCP1CIAS
)

type IdentityType = original.IdentityType

const (
	SystemAssigned IdentityType = original.SystemAssigned
)

type InstancePoolLicenseType = original.InstancePoolLicenseType

const (
	BasePrice       InstancePoolLicenseType = original.BasePrice
	LicenseIncluded InstancePoolLicenseType = original.LicenseIncluded
)

type ManagedDatabaseCreateMode = original.ManagedDatabaseCreateMode

const (
	Default               ManagedDatabaseCreateMode = original.Default
	PointInTimeRestore    ManagedDatabaseCreateMode = original.PointInTimeRestore
	Recovery              ManagedDatabaseCreateMode = original.Recovery
	RestoreExternalBackup ManagedDatabaseCreateMode = original.RestoreExternalBackup
)

type ManagedDatabaseStatus = original.ManagedDatabaseStatus

const (
	Creating     ManagedDatabaseStatus = original.Creating
	Inaccessible ManagedDatabaseStatus = original.Inaccessible
	Offline      ManagedDatabaseStatus = original.Offline
	Online       ManagedDatabaseStatus = original.Online
	Restoring    ManagedDatabaseStatus = original.Restoring
	Shutdown     ManagedDatabaseStatus = original.Shutdown
	Updating     ManagedDatabaseStatus = original.Updating
)

type ManagedInstanceLicenseType = original.ManagedInstanceLicenseType

const (
	ManagedInstanceLicenseTypeBasePrice       ManagedInstanceLicenseType = original.ManagedInstanceLicenseTypeBasePrice
	ManagedInstanceLicenseTypeLicenseIncluded ManagedInstanceLicenseType = original.ManagedInstanceLicenseTypeLicenseIncluded
)

type ManagedInstanceProxyOverride = original.ManagedInstanceProxyOverride

const (
	ManagedInstanceProxyOverrideDefault  ManagedInstanceProxyOverride = original.ManagedInstanceProxyOverrideDefault
	ManagedInstanceProxyOverrideProxy    ManagedInstanceProxyOverride = original.ManagedInstanceProxyOverrideProxy
	ManagedInstanceProxyOverrideRedirect ManagedInstanceProxyOverride = original.ManagedInstanceProxyOverrideRedirect
)

type ManagedServerCreateMode = original.ManagedServerCreateMode

const (
	ManagedServerCreateModeDefault            ManagedServerCreateMode = original.ManagedServerCreateModeDefault
	ManagedServerCreateModePointInTimeRestore ManagedServerCreateMode = original.ManagedServerCreateModePointInTimeRestore
)

type ReplicaType = original.ReplicaType

const (
	Primary           ReplicaType = original.Primary
	ReadableSecondary ReplicaType = original.ReadableSecondary
)

type SecurityAlertPolicyState = original.SecurityAlertPolicyState

const (
	SecurityAlertPolicyStateDisabled SecurityAlertPolicyState = original.SecurityAlertPolicyStateDisabled
	SecurityAlertPolicyStateEnabled  SecurityAlertPolicyState = original.SecurityAlertPolicyStateEnabled
	SecurityAlertPolicyStateNew      SecurityAlertPolicyState = original.SecurityAlertPolicyStateNew
)

type SensitivityLabelSource = original.SensitivityLabelSource

const (
	Current     SensitivityLabelSource = original.Current
	Recommended SensitivityLabelSource = original.Recommended
)

type BaseClient = original.BaseClient
type CompleteDatabaseRestoreDefinition = original.CompleteDatabaseRestoreDefinition
type DatabaseSecurityAlertListResult = original.DatabaseSecurityAlertListResult
type DatabaseSecurityAlertListResultIterator = original.DatabaseSecurityAlertListResultIterator
type DatabaseSecurityAlertListResultPage = original.DatabaseSecurityAlertListResultPage
type DatabaseSecurityAlertPoliciesClient = original.DatabaseSecurityAlertPoliciesClient
type DatabaseSecurityAlertPolicy = original.DatabaseSecurityAlertPolicy
type DatabasesClient = original.DatabasesClient
type DatabasesFailoverFuture = original.DatabasesFailoverFuture
type ElasticPoolsClient = original.ElasticPoolsClient
type ElasticPoolsFailoverFuture = original.ElasticPoolsFailoverFuture
type InstancePool = original.InstancePool
type InstancePoolListResult = original.InstancePoolListResult
type InstancePoolListResultIterator = original.InstancePoolListResultIterator
type InstancePoolListResultPage = original.InstancePoolListResultPage
type InstancePoolProperties = original.InstancePoolProperties
type InstancePoolUpdate = original.InstancePoolUpdate
type InstancePoolsClient = original.InstancePoolsClient
type InstancePoolsCreateOrUpdateFuture = original.InstancePoolsCreateOrUpdateFuture
type InstancePoolsDeleteFuture = original.InstancePoolsDeleteFuture
type InstancePoolsUpdateFuture = original.InstancePoolsUpdateFuture
type ManagedDatabase = original.ManagedDatabase
type ManagedDatabaseListResult = original.ManagedDatabaseListResult
type ManagedDatabaseListResultIterator = original.ManagedDatabaseListResultIterator
type ManagedDatabaseListResultPage = original.ManagedDatabaseListResultPage
type ManagedDatabaseProperties = original.ManagedDatabaseProperties
type ManagedDatabaseRestoreDetailsClient = original.ManagedDatabaseRestoreDetailsClient
type ManagedDatabaseRestoreDetailsProperties = original.ManagedDatabaseRestoreDetailsProperties
type ManagedDatabaseRestoreDetailsResult = original.ManagedDatabaseRestoreDetailsResult
type ManagedDatabaseSensitivityLabelsClient = original.ManagedDatabaseSensitivityLabelsClient
type ManagedDatabaseUpdate = original.ManagedDatabaseUpdate
type ManagedDatabasesClient = original.ManagedDatabasesClient
type ManagedDatabasesCompleteRestoreFuture = original.ManagedDatabasesCompleteRestoreFuture
type ManagedDatabasesCreateOrUpdateFuture = original.ManagedDatabasesCreateOrUpdateFuture
type ManagedDatabasesDeleteFuture = original.ManagedDatabasesDeleteFuture
type ManagedDatabasesUpdateFuture = original.ManagedDatabasesUpdateFuture
type ManagedInstance = original.ManagedInstance
type ManagedInstanceListResult = original.ManagedInstanceListResult
type ManagedInstanceListResultIterator = original.ManagedInstanceListResultIterator
type ManagedInstanceListResultPage = original.ManagedInstanceListResultPage
type ManagedInstanceProperties = original.ManagedInstanceProperties
type ManagedInstanceUpdate = original.ManagedInstanceUpdate
type ManagedInstanceVulnerabilityAssessment = original.ManagedInstanceVulnerabilityAssessment
type ManagedInstanceVulnerabilityAssessmentListResult = original.ManagedInstanceVulnerabilityAssessmentListResult
type ManagedInstanceVulnerabilityAssessmentListResultIterator = original.ManagedInstanceVulnerabilityAssessmentListResultIterator
type ManagedInstanceVulnerabilityAssessmentListResultPage = original.ManagedInstanceVulnerabilityAssessmentListResultPage
type ManagedInstanceVulnerabilityAssessmentProperties = original.ManagedInstanceVulnerabilityAssessmentProperties
type ManagedInstanceVulnerabilityAssessmentsClient = original.ManagedInstanceVulnerabilityAssessmentsClient
type ManagedInstancesClient = original.ManagedInstancesClient
type ManagedInstancesCreateOrUpdateFuture = original.ManagedInstancesCreateOrUpdateFuture
type ManagedInstancesDeleteFuture = original.ManagedInstancesDeleteFuture
type ManagedInstancesUpdateFuture = original.ManagedInstancesUpdateFuture
type Name = original.Name
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionListResultIterator = original.PrivateEndpointConnectionListResultIterator
type PrivateEndpointConnectionListResultPage = original.PrivateEndpointConnectionListResultPage
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointProperty = original.PrivateEndpointProperty
type PrivateLinkServiceConnectionStateProperty = original.PrivateLinkServiceConnectionStateProperty
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type SecurityAlertPolicyProperties = original.SecurityAlertPolicyProperties
type SensitivityLabel = original.SensitivityLabel
type SensitivityLabelListResult = original.SensitivityLabelListResult
type SensitivityLabelListResultIterator = original.SensitivityLabelListResultIterator
type SensitivityLabelListResultPage = original.SensitivityLabelListResultPage
type SensitivityLabelProperties = original.SensitivityLabelProperties
type ServerVulnerabilityAssessment = original.ServerVulnerabilityAssessment
type ServerVulnerabilityAssessmentListResult = original.ServerVulnerabilityAssessmentListResult
type ServerVulnerabilityAssessmentListResultIterator = original.ServerVulnerabilityAssessmentListResultIterator
type ServerVulnerabilityAssessmentListResultPage = original.ServerVulnerabilityAssessmentListResultPage
type ServerVulnerabilityAssessmentProperties = original.ServerVulnerabilityAssessmentProperties
type ServerVulnerabilityAssessmentsClient = original.ServerVulnerabilityAssessmentsClient
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageListResultIterator = original.UsageListResultIterator
type UsageListResultPage = original.UsageListResultPage
type UsagesClient = original.UsagesClient
type VulnerabilityAssessmentRecurringScansProperties = original.VulnerabilityAssessmentRecurringScansProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDatabaseSecurityAlertListResultIterator(page DatabaseSecurityAlertListResultPage) DatabaseSecurityAlertListResultIterator {
	return original.NewDatabaseSecurityAlertListResultIterator(page)
}
func NewDatabaseSecurityAlertListResultPage(getNextPage func(context.Context, DatabaseSecurityAlertListResult) (DatabaseSecurityAlertListResult, error)) DatabaseSecurityAlertListResultPage {
	return original.NewDatabaseSecurityAlertListResultPage(getNextPage)
}
func NewDatabaseSecurityAlertPoliciesClient(subscriptionID string) DatabaseSecurityAlertPoliciesClient {
	return original.NewDatabaseSecurityAlertPoliciesClient(subscriptionID)
}
func NewDatabaseSecurityAlertPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseSecurityAlertPoliciesClient {
	return original.NewDatabaseSecurityAlertPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClient(subscriptionID)
}
func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInstancePoolListResultIterator(page InstancePoolListResultPage) InstancePoolListResultIterator {
	return original.NewInstancePoolListResultIterator(page)
}
func NewInstancePoolListResultPage(getNextPage func(context.Context, InstancePoolListResult) (InstancePoolListResult, error)) InstancePoolListResultPage {
	return original.NewInstancePoolListResultPage(getNextPage)
}
func NewInstancePoolsClient(subscriptionID string) InstancePoolsClient {
	return original.NewInstancePoolsClient(subscriptionID)
}
func NewInstancePoolsClientWithBaseURI(baseURI string, subscriptionID string) InstancePoolsClient {
	return original.NewInstancePoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedDatabaseListResultIterator(page ManagedDatabaseListResultPage) ManagedDatabaseListResultIterator {
	return original.NewManagedDatabaseListResultIterator(page)
}
func NewManagedDatabaseListResultPage(getNextPage func(context.Context, ManagedDatabaseListResult) (ManagedDatabaseListResult, error)) ManagedDatabaseListResultPage {
	return original.NewManagedDatabaseListResultPage(getNextPage)
}
func NewManagedDatabaseRestoreDetailsClient(subscriptionID string) ManagedDatabaseRestoreDetailsClient {
	return original.NewManagedDatabaseRestoreDetailsClient(subscriptionID)
}
func NewManagedDatabaseRestoreDetailsClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseRestoreDetailsClient {
	return original.NewManagedDatabaseRestoreDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedDatabaseSensitivityLabelsClient(subscriptionID string) ManagedDatabaseSensitivityLabelsClient {
	return original.NewManagedDatabaseSensitivityLabelsClient(subscriptionID)
}
func NewManagedDatabaseSensitivityLabelsClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseSensitivityLabelsClient {
	return original.NewManagedDatabaseSensitivityLabelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedDatabasesClient(subscriptionID string) ManagedDatabasesClient {
	return original.NewManagedDatabasesClient(subscriptionID)
}
func NewManagedDatabasesClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabasesClient {
	return original.NewManagedDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedInstanceListResultIterator(page ManagedInstanceListResultPage) ManagedInstanceListResultIterator {
	return original.NewManagedInstanceListResultIterator(page)
}
func NewManagedInstanceListResultPage(getNextPage func(context.Context, ManagedInstanceListResult) (ManagedInstanceListResult, error)) ManagedInstanceListResultPage {
	return original.NewManagedInstanceListResultPage(getNextPage)
}
func NewManagedInstanceVulnerabilityAssessmentListResultIterator(page ManagedInstanceVulnerabilityAssessmentListResultPage) ManagedInstanceVulnerabilityAssessmentListResultIterator {
	return original.NewManagedInstanceVulnerabilityAssessmentListResultIterator(page)
}
func NewManagedInstanceVulnerabilityAssessmentListResultPage(getNextPage func(context.Context, ManagedInstanceVulnerabilityAssessmentListResult) (ManagedInstanceVulnerabilityAssessmentListResult, error)) ManagedInstanceVulnerabilityAssessmentListResultPage {
	return original.NewManagedInstanceVulnerabilityAssessmentListResultPage(getNextPage)
}
func NewManagedInstanceVulnerabilityAssessmentsClient(subscriptionID string) ManagedInstanceVulnerabilityAssessmentsClient {
	return original.NewManagedInstanceVulnerabilityAssessmentsClient(subscriptionID)
}
func NewManagedInstanceVulnerabilityAssessmentsClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstanceVulnerabilityAssessmentsClient {
	return original.NewManagedInstanceVulnerabilityAssessmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedInstancesClient(subscriptionID string) ManagedInstancesClient {
	return original.NewManagedInstancesClient(subscriptionID)
}
func NewManagedInstancesClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstancesClient {
	return original.NewManagedInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionListResultIterator(page PrivateEndpointConnectionListResultPage) PrivateEndpointConnectionListResultIterator {
	return original.NewPrivateEndpointConnectionListResultIterator(page)
}
func NewPrivateEndpointConnectionListResultPage(getNextPage func(context.Context, PrivateEndpointConnectionListResult) (PrivateEndpointConnectionListResult, error)) PrivateEndpointConnectionListResultPage {
	return original.NewPrivateEndpointConnectionListResultPage(getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSensitivityLabelListResultIterator(page SensitivityLabelListResultPage) SensitivityLabelListResultIterator {
	return original.NewSensitivityLabelListResultIterator(page)
}
func NewSensitivityLabelListResultPage(getNextPage func(context.Context, SensitivityLabelListResult) (SensitivityLabelListResult, error)) SensitivityLabelListResultPage {
	return original.NewSensitivityLabelListResultPage(getNextPage)
}
func NewServerVulnerabilityAssessmentListResultIterator(page ServerVulnerabilityAssessmentListResultPage) ServerVulnerabilityAssessmentListResultIterator {
	return original.NewServerVulnerabilityAssessmentListResultIterator(page)
}
func NewServerVulnerabilityAssessmentListResultPage(getNextPage func(context.Context, ServerVulnerabilityAssessmentListResult) (ServerVulnerabilityAssessmentListResult, error)) ServerVulnerabilityAssessmentListResultPage {
	return original.NewServerVulnerabilityAssessmentListResultPage(getNextPage)
}
func NewServerVulnerabilityAssessmentsClient(subscriptionID string) ServerVulnerabilityAssessmentsClient {
	return original.NewServerVulnerabilityAssessmentsClient(subscriptionID)
}
func NewServerVulnerabilityAssessmentsClientWithBaseURI(baseURI string, subscriptionID string) ServerVulnerabilityAssessmentsClient {
	return original.NewServerVulnerabilityAssessmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageListResultIterator(page UsageListResultPage) UsageListResultIterator {
	return original.NewUsageListResultIterator(page)
}
func NewUsageListResultPage(getNextPage func(context.Context, UsageListResult) (UsageListResult, error)) UsageListResultPage {
	return original.NewUsageListResultPage(getNextPage)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCatalogCollationTypeValues() []CatalogCollationType {
	return original.PossibleCatalogCollationTypeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleInstancePoolLicenseTypeValues() []InstancePoolLicenseType {
	return original.PossibleInstancePoolLicenseTypeValues()
}
func PossibleManagedDatabaseCreateModeValues() []ManagedDatabaseCreateMode {
	return original.PossibleManagedDatabaseCreateModeValues()
}
func PossibleManagedDatabaseStatusValues() []ManagedDatabaseStatus {
	return original.PossibleManagedDatabaseStatusValues()
}
func PossibleManagedInstanceLicenseTypeValues() []ManagedInstanceLicenseType {
	return original.PossibleManagedInstanceLicenseTypeValues()
}
func PossibleManagedInstanceProxyOverrideValues() []ManagedInstanceProxyOverride {
	return original.PossibleManagedInstanceProxyOverrideValues()
}
func PossibleManagedServerCreateModeValues() []ManagedServerCreateMode {
	return original.PossibleManagedServerCreateModeValues()
}
func PossibleReplicaTypeValues() []ReplicaType {
	return original.PossibleReplicaTypeValues()
}
func PossibleSecurityAlertPolicyStateValues() []SecurityAlertPolicyState {
	return original.PossibleSecurityAlertPolicyStateValues()
}
func PossibleSensitivityLabelSourceValues() []SensitivityLabelSource {
	return original.PossibleSensitivityLabelSourceValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
