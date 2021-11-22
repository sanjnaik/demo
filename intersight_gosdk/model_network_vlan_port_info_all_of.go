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

// NetworkVlanPortInfoAllOf Definition of the list of properties defined in 'network.VlanPortInfo', excluding properties defined in parent classes.
type NetworkVlanPortInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of available VLAN access ports on a Fabric Interconnect.
	AccessVlanPortCount *int64 `json:"AccessVlanPortCount,omitempty"`
	// The number of available VLAN border ports on a Fabric Interconnect.
	BorderVlanPortCount *int64 `json:"BorderVlanPortCount,omitempty"`
	// The number of compressed VLAN Group count on a Fabric Interconnect calculated by VLAN port group library.
	CompressedOptimizationSetsValue *int64 `json:"CompressedOptimizationSetsValue,omitempty"`
	// The number of compressed VLAN ports on a Fabric Interconnect.
	// Deprecated
	CompressedVlanPortCount *string `json:"CompressedVlanPortCount,omitempty"`
	// The number of compressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
	CompressedVlanPortCountValue *int64 `json:"CompressedVlanPortCountValue,omitempty"`
	// The total number of VLAN ports on a Fabric Interconnect.
	TotalVlanPortCount *int64 `json:"TotalVlanPortCount,omitempty"`
	// The number of uncompressed VLAN ports on a Fabric Interconnect.
	// Deprecated
	UncompressedVlanPortCount *string `json:"UncompressedVlanPortCount,omitempty"`
	// The number of uncompressed VLAN port count on a Fabric Interconnect calculated by VLAN port group library.
	UncompressedVlanPortCountValue *int64 `json:"UncompressedVlanPortCountValue,omitempty"`
	// The maximum number of VLAN ports allowed on a Fabric Interconnect.
	VlanPortLimit        *int64                               `json:"VlanPortLimit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkVlanPortInfoAllOf NetworkVlanPortInfoAllOf

// NewNetworkVlanPortInfoAllOf instantiates a new NetworkVlanPortInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVlanPortInfoAllOf(classId string, objectType string) *NetworkVlanPortInfoAllOf {
	this := NetworkVlanPortInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVlanPortInfoAllOfWithDefaults instantiates a new NetworkVlanPortInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVlanPortInfoAllOfWithDefaults() *NetworkVlanPortInfoAllOf {
	this := NetworkVlanPortInfoAllOf{}
	var classId string = "network.VlanPortInfo"
	this.ClassId = classId
	var objectType string = "network.VlanPortInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVlanPortInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVlanPortInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVlanPortInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVlanPortInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessVlanPortCount returns the AccessVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetAccessVlanPortCount() int64 {
	if o == nil || o.AccessVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.AccessVlanPortCount
}

// GetAccessVlanPortCountOk returns a tuple with the AccessVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetAccessVlanPortCountOk() (*int64, bool) {
	if o == nil || o.AccessVlanPortCount == nil {
		return nil, false
	}
	return o.AccessVlanPortCount, true
}

// HasAccessVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasAccessVlanPortCount() bool {
	if o != nil && o.AccessVlanPortCount != nil {
		return true
	}

	return false
}

// SetAccessVlanPortCount gets a reference to the given int64 and assigns it to the AccessVlanPortCount field.
func (o *NetworkVlanPortInfoAllOf) SetAccessVlanPortCount(v int64) {
	o.AccessVlanPortCount = &v
}

// GetBorderVlanPortCount returns the BorderVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetBorderVlanPortCount() int64 {
	if o == nil || o.BorderVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.BorderVlanPortCount
}

// GetBorderVlanPortCountOk returns a tuple with the BorderVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetBorderVlanPortCountOk() (*int64, bool) {
	if o == nil || o.BorderVlanPortCount == nil {
		return nil, false
	}
	return o.BorderVlanPortCount, true
}

// HasBorderVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasBorderVlanPortCount() bool {
	if o != nil && o.BorderVlanPortCount != nil {
		return true
	}

	return false
}

// SetBorderVlanPortCount gets a reference to the given int64 and assigns it to the BorderVlanPortCount field.
func (o *NetworkVlanPortInfoAllOf) SetBorderVlanPortCount(v int64) {
	o.BorderVlanPortCount = &v
}

// GetCompressedOptimizationSetsValue returns the CompressedOptimizationSetsValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetCompressedOptimizationSetsValue() int64 {
	if o == nil || o.CompressedOptimizationSetsValue == nil {
		var ret int64
		return ret
	}
	return *o.CompressedOptimizationSetsValue
}

// GetCompressedOptimizationSetsValueOk returns a tuple with the CompressedOptimizationSetsValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetCompressedOptimizationSetsValueOk() (*int64, bool) {
	if o == nil || o.CompressedOptimizationSetsValue == nil {
		return nil, false
	}
	return o.CompressedOptimizationSetsValue, true
}

// HasCompressedOptimizationSetsValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasCompressedOptimizationSetsValue() bool {
	if o != nil && o.CompressedOptimizationSetsValue != nil {
		return true
	}

	return false
}

// SetCompressedOptimizationSetsValue gets a reference to the given int64 and assigns it to the CompressedOptimizationSetsValue field.
func (o *NetworkVlanPortInfoAllOf) SetCompressedOptimizationSetsValue(v int64) {
	o.CompressedOptimizationSetsValue = &v
}

// GetCompressedVlanPortCount returns the CompressedVlanPortCount field value if set, zero value otherwise.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCount() string {
	if o == nil || o.CompressedVlanPortCount == nil {
		var ret string
		return ret
	}
	return *o.CompressedVlanPortCount
}

// GetCompressedVlanPortCountOk returns a tuple with the CompressedVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountOk() (*string, bool) {
	if o == nil || o.CompressedVlanPortCount == nil {
		return nil, false
	}
	return o.CompressedVlanPortCount, true
}

// HasCompressedVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasCompressedVlanPortCount() bool {
	if o != nil && o.CompressedVlanPortCount != nil {
		return true
	}

	return false
}

// SetCompressedVlanPortCount gets a reference to the given string and assigns it to the CompressedVlanPortCount field.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) SetCompressedVlanPortCount(v string) {
	o.CompressedVlanPortCount = &v
}

// GetCompressedVlanPortCountValue returns the CompressedVlanPortCountValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountValue() int64 {
	if o == nil || o.CompressedVlanPortCountValue == nil {
		var ret int64
		return ret
	}
	return *o.CompressedVlanPortCountValue
}

// GetCompressedVlanPortCountValueOk returns a tuple with the CompressedVlanPortCountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetCompressedVlanPortCountValueOk() (*int64, bool) {
	if o == nil || o.CompressedVlanPortCountValue == nil {
		return nil, false
	}
	return o.CompressedVlanPortCountValue, true
}

// HasCompressedVlanPortCountValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasCompressedVlanPortCountValue() bool {
	if o != nil && o.CompressedVlanPortCountValue != nil {
		return true
	}

	return false
}

// SetCompressedVlanPortCountValue gets a reference to the given int64 and assigns it to the CompressedVlanPortCountValue field.
func (o *NetworkVlanPortInfoAllOf) SetCompressedVlanPortCountValue(v int64) {
	o.CompressedVlanPortCountValue = &v
}

// GetTotalVlanPortCount returns the TotalVlanPortCount field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetTotalVlanPortCount() int64 {
	if o == nil || o.TotalVlanPortCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalVlanPortCount
}

// GetTotalVlanPortCountOk returns a tuple with the TotalVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetTotalVlanPortCountOk() (*int64, bool) {
	if o == nil || o.TotalVlanPortCount == nil {
		return nil, false
	}
	return o.TotalVlanPortCount, true
}

// HasTotalVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasTotalVlanPortCount() bool {
	if o != nil && o.TotalVlanPortCount != nil {
		return true
	}

	return false
}

// SetTotalVlanPortCount gets a reference to the given int64 and assigns it to the TotalVlanPortCount field.
func (o *NetworkVlanPortInfoAllOf) SetTotalVlanPortCount(v int64) {
	o.TotalVlanPortCount = &v
}

// GetUncompressedVlanPortCount returns the UncompressedVlanPortCount field value if set, zero value otherwise.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCount() string {
	if o == nil || o.UncompressedVlanPortCount == nil {
		var ret string
		return ret
	}
	return *o.UncompressedVlanPortCount
}

// GetUncompressedVlanPortCountOk returns a tuple with the UncompressedVlanPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountOk() (*string, bool) {
	if o == nil || o.UncompressedVlanPortCount == nil {
		return nil, false
	}
	return o.UncompressedVlanPortCount, true
}

// HasUncompressedVlanPortCount returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasUncompressedVlanPortCount() bool {
	if o != nil && o.UncompressedVlanPortCount != nil {
		return true
	}

	return false
}

// SetUncompressedVlanPortCount gets a reference to the given string and assigns it to the UncompressedVlanPortCount field.
// Deprecated
func (o *NetworkVlanPortInfoAllOf) SetUncompressedVlanPortCount(v string) {
	o.UncompressedVlanPortCount = &v
}

// GetUncompressedVlanPortCountValue returns the UncompressedVlanPortCountValue field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountValue() int64 {
	if o == nil || o.UncompressedVlanPortCountValue == nil {
		var ret int64
		return ret
	}
	return *o.UncompressedVlanPortCountValue
}

// GetUncompressedVlanPortCountValueOk returns a tuple with the UncompressedVlanPortCountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetUncompressedVlanPortCountValueOk() (*int64, bool) {
	if o == nil || o.UncompressedVlanPortCountValue == nil {
		return nil, false
	}
	return o.UncompressedVlanPortCountValue, true
}

// HasUncompressedVlanPortCountValue returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasUncompressedVlanPortCountValue() bool {
	if o != nil && o.UncompressedVlanPortCountValue != nil {
		return true
	}

	return false
}

// SetUncompressedVlanPortCountValue gets a reference to the given int64 and assigns it to the UncompressedVlanPortCountValue field.
func (o *NetworkVlanPortInfoAllOf) SetUncompressedVlanPortCountValue(v int64) {
	o.UncompressedVlanPortCountValue = &v
}

// GetVlanPortLimit returns the VlanPortLimit field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetVlanPortLimit() int64 {
	if o == nil || o.VlanPortLimit == nil {
		var ret int64
		return ret
	}
	return *o.VlanPortLimit
}

// GetVlanPortLimitOk returns a tuple with the VlanPortLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetVlanPortLimitOk() (*int64, bool) {
	if o == nil || o.VlanPortLimit == nil {
		return nil, false
	}
	return o.VlanPortLimit, true
}

// HasVlanPortLimit returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasVlanPortLimit() bool {
	if o != nil && o.VlanPortLimit != nil {
		return true
	}

	return false
}

// SetVlanPortLimit gets a reference to the given int64 and assigns it to the VlanPortLimit field.
func (o *NetworkVlanPortInfoAllOf) SetVlanPortLimit(v int64) {
	o.VlanPortLimit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *NetworkVlanPortInfoAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVlanPortInfoAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkVlanPortInfoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVlanPortInfoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVlanPortInfoAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVlanPortInfoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkVlanPortInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessVlanPortCount != nil {
		toSerialize["AccessVlanPortCount"] = o.AccessVlanPortCount
	}
	if o.BorderVlanPortCount != nil {
		toSerialize["BorderVlanPortCount"] = o.BorderVlanPortCount
	}
	if o.CompressedOptimizationSetsValue != nil {
		toSerialize["CompressedOptimizationSetsValue"] = o.CompressedOptimizationSetsValue
	}
	if o.CompressedVlanPortCount != nil {
		toSerialize["CompressedVlanPortCount"] = o.CompressedVlanPortCount
	}
	if o.CompressedVlanPortCountValue != nil {
		toSerialize["CompressedVlanPortCountValue"] = o.CompressedVlanPortCountValue
	}
	if o.TotalVlanPortCount != nil {
		toSerialize["TotalVlanPortCount"] = o.TotalVlanPortCount
	}
	if o.UncompressedVlanPortCount != nil {
		toSerialize["UncompressedVlanPortCount"] = o.UncompressedVlanPortCount
	}
	if o.UncompressedVlanPortCountValue != nil {
		toSerialize["UncompressedVlanPortCountValue"] = o.UncompressedVlanPortCountValue
	}
	if o.VlanPortLimit != nil {
		toSerialize["VlanPortLimit"] = o.VlanPortLimit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkVlanPortInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkVlanPortInfoAllOf := _NetworkVlanPortInfoAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkVlanPortInfoAllOf); err == nil {
		*o = NetworkVlanPortInfoAllOf(varNetworkVlanPortInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessVlanPortCount")
		delete(additionalProperties, "BorderVlanPortCount")
		delete(additionalProperties, "CompressedOptimizationSetsValue")
		delete(additionalProperties, "CompressedVlanPortCount")
		delete(additionalProperties, "CompressedVlanPortCountValue")
		delete(additionalProperties, "TotalVlanPortCount")
		delete(additionalProperties, "UncompressedVlanPortCount")
		delete(additionalProperties, "UncompressedVlanPortCountValue")
		delete(additionalProperties, "VlanPortLimit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkVlanPortInfoAllOf struct {
	value *NetworkVlanPortInfoAllOf
	isSet bool
}

func (v NullableNetworkVlanPortInfoAllOf) Get() *NetworkVlanPortInfoAllOf {
	return v.value
}

func (v *NullableNetworkVlanPortInfoAllOf) Set(val *NetworkVlanPortInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVlanPortInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVlanPortInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVlanPortInfoAllOf(val *NetworkVlanPortInfoAllOf) *NullableNetworkVlanPortInfoAllOf {
	return &NullableNetworkVlanPortInfoAllOf{value: val, isSet: true}
}

func (v NullableNetworkVlanPortInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVlanPortInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
