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
	"time"
)

// RecoveryBackupScheduleAllOf Definition of the list of properties defined in 'recovery.BackupSchedule', excluding properties defined in parent classes.
type RecoveryBackupScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time at which the backup is to be run on a given day. This is used when the frequency unit is daily.
	ExecutionTime *time.Time `json:"ExecutionTime,omitempty"`
	// The frequency at which the backup schedule must run. * `Daily` - Allows the user to run the backup daily at a given time. * `Periodic` - Allows the user to run the backup after a certain number of hours.
	FrequencyUnit *string `json:"FrequencyUnit,omitempty"`
	// The frequency, in hours, at which the backup schedule runs. * `8` -  * `4` -  * `12` -  * `16` -  * `20` -
	Hours                *int32 `json:"Hours,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryBackupScheduleAllOf RecoveryBackupScheduleAllOf

// NewRecoveryBackupScheduleAllOf instantiates a new RecoveryBackupScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryBackupScheduleAllOf(classId string, objectType string) *RecoveryBackupScheduleAllOf {
	this := RecoveryBackupScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var frequencyUnit string = "Daily"
	this.FrequencyUnit = &frequencyUnit
	var hours int32 = 8
	this.Hours = &hours
	return &this
}

// NewRecoveryBackupScheduleAllOfWithDefaults instantiates a new RecoveryBackupScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryBackupScheduleAllOfWithDefaults() *RecoveryBackupScheduleAllOf {
	this := RecoveryBackupScheduleAllOf{}
	var classId string = "recovery.BackupSchedule"
	this.ClassId = classId
	var objectType string = "recovery.BackupSchedule"
	this.ObjectType = objectType
	var frequencyUnit string = "Daily"
	this.FrequencyUnit = &frequencyUnit
	var hours int32 = 8
	this.Hours = &hours
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryBackupScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryBackupScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryBackupScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryBackupScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetExecutionTime() time.Time {
	if o == nil || o.ExecutionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetExecutionTimeOk() (*time.Time, bool) {
	if o == nil || o.ExecutionTime == nil {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasExecutionTime() bool {
	if o != nil && o.ExecutionTime != nil {
		return true
	}

	return false
}

// SetExecutionTime gets a reference to the given time.Time and assigns it to the ExecutionTime field.
func (o *RecoveryBackupScheduleAllOf) SetExecutionTime(v time.Time) {
	o.ExecutionTime = &v
}

// GetFrequencyUnit returns the FrequencyUnit field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnit() string {
	if o == nil || o.FrequencyUnit == nil {
		var ret string
		return ret
	}
	return *o.FrequencyUnit
}

// GetFrequencyUnitOk returns a tuple with the FrequencyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetFrequencyUnitOk() (*string, bool) {
	if o == nil || o.FrequencyUnit == nil {
		return nil, false
	}
	return o.FrequencyUnit, true
}

// HasFrequencyUnit returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasFrequencyUnit() bool {
	if o != nil && o.FrequencyUnit != nil {
		return true
	}

	return false
}

// SetFrequencyUnit gets a reference to the given string and assigns it to the FrequencyUnit field.
func (o *RecoveryBackupScheduleAllOf) SetFrequencyUnit(v string) {
	o.FrequencyUnit = &v
}

// GetHours returns the Hours field value if set, zero value otherwise.
func (o *RecoveryBackupScheduleAllOf) GetHours() int32 {
	if o == nil || o.Hours == nil {
		var ret int32
		return ret
	}
	return *o.Hours
}

// GetHoursOk returns a tuple with the Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupScheduleAllOf) GetHoursOk() (*int32, bool) {
	if o == nil || o.Hours == nil {
		return nil, false
	}
	return o.Hours, true
}

// HasHours returns a boolean if a field has been set.
func (o *RecoveryBackupScheduleAllOf) HasHours() bool {
	if o != nil && o.Hours != nil {
		return true
	}

	return false
}

// SetHours gets a reference to the given int32 and assigns it to the Hours field.
func (o *RecoveryBackupScheduleAllOf) SetHours(v int32) {
	o.Hours = &v
}

func (o RecoveryBackupScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExecutionTime != nil {
		toSerialize["ExecutionTime"] = o.ExecutionTime
	}
	if o.FrequencyUnit != nil {
		toSerialize["FrequencyUnit"] = o.FrequencyUnit
	}
	if o.Hours != nil {
		toSerialize["Hours"] = o.Hours
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryBackupScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryBackupScheduleAllOf := _RecoveryBackupScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varRecoveryBackupScheduleAllOf); err == nil {
		*o = RecoveryBackupScheduleAllOf(varRecoveryBackupScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExecutionTime")
		delete(additionalProperties, "FrequencyUnit")
		delete(additionalProperties, "Hours")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryBackupScheduleAllOf struct {
	value *RecoveryBackupScheduleAllOf
	isSet bool
}

func (v NullableRecoveryBackupScheduleAllOf) Get() *RecoveryBackupScheduleAllOf {
	return v.value
}

func (v *NullableRecoveryBackupScheduleAllOf) Set(val *RecoveryBackupScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryBackupScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryBackupScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryBackupScheduleAllOf(val *RecoveryBackupScheduleAllOf) *NullableRecoveryBackupScheduleAllOf {
	return &NullableRecoveryBackupScheduleAllOf{value: val, isSet: true}
}

func (v NullableRecoveryBackupScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryBackupScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
