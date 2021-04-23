/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-15T06:27:08Z.
 *
 * API version: 1.0.9-4247
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationMemoryAllocationAllOf Definition of the list of properties defined in 'virtualization.MemoryAllocation', excluding properties defined in parent classes.
type VirtualizationMemoryAllocationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Free memory on the entity in bytes.
	Free *int64 `json:"Free,omitempty"`
	// Reserved memory on the entity in bytes. These reserved memory can be used for system purposes.
	Reserved *int64 `json:"Reserved,omitempty"`
	// The total memory capacity of the entity in bytes.
	Total *int64 `json:"Total,omitempty"`
	// Used or allocated memory on the entity represented in bytes.
	Used                 *int64 `json:"Used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationMemoryAllocationAllOf VirtualizationMemoryAllocationAllOf

// NewVirtualizationMemoryAllocationAllOf instantiates a new VirtualizationMemoryAllocationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationMemoryAllocationAllOf(classId string, objectType string) *VirtualizationMemoryAllocationAllOf {
	this := VirtualizationMemoryAllocationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationMemoryAllocationAllOfWithDefaults instantiates a new VirtualizationMemoryAllocationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationMemoryAllocationAllOfWithDefaults() *VirtualizationMemoryAllocationAllOf {
	this := VirtualizationMemoryAllocationAllOf{}
	var classId string = "virtualization.MemoryAllocation"
	this.ClassId = classId
	var objectType string = "virtualization.MemoryAllocation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationMemoryAllocationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationMemoryAllocationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationMemoryAllocationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationMemoryAllocationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *VirtualizationMemoryAllocationAllOf) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *VirtualizationMemoryAllocationAllOf) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *VirtualizationMemoryAllocationAllOf) SetFree(v int64) {
	o.Free = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *VirtualizationMemoryAllocationAllOf) GetReserved() int64 {
	if o == nil || o.Reserved == nil {
		var ret int64
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetReservedOk() (*int64, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *VirtualizationMemoryAllocationAllOf) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int64 and assigns it to the Reserved field.
func (o *VirtualizationMemoryAllocationAllOf) SetReserved(v int64) {
	o.Reserved = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *VirtualizationMemoryAllocationAllOf) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *VirtualizationMemoryAllocationAllOf) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *VirtualizationMemoryAllocationAllOf) SetTotal(v int64) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *VirtualizationMemoryAllocationAllOf) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationMemoryAllocationAllOf) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *VirtualizationMemoryAllocationAllOf) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *VirtualizationMemoryAllocationAllOf) SetUsed(v int64) {
	o.Used = &v
}

func (o VirtualizationMemoryAllocationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Reserved != nil {
		toSerialize["Reserved"] = o.Reserved
	}
	if o.Total != nil {
		toSerialize["Total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationMemoryAllocationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationMemoryAllocationAllOf := _VirtualizationMemoryAllocationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationMemoryAllocationAllOf); err == nil {
		*o = VirtualizationMemoryAllocationAllOf(varVirtualizationMemoryAllocationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Free")
		delete(additionalProperties, "Reserved")
		delete(additionalProperties, "Total")
		delete(additionalProperties, "Used")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationMemoryAllocationAllOf struct {
	value *VirtualizationMemoryAllocationAllOf
	isSet bool
}

func (v NullableVirtualizationMemoryAllocationAllOf) Get() *VirtualizationMemoryAllocationAllOf {
	return v.value
}

func (v *NullableVirtualizationMemoryAllocationAllOf) Set(val *VirtualizationMemoryAllocationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationMemoryAllocationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationMemoryAllocationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationMemoryAllocationAllOf(val *VirtualizationMemoryAllocationAllOf) *NullableVirtualizationMemoryAllocationAllOf {
	return &NullableVirtualizationMemoryAllocationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationMemoryAllocationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationMemoryAllocationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
