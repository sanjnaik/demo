/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-12T05:47:20Z.
 *
 * API version: 1.0.9-4240
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationBaseVirtualNetworkInterfaceCardAllOf Definition of the list of properties defined in 'virtualization.BaseVirtualNetworkInterfaceCard', excluding properties defined in parent classes.
type VirtualizationBaseVirtualNetworkInterfaceCardAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Type or model of the virtual network interface card. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
	AdapterType *string `json:"AdapterType,omitempty"`
	// MAC address assigned to the virtual network interface card.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the virtual network interface card.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualNetworkInterfaceCardAllOf VirtualizationBaseVirtualNetworkInterfaceCardAllOf

// NewVirtualizationBaseVirtualNetworkInterfaceCardAllOf instantiates a new VirtualizationBaseVirtualNetworkInterfaceCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualNetworkInterfaceCardAllOf(classId string, objectType string) *VirtualizationBaseVirtualNetworkInterfaceCardAllOf {
	this := VirtualizationBaseVirtualNetworkInterfaceCardAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// NewVirtualizationBaseVirtualNetworkInterfaceCardAllOfWithDefaults instantiates a new VirtualizationBaseVirtualNetworkInterfaceCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualNetworkInterfaceCardAllOfWithDefaults() *VirtualizationBaseVirtualNetworkInterfaceCardAllOf {
	this := VirtualizationBaseVirtualNetworkInterfaceCardAllOf{}
	var classId string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ObjectType = objectType
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterType returns the AdapterType field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetAdapterType() string {
	if o == nil || o.AdapterType == nil {
		var ret string
		return ret
	}
	return *o.AdapterType
}

// GetAdapterTypeOk returns a tuple with the AdapterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetAdapterTypeOk() (*string, bool) {
	if o == nil || o.AdapterType == nil {
		return nil, false
	}
	return o.AdapterType, true
}

// HasAdapterType returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasAdapterType() bool {
	if o != nil && o.AdapterType != nil {
		return true
	}

	return false
}

// SetAdapterType gets a reference to the given string and assigns it to the AdapterType field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetAdapterType(v string) {
	o.AdapterType = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationBaseVirtualNetworkInterfaceCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterType != nil {
		toSerialize["AdapterType"] = o.AdapterType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseVirtualNetworkInterfaceCardAllOf := _VirtualizationBaseVirtualNetworkInterfaceCardAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualNetworkInterfaceCardAllOf); err == nil {
		*o = VirtualizationBaseVirtualNetworkInterfaceCardAllOf(varVirtualizationBaseVirtualNetworkInterfaceCardAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf struct {
	value *VirtualizationBaseVirtualNetworkInterfaceCardAllOf
	isSet bool
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) Get() *VirtualizationBaseVirtualNetworkInterfaceCardAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) Set(val *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf(val *VirtualizationBaseVirtualNetworkInterfaceCardAllOf) *NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf {
	return &NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualNetworkInterfaceCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}