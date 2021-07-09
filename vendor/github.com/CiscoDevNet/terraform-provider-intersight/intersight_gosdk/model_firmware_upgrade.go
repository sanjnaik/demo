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
	"reflect"
	"strings"
)

// FirmwareUpgrade Firmware upgrade operation for rack and blade servers that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
type FirmwareUpgrade struct {
	FirmwareUpgradeBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	ExcludeComponentList []string                             `json:"ExcludeComponentList,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server               *ComputePhysicalRelationship         `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgrade FirmwareUpgrade

// NewFirmwareUpgrade instantiates a new FirmwareUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgrade(classId string, objectType string) *FirmwareUpgrade {
	this := FirmwareUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeWithDefaults instantiates a new FirmwareUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeWithDefaults() *FirmwareUpgrade {
	this := FirmwareUpgrade{}
	var classId string = "firmware.Upgrade"
	this.ClassId = classId
	var objectType string = "firmware.Upgrade"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgrade) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgrade) GetExcludeComponentListOk() (*[]string, bool) {
	if o == nil || o.ExcludeComponentList == nil {
		return nil, false
	}
	return &o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwareUpgrade) HasExcludeComponentList() bool {
	if o != nil && o.ExcludeComponentList != nil {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwareUpgrade) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareUpgrade) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareUpgrade) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareUpgrade) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *FirmwareUpgrade) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgrade) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareUpgrade) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareUpgrade) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o FirmwareUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeBase, errFirmwareUpgradeBase := json.Marshal(o.FirmwareUpgradeBase)
	if errFirmwareUpgradeBase != nil {
		return []byte{}, errFirmwareUpgradeBase
	}
	errFirmwareUpgradeBase = json.Unmarshal([]byte(serializedFirmwareUpgradeBase), &toSerialize)
	if errFirmwareUpgradeBase != nil {
		return []byte{}, errFirmwareUpgradeBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                               `json:"ObjectType"`
		ExcludeComponentList []string                             `json:"ExcludeComponentList,omitempty"`
		Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Server               *ComputePhysicalRelationship         `json:"Server,omitempty"`
	}

	varFirmwareUpgradeWithoutEmbeddedStruct := FirmwareUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgrade := _FirmwareUpgrade{}
		varFirmwareUpgrade.ClassId = varFirmwareUpgradeWithoutEmbeddedStruct.ClassId
		varFirmwareUpgrade.ObjectType = varFirmwareUpgradeWithoutEmbeddedStruct.ObjectType
		varFirmwareUpgrade.ExcludeComponentList = varFirmwareUpgradeWithoutEmbeddedStruct.ExcludeComponentList
		varFirmwareUpgrade.Device = varFirmwareUpgradeWithoutEmbeddedStruct.Device
		varFirmwareUpgrade.Server = varFirmwareUpgradeWithoutEmbeddedStruct.Server
		*o = FirmwareUpgrade(varFirmwareUpgrade)
	} else {
		return err
	}

	varFirmwareUpgrade := _FirmwareUpgrade{}

	err = json.Unmarshal(bytes, &varFirmwareUpgrade)
	if err == nil {
		o.FirmwareUpgradeBase = varFirmwareUpgrade.FirmwareUpgradeBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectFirmwareUpgradeBase := reflect.ValueOf(o.FirmwareUpgradeBase)
		for i := 0; i < reflectFirmwareUpgradeBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeBase.Type().Field(i)

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

type NullableFirmwareUpgrade struct {
	value *FirmwareUpgrade
	isSet bool
}

func (v NullableFirmwareUpgrade) Get() *FirmwareUpgrade {
	return v.value
}

func (v *NullableFirmwareUpgrade) Set(val *FirmwareUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgrade(val *FirmwareUpgrade) *NullableFirmwareUpgrade {
	return &NullableFirmwareUpgrade{value: val, isSet: true}
}

func (v NullableFirmwareUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
