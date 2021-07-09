/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesClusterProfileAllOf Definition of the list of properties defined in 'kubernetes.ClusterProfile', excluding properties defined in parent classes.
type KubernetesClusterProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                                            `json:"ObjectType"`
	ActionInfo      NullableKubernetesActionInfo                      `json:"ActionInfo,omitempty"`
	CertConfig      NullableKubernetesClusterCertificateConfiguration `json:"CertConfig,omitempty"`
	EssentialAddons []KubernetesEssentialAddon                        `json:"EssentialAddons,omitempty"`
	KubeConfig      NullableKubernetesConfiguration                   `json:"KubeConfig,omitempty"`
	// Management mode for the cluster. In some cases Intersight kubernetes service is not required to provision and manage the management entities and endpoints (for e.g. EKS). In most other cases it will be required to provision and manage these entities and endpoints. * `Provided` - Cluster management entities and endpoints are provided by the infrastructure platform. * `Managed` - Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service.
	ManagedMode      *string                                   `json:"ManagedMode,omitempty"`
	ManagementConfig NullableKubernetesClusterManagementConfig `json:"ManagementConfig,omitempty"`
	// Status of the Kubernetes cluster and its nodes. * `Undeployed` - The cluster is undeployed. * `Configuring` - The cluster is being configured. * `Deploying` - The cluster is being deployed. * `Undeploying` - The cluster is being undeployed. * `DeployFailedTerminal` - The cluster deployment failed terminally and can not be recovered. * `DeployFailed` - The cluster deployment failed. * `Upgrading` - The cluster is being upgraded. * `Deleting` - The cluster is being deleted. * `DeleteFailed` - The cluster delete failed. * `Ready` - The cluster is ready for use. * `Active` - The cluster is being active. * `Shutdown` - All the nodes in the cluster are powered off. * `Terminated` - The cluster is terminated. * `Deployed` - The cluster is deployed. The cluster may not yet be ready for use. * `UndeployFailed` - The cluster undeploy action failed. * `NotReady` - The cluster is created and some nodes are not ready.
	Status            *string                              `json:"Status,omitempty"`
	AciCniProfile     *KubernetesAciCniProfileRelationship `json:"AciCniProfile,omitempty"`
	AssociatedCluster *KubernetesClusterRelationship       `json:"AssociatedCluster,omitempty"`
	// An array of relationships to ippoolPool resources.
	ClusterIpPools         []IppoolPoolRelationship                      `json:"ClusterIpPools,omitempty"`
	ContainerRuntimeConfig *KubernetesContainerRuntimePolicyRelationship `json:"ContainerRuntimeConfig,omitempty"`
	// An array of relationships to ippoolIpLease resources.
	LoadbalancerIpLeases []IppoolIpLeaseRelationship          `json:"LoadbalancerIpLeases,omitempty"`
	MasterVipLease       *IppoolIpLeaseRelationship           `json:"MasterVipLease,omitempty"`
	NetConfig            *KubernetesNetworkPolicyRelationship `json:"NetConfig,omitempty"`
	// An array of relationships to kubernetesNodeGroupProfile resources.
	NodeGroups           []KubernetesNodeGroupProfileRelationship       `json:"NodeGroups,omitempty"`
	Organization         *OrganizationOrganizationRelationship          `json:"Organization,omitempty"`
	SysConfig            *KubernetesSysConfigPolicyRelationship         `json:"SysConfig,omitempty"`
	TrustedRegistries    *KubernetesTrustedRegistriesPolicyRelationship `json:"TrustedRegistries,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship              `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterProfileAllOf KubernetesClusterProfileAllOf

// NewKubernetesClusterProfileAllOf instantiates a new KubernetesClusterProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterProfileAllOf(classId string, objectType string) *KubernetesClusterProfileAllOf {
	this := KubernetesClusterProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var managedMode string = "Provided"
	this.ManagedMode = &managedMode
	var status string = "Undeployed"
	this.Status = &status
	return &this
}

