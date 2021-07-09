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

// InventoryUemInfoAllOf Definition of the list of properties defined in 'inventory.UemInfo', excluding properties defined in parent classes.
type InventoryUemInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Connections status of Uem endpoint. * `Unknown` - Connection status of UEM is not known. * `Connected` - UEM endpoint has an open TCP connection with device connector. * `Enabled` - Eventing is enabled on the UEM endpoint. * `Paused` - Eventing is paused on the UEM endpoint. * `Disabled` - Eventing is disabled on the UEM endpoint. * `Disconnected` - There is no TCP connection between the UEM endpoint and device connector.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Last Sequence ID in the UEM event frame received from the endpoint identified by endpointInfo.
	LastSequenceId *int64 `json:"LastSequenceId,omitempty"`
	// Version of UEM package on the endpoint.
	PackageVersion *string `json:"PackageVersion,omitempty"`
	// Version of UEM protocol running on the endpoint.
	ProtocolVersion *string `json:"ProtocolVersion,omitempty"`
	// The switch ID to identify the path to the endpoint.
	SwitchId             *string `json:"SwitchId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryUemInfoAllOf InventoryUemInfoAllOf

// NewInventoryUemInfoAllOf instantiates a new InventoryUemInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUemInfoAllOf(classId string, objectType string) *InventoryUemInfoAllOf {
	this := InventoryUemInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewInventoryUemInfoAllOfWithDefaults instantiates a new InventoryUemInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUemInfoAllOfWithDefaults() *InventoryUemInfoAllOf {
	this := InventoryUemInfoAllOf{}
	var classId string = "inventory.UemInfo"
	this.ClassId = classId
	var objectType string = "inventory.UemInfo"
	this.ObjectType = objectType
	var connectionStatus string = "Unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryUemInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryUemInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryUemInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryUemInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *InventoryUemInfoAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *InventoryUemInfoAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *InventoryUemInfoAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetLastSequenceId returns the LastSequenceId field value if set, zero value otherwise.
func (o *InventoryUemInfoAllOf) GetLastSequenceId() int64 {
	if o == nil || o.LastSequenceId == nil {
		var ret int64
		return ret
	}
	return *o.LastSequenceId
}

// GetLastSequenceIdOk returns a tuple with the LastSequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetLastSequenceIdOk() (*int64, bool) {
	if o == nil || o.LastSequenceId == nil {
		return nil, false
	}
	return o.LastSequenceId, true
}

// HasLastSequenceId returns a boolean if a field has been set.
func (o *InventoryUemInfoAllOf) HasLastSequenceId() bool {
	if o != nil && o.LastSequenceId != nil {
		return true
	}

	return false
}

// SetLastSequenceId gets a reference to the given int64 and assigns it to the LastSequenceId field.
func (o *InventoryUemInfoAllOf) SetLastSequenceId(v int64) {
	o.LastSequenceId = &v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise.
func (o *InventoryUemInfoAllOf) GetPackageVersion() string {
	if o == nil || o.PackageVersion == nil {
		var ret string
		return ret
	}
	return *o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetPackageVersionOk() (*string, bool) {
	if o == nil || o.PackageVersion == nil {
		return nil, false
	}
	return o.PackageVersion, true
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *InventoryUemInfoAllOf) HasPackageVersion() bool {
	if o != nil && o.PackageVersion != nil {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given string and assigns it to the PackageVersion field.
func (o *InventoryUemInfoAllOf) SetPackageVersion(v string) {
	o.PackageVersion = &v
}

// GetProtocolVersion returns the ProtocolVersion field value if set, zero value otherwise.
func (o *InventoryUemInfoAllOf) GetProtocolVersion() string {
	if o == nil || o.ProtocolVersion == nil {
		var ret string
		return ret
	}
	return *o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetProtocolVersionOk() (*string, bool) {
	if o == nil || o.ProtocolVersion == nil {
		return nil, false
	}
	return o.ProtocolVersion, true
}

// HasProtocolVersion returns a boolean if a field has been set.
func (o *InventoryUemInfoAllOf) HasProtocolVersion() bool {
	if o != nil && o.ProtocolVersion != nil {
		return true
	}

	return false
}

// SetProtocolVersion gets a reference to the given string and assigns it to the ProtocolVersion field.
func (o *InventoryUemInfoAllOf) SetProtocolVersion(v string) {
	o.ProtocolVersion = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *InventoryUemInfoAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUemInfoAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *InventoryUemInfoAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *InventoryUemInfoAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

func (o InventoryUemInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.LastSequenceId != nil {
		toSerialize["LastSequenceId"] = o.LastSequenceId
	}
	if o.PackageVersion != nil {
		toSerialize["PackageVersion"] = o.PackageVersion
	}
	if o.ProtocolVersion != nil {
		toSerialize["ProtocolVersion"] = o.ProtocolVersion
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryUemInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryUemInfoAllOf := _InventoryUemInfoAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryUemInfoAllOf); err == nil {
		*o = InventoryUemInfoAllOf(varInventoryUemInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "LastSequenceId")
		delete(additionalProperties, "PackageVersion")
		delete(additionalProperties, "ProtocolVersion")
		delete(additionalProperties, "SwitchId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryUemInfoAllOf struct {
	value *InventoryUemInfoAllOf
	isSet bool
}

func (v NullableInventoryUemInfoAllOf) Get() *InventoryUemInfoAllOf {
	return v.value
}

func (v *NullableInventoryUemInfoAllOf) Set(val *InventoryUemInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUemInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUemInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUemInfoAllOf(val *InventoryUemInfoAllOf) *NullableInventoryUemInfoAllOf {
	return &NullableInventoryUemInfoAllOf{value: val, isSet: true}
}

func (v NullableInventoryUemInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUemInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
