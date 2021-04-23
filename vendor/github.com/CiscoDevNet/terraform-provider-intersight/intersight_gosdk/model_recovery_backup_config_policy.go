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

// RecoveryBackupConfigPolicy Backup config policy which contains all the required inputs to do backup on a local or remote server.
type RecoveryBackupConfigPolicy struct {
	RecoveryAbstractBackupConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to recoveryBackupProfile resources.
	BackupProfiles       []RecoveryBackupProfileRelationship   `json:"BackupProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryBackupConfigPolicy RecoveryBackupConfigPolicy

// NewRecoveryBackupConfigPolicy instantiates a new RecoveryBackupConfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryBackupConfigPolicy(classId string, objectType string) *RecoveryBackupConfigPolicy {
	this := RecoveryBackupConfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryBackupConfigPolicyWithDefaults instantiates a new RecoveryBackupConfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryBackupConfigPolicyWithDefaults() *RecoveryBackupConfigPolicy {
	this := RecoveryBackupConfigPolicy{}
	var classId string = "recovery.BackupConfigPolicy"
	this.ClassId = classId
	var objectType string = "recovery.BackupConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryBackupConfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryBackupConfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryBackupConfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryBackupConfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryBackupConfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryBackupConfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupProfiles returns the BackupProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryBackupConfigPolicy) GetBackupProfiles() []RecoveryBackupProfileRelationship {
	if o == nil {
		var ret []RecoveryBackupProfileRelationship
		return ret
	}
	return o.BackupProfiles
}

// GetBackupProfilesOk returns a tuple with the BackupProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryBackupConfigPolicy) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool) {
	if o == nil || o.BackupProfiles == nil {
		return nil, false
	}
	return &o.BackupProfiles, true
}

// HasBackupProfiles returns a boolean if a field has been set.
func (o *RecoveryBackupConfigPolicy) HasBackupProfiles() bool {
	if o != nil && o.BackupProfiles != nil {
		return true
	}

	return false
}

// SetBackupProfiles gets a reference to the given []RecoveryBackupProfileRelationship and assigns it to the BackupProfiles field.
func (o *RecoveryBackupConfigPolicy) SetBackupProfiles(v []RecoveryBackupProfileRelationship) {
	o.BackupProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryBackupConfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryBackupConfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryBackupConfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryBackupConfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o RecoveryBackupConfigPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryAbstractBackupConfig, errRecoveryAbstractBackupConfig := json.Marshal(o.RecoveryAbstractBackupConfig)
	if errRecoveryAbstractBackupConfig != nil {
		return []byte{}, errRecoveryAbstractBackupConfig
	}
	errRecoveryAbstractBackupConfig = json.Unmarshal([]byte(serializedRecoveryAbstractBackupConfig), &toSerialize)
	if errRecoveryAbstractBackupConfig != nil {
		return []byte{}, errRecoveryAbstractBackupConfig
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackupProfiles != nil {
		toSerialize["BackupProfiles"] = o.BackupProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryBackupConfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type RecoveryBackupConfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to recoveryBackupProfile resources.
		BackupProfiles []RecoveryBackupProfileRelationship   `json:"BackupProfiles,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varRecoveryBackupConfigPolicyWithoutEmbeddedStruct := RecoveryBackupConfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecoveryBackupConfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryBackupConfigPolicy := _RecoveryBackupConfigPolicy{}
		varRecoveryBackupConfigPolicy.ClassId = varRecoveryBackupConfigPolicyWithoutEmbeddedStruct.ClassId
		varRecoveryBackupConfigPolicy.ObjectType = varRecoveryBackupConfigPolicyWithoutEmbeddedStruct.ObjectType
		varRecoveryBackupConfigPolicy.BackupProfiles = varRecoveryBackupConfigPolicyWithoutEmbeddedStruct.BackupProfiles
		varRecoveryBackupConfigPolicy.Organization = varRecoveryBackupConfigPolicyWithoutEmbeddedStruct.Organization
		*o = RecoveryBackupConfigPolicy(varRecoveryBackupConfigPolicy)
	} else {
		return err
	}

	varRecoveryBackupConfigPolicy := _RecoveryBackupConfigPolicy{}

	err = json.Unmarshal(bytes, &varRecoveryBackupConfigPolicy)
	if err == nil {
		o.RecoveryAbstractBackupConfig = varRecoveryBackupConfigPolicy.RecoveryAbstractBackupConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupProfiles")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectRecoveryAbstractBackupConfig := reflect.ValueOf(o.RecoveryAbstractBackupConfig)
		for i := 0; i < reflectRecoveryAbstractBackupConfig.Type().NumField(); i++ {
			t := reflectRecoveryAbstractBackupConfig.Type().Field(i)

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

type NullableRecoveryBackupConfigPolicy struct {
	value *RecoveryBackupConfigPolicy
	isSet bool
}

func (v NullableRecoveryBackupConfigPolicy) Get() *RecoveryBackupConfigPolicy {
	return v.value
}

func (v *NullableRecoveryBackupConfigPolicy) Set(val *RecoveryBackupConfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryBackupConfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryBackupConfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryBackupConfigPolicy(val *RecoveryBackupConfigPolicy) *NullableRecoveryBackupConfigPolicy {
	return &NullableRecoveryBackupConfigPolicy{value: val, isSet: true}
}

func (v NullableRecoveryBackupConfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryBackupConfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