// NewKubernetesClusterProfileAllOfWithDefaults instantiates a new KubernetesClusterProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterProfileAllOfWithDefaults() *KubernetesClusterProfileAllOf {
	this := KubernetesClusterProfileAllOf{}
	var classId string = "kubernetes.ClusterProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterProfile"
	this.ObjectType = objectType
	var managedMode string = "Provided"
	this.ManagedMode = &managedMode
	var status string = "Undeployed"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionInfo returns the ActionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetActionInfo() KubernetesActionInfo {
	if o == nil || o.ActionInfo.Get() == nil {
		var ret KubernetesActionInfo
		return ret
	}
	return *o.ActionInfo.Get()
}

// GetActionInfoOk returns a tuple with the ActionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetActionInfoOk() (*KubernetesActionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionInfo.Get(), o.ActionInfo.IsSet()
}

// HasActionInfo returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasActionInfo() bool {
	if o != nil && o.ActionInfo.IsSet() {
		return true
	}

	return false
}

// SetActionInfo gets a reference to the given NullableKubernetesActionInfo and assigns it to the ActionInfo field.
func (o *KubernetesClusterProfileAllOf) SetActionInfo(v KubernetesActionInfo) {
	o.ActionInfo.Set(&v)
}

// SetActionInfoNil sets the value for ActionInfo to be an explicit nil
func (o *KubernetesClusterProfileAllOf) SetActionInfoNil() {
	o.ActionInfo.Set(nil)
}

// UnsetActionInfo ensures that no value is present for ActionInfo, not even an explicit nil
func (o *KubernetesClusterProfileAllOf) UnsetActionInfo() {
	o.ActionInfo.Unset()
}

// GetCertConfig returns the CertConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetCertConfig() KubernetesClusterCertificateConfiguration {
	if o == nil || o.CertConfig.Get() == nil {
		var ret KubernetesClusterCertificateConfiguration
		return ret
	}
	return *o.CertConfig.Get()
}

// GetCertConfigOk returns a tuple with the CertConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetCertConfigOk() (*KubernetesClusterCertificateConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertConfig.Get(), o.CertConfig.IsSet()
}

// HasCertConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasCertConfig() bool {
	if o != nil && o.CertConfig.IsSet() {
		return true
	}

	return false
}

// SetCertConfig gets a reference to the given NullableKubernetesClusterCertificateConfiguration and assigns it to the CertConfig field.
func (o *KubernetesClusterProfileAllOf) SetCertConfig(v KubernetesClusterCertificateConfiguration) {
	o.CertConfig.Set(&v)
}

// SetCertConfigNil sets the value for CertConfig to be an explicit nil
func (o *KubernetesClusterProfileAllOf) SetCertConfigNil() {
	o.CertConfig.Set(nil)
}

// UnsetCertConfig ensures that no value is present for CertConfig, not even an explicit nil
func (o *KubernetesClusterProfileAllOf) UnsetCertConfig() {
	o.CertConfig.Unset()
}

// GetEssentialAddons returns the EssentialAddons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetEssentialAddons() []KubernetesEssentialAddon {
	if o == nil {
		var ret []KubernetesEssentialAddon
		return ret
	}
	return o.EssentialAddons
}

// GetEssentialAddonsOk returns a tuple with the EssentialAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetEssentialAddonsOk() (*[]KubernetesEssentialAddon, bool) {
	if o == nil || o.EssentialAddons == nil {
		return nil, false
	}
	return &o.EssentialAddons, true
}

// HasEssentialAddons returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasEssentialAddons() bool {
	if o != nil && o.EssentialAddons != nil {
		return true
	}

	return false
}

// SetEssentialAddons gets a reference to the given []KubernetesEssentialAddon and assigns it to the EssentialAddons field.
func (o *KubernetesClusterProfileAllOf) SetEssentialAddons(v []KubernetesEssentialAddon) {
	o.EssentialAddons = v
}

// GetKubeConfig returns the KubeConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetKubeConfig() KubernetesConfiguration {
	if o == nil || o.KubeConfig.Get() == nil {
		var ret KubernetesConfiguration
		return ret
	}
	return *o.KubeConfig.Get()
}

// GetKubeConfigOk returns a tuple with the KubeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetKubeConfigOk() (*KubernetesConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubeConfig.Get(), o.KubeConfig.IsSet()
}

// HasKubeConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasKubeConfig() bool {
	if o != nil && o.KubeConfig.IsSet() {
		return true
	}

	return false
}

// SetKubeConfig gets a reference to the given NullableKubernetesConfiguration and assigns it to the KubeConfig field.
func (o *KubernetesClusterProfileAllOf) SetKubeConfig(v KubernetesConfiguration) {
	o.KubeConfig.Set(&v)
}

// SetKubeConfigNil sets the value for KubeConfig to be an explicit nil
func (o *KubernetesClusterProfileAllOf) SetKubeConfigNil() {
	o.KubeConfig.Set(nil)
}

// UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
func (o *KubernetesClusterProfileAllOf) UnsetKubeConfig() {
	o.KubeConfig.Unset()
}

// GetManagedMode returns the ManagedMode field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetManagedMode() string {
	if o == nil || o.ManagedMode == nil {
		var ret string
		return ret
	}
	return *o.ManagedMode
}

// GetManagedModeOk returns a tuple with the ManagedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetManagedModeOk() (*string, bool) {
	if o == nil || o.ManagedMode == nil {
		return nil, false
	}
	return o.ManagedMode, true
}

// HasManagedMode returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasManagedMode() bool {
	if o != nil && o.ManagedMode != nil {
		return true
	}

	return false
}

// SetManagedMode gets a reference to the given string and assigns it to the ManagedMode field.
func (o *KubernetesClusterProfileAllOf) SetManagedMode(v string) {
	o.ManagedMode = &v
}

// GetManagementConfig returns the ManagementConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetManagementConfig() KubernetesClusterManagementConfig {
	if o == nil || o.ManagementConfig.Get() == nil {
		var ret KubernetesClusterManagementConfig
		return ret
	}
	return *o.ManagementConfig.Get()
}

// GetManagementConfigOk returns a tuple with the ManagementConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetManagementConfigOk() (*KubernetesClusterManagementConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementConfig.Get(), o.ManagementConfig.IsSet()
}

// HasManagementConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasManagementConfig() bool {
	if o != nil && o.ManagementConfig.IsSet() {
		return true
	}

	return false
}

// SetManagementConfig gets a reference to the given NullableKubernetesClusterManagementConfig and assigns it to the ManagementConfig field.
func (o *KubernetesClusterProfileAllOf) SetManagementConfig(v KubernetesClusterManagementConfig) {
	o.ManagementConfig.Set(&v)
}

// SetManagementConfigNil sets the value for ManagementConfig to be an explicit nil
func (o *KubernetesClusterProfileAllOf) SetManagementConfigNil() {
	o.ManagementConfig.Set(nil)
}

// UnsetManagementConfig ensures that no value is present for ManagementConfig, not even an explicit nil
func (o *KubernetesClusterProfileAllOf) UnsetManagementConfig() {
	o.ManagementConfig.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KubernetesClusterProfileAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetAciCniProfile returns the AciCniProfile field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetAciCniProfile() KubernetesAciCniProfileRelationship {
	if o == nil || o.AciCniProfile == nil {
		var ret KubernetesAciCniProfileRelationship
		return ret
	}
	return *o.AciCniProfile
}

// GetAciCniProfileOk returns a tuple with the AciCniProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetAciCniProfileOk() (*KubernetesAciCniProfileRelationship, bool) {
	if o == nil || o.AciCniProfile == nil {
		return nil, false
	}
	return o.AciCniProfile, true
}

// HasAciCniProfile returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasAciCniProfile() bool {
	if o != nil && o.AciCniProfile != nil {
		return true
	}

	return false
}

// SetAciCniProfile gets a reference to the given KubernetesAciCniProfileRelationship and assigns it to the AciCniProfile field.
func (o *KubernetesClusterProfileAllOf) SetAciCniProfile(v KubernetesAciCniProfileRelationship) {
	o.AciCniProfile = &v
}

// GetAssociatedCluster returns the AssociatedCluster field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetAssociatedCluster() KubernetesClusterRelationship {
	if o == nil || o.AssociatedCluster == nil {
		var ret KubernetesClusterRelationship
		return ret
	}
	return *o.AssociatedCluster
}

