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

// StorageLocalDiskAllOf Definition of the list of properties defined in 'storage.LocalDisk', excluding properties defined in parent classes.
type StorageLocalDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server.
	SlotNumber           *int64 `json:"SlotNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageLocalDiskAllOf StorageLocalDiskAllOf

// NewStorageLocalDiskAllOf instantiates a new StorageLocalDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageLocalDiskAllOf(classId string, objectType string) *StorageLocalDiskAllOf {
	this := StorageLocalDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageLocalDiskAllOfWithDefaults instantiates a new StorageLocalDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageLocalDiskAllOfWithDefaults() *StorageLocalDiskAllOf {
	this := StorageLocalDiskAllOf{}
	var classId string = "storage.LocalDisk"
	this.ClassId = classId
	var objectType string = "storage.LocalDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageLocalDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageLocalDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageLocalDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageLocalDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageLocalDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageLocalDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSlotNumber returns the SlotNumber field value if set, zero value otherwise.
func (o *StorageLocalDiskAllOf) GetSlotNumber() int64 {
	if o == nil || o.SlotNumber == nil {
		var ret int64
		return ret
	}
	return *o.SlotNumber
}

// GetSlotNumberOk returns a tuple with the SlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageLocalDiskAllOf) GetSlotNumberOk() (*int64, bool) {
	if o == nil || o.SlotNumber == nil {
		return nil, false
	}
	return o.SlotNumber, true
}

// HasSlotNumber returns a boolean if a field has been set.
func (o *StorageLocalDiskAllOf) HasSlotNumber() bool {
	if o != nil && o.SlotNumber != nil {
		return true
	}

	return false
}

// SetSlotNumber gets a reference to the given int64 and assigns it to the SlotNumber field.
func (o *StorageLocalDiskAllOf) SetSlotNumber(v int64) {
	o.SlotNumber = &v
}

func (o StorageLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SlotNumber != nil {
		toSerialize["SlotNumber"] = o.SlotNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageLocalDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageLocalDiskAllOf := _StorageLocalDiskAllOf{}

	if err = json.Unmarshal(bytes, &varStorageLocalDiskAllOf); err == nil {
		*o = StorageLocalDiskAllOf(varStorageLocalDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SlotNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageLocalDiskAllOf struct {
	value *StorageLocalDiskAllOf
	isSet bool
}

func (v NullableStorageLocalDiskAllOf) Get() *StorageLocalDiskAllOf {
	return v.value
}

func (v *NullableStorageLocalDiskAllOf) Set(val *StorageLocalDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageLocalDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageLocalDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageLocalDiskAllOf(val *StorageLocalDiskAllOf) *NullableStorageLocalDiskAllOf {
	return &NullableStorageLocalDiskAllOf{value: val, isSet: true}
}

func (v NullableStorageLocalDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageLocalDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
