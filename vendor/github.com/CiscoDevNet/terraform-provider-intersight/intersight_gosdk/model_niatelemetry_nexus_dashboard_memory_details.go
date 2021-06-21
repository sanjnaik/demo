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

// NiatelemetryNexusDashboardMemoryDetails Details of Nexus Dashboard's memory.
type NiatelemetryNexusDashboardMemoryDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the node in Nexus Dashboard cluster.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Memory capacity of a node in Nexus Dashboard.
	MemoryCapacity       *int64                               `json:"MemoryCapacity,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboardMemoryDetails NiatelemetryNexusDashboardMemoryDetails

// NewNiatelemetryNexusDashboardMemoryDetails instantiates a new NiatelemetryNexusDashboardMemoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardMemoryDetails(classId string, objectType string) *NiatelemetryNexusDashboardMemoryDetails {
	this := NiatelemetryNexusDashboardMemoryDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardMemoryDetailsWithDefaults instantiates a new NiatelemetryNexusDashboardMemoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardMemoryDetailsWithDefaults() *NiatelemetryNexusDashboardMemoryDetails {
	this := NiatelemetryNexusDashboardMemoryDetails{}
	var classId string = "niatelemetry.NexusDashboardMemoryDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboardMemoryDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardMemoryDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardMemoryDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardMemoryDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardMemoryDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *NiatelemetryNexusDashboardMemoryDetails) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetMemoryCapacity() int64 {
	if o == nil || o.MemoryCapacity == nil {
		var ret int64
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetMemoryCapacityOk() (*int64, bool) {
	if o == nil || o.MemoryCapacity == nil {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity != nil {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given int64 and assigns it to the MemoryCapacity field.
func (o *NiatelemetryNexusDashboardMemoryDetails) SetMemoryCapacity(v int64) {
	o.MemoryCapacity = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardMemoryDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardMemoryDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardMemoryDetails) MarshalJSON() ([]byte, error) {
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
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.MemoryCapacity != nil {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardMemoryDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the node in Nexus Dashboard cluster.
		DeviceName *string `json:"DeviceName,omitempty"`
		// Memory capacity of a node in Nexus Dashboard.
		MemoryCapacity   *int64                               `json:"MemoryCapacity,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct := NiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNexusDashboardMemoryDetails := _NiatelemetryNexusDashboardMemoryDetails{}
		varNiatelemetryNexusDashboardMemoryDetails.ClassId = varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryNexusDashboardMemoryDetails.ObjectType = varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNexusDashboardMemoryDetails.DeviceName = varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct.DeviceName
		varNiatelemetryNexusDashboardMemoryDetails.MemoryCapacity = varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct.MemoryCapacity
		varNiatelemetryNexusDashboardMemoryDetails.RegisteredDevice = varNiatelemetryNexusDashboardMemoryDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryNexusDashboardMemoryDetails(varNiatelemetryNexusDashboardMemoryDetails)
	} else {
		return err
	}

	varNiatelemetryNexusDashboardMemoryDetails := _NiatelemetryNexusDashboardMemoryDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardMemoryDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryNexusDashboardMemoryDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "MemoryCapacity")
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

type NullableNiatelemetryNexusDashboardMemoryDetails struct {
	value *NiatelemetryNexusDashboardMemoryDetails
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardMemoryDetails) Get() *NiatelemetryNexusDashboardMemoryDetails {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardMemoryDetails) Set(val *NiatelemetryNexusDashboardMemoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardMemoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardMemoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardMemoryDetails(val *NiatelemetryNexusDashboardMemoryDetails) *NullableNiatelemetryNexusDashboardMemoryDetails {
	return &NullableNiatelemetryNexusDashboardMemoryDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardMemoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardMemoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
