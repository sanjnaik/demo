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

// VirtualizationCloudInitConfig It is a configuration for early initialization of virtual machine instances. It enables automatic configuration of virtual machine instances as it boots.
type VirtualizationCloudInitConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual machine cloud init configuration type. * `` - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected. * `NoCloudSource` - Allows the user to provide user-data to the instance without running a network service. * `CloudConfigDrive` - Allows the user to provide user-data and network-data from cloud.
	ConfigType *string `json:"ConfigType,omitempty"`
	// Network configuration data for a virtual machine.
	NetworkData *string `json:"NetworkData,omitempty"`
	// Set to true, if the cloud init network data is in base64 format.
	NetworkDataBase64Encoded *bool `json:"NetworkDataBase64Encoded,omitempty"`
	// User configuration data for a virtual machine such as adding user, group etc.
	UserData *string `json:"UserData,omitempty"`
	// Set to true, if the cloud init user data is in base64 format.
	UserDataBase64Encoded *bool `json:"UserDataBase64Encoded,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VirtualizationCloudInitConfig VirtualizationCloudInitConfig

// NewVirtualizationCloudInitConfig instantiates a new VirtualizationCloudInitConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCloudInitConfig(classId string, objectType string) *VirtualizationCloudInitConfig {
	this := VirtualizationCloudInitConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configType string = ""
	this.ConfigType = &configType
	return &this
}

// NewVirtualizationCloudInitConfigWithDefaults instantiates a new VirtualizationCloudInitConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCloudInitConfigWithDefaults() *VirtualizationCloudInitConfig {
	this := VirtualizationCloudInitConfig{}
	var classId string = "virtualization.CloudInitConfig"
	this.ClassId = classId
	var objectType string = "virtualization.CloudInitConfig"
	this.ObjectType = objectType
	var configType string = ""
	this.ConfigType = &configType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCloudInitConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCloudInitConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCloudInitConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCloudInitConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigType returns the ConfigType field value if set, zero value otherwise.
func (o *VirtualizationCloudInitConfig) GetConfigType() string {
	if o == nil || o.ConfigType == nil {
		var ret string
		return ret
	}
	return *o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetConfigTypeOk() (*string, bool) {
	if o == nil || o.ConfigType == nil {
		return nil, false
	}
	return o.ConfigType, true
}

// HasConfigType returns a boolean if a field has been set.
func (o *VirtualizationCloudInitConfig) HasConfigType() bool {
	if o != nil && o.ConfigType != nil {
		return true
	}

	return false
}

// SetConfigType gets a reference to the given string and assigns it to the ConfigType field.
func (o *VirtualizationCloudInitConfig) SetConfigType(v string) {
	o.ConfigType = &v
}

// GetNetworkData returns the NetworkData field value if set, zero value otherwise.
func (o *VirtualizationCloudInitConfig) GetNetworkData() string {
	if o == nil || o.NetworkData == nil {
		var ret string
		return ret
	}
	return *o.NetworkData
}

// GetNetworkDataOk returns a tuple with the NetworkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetNetworkDataOk() (*string, bool) {
	if o == nil || o.NetworkData == nil {
		return nil, false
	}
	return o.NetworkData, true
}

// HasNetworkData returns a boolean if a field has been set.
func (o *VirtualizationCloudInitConfig) HasNetworkData() bool {
	if o != nil && o.NetworkData != nil {
		return true
	}

	return false
}

// SetNetworkData gets a reference to the given string and assigns it to the NetworkData field.
func (o *VirtualizationCloudInitConfig) SetNetworkData(v string) {
	o.NetworkData = &v
}

// GetNetworkDataBase64Encoded returns the NetworkDataBase64Encoded field value if set, zero value otherwise.
func (o *VirtualizationCloudInitConfig) GetNetworkDataBase64Encoded() bool {
	if o == nil || o.NetworkDataBase64Encoded == nil {
		var ret bool
		return ret
	}
	return *o.NetworkDataBase64Encoded
}

// GetNetworkDataBase64EncodedOk returns a tuple with the NetworkDataBase64Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetNetworkDataBase64EncodedOk() (*bool, bool) {
	if o == nil || o.NetworkDataBase64Encoded == nil {
		return nil, false
	}
	return o.NetworkDataBase64Encoded, true
}

// HasNetworkDataBase64Encoded returns a boolean if a field has been set.
func (o *VirtualizationCloudInitConfig) HasNetworkDataBase64Encoded() bool {
	if o != nil && o.NetworkDataBase64Encoded != nil {
		return true
	}

	return false
}

// SetNetworkDataBase64Encoded gets a reference to the given bool and assigns it to the NetworkDataBase64Encoded field.
func (o *VirtualizationCloudInitConfig) SetNetworkDataBase64Encoded(v bool) {
	o.NetworkDataBase64Encoded = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *VirtualizationCloudInitConfig) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *VirtualizationCloudInitConfig) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *VirtualizationCloudInitConfig) SetUserData(v string) {
	o.UserData = &v
}

// GetUserDataBase64Encoded returns the UserDataBase64Encoded field value if set, zero value otherwise.
func (o *VirtualizationCloudInitConfig) GetUserDataBase64Encoded() bool {
	if o == nil || o.UserDataBase64Encoded == nil {
		var ret bool
		return ret
	}
	return *o.UserDataBase64Encoded
}

// GetUserDataBase64EncodedOk returns a tuple with the UserDataBase64Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudInitConfig) GetUserDataBase64EncodedOk() (*bool, bool) {
	if o == nil || o.UserDataBase64Encoded == nil {
		return nil, false
	}
	return o.UserDataBase64Encoded, true
}

// HasUserDataBase64Encoded returns a boolean if a field has been set.
func (o *VirtualizationCloudInitConfig) HasUserDataBase64Encoded() bool {
	if o != nil && o.UserDataBase64Encoded != nil {
		return true
	}

	return false
}

// SetUserDataBase64Encoded gets a reference to the given bool and assigns it to the UserDataBase64Encoded field.
func (o *VirtualizationCloudInitConfig) SetUserDataBase64Encoded(v bool) {
	o.UserDataBase64Encoded = &v
}

func (o VirtualizationCloudInitConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigType != nil {
		toSerialize["ConfigType"] = o.ConfigType
	}
	if o.NetworkData != nil {
		toSerialize["NetworkData"] = o.NetworkData
	}
	if o.NetworkDataBase64Encoded != nil {
		toSerialize["NetworkDataBase64Encoded"] = o.NetworkDataBase64Encoded
	}
	if o.UserData != nil {
		toSerialize["UserData"] = o.UserData
	}
	if o.UserDataBase64Encoded != nil {
		toSerialize["UserDataBase64Encoded"] = o.UserDataBase64Encoded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCloudInitConfig) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationCloudInitConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Virtual machine cloud init configuration type. * `` - No cloud init specified. Cloud-init configurations are not sent to hypervisor, if none is selected. * `NoCloudSource` - Allows the user to provide user-data to the instance without running a network service. * `CloudConfigDrive` - Allows the user to provide user-data and network-data from cloud.
		ConfigType *string `json:"ConfigType,omitempty"`
		// Network configuration data for a virtual machine.
		NetworkData *string `json:"NetworkData,omitempty"`
		// Set to true, if the cloud init network data is in base64 format.
		NetworkDataBase64Encoded *bool `json:"NetworkDataBase64Encoded,omitempty"`
		// User configuration data for a virtual machine such as adding user, group etc.
		UserData *string `json:"UserData,omitempty"`
		// Set to true, if the cloud init user data is in base64 format.
		UserDataBase64Encoded *bool `json:"UserDataBase64Encoded,omitempty"`
	}

	varVirtualizationCloudInitConfigWithoutEmbeddedStruct := VirtualizationCloudInitConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationCloudInitConfigWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationCloudInitConfig := _VirtualizationCloudInitConfig{}
		varVirtualizationCloudInitConfig.ClassId = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.ClassId
		varVirtualizationCloudInitConfig.ObjectType = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.ObjectType
		varVirtualizationCloudInitConfig.ConfigType = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.ConfigType
		varVirtualizationCloudInitConfig.NetworkData = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.NetworkData
		varVirtualizationCloudInitConfig.NetworkDataBase64Encoded = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.NetworkDataBase64Encoded
		varVirtualizationCloudInitConfig.UserData = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.UserData
		varVirtualizationCloudInitConfig.UserDataBase64Encoded = varVirtualizationCloudInitConfigWithoutEmbeddedStruct.UserDataBase64Encoded
		*o = VirtualizationCloudInitConfig(varVirtualizationCloudInitConfig)
	} else {
		return err
	}

	varVirtualizationCloudInitConfig := _VirtualizationCloudInitConfig{}

	err = json.Unmarshal(bytes, &varVirtualizationCloudInitConfig)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationCloudInitConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigType")
		delete(additionalProperties, "NetworkData")
		delete(additionalProperties, "NetworkDataBase64Encoded")
		delete(additionalProperties, "UserData")
		delete(additionalProperties, "UserDataBase64Encoded")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVirtualizationCloudInitConfig struct {
	value *VirtualizationCloudInitConfig
	isSet bool
}

func (v NullableVirtualizationCloudInitConfig) Get() *VirtualizationCloudInitConfig {
	return v.value
}

func (v *NullableVirtualizationCloudInitConfig) Set(val *VirtualizationCloudInitConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCloudInitConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCloudInitConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCloudInitConfig(val *VirtualizationCloudInitConfig) *NullableVirtualizationCloudInitConfig {
	return &NullableVirtualizationCloudInitConfig{value: val, isSet: true}
}

func (v NullableVirtualizationCloudInitConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCloudInitConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
