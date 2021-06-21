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
	"reflect"
	"strings"
)

// TamIdentifiers Identifiers are used to identify an affected object in an alert defition. please refer to https://Intersight.com/apidocs/introduction/#get-requests-and-query-capabilities for more details.
type TamIdentifiers struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the filter paramter.
	Name *string `json:"Name,omitempty"`
	// Value of the filter paramter.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamIdentifiers TamIdentifiers

// NewTamIdentifiers instantiates a new TamIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamIdentifiers(classId string, objectType string) *TamIdentifiers {
	this := TamIdentifiers{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamIdentifiersWithDefaults instantiates a new TamIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamIdentifiersWithDefaults() *TamIdentifiers {
	this := TamIdentifiers{}
	var classId string = "tam.Identifiers"
	this.ClassId = classId
	var objectType string = "tam.Identifiers"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamIdentifiers) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamIdentifiers) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamIdentifiers) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamIdentifiers) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamIdentifiers) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamIdentifiers) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamIdentifiers) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamIdentifiers) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamIdentifiers) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamIdentifiers) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TamIdentifiers) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamIdentifiers) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TamIdentifiers) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TamIdentifiers) SetValue(v string) {
	o.Value = &v
}

func (o TamIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamIdentifiers) UnmarshalJSON(bytes []byte) (err error) {
	type TamIdentifiersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the filter paramter.
		Name *string `json:"Name,omitempty"`
		// Value of the filter paramter.
		Value *string `json:"Value,omitempty"`
	}

	varTamIdentifiersWithoutEmbeddedStruct := TamIdentifiersWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamIdentifiersWithoutEmbeddedStruct)
	if err == nil {
		varTamIdentifiers := _TamIdentifiers{}
		varTamIdentifiers.ClassId = varTamIdentifiersWithoutEmbeddedStruct.ClassId
		varTamIdentifiers.ObjectType = varTamIdentifiersWithoutEmbeddedStruct.ObjectType
		varTamIdentifiers.Name = varTamIdentifiersWithoutEmbeddedStruct.Name
		varTamIdentifiers.Value = varTamIdentifiersWithoutEmbeddedStruct.Value
		*o = TamIdentifiers(varTamIdentifiers)
	} else {
		return err
	}

	varTamIdentifiers := _TamIdentifiers{}

	err = json.Unmarshal(bytes, &varTamIdentifiers)
	if err == nil {
		o.MoBaseComplexType = varTamIdentifiers.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableTamIdentifiers struct {
	value *TamIdentifiers
	isSet bool
}

func (v NullableTamIdentifiers) Get() *TamIdentifiers {
	return v.value
}

func (v *NullableTamIdentifiers) Set(val *TamIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableTamIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableTamIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamIdentifiers(val *TamIdentifiers) *NullableTamIdentifiers {
	return &NullableTamIdentifiers{value: val, isSet: true}
}

func (v NullableTamIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
