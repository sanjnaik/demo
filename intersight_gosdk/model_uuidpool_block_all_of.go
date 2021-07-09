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

// UuidpoolBlockAllOf Definition of the list of properties defined in 'uuidpool.Block', excluding properties defined in parent classes.
type UuidpoolBlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                    `json:"ObjectType"`
	UuidSuffixBlock      *UuidpoolUuidBlock        `json:"UuidSuffixBlock,omitempty"`
	Pool                 *UuidpoolPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolBlockAllOf UuidpoolBlockAllOf

// NewUuidpoolBlockAllOf instantiates a new UuidpoolBlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolBlockAllOf(classId string, objectType string) *UuidpoolBlockAllOf {
	this := UuidpoolBlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUuidpoolBlockAllOfWithDefaults instantiates a new UuidpoolBlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolBlockAllOfWithDefaults() *UuidpoolBlockAllOf {
	this := UuidpoolBlockAllOf{}
	var classId string = "uuidpool.Block"
	this.ClassId = classId
	var objectType string = "uuidpool.Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolBlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolBlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolBlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolBlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolBlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolBlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuidSuffixBlock returns the UuidSuffixBlock field value if set, zero value otherwise.
func (o *UuidpoolBlockAllOf) GetUuidSuffixBlock() UuidpoolUuidBlock {
	if o == nil || o.UuidSuffixBlock == nil {
		var ret UuidpoolUuidBlock
		return ret
	}
	return *o.UuidSuffixBlock
}

// GetUuidSuffixBlockOk returns a tuple with the UuidSuffixBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolBlockAllOf) GetUuidSuffixBlockOk() (*UuidpoolUuidBlock, bool) {
	if o == nil || o.UuidSuffixBlock == nil {
		return nil, false
	}
	return o.UuidSuffixBlock, true
}

// HasUuidSuffixBlock returns a boolean if a field has been set.
func (o *UuidpoolBlockAllOf) HasUuidSuffixBlock() bool {
	if o != nil && o.UuidSuffixBlock != nil {
		return true
	}

	return false
}

// SetUuidSuffixBlock gets a reference to the given UuidpoolUuidBlock and assigns it to the UuidSuffixBlock field.
func (o *UuidpoolBlockAllOf) SetUuidSuffixBlock(v UuidpoolUuidBlock) {
	o.UuidSuffixBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolBlockAllOf) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolBlockAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolBlockAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolBlockAllOf) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

func (o UuidpoolBlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UuidSuffixBlock != nil {
		toSerialize["UuidSuffixBlock"] = o.UuidSuffixBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolBlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUuidpoolBlockAllOf := _UuidpoolBlockAllOf{}

	if err = json.Unmarshal(bytes, &varUuidpoolBlockAllOf); err == nil {
		*o = UuidpoolBlockAllOf(varUuidpoolBlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UuidSuffixBlock")
		delete(additionalProperties, "Pool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUuidpoolBlockAllOf struct {
	value *UuidpoolBlockAllOf
	isSet bool
}

func (v NullableUuidpoolBlockAllOf) Get() *UuidpoolBlockAllOf {
	return v.value
}

func (v *NullableUuidpoolBlockAllOf) Set(val *UuidpoolBlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolBlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolBlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolBlockAllOf(val *UuidpoolBlockAllOf) *NullableUuidpoolBlockAllOf {
	return &NullableUuidpoolBlockAllOf{value: val, isSet: true}
}

func (v NullableUuidpoolBlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolBlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