// GetAssociatedClusterOk returns a tuple with the AssociatedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool) {
	if o == nil || o.AssociatedCluster == nil {
		return nil, false
	}
	return o.AssociatedCluster, true
}

// HasAssociatedCluster returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasAssociatedCluster() bool {
	if o != nil && o.AssociatedCluster != nil {
		return true
	}

	return false
}

// SetAssociatedCluster gets a reference to the given KubernetesClusterRelationship and assigns it to the AssociatedCluster field.
func (o *KubernetesClusterProfileAllOf) SetAssociatedCluster(v KubernetesClusterRelationship) {
	o.AssociatedCluster = &v
}

// GetClusterIpPools returns the ClusterIpPools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetClusterIpPools() []IppoolPoolRelationship {
	if o == nil {
		var ret []IppoolPoolRelationship
		return ret
	}
	return o.ClusterIpPools
}

// GetClusterIpPoolsOk returns a tuple with the ClusterIpPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetClusterIpPoolsOk() (*[]IppoolPoolRelationship, bool) {
	if o == nil || o.ClusterIpPools == nil {
		return nil, false
	}
	return &o.ClusterIpPools, true
}

// HasClusterIpPools returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasClusterIpPools() bool {
	if o != nil && o.ClusterIpPools != nil {
		return true
	}

	return false
}

// SetClusterIpPools gets a reference to the given []IppoolPoolRelationship and assigns it to the ClusterIpPools field.
func (o *KubernetesClusterProfileAllOf) SetClusterIpPools(v []IppoolPoolRelationship) {
	o.ClusterIpPools = v
}

// GetContainerRuntimeConfig returns the ContainerRuntimeConfig field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetContainerRuntimeConfig() KubernetesContainerRuntimePolicyRelationship {
	if o == nil || o.ContainerRuntimeConfig == nil {
		var ret KubernetesContainerRuntimePolicyRelationship
		return ret
	}
	return *o.ContainerRuntimeConfig
}

// GetContainerRuntimeConfigOk returns a tuple with the ContainerRuntimeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetContainerRuntimeConfigOk() (*KubernetesContainerRuntimePolicyRelationship, bool) {
	if o == nil || o.ContainerRuntimeConfig == nil {
		return nil, false
	}
	return o.ContainerRuntimeConfig, true
}

// HasContainerRuntimeConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasContainerRuntimeConfig() bool {
	if o != nil && o.ContainerRuntimeConfig != nil {
		return true
	}

	return false
}

// SetContainerRuntimeConfig gets a reference to the given KubernetesContainerRuntimePolicyRelationship and assigns it to the ContainerRuntimeConfig field.
func (o *KubernetesClusterProfileAllOf) SetContainerRuntimeConfig(v KubernetesContainerRuntimePolicyRelationship) {
	o.ContainerRuntimeConfig = &v
}

// GetLoadbalancerIpLeases returns the LoadbalancerIpLeases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetLoadbalancerIpLeases() []IppoolIpLeaseRelationship {
	if o == nil {
		var ret []IppoolIpLeaseRelationship
		return ret
	}
	return o.LoadbalancerIpLeases
}

// GetLoadbalancerIpLeasesOk returns a tuple with the LoadbalancerIpLeases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetLoadbalancerIpLeasesOk() (*[]IppoolIpLeaseRelationship, bool) {
	if o == nil || o.LoadbalancerIpLeases == nil {
		return nil, false
	}
	return &o.LoadbalancerIpLeases, true
}

// HasLoadbalancerIpLeases returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasLoadbalancerIpLeases() bool {
	if o != nil && o.LoadbalancerIpLeases != nil {
		return true
	}

	return false
}

// SetLoadbalancerIpLeases gets a reference to the given []IppoolIpLeaseRelationship and assigns it to the LoadbalancerIpLeases field.
func (o *KubernetesClusterProfileAllOf) SetLoadbalancerIpLeases(v []IppoolIpLeaseRelationship) {
	o.LoadbalancerIpLeases = v
}

