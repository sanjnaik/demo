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
)

// OsWindowsParametersAllOf Definition of the list of properties defined in 'os.WindowsParameters', excluding properties defined in parent classes.
type OsWindowsParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Lists all the editions supported for Windows Server installation. * `Standard` - Windows Standard Edition ideal for advanced features with limited virtualization. * `StandardCore` - Windows Standard Core Edition is a minimal installation option while installing Standard Core that is ideal for advanced features with limited virtualization. * `Datacenter` - Windows Standard Core Edition ideal for high requirements on IT workloads with largenumber fo virtual systems. * `DatacenterCore` - Windows Datacenter Core Edition is a minimal installation option while installing Datacenter Core that isideal for high requirements on IT workloads with largenumber for virtual systems. * `Core` - Microsoft Hyper-V is a native hypervisor to create and run virtual machines.
	Edition              *string `json:"Edition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsWindowsParametersAllOf OsWindowsParametersAllOf

// NewOsWindowsParametersAllOf instantiates a new OsWindowsParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsWindowsParametersAllOf(classId string, objectType string) *OsWindowsParametersAllOf {
	this := OsWindowsParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var edition string = "Standard"
	this.Edition = &edition
	return &this
}

// NewOsWindowsParametersAllOfWithDefaults instantiates a new OsWindowsParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsWindowsParametersAllOfWithDefaults() *OsWindowsParametersAllOf {
	this := OsWindowsParametersAllOf{}
	var classId string = "os.WindowsParameters"
	this.ClassId = classId
	var objectType string = "os.WindowsParameters"
	this.ObjectType = objectType
	var edition string = "Standard"
	this.Edition = &edition
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsWindowsParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsWindowsParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsWindowsParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsWindowsParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsWindowsParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsWindowsParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *OsWindowsParametersAllOf) GetEdition() string {
	if o == nil || o.Edition == nil {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsWindowsParametersAllOf) GetEditionOk() (*string, bool) {
	if o == nil || o.Edition == nil {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *OsWindowsParametersAllOf) HasEdition() bool {
	if o != nil && o.Edition != nil {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *OsWindowsParametersAllOf) SetEdition(v string) {
	o.Edition = &v
}

func (o OsWindowsParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Edition != nil {
		toSerialize["Edition"] = o.Edition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsWindowsParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsWindowsParametersAllOf := _OsWindowsParametersAllOf{}

	if err = json.Unmarshal(bytes, &varOsWindowsParametersAllOf); err == nil {
		*o = OsWindowsParametersAllOf(varOsWindowsParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Edition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsWindowsParametersAllOf struct {
	value *OsWindowsParametersAllOf
	isSet bool
}

func (v NullableOsWindowsParametersAllOf) Get() *OsWindowsParametersAllOf {
	return v.value
}

func (v *NullableOsWindowsParametersAllOf) Set(val *OsWindowsParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsWindowsParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsWindowsParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsWindowsParametersAllOf(val *OsWindowsParametersAllOf) *NullableOsWindowsParametersAllOf {
	return &NullableOsWindowsParametersAllOf{value: val, isSet: true}
}

func (v NullableOsWindowsParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsWindowsParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
