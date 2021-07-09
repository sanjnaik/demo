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

// EquipmentIdentityAllOf Definition of the list of properties defined in 'equipment.Identity', excluding properties defined in parent classes.
type EquipmentIdentityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Updated by UI/API to trigger specific chassis action type. * `None` - No operation value for maintenance actions on an equipment. * `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight. * `Recommission` - Recommission the equipment. * `Reack` - Reacknowledge the equipment and discover it again. * `Remove` - Remove the equipment permanently from Intersight management. * `Replace` - Replace the equipment with the other one.
	AdminAction *string `json:"AdminAction,omitempty"`
	// The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	AdminActionState *string `json:"AdminActionState,omitempty"`
	// Numeric Identifier assigned by the management system to the equipment.
	Identifier *int64 `json:"Identifier,omitempty"`
	// The equipment's lifecycle status. * `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * `Active` - Default Lifecycle State for a physical entity. * `Decommissioned` - Decommission Lifecycle state. * `DecommissionInProgress` - Decommission Inprogress Lifecycle state. * `RecommissionInProgress` - Recommission Inprogress Lifecycle state. * `OperationFailed` - Failed Operation Lifecycle state. * `ReackInProgress` - ReackInProgress Lifecycle state. * `RemoveInProgress` - RemoveInProgress Lifecycle state. * `Discovered` - Discovered Lifecycle state. * `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state. * `DiscoveryFailed` - DiscoveryFailed Lifecycle state. * `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity. * `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity. * `Inactive` - Inactive Lifecycle state. * `ReplaceInProgress` - ReplaceInProgress Lifecycle state.
	Lifecycle *string `json:"Lifecycle,omitempty"`
	// The vendor provided model name for the equipment.
	Model *string `json:"Model,omitempty"`
	// The serial number of the equipment.
	Serial *string `json:"Serial,omitempty"`
	// The manufacturer of the equipment.
	Vendor               *string                              `json:"Vendor,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIdentityAllOf EquipmentIdentityAllOf

// NewEquipmentIdentityAllOf instantiates a new EquipmentIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIdentityAllOf(classId string, objectType string) *EquipmentIdentityAllOf {
	this := EquipmentIdentityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var adminActionState string = "None"
	this.AdminActionState = &adminActionState
	var lifecycle string = "None"
	this.Lifecycle = &lifecycle
	return &this
}

// NewEquipmentIdentityAllOfWithDefaults instantiates a new EquipmentIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIdentityAllOfWithDefaults() *EquipmentIdentityAllOf {
	this := EquipmentIdentityAllOf{}
	var adminAction string = "None"
	this.AdminAction = &adminAction
	var adminActionState string = "None"
	this.AdminActionState = &adminActionState
	var lifecycle string = "None"
	this.Lifecycle = &lifecycle
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIdentityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIdentityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIdentityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIdentityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *EquipmentIdentityAllOf) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetAdminActionState returns the AdminActionState field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetAdminActionState() string {
	if o == nil || o.AdminActionState == nil {
		var ret string
		return ret
	}
	return *o.AdminActionState
}

// GetAdminActionStateOk returns a tuple with the AdminActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetAdminActionStateOk() (*string, bool) {
	if o == nil || o.AdminActionState == nil {
		return nil, false
	}
	return o.AdminActionState, true
}

// HasAdminActionState returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasAdminActionState() bool {
	if o != nil && o.AdminActionState != nil {
		return true
	}

	return false
}

// SetAdminActionState gets a reference to the given string and assigns it to the AdminActionState field.
func (o *EquipmentIdentityAllOf) SetAdminActionState(v string) {
	o.AdminActionState = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetIdentifier() int64 {
	if o == nil || o.Identifier == nil {
		var ret int64
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetIdentifierOk() (*int64, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given int64 and assigns it to the Identifier field.
func (o *EquipmentIdentityAllOf) SetIdentifier(v int64) {
	o.Identifier = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EquipmentIdentityAllOf) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentIdentityAllOf) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentIdentityAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentIdentityAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIdentityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIdentityAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIdentityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.AdminActionState != nil {
		toSerialize["AdminActionState"] = o.AdminActionState
	}
	if o.Identifier != nil {
		toSerialize["Identifier"] = o.Identifier
	}
	if o.Lifecycle != nil {
		toSerialize["Lifecycle"] = o.Lifecycle
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIdentityAllOf := _EquipmentIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIdentityAllOf); err == nil {
		*o = EquipmentIdentityAllOf(varEquipmentIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "AdminActionState")
		delete(additionalProperties, "Identifier")
		delete(additionalProperties, "Lifecycle")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIdentityAllOf struct {
	value *EquipmentIdentityAllOf
	isSet bool
}

func (v NullableEquipmentIdentityAllOf) Get() *EquipmentIdentityAllOf {
	return v.value
}

func (v *NullableEquipmentIdentityAllOf) Set(val *EquipmentIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIdentityAllOf(val *EquipmentIdentityAllOf) *NullableEquipmentIdentityAllOf {
	return &NullableEquipmentIdentityAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
