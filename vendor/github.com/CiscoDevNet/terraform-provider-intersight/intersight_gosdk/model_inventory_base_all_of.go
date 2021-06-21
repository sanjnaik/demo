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

// InventoryBaseAllOf Definition of the list of properties defined in 'inventory.Base', excluding properties defined in parent classes.
type InventoryBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The database identifier of the registered device of an object.
	DeviceMoId *string `json:"DeviceMoId,omitempty"`
	// The Distinguished Name unambiguously identifies an object in the system.
	Dn *string `json:"Dn,omitempty"`
	// The Relative Name uniquely identifies an object within a given context.
	Rn                   *string `json:"Rn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryBaseAllOf InventoryBaseAllOf

// NewInventoryBaseAllOf instantiates a new InventoryBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryBaseAllOf(classId string, objectType string) *InventoryBaseAllOf {
	this := InventoryBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryBaseAllOfWithDefaults instantiates a new InventoryBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryBaseAllOfWithDefaults() *InventoryBaseAllOf {
	this := InventoryBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *InventoryBaseAllOf) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBaseAllOf) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *InventoryBaseAllOf) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *InventoryBaseAllOf) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *InventoryBaseAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBaseAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *InventoryBaseAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *InventoryBaseAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetRn returns the Rn field value if set, zero value otherwise.
func (o *InventoryBaseAllOf) GetRn() string {
	if o == nil || o.Rn == nil {
		var ret string
		return ret
	}
	return *o.Rn
}

// GetRnOk returns a tuple with the Rn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryBaseAllOf) GetRnOk() (*string, bool) {
	if o == nil || o.Rn == nil {
		return nil, false
	}
	return o.Rn, true
}

// HasRn returns a boolean if a field has been set.
func (o *InventoryBaseAllOf) HasRn() bool {
	if o != nil && o.Rn != nil {
		return true
	}

	return false
}

// SetRn gets a reference to the given string and assigns it to the Rn field.
func (o *InventoryBaseAllOf) SetRn(v string) {
	o.Rn = &v
}

func (o InventoryBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Rn != nil {
		toSerialize["Rn"] = o.Rn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryBaseAllOf := _InventoryBaseAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryBaseAllOf); err == nil {
		*o = InventoryBaseAllOf(varInventoryBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceMoId")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Rn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryBaseAllOf struct {
	value *InventoryBaseAllOf
	isSet bool
}

func (v NullableInventoryBaseAllOf) Get() *InventoryBaseAllOf {
	return v.value
}

func (v *NullableInventoryBaseAllOf) Set(val *InventoryBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryBaseAllOf(val *InventoryBaseAllOf) *NullableInventoryBaseAllOf {
	return &NullableInventoryBaseAllOf{value: val, isSet: true}
}

func (v NullableInventoryBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
