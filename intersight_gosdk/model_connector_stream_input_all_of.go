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

// ConnectorStreamInputAllOf Definition of the list of properties defined in 'connector.StreamInput', excluding properties defined in parent classes.
type ConnectorStreamInputAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The input to write to the stream plugin.
	Input                *string `json:"Input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamInputAllOf ConnectorStreamInputAllOf

// NewConnectorStreamInputAllOf instantiates a new ConnectorStreamInputAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamInputAllOf(classId string, objectType string) *ConnectorStreamInputAllOf {
	this := ConnectorStreamInputAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamInputAllOfWithDefaults instantiates a new ConnectorStreamInputAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamInputAllOfWithDefaults() *ConnectorStreamInputAllOf {
	this := ConnectorStreamInputAllOf{}
	var classId string = "connector.StreamInput"
	this.ClassId = classId
	var objectType string = "connector.StreamInput"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStreamInputAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamInputAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStreamInputAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStreamInputAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamInputAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStreamInputAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ConnectorStreamInputAllOf) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamInputAllOf) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ConnectorStreamInputAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *ConnectorStreamInputAllOf) SetInput(v string) {
	o.Input = &v
}

func (o ConnectorStreamInputAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamInputAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorStreamInputAllOf := _ConnectorStreamInputAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorStreamInputAllOf); err == nil {
		*o = ConnectorStreamInputAllOf(varConnectorStreamInputAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorStreamInputAllOf struct {
	value *ConnectorStreamInputAllOf
	isSet bool
}

func (v NullableConnectorStreamInputAllOf) Get() *ConnectorStreamInputAllOf {
	return v.value
}

func (v *NullableConnectorStreamInputAllOf) Set(val *ConnectorStreamInputAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamInputAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamInputAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamInputAllOf(val *ConnectorStreamInputAllOf) *NullableConnectorStreamInputAllOf {
	return &NullableConnectorStreamInputAllOf{value: val, isSet: true}
}

func (v NullableConnectorStreamInputAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamInputAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
