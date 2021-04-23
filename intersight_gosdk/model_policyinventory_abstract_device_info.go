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
	"reflect"
	"strings"
)

// PolicyinventoryAbstractDeviceInfo Information pertaining to a Registered Device in Intersight which is pertinent to policy microservice.
type PolicyinventoryAbstractDeviceInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Configuration state of server profile config context.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Control action of server profile config context.
	ControlAction *string `json:"ControlAction,omitempty"`
	// Error state of server profile config context.
	ErrorState *string                  `json:"ErrorState,omitempty"`
	JobInfo    []PolicyinventoryJobInfo `json:"JobInfo,omitempty"`
	// Operational state of server profile config context.
	OperState *string `json:"OperState,omitempty"`
	// Server profile MO ID of the server.
	ProfileMoId          *string                              `json:"ProfileMoId,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyinventoryAbstractDeviceInfo PolicyinventoryAbstractDeviceInfo

// NewPolicyinventoryAbstractDeviceInfo instantiates a new PolicyinventoryAbstractDeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyinventoryAbstractDeviceInfo(classId string, objectType string) *PolicyinventoryAbstractDeviceInfo {
	this := PolicyinventoryAbstractDeviceInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyinventoryAbstractDeviceInfoWithDefaults instantiates a new PolicyinventoryAbstractDeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyinventoryAbstractDeviceInfoWithDefaults() *PolicyinventoryAbstractDeviceInfo {
	this := PolicyinventoryAbstractDeviceInfo{}
	var classId string = "inventory.DeviceInfo"
	this.ClassId = classId
	var objectType string = "inventory.DeviceInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyinventoryAbstractDeviceInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyinventoryAbstractDeviceInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyinventoryAbstractDeviceInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyinventoryAbstractDeviceInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *PolicyinventoryAbstractDeviceInfo) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetControlAction returns the ControlAction field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetControlAction() string {
	if o == nil || o.ControlAction == nil {
		var ret string
		return ret
	}
	return *o.ControlAction
}

// GetControlActionOk returns a tuple with the ControlAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetControlActionOk() (*string, bool) {
	if o == nil || o.ControlAction == nil {
		return nil, false
	}
	return o.ControlAction, true
}

// HasControlAction returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasControlAction() bool {
	if o != nil && o.ControlAction != nil {
		return true
	}

	return false
}

// SetControlAction gets a reference to the given string and assigns it to the ControlAction field.
func (o *PolicyinventoryAbstractDeviceInfo) SetControlAction(v string) {
	o.ControlAction = &v
}

// GetErrorState returns the ErrorState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetErrorState() string {
	if o == nil || o.ErrorState == nil {
		var ret string
		return ret
	}
	return *o.ErrorState
}

// GetErrorStateOk returns a tuple with the ErrorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetErrorStateOk() (*string, bool) {
	if o == nil || o.ErrorState == nil {
		return nil, false
	}
	return o.ErrorState, true
}

// HasErrorState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasErrorState() bool {
	if o != nil && o.ErrorState != nil {
		return true
	}

	return false
}

// SetErrorState gets a reference to the given string and assigns it to the ErrorState field.
func (o *PolicyinventoryAbstractDeviceInfo) SetErrorState(v string) {
	o.ErrorState = &v
}

// GetJobInfo returns the JobInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyinventoryAbstractDeviceInfo) GetJobInfo() []PolicyinventoryJobInfo {
	if o == nil {
		var ret []PolicyinventoryJobInfo
		return ret
	}
	return o.JobInfo
}

// GetJobInfoOk returns a tuple with the JobInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyinventoryAbstractDeviceInfo) GetJobInfoOk() (*[]PolicyinventoryJobInfo, bool) {
	if o == nil || o.JobInfo == nil {
		return nil, false
	}
	return &o.JobInfo, true
}

// HasJobInfo returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasJobInfo() bool {
	if o != nil && o.JobInfo != nil {
		return true
	}

	return false
}

// SetJobInfo gets a reference to the given []PolicyinventoryJobInfo and assigns it to the JobInfo field.
func (o *PolicyinventoryAbstractDeviceInfo) SetJobInfo(v []PolicyinventoryJobInfo) {
	o.JobInfo = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PolicyinventoryAbstractDeviceInfo) SetOperState(v string) {
	o.OperState = &v
}

// GetProfileMoId returns the ProfileMoId field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetProfileMoId() string {
	if o == nil || o.ProfileMoId == nil {
		var ret string
		return ret
	}
	return *o.ProfileMoId
}

// GetProfileMoIdOk returns a tuple with the ProfileMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetProfileMoIdOk() (*string, bool) {
	if o == nil || o.ProfileMoId == nil {
		return nil, false
	}
	return o.ProfileMoId, true
}

// HasProfileMoId returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasProfileMoId() bool {
	if o != nil && o.ProfileMoId != nil {
		return true
	}

	return false
}

// SetProfileMoId gets a reference to the given string and assigns it to the ProfileMoId field.
func (o *PolicyinventoryAbstractDeviceInfo) SetProfileMoId(v string) {
	o.ProfileMoId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PolicyinventoryAbstractDeviceInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryAbstractDeviceInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PolicyinventoryAbstractDeviceInfo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PolicyinventoryAbstractDeviceInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PolicyinventoryAbstractDeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.ControlAction != nil {
		toSerialize["ControlAction"] = o.ControlAction
	}
	if o.ErrorState != nil {
		toSerialize["ErrorState"] = o.ErrorState
	}
	if o.JobInfo != nil {
		toSerialize["JobInfo"] = o.JobInfo
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.ProfileMoId != nil {
		toSerialize["ProfileMoId"] = o.ProfileMoId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyinventoryAbstractDeviceInfo) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Configuration state of server profile config context.
		ConfigState *string `json:"ConfigState,omitempty"`
		// Control action of server profile config context.
		ControlAction *string `json:"ControlAction,omitempty"`
		// Error state of server profile config context.
		ErrorState *string                  `json:"ErrorState,omitempty"`
		JobInfo    []PolicyinventoryJobInfo `json:"JobInfo,omitempty"`
		// Operational state of server profile config context.
		OperState *string `json:"OperState,omitempty"`
		// Server profile MO ID of the server.
		ProfileMoId      *string                              `json:"ProfileMoId,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct := PolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct)
	if err == nil {
		varPolicyinventoryAbstractDeviceInfo := _PolicyinventoryAbstractDeviceInfo{}
		varPolicyinventoryAbstractDeviceInfo.ClassId = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ClassId
		varPolicyinventoryAbstractDeviceInfo.ObjectType = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ObjectType
		varPolicyinventoryAbstractDeviceInfo.ConfigState = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ConfigState
		varPolicyinventoryAbstractDeviceInfo.ControlAction = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ControlAction
		varPolicyinventoryAbstractDeviceInfo.ErrorState = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ErrorState
		varPolicyinventoryAbstractDeviceInfo.JobInfo = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.JobInfo
		varPolicyinventoryAbstractDeviceInfo.OperState = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.OperState
		varPolicyinventoryAbstractDeviceInfo.ProfileMoId = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.ProfileMoId
		varPolicyinventoryAbstractDeviceInfo.RegisteredDevice = varPolicyinventoryAbstractDeviceInfoWithoutEmbeddedStruct.RegisteredDevice
		*o = PolicyinventoryAbstractDeviceInfo(varPolicyinventoryAbstractDeviceInfo)
	} else {
		return err
	}

	varPolicyinventoryAbstractDeviceInfo := _PolicyinventoryAbstractDeviceInfo{}

	err = json.Unmarshal(bytes, &varPolicyinventoryAbstractDeviceInfo)
	if err == nil {
		o.MoBaseMo = varPolicyinventoryAbstractDeviceInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "ControlAction")
		delete(additionalProperties, "ErrorState")
		delete(additionalProperties, "JobInfo")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "ProfileMoId")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullablePolicyinventoryAbstractDeviceInfo struct {
	value *PolicyinventoryAbstractDeviceInfo
	isSet bool
}

func (v NullablePolicyinventoryAbstractDeviceInfo) Get() *PolicyinventoryAbstractDeviceInfo {
	return v.value
}

func (v *NullablePolicyinventoryAbstractDeviceInfo) Set(val *PolicyinventoryAbstractDeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyinventoryAbstractDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyinventoryAbstractDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyinventoryAbstractDeviceInfo(val *PolicyinventoryAbstractDeviceInfo) *NullablePolicyinventoryAbstractDeviceInfo {
	return &NullablePolicyinventoryAbstractDeviceInfo{value: val, isSet: true}
}

func (v NullablePolicyinventoryAbstractDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyinventoryAbstractDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
