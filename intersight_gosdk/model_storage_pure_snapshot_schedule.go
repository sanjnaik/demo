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
	"reflect"
	"strings"
)

// StoragePureSnapshotSchedule PureStorage FlashArray snapshot schedule configuration entity.
type StoragePureSnapshotSchedule struct {
	StorageBaseSnapshotSchedule
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	DailyLimit *int64 `json:"DailyLimit,omitempty"`
	// Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.
	SnapshotExpiryTime   *string                                 `json:"SnapshotExpiryTime,omitempty"`
	Array                *StoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureSnapshotSchedule StoragePureSnapshotSchedule

// NewStoragePureSnapshotSchedule instantiates a new StoragePureSnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureSnapshotSchedule(classId string, objectType string) *StoragePureSnapshotSchedule {
	this := StoragePureSnapshotSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureSnapshotScheduleWithDefaults instantiates a new StoragePureSnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureSnapshotScheduleWithDefaults() *StoragePureSnapshotSchedule {
	this := StoragePureSnapshotSchedule{}
	var classId string = "storage.PureSnapshotSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureSnapshotSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureSnapshotSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureSnapshotSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureSnapshotSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureSnapshotSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetDailyLimit() int64 {
	if o == nil || o.DailyLimit == nil {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetDailyLimitOk() (*int64, bool) {
	if o == nil || o.DailyLimit == nil {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasDailyLimit() bool {
	if o != nil && o.DailyLimit != nil {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *StoragePureSnapshotSchedule) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetSnapshotExpiryTime returns the SnapshotExpiryTime field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTime() string {
	if o == nil || o.SnapshotExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.SnapshotExpiryTime
}

// GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetSnapshotExpiryTimeOk() (*string, bool) {
	if o == nil || o.SnapshotExpiryTime == nil {
		return nil, false
	}
	return o.SnapshotExpiryTime, true
}

// HasSnapshotExpiryTime returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasSnapshotExpiryTime() bool {
	if o != nil && o.SnapshotExpiryTime != nil {
		return true
	}

	return false
}

// SetSnapshotExpiryTime gets a reference to the given string and assigns it to the SnapshotExpiryTime field.
func (o *StoragePureSnapshotSchedule) SetSnapshotExpiryTime(v string) {
	o.SnapshotExpiryTime = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureSnapshotSchedule) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureSnapshotSchedule) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureSnapshotSchedule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureSnapshotSchedule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureSnapshotSchedule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureSnapshotSchedule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureSnapshotSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseSnapshotSchedule, errStorageBaseSnapshotSchedule := json.Marshal(o.StorageBaseSnapshotSchedule)
	if errStorageBaseSnapshotSchedule != nil {
		return []byte{}, errStorageBaseSnapshotSchedule
	}
	errStorageBaseSnapshotSchedule = json.Unmarshal([]byte(serializedStorageBaseSnapshotSchedule), &toSerialize)
	if errStorageBaseSnapshotSchedule != nil {
		return []byte{}, errStorageBaseSnapshotSchedule
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DailyLimit != nil {
		toSerialize["DailyLimit"] = o.DailyLimit
	}
	if o.SnapshotExpiryTime != nil {
		toSerialize["SnapshotExpiryTime"] = o.SnapshotExpiryTime
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureSnapshotSchedule) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureSnapshotScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total number of snapshots per day to be available on source above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
		DailyLimit *int64 `json:"DailyLimit,omitempty"`
		// Duration to keep the daily limit snapshots on source array. StorageArray deletes the snapshots permanently from source beyond this period.
		SnapshotExpiryTime *string                                 `json:"SnapshotExpiryTime,omitempty"`
		Array              *StoragePureArrayRelationship           `json:"Array,omitempty"`
		ProtectionGroup    *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice   *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureSnapshotScheduleWithoutEmbeddedStruct := StoragePureSnapshotScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureSnapshotScheduleWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureSnapshotSchedule := _StoragePureSnapshotSchedule{}
		varStoragePureSnapshotSchedule.ClassId = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ClassId
		varStoragePureSnapshotSchedule.ObjectType = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ObjectType
		varStoragePureSnapshotSchedule.DailyLimit = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.DailyLimit
		varStoragePureSnapshotSchedule.SnapshotExpiryTime = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.SnapshotExpiryTime
		varStoragePureSnapshotSchedule.Array = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.Array
		varStoragePureSnapshotSchedule.ProtectionGroup = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureSnapshotSchedule.RegisteredDevice = varStoragePureSnapshotScheduleWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureSnapshotSchedule(varStoragePureSnapshotSchedule)
	} else {
		return err
	}

	varStoragePureSnapshotSchedule := _StoragePureSnapshotSchedule{}

	err = json.Unmarshal(bytes, &varStoragePureSnapshotSchedule)
	if err == nil {
		o.StorageBaseSnapshotSchedule = varStoragePureSnapshotSchedule.StorageBaseSnapshotSchedule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DailyLimit")
		delete(additionalProperties, "SnapshotExpiryTime")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseSnapshotSchedule := reflect.ValueOf(o.StorageBaseSnapshotSchedule)
		for i := 0; i < reflectStorageBaseSnapshotSchedule.Type().NumField(); i++ {
			t := reflectStorageBaseSnapshotSchedule.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePureSnapshotSchedule struct {
	value *StoragePureSnapshotSchedule
	isSet bool
}

func (v NullableStoragePureSnapshotSchedule) Get() *StoragePureSnapshotSchedule {
	return v.value
}

func (v *NullableStoragePureSnapshotSchedule) Set(val *StoragePureSnapshotSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureSnapshotSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureSnapshotSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureSnapshotSchedule(val *StoragePureSnapshotSchedule) *NullableStoragePureSnapshotSchedule {
	return &NullableStoragePureSnapshotSchedule{value: val, isSet: true}
}

func (v NullableStoragePureSnapshotSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureSnapshotSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
