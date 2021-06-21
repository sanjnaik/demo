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
	"time"
)

// HyperflexVolumeAllOf Definition of the list of properties defined in 'hyperflex.Volume', excluding properties defined in parent classes.
type HyperflexVolumeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Provisioned Capacity of the volume in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Client (tenant) ID to which the volume belongs.
	ClientId *string `json:"ClientId,omitempty"`
	// Last modified time as UTC of the volume.
	LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
	// UUID of LUN associated with the volume.
	LunUuid *string `json:"LunUuid,omitempty"`
	// Serial number of the volume.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The unique identifier for this volume.
	Uuid *string `json:"Uuid,omitempty"`
	// Access Mode of the volume. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. * `` - Unknown disk access mode.
	VolumeAccessMode *string `json:"VolumeAccessMode,omitempty"`
	// The mode of the HyperFlex volume. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	VolumeMode *string `json:"VolumeMode,omitempty"`
	// The type of the HyperFlex volume.
	VolumeType           *string                                `json:"VolumeType,omitempty"`
	Cluster              *HyperflexClusterRelationship          `json:"Cluster,omitempty"`
	StorageContainer     *HyperflexStorageContainerRelationship `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVolumeAllOf HyperflexVolumeAllOf

// NewHyperflexVolumeAllOf instantiates a new HyperflexVolumeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVolumeAllOf(classId string, objectType string) *HyperflexVolumeAllOf {
	this := HyperflexVolumeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var volumeAccessMode string = "ReadWriteOnce"
	this.VolumeAccessMode = &volumeAccessMode
	var volumeMode string = "Block"
	this.VolumeMode = &volumeMode
	return &this
}

// NewHyperflexVolumeAllOfWithDefaults instantiates a new HyperflexVolumeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVolumeAllOfWithDefaults() *HyperflexVolumeAllOf {
	this := HyperflexVolumeAllOf{}
	var classId string = "hyperflex.Volume"
	this.ClassId = classId
	var objectType string = "hyperflex.Volume"
	this.ObjectType = objectType
	var volumeAccessMode string = "ReadWriteOnce"
	this.VolumeAccessMode = &volumeAccessMode
	var volumeMode string = "Block"
	this.VolumeMode = &volumeMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVolumeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVolumeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVolumeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVolumeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *HyperflexVolumeAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *HyperflexVolumeAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *HyperflexVolumeAllOf) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetLunUuid returns the LunUuid field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetLunUuid() string {
	if o == nil || o.LunUuid == nil {
		var ret string
		return ret
	}
	return *o.LunUuid
}

// GetLunUuidOk returns a tuple with the LunUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetLunUuidOk() (*string, bool) {
	if o == nil || o.LunUuid == nil {
		return nil, false
	}
	return o.LunUuid, true
}

// HasLunUuid returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasLunUuid() bool {
	if o != nil && o.LunUuid != nil {
		return true
	}

	return false
}

// SetLunUuid gets a reference to the given string and assigns it to the LunUuid field.
func (o *HyperflexVolumeAllOf) SetLunUuid(v string) {
	o.LunUuid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexVolumeAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexVolumeAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeAccessMode returns the VolumeAccessMode field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetVolumeAccessMode() string {
	if o == nil || o.VolumeAccessMode == nil {
		var ret string
		return ret
	}
	return *o.VolumeAccessMode
}

// GetVolumeAccessModeOk returns a tuple with the VolumeAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetVolumeAccessModeOk() (*string, bool) {
	if o == nil || o.VolumeAccessMode == nil {
		return nil, false
	}
	return o.VolumeAccessMode, true
}

// HasVolumeAccessMode returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasVolumeAccessMode() bool {
	if o != nil && o.VolumeAccessMode != nil {
		return true
	}

	return false
}

// SetVolumeAccessMode gets a reference to the given string and assigns it to the VolumeAccessMode field.
func (o *HyperflexVolumeAllOf) SetVolumeAccessMode(v string) {
	o.VolumeAccessMode = &v
}

// GetVolumeMode returns the VolumeMode field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetVolumeMode() string {
	if o == nil || o.VolumeMode == nil {
		var ret string
		return ret
	}
	return *o.VolumeMode
}

// GetVolumeModeOk returns a tuple with the VolumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetVolumeModeOk() (*string, bool) {
	if o == nil || o.VolumeMode == nil {
		return nil, false
	}
	return o.VolumeMode, true
}

// HasVolumeMode returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasVolumeMode() bool {
	if o != nil && o.VolumeMode != nil {
		return true
	}

	return false
}

// SetVolumeMode gets a reference to the given string and assigns it to the VolumeMode field.
func (o *HyperflexVolumeAllOf) SetVolumeMode(v string) {
	o.VolumeMode = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *HyperflexVolumeAllOf) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexVolumeAllOf) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *HyperflexVolumeAllOf) GetStorageContainer() HyperflexStorageContainerRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret HyperflexStorageContainerRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolumeAllOf) GetStorageContainerOk() (*HyperflexStorageContainerRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *HyperflexVolumeAllOf) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given HyperflexStorageContainerRelationship and assigns it to the StorageContainer field.
func (o *HyperflexVolumeAllOf) SetStorageContainer(v HyperflexStorageContainerRelationship) {
	o.StorageContainer = &v
}

func (o HyperflexVolumeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.LastModifiedTime != nil {
		toSerialize["LastModifiedTime"] = o.LastModifiedTime
	}
	if o.LunUuid != nil {
		toSerialize["LunUuid"] = o.LunUuid
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeAccessMode != nil {
		toSerialize["VolumeAccessMode"] = o.VolumeAccessMode
	}
	if o.VolumeMode != nil {
		toSerialize["VolumeMode"] = o.VolumeMode
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVolumeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVolumeAllOf := _HyperflexVolumeAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVolumeAllOf); err == nil {
		*o = HyperflexVolumeAllOf(varHyperflexVolumeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "LastModifiedTime")
		delete(additionalProperties, "LunUuid")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeAccessMode")
		delete(additionalProperties, "VolumeMode")
		delete(additionalProperties, "VolumeType")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "StorageContainer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVolumeAllOf struct {
	value *HyperflexVolumeAllOf
	isSet bool
}

func (v NullableHyperflexVolumeAllOf) Get() *HyperflexVolumeAllOf {
	return v.value
}

func (v *NullableHyperflexVolumeAllOf) Set(val *HyperflexVolumeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVolumeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVolumeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVolumeAllOf(val *HyperflexVolumeAllOf) *NullableHyperflexVolumeAllOf {
	return &NullableHyperflexVolumeAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVolumeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVolumeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
