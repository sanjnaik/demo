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

// StorageVolumeUtilization Storage space utilization of HyperFlex Volume entity.
type StorageVolumeUtilization struct {
	StorageBaseCapacity
	AdditionalProperties map[string]interface{}
}

type _StorageVolumeUtilization StorageVolumeUtilization

// NewStorageVolumeUtilization instantiates a new StorageVolumeUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVolumeUtilization() *StorageVolumeUtilization {
	this := StorageVolumeUtilization{}
	return &this
}

// NewStorageVolumeUtilizationWithDefaults instantiates a new StorageVolumeUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVolumeUtilizationWithDefaults() *StorageVolumeUtilization {
	this := StorageVolumeUtilization{}
	return &this
}

func (o StorageVolumeUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseCapacity, errStorageBaseCapacity := json.Marshal(o.StorageBaseCapacity)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	errStorageBaseCapacity = json.Unmarshal([]byte(serializedStorageBaseCapacity), &toSerialize)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVolumeUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVolumeUtilizationWithoutEmbeddedStruct struct {
	}

	varStorageVolumeUtilizationWithoutEmbeddedStruct := StorageVolumeUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVolumeUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStorageVolumeUtilization := _StorageVolumeUtilization{}
		*o = StorageVolumeUtilization(varStorageVolumeUtilization)
	} else {
		return err
	}

	varStorageVolumeUtilization := _StorageVolumeUtilization{}

	err = json.Unmarshal(bytes, &varStorageVolumeUtilization)
	if err == nil {
		o.StorageBaseCapacity = varStorageVolumeUtilization.StorageBaseCapacity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectStorageBaseCapacity := reflect.ValueOf(o.StorageBaseCapacity)
		for i := 0; i < reflectStorageBaseCapacity.Type().NumField(); i++ {
			t := reflectStorageBaseCapacity.Type().Field(i)

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

type NullableStorageVolumeUtilization struct {
	value *StorageVolumeUtilization
	isSet bool
}

func (v NullableStorageVolumeUtilization) Get() *StorageVolumeUtilization {
	return v.value
}

func (v *NullableStorageVolumeUtilization) Set(val *StorageVolumeUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVolumeUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVolumeUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVolumeUtilization(val *StorageVolumeUtilization) *NullableStorageVolumeUtilization {
	return &NullableStorageVolumeUtilization{value: val, isSet: true}
}

func (v NullableStorageVolumeUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVolumeUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
