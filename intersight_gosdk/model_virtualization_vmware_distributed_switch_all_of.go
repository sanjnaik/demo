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
)

// VirtualizationVmwareDistributedSwitchAllOf Definition of the list of properties defined in 'virtualization.VmwareDistributedSwitch', excluding properties defined in parent classes.
type VirtualizationVmwareDistributedSwitchAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switch description (user provided), if any.
	Description *string `json:"Description,omitempty"`
	// Maximum number of ports allowed on this distributed virtual switch.
	MaxPort *int64 `json:"MaxPort,omitempty"`
	// Maximum transmission unit configured on a distributed virtual switch.
	Mtu                   *int64                                         `json:"Mtu,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// The total number of hosts attached to the distributed virtual switch.
	NumHosts *int64 `json:"NumHosts,omitempty"`
	// The total number of distributed networks in the distributed virtual switch.
	NumNetworks *int64 `json:"NumNetworks,omitempty"`
	// Number of stand-alone ports in use.
	NumStandAlonePorts *int64 `json:"NumStandAlonePorts,omitempty"`
	// Number of uplinks configured in this distributed virtual switch.
	NumUplinks     *int64                                `json:"NumUplinks,omitempty"`
	SwitchCapacity NullableVirtualizationStorageCapacity `json:"SwitchCapacity,omitempty"`
	// Universally Unique Id of this distributed virtual switch.
	Uuid *string `json:"Uuid,omitempty"`
	// The running config's version details are represented.
	Version    *string                                     `json:"Version,omitempty"`
	Datacenter *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	// An array of relationships to virtualizationVmwareHost resources.
	Hosts                []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDistributedSwitchAllOf VirtualizationVmwareDistributedSwitchAllOf

// NewVirtualizationVmwareDistributedSwitchAllOf instantiates a new VirtualizationVmwareDistributedSwitchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDistributedSwitchAllOf(classId string, objectType string) *VirtualizationVmwareDistributedSwitchAllOf {
	this := VirtualizationVmwareDistributedSwitchAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareDistributedSwitchAllOfWithDefaults instantiates a new VirtualizationVmwareDistributedSwitchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDistributedSwitchAllOfWithDefaults() *VirtualizationVmwareDistributedSwitchAllOf {
	this := VirtualizationVmwareDistributedSwitchAllOf{}
	var classId string = "virtualization.VmwareDistributedSwitch"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDistributedSwitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetMaxPort returns the MaxPort field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMaxPort() int64 {
	if o == nil || o.MaxPort == nil {
		var ret int64
		return ret
	}
	return *o.MaxPort
}

// GetMaxPortOk returns a tuple with the MaxPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMaxPortOk() (*int64, bool) {
	if o == nil || o.MaxPort == nil {
		return nil, false
	}
	return o.MaxPort, true
}

// HasMaxPort returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasMaxPort() bool {
	if o != nil && o.MaxPort != nil {
		return true
	}

	return false
}

// SetMaxPort gets a reference to the given int64 and assigns it to the MaxPort field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetMaxPort(v int64) {
	o.MaxPort = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || o.NicTeamingAndFailover.Get() == nil {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareDistributedSwitchAllOf) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetNumHosts returns the NumHosts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumHosts() int64 {
	if o == nil || o.NumHosts == nil {
		var ret int64
		return ret
	}
	return *o.NumHosts
}

// GetNumHostsOk returns a tuple with the NumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumHostsOk() (*int64, bool) {
	if o == nil || o.NumHosts == nil {
		return nil, false
	}
	return o.NumHosts, true
}

// HasNumHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumHosts() bool {
	if o != nil && o.NumHosts != nil {
		return true
	}

	return false
}

// SetNumHosts gets a reference to the given int64 and assigns it to the NumHosts field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumHosts(v int64) {
	o.NumHosts = &v
}

// GetNumNetworks returns the NumNetworks field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumNetworks() int64 {
	if o == nil || o.NumNetworks == nil {
		var ret int64
		return ret
	}
	return *o.NumNetworks
}

// GetNumNetworksOk returns a tuple with the NumNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumNetworksOk() (*int64, bool) {
	if o == nil || o.NumNetworks == nil {
		return nil, false
	}
	return o.NumNetworks, true
}

// HasNumNetworks returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumNetworks() bool {
	if o != nil && o.NumNetworks != nil {
		return true
	}

	return false
}

// SetNumNetworks gets a reference to the given int64 and assigns it to the NumNetworks field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumNetworks(v int64) {
	o.NumNetworks = &v
}

// GetNumStandAlonePorts returns the NumStandAlonePorts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumStandAlonePorts() int64 {
	if o == nil || o.NumStandAlonePorts == nil {
		var ret int64
		return ret
	}
	return *o.NumStandAlonePorts
}

// GetNumStandAlonePortsOk returns a tuple with the NumStandAlonePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumStandAlonePortsOk() (*int64, bool) {
	if o == nil || o.NumStandAlonePorts == nil {
		return nil, false
	}
	return o.NumStandAlonePorts, true
}

// HasNumStandAlonePorts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumStandAlonePorts() bool {
	if o != nil && o.NumStandAlonePorts != nil {
		return true
	}

	return false
}

// SetNumStandAlonePorts gets a reference to the given int64 and assigns it to the NumStandAlonePorts field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumStandAlonePorts(v int64) {
	o.NumStandAlonePorts = &v
}

// GetNumUplinks returns the NumUplinks field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumUplinks() int64 {
	if o == nil || o.NumUplinks == nil {
		var ret int64
		return ret
	}
	return *o.NumUplinks
}

// GetNumUplinksOk returns a tuple with the NumUplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetNumUplinksOk() (*int64, bool) {
	if o == nil || o.NumUplinks == nil {
		return nil, false
	}
	return o.NumUplinks, true
}

// HasNumUplinks returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasNumUplinks() bool {
	if o != nil && o.NumUplinks != nil {
		return true
	}

	return false
}

// SetNumUplinks gets a reference to the given int64 and assigns it to the NumUplinks field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetNumUplinks(v int64) {
	o.NumUplinks = &v
}

// GetSwitchCapacity returns the SwitchCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetSwitchCapacity() VirtualizationStorageCapacity {
	if o == nil || o.SwitchCapacity.Get() == nil {
		var ret VirtualizationStorageCapacity
		return ret
	}
	return *o.SwitchCapacity.Get()
}

// GetSwitchCapacityOk returns a tuple with the SwitchCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetSwitchCapacityOk() (*VirtualizationStorageCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.SwitchCapacity.Get(), o.SwitchCapacity.IsSet()
}

// HasSwitchCapacity returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasSwitchCapacity() bool {
	if o != nil && o.SwitchCapacity.IsSet() {
		return true
	}

	return false
}

// SetSwitchCapacity gets a reference to the given NullableVirtualizationStorageCapacity and assigns it to the SwitchCapacity field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetSwitchCapacity(v VirtualizationStorageCapacity) {
	o.SwitchCapacity.Set(&v)
}

// SetSwitchCapacityNil sets the value for SwitchCapacity to be an explicit nil
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetSwitchCapacityNil() {
	o.SwitchCapacity.Set(nil)
}

// UnsetSwitchCapacity ensures that no value is present for SwitchCapacity, not even an explicit nil
func (o *VirtualizationVmwareDistributedSwitchAllOf) UnsetSwitchCapacity() {
	o.SwitchCapacity.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetHosts() []VirtualizationVmwareHostRelationship {
	if o == nil {
		var ret []VirtualizationVmwareHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedSwitchAllOf) GetHostsOk() (*[]VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedSwitchAllOf) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VirtualizationVmwareHostRelationship and assigns it to the Hosts field.
func (o *VirtualizationVmwareDistributedSwitchAllOf) SetHosts(v []VirtualizationVmwareHostRelationship) {
	o.Hosts = v
}

func (o VirtualizationVmwareDistributedSwitchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.MaxPort != nil {
		toSerialize["MaxPort"] = o.MaxPort
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if o.NumHosts != nil {
		toSerialize["NumHosts"] = o.NumHosts
	}
	if o.NumNetworks != nil {
		toSerialize["NumNetworks"] = o.NumNetworks
	}
	if o.NumStandAlonePorts != nil {
		toSerialize["NumStandAlonePorts"] = o.NumStandAlonePorts
	}
	if o.NumUplinks != nil {
		toSerialize["NumUplinks"] = o.NumUplinks
	}
	if o.SwitchCapacity.IsSet() {
		toSerialize["SwitchCapacity"] = o.SwitchCapacity.Get()
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDistributedSwitchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareDistributedSwitchAllOf := _VirtualizationVmwareDistributedSwitchAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareDistributedSwitchAllOf); err == nil {
		*o = VirtualizationVmwareDistributedSwitchAllOf(varVirtualizationVmwareDistributedSwitchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "MaxPort")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "NumHosts")
		delete(additionalProperties, "NumNetworks")
		delete(additionalProperties, "NumStandAlonePorts")
		delete(additionalProperties, "NumUplinks")
		delete(additionalProperties, "SwitchCapacity")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "Hosts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareDistributedSwitchAllOf struct {
	value *VirtualizationVmwareDistributedSwitchAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareDistributedSwitchAllOf) Get() *VirtualizationVmwareDistributedSwitchAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareDistributedSwitchAllOf) Set(val *VirtualizationVmwareDistributedSwitchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDistributedSwitchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDistributedSwitchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDistributedSwitchAllOf(val *VirtualizationVmwareDistributedSwitchAllOf) *NullableVirtualizationVmwareDistributedSwitchAllOf {
	return &NullableVirtualizationVmwareDistributedSwitchAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDistributedSwitchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDistributedSwitchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
