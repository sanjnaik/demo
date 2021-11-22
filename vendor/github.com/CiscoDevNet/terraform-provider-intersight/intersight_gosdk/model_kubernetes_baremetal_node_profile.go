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
	"reflect"
	"strings"
)

// KubernetesBaremetalNodeProfile A profile specifying configuration settings for a baremetal node. Users can do operations like Drain, Cordon, Rebuild on a node.
type KubernetesBaremetalNodeProfile struct {
	KubernetesNodeProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Network interface from NetworkInfo (by name) to use for kubernetes VIP.
	KubernetesNic        *string                                `json:"KubernetesNic,omitempty"`
	NetworkInfo          NullableKubernetesBaremetalNetworkInfo `json:"NetworkInfo,omitempty"`
	Server               *ComputeRackUnitRelationship           `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaremetalNodeProfile KubernetesBaremetalNodeProfile

// NewKubernetesBaremetalNodeProfile instantiates a new KubernetesBaremetalNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaremetalNodeProfile(classId string, objectType string) *KubernetesBaremetalNodeProfile {
	this := KubernetesBaremetalNodeProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// NewKubernetesBaremetalNodeProfileWithDefaults instantiates a new KubernetesBaremetalNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaremetalNodeProfileWithDefaults() *KubernetesBaremetalNodeProfile {
	this := KubernetesBaremetalNodeProfile{}
	var classId string = "kubernetes.BaremetalNodeProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.BaremetalNodeProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaremetalNodeProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaremetalNodeProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaremetalNodeProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaremetalNodeProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKubernetesNic returns the KubernetesNic field value if set, zero value otherwise.
func (o *KubernetesBaremetalNodeProfile) GetKubernetesNic() string {
	if o == nil || o.KubernetesNic == nil {
		var ret string
		return ret
	}
	return *o.KubernetesNic
}

// GetKubernetesNicOk returns a tuple with the KubernetesNic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfile) GetKubernetesNicOk() (*string, bool) {
	if o == nil || o.KubernetesNic == nil {
		return nil, false
	}
	return o.KubernetesNic, true
}

// HasKubernetesNic returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfile) HasKubernetesNic() bool {
	if o != nil && o.KubernetesNic != nil {
		return true
	}

	return false
}

// SetKubernetesNic gets a reference to the given string and assigns it to the KubernetesNic field.
func (o *KubernetesBaremetalNodeProfile) SetKubernetesNic(v string) {
	o.KubernetesNic = &v
}

// GetNetworkInfo returns the NetworkInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesBaremetalNodeProfile) GetNetworkInfo() KubernetesBaremetalNetworkInfo {
	if o == nil || o.NetworkInfo.Get() == nil {
		var ret KubernetesBaremetalNetworkInfo
		return ret
	}
	return *o.NetworkInfo.Get()
}

// GetNetworkInfoOk returns a tuple with the NetworkInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesBaremetalNodeProfile) GetNetworkInfoOk() (*KubernetesBaremetalNetworkInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkInfo.Get(), o.NetworkInfo.IsSet()
}

// HasNetworkInfo returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfile) HasNetworkInfo() bool {
	if o != nil && o.NetworkInfo.IsSet() {
		return true
	}

	return false
}

// SetNetworkInfo gets a reference to the given NullableKubernetesBaremetalNetworkInfo and assigns it to the NetworkInfo field.
func (o *KubernetesBaremetalNodeProfile) SetNetworkInfo(v KubernetesBaremetalNetworkInfo) {
	o.NetworkInfo.Set(&v)
}

// SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil
func (o *KubernetesBaremetalNodeProfile) SetNetworkInfoNil() {
	o.NetworkInfo.Set(nil)
}

// UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
func (o *KubernetesBaremetalNodeProfile) UnsetNetworkInfo() {
	o.NetworkInfo.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KubernetesBaremetalNodeProfile) GetServer() ComputeRackUnitRelationship {
	if o == nil || o.Server == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaremetalNodeProfile) GetServerOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KubernetesBaremetalNodeProfile) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputeRackUnitRelationship and assigns it to the Server field.
func (o *KubernetesBaremetalNodeProfile) SetServer(v ComputeRackUnitRelationship) {
	o.Server = &v
}

func (o KubernetesBaremetalNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesNodeProfile, errKubernetesNodeProfile := json.Marshal(o.KubernetesNodeProfile)
	if errKubernetesNodeProfile != nil {
		return []byte{}, errKubernetesNodeProfile
	}
	errKubernetesNodeProfile = json.Unmarshal([]byte(serializedKubernetesNodeProfile), &toSerialize)
	if errKubernetesNodeProfile != nil {
		return []byte{}, errKubernetesNodeProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KubernetesNic != nil {
		toSerialize["KubernetesNic"] = o.KubernetesNic
	}
	if o.NetworkInfo.IsSet() {
		toSerialize["NetworkInfo"] = o.NetworkInfo.Get()
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaremetalNodeProfile) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesBaremetalNodeProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Network interface from NetworkInfo (by name) to use for kubernetes VIP.
		KubernetesNic *string                                `json:"KubernetesNic,omitempty"`
		NetworkInfo   NullableKubernetesBaremetalNetworkInfo `json:"NetworkInfo,omitempty"`
		Server        *ComputeRackUnitRelationship           `json:"Server,omitempty"`
	}

	varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct := KubernetesBaremetalNodeProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesBaremetalNodeProfile := _KubernetesBaremetalNodeProfile{}
		varKubernetesBaremetalNodeProfile.ClassId = varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct.ClassId
		varKubernetesBaremetalNodeProfile.ObjectType = varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesBaremetalNodeProfile.KubernetesNic = varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct.KubernetesNic
		varKubernetesBaremetalNodeProfile.NetworkInfo = varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct.NetworkInfo
		varKubernetesBaremetalNodeProfile.Server = varKubernetesBaremetalNodeProfileWithoutEmbeddedStruct.Server
		*o = KubernetesBaremetalNodeProfile(varKubernetesBaremetalNodeProfile)
	} else {
		return err
	}

	varKubernetesBaremetalNodeProfile := _KubernetesBaremetalNodeProfile{}

	err = json.Unmarshal(bytes, &varKubernetesBaremetalNodeProfile)
	if err == nil {
		o.KubernetesNodeProfile = varKubernetesBaremetalNodeProfile.KubernetesNodeProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KubernetesNic")
		delete(additionalProperties, "NetworkInfo")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectKubernetesNodeProfile := reflect.ValueOf(o.KubernetesNodeProfile)
		for i := 0; i < reflectKubernetesNodeProfile.Type().NumField(); i++ {
			t := reflectKubernetesNodeProfile.Type().Field(i)

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

type NullableKubernetesBaremetalNodeProfile struct {
	value *KubernetesBaremetalNodeProfile
	isSet bool
}

func (v NullableKubernetesBaremetalNodeProfile) Get() *KubernetesBaremetalNodeProfile {
	return v.value
}

func (v *NullableKubernetesBaremetalNodeProfile) Set(val *KubernetesBaremetalNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaremetalNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaremetalNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaremetalNodeProfile(val *KubernetesBaremetalNodeProfile) *NullableKubernetesBaremetalNodeProfile {
	return &NullableKubernetesBaremetalNodeProfile{value: val, isSet: true}
}

func (v NullableKubernetesBaremetalNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaremetalNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
