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

// VirtualizationEsxiVmComputeConfigurationAllOf Definition of the list of properties defined in 'virtualization.EsxiVmComputeConfiguration', excluding properties defined in parent classes.
type VirtualizationEsxiVmComputeConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ResourcePool where virtual machine is deployed.
	ResourcePool         *string `json:"ResourcePool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiVmComputeConfigurationAllOf VirtualizationEsxiVmComputeConfigurationAllOf

// NewVirtualizationEsxiVmComputeConfigurationAllOf instantiates a new VirtualizationEsxiVmComputeConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiVmComputeConfigurationAllOf(classId string, objectType string) *VirtualizationEsxiVmComputeConfigurationAllOf {
	this := VirtualizationEsxiVmComputeConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiVmComputeConfigurationAllOfWithDefaults instantiates a new VirtualizationEsxiVmComputeConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiVmComputeConfigurationAllOfWithDefaults() *VirtualizationEsxiVmComputeConfigurationAllOf {
	this := VirtualizationEsxiVmComputeConfigurationAllOf{}
	var classId string = "virtualization.EsxiVmComputeConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiVmComputeConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *VirtualizationEsxiVmComputeConfigurationAllOf) SetResourcePool(v string) {
	o.ResourcePool = &v
}

func (o VirtualizationEsxiVmComputeConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ResourcePool != nil {
		toSerialize["ResourcePool"] = o.ResourcePool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiVmComputeConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationEsxiVmComputeConfigurationAllOf := _VirtualizationEsxiVmComputeConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationEsxiVmComputeConfigurationAllOf); err == nil {
		*o = VirtualizationEsxiVmComputeConfigurationAllOf(varVirtualizationEsxiVmComputeConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ResourcePool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationEsxiVmComputeConfigurationAllOf struct {
	value *VirtualizationEsxiVmComputeConfigurationAllOf
	isSet bool
}

func (v NullableVirtualizationEsxiVmComputeConfigurationAllOf) Get() *VirtualizationEsxiVmComputeConfigurationAllOf {
	return v.value
}

func (v *NullableVirtualizationEsxiVmComputeConfigurationAllOf) Set(val *VirtualizationEsxiVmComputeConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiVmComputeConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiVmComputeConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiVmComputeConfigurationAllOf(val *VirtualizationEsxiVmComputeConfigurationAllOf) *NullableVirtualizationEsxiVmComputeConfigurationAllOf {
	return &NullableVirtualizationEsxiVmComputeConfigurationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiVmComputeConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiVmComputeConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
