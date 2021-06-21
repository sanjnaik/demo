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

// OsValidInstallTarget ValidInstallTarget is used to fetch all the valid Install targets for the servers. The List of Install targets includes Physical Disks and Virtual Drives.
type OsValidInstallTarget struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error message if any errors are encountered while fetching and validating Install targets for the server.
	Error              *string                  `json:"Error,omitempty"`
	M2Jbod             []OsPhysicalDiskResponse `json:"M2Jbod,omitempty"`
	M2VirtualDrives    []OsVirtualDriveResponse `json:"M2VirtualDrives,omitempty"`
	MraidJbod          []OsPhysicalDiskResponse `json:"MraidJbod,omitempty"`
	MraidVirtualDrives []OsVirtualDriveResponse `json:"MraidVirtualDrives,omitempty"`
	// An array of relationships to computePhysical resources.
	Servers              []ComputePhysicalRelationship `json:"Servers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsValidInstallTarget OsValidInstallTarget

// NewOsValidInstallTarget instantiates a new OsValidInstallTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsValidInstallTarget(classId string, objectType string) *OsValidInstallTarget {
	this := OsValidInstallTarget{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsValidInstallTargetWithDefaults instantiates a new OsValidInstallTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsValidInstallTargetWithDefaults() *OsValidInstallTarget {
	this := OsValidInstallTarget{}
	var classId string = "os.ValidInstallTarget"
	this.ClassId = classId
	var objectType string = "os.ValidInstallTarget"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsValidInstallTarget) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsValidInstallTarget) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsValidInstallTarget) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsValidInstallTarget) SetObjectType(v string) {
	o.ObjectType = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OsValidInstallTarget) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsValidInstallTarget) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *OsValidInstallTarget) SetError(v string) {
	o.Error = &v
}

// GetM2Jbod returns the M2Jbod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2Jbod() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.M2Jbod
}

// GetM2JbodOk returns a tuple with the M2Jbod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2JbodOk() (*[]OsPhysicalDiskResponse, bool) {
	if o == nil || o.M2Jbod == nil {
		return nil, false
	}
	return &o.M2Jbod, true
}

// HasM2Jbod returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2Jbod() bool {
	if o != nil && o.M2Jbod != nil {
		return true
	}

	return false
}

// SetM2Jbod gets a reference to the given []OsPhysicalDiskResponse and assigns it to the M2Jbod field.
func (o *OsValidInstallTarget) SetM2Jbod(v []OsPhysicalDiskResponse) {
	o.M2Jbod = v
}

// GetM2VirtualDrives returns the M2VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetM2VirtualDrives() []OsVirtualDriveResponse {
	if o == nil {
		var ret []OsVirtualDriveResponse
		return ret
	}
	return o.M2VirtualDrives
}

// GetM2VirtualDrivesOk returns a tuple with the M2VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetM2VirtualDrivesOk() (*[]OsVirtualDriveResponse, bool) {
	if o == nil || o.M2VirtualDrives == nil {
		return nil, false
	}
	return &o.M2VirtualDrives, true
}

// HasM2VirtualDrives returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasM2VirtualDrives() bool {
	if o != nil && o.M2VirtualDrives != nil {
		return true
	}

	return false
}

// SetM2VirtualDrives gets a reference to the given []OsVirtualDriveResponse and assigns it to the M2VirtualDrives field.
func (o *OsValidInstallTarget) SetM2VirtualDrives(v []OsVirtualDriveResponse) {
	o.M2VirtualDrives = v
}

// GetMraidJbod returns the MraidJbod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetMraidJbod() []OsPhysicalDiskResponse {
	if o == nil {
		var ret []OsPhysicalDiskResponse
		return ret
	}
	return o.MraidJbod
}

// GetMraidJbodOk returns a tuple with the MraidJbod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetMraidJbodOk() (*[]OsPhysicalDiskResponse, bool) {
	if o == nil || o.MraidJbod == nil {
		return nil, false
	}
	return &o.MraidJbod, true
}

// HasMraidJbod returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasMraidJbod() bool {
	if o != nil && o.MraidJbod != nil {
		return true
	}

	return false
}

// SetMraidJbod gets a reference to the given []OsPhysicalDiskResponse and assigns it to the MraidJbod field.
func (o *OsValidInstallTarget) SetMraidJbod(v []OsPhysicalDiskResponse) {
	o.MraidJbod = v
}

// GetMraidVirtualDrives returns the MraidVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetMraidVirtualDrives() []OsVirtualDriveResponse {
	if o == nil {
		var ret []OsVirtualDriveResponse
		return ret
	}
	return o.MraidVirtualDrives
}

// GetMraidVirtualDrivesOk returns a tuple with the MraidVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetMraidVirtualDrivesOk() (*[]OsVirtualDriveResponse, bool) {
	if o == nil || o.MraidVirtualDrives == nil {
		return nil, false
	}
	return &o.MraidVirtualDrives, true
}

// HasMraidVirtualDrives returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasMraidVirtualDrives() bool {
	if o != nil && o.MraidVirtualDrives != nil {
		return true
	}

	return false
}

// SetMraidVirtualDrives gets a reference to the given []OsVirtualDriveResponse and assigns it to the MraidVirtualDrives field.
func (o *OsValidInstallTarget) SetMraidVirtualDrives(v []OsVirtualDriveResponse) {
	o.MraidVirtualDrives = v
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsValidInstallTarget) GetServers() []ComputePhysicalRelationship {
	if o == nil {
		var ret []ComputePhysicalRelationship
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsValidInstallTarget) GetServersOk() (*[]ComputePhysicalRelationship, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return &o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *OsValidInstallTarget) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ComputePhysicalRelationship and assigns it to the Servers field.
func (o *OsValidInstallTarget) SetServers(v []ComputePhysicalRelationship) {
	o.Servers = v
}

func (o OsValidInstallTarget) MarshalJSON() ([]byte, error) {
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
	if o.Error != nil {
		toSerialize["Error"] = o.Error
	}
	if o.M2Jbod != nil {
		toSerialize["M2Jbod"] = o.M2Jbod
	}
	if o.M2VirtualDrives != nil {
		toSerialize["M2VirtualDrives"] = o.M2VirtualDrives
	}
	if o.MraidJbod != nil {
		toSerialize["MraidJbod"] = o.MraidJbod
	}
	if o.MraidVirtualDrives != nil {
		toSerialize["MraidVirtualDrives"] = o.MraidVirtualDrives
	}
	if o.Servers != nil {
		toSerialize["Servers"] = o.Servers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsValidInstallTarget) UnmarshalJSON(bytes []byte) (err error) {
	type OsValidInstallTargetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Error message if any errors are encountered while fetching and validating Install targets for the server.
		Error              *string                  `json:"Error,omitempty"`
		M2Jbod             []OsPhysicalDiskResponse `json:"M2Jbod,omitempty"`
		M2VirtualDrives    []OsVirtualDriveResponse `json:"M2VirtualDrives,omitempty"`
		MraidJbod          []OsPhysicalDiskResponse `json:"MraidJbod,omitempty"`
		MraidVirtualDrives []OsVirtualDriveResponse `json:"MraidVirtualDrives,omitempty"`
		// An array of relationships to computePhysical resources.
		Servers []ComputePhysicalRelationship `json:"Servers,omitempty"`
	}

	varOsValidInstallTargetWithoutEmbeddedStruct := OsValidInstallTargetWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsValidInstallTargetWithoutEmbeddedStruct)
	if err == nil {
		varOsValidInstallTarget := _OsValidInstallTarget{}
		varOsValidInstallTarget.ClassId = varOsValidInstallTargetWithoutEmbeddedStruct.ClassId
		varOsValidInstallTarget.ObjectType = varOsValidInstallTargetWithoutEmbeddedStruct.ObjectType
		varOsValidInstallTarget.Error = varOsValidInstallTargetWithoutEmbeddedStruct.Error
		varOsValidInstallTarget.M2Jbod = varOsValidInstallTargetWithoutEmbeddedStruct.M2Jbod
		varOsValidInstallTarget.M2VirtualDrives = varOsValidInstallTargetWithoutEmbeddedStruct.M2VirtualDrives
		varOsValidInstallTarget.MraidJbod = varOsValidInstallTargetWithoutEmbeddedStruct.MraidJbod
		varOsValidInstallTarget.MraidVirtualDrives = varOsValidInstallTargetWithoutEmbeddedStruct.MraidVirtualDrives
		varOsValidInstallTarget.Servers = varOsValidInstallTargetWithoutEmbeddedStruct.Servers
		*o = OsValidInstallTarget(varOsValidInstallTarget)
	} else {
		return err
	}

	varOsValidInstallTarget := _OsValidInstallTarget{}

	err = json.Unmarshal(bytes, &varOsValidInstallTarget)
	if err == nil {
		o.MoBaseMo = varOsValidInstallTarget.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "M2Jbod")
		delete(additionalProperties, "M2VirtualDrives")
		delete(additionalProperties, "MraidJbod")
		delete(additionalProperties, "MraidVirtualDrives")
		delete(additionalProperties, "Servers")

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

type NullableOsValidInstallTarget struct {
	value *OsValidInstallTarget
	isSet bool
}

func (v NullableOsValidInstallTarget) Get() *OsValidInstallTarget {
	return v.value
}

func (v *NullableOsValidInstallTarget) Set(val *OsValidInstallTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableOsValidInstallTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableOsValidInstallTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsValidInstallTarget(val *OsValidInstallTarget) *NullableOsValidInstallTarget {
	return &NullableOsValidInstallTarget{value: val, isSet: true}
}

func (v NullableOsValidInstallTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsValidInstallTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
