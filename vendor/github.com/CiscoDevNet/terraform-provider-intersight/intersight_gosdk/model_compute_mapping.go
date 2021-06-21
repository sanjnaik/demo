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

// ComputeMapping Virtual Media image uploaded on the server.
type ComputeMapping struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Remote image location from where the image is uploaded to server.
	FileLocation *string `json:"FileLocation,omitempty"`
	// The identity assigned to this Virtual Media Image by server.
	Identifier *string `json:"Identifier,omitempty"`
	// Image name of uploaded Virtual Media Image.
	ImageName  *string  `json:"ImageName,omitempty"`
	MediaTypes []string `json:"MediaTypes,omitempty"`
	// Name of Virtual Media mapping assigne by server.
	Name                 *string                              `json:"Name,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Vmedia               *ComputeVmediaRelationship           `json:"Vmedia,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeMapping ComputeMapping

// NewComputeMapping instantiates a new ComputeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeMapping(classId string, objectType string) *ComputeMapping {
	this := ComputeMapping{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeMappingWithDefaults instantiates a new ComputeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeMappingWithDefaults() *ComputeMapping {
	this := ComputeMapping{}
	var classId string = "compute.Mapping"
	this.ClassId = classId
	var objectType string = "compute.Mapping"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeMapping) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeMapping) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeMapping) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeMapping) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *ComputeMapping) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *ComputeMapping) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *ComputeMapping) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ComputeMapping) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ComputeMapping) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ComputeMapping) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *ComputeMapping) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *ComputeMapping) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *ComputeMapping) SetImageName(v string) {
	o.ImageName = &v
}

// GetMediaTypes returns the MediaTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeMapping) GetMediaTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MediaTypes
}

// GetMediaTypesOk returns a tuple with the MediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeMapping) GetMediaTypesOk() (*[]string, bool) {
	if o == nil || o.MediaTypes == nil {
		return nil, false
	}
	return &o.MediaTypes, true
}

// HasMediaTypes returns a boolean if a field has been set.
func (o *ComputeMapping) HasMediaTypes() bool {
	if o != nil && o.MediaTypes != nil {
		return true
	}

	return false
}

// SetMediaTypes gets a reference to the given []string and assigns it to the MediaTypes field.
func (o *ComputeMapping) SetMediaTypes(v []string) {
	o.MediaTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeMapping) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeMapping) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeMapping) SetName(v string) {
	o.Name = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeMapping) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeMapping) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeMapping) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeMapping) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeMapping) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeMapping) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVmedia returns the Vmedia field value if set, zero value otherwise.
func (o *ComputeMapping) GetVmedia() ComputeVmediaRelationship {
	if o == nil || o.Vmedia == nil {
		var ret ComputeVmediaRelationship
		return ret
	}
	return *o.Vmedia
}

// GetVmediaOk returns a tuple with the Vmedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeMapping) GetVmediaOk() (*ComputeVmediaRelationship, bool) {
	if o == nil || o.Vmedia == nil {
		return nil, false
	}
	return o.Vmedia, true
}

// HasVmedia returns a boolean if a field has been set.
func (o *ComputeMapping) HasVmedia() bool {
	if o != nil && o.Vmedia != nil {
		return true
	}

	return false
}

// SetVmedia gets a reference to the given ComputeVmediaRelationship and assigns it to the Vmedia field.
func (o *ComputeMapping) SetVmedia(v ComputeVmediaRelationship) {
	o.Vmedia = &v
}

func (o ComputeMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.Identifier != nil {
		toSerialize["Identifier"] = o.Identifier
	}
	if o.ImageName != nil {
		toSerialize["ImageName"] = o.ImageName
	}
	if o.MediaTypes != nil {
		toSerialize["MediaTypes"] = o.MediaTypes
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Vmedia != nil {
		toSerialize["Vmedia"] = o.Vmedia
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeMapping) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeMappingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Remote image location from where the image is uploaded to server.
		FileLocation *string `json:"FileLocation,omitempty"`
		// The identity assigned to this Virtual Media Image by server.
		Identifier *string `json:"Identifier,omitempty"`
		// Image name of uploaded Virtual Media Image.
		ImageName  *string  `json:"ImageName,omitempty"`
		MediaTypes []string `json:"MediaTypes,omitempty"`
		// Name of Virtual Media mapping assigne by server.
		Name                *string                              `json:"Name,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Vmedia              *ComputeVmediaRelationship           `json:"Vmedia,omitempty"`
	}

	varComputeMappingWithoutEmbeddedStruct := ComputeMappingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeMappingWithoutEmbeddedStruct)
	if err == nil {
		varComputeMapping := _ComputeMapping{}
		varComputeMapping.ClassId = varComputeMappingWithoutEmbeddedStruct.ClassId
		varComputeMapping.ObjectType = varComputeMappingWithoutEmbeddedStruct.ObjectType
		varComputeMapping.FileLocation = varComputeMappingWithoutEmbeddedStruct.FileLocation
		varComputeMapping.Identifier = varComputeMappingWithoutEmbeddedStruct.Identifier
		varComputeMapping.ImageName = varComputeMappingWithoutEmbeddedStruct.ImageName
		varComputeMapping.MediaTypes = varComputeMappingWithoutEmbeddedStruct.MediaTypes
		varComputeMapping.Name = varComputeMappingWithoutEmbeddedStruct.Name
		varComputeMapping.InventoryDeviceInfo = varComputeMappingWithoutEmbeddedStruct.InventoryDeviceInfo
		varComputeMapping.RegisteredDevice = varComputeMappingWithoutEmbeddedStruct.RegisteredDevice
		varComputeMapping.Vmedia = varComputeMappingWithoutEmbeddedStruct.Vmedia
		*o = ComputeMapping(varComputeMapping)
	} else {
		return err
	}

	varComputeMapping := _ComputeMapping{}

	err = json.Unmarshal(bytes, &varComputeMapping)
	if err == nil {
		o.InventoryBase = varComputeMapping.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "Identifier")
		delete(additionalProperties, "ImageName")
		delete(additionalProperties, "MediaTypes")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Vmedia")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableComputeMapping struct {
	value *ComputeMapping
	isSet bool
}

func (v NullableComputeMapping) Get() *ComputeMapping {
	return v.value
}

func (v *NullableComputeMapping) Set(val *ComputeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeMapping(val *ComputeMapping) *NullableComputeMapping {
	return &NullableComputeMapping{value: val, isSet: true}
}

func (v NullableComputeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