// GetMasterVipLease returns the MasterVipLease field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetMasterVipLease() IppoolIpLeaseRelationship {
	if o == nil || o.MasterVipLease == nil {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.MasterVipLease
}

// GetMasterVipLeaseOk returns a tuple with the MasterVipLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetMasterVipLeaseOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil || o.MasterVipLease == nil {
		return nil, false
	}
	return o.MasterVipLease, true
}

// HasMasterVipLease returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasMasterVipLease() bool {
	if o != nil && o.MasterVipLease != nil {
		return true
	}

	return false
}

// SetMasterVipLease gets a reference to the given IppoolIpLeaseRelationship and assigns it to the MasterVipLease field.
func (o *KubernetesClusterProfileAllOf) SetMasterVipLease(v IppoolIpLeaseRelationship) {
	o.MasterVipLease = &v
}

// GetNetConfig returns the NetConfig field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetNetConfig() KubernetesNetworkPolicyRelationship {
	if o == nil || o.NetConfig == nil {
		var ret KubernetesNetworkPolicyRelationship
		return ret
	}
	return *o.NetConfig
}

// GetNetConfigOk returns a tuple with the NetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetNetConfigOk() (*KubernetesNetworkPolicyRelationship, bool) {
	if o == nil || o.NetConfig == nil {
		return nil, false
	}
	return o.NetConfig, true
}

// HasNetConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasNetConfig() bool {
	if o != nil && o.NetConfig != nil {
		return true
	}

	return false
}

// SetNetConfig gets a reference to the given KubernetesNetworkPolicyRelationship and assigns it to the NetConfig field.
func (o *KubernetesClusterProfileAllOf) SetNetConfig(v KubernetesNetworkPolicyRelationship) {
	o.NetConfig = &v
}

// GetNodeGroups returns the NodeGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterProfileAllOf) GetNodeGroups() []KubernetesNodeGroupProfileRelationship {
	if o == nil {
		var ret []KubernetesNodeGroupProfileRelationship
		return ret
	}
	return o.NodeGroups
}

// GetNodeGroupsOk returns a tuple with the NodeGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProfileAllOf) GetNodeGroupsOk() (*[]KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.NodeGroups == nil {
		return nil, false
	}
	return &o.NodeGroups, true
}

// HasNodeGroups returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasNodeGroups() bool {
	if o != nil && o.NodeGroups != nil {
		return true
	}

	return false
}

// SetNodeGroups gets a reference to the given []KubernetesNodeGroupProfileRelationship and assigns it to the NodeGroups field.
func (o *KubernetesClusterProfileAllOf) SetNodeGroups(v []KubernetesNodeGroupProfileRelationship) {
	o.NodeGroups = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesClusterProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSysConfig returns the SysConfig field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetSysConfig() KubernetesSysConfigPolicyRelationship {
	if o == nil || o.SysConfig == nil {
		var ret KubernetesSysConfigPolicyRelationship
		return ret
	}
	return *o.SysConfig
}

// GetSysConfigOk returns a tuple with the SysConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetSysConfigOk() (*KubernetesSysConfigPolicyRelationship, bool) {
	if o == nil || o.SysConfig == nil {
		return nil, false
	}
	return o.SysConfig, true
}

// HasSysConfig returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasSysConfig() bool {
	if o != nil && o.SysConfig != nil {
		return true
	}

	return false
}

// SetSysConfig gets a reference to the given KubernetesSysConfigPolicyRelationship and assigns it to the SysConfig field.
func (o *KubernetesClusterProfileAllOf) SetSysConfig(v KubernetesSysConfigPolicyRelationship) {
	o.SysConfig = &v
}

// GetTrustedRegistries returns the TrustedRegistries field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetTrustedRegistries() KubernetesTrustedRegistriesPolicyRelationship {
	if o == nil || o.TrustedRegistries == nil {
		var ret KubernetesTrustedRegistriesPolicyRelationship
		return ret
	}
	return *o.TrustedRegistries
}

