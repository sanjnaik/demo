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
	"reflect"
	"strings"
)

// HyperflexDatastoreInfo Information related to the Protected Datastore.
type HyperflexDatastoreInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This datastore's backend unique id.
	DsBackendId *string `json:"DsBackendId,omitempty"`
	// This datastore's frontend id.
	DsFrontendId         *string `json:"DsFrontendId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDatastoreInfo HyperflexDatastoreInfo

// NewHyperflexDatastoreInfo instantiates a new HyperflexDatastoreInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDatastoreInfo(classId string, objectType string) *HyperflexDatastoreInfo {
	this := HyperflexDatastoreInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDatastoreInfoWithDefaults instantiates a new HyperflexDatastoreInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDatastoreInfoWithDefaults() *HyperflexDatastoreInfo {
	this := HyperflexDatastoreInfo{}
	var classId string = "hyperflex.DatastoreInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.DatastoreInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDatastoreInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDatastoreInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDatastoreInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDatastoreInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDsBackendId returns the DsBackendId field value if set, zero value otherwise.
func (o *HyperflexDatastoreInfo) GetDsBackendId() string {
	if o == nil || o.DsBackendId == nil {
		var ret string
		return ret
	}
	return *o.DsBackendId
}

// GetDsBackendIdOk returns a tuple with the DsBackendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfo) GetDsBackendIdOk() (*string, bool) {
	if o == nil || o.DsBackendId == nil {
		return nil, false
	}
	return o.DsBackendId, true
}

// HasDsBackendId returns a boolean if a field has been set.
func (o *HyperflexDatastoreInfo) HasDsBackendId() bool {
	if o != nil && o.DsBackendId != nil {
		return true
	}

	return false
}

// SetDsBackendId gets a reference to the given string and assigns it to the DsBackendId field.
func (o *HyperflexDatastoreInfo) SetDsBackendId(v string) {
	o.DsBackendId = &v
}

// GetDsFrontendId returns the DsFrontendId field value if set, zero value otherwise.
func (o *HyperflexDatastoreInfo) GetDsFrontendId() string {
	if o == nil || o.DsFrontendId == nil {
		var ret string
		return ret
	}
	return *o.DsFrontendId
}

// GetDsFrontendIdOk returns a tuple with the DsFrontendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfo) GetDsFrontendIdOk() (*string, bool) {
	if o == nil || o.DsFrontendId == nil {
		return nil, false
	}
	return o.DsFrontendId, true
}

// HasDsFrontendId returns a boolean if a field has been set.
func (o *HyperflexDatastoreInfo) HasDsFrontendId() bool {
	if o != nil && o.DsFrontendId != nil {
		return true
	}

	return false
}

// SetDsFrontendId gets a reference to the given string and assigns it to the DsFrontendId field.
func (o *HyperflexDatastoreInfo) SetDsFrontendId(v string) {
	o.DsFrontendId = &v
}

func (o HyperflexDatastoreInfo) MarshalJSON() ([]byte, error) {
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
	if o.DsBackendId != nil {
		toSerialize["DsBackendId"] = o.DsBackendId
	}
	if o.DsFrontendId != nil {
		toSerialize["DsFrontendId"] = o.DsFrontendId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDatastoreInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexDatastoreInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This datastore's backend unique id.
		DsBackendId *string `json:"DsBackendId,omitempty"`
		// This datastore's frontend id.
		DsFrontendId *string `json:"DsFrontendId,omitempty"`
	}

	varHyperflexDatastoreInfoWithoutEmbeddedStruct := HyperflexDatastoreInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexDatastoreInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexDatastoreInfo := _HyperflexDatastoreInfo{}
		varHyperflexDatastoreInfo.ClassId = varHyperflexDatastoreInfoWithoutEmbeddedStruct.ClassId
		varHyperflexDatastoreInfo.ObjectType = varHyperflexDatastoreInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexDatastoreInfo.DsBackendId = varHyperflexDatastoreInfoWithoutEmbeddedStruct.DsBackendId
		varHyperflexDatastoreInfo.DsFrontendId = varHyperflexDatastoreInfoWithoutEmbeddedStruct.DsFrontendId
		*o = HyperflexDatastoreInfo(varHyperflexDatastoreInfo)
	} else {
		return err
	}

	varHyperflexDatastoreInfo := _HyperflexDatastoreInfo{}

	err = json.Unmarshal(bytes, &varHyperflexDatastoreInfo)
	if err == nil {
		o.MoBaseComplexType = varHyperflexDatastoreInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DsBackendId")
		delete(additionalProperties, "DsFrontendId")

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

type NullableHyperflexDatastoreInfo struct {
	value *HyperflexDatastoreInfo
	isSet bool
}

func (v NullableHyperflexDatastoreInfo) Get() *HyperflexDatastoreInfo {
	return v.value
}

func (v *NullableHyperflexDatastoreInfo) Set(val *HyperflexDatastoreInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDatastoreInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDatastoreInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDatastoreInfo(val *HyperflexDatastoreInfo) *NullableHyperflexDatastoreInfo {
	return &NullableHyperflexDatastoreInfo{value: val, isSet: true}
}

func (v NullableHyperflexDatastoreInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDatastoreInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
