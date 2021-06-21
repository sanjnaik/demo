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

// VnicScpStatusAllOf Definition of the list of properties defined in 'vnic.ScpStatus', excluding properties defined in parent classes.
type VnicScpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the LCP is ready for Deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
	Status               *string                                  `json:"Status,omitempty"`
	VhbaInfo             []VnicVifStatus                          `json:"VhbaInfo,omitempty"`
	Profile              *PolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicScpStatusAllOf VnicScpStatusAllOf

// NewVnicScpStatusAllOf instantiates a new VnicScpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicScpStatusAllOf(classId string, objectType string) *VnicScpStatusAllOf {
	this := VnicScpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewVnicScpStatusAllOfWithDefaults instantiates a new VnicScpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicScpStatusAllOfWithDefaults() *VnicScpStatusAllOf {
	this := VnicScpStatusAllOf{}
	var classId string = "vnic.ScpStatus"
	this.ClassId = classId
	var objectType string = "vnic.ScpStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicScpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicScpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicScpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicScpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicScpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicScpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VnicScpStatusAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScpStatusAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VnicScpStatusAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VnicScpStatusAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VnicScpStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScpStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VnicScpStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VnicScpStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVhbaInfo returns the VhbaInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicScpStatusAllOf) GetVhbaInfo() []VnicVifStatus {
	if o == nil {
		var ret []VnicVifStatus
		return ret
	}
	return o.VhbaInfo
}

// GetVhbaInfoOk returns a tuple with the VhbaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicScpStatusAllOf) GetVhbaInfoOk() (*[]VnicVifStatus, bool) {
	if o == nil || o.VhbaInfo == nil {
		return nil, false
	}
	return &o.VhbaInfo, true
}

// HasVhbaInfo returns a boolean if a field has been set.
func (o *VnicScpStatusAllOf) HasVhbaInfo() bool {
	if o != nil && o.VhbaInfo != nil {
		return true
	}

	return false
}

// SetVhbaInfo gets a reference to the given []VnicVifStatus and assigns it to the VhbaInfo field.
func (o *VnicScpStatusAllOf) SetVhbaInfo(v []VnicVifStatus) {
	o.VhbaInfo = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VnicScpStatusAllOf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScpStatusAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicScpStatusAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicScpStatusAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile = &v
}

func (o VnicScpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VhbaInfo != nil {
		toSerialize["VhbaInfo"] = o.VhbaInfo
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicScpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicScpStatusAllOf := _VnicScpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varVnicScpStatusAllOf); err == nil {
		*o = VnicScpStatusAllOf(varVnicScpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VhbaInfo")
		delete(additionalProperties, "Profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicScpStatusAllOf struct {
	value *VnicScpStatusAllOf
	isSet bool
}

func (v NullableVnicScpStatusAllOf) Get() *VnicScpStatusAllOf {
	return v.value
}

func (v *NullableVnicScpStatusAllOf) Set(val *VnicScpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicScpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicScpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicScpStatusAllOf(val *VnicScpStatusAllOf) *NullableVnicScpStatusAllOf {
	return &NullableVnicScpStatusAllOf{value: val, isSet: true}
}

func (v NullableVnicScpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicScpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
