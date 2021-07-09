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

// VirtualizationBaseDatastore Common attributes of a datastore allocated to a hypervisor. Serves as a base class for all concrete datastore types. A datastore is assigned to a datacenter and virtual machines use it for storage. The datastore could be backed by NFS, VMFS, etc.
type VirtualizationBaseDatastore struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                                `json:"ObjectType"`
	Capacity   NullableVirtualizationStorageCapacity `json:"Capacity,omitempty"`
	// Number of hosts attached to or supported-by this datastore.
	HostCount *int64 `json:"HostCount,omitempty"`
	// The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.
	Name *string `json:"Name,omitempty"`
	// A string indicating the type of the datastore (VMFS, NFS, etc). * `Unknown` - The nature of the file system is unknown. * `VMFS` - It is a Virtual Machine Filesystem. * `NFS` - It is a Network File System. * `vSAN` - It is a virtual Storage Area Network file system. * `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server.
	Type *string `json:"Type,omitempty"`
	// Number of virtual machines relying on (using) this datastore.
	VmCount              *int64 `json:"VmCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseDatastore VirtualizationBaseDatastore

// NewVirtualizationBaseDatastore instantiates a new VirtualizationBaseDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseDatastore(classId string, objectType string) *VirtualizationBaseDatastore {
	this := VirtualizationBaseDatastore{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// NewVirtualizationBaseDatastoreWithDefaults instantiates a new VirtualizationBaseDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseDatastoreWithDefaults() *VirtualizationBaseDatastore {
	this := VirtualizationBaseDatastore{}
	var classId string = "virtualization.VmwareDatastore"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatastore"
	this.ObjectType = objectType
	var type_ string = "Unknown"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseDatastore) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseDatastore) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseDatastore) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseDatastore) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseDatastore) GetCapacity() VirtualizationStorageCapacity {
	if o == nil || o.Capacity.Get() == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.Capacity.Get()
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseDatastore) GetCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capacity.Get(), o.Capacity.IsSet()
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasCapacity() bool {
	if o != nil && o.Capacity.IsSet() {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given NullableVirtualizationStorageCapacity and assigns it to the Capacity field.
func (o *VirtualizationBaseDatastore) SetCapacity(v VirtualizationStorageCapacity) {
	o.Capacity.Set(&v)
}

// SetCapacityNil sets the value for Capacity to be an explicit nil
func (o *VirtualizationBaseDatastore) SetCapacityNil() {
	o.Capacity.Set(nil)
}

// UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
func (o *VirtualizationBaseDatastore) UnsetCapacity() {
	o.Capacity.Unset()
}

// GetHostCount returns the HostCount field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastore) GetHostCount() int64 {
	if o == nil || o.HostCount == nil {
		var ret int64
		return ret
	}
	return *o.HostCount
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetHostCountOk() (*int64, bool) {
	if o == nil || o.HostCount == nil {
		return nil, false
	}
	return o.HostCount, true
}

// HasHostCount returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasHostCount() bool {
	if o != nil && o.HostCount != nil {
		return true
	}

	return false
}

// SetHostCount gets a reference to the given int64 and assigns it to the HostCount field.
func (o *VirtualizationBaseDatastore) SetHostCount(v int64) {
	o.HostCount = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastore) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseDatastore) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastore) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseDatastore) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationBaseDatastore) SetType(v string) {
	o.Type = &v
}

// GetVmCount returns the VmCount field value if set, zero value otherwise.
func (o *VirtualizationBaseDatastore) GetVmCount() int64 {
	if o == nil || o.VmCount == nil {
		var ret int64
		return ret
	}
	return *o.VmCount
}

// GetVmCountOk returns a tuple with the VmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseDatastore) GetVmCountOk() (*int64, bool) {
	if o == nil || o.VmCount == nil {
		return nil, false
	}
	return o.VmCount, true
}

// HasVmCount returns a boolean if a field has been set.
func (o *VirtualizationBaseDatastore) HasVmCount() bool {
	if o != nil && o.VmCount != nil {
		return true
	}

	return false
}

// SetVmCount gets a reference to the given int64 and assigns it to the VmCount field.
func (o *VirtualizationBaseDatastore) SetVmCount(v int64) {
	o.VmCount = &v
}

func (o VirtualizationBaseDatastore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity.IsSet() {
		toSerialize["Capacity"] = o.Capacity.Get()
	}
	if o.HostCount != nil {
		toSerialize["HostCount"] = o.HostCount
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VmCount != nil {
		toSerialize["VmCount"] = o.VmCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseDatastore) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseDatastoreWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                                `json:"ObjectType"`
		Capacity   NullableVirtualizationStorageCapacity `json:"Capacity,omitempty"`
		// Number of hosts attached to or supported-by this datastore.
		HostCount *int64 `json:"HostCount,omitempty"`
		// The internally generated identity of this datastore. This entity is not manipulated by users. It aids in uniquely identifying the datastore object. For VMware, this is a MOR (managed object reference).
		Identity *string `json:"Identity,omitempty"`
		// Name of this datastore supplied by user. It is not the identity of the datastore. The name is subject to user manipulations.
		Name *string `json:"Name,omitempty"`
		// A string indicating the type of the datastore (VMFS, NFS, etc). * `Unknown` - The nature of the file system is unknown. * `VMFS` - It is a Virtual Machine Filesystem. * `NFS` - It is a Network File System. * `vSAN` - It is a virtual Storage Area Network file system. * `VirtualVolume` - A Virtual Volume datastore represents a storage container in a hypervisor server.
		Type *string `json:"Type,omitempty"`
		// Number of virtual machines relying on (using) this datastore.
		VmCount *int64 `json:"VmCount,omitempty"`
	}

	varVirtualizationBaseDatastoreWithoutEmbeddedStruct := VirtualizationBaseDatastoreWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseDatastoreWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseDatastore := _VirtualizationBaseDatastore{}
		varVirtualizationBaseDatastore.ClassId = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseDatastore.ObjectType = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseDatastore.Capacity = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.Capacity
		varVirtualizationBaseDatastore.HostCount = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.HostCount
		varVirtualizationBaseDatastore.Identity = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.Identity
		varVirtualizationBaseDatastore.Name = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.Name
		varVirtualizationBaseDatastore.Type = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.Type
		varVirtualizationBaseDatastore.VmCount = varVirtualizationBaseDatastoreWithoutEmbeddedStruct.VmCount
		*o = VirtualizationBaseDatastore(varVirtualizationBaseDatastore)
	} else {
		return err
	}

	varVirtualizationBaseDatastore := _VirtualizationBaseDatastore{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseDatastore)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseDatastore.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "HostCount")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VmCount")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseDatastore struct {
	value *VirtualizationBaseDatastore
	isSet bool
}

func (v NullableVirtualizationBaseDatastore) Get() *VirtualizationBaseDatastore {
	return v.value
}

func (v *NullableVirtualizationBaseDatastore) Set(val *VirtualizationBaseDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseDatastore(val *VirtualizationBaseDatastore) *NullableVirtualizationBaseDatastore {
	return &NullableVirtualizationBaseDatastore{value: val, isSet: true}
}

func (v NullableVirtualizationBaseDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
