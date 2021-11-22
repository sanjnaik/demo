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

// BiosVfSelectMemoryRasConfigurationAllOf Definition of the list of properties defined in 'bios.VfSelectMemoryRasConfiguration', excluding properties defined in parent classes.
type BiosVfSelectMemoryRasConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Parent server serial number.
	Serial *string `json:"Serial,omitempty"`
	// The actual BIOS memory RAS configuration as reported by the platform BIOS. Possible values are \"maximum-performance\", \"mirror-mode-1lm\", \"adddc-sparing\", \"platform-default\", \"lockstep\", \"sparing\", \"mirroring\".
	VpSelectMemoryRasConfiguration *string                              `json:"VpSelectMemoryRasConfiguration,omitempty"`
	ComputeBlade                   *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit                *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo            *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice               *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _BiosVfSelectMemoryRasConfigurationAllOf BiosVfSelectMemoryRasConfigurationAllOf

// NewBiosVfSelectMemoryRasConfigurationAllOf instantiates a new BiosVfSelectMemoryRasConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosVfSelectMemoryRasConfigurationAllOf(classId string, objectType string) *BiosVfSelectMemoryRasConfigurationAllOf {
	this := BiosVfSelectMemoryRasConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosVfSelectMemoryRasConfigurationAllOfWithDefaults instantiates a new BiosVfSelectMemoryRasConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosVfSelectMemoryRasConfigurationAllOfWithDefaults() *BiosVfSelectMemoryRasConfigurationAllOf {
	this := BiosVfSelectMemoryRasConfigurationAllOf{}
	var classId string = "bios.VfSelectMemoryRasConfiguration"
	this.ClassId = classId
	var objectType string = "bios.VfSelectMemoryRasConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetVpSelectMemoryRasConfiguration returns the VpSelectMemoryRasConfiguration field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetVpSelectMemoryRasConfiguration() string {
	if o == nil || o.VpSelectMemoryRasConfiguration == nil {
		var ret string
		return ret
	}
	return *o.VpSelectMemoryRasConfiguration
}

// GetVpSelectMemoryRasConfigurationOk returns a tuple with the VpSelectMemoryRasConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetVpSelectMemoryRasConfigurationOk() (*string, bool) {
	if o == nil || o.VpSelectMemoryRasConfiguration == nil {
		return nil, false
	}
	return o.VpSelectMemoryRasConfiguration, true
}

// HasVpSelectMemoryRasConfiguration returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasVpSelectMemoryRasConfiguration() bool {
	if o != nil && o.VpSelectMemoryRasConfiguration != nil {
		return true
	}

	return false
}

// SetVpSelectMemoryRasConfiguration gets a reference to the given string and assigns it to the VpSelectMemoryRasConfiguration field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetVpSelectMemoryRasConfiguration(v string) {
	o.VpSelectMemoryRasConfiguration = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosVfSelectMemoryRasConfigurationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosVfSelectMemoryRasConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.VpSelectMemoryRasConfiguration != nil {
		toSerialize["VpSelectMemoryRasConfiguration"] = o.VpSelectMemoryRasConfiguration
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosVfSelectMemoryRasConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBiosVfSelectMemoryRasConfigurationAllOf := _BiosVfSelectMemoryRasConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varBiosVfSelectMemoryRasConfigurationAllOf); err == nil {
		*o = BiosVfSelectMemoryRasConfigurationAllOf(varBiosVfSelectMemoryRasConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "VpSelectMemoryRasConfiguration")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBiosVfSelectMemoryRasConfigurationAllOf struct {
	value *BiosVfSelectMemoryRasConfigurationAllOf
	isSet bool
}

func (v NullableBiosVfSelectMemoryRasConfigurationAllOf) Get() *BiosVfSelectMemoryRasConfigurationAllOf {
	return v.value
}

func (v *NullableBiosVfSelectMemoryRasConfigurationAllOf) Set(val *BiosVfSelectMemoryRasConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosVfSelectMemoryRasConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosVfSelectMemoryRasConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosVfSelectMemoryRasConfigurationAllOf(val *BiosVfSelectMemoryRasConfigurationAllOf) *NullableBiosVfSelectMemoryRasConfigurationAllOf {
	return &NullableBiosVfSelectMemoryRasConfigurationAllOf{value: val, isSet: true}
}

func (v NullableBiosVfSelectMemoryRasConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosVfSelectMemoryRasConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