// GetTrustedRegistriesOk returns a tuple with the TrustedRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetTrustedRegistriesOk() (*KubernetesTrustedRegistriesPolicyRelationship, bool) {
	if o == nil || o.TrustedRegistries == nil {
		return nil, false
	}
	return o.TrustedRegistries, true
}

// HasTrustedRegistries returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasTrustedRegistries() bool {
	if o != nil && o.TrustedRegistries != nil {
		return true
	}

	return false
}

// SetTrustedRegistries gets a reference to the given KubernetesTrustedRegistriesPolicyRelationship and assigns it to the TrustedRegistries field.
func (o *KubernetesClusterProfileAllOf) SetTrustedRegistries(v KubernetesTrustedRegistriesPolicyRelationship) {
	o.TrustedRegistries = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *KubernetesClusterProfileAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProfileAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *KubernetesClusterProfileAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *KubernetesClusterProfileAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o KubernetesClusterProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActionInfo.IsSet() {
		toSerialize["ActionInfo"] = o.ActionInfo.Get()
	}
	if o.CertConfig.IsSet() {
		toSerialize["CertConfig"] = o.CertConfig.Get()
	}
	if o.EssentialAddons != nil {
		toSerialize["EssentialAddons"] = o.EssentialAddons
	}
	if o.KubeConfig.IsSet() {
		toSerialize["KubeConfig"] = o.KubeConfig.Get()
	}
	if o.ManagedMode != nil {
		toSerialize["ManagedMode"] = o.ManagedMode
	}
	if o.ManagementConfig.IsSet() {
		toSerialize["ManagementConfig"] = o.ManagementConfig.Get()
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.AciCniProfile != nil {
		toSerialize["AciCniProfile"] = o.AciCniProfile
	}
	if o.AssociatedCluster != nil {
		toSerialize["AssociatedCluster"] = o.AssociatedCluster
	}
	if o.ClusterIpPools != nil {
		toSerialize["ClusterIpPools"] = o.ClusterIpPools
	}
	if o.ContainerRuntimeConfig != nil {
		toSerialize["ContainerRuntimeConfig"] = o.ContainerRuntimeConfig
	}
	if o.LoadbalancerIpLeases != nil {
		toSerialize["LoadbalancerIpLeases"] = o.LoadbalancerIpLeases
	}
	if o.MasterVipLease != nil {
		toSerialize["MasterVipLease"] = o.MasterVipLease
	}
	if o.NetConfig != nil {
		toSerialize["NetConfig"] = o.NetConfig
	}
	if o.NodeGroups != nil {
		toSerialize["NodeGroups"] = o.NodeGroups
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.SysConfig != nil {
		toSerialize["SysConfig"] = o.SysConfig
	}
	if o.TrustedRegistries != nil {
		toSerialize["TrustedRegistries"] = o.TrustedRegistries
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterProfileAllOf := _KubernetesClusterProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterProfileAllOf); err == nil {
		*o = KubernetesClusterProfileAllOf(varKubernetesClusterProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionInfo")
		delete(additionalProperties, "CertConfig")
		delete(additionalProperties, "EssentialAddons")
		delete(additionalProperties, "KubeConfig")
		delete(additionalProperties, "ManagedMode")
		delete(additionalProperties, "ManagementConfig")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "AciCniProfile")
		delete(additionalProperties, "AssociatedCluster")
		delete(additionalProperties, "ClusterIpPools")
		delete(additionalProperties, "ContainerRuntimeConfig")
		delete(additionalProperties, "LoadbalancerIpLeases")
		delete(additionalProperties, "MasterVipLease")
		delete(additionalProperties, "NetConfig")
		delete(additionalProperties, "NodeGroups")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "SysConfig")
		delete(additionalProperties, "TrustedRegistries")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterProfileAllOf struct {
	value *KubernetesClusterProfileAllOf
	isSet bool
}

func (v NullableKubernetesClusterProfileAllOf) Get() *KubernetesClusterProfileAllOf {
	return v.value
}

func (v *NullableKubernetesClusterProfileAllOf) Set(val *KubernetesClusterProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterProfileAllOf(val *KubernetesClusterProfileAllOf) *NullableKubernetesClusterProfileAllOf {
	return &NullableKubernetesClusterProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
