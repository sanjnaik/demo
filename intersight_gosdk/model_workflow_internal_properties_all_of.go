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

// WorkflowInternalPropertiesAllOf Definition of the list of properties defined in 'workflow.InternalProperties', excluding properties defined in parent classes.
type WorkflowInternalPropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field will hold the base task type like HttpBaseTask or RemoteAnsibleBaseTask.
	BaseTaskType *string                         `json:"BaseTaskType,omitempty"`
	Constraints  NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
	// Denotes this is an internal task. Internal tasks will be hidden from the UI when executing a workflow.
	Internal *bool `json:"Internal,omitempty"`
	// The service that owns and is responsible for execution of the task.
	Owner                *string `json:"Owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowInternalPropertiesAllOf WorkflowInternalPropertiesAllOf

// NewWorkflowInternalPropertiesAllOf instantiates a new WorkflowInternalPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowInternalPropertiesAllOf(classId string, objectType string) *WorkflowInternalPropertiesAllOf {
	this := WorkflowInternalPropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowInternalPropertiesAllOfWithDefaults instantiates a new WorkflowInternalPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowInternalPropertiesAllOfWithDefaults() *WorkflowInternalPropertiesAllOf {
	this := WorkflowInternalPropertiesAllOf{}
	var classId string = "workflow.InternalProperties"
	this.ClassId = classId
	var objectType string = "workflow.InternalProperties"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowInternalPropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowInternalPropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowInternalPropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowInternalPropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowInternalPropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowInternalPropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaseTaskType returns the BaseTaskType field value if set, zero value otherwise.
func (o *WorkflowInternalPropertiesAllOf) GetBaseTaskType() string {
	if o == nil || o.BaseTaskType == nil {
		var ret string
		return ret
	}
	return *o.BaseTaskType
}

// GetBaseTaskTypeOk returns a tuple with the BaseTaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalPropertiesAllOf) GetBaseTaskTypeOk() (*string, bool) {
	if o == nil || o.BaseTaskType == nil {
		return nil, false
	}
	return o.BaseTaskType, true
}

// HasBaseTaskType returns a boolean if a field has been set.
func (o *WorkflowInternalPropertiesAllOf) HasBaseTaskType() bool {
	if o != nil && o.BaseTaskType != nil {
		return true
	}

	return false
}

// SetBaseTaskType gets a reference to the given string and assigns it to the BaseTaskType field.
func (o *WorkflowInternalPropertiesAllOf) SetBaseTaskType(v string) {
	o.BaseTaskType = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowInternalPropertiesAllOf) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowInternalPropertiesAllOf) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowInternalPropertiesAllOf) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowInternalPropertiesAllOf) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowInternalPropertiesAllOf) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowInternalPropertiesAllOf) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *WorkflowInternalPropertiesAllOf) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalPropertiesAllOf) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *WorkflowInternalPropertiesAllOf) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *WorkflowInternalPropertiesAllOf) SetInternal(v bool) {
	o.Internal = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *WorkflowInternalPropertiesAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowInternalPropertiesAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *WorkflowInternalPropertiesAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *WorkflowInternalPropertiesAllOf) SetOwner(v string) {
	o.Owner = &v
}

func (o WorkflowInternalPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BaseTaskType != nil {
		toSerialize["BaseTaskType"] = o.BaseTaskType
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowInternalPropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowInternalPropertiesAllOf := _WorkflowInternalPropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowInternalPropertiesAllOf); err == nil {
		*o = WorkflowInternalPropertiesAllOf(varWorkflowInternalPropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseTaskType")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowInternalPropertiesAllOf struct {
	value *WorkflowInternalPropertiesAllOf
	isSet bool
}

func (v NullableWorkflowInternalPropertiesAllOf) Get() *WorkflowInternalPropertiesAllOf {
	return v.value
}

func (v *NullableWorkflowInternalPropertiesAllOf) Set(val *WorkflowInternalPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowInternalPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowInternalPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowInternalPropertiesAllOf(val *WorkflowInternalPropertiesAllOf) *NullableWorkflowInternalPropertiesAllOf {
	return &NullableWorkflowInternalPropertiesAllOf{value: val, isSet: true}
}

func (v NullableWorkflowInternalPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowInternalPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
