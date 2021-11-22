/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4903
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowInitiatorContext Details about the workflow initiator name, type, initiator object Id if present and the result handler.
type WorkflowInitiatorContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the Intersight managed object that initiated the workflow.
	InitiatorMoid *string `json:"InitiatorMoid,omitempty"`
	// Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow.
	InitiatorName *string `json:"InitiatorName,omitempty"`
	// Type of Intersight managed object that initiated the workflow.
	InitiatorType        *string `json:"InitiatorType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowInitiatorContext WorkflowInitiatorContext

// NewWorkflowInitiatorContext instantiates a new WorkflowInitiatorContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInitiatorContext(classId string, objectType string) *WorkflowInitiatorContext {
	this := WorkflowInitiatorContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowInitiatorContextWithDefaults instantiates a new WorkflowInitiatorContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInitiatorContextWithDefaults() *WorkflowInitiatorContext {
	this := WorkflowInitiatorContext{}
	var classId string = "workflow.InitiatorContext"
	this.ClassId = classId
	var objectType string = "workflow.InitiatorContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowInitiatorContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowInitiatorContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowInitiatorContext) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowInitiatorContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowInitiatorContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowInitiatorContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitiatorMoid returns the InitiatorMoid field value if set, zero value otherwise.
func (o *WorkflowInitiatorContext) GetInitiatorMoid() string {
	if o == nil || o.InitiatorMoid == nil {
		var ret string
		return ret
	}
	return *o.InitiatorMoid
}

// GetInitiatorMoidOk returns a tuple with the InitiatorMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInitiatorContext) GetInitiatorMoidOk() (*string, bool) {
	if o == nil || o.InitiatorMoid == nil {
		return nil, false
	}
	return o.InitiatorMoid, true
}

// HasInitiatorMoid returns a boolean if a field has been set.
func (o *WorkflowInitiatorContext) HasInitiatorMoid() bool {
	if o != nil && o.InitiatorMoid != nil {
		return true
	}

	return false
}

// SetInitiatorMoid gets a reference to the given string and assigns it to the InitiatorMoid field.
func (o *WorkflowInitiatorContext) SetInitiatorMoid(v string) {
	o.InitiatorMoid = &v
}

// GetInitiatorName returns the InitiatorName field value if set, zero value otherwise.
func (o *WorkflowInitiatorContext) GetInitiatorName() string {
	if o == nil || o.InitiatorName == nil {
		var ret string
		return ret
	}
	return *o.InitiatorName
}

// GetInitiatorNameOk returns a tuple with the InitiatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInitiatorContext) GetInitiatorNameOk() (*string, bool) {
	if o == nil || o.InitiatorName == nil {
		return nil, false
	}
	return o.InitiatorName, true
}

// HasInitiatorName returns a boolean if a field has been set.
func (o *WorkflowInitiatorContext) HasInitiatorName() bool {
	if o != nil && o.InitiatorName != nil {
		return true
	}

	return false
}

// SetInitiatorName gets a reference to the given string and assigns it to the InitiatorName field.
func (o *WorkflowInitiatorContext) SetInitiatorName(v string) {
	o.InitiatorName = &v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *WorkflowInitiatorContext) GetInitiatorType() string {
	if o == nil || o.InitiatorType == nil {
		var ret string
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInitiatorContext) GetInitiatorTypeOk() (*string, bool) {
	if o == nil || o.InitiatorType == nil {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *WorkflowInitiatorContext) HasInitiatorType() bool {
	if o != nil && o.InitiatorType != nil {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given string and assigns it to the InitiatorType field.
func (o *WorkflowInitiatorContext) SetInitiatorType(v string) {
	o.InitiatorType = &v
}

func (o WorkflowInitiatorContext) MarshalJSON() ([]byte, error) {
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
	if o.InitiatorMoid != nil {
		toSerialize["InitiatorMoid"] = o.InitiatorMoid
	}
	if o.InitiatorName != nil {
		toSerialize["InitiatorName"] = o.InitiatorName
	}
	if o.InitiatorType != nil {
		toSerialize["InitiatorType"] = o.InitiatorType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowInitiatorContext) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowInitiatorContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The moid of the Intersight managed object that initiated the workflow.
		InitiatorMoid *string `json:"InitiatorMoid,omitempty"`
		// Name of the initiator who started the workflow. The initiator can be Intersight managed object that triggered the workflow.
		InitiatorName *string `json:"InitiatorName,omitempty"`
		// Type of Intersight managed object that initiated the workflow.
		InitiatorType *string `json:"InitiatorType,omitempty"`
	}

	varWorkflowInitiatorContextWithoutEmbeddedStruct := WorkflowInitiatorContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowInitiatorContextWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowInitiatorContext := _WorkflowInitiatorContext{}
		varWorkflowInitiatorContext.ClassId = varWorkflowInitiatorContextWithoutEmbeddedStruct.ClassId
		varWorkflowInitiatorContext.ObjectType = varWorkflowInitiatorContextWithoutEmbeddedStruct.ObjectType
		varWorkflowInitiatorContext.InitiatorMoid = varWorkflowInitiatorContextWithoutEmbeddedStruct.InitiatorMoid
		varWorkflowInitiatorContext.InitiatorName = varWorkflowInitiatorContextWithoutEmbeddedStruct.InitiatorName
		varWorkflowInitiatorContext.InitiatorType = varWorkflowInitiatorContextWithoutEmbeddedStruct.InitiatorType
		*o = WorkflowInitiatorContext(varWorkflowInitiatorContext)
	} else {
		return err
	}

	varWorkflowInitiatorContext := _WorkflowInitiatorContext{}

	err = json.Unmarshal(bytes, &varWorkflowInitiatorContext)
	if err == nil {
		o.MoBaseComplexType = varWorkflowInitiatorContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitiatorMoid")
		delete(additionalProperties, "InitiatorName")
		delete(additionalProperties, "InitiatorType")

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

type NullableWorkflowInitiatorContext struct {
	value *WorkflowInitiatorContext
	isSet bool
}

func (v NullableWorkflowInitiatorContext) Get() *WorkflowInitiatorContext {
	return v.value
}

func (v *NullableWorkflowInitiatorContext) Set(val *WorkflowInitiatorContext) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInitiatorContext) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInitiatorContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInitiatorContext(val *WorkflowInitiatorContext) *NullableWorkflowInitiatorContext {
	return &NullableWorkflowInitiatorContext{value: val, isSet: true}
}

func (v NullableWorkflowInitiatorContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInitiatorContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
