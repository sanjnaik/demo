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

// IaasMostRunTasksAllOf Definition of the list of properties defined in 'iaas.MostRunTasks', excluding properties defined in parent classes.
type IaasMostRunTasksAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A functional area to which a task belongs to.
	TaskCategory *string `json:"TaskCategory,omitempty"`
	// Number of times this task has executed.
	TaskExecutionCount *int64 `json:"TaskExecutionCount,omitempty"`
	// Name of the task executed in UCSD.
	TaskName *string `json:"TaskName,omitempty"`
	// Type of the task whether it is system task or custom task.
	TaskType             *string                   `json:"TaskType,omitempty"`
	Guid                 *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasMostRunTasksAllOf IaasMostRunTasksAllOf

// NewIaasMostRunTasksAllOf instantiates a new IaasMostRunTasksAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasMostRunTasksAllOf(classId string, objectType string) *IaasMostRunTasksAllOf {
	this := IaasMostRunTasksAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasMostRunTasksAllOfWithDefaults instantiates a new IaasMostRunTasksAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasMostRunTasksAllOfWithDefaults() *IaasMostRunTasksAllOf {
	this := IaasMostRunTasksAllOf{}
	var classId string = "iaas.MostRunTasks"
	this.ClassId = classId
	var objectType string = "iaas.MostRunTasks"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasMostRunTasksAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasMostRunTasksAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasMostRunTasksAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasMostRunTasksAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTaskCategory returns the TaskCategory field value if set, zero value otherwise.
func (o *IaasMostRunTasksAllOf) GetTaskCategory() string {
	if o == nil || o.TaskCategory == nil {
		var ret string
		return ret
	}
	return *o.TaskCategory
}

// GetTaskCategoryOk returns a tuple with the TaskCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetTaskCategoryOk() (*string, bool) {
	if o == nil || o.TaskCategory == nil {
		return nil, false
	}
	return o.TaskCategory, true
}

// HasTaskCategory returns a boolean if a field has been set.
func (o *IaasMostRunTasksAllOf) HasTaskCategory() bool {
	if o != nil && o.TaskCategory != nil {
		return true
	}

	return false
}

// SetTaskCategory gets a reference to the given string and assigns it to the TaskCategory field.
func (o *IaasMostRunTasksAllOf) SetTaskCategory(v string) {
	o.TaskCategory = &v
}

// GetTaskExecutionCount returns the TaskExecutionCount field value if set, zero value otherwise.
func (o *IaasMostRunTasksAllOf) GetTaskExecutionCount() int64 {
	if o == nil || o.TaskExecutionCount == nil {
		var ret int64
		return ret
	}
	return *o.TaskExecutionCount
}

// GetTaskExecutionCountOk returns a tuple with the TaskExecutionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetTaskExecutionCountOk() (*int64, bool) {
	if o == nil || o.TaskExecutionCount == nil {
		return nil, false
	}
	return o.TaskExecutionCount, true
}

// HasTaskExecutionCount returns a boolean if a field has been set.
func (o *IaasMostRunTasksAllOf) HasTaskExecutionCount() bool {
	if o != nil && o.TaskExecutionCount != nil {
		return true
	}

	return false
}

// SetTaskExecutionCount gets a reference to the given int64 and assigns it to the TaskExecutionCount field.
func (o *IaasMostRunTasksAllOf) SetTaskExecutionCount(v int64) {
	o.TaskExecutionCount = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *IaasMostRunTasksAllOf) GetTaskName() string {
	if o == nil || o.TaskName == nil {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetTaskNameOk() (*string, bool) {
	if o == nil || o.TaskName == nil {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *IaasMostRunTasksAllOf) HasTaskName() bool {
	if o != nil && o.TaskName != nil {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *IaasMostRunTasksAllOf) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *IaasMostRunTasksAllOf) GetTaskType() string {
	if o == nil || o.TaskType == nil {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetTaskTypeOk() (*string, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *IaasMostRunTasksAllOf) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *IaasMostRunTasksAllOf) SetTaskType(v string) {
	o.TaskType = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasMostRunTasksAllOf) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasMostRunTasksAllOf) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasMostRunTasksAllOf) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasMostRunTasksAllOf) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasMostRunTasksAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TaskCategory != nil {
		toSerialize["TaskCategory"] = o.TaskCategory
	}
	if o.TaskExecutionCount != nil {
		toSerialize["TaskExecutionCount"] = o.TaskExecutionCount
	}
	if o.TaskName != nil {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.TaskType != nil {
		toSerialize["TaskType"] = o.TaskType
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasMostRunTasksAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasMostRunTasksAllOf := _IaasMostRunTasksAllOf{}

	if err = json.Unmarshal(bytes, &varIaasMostRunTasksAllOf); err == nil {
		*o = IaasMostRunTasksAllOf(varIaasMostRunTasksAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TaskCategory")
		delete(additionalProperties, "TaskExecutionCount")
		delete(additionalProperties, "TaskName")
		delete(additionalProperties, "TaskType")
		delete(additionalProperties, "Guid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasMostRunTasksAllOf struct {
	value *IaasMostRunTasksAllOf
	isSet bool
}

func (v NullableIaasMostRunTasksAllOf) Get() *IaasMostRunTasksAllOf {
	return v.value
}

func (v *NullableIaasMostRunTasksAllOf) Set(val *IaasMostRunTasksAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasMostRunTasksAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasMostRunTasksAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasMostRunTasksAllOf(val *IaasMostRunTasksAllOf) *NullableIaasMostRunTasksAllOf {
	return &NullableIaasMostRunTasksAllOf{value: val, isSet: true}
}

func (v NullableIaasMostRunTasksAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasMostRunTasksAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
