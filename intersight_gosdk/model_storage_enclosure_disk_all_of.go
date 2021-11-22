/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4903
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageEnclosureDiskAllOf Definition of the list of properties defined in 'storage.EnclosureDisk', excluding properties defined in parent classes.
type StorageEnclosureDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The block size of the physical disk in bytes.
	BlockSize *string `json:"BlockSize,omitempty"`
	// This field represents the disk Id in the storage enclosure.
	DiskId *string `json:"DiskId,omitempty"`
	// This field identifies the current disk configuration applied in the physical disk.
	DiskState *string `json:"DiskState,omitempty"`
	// The current health state of the enclosure disk.
	Health *string `json:"Health,omitempty"`
	// The number of blocks present on the physical disk.
	NumBlocks *string `json:"NumBlocks,omitempty"`
	// This field identifies the Product ID for physicalDisk.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the SAS address assigned to the disk SAS port-1.
	SasAddress1 *string `json:"SasAddress1,omitempty"`
	// This field identifies the SAS address assigned to the disk SAS port-2.
	SasAddress2 *string `json:"SasAddress2,omitempty"`
	// The size of the physical disk in MB.
	Size                 *string                              `json:"Size,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDisk         *StoragePhysicalDiskRelationship     `json:"PhysicalDisk,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageEnclosure     *StorageEnclosureRelationship        `json:"StorageEnclosure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageEnclosureDiskAllOf StorageEnclosureDiskAllOf

// NewStorageEnclosureDiskAllOf instantiates a new StorageEnclosureDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageEnclosureDiskAllOf(classId string, objectType string) *StorageEnclosureDiskAllOf {
	this := StorageEnclosureDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageEnclosureDiskAllOfWithDefaults instantiates a new StorageEnclosureDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageEnclosureDiskAllOfWithDefaults() *StorageEnclosureDiskAllOf {
	this := StorageEnclosureDiskAllOf{}
	var classId string = "storage.EnclosureDisk"
	this.ClassId = classId
	var objectType string = "storage.EnclosureDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageEnclosureDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageEnclosureDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageEnclosureDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageEnclosureDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetBlockSize() string {
	if o == nil || o.BlockSize == nil {
		var ret string
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetBlockSizeOk() (*string, bool) {
	if o == nil || o.BlockSize == nil {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasBlockSize() bool {
	if o != nil && o.BlockSize != nil {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given string and assigns it to the BlockSize field.
func (o *StorageEnclosureDiskAllOf) SetBlockSize(v string) {
	o.BlockSize = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetDiskId() string {
	if o == nil || o.DiskId == nil {
		var ret string
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetDiskIdOk() (*string, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given string and assigns it to the DiskId field.
func (o *StorageEnclosureDiskAllOf) SetDiskId(v string) {
	o.DiskId = &v
}

// GetDiskState returns the DiskState field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetDiskState() string {
	if o == nil || o.DiskState == nil {
		var ret string
		return ret
	}
	return *o.DiskState
}

// GetDiskStateOk returns a tuple with the DiskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetDiskStateOk() (*string, bool) {
	if o == nil || o.DiskState == nil {
		return nil, false
	}
	return o.DiskState, true
}

// HasDiskState returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasDiskState() bool {
	if o != nil && o.DiskState != nil {
		return true
	}

	return false
}

// SetDiskState gets a reference to the given string and assigns it to the DiskState field.
func (o *StorageEnclosureDiskAllOf) SetDiskState(v string) {
	o.DiskState = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *StorageEnclosureDiskAllOf) SetHealth(v string) {
	o.Health = &v
}

// GetNumBlocks returns the NumBlocks field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetNumBlocks() string {
	if o == nil || o.NumBlocks == nil {
		var ret string
		return ret
	}
	return *o.NumBlocks
}

// GetNumBlocksOk returns a tuple with the NumBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetNumBlocksOk() (*string, bool) {
	if o == nil || o.NumBlocks == nil {
		return nil, false
	}
	return o.NumBlocks, true
}

// HasNumBlocks returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasNumBlocks() bool {
	if o != nil && o.NumBlocks != nil {
		return true
	}

	return false
}

// SetNumBlocks gets a reference to the given string and assigns it to the NumBlocks field.
func (o *StorageEnclosureDiskAllOf) SetNumBlocks(v string) {
	o.NumBlocks = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *StorageEnclosureDiskAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSasAddress1 returns the SasAddress1 field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetSasAddress1() string {
	if o == nil || o.SasAddress1 == nil {
		var ret string
		return ret
	}
	return *o.SasAddress1
}

// GetSasAddress1Ok returns a tuple with the SasAddress1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetSasAddress1Ok() (*string, bool) {
	if o == nil || o.SasAddress1 == nil {
		return nil, false
	}
	return o.SasAddress1, true
}

// HasSasAddress1 returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasSasAddress1() bool {
	if o != nil && o.SasAddress1 != nil {
		return true
	}

	return false
}

// SetSasAddress1 gets a reference to the given string and assigns it to the SasAddress1 field.
func (o *StorageEnclosureDiskAllOf) SetSasAddress1(v string) {
	o.SasAddress1 = &v
}

// GetSasAddress2 returns the SasAddress2 field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetSasAddress2() string {
	if o == nil || o.SasAddress2 == nil {
		var ret string
		return ret
	}
	return *o.SasAddress2
}

// GetSasAddress2Ok returns a tuple with the SasAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetSasAddress2Ok() (*string, bool) {
	if o == nil || o.SasAddress2 == nil {
		return nil, false
	}
	return o.SasAddress2, true
}

// HasSasAddress2 returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasSasAddress2() bool {
	if o != nil && o.SasAddress2 != nil {
		return true
	}

	return false
}

