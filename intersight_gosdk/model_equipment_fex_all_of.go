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

// EquipmentFexAllOf Definition of the list of properties defined in 'equipment.Fex', excluding properties defined in parent classes.
type EquipmentFexAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Discovery state of IO card or fabric extender.
	DiscoveryState *string `json:"DiscoveryState,omitempty"`
	// An array of relationships to equipmentFan resources.
	Fans                []EquipmentFanRelationship       `json:"Fans,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to equipmentIoCard resources.
	Ioms           []EquipmentIoCardRelationship    `json:"Ioms,omitempty"`
	LocatorLed     *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	NetworkElement *NetworkElementRelationship      `json:"NetworkElement,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus                 []EquipmentPsuRelationship           `json:"Psus,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentFexAllOf EquipmentFexAllOf

// NewEquipmentFexAllOf instantiates a new EquipmentFexAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFexAllOf(classId string, objectType string) *EquipmentFexAllOf {
	this := EquipmentFexAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentFexAllOfWithDefaults instantiates a new EquipmentFexAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFexAllOfWithDefaults() *EquipmentFexAllOf {
	this := EquipmentFexAllOf{}
	var classId string = "equipment.Fex"
	this.ClassId = classId
	var objectType string = "equipment.Fex"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentFexAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentFexAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentFexAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentFexAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiscoveryState returns the DiscoveryState field value if set, zero value otherwise.
func (o *EquipmentFexAllOf) GetDiscoveryState() string {
	if o == nil || o.DiscoveryState == nil {
		var ret string
		return ret
	}
	return *o.DiscoveryState
}

// GetDiscoveryStateOk returns a tuple with the DiscoveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetDiscoveryStateOk() (*string, bool) {
	if o == nil || o.DiscoveryState == nil {
		return nil, false
	}
	return o.DiscoveryState, true
}

// HasDiscoveryState returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasDiscoveryState() bool {
	if o != nil && o.DiscoveryState != nil {
		return true
	}

	return false
}

// SetDiscoveryState gets a reference to the given string and assigns it to the DiscoveryState field.
func (o *EquipmentFexAllOf) SetDiscoveryState(v string) {
	o.DiscoveryState = &v
}

// GetFans returns the Fans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFexAllOf) GetFans() []EquipmentFanRelationship {
	if o == nil {
		var ret []EquipmentFanRelationship
		return ret
	}
	return o.Fans
}

// GetFansOk returns a tuple with the Fans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFexAllOf) GetFansOk() (*[]EquipmentFanRelationship, bool) {
	if o == nil || o.Fans == nil {
		return nil, false
	}
	return &o.Fans, true
}

// HasFans returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasFans() bool {
	if o != nil && o.Fans != nil {
		return true
	}

	return false
}

// SetFans gets a reference to the given []EquipmentFanRelationship and assigns it to the Fans field.
func (o *EquipmentFexAllOf) SetFans(v []EquipmentFanRelationship) {
	o.Fans = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentFexAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentFexAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetIoms returns the Ioms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFexAllOf) GetIoms() []EquipmentIoCardRelationship {
	if o == nil {
		var ret []EquipmentIoCardRelationship
		return ret
	}
	return o.Ioms
}

// GetIomsOk returns a tuple with the Ioms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFexAllOf) GetIomsOk() (*[]EquipmentIoCardRelationship, bool) {
	if o == nil || o.Ioms == nil {
		return nil, false
	}
	return &o.Ioms, true
}

// HasIoms returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasIoms() bool {
	if o != nil && o.Ioms != nil {
		return true
	}

	return false
}

// SetIoms gets a reference to the given []EquipmentIoCardRelationship and assigns it to the Ioms field.
func (o *EquipmentFexAllOf) SetIoms(v []EquipmentIoCardRelationship) {
	o.Ioms = v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *EquipmentFexAllOf) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *EquipmentFexAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentFexAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentFexAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetPsus returns the Psus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFexAllOf) GetPsus() []EquipmentPsuRelationship {
	if o == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFexAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool) {
	if o == nil || o.Psus == nil {
		return nil, false
	}
	return &o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasPsus() bool {
	if o != nil && o.Psus != nil {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *EquipmentFexAllOf) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentFexAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentFexAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentFexAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentFexAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiscoveryState != nil {
		toSerialize["DiscoveryState"] = o.DiscoveryState
	}
	if o.Fans != nil {
		toSerialize["Fans"] = o.Fans
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.Ioms != nil {
		toSerialize["Ioms"] = o.Ioms
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentFexAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentFexAllOf := _EquipmentFexAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentFexAllOf); err == nil {
		*o = EquipmentFexAllOf(varEquipmentFexAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiscoveryState")
		delete(additionalProperties, "Fans")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Ioms")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "Psus")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentFexAllOf struct {
	value *EquipmentFexAllOf
	isSet bool
}

func (v NullableEquipmentFexAllOf) Get() *EquipmentFexAllOf {
	return v.value
}

func (v *NullableEquipmentFexAllOf) Set(val *EquipmentFexAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFexAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFexAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFexAllOf(val *EquipmentFexAllOf) *NullableEquipmentFexAllOf {
	return &NullableEquipmentFexAllOf{value: val, isSet: true}
}

func (v NullableEquipmentFexAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFexAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
