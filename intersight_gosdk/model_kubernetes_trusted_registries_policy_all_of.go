/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-15T06:27:08Z.
 *
 * API version: 1.0.9-4247
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KubernetesTrustedRegistriesPolicyAllOf Definition of the list of properties defined in 'kubernetes.TrustedRegistriesPolicy', excluding properties defined in parent classes.
type KubernetesTrustedRegistriesPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string   `json:"ObjectType"`
	RootCaRegistries   []string `json:"RootCaRegistries,omitempty"`
	UnsignedRegistries []string `json:"UnsignedRegistries,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesTrustedRegistriesPolicyAllOf KubernetesTrustedRegistriesPolicyAllOf

// NewKubernetesTrustedRegistriesPolicyAllOf instantiates a new KubernetesTrustedRegistriesPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesTrustedRegistriesPolicyAllOf(classId string, objectType string) *KubernetesTrustedRegistriesPolicyAllOf {
	this := KubernetesTrustedRegistriesPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesTrustedRegistriesPolicyAllOfWithDefaults instantiates a new KubernetesTrustedRegistriesPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesTrustedRegistriesPolicyAllOfWithDefaults() *KubernetesTrustedRegistriesPolicyAllOf {
	this := KubernetesTrustedRegistriesPolicyAllOf{}
	var classId string = "kubernetes.TrustedRegistriesPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.TrustedRegistriesPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRootCaRegistries returns the RootCaRegistries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetRootCaRegistries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RootCaRegistries
}

// GetRootCaRegistriesOk returns a tuple with the RootCaRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetRootCaRegistriesOk() (*[]string, bool) {
	if o == nil || o.RootCaRegistries == nil {
		return nil, false
	}
	return &o.RootCaRegistries, true
}

// HasRootCaRegistries returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) HasRootCaRegistries() bool {
	if o != nil && o.RootCaRegistries != nil {
		return true
	}

	return false
}

// SetRootCaRegistries gets a reference to the given []string and assigns it to the RootCaRegistries field.
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetRootCaRegistries(v []string) {
	o.RootCaRegistries = v
}

// GetUnsignedRegistries returns the UnsignedRegistries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetUnsignedRegistries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UnsignedRegistries
}

// GetUnsignedRegistriesOk returns a tuple with the UnsignedRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetUnsignedRegistriesOk() (*[]string, bool) {
	if o == nil || o.UnsignedRegistries == nil {
		return nil, false
	}
	return &o.UnsignedRegistries, true
}

// HasUnsignedRegistries returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) HasUnsignedRegistries() bool {
	if o != nil && o.UnsignedRegistries != nil {
		return true
	}

	return false
}

// SetUnsignedRegistries gets a reference to the given []string and assigns it to the UnsignedRegistries field.
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetUnsignedRegistries(v []string) {
	o.UnsignedRegistries = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesTrustedRegistriesPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesTrustedRegistriesPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesTrustedRegistriesPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RootCaRegistries != nil {
		toSerialize["RootCaRegistries"] = o.RootCaRegistries
	}
	if o.UnsignedRegistries != nil {
		toSerialize["UnsignedRegistries"] = o.UnsignedRegistries
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesTrustedRegistriesPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesTrustedRegistriesPolicyAllOf := _KubernetesTrustedRegistriesPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesTrustedRegistriesPolicyAllOf); err == nil {
		*o = KubernetesTrustedRegistriesPolicyAllOf(varKubernetesTrustedRegistriesPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RootCaRegistries")
		delete(additionalProperties, "UnsignedRegistries")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesTrustedRegistriesPolicyAllOf struct {
	value *KubernetesTrustedRegistriesPolicyAllOf
	isSet bool
}

func (v NullableKubernetesTrustedRegistriesPolicyAllOf) Get() *KubernetesTrustedRegistriesPolicyAllOf {
	return v.value
}

func (v *NullableKubernetesTrustedRegistriesPolicyAllOf) Set(val *KubernetesTrustedRegistriesPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesTrustedRegistriesPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesTrustedRegistriesPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesTrustedRegistriesPolicyAllOf(val *KubernetesTrustedRegistriesPolicyAllOf) *NullableKubernetesTrustedRegistriesPolicyAllOf {
	return &NullableKubernetesTrustedRegistriesPolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesTrustedRegistriesPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesTrustedRegistriesPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
