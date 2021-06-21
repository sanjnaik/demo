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

// FabricVsanAllOf Definition of the list of properties defined in 'fabric.Vsan', excluding properties defined in parent classes.
type FabricVsanAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables or Disables the default zoning state. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	DefaultZoning *string `json:"DefaultZoning,omitempty"`
	// Logical grouping mode for fc ports.
	FcZoneSharingMode *string `json:"FcZoneSharingMode,omitempty"`
	// FCOE Vlan associated to the VSAN configuration.
	FcoeVlan *int64 `json:"FcoeVlan,omitempty"`
	// User given name for the VSAN configuration.
	Name *string `json:"Name,omitempty"`
	// Virtual San Identifier in the switch.
	VsanId               *int64                             `json:"VsanId,omitempty"`
	FcNetworkPolicy      *FabricFcNetworkPolicyRelationship `json:"FcNetworkPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVsanAllOf FabricVsanAllOf

// NewFabricVsanAllOf instantiates a new FabricVsanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVsanAllOf(classId string, objectType string) *FabricVsanAllOf {
	this := FabricVsanAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultZoning string = "Enabled"
	this.DefaultZoning = &defaultZoning
	return &this
}

// NewFabricVsanAllOfWithDefaults instantiates a new FabricVsanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVsanAllOfWithDefaults() *FabricVsanAllOf {
	this := FabricVsanAllOf{}
	var classId string = "fabric.Vsan"
	this.ClassId = classId
	var objectType string = "fabric.Vsan"
	this.ObjectType = objectType
	var defaultZoning string = "Enabled"
	this.DefaultZoning = &defaultZoning
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVsanAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVsanAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVsanAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVsanAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultZoning returns the DefaultZoning field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetDefaultZoning() string {
	if o == nil || o.DefaultZoning == nil {
		var ret string
		return ret
	}
	return *o.DefaultZoning
}

// GetDefaultZoningOk returns a tuple with the DefaultZoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetDefaultZoningOk() (*string, bool) {
	if o == nil || o.DefaultZoning == nil {
		return nil, false
	}
	return o.DefaultZoning, true
}

// HasDefaultZoning returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasDefaultZoning() bool {
	if o != nil && o.DefaultZoning != nil {
		return true
	}

	return false
}

// SetDefaultZoning gets a reference to the given string and assigns it to the DefaultZoning field.
func (o *FabricVsanAllOf) SetDefaultZoning(v string) {
	o.DefaultZoning = &v
}

// GetFcZoneSharingMode returns the FcZoneSharingMode field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetFcZoneSharingMode() string {
	if o == nil || o.FcZoneSharingMode == nil {
		var ret string
		return ret
	}
	return *o.FcZoneSharingMode
}

// GetFcZoneSharingModeOk returns a tuple with the FcZoneSharingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetFcZoneSharingModeOk() (*string, bool) {
	if o == nil || o.FcZoneSharingMode == nil {
		return nil, false
	}
	return o.FcZoneSharingMode, true
}

// HasFcZoneSharingMode returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasFcZoneSharingMode() bool {
	if o != nil && o.FcZoneSharingMode != nil {
		return true
	}

	return false
}

// SetFcZoneSharingMode gets a reference to the given string and assigns it to the FcZoneSharingMode field.
func (o *FabricVsanAllOf) SetFcZoneSharingMode(v string) {
	o.FcZoneSharingMode = &v
}

// GetFcoeVlan returns the FcoeVlan field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetFcoeVlan() int64 {
	if o == nil || o.FcoeVlan == nil {
		var ret int64
		return ret
	}
	return *o.FcoeVlan
}

// GetFcoeVlanOk returns a tuple with the FcoeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetFcoeVlanOk() (*int64, bool) {
	if o == nil || o.FcoeVlan == nil {
		return nil, false
	}
	return o.FcoeVlan, true
}

// HasFcoeVlan returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasFcoeVlan() bool {
	if o != nil && o.FcoeVlan != nil {
		return true
	}

	return false
}

// SetFcoeVlan gets a reference to the given int64 and assigns it to the FcoeVlan field.
func (o *FabricVsanAllOf) SetFcoeVlan(v int64) {
	o.FcoeVlan = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVsanAllOf) SetName(v string) {
	o.Name = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricVsanAllOf) SetVsanId(v int64) {
	o.VsanId = &v
}

// GetFcNetworkPolicy returns the FcNetworkPolicy field value if set, zero value otherwise.
func (o *FabricVsanAllOf) GetFcNetworkPolicy() FabricFcNetworkPolicyRelationship {
	if o == nil || o.FcNetworkPolicy == nil {
		var ret FabricFcNetworkPolicyRelationship
		return ret
	}
	return *o.FcNetworkPolicy
}

// GetFcNetworkPolicyOk returns a tuple with the FcNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanAllOf) GetFcNetworkPolicyOk() (*FabricFcNetworkPolicyRelationship, bool) {
	if o == nil || o.FcNetworkPolicy == nil {
		return nil, false
	}
	return o.FcNetworkPolicy, true
}

// HasFcNetworkPolicy returns a boolean if a field has been set.
func (o *FabricVsanAllOf) HasFcNetworkPolicy() bool {
	if o != nil && o.FcNetworkPolicy != nil {
		return true
	}

	return false
}

// SetFcNetworkPolicy gets a reference to the given FabricFcNetworkPolicyRelationship and assigns it to the FcNetworkPolicy field.
func (o *FabricVsanAllOf) SetFcNetworkPolicy(v FabricFcNetworkPolicyRelationship) {
	o.FcNetworkPolicy = &v
}

func (o FabricVsanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DefaultZoning != nil {
		toSerialize["DefaultZoning"] = o.DefaultZoning
	}
	if o.FcZoneSharingMode != nil {
		toSerialize["FcZoneSharingMode"] = o.FcZoneSharingMode
	}
	if o.FcoeVlan != nil {
		toSerialize["FcoeVlan"] = o.FcoeVlan
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}
	if o.FcNetworkPolicy != nil {
		toSerialize["FcNetworkPolicy"] = o.FcNetworkPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVsanAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricVsanAllOf := _FabricVsanAllOf{}

	if err = json.Unmarshal(bytes, &varFabricVsanAllOf); err == nil {
		*o = FabricVsanAllOf(varFabricVsanAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultZoning")
		delete(additionalProperties, "FcZoneSharingMode")
		delete(additionalProperties, "FcoeVlan")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VsanId")
		delete(additionalProperties, "FcNetworkPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricVsanAllOf struct {
	value *FabricVsanAllOf
	isSet bool
}

func (v NullableFabricVsanAllOf) Get() *FabricVsanAllOf {
	return v.value
}

func (v *NullableFabricVsanAllOf) Set(val *FabricVsanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVsanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVsanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVsanAllOf(val *FabricVsanAllOf) *NullableFabricVsanAllOf {
	return &NullableFabricVsanAllOf{value: val, isSet: true}
}

func (v NullableFabricVsanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVsanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
