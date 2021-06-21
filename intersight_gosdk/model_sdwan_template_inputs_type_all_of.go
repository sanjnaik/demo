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

// SdwanTemplateInputsTypeAllOf Definition of the list of properties defined in 'sdwan.TemplateInputsType', excluding properties defined in parent classes.
type SdwanTemplateInputsTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Defines if the input is editable.
	Editable *bool `json:"Editable,omitempty"`
	// Name of the dynamic input key specified in the vManage template.
	Key *string `json:"Key,omitempty"`
	// Defines if the input is optional or required.
	Required *bool `json:"Required,omitempty"`
	// Refers to the name of the vManage template that this inputs belongs to.
	Template *string `json:"Template,omitempty"`
	// Label for the property being saved in the current instance of template Input.
	Title *string `json:"Title,omitempty"`
	// Defines the object type for the input.
	Type *string `json:"Type,omitempty"`
	// Value of the dynamic input key specfied in the vManage template.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanTemplateInputsTypeAllOf SdwanTemplateInputsTypeAllOf

// NewSdwanTemplateInputsTypeAllOf instantiates a new SdwanTemplateInputsTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanTemplateInputsTypeAllOf(classId string, objectType string) *SdwanTemplateInputsTypeAllOf {
	this := SdwanTemplateInputsTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var editable bool = false
	this.Editable = &editable
	var required bool = false
	this.Required = &required
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// NewSdwanTemplateInputsTypeAllOfWithDefaults instantiates a new SdwanTemplateInputsTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanTemplateInputsTypeAllOfWithDefaults() *SdwanTemplateInputsTypeAllOf {
	this := SdwanTemplateInputsTypeAllOf{}
	var classId string = "sdwan.TemplateInputsType"
	this.ClassId = classId
	var objectType string = "sdwan.TemplateInputsType"
	this.ObjectType = objectType
	var editable bool = false
	this.Editable = &editable
	var required bool = false
	this.Required = &required
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdwanTemplateInputsTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdwanTemplateInputsTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdwanTemplateInputsTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdwanTemplateInputsTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *SdwanTemplateInputsTypeAllOf) SetEditable(v bool) {
	o.Editable = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SdwanTemplateInputsTypeAllOf) SetKey(v string) {
	o.Key = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SdwanTemplateInputsTypeAllOf) SetRequired(v bool) {
	o.Required = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *SdwanTemplateInputsTypeAllOf) SetTemplate(v string) {
	o.Template = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SdwanTemplateInputsTypeAllOf) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SdwanTemplateInputsTypeAllOf) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SdwanTemplateInputsTypeAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanTemplateInputsTypeAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SdwanTemplateInputsTypeAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SdwanTemplateInputsTypeAllOf) SetValue(v string) {
	o.Value = &v
}

func (o SdwanTemplateInputsTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Editable != nil {
		toSerialize["Editable"] = o.Editable
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Required != nil {
		toSerialize["Required"] = o.Required
	}
	if o.Template != nil {
		toSerialize["Template"] = o.Template
	}
	if o.Title != nil {
		toSerialize["Title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdwanTemplateInputsTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdwanTemplateInputsTypeAllOf := _SdwanTemplateInputsTypeAllOf{}

	if err = json.Unmarshal(bytes, &varSdwanTemplateInputsTypeAllOf); err == nil {
		*o = SdwanTemplateInputsTypeAllOf(varSdwanTemplateInputsTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Editable")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Required")
		delete(additionalProperties, "Template")
		delete(additionalProperties, "Title")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdwanTemplateInputsTypeAllOf struct {
	value *SdwanTemplateInputsTypeAllOf
	isSet bool
}

func (v NullableSdwanTemplateInputsTypeAllOf) Get() *SdwanTemplateInputsTypeAllOf {
	return v.value
}

func (v *NullableSdwanTemplateInputsTypeAllOf) Set(val *SdwanTemplateInputsTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanTemplateInputsTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanTemplateInputsTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanTemplateInputsTypeAllOf(val *SdwanTemplateInputsTypeAllOf) *NullableSdwanTemplateInputsTypeAllOf {
	return &NullableSdwanTemplateInputsTypeAllOf{value: val, isSet: true}
}

func (v NullableSdwanTemplateInputsTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanTemplateInputsTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
