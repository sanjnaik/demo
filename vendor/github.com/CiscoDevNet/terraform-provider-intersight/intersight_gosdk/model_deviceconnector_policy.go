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

// DeviceconnectorPolicy Policy to control configuration changes allowed from Cisco IMC.
type DeviceconnectorPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables configuration lockout on the endpoint.
	LockoutEnabled *bool                                 `json:"LockoutEnabled,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceconnectorPolicy DeviceconnectorPolicy

// NewDeviceconnectorPolicy instantiates a new DeviceconnectorPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceconnectorPolicy(classId string, objectType string) *DeviceconnectorPolicy {
	this := DeviceconnectorPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lockoutEnabled bool = true
	this.LockoutEnabled = &lockoutEnabled
	return &this
}

// NewDeviceconnectorPolicyWithDefaults instantiates a new DeviceconnectorPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceconnectorPolicyWithDefaults() *DeviceconnectorPolicy {
	this := DeviceconnectorPolicy{}
	var classId string = "deviceconnector.Policy"
	this.ClassId = classId
	var objectType string = "deviceconnector.Policy"
	this.ObjectType = objectType
	var lockoutEnabled bool = true
	this.LockoutEnabled = &lockoutEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *DeviceconnectorPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *DeviceconnectorPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *DeviceconnectorPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DeviceconnectorPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLockoutEnabled returns the LockoutEnabled field value if set, zero value otherwise.
func (o *DeviceconnectorPolicy) GetLockoutEnabled() bool {
	if o == nil || o.LockoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LockoutEnabled
}

// GetLockoutEnabledOk returns a tuple with the LockoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicy) GetLockoutEnabledOk() (*bool, bool) {
	if o == nil || o.LockoutEnabled == nil {
		return nil, false
	}
	return o.LockoutEnabled, true
}

// HasLockoutEnabled returns a boolean if a field has been set.
func (o *DeviceconnectorPolicy) HasLockoutEnabled() bool {
	if o != nil && o.LockoutEnabled != nil {
		return true
	}

	return false
}

// SetLockoutEnabled gets a reference to the given bool and assigns it to the LockoutEnabled field.
func (o *DeviceconnectorPolicy) SetLockoutEnabled(v bool) {
	o.LockoutEnabled = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DeviceconnectorPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceconnectorPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DeviceconnectorPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *DeviceconnectorPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceconnectorPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceconnectorPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *DeviceconnectorPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *DeviceconnectorPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o DeviceconnectorPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LockoutEnabled != nil {
		toSerialize["LockoutEnabled"] = o.LockoutEnabled
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceconnectorPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type DeviceconnectorPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables configuration lockout on the endpoint.
		LockoutEnabled *bool                                 `json:"LockoutEnabled,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varDeviceconnectorPolicyWithoutEmbeddedStruct := DeviceconnectorPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDeviceconnectorPolicyWithoutEmbeddedStruct)
	if err == nil {
		varDeviceconnectorPolicy := _DeviceconnectorPolicy{}
		varDeviceconnectorPolicy.ClassId = varDeviceconnectorPolicyWithoutEmbeddedStruct.ClassId
		varDeviceconnectorPolicy.ObjectType = varDeviceconnectorPolicyWithoutEmbeddedStruct.ObjectType
		varDeviceconnectorPolicy.LockoutEnabled = varDeviceconnectorPolicyWithoutEmbeddedStruct.LockoutEnabled
		varDeviceconnectorPolicy.Organization = varDeviceconnectorPolicyWithoutEmbeddedStruct.Organization
		varDeviceconnectorPolicy.Profiles = varDeviceconnectorPolicyWithoutEmbeddedStruct.Profiles
		*o = DeviceconnectorPolicy(varDeviceconnectorPolicy)
	} else {
		return err
	}

	varDeviceconnectorPolicy := _DeviceconnectorPolicy{}

	err = json.Unmarshal(bytes, &varDeviceconnectorPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varDeviceconnectorPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LockoutEnabled")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableDeviceconnectorPolicy struct {
	value *DeviceconnectorPolicy
	isSet bool
}

func (v NullableDeviceconnectorPolicy) Get() *DeviceconnectorPolicy {
	return v.value
}

func (v *NullableDeviceconnectorPolicy) Set(val *DeviceconnectorPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceconnectorPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceconnectorPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceconnectorPolicy(val *DeviceconnectorPolicy) *NullableDeviceconnectorPolicy {
	return &NullableDeviceconnectorPolicy{value: val, isSet: true}
}

func (v NullableDeviceconnectorPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceconnectorPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
