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
	"reflect"
	"strings"
)

// HyperflexHxapVirtualMachineNetworkInterface A HyperFlex Application Platform virtual network interface entity that a virtual machine is attached to.
type HyperflexHxapVirtualMachineNetworkInterface struct {
	VirtualizationBaseVirtualNetworkInterfaceCard
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operating system assigned name for network interface.
	InterfaceName *string  `json:"InterfaceName,omitempty"`
	IpAddress     []string `json:"IpAddress,omitempty"`
	// Primary IP address of the network interface.
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty"`
	// Current status of virtual network interface status. * `Up` - Virtual network interface is up and running. * `Down` - Virtual network interface is down and not running.
	Status *string `json:"Status,omitempty"`
	// A reference to the virtual machine where this network object is attached to.
	VirtualMachineName   *string                                  `json:"VirtualMachineName,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship        `json:"Cluster,omitempty"`
	Network              *HyperflexHxapNetworkRelationship        `json:"Network,omitempty"`
	VirtualMachine       *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapVirtualMachineNetworkInterface HyperflexHxapVirtualMachineNetworkInterface

// NewHyperflexHxapVirtualMachineNetworkInterface instantiates a new HyperflexHxapVirtualMachineNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapVirtualMachineNetworkInterface(classId string, objectType string) *HyperflexHxapVirtualMachineNetworkInterface {
	this := HyperflexHxapVirtualMachineNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Up"
	this.Status = &status
	return &this
}

// NewHyperflexHxapVirtualMachineNetworkInterfaceWithDefaults instantiates a new HyperflexHxapVirtualMachineNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapVirtualMachineNetworkInterfaceWithDefaults() *HyperflexHxapVirtualMachineNetworkInterface {
	this := HyperflexHxapVirtualMachineNetworkInterface{}
	var classId string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ObjectType = objectType
	var status string = "Up"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetPrimaryIpAddress returns the PrimaryIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetPrimaryIpAddress() string {
	if o == nil || o.PrimaryIpAddress == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIpAddress
}

// GetPrimaryIpAddressOk returns a tuple with the PrimaryIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetPrimaryIpAddressOk() (*string, bool) {
	if o == nil || o.PrimaryIpAddress == nil {
		return nil, false
	}
	return o.PrimaryIpAddress, true
}

// HasPrimaryIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasPrimaryIpAddress() bool {
	if o != nil && o.PrimaryIpAddress != nil {
		return true
	}

	return false
}

// SetPrimaryIpAddress gets a reference to the given string and assigns it to the PrimaryIpAddress field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetPrimaryIpAddress(v string) {
	o.PrimaryIpAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetStatus(v string) {
	o.Status = &v
}

// GetVirtualMachineName returns the VirtualMachineName field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineName() string {
	if o == nil || o.VirtualMachineName == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachineName
}

// GetVirtualMachineNameOk returns a tuple with the VirtualMachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineNameOk() (*string, bool) {
	if o == nil || o.VirtualMachineName == nil {
		return nil, false
	}
	return o.VirtualMachineName, true
}

// HasVirtualMachineName returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasVirtualMachineName() bool {
	if o != nil && o.VirtualMachineName != nil {
		return true
	}

	return false
}

// SetVirtualMachineName gets a reference to the given string and assigns it to the VirtualMachineName field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetVirtualMachineName(v string) {
	o.VirtualMachineName = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetNetwork() HyperflexHxapNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret HyperflexHxapNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given HyperflexHxapNetworkRelationship and assigns it to the Network field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetNetwork(v HyperflexHxapNetworkRelationship) {
	o.Network = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret HyperflexHxapVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterface) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given HyperflexHxapVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *HyperflexHxapVirtualMachineNetworkInterface) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o HyperflexHxapVirtualMachineNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualNetworkInterfaceCard, errVirtualizationBaseVirtualNetworkInterfaceCard := json.Marshal(o.VirtualizationBaseVirtualNetworkInterfaceCard)
	if errVirtualizationBaseVirtualNetworkInterfaceCard != nil {
		return []byte{}, errVirtualizationBaseVirtualNetworkInterfaceCard
	}
	errVirtualizationBaseVirtualNetworkInterfaceCard = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualNetworkInterfaceCard), &toSerialize)
	if errVirtualizationBaseVirtualNetworkInterfaceCard != nil {
		return []byte{}, errVirtualizationBaseVirtualNetworkInterfaceCard
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.PrimaryIpAddress != nil {
		toSerialize["PrimaryIpAddress"] = o.PrimaryIpAddress
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VirtualMachineName != nil {
		toSerialize["VirtualMachineName"] = o.VirtualMachineName
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapVirtualMachineNetworkInterface) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Operating system assigned name for network interface.
		InterfaceName *string  `json:"InterfaceName,omitempty"`
		IpAddress     []string `json:"IpAddress,omitempty"`
		// Primary IP address of the network interface.
		PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty"`
		// Current status of virtual network interface status. * `Up` - Virtual network interface is up and running. * `Down` - Virtual network interface is down and not running.
		Status *string `json:"Status,omitempty"`
		// A reference to the virtual machine where this network object is attached to.
		VirtualMachineName *string                                  `json:"VirtualMachineName,omitempty"`
		Cluster            *HyperflexHxapClusterRelationship        `json:"Cluster,omitempty"`
		Network            *HyperflexHxapNetworkRelationship        `json:"Network,omitempty"`
		VirtualMachine     *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct := HyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapVirtualMachineNetworkInterface := _HyperflexHxapVirtualMachineNetworkInterface{}
		varHyperflexHxapVirtualMachineNetworkInterface.ClassId = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varHyperflexHxapVirtualMachineNetworkInterface.ObjectType = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapVirtualMachineNetworkInterface.InterfaceName = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.InterfaceName
		varHyperflexHxapVirtualMachineNetworkInterface.IpAddress = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.IpAddress
		varHyperflexHxapVirtualMachineNetworkInterface.PrimaryIpAddress = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.PrimaryIpAddress
		varHyperflexHxapVirtualMachineNetworkInterface.Status = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.Status
		varHyperflexHxapVirtualMachineNetworkInterface.VirtualMachineName = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.VirtualMachineName
		varHyperflexHxapVirtualMachineNetworkInterface.Cluster = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.Cluster
		varHyperflexHxapVirtualMachineNetworkInterface.Network = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.Network
		varHyperflexHxapVirtualMachineNetworkInterface.VirtualMachine = varHyperflexHxapVirtualMachineNetworkInterfaceWithoutEmbeddedStruct.VirtualMachine
		*o = HyperflexHxapVirtualMachineNetworkInterface(varHyperflexHxapVirtualMachineNetworkInterface)
	} else {
		return err
	}

	varHyperflexHxapVirtualMachineNetworkInterface := _HyperflexHxapVirtualMachineNetworkInterface{}

	err = json.Unmarshal(bytes, &varHyperflexHxapVirtualMachineNetworkInterface)
	if err == nil {
		o.VirtualizationBaseVirtualNetworkInterfaceCard = varHyperflexHxapVirtualMachineNetworkInterface.VirtualizationBaseVirtualNetworkInterfaceCard
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "PrimaryIpAddress")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VirtualMachineName")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Network")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualNetworkInterfaceCard := reflect.ValueOf(o.VirtualizationBaseVirtualNetworkInterfaceCard)
		for i := 0; i < reflectVirtualizationBaseVirtualNetworkInterfaceCard.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualNetworkInterfaceCard.Type().Field(i)

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

type NullableHyperflexHxapVirtualMachineNetworkInterface struct {
	value *HyperflexHxapVirtualMachineNetworkInterface
	isSet bool
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterface) Get() *HyperflexHxapVirtualMachineNetworkInterface {
	return v.value
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterface) Set(val *HyperflexHxapVirtualMachineNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapVirtualMachineNetworkInterface(val *HyperflexHxapVirtualMachineNetworkInterface) *NullableHyperflexHxapVirtualMachineNetworkInterface {
	return &NullableHyperflexHxapVirtualMachineNetworkInterface{value: val, isSet: true}
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
