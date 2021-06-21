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

// StorageNetAppFcInterfaceAllOf Definition of the list of properties defined in 'storage.NetAppFcInterface', excluding properties defined in parent classes.
type StorageNetAppFcInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// FC interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// State of the FC interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Uuid of  NetApp FC Interface.
	Uuid                 *string                             `json:"Uuid,omitempty"`
	ArrayController      *StorageNetAppNodeRelationship      `json:"ArrayController,omitempty"`
	PhysicalPort         *StorageNetAppFcPortRelationship    `json:"PhysicalPort,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcInterfaceAllOf StorageNetAppFcInterfaceAllOf

// NewStorageNetAppFcInterfaceAllOf instantiates a new StorageNetAppFcInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcInterfaceAllOf(classId string, objectType string) *StorageNetAppFcInterfaceAllOf {
	this := StorageNetAppFcInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	return &this
}

// NewStorageNetAppFcInterfaceAllOfWithDefaults instantiates a new StorageNetAppFcInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcInterfaceAllOfWithDefaults() *StorageNetAppFcInterfaceAllOf {
	this := StorageNetAppFcInterfaceAllOf{}
	var classId string = "storage.NetAppFcInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcInterface"
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppFcInterfaceAllOf) SetEnabled(v string) {
	o.Enabled = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppFcInterfaceAllOf) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppFcInterfaceAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppFcInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetPhysicalPort returns the PhysicalPort field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetPhysicalPort() StorageNetAppFcPortRelationship {
	if o == nil || o.PhysicalPort == nil {
		var ret StorageNetAppFcPortRelationship
		return ret
	}
	return *o.PhysicalPort
}

// GetPhysicalPortOk returns a tuple with the PhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetPhysicalPortOk() (*StorageNetAppFcPortRelationship, bool) {
	if o == nil || o.PhysicalPort == nil {
		return nil, false
	}
	return o.PhysicalPort, true
}

// HasPhysicalPort returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasPhysicalPort() bool {
	if o != nil && o.PhysicalPort != nil {
		return true
	}

	return false
}

// SetPhysicalPort gets a reference to the given StorageNetAppFcPortRelationship and assigns it to the PhysicalPort field.
func (o *StorageNetAppFcInterfaceAllOf) SetPhysicalPort(v StorageNetAppFcPortRelationship) {
	o.PhysicalPort = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppFcInterfaceAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcInterfaceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppFcInterfaceAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppFcInterfaceAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppFcInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.PhysicalPort != nil {
		toSerialize["PhysicalPort"] = o.PhysicalPort
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppFcInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppFcInterfaceAllOf := _StorageNetAppFcInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppFcInterfaceAllOf); err == nil {
		*o = StorageNetAppFcInterfaceAllOf(varStorageNetAppFcInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "PhysicalPort")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppFcInterfaceAllOf struct {
	value *StorageNetAppFcInterfaceAllOf
	isSet bool
}

func (v NullableStorageNetAppFcInterfaceAllOf) Get() *StorageNetAppFcInterfaceAllOf {
	return v.value
}

func (v *NullableStorageNetAppFcInterfaceAllOf) Set(val *StorageNetAppFcInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcInterfaceAllOf(val *StorageNetAppFcInterfaceAllOf) *NullableStorageNetAppFcInterfaceAllOf {
	return &NullableStorageNetAppFcInterfaceAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppFcInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
