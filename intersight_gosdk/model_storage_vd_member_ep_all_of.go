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

// StorageVdMemberEpAllOf Definition of the list of properties defined in 'storage.VdMemberEp', excluding properties defined in parent classes.
type StorageVdMemberEpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// For certain states, indicates the reason why the operState is in that state.
	OperQualifierReason *string `json:"OperQualifierReason,omitempty"`
	// The presence state of the local disk.
	Presence *string `json:"Presence,omitempty"`
	// Role of the disk normal or hot-spare, used by virtual-drive.
	Role *string `json:"Role,omitempty"`
	// The span id number of the virtual drive.
	SpanId *string `json:"SpanId,omitempty"`
	// The local disk slot number as id.
	VdMemberEpId         *int64                               `json:"VdMemberEpId,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageVirtualDrive  *StorageVirtualDriveRelationship     `json:"StorageVirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVdMemberEpAllOf StorageVdMemberEpAllOf

// NewStorageVdMemberEpAllOf instantiates a new StorageVdMemberEpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVdMemberEpAllOf(classId string, objectType string) *StorageVdMemberEpAllOf {
	this := StorageVdMemberEpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVdMemberEpAllOfWithDefaults instantiates a new StorageVdMemberEpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVdMemberEpAllOfWithDefaults() *StorageVdMemberEpAllOf {
	this := StorageVdMemberEpAllOf{}
	var classId string = "storage.VdMemberEp"
	this.ClassId = classId
	var objectType string = "storage.VdMemberEp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVdMemberEpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVdMemberEpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVdMemberEpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVdMemberEpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperQualifierReason returns the OperQualifierReason field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetOperQualifierReason() string {
	if o == nil || o.OperQualifierReason == nil {
		var ret string
		return ret
	}
	return *o.OperQualifierReason
}

// GetOperQualifierReasonOk returns a tuple with the OperQualifierReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetOperQualifierReasonOk() (*string, bool) {
	if o == nil || o.OperQualifierReason == nil {
		return nil, false
	}
	return o.OperQualifierReason, true
}

// HasOperQualifierReason returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasOperQualifierReason() bool {
	if o != nil && o.OperQualifierReason != nil {
		return true
	}

	return false
}

// SetOperQualifierReason gets a reference to the given string and assigns it to the OperQualifierReason field.
func (o *StorageVdMemberEpAllOf) SetOperQualifierReason(v string) {
	o.OperQualifierReason = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *StorageVdMemberEpAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *StorageVdMemberEpAllOf) SetRole(v string) {
	o.Role = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasSpanId() bool {
	if o != nil && o.SpanId != nil {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *StorageVdMemberEpAllOf) SetSpanId(v string) {
	o.SpanId = &v
}

// GetVdMemberEpId returns the VdMemberEpId field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetVdMemberEpId() int64 {
	if o == nil || o.VdMemberEpId == nil {
		var ret int64
		return ret
	}
	return *o.VdMemberEpId
}

// GetVdMemberEpIdOk returns a tuple with the VdMemberEpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetVdMemberEpIdOk() (*int64, bool) {
	if o == nil || o.VdMemberEpId == nil {
		return nil, false
	}
	return o.VdMemberEpId, true
}

// HasVdMemberEpId returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasVdMemberEpId() bool {
	if o != nil && o.VdMemberEpId != nil {
		return true
	}

	return false
}

// SetVdMemberEpId gets a reference to the given int64 and assigns it to the VdMemberEpId field.
func (o *StorageVdMemberEpAllOf) SetVdMemberEpId(v int64) {
	o.VdMemberEpId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVdMemberEpAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVdMemberEpAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageVirtualDrive returns the StorageVirtualDrive field value if set, zero value otherwise.
func (o *StorageVdMemberEpAllOf) GetStorageVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || o.StorageVirtualDrive == nil {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.StorageVirtualDrive
}

// GetStorageVirtualDriveOk returns a tuple with the StorageVirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVdMemberEpAllOf) GetStorageVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil || o.StorageVirtualDrive == nil {
		return nil, false
	}
	return o.StorageVirtualDrive, true
}

// HasStorageVirtualDrive returns a boolean if a field has been set.
func (o *StorageVdMemberEpAllOf) HasStorageVirtualDrive() bool {
	if o != nil && o.StorageVirtualDrive != nil {
		return true
	}

	return false
}

// SetStorageVirtualDrive gets a reference to the given StorageVirtualDriveRelationship and assigns it to the StorageVirtualDrive field.
func (o *StorageVdMemberEpAllOf) SetStorageVirtualDrive(v StorageVirtualDriveRelationship) {
	o.StorageVirtualDrive = &v
}

func (o StorageVdMemberEpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperQualifierReason != nil {
		toSerialize["OperQualifierReason"] = o.OperQualifierReason
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.SpanId != nil {
		toSerialize["SpanId"] = o.SpanId
	}
	if o.VdMemberEpId != nil {
		toSerialize["VdMemberEpId"] = o.VdMemberEpId
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageVirtualDrive != nil {
		toSerialize["StorageVirtualDrive"] = o.StorageVirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVdMemberEpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageVdMemberEpAllOf := _StorageVdMemberEpAllOf{}

	if err = json.Unmarshal(bytes, &varStorageVdMemberEpAllOf); err == nil {
		*o = StorageVdMemberEpAllOf(varStorageVdMemberEpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperQualifierReason")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "SpanId")
		delete(additionalProperties, "VdMemberEpId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageVirtualDrive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageVdMemberEpAllOf struct {
	value *StorageVdMemberEpAllOf
	isSet bool
}

func (v NullableStorageVdMemberEpAllOf) Get() *StorageVdMemberEpAllOf {
	return v.value
}

func (v *NullableStorageVdMemberEpAllOf) Set(val *StorageVdMemberEpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVdMemberEpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVdMemberEpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVdMemberEpAllOf(val *StorageVdMemberEpAllOf) *NullableStorageVdMemberEpAllOf {
	return &NullableStorageVdMemberEpAllOf{value: val, isSet: true}
}

func (v NullableStorageVdMemberEpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVdMemberEpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
