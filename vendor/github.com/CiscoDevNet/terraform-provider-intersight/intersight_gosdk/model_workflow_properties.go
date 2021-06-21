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

// WorkflowProperties Properties for the task definition like the inputs, outputs, timeout and retry policies. Tasks are the building blocks for workflows.
type WorkflowProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows.
	ExternalMeta     *bool                  `json:"ExternalMeta,omitempty"`
	InputDefinition  []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The number of times a task should be tried before marking as failed.
	RetryCount *int64 `json:"RetryCount,omitempty"`
	// The delay in seconds after which the the task is re-tried.
	RetryDelay *int64 `json:"RetryDelay,omitempty"`
	// The retry policy for the task. * `Fixed` - The enum specifies the option as Fixed where the task retry happens after fixed time specified by RetryDelay.
	RetryPolicy *string `json:"RetryPolicy,omitempty"`
	// Supported status of the definition. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
	SupportStatus *string `json:"SupportStatus,omitempty"`
	// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
	Timeout *int64 `json:"Timeout,omitempty"`
	// The timeout policy for the task. * `Timeout` - The enum specifies the option as Timeout where task will be timed out after the specified time in Timeout property. * `Retry` - The enum specifies the option as Retry where task will be re-tried.
	TimeoutPolicy        *string `json:"TimeoutPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowProperties WorkflowProperties

// NewWorkflowProperties instantiates a new WorkflowProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowProperties(classId string, objectType string) *WorkflowProperties {
	this := WorkflowProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryCount int64 = 3
	this.RetryCount = &retryCount
	var retryDelay int64 = 60
	this.RetryDelay = &retryDelay
	var retryPolicy string = "Fixed"
	this.RetryPolicy = &retryPolicy
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var timeout int64 = 600
	this.Timeout = &timeout
	var timeoutPolicy string = "Timeout"
	this.TimeoutPolicy = &timeoutPolicy
	return &this
}

// NewWorkflowPropertiesWithDefaults instantiates a new WorkflowProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPropertiesWithDefaults() *WorkflowProperties {
	this := WorkflowProperties{}
	var classId string = "workflow.Properties"
	this.ClassId = classId
	var objectType string = "workflow.Properties"
	this.ObjectType = objectType
	var externalMeta bool = false
	this.ExternalMeta = &externalMeta
	var retryCount int64 = 3
	this.RetryCount = &retryCount
	var retryDelay int64 = 60
	this.RetryDelay = &retryDelay
	var retryPolicy string = "Fixed"
	this.RetryPolicy = &retryPolicy
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var timeout int64 = 600
	this.Timeout = &timeout
	var timeoutPolicy string = "Timeout"
	this.TimeoutPolicy = &timeoutPolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowProperties) GetExternalMeta() bool {
	if o == nil || o.ExternalMeta == nil {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetExternalMetaOk() (*bool, bool) {
	if o == nil || o.ExternalMeta == nil {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowProperties) HasExternalMeta() bool {
	if o != nil && o.ExternalMeta != nil {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowProperties) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowProperties) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowProperties) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return &o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowProperties) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowProperties) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowProperties) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowProperties) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return &o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowProperties) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowProperties) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *WorkflowProperties) GetRetryCount() int64 {
	if o == nil || o.RetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetRetryCountOk() (*int64, bool) {
	if o == nil || o.RetryCount == nil {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *WorkflowProperties) HasRetryCount() bool {
	if o != nil && o.RetryCount != nil {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *WorkflowProperties) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *WorkflowProperties) GetRetryDelay() int64 {
	if o == nil || o.RetryDelay == nil {
		var ret int64
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetRetryDelayOk() (*int64, bool) {
	if o == nil || o.RetryDelay == nil {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *WorkflowProperties) HasRetryDelay() bool {
	if o != nil && o.RetryDelay != nil {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given int64 and assigns it to the RetryDelay field.
func (o *WorkflowProperties) SetRetryDelay(v int64) {
	o.RetryDelay = &v
}

// GetRetryPolicy returns the RetryPolicy field value if set, zero value otherwise.
func (o *WorkflowProperties) GetRetryPolicy() string {
	if o == nil || o.RetryPolicy == nil {
		var ret string
		return ret
	}
	return *o.RetryPolicy
}

// GetRetryPolicyOk returns a tuple with the RetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetRetryPolicyOk() (*string, bool) {
	if o == nil || o.RetryPolicy == nil {
		return nil, false
	}
	return o.RetryPolicy, true
}

// HasRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowProperties) HasRetryPolicy() bool {
	if o != nil && o.RetryPolicy != nil {
		return true
	}

	return false
}

// SetRetryPolicy gets a reference to the given string and assigns it to the RetryPolicy field.
func (o *WorkflowProperties) SetRetryPolicy(v string) {
	o.RetryPolicy = &v
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *WorkflowProperties) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *WorkflowProperties) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *WorkflowProperties) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *WorkflowProperties) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *WorkflowProperties) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *WorkflowProperties) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTimeoutPolicy returns the TimeoutPolicy field value if set, zero value otherwise.
func (o *WorkflowProperties) GetTimeoutPolicy() string {
	if o == nil || o.TimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.TimeoutPolicy
}

// GetTimeoutPolicyOk returns a tuple with the TimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowProperties) GetTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.TimeoutPolicy == nil {
		return nil, false
	}
	return o.TimeoutPolicy, true
}

// HasTimeoutPolicy returns a boolean if a field has been set.
func (o *WorkflowProperties) HasTimeoutPolicy() bool {
	if o != nil && o.TimeoutPolicy != nil {
		return true
	}

	return false
}

// SetTimeoutPolicy gets a reference to the given string and assigns it to the TimeoutPolicy field.
func (o *WorkflowProperties) SetTimeoutPolicy(v string) {
	o.TimeoutPolicy = &v
}

func (o WorkflowProperties) MarshalJSON() ([]byte, error) {
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
	if o.ExternalMeta != nil {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.RetryCount != nil {
		toSerialize["RetryCount"] = o.RetryCount
	}
	if o.RetryDelay != nil {
		toSerialize["RetryDelay"] = o.RetryDelay
	}
	if o.RetryPolicy != nil {
		toSerialize["RetryPolicy"] = o.RetryPolicy
	}
	if o.SupportStatus != nil {
		toSerialize["SupportStatus"] = o.SupportStatus
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.TimeoutPolicy != nil {
		toSerialize["TimeoutPolicy"] = o.TimeoutPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowProperties) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When set to false the task definition can only be used by internal system workflows. When set to true then the task can be included in user defined workflows.
		ExternalMeta     *bool                  `json:"ExternalMeta,omitempty"`
		InputDefinition  []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
		OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
		// The number of times a task should be tried before marking as failed.
		RetryCount *int64 `json:"RetryCount,omitempty"`
		// The delay in seconds after which the the task is re-tried.
		RetryDelay *int64 `json:"RetryDelay,omitempty"`
		// The retry policy for the task. * `Fixed` - The enum specifies the option as Fixed where the task retry happens after fixed time specified by RetryDelay.
		RetryPolicy *string `json:"RetryPolicy,omitempty"`
		// Supported status of the definition. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
		SupportStatus *string `json:"SupportStatus,omitempty"`
		// The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days.
		Timeout *int64 `json:"Timeout,omitempty"`
		// The timeout policy for the task. * `Timeout` - The enum specifies the option as Timeout where task will be timed out after the specified time in Timeout property. * `Retry` - The enum specifies the option as Retry where task will be re-tried.
		TimeoutPolicy *string `json:"TimeoutPolicy,omitempty"`
	}

	varWorkflowPropertiesWithoutEmbeddedStruct := WorkflowPropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowProperties := _WorkflowProperties{}
		varWorkflowProperties.ClassId = varWorkflowPropertiesWithoutEmbeddedStruct.ClassId
		varWorkflowProperties.ObjectType = varWorkflowPropertiesWithoutEmbeddedStruct.ObjectType
		varWorkflowProperties.ExternalMeta = varWorkflowPropertiesWithoutEmbeddedStruct.ExternalMeta
		varWorkflowProperties.InputDefinition = varWorkflowPropertiesWithoutEmbeddedStruct.InputDefinition
		varWorkflowProperties.OutputDefinition = varWorkflowPropertiesWithoutEmbeddedStruct.OutputDefinition
		varWorkflowProperties.RetryCount = varWorkflowPropertiesWithoutEmbeddedStruct.RetryCount
		varWorkflowProperties.RetryDelay = varWorkflowPropertiesWithoutEmbeddedStruct.RetryDelay
		varWorkflowProperties.RetryPolicy = varWorkflowPropertiesWithoutEmbeddedStruct.RetryPolicy
		varWorkflowProperties.SupportStatus = varWorkflowPropertiesWithoutEmbeddedStruct.SupportStatus
		varWorkflowProperties.Timeout = varWorkflowPropertiesWithoutEmbeddedStruct.Timeout
		varWorkflowProperties.TimeoutPolicy = varWorkflowPropertiesWithoutEmbeddedStruct.TimeoutPolicy
		*o = WorkflowProperties(varWorkflowProperties)
	} else {
		return err
	}

	varWorkflowProperties := _WorkflowProperties{}

	err = json.Unmarshal(bytes, &varWorkflowProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalMeta")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "RetryCount")
		delete(additionalProperties, "RetryDelay")
		delete(additionalProperties, "RetryPolicy")
		delete(additionalProperties, "SupportStatus")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "TimeoutPolicy")

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

type NullableWorkflowProperties struct {
	value *WorkflowProperties
	isSet bool
}

func (v NullableWorkflowProperties) Get() *WorkflowProperties {
	return v.value
}

func (v *NullableWorkflowProperties) Set(val *WorkflowProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowProperties(val *WorkflowProperties) *NullableWorkflowProperties {
	return &NullableWorkflowProperties{value: val, isSet: true}
}

func (v NullableWorkflowProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
