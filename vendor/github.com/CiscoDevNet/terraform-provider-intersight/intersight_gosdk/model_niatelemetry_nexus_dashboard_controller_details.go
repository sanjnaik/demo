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

// NiatelemetryNexusDashboardControllerDetails Details of controller added to NexusDashboard.
type NiatelemetryNexusDashboardControllerDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Health of the site serviced by ND.
	SiteHealth *int64 `json:"SiteHealth,omitempty"`
	// Name of fabric domain of the controller.
	SiteName *string `json:"SiteName,omitempty"`
	// Version of the controller serviced by ND.
	VersionOfController  *string                              `json:"VersionOfController,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboardControllerDetails NiatelemetryNexusDashboardControllerDetails

// NewNiatelemetryNexusDashboardControllerDetails instantiates a new NiatelemetryNexusDashboardControllerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardControllerDetails(classId string, objectType string) *NiatelemetryNexusDashboardControllerDetails {
	this := NiatelemetryNexusDashboardControllerDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardControllerDetailsWithDefaults instantiates a new NiatelemetryNexusDashboardControllerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardControllerDetailsWithDefaults() *NiatelemetryNexusDashboardControllerDetails {
	this := NiatelemetryNexusDashboardControllerDetails{}
	var classId string = "niatelemetry.NexusDashboardControllerDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboardControllerDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardControllerDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardControllerDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardControllerDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardControllerDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSiteHealth returns the SiteHealth field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetails) GetSiteHealth() int64 {
	if o == nil || o.SiteHealth == nil {
		var ret int64
		return ret
	}
	return *o.SiteHealth
}

// GetSiteHealthOk returns a tuple with the SiteHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetSiteHealthOk() (*int64, bool) {
	if o == nil || o.SiteHealth == nil {
		return nil, false
	}
	return o.SiteHealth, true
}

// HasSiteHealth returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) HasSiteHealth() bool {
	if o != nil && o.SiteHealth != nil {
		return true
	}

	return false
}

// SetSiteHealth gets a reference to the given int64 and assigns it to the SiteHealth field.
func (o *NiatelemetryNexusDashboardControllerDetails) SetSiteHealth(v int64) {
	o.SiteHealth = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryNexusDashboardControllerDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVersionOfController returns the VersionOfController field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetails) GetVersionOfController() string {
	if o == nil || o.VersionOfController == nil {
		var ret string
		return ret
	}
	return *o.VersionOfController
}

// GetVersionOfControllerOk returns a tuple with the VersionOfController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetVersionOfControllerOk() (*string, bool) {
	if o == nil || o.VersionOfController == nil {
		return nil, false
	}
	return o.VersionOfController, true
}

// HasVersionOfController returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) HasVersionOfController() bool {
	if o != nil && o.VersionOfController != nil {
		return true
	}

	return false
}

// SetVersionOfController gets a reference to the given string and assigns it to the VersionOfController field.
func (o *NiatelemetryNexusDashboardControllerDetails) SetVersionOfController(v string) {
	o.VersionOfController = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardControllerDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardControllerDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardControllerDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardControllerDetails) MarshalJSON() ([]byte, error) {
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
	if o.SiteHealth != nil {
		toSerialize["SiteHealth"] = o.SiteHealth
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.VersionOfController != nil {
		toSerialize["VersionOfController"] = o.VersionOfController
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardControllerDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Health of the site serviced by ND.
		SiteHealth *int64 `json:"SiteHealth,omitempty"`
		// Name of fabric domain of the controller.
		SiteName *string `json:"SiteName,omitempty"`
		// Version of the controller serviced by ND.
		VersionOfController *string                              `json:"VersionOfController,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct := NiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNexusDashboardControllerDetails := _NiatelemetryNexusDashboardControllerDetails{}
		varNiatelemetryNexusDashboardControllerDetails.ClassId = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryNexusDashboardControllerDetails.ObjectType = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNexusDashboardControllerDetails.SiteHealth = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.SiteHealth
		varNiatelemetryNexusDashboardControllerDetails.SiteName = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryNexusDashboardControllerDetails.VersionOfController = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.VersionOfController
		varNiatelemetryNexusDashboardControllerDetails.RegisteredDevice = varNiatelemetryNexusDashboardControllerDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryNexusDashboardControllerDetails(varNiatelemetryNexusDashboardControllerDetails)
	} else {
		return err
	}

	varNiatelemetryNexusDashboardControllerDetails := _NiatelemetryNexusDashboardControllerDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardControllerDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryNexusDashboardControllerDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SiteHealth")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "VersionOfController")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryNexusDashboardControllerDetails struct {
	value *NiatelemetryNexusDashboardControllerDetails
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardControllerDetails) Get() *NiatelemetryNexusDashboardControllerDetails {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardControllerDetails) Set(val *NiatelemetryNexusDashboardControllerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardControllerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardControllerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardControllerDetails(val *NiatelemetryNexusDashboardControllerDetails) *NullableNiatelemetryNexusDashboardControllerDetails {
	return &NullableNiatelemetryNexusDashboardControllerDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardControllerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardControllerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
