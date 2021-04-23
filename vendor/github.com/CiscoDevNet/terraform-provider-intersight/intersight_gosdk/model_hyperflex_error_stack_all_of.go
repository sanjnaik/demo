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

// HyperflexErrorStackAllOf Definition of the list of properties defined in 'hyperflex.ErrorStack', excluding properties defined in parent classes.
type HyperflexErrorStackAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The error message string for this error stack.
	Message *string `json:"Message,omitempty"`
	// The error message ID for this error stack.
	MessageId            *int64 `json:"MessageId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexErrorStackAllOf HyperflexErrorStackAllOf

// NewHyperflexErrorStackAllOf instantiates a new HyperflexErrorStackAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexErrorStackAllOf(classId string, objectType string) *HyperflexErrorStackAllOf {
	this := HyperflexErrorStackAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexErrorStackAllOfWithDefaults instantiates a new HyperflexErrorStackAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexErrorStackAllOfWithDefaults() *HyperflexErrorStackAllOf {
	this := HyperflexErrorStackAllOf{}
	var classId string = "hyperflex.ErrorStack"
	this.ClassId = classId
	var objectType string = "hyperflex.ErrorStack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexErrorStackAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStackAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexErrorStackAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexErrorStackAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStackAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexErrorStackAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HyperflexErrorStackAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStackAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HyperflexErrorStackAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HyperflexErrorStackAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *HyperflexErrorStackAllOf) GetMessageId() int64 {
	if o == nil || o.MessageId == nil {
		var ret int64
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexErrorStackAllOf) GetMessageIdOk() (*int64, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *HyperflexErrorStackAllOf) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given int64 and assigns it to the MessageId field.
func (o *HyperflexErrorStackAllOf) SetMessageId(v int64) {
	o.MessageId = &v
}

func (o HyperflexErrorStackAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.MessageId != nil {
		toSerialize["MessageId"] = o.MessageId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexErrorStackAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexErrorStackAllOf := _HyperflexErrorStackAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexErrorStackAllOf); err == nil {
		*o = HyperflexErrorStackAllOf(varHyperflexErrorStackAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "MessageId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexErrorStackAllOf struct {
	value *HyperflexErrorStackAllOf
	isSet bool
}

func (v NullableHyperflexErrorStackAllOf) Get() *HyperflexErrorStackAllOf {
	return v.value
}

func (v *NullableHyperflexErrorStackAllOf) Set(val *HyperflexErrorStackAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexErrorStackAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexErrorStackAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexErrorStackAllOf(val *HyperflexErrorStackAllOf) *NullableHyperflexErrorStackAllOf {
	return &NullableHyperflexErrorStackAllOf{value: val, isSet: true}
}

func (v NullableHyperflexErrorStackAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexErrorStackAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
