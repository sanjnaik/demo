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

// VirtualizationBaseHypervisorManagerAllOf Definition of the list of properties defined in 'virtualization.BaseHypervisorManager', excluding properties defined in parent classes.
type VirtualizationBaseHypervisorManagerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The build number of the Hypervisor Manger (e.g., 4541947, 6.3.9600.18692). The build number may indicate some feature support that applications might rely on. The build number may not always be an integer.
	Build *string `json:"Build,omitempty"`
	// Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608.
	Identity *string `json:"Identity,omitempty"`
	// The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor.
	Name *string `json:"Name,omitempty"`
	// Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947).
	Version              *string                              `json:"Version,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseHypervisorManagerAllOf VirtualizationBaseHypervisorManagerAllOf

// NewVirtualizationBaseHypervisorManagerAllOf instantiates a new VirtualizationBaseHypervisorManagerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseHypervisorManagerAllOf(classId string, objectType string) *VirtualizationBaseHypervisorManagerAllOf {
	this := VirtualizationBaseHypervisorManagerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseHypervisorManagerAllOfWithDefaults instantiates a new VirtualizationBaseHypervisorManagerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseHypervisorManagerAllOfWithDefaults() *VirtualizationBaseHypervisorManagerAllOf {
	this := VirtualizationBaseHypervisorManagerAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseHypervisorManagerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseHypervisorManagerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseHypervisorManagerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseHypervisorManagerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetBuild() string {
	if o == nil || o.Build == nil {
		var ret string
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetBuildOk() (*string, bool) {
	if o == nil || o.Build == nil {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) HasBuild() bool {
	if o != nil && o.Build != nil {
		return true
	}

	return false
}

// SetBuild gets a reference to the given string and assigns it to the Build field.
func (o *VirtualizationBaseHypervisorManagerAllOf) SetBuild(v string) {
	o.Build = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseHypervisorManagerAllOf) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseHypervisorManagerAllOf) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VirtualizationBaseHypervisorManagerAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationBaseHypervisorManagerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationBaseHypervisorManagerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o VirtualizationBaseHypervisorManagerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Build != nil {
		toSerialize["Build"] = o.Build
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseHypervisorManagerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationBaseHypervisorManagerAllOf := _VirtualizationBaseHypervisorManagerAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationBaseHypervisorManagerAllOf); err == nil {
		*o = VirtualizationBaseHypervisorManagerAllOf(varVirtualizationBaseHypervisorManagerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Build")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationBaseHypervisorManagerAllOf struct {
	value *VirtualizationBaseHypervisorManagerAllOf
	isSet bool
}

func (v NullableVirtualizationBaseHypervisorManagerAllOf) Get() *VirtualizationBaseHypervisorManagerAllOf {
	return v.value
}

func (v *NullableVirtualizationBaseHypervisorManagerAllOf) Set(val *VirtualizationBaseHypervisorManagerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseHypervisorManagerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseHypervisorManagerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseHypervisorManagerAllOf(val *VirtualizationBaseHypervisorManagerAllOf) *NullableVirtualizationBaseHypervisorManagerAllOf {
	return &NullableVirtualizationBaseHypervisorManagerAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationBaseHypervisorManagerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseHypervisorManagerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
