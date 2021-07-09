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

// KubernetesIngressAllOf Definition of the list of properties defined in 'kubernetes.Ingress', excluding properties defined in parent classes.
type KubernetesIngressAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Status               NullableKubernetesIngressStatus      `json:"Status,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesIngressAllOf KubernetesIngressAllOf

// NewKubernetesIngressAllOf instantiates a new KubernetesIngressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIngressAllOf(classId string, objectType string) *KubernetesIngressAllOf {
	this := KubernetesIngressAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesIngressAllOfWithDefaults instantiates a new KubernetesIngressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIngressAllOfWithDefaults() *KubernetesIngressAllOf {
	this := KubernetesIngressAllOf{}
	var classId string = "kubernetes.Ingress"
	this.ClassId = classId
	var objectType string = "kubernetes.Ingress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesIngressAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngressAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesIngressAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesIngressAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngressAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesIngressAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesIngressAllOf) GetStatus() KubernetesIngressStatus {
	if o == nil || o.Status.Get() == nil {
		var ret KubernetesIngressStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesIngressAllOf) GetStatusOk() (*KubernetesIngressStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesIngressAllOf) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableKubernetesIngressStatus and assigns it to the Status field.
func (o *KubernetesIngressAllOf) SetStatus(v KubernetesIngressStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesIngressAllOf) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesIngressAllOf) UnsetStatus() {
	o.Status.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesIngressAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesIngressAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesIngressAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesIngressAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesIngressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesIngressAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesIngressAllOf := _KubernetesIngressAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesIngressAllOf); err == nil {
		*o = KubernetesIngressAllOf(varKubernetesIngressAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesIngressAllOf struct {
	value *KubernetesIngressAllOf
	isSet bool
}

func (v NullableKubernetesIngressAllOf) Get() *KubernetesIngressAllOf {
	return v.value
}

func (v *NullableKubernetesIngressAllOf) Set(val *KubernetesIngressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIngressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIngressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIngressAllOf(val *KubernetesIngressAllOf) *NullableKubernetesIngressAllOf {
	return &NullableKubernetesIngressAllOf{value: val, isSet: true}
}

func (v NullableKubernetesIngressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIngressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
