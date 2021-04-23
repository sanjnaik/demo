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

// AccessAddressTypeAllOf Definition of the list of properties defined in 'access.AddressType', excluding properties defined in parent classes.
type AccessAddressTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This flag enables the use of IPv4 address for end-point access.
	EnableIpV4 *bool `json:"EnableIpV4,omitempty"`
	// This flag enables the use of IPv6 address for end-point access.
	EnableIpV6           *bool `json:"EnableIpV6,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessAddressTypeAllOf AccessAddressTypeAllOf

// NewAccessAddressTypeAllOf instantiates a new AccessAddressTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessAddressTypeAllOf(classId string, objectType string) *AccessAddressTypeAllOf {
	this := AccessAddressTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableIpV4 bool = true
	this.EnableIpV4 = &enableIpV4
	var enableIpV6 bool = false
	this.EnableIpV6 = &enableIpV6
	return &this
}

// NewAccessAddressTypeAllOfWithDefaults instantiates a new AccessAddressTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessAddressTypeAllOfWithDefaults() *AccessAddressTypeAllOf {
	this := AccessAddressTypeAllOf{}
	var classId string = "access.AddressType"
	this.ClassId = classId
	var objectType string = "access.AddressType"
	this.ObjectType = objectType
	var enableIpV4 bool = true
	this.EnableIpV4 = &enableIpV4
	var enableIpV6 bool = false
	this.EnableIpV6 = &enableIpV6
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessAddressTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessAddressTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessAddressTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AccessAddressTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessAddressTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessAddressTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableIpV4 returns the EnableIpV4 field value if set, zero value otherwise.
func (o *AccessAddressTypeAllOf) GetEnableIpV4() bool {
	if o == nil || o.EnableIpV4 == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpV4
}

// GetEnableIpV4Ok returns a tuple with the EnableIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAddressTypeAllOf) GetEnableIpV4Ok() (*bool, bool) {
	if o == nil || o.EnableIpV4 == nil {
		return nil, false
	}
	return o.EnableIpV4, true
}

// HasEnableIpV4 returns a boolean if a field has been set.
func (o *AccessAddressTypeAllOf) HasEnableIpV4() bool {
	if o != nil && o.EnableIpV4 != nil {
		return true
	}

	return false
}

// SetEnableIpV4 gets a reference to the given bool and assigns it to the EnableIpV4 field.
func (o *AccessAddressTypeAllOf) SetEnableIpV4(v bool) {
	o.EnableIpV4 = &v
}

// GetEnableIpV6 returns the EnableIpV6 field value if set, zero value otherwise.
func (o *AccessAddressTypeAllOf) GetEnableIpV6() bool {
	if o == nil || o.EnableIpV6 == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpV6
}

// GetEnableIpV6Ok returns a tuple with the EnableIpV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessAddressTypeAllOf) GetEnableIpV6Ok() (*bool, bool) {
	if o == nil || o.EnableIpV6 == nil {
		return nil, false
	}
	return o.EnableIpV6, true
}

// HasEnableIpV6 returns a boolean if a field has been set.
func (o *AccessAddressTypeAllOf) HasEnableIpV6() bool {
	if o != nil && o.EnableIpV6 != nil {
		return true
	}

	return false
}

// SetEnableIpV6 gets a reference to the given bool and assigns it to the EnableIpV6 field.
func (o *AccessAddressTypeAllOf) SetEnableIpV6(v bool) {
	o.EnableIpV6 = &v
}

func (o AccessAddressTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableIpV4 != nil {
		toSerialize["EnableIpV4"] = o.EnableIpV4
	}
	if o.EnableIpV6 != nil {
		toSerialize["EnableIpV6"] = o.EnableIpV6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessAddressTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessAddressTypeAllOf := _AccessAddressTypeAllOf{}

	if err = json.Unmarshal(bytes, &varAccessAddressTypeAllOf); err == nil {
		*o = AccessAddressTypeAllOf(varAccessAddressTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableIpV4")
		delete(additionalProperties, "EnableIpV6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessAddressTypeAllOf struct {
	value *AccessAddressTypeAllOf
	isSet bool
}

func (v NullableAccessAddressTypeAllOf) Get() *AccessAddressTypeAllOf {
	return v.value
}

func (v *NullableAccessAddressTypeAllOf) Set(val *AccessAddressTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessAddressTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessAddressTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessAddressTypeAllOf(val *AccessAddressTypeAllOf) *NullableAccessAddressTypeAllOf {
	return &NullableAccessAddressTypeAllOf{value: val, isSet: true}
}

func (v NullableAccessAddressTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessAddressTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
