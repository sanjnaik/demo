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

// KubernetesKubernetesResource Kubernetes Resource Base Class.
type KubernetesKubernetesResource struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                       `json:"ObjectType"`
	Metadata   NullableKubernetesObjectMeta `json:"Metadata,omitempty"`
	// Name of the referenced kubernetes resource.
	Name *string `json:"Name,omitempty"`
	// UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesKubernetesResource KubernetesKubernetesResource

// NewKubernetesKubernetesResource instantiates a new KubernetesKubernetesResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesKubernetesResource(classId string, objectType string) *KubernetesKubernetesResource {
	this := KubernetesKubernetesResource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesKubernetesResourceWithDefaults instantiates a new KubernetesKubernetesResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesKubernetesResourceWithDefaults() *KubernetesKubernetesResource {
	this := KubernetesKubernetesResource{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesKubernetesResource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesKubernetesResource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesKubernetesResource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesKubernetesResource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesKubernetesResource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesKubernetesResource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesKubernetesResource) GetMetadata() KubernetesObjectMeta {
	if o == nil || o.Metadata.Get() == nil {
		var ret KubernetesObjectMeta
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesKubernetesResource) GetMetadataOk() (*KubernetesObjectMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesKubernetesResource) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableKubernetesObjectMeta and assigns it to the Metadata field.
func (o *KubernetesKubernetesResource) SetMetadata(v KubernetesObjectMeta) {
	o.Metadata.Set(&v)
}

// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *KubernetesKubernetesResource) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *KubernetesKubernetesResource) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesKubernetesResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesKubernetesResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesKubernetesResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesKubernetesResource) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *KubernetesKubernetesResource) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesKubernetesResource) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *KubernetesKubernetesResource) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *KubernetesKubernetesResource) SetUuid(v string) {
	o.Uuid = &v
}

func (o KubernetesKubernetesResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Metadata.IsSet() {
		toSerialize["Metadata"] = o.Metadata.Get()
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesKubernetesResource) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesKubernetesResourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                       `json:"ObjectType"`
		Metadata   NullableKubernetesObjectMeta `json:"Metadata,omitempty"`
		// Name of the referenced kubernetes resource.
		Name *string `json:"Name,omitempty"`
		// UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varKubernetesKubernetesResourceWithoutEmbeddedStruct := KubernetesKubernetesResourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesKubernetesResourceWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesKubernetesResource := _KubernetesKubernetesResource{}
		varKubernetesKubernetesResource.ClassId = varKubernetesKubernetesResourceWithoutEmbeddedStruct.ClassId
		varKubernetesKubernetesResource.ObjectType = varKubernetesKubernetesResourceWithoutEmbeddedStruct.ObjectType
		varKubernetesKubernetesResource.Metadata = varKubernetesKubernetesResourceWithoutEmbeddedStruct.Metadata
		varKubernetesKubernetesResource.Name = varKubernetesKubernetesResourceWithoutEmbeddedStruct.Name
		varKubernetesKubernetesResource.Uuid = varKubernetesKubernetesResourceWithoutEmbeddedStruct.Uuid
		*o = KubernetesKubernetesResource(varKubernetesKubernetesResource)
	} else {
		return err
	}

	varKubernetesKubernetesResource := _KubernetesKubernetesResource{}

	err = json.Unmarshal(bytes, &varKubernetesKubernetesResource)
	if err == nil {
		o.MoBaseMo = varKubernetesKubernetesResource.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Metadata")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableKubernetesKubernetesResource struct {
	value *KubernetesKubernetesResource
	isSet bool
}

func (v NullableKubernetesKubernetesResource) Get() *KubernetesKubernetesResource {
	return v.value
}

func (v *NullableKubernetesKubernetesResource) Set(val *KubernetesKubernetesResource) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesKubernetesResource) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesKubernetesResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesKubernetesResource(val *KubernetesKubernetesResource) *NullableKubernetesKubernetesResource {
	return &NullableKubernetesKubernetesResource{value: val, isSet: true}
}

func (v NullableKubernetesKubernetesResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesKubernetesResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
