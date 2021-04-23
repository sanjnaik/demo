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

// EquipmentSlotAllOf Definition of the list of properties defined in 'equipment.Slot', excluding properties defined in parent classes.
type EquipmentSlotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// This field identifies a specific slot on the server.
	SlotId               *string `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSlotAllOf EquipmentSlotAllOf

// NewEquipmentSlotAllOf instantiates a new EquipmentSlotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSlotAllOf(classId string, objectType string) *EquipmentSlotAllOf {
	this := EquipmentSlotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSlotAllOfWithDefaults instantiates a new EquipmentSlotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSlotAllOfWithDefaults() *EquipmentSlotAllOf {
	this := EquipmentSlotAllOf{}
	var classId string = "storage.DiskSlot"
	this.ClassId = classId
	var objectType string = "storage.DiskSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSlotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSlotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSlotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSlotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSlotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSlotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentSlotAllOf) GetSlotId() string {
	if o == nil || o.SlotId == nil {
		var ret string
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSlotAllOf) GetSlotIdOk() (*string, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentSlotAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given string and assigns it to the SlotId field.
func (o *EquipmentSlotAllOf) SetSlotId(v string) {
	o.SlotId = &v
}

func (o EquipmentSlotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSlotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentSlotAllOf := _EquipmentSlotAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentSlotAllOf); err == nil {
		*o = EquipmentSlotAllOf(varEquipmentSlotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SlotId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentSlotAllOf struct {
	value *EquipmentSlotAllOf
	isSet bool
}

func (v NullableEquipmentSlotAllOf) Get() *EquipmentSlotAllOf {
	return v.value
}

func (v *NullableEquipmentSlotAllOf) Set(val *EquipmentSlotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSlotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSlotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSlotAllOf(val *EquipmentSlotAllOf) *NullableEquipmentSlotAllOf {
	return &NullableEquipmentSlotAllOf{value: val, isSet: true}
}

func (v NullableEquipmentSlotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSlotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
