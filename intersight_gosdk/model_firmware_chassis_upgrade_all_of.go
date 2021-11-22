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
)

// FirmwareChassisUpgradeAllOf Definition of the list of properties defined in 'firmware.ChassisUpgrade', excluding properties defined in parent classes.
type FirmwareChassisUpgradeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	ExcludeComponentList []string                             `json:"ExcludeComponentList,omitempty"`
	Chassis              *EquipmentChassisRelationship        `json:"Chassis,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareChassisUpgradeAllOf FirmwareChassisUpgradeAllOf

// NewFirmwareChassisUpgradeAllOf instantiates a new FirmwareChassisUpgradeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareChassisUpgradeAllOf(classId string, objectType string) *FirmwareChassisUpgradeAllOf {
	this := FirmwareChassisUpgradeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareChassisUpgradeAllOfWithDefaults instantiates a new FirmwareChassisUpgradeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareChassisUpgradeAllOfWithDefaults() *FirmwareChassisUpgradeAllOf {
	this := FirmwareChassisUpgradeAllOf{}
	var classId string = "firmware.ChassisUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.ChassisUpgrade"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareChassisUpgradeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareChassisUpgradeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareChassisUpgradeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareChassisUpgradeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareChassisUpgradeAllOf) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareChassisUpgradeAllOf) GetExcludeComponentListOk() (*[]string, bool) {
	if o == nil || o.ExcludeComponentList == nil {
		return nil, false
	}
	return &o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeAllOf) HasExcludeComponentList() bool {
	if o != nil && o.ExcludeComponentList != nil {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwareChassisUpgradeAllOf) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeAllOf) GetChassis() EquipmentChassisRelationship {
	if o == nil || o.Chassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeAllOf) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *FirmwareChassisUpgradeAllOf) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareChassisUpgradeAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgradeAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareChassisUpgradeAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareChassisUpgradeAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o FirmwareChassisUpgradeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareChassisUpgradeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareChassisUpgradeAllOf := _FirmwareChassisUpgradeAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareChassisUpgradeAllOf); err == nil {
		*o = FirmwareChassisUpgradeAllOf(varFirmwareChassisUpgradeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "Device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareChassisUpgradeAllOf struct {
	value *FirmwareChassisUpgradeAllOf
	isSet bool
}

func (v NullableFirmwareChassisUpgradeAllOf) Get() *FirmwareChassisUpgradeAllOf {
	return v.value
}

func (v *NullableFirmwareChassisUpgradeAllOf) Set(val *FirmwareChassisUpgradeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareChassisUpgradeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareChassisUpgradeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareChassisUpgradeAllOf(val *FirmwareChassisUpgradeAllOf) *NullableFirmwareChassisUpgradeAllOf {
	return &NullableFirmwareChassisUpgradeAllOf{value: val, isSet: true}
}

func (v NullableFirmwareChassisUpgradeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareChassisUpgradeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
