/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexSnapshotPointAllOf Definition of the list of properties defined in 'hyperflex.SnapshotPoint', excluding properties defined in parent classes.
type HyperflexSnapshotPointAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string                             `json:"ObjectType"`
	ClusterEntityReference       NullableHyperflexEntityReference   `json:"ClusterEntityReference,omitempty"`
	ReplicationStatus            NullableHyperflexReplicationStatus `json:"ReplicationStatus,omitempty"`
	SnapshotFiles                NullableHyperflexSnapshotFiles     `json:"SnapshotFiles,omitempty"`
	SnapshotPointEntityReference NullableHyperflexEntityReference   `json:"SnapshotPointEntityReference,omitempty"`
	SnapshotStatus               NullableHyperflexSnapshotStatus    `json:"SnapshotStatus,omitempty"`
	VmRuntimeInfo                NullableHyperflexVirtualMachine    `json:"VmRuntimeInfo,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _HyperflexSnapshotPointAllOf HyperflexSnapshotPointAllOf

// NewHyperflexSnapshotPointAllOf instantiates a new HyperflexSnapshotPointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotPointAllOf(classId string, objectType string) *HyperflexSnapshotPointAllOf {
	this := HyperflexSnapshotPointAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSnapshotPointAllOfWithDefaults instantiates a new HyperflexSnapshotPointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotPointAllOfWithDefaults() *HyperflexSnapshotPointAllOf {
	this := HyperflexSnapshotPointAllOf{}
	var classId string = "hyperflex.SnapshotPoint"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotPoint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotPointAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotPointAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotPointAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotPointAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotPointAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotPointAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterEntityReference returns the ClusterEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetClusterEntityReference() HyperflexEntityReference {
	if o == nil || o.ClusterEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.ClusterEntityReference.Get()
}

// GetClusterEntityReferenceOk returns a tuple with the ClusterEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetClusterEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterEntityReference.Get(), o.ClusterEntityReference.IsSet()
}

// HasClusterEntityReference returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasClusterEntityReference() bool {
	if o != nil && o.ClusterEntityReference.IsSet() {
		return true
	}

	return false
}

// SetClusterEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the ClusterEntityReference field.
func (o *HyperflexSnapshotPointAllOf) SetClusterEntityReference(v HyperflexEntityReference) {
	o.ClusterEntityReference.Set(&v)
}

// SetClusterEntityReferenceNil sets the value for ClusterEntityReference to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetClusterEntityReferenceNil() {
	o.ClusterEntityReference.Set(nil)
}

// UnsetClusterEntityReference ensures that no value is present for ClusterEntityReference, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetClusterEntityReference() {
	o.ClusterEntityReference.Unset()
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetReplicationStatus() HyperflexReplicationStatus {
	if o == nil || o.ReplicationStatus.Get() == nil {
		var ret HyperflexReplicationStatus
		return ret
	}
	return *o.ReplicationStatus.Get()
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetReplicationStatusOk() (*HyperflexReplicationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationStatus.Get(), o.ReplicationStatus.IsSet()
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasReplicationStatus() bool {
	if o != nil && o.ReplicationStatus.IsSet() {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given NullableHyperflexReplicationStatus and assigns it to the ReplicationStatus field.
func (o *HyperflexSnapshotPointAllOf) SetReplicationStatus(v HyperflexReplicationStatus) {
	o.ReplicationStatus.Set(&v)
}

// SetReplicationStatusNil sets the value for ReplicationStatus to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetReplicationStatusNil() {
	o.ReplicationStatus.Set(nil)
}

// UnsetReplicationStatus ensures that no value is present for ReplicationStatus, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetReplicationStatus() {
	o.ReplicationStatus.Unset()
}

// GetSnapshotFiles returns the SnapshotFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetSnapshotFiles() HyperflexSnapshotFiles {
	if o == nil || o.SnapshotFiles.Get() == nil {
		var ret HyperflexSnapshotFiles
		return ret
	}
	return *o.SnapshotFiles.Get()
}

// GetSnapshotFilesOk returns a tuple with the SnapshotFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetSnapshotFilesOk() (*HyperflexSnapshotFiles, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotFiles.Get(), o.SnapshotFiles.IsSet()
}

// HasSnapshotFiles returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasSnapshotFiles() bool {
	if o != nil && o.SnapshotFiles.IsSet() {
		return true
	}

	return false
}

// SetSnapshotFiles gets a reference to the given NullableHyperflexSnapshotFiles and assigns it to the SnapshotFiles field.
func (o *HyperflexSnapshotPointAllOf) SetSnapshotFiles(v HyperflexSnapshotFiles) {
	o.SnapshotFiles.Set(&v)
}

// SetSnapshotFilesNil sets the value for SnapshotFiles to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetSnapshotFilesNil() {
	o.SnapshotFiles.Set(nil)
}

// UnsetSnapshotFiles ensures that no value is present for SnapshotFiles, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotFiles() {
	o.SnapshotFiles.Unset()
}

// GetSnapshotPointEntityReference returns the SnapshotPointEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetSnapshotPointEntityReference() HyperflexEntityReference {
	if o == nil || o.SnapshotPointEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.SnapshotPointEntityReference.Get()
}

// GetSnapshotPointEntityReferenceOk returns a tuple with the SnapshotPointEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetSnapshotPointEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotPointEntityReference.Get(), o.SnapshotPointEntityReference.IsSet()
}

// HasSnapshotPointEntityReference returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasSnapshotPointEntityReference() bool {
	if o != nil && o.SnapshotPointEntityReference.IsSet() {
		return true
	}

	return false
}

// SetSnapshotPointEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the SnapshotPointEntityReference field.
func (o *HyperflexSnapshotPointAllOf) SetSnapshotPointEntityReference(v HyperflexEntityReference) {
	o.SnapshotPointEntityReference.Set(&v)
}

// SetSnapshotPointEntityReferenceNil sets the value for SnapshotPointEntityReference to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetSnapshotPointEntityReferenceNil() {
	o.SnapshotPointEntityReference.Set(nil)
}

// UnsetSnapshotPointEntityReference ensures that no value is present for SnapshotPointEntityReference, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotPointEntityReference() {
	o.SnapshotPointEntityReference.Unset()
}

// GetSnapshotStatus returns the SnapshotStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetSnapshotStatus() HyperflexSnapshotStatus {
	if o == nil || o.SnapshotStatus.Get() == nil {
		var ret HyperflexSnapshotStatus
		return ret
	}
	return *o.SnapshotStatus.Get()
}

// GetSnapshotStatusOk returns a tuple with the SnapshotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetSnapshotStatusOk() (*HyperflexSnapshotStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotStatus.Get(), o.SnapshotStatus.IsSet()
}

// HasSnapshotStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasSnapshotStatus() bool {
	if o != nil && o.SnapshotStatus.IsSet() {
		return true
	}

	return false
}

// SetSnapshotStatus gets a reference to the given NullableHyperflexSnapshotStatus and assigns it to the SnapshotStatus field.
func (o *HyperflexSnapshotPointAllOf) SetSnapshotStatus(v HyperflexSnapshotStatus) {
	o.SnapshotStatus.Set(&v)
}

// SetSnapshotStatusNil sets the value for SnapshotStatus to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetSnapshotStatusNil() {
	o.SnapshotStatus.Set(nil)
}

// UnsetSnapshotStatus ensures that no value is present for SnapshotStatus, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetSnapshotStatus() {
	o.SnapshotStatus.Unset()
}

// GetVmRuntimeInfo returns the VmRuntimeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotPointAllOf) GetVmRuntimeInfo() HyperflexVirtualMachine {
	if o == nil || o.VmRuntimeInfo.Get() == nil {
		var ret HyperflexVirtualMachine
		return ret
	}
	return *o.VmRuntimeInfo.Get()
}

// GetVmRuntimeInfoOk returns a tuple with the VmRuntimeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotPointAllOf) GetVmRuntimeInfoOk() (*HyperflexVirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmRuntimeInfo.Get(), o.VmRuntimeInfo.IsSet()
}

// HasVmRuntimeInfo returns a boolean if a field has been set.
func (o *HyperflexSnapshotPointAllOf) HasVmRuntimeInfo() bool {
	if o != nil && o.VmRuntimeInfo.IsSet() {
		return true
	}

	return false
}

// SetVmRuntimeInfo gets a reference to the given NullableHyperflexVirtualMachine and assigns it to the VmRuntimeInfo field.
func (o *HyperflexSnapshotPointAllOf) SetVmRuntimeInfo(v HyperflexVirtualMachine) {
	o.VmRuntimeInfo.Set(&v)
}

// SetVmRuntimeInfoNil sets the value for VmRuntimeInfo to be an explicit nil
func (o *HyperflexSnapshotPointAllOf) SetVmRuntimeInfoNil() {
	o.VmRuntimeInfo.Set(nil)
}

// UnsetVmRuntimeInfo ensures that no value is present for VmRuntimeInfo, not even an explicit nil
func (o *HyperflexSnapshotPointAllOf) UnsetVmRuntimeInfo() {
	o.VmRuntimeInfo.Unset()
}

func (o HyperflexSnapshotPointAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterEntityReference.IsSet() {
		toSerialize["ClusterEntityReference"] = o.ClusterEntityReference.Get()
	}
	if o.ReplicationStatus.IsSet() {
		toSerialize["ReplicationStatus"] = o.ReplicationStatus.Get()
	}
	if o.SnapshotFiles.IsSet() {
		toSerialize["SnapshotFiles"] = o.SnapshotFiles.Get()
	}
	if o.SnapshotPointEntityReference.IsSet() {
		toSerialize["SnapshotPointEntityReference"] = o.SnapshotPointEntityReference.Get()
	}
	if o.SnapshotStatus.IsSet() {
		toSerialize["SnapshotStatus"] = o.SnapshotStatus.Get()
	}
	if o.VmRuntimeInfo.IsSet() {
		toSerialize["VmRuntimeInfo"] = o.VmRuntimeInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSnapshotPointAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSnapshotPointAllOf := _HyperflexSnapshotPointAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSnapshotPointAllOf); err == nil {
		*o = HyperflexSnapshotPointAllOf(varHyperflexSnapshotPointAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterEntityReference")
		delete(additionalProperties, "ReplicationStatus")
		delete(additionalProperties, "SnapshotFiles")
		delete(additionalProperties, "SnapshotPointEntityReference")
		delete(additionalProperties, "SnapshotStatus")
		delete(additionalProperties, "VmRuntimeInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSnapshotPointAllOf struct {
	value *HyperflexSnapshotPointAllOf
	isSet bool
}

func (v NullableHyperflexSnapshotPointAllOf) Get() *HyperflexSnapshotPointAllOf {
	return v.value
}

func (v *NullableHyperflexSnapshotPointAllOf) Set(val *HyperflexSnapshotPointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotPointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotPointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotPointAllOf(val *HyperflexSnapshotPointAllOf) *NullableHyperflexSnapshotPointAllOf {
	return &NullableHyperflexSnapshotPointAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotPointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotPointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