// SetSasAddress2 gets a reference to the given string and assigns it to the SasAddress2 field.
func (o *StorageEnclosureDiskAllOf) SetSasAddress2(v string) {
	o.SasAddress2 = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageEnclosureDiskAllOf) SetSize(v string) {
	o.Size = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageEnclosureDiskAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDisk returns the PhysicalDisk field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetPhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.PhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.PhysicalDisk
}

// GetPhysicalDiskOk returns a tuple with the PhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetPhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisk == nil {
		return nil, false
	}
	return o.PhysicalDisk, true
}

// HasPhysicalDisk returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasPhysicalDisk() bool {
	if o != nil && o.PhysicalDisk != nil {
		return true
	}

	return false
}

// SetPhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the PhysicalDisk field.
func (o *StorageEnclosureDiskAllOf) SetPhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.PhysicalDisk = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageEnclosureDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageEnclosure returns the StorageEnclosure field value if set, zero value otherwise.
func (o *StorageEnclosureDiskAllOf) GetStorageEnclosure() StorageEnclosureRelationship {
	if o == nil || o.StorageEnclosure == nil {
		var ret StorageEnclosureRelationship
		return ret
	}
	return *o.StorageEnclosure
}

// GetStorageEnclosureOk returns a tuple with the StorageEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageEnclosureDiskAllOf) GetStorageEnclosureOk() (*StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosure == nil {
		return nil, false
	}
	return o.StorageEnclosure, true
}

// HasStorageEnclosure returns a boolean if a field has been set.
func (o *StorageEnclosureDiskAllOf) HasStorageEnclosure() bool {
	if o != nil && o.StorageEnclosure != nil {
		return true
	}

	return false
}

// SetStorageEnclosure gets a reference to the given StorageEnclosureRelationship and assigns it to the StorageEnclosure field.
func (o *StorageEnclosureDiskAllOf) SetStorageEnclosure(v StorageEnclosureRelationship) {
	o.StorageEnclosure = &v
}

func (o StorageEnclosureDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BlockSize != nil {
		toSerialize["BlockSize"] = o.BlockSize
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.DiskState != nil {
		toSerialize["DiskState"] = o.DiskState
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.NumBlocks != nil {
		toSerialize["NumBlocks"] = o.NumBlocks
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SasAddress1 != nil {
		toSerialize["SasAddress1"] = o.SasAddress1
	}
	if o.SasAddress2 != nil {
		toSerialize["SasAddress2"] = o.SasAddress2
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDisk != nil {
		toSerialize["PhysicalDisk"] = o.PhysicalDisk
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageEnclosure != nil {
		toSerialize["StorageEnclosure"] = o.StorageEnclosure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageEnclosureDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageEnclosureDiskAllOf := _StorageEnclosureDiskAllOf{}

	if err = json.Unmarshal(bytes, &varStorageEnclosureDiskAllOf); err == nil {
		*o = StorageEnclosureDiskAllOf(varStorageEnclosureDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockSize")
		delete(additionalProperties, "DiskId")
		delete(additionalProperties, "DiskState")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "NumBlocks")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SasAddress1")
		delete(additionalProperties, "SasAddress2")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDisk")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageEnclosure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageEnclosureDiskAllOf struct {
	value *StorageEnclosureDiskAllOf
	isSet bool
}

func (v NullableStorageEnclosureDiskAllOf) Get() *StorageEnclosureDiskAllOf {
	return v.value
}

func (v *NullableStorageEnclosureDiskAllOf) Set(val *StorageEnclosureDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageEnclosureDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageEnclosureDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageEnclosureDiskAllOf(val *StorageEnclosureDiskAllOf) *NullableStorageEnclosureDiskAllOf {
	return &NullableStorageEnclosureDiskAllOf{value: val, isSet: true}
}

func (v NullableStorageEnclosureDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageEnclosureDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
