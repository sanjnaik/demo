/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-09T07:46:40Z.
 *
 * API version: 1.0.9-4334
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesAddonDefinitionAllOf Definition of the list of properties defined in 'kubernetes.AddonDefinition', excluding properties defined in parent classes.
type KubernetesAddonDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the addon component.
	ChartUrl *string `json:"ChartUrl,omitempty"`
	// Default installation strategy for the release. * `None` - Unspecified install strategy. * `NoAction` - No install action performed. * `InstallOnly` - Only install in green field. No action in case of failure or removal. * `Always` - Attempt install if chart is not already installed.
	DefaultInstallStrategy *string `json:"DefaultInstallStrategy,omitempty"`
	// Default namespace to install the release.
	DefaultNamespace *string `json:"DefaultNamespace,omitempty"`
	// Default upgrade strategy for the release. * `None` - Unspecified upgrade strategy. * `NoAction` - This choice enables No upgrades to be performed. * `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure. * `AlwaysReinstall` - Always remove older release and reinstall.
	DefaultUpgradeStrategy *string `json:"DefaultUpgradeStrategy,omitempty"`
	// Description of the addon component.
	Description *string `json:"Description,omitempty"`
	// Digest used to verify the integrity of an addon.
	Digest *string `json:"Digest,omitempty"`
	// Icon used to represent the addon in UI.
	IconUrl *string  `json:"IconUrl,omitempty"`
	Labels  []string `json:"Labels,omitempty"`
	// Name of an addon component.
	Name *string `json:"Name,omitempty"`
	// Version of the addon component.
	Version              *string                               `json:"Version,omitempty"`
	Catalog              *KubernetesCatalogRelationship        `json:"Catalog,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonDefinitionAllOf KubernetesAddonDefinitionAllOf

// NewKubernetesAddonDefinitionAllOf instantiates a new KubernetesAddonDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonDefinitionAllOf(classId string, objectType string) *KubernetesAddonDefinitionAllOf {
	this := KubernetesAddonDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultInstallStrategy string = "None"
	this.DefaultInstallStrategy = &defaultInstallStrategy
	var defaultUpgradeStrategy string = "None"
	this.DefaultUpgradeStrategy = &defaultUpgradeStrategy
	return &this
}

// NewKubernetesAddonDefinitionAllOfWithDefaults instantiates a new KubernetesAddonDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonDefinitionAllOfWithDefaults() *KubernetesAddonDefinitionAllOf {
	this := KubernetesAddonDefinitionAllOf{}
	var classId string = "kubernetes.AddonDefinition"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonDefinition"
	this.ObjectType = objectType
	var defaultInstallStrategy string = "None"
	this.DefaultInstallStrategy = &defaultInstallStrategy
	var defaultUpgradeStrategy string = "None"
	this.DefaultUpgradeStrategy = &defaultUpgradeStrategy
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChartUrl returns the ChartUrl field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetChartUrl() string {
	if o == nil || o.ChartUrl == nil {
		var ret string
		return ret
	}
	return *o.ChartUrl
}

// GetChartUrlOk returns a tuple with the ChartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetChartUrlOk() (*string, bool) {
	if o == nil || o.ChartUrl == nil {
		return nil, false
	}
	return o.ChartUrl, true
}

// HasChartUrl returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasChartUrl() bool {
	if o != nil && o.ChartUrl != nil {
		return true
	}

	return false
}

// SetChartUrl gets a reference to the given string and assigns it to the ChartUrl field.
func (o *KubernetesAddonDefinitionAllOf) SetChartUrl(v string) {
	o.ChartUrl = &v
}

// GetDefaultInstallStrategy returns the DefaultInstallStrategy field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultInstallStrategy() string {
	if o == nil || o.DefaultInstallStrategy == nil {
		var ret string
		return ret
	}
	return *o.DefaultInstallStrategy
}

// GetDefaultInstallStrategyOk returns a tuple with the DefaultInstallStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultInstallStrategyOk() (*string, bool) {
	if o == nil || o.DefaultInstallStrategy == nil {
		return nil, false
	}
	return o.DefaultInstallStrategy, true
}

// HasDefaultInstallStrategy returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasDefaultInstallStrategy() bool {
	if o != nil && o.DefaultInstallStrategy != nil {
		return true
	}

	return false
}

// SetDefaultInstallStrategy gets a reference to the given string and assigns it to the DefaultInstallStrategy field.
func (o *KubernetesAddonDefinitionAllOf) SetDefaultInstallStrategy(v string) {
	o.DefaultInstallStrategy = &v
}

// GetDefaultNamespace returns the DefaultNamespace field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultNamespace() string {
	if o == nil || o.DefaultNamespace == nil {
		var ret string
		return ret
	}
	return *o.DefaultNamespace
}

// GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultNamespaceOk() (*string, bool) {
	if o == nil || o.DefaultNamespace == nil {
		return nil, false
	}
	return o.DefaultNamespace, true
}

// HasDefaultNamespace returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasDefaultNamespace() bool {
	if o != nil && o.DefaultNamespace != nil {
		return true
	}

	return false
}

// SetDefaultNamespace gets a reference to the given string and assigns it to the DefaultNamespace field.
func (o *KubernetesAddonDefinitionAllOf) SetDefaultNamespace(v string) {
	o.DefaultNamespace = &v
}

// GetDefaultUpgradeStrategy returns the DefaultUpgradeStrategy field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultUpgradeStrategy() string {
	if o == nil || o.DefaultUpgradeStrategy == nil {
		var ret string
		return ret
	}
	return *o.DefaultUpgradeStrategy
}

// GetDefaultUpgradeStrategyOk returns a tuple with the DefaultUpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetDefaultUpgradeStrategyOk() (*string, bool) {
	if o == nil || o.DefaultUpgradeStrategy == nil {
		return nil, false
	}
	return o.DefaultUpgradeStrategy, true
}

// HasDefaultUpgradeStrategy returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasDefaultUpgradeStrategy() bool {
	if o != nil && o.DefaultUpgradeStrategy != nil {
		return true
	}

	return false
}

// SetDefaultUpgradeStrategy gets a reference to the given string and assigns it to the DefaultUpgradeStrategy field.
func (o *KubernetesAddonDefinitionAllOf) SetDefaultUpgradeStrategy(v string) {
	o.DefaultUpgradeStrategy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KubernetesAddonDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *KubernetesAddonDefinitionAllOf) SetDigest(v string) {
	o.Digest = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *KubernetesAddonDefinitionAllOf) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonDefinitionAllOf) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonDefinitionAllOf) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *KubernetesAddonDefinitionAllOf) SetLabels(v []string) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesAddonDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KubernetesAddonDefinitionAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetCatalog() KubernetesCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret KubernetesCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetCatalogOk() (*KubernetesCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given KubernetesCatalogRelationship and assigns it to the Catalog field.
func (o *KubernetesAddonDefinitionAllOf) SetCatalog(v KubernetesCatalogRelationship) {
	o.Catalog = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAddonDefinitionAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonDefinitionAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAddonDefinitionAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAddonDefinitionAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesAddonDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChartUrl != nil {
		toSerialize["ChartUrl"] = o.ChartUrl
	}
	if o.DefaultInstallStrategy != nil {
		toSerialize["DefaultInstallStrategy"] = o.DefaultInstallStrategy
	}
	if o.DefaultNamespace != nil {
		toSerialize["DefaultNamespace"] = o.DefaultNamespace
	}
	if o.DefaultUpgradeStrategy != nil {
		toSerialize["DefaultUpgradeStrategy"] = o.DefaultUpgradeStrategy
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Digest != nil {
		toSerialize["Digest"] = o.Digest
	}
	if o.IconUrl != nil {
		toSerialize["IconUrl"] = o.IconUrl
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAddonDefinitionAllOf := _KubernetesAddonDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAddonDefinitionAllOf); err == nil {
		*o = KubernetesAddonDefinitionAllOf(varKubernetesAddonDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChartUrl")
		delete(additionalProperties, "DefaultInstallStrategy")
		delete(additionalProperties, "DefaultNamespace")
		delete(additionalProperties, "DefaultUpgradeStrategy")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Digest")
		delete(additionalProperties, "IconUrl")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAddonDefinitionAllOf struct {
	value *KubernetesAddonDefinitionAllOf
	isSet bool
}

func (v NullableKubernetesAddonDefinitionAllOf) Get() *KubernetesAddonDefinitionAllOf {
	return v.value
}

func (v *NullableKubernetesAddonDefinitionAllOf) Set(val *KubernetesAddonDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonDefinitionAllOf(val *KubernetesAddonDefinitionAllOf) *NullableKubernetesAddonDefinitionAllOf {
	return &NullableKubernetesAddonDefinitionAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAddonDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
