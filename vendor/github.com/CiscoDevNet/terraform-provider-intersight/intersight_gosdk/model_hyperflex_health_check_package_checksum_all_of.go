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
	"time"
)

// HyperflexHealthCheckPackageChecksumAllOf Definition of the list of properties defined in 'hyperflex.HealthCheckPackageChecksum', excluding properties defined in parent classes.
type HyperflexHealthCheckPackageChecksumAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SHA512 checksum of the health check package.
	Checksum *string `json:"Checksum,omitempty"`
	// Name of the health check Debian package.
	Name *string `json:"Name,omitempty"`
	// HyperFlex health check Debian package file name.
	PackageName *string `json:"PackageName,omitempty"`
	// Timestamp of last update of HyperFlex health check package checksum.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// HyperFlex health check Debian Package Version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealthCheckPackageChecksumAllOf HyperflexHealthCheckPackageChecksumAllOf

// NewHyperflexHealthCheckPackageChecksumAllOf instantiates a new HyperflexHealthCheckPackageChecksumAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckPackageChecksumAllOf(classId string, objectType string) *HyperflexHealthCheckPackageChecksumAllOf {
	this := HyperflexHealthCheckPackageChecksumAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHealthCheckPackageChecksumAllOfWithDefaults instantiates a new HyperflexHealthCheckPackageChecksumAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckPackageChecksumAllOfWithDefaults() *HyperflexHealthCheckPackageChecksumAllOf {
	this := HyperflexHealthCheckPackageChecksumAllOf{}
	var classId string = "hyperflex.HealthCheckPackageChecksum"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckPackageChecksum"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetChecksum(v string) {
	o.Checksum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetName(v string) {
	o.Name = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetPackageName(v string) {
	o.PackageName = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksumAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHealthCheckPackageChecksumAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexHealthCheckPackageChecksumAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Checksum != nil {
		toSerialize["Checksum"] = o.Checksum
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PackageName != nil {
		toSerialize["PackageName"] = o.PackageName
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckPackageChecksumAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHealthCheckPackageChecksumAllOf := _HyperflexHealthCheckPackageChecksumAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHealthCheckPackageChecksumAllOf); err == nil {
		*o = HyperflexHealthCheckPackageChecksumAllOf(varHyperflexHealthCheckPackageChecksumAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Checksum")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PackageName")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHealthCheckPackageChecksumAllOf struct {
	value *HyperflexHealthCheckPackageChecksumAllOf
	isSet bool
}

func (v NullableHyperflexHealthCheckPackageChecksumAllOf) Get() *HyperflexHealthCheckPackageChecksumAllOf {
	return v.value
}

func (v *NullableHyperflexHealthCheckPackageChecksumAllOf) Set(val *HyperflexHealthCheckPackageChecksumAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckPackageChecksumAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckPackageChecksumAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckPackageChecksumAllOf(val *HyperflexHealthCheckPackageChecksumAllOf) *NullableHyperflexHealthCheckPackageChecksumAllOf {
	return &NullableHyperflexHealthCheckPackageChecksumAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckPackageChecksumAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckPackageChecksumAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
