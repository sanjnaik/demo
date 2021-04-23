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
)

// HyperflexFilePathAllOf Definition of the list of properties defined in 'hyperflex.FilePath', excluding properties defined in parent classes.
type HyperflexFilePathAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	DsInfo     NullableHyperflexDatastoreInfo `json:"DsInfo,omitempty"`
	// Relative file path within the datastore.
	RelativeFilePath     *string `json:"RelativeFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFilePathAllOf HyperflexFilePathAllOf

// NewHyperflexFilePathAllOf instantiates a new HyperflexFilePathAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFilePathAllOf(classId string, objectType string) *HyperflexFilePathAllOf {
	this := HyperflexFilePathAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexFilePathAllOfWithDefaults instantiates a new HyperflexFilePathAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFilePathAllOfWithDefaults() *HyperflexFilePathAllOf {
	this := HyperflexFilePathAllOf{}
	var classId string = "hyperflex.FilePath"
	this.ClassId = classId
	var objectType string = "hyperflex.FilePath"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexFilePathAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePathAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexFilePathAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexFilePathAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexFilePathAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexFilePathAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDsInfo returns the DsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFilePathAllOf) GetDsInfo() HyperflexDatastoreInfo {
	if o == nil || o.DsInfo.Get() == nil {
		var ret HyperflexDatastoreInfo
		return ret
	}
	return *o.DsInfo.Get()
}

// GetDsInfoOk returns a tuple with the DsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFilePathAllOf) GetDsInfoOk() (*HyperflexDatastoreInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DsInfo.Get(), o.DsInfo.IsSet()
}

// HasDsInfo returns a boolean if a field has been set.
func (o *HyperflexFilePathAllOf) HasDsInfo() bool {
	if o != nil && o.DsInfo.IsSet() {
		return true
	}

	return false
}

// SetDsInfo gets a reference to the given NullableHyperflexDatastoreInfo and assigns it to the DsInfo field.
func (o *HyperflexFilePathAllOf) SetDsInfo(v HyperflexDatastoreInfo) {
	o.DsInfo.Set(&v)
}

// SetDsInfoNil sets the value for DsInfo to be an explicit nil
func (o *HyperflexFilePathAllOf) SetDsInfoNil() {
	o.DsInfo.Set(nil)
}

// UnsetDsInfo ensures that no value is present for DsInfo, not even an explicit nil
func (o *HyperflexFilePathAllOf) UnsetDsInfo() {
	o.DsInfo.Unset()
}

// GetRelativeFilePath returns the RelativeFilePath field value if set, zero value otherwise.
func (o *HyperflexFilePathAllOf) GetRelativeFilePath() string {
	if o == nil || o.RelativeFilePath == nil {
		var ret string
		return ret
	}
	return *o.RelativeFilePath
}

// GetRelativeFilePathOk returns a tuple with the RelativeFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFilePathAllOf) GetRelativeFilePathOk() (*string, bool) {
	if o == nil || o.RelativeFilePath == nil {
		return nil, false
	}
	return o.RelativeFilePath, true
}

// HasRelativeFilePath returns a boolean if a field has been set.
func (o *HyperflexFilePathAllOf) HasRelativeFilePath() bool {
	if o != nil && o.RelativeFilePath != nil {
		return true
	}

	return false
}

// SetRelativeFilePath gets a reference to the given string and assigns it to the RelativeFilePath field.
func (o *HyperflexFilePathAllOf) SetRelativeFilePath(v string) {
	o.RelativeFilePath = &v
}

func (o HyperflexFilePathAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DsInfo.IsSet() {
		toSerialize["DsInfo"] = o.DsInfo.Get()
	}
	if o.RelativeFilePath != nil {
		toSerialize["RelativeFilePath"] = o.RelativeFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexFilePathAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexFilePathAllOf := _HyperflexFilePathAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexFilePathAllOf); err == nil {
		*o = HyperflexFilePathAllOf(varHyperflexFilePathAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DsInfo")
		delete(additionalProperties, "RelativeFilePath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexFilePathAllOf struct {
	value *HyperflexFilePathAllOf
	isSet bool
}

func (v NullableHyperflexFilePathAllOf) Get() *HyperflexFilePathAllOf {
	return v.value
}

func (v *NullableHyperflexFilePathAllOf) Set(val *HyperflexFilePathAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFilePathAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFilePathAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFilePathAllOf(val *HyperflexFilePathAllOf) *NullableHyperflexFilePathAllOf {
	return &NullableHyperflexFilePathAllOf{value: val, isSet: true}
}

func (v NullableHyperflexFilePathAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFilePathAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
