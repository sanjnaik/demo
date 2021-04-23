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
	"reflect"
	"strings"
	"time"
)

// MoVersionContext VersionContext contains the versioning info for an object.
type MoVersionContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string    `json:"ObjectType"`
	InterestedMos []MoMoRef `json:"InterestedMos,omitempty"`
	RefMo         *MoMoRef  `json:"RefMo,omitempty"`
	// The time this versioned Managed Object was created.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	Version *string `json:"Version,omitempty"`
	// Specifies type of version. Currently the only supported value is \"Configured\" that is used to keep track of snapshots of policies and profiles that are intended to be configured to target endpoints. * `Modified` - Version created every time an object is modified. * `Configured` - Version created every time an object is configured to the service profile. * `Deployed` - Version created for objects related to a service profile when it is deployed.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoVersionContext MoVersionContext

// NewMoVersionContext instantiates a new MoVersionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoVersionContext(classId string, objectType string) *MoVersionContext {
	this := MoVersionContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// NewMoVersionContextWithDefaults instantiates a new MoVersionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoVersionContextWithDefaults() *MoVersionContext {
	this := MoVersionContext{}
	var classId string = "mo.VersionContext"
	this.ClassId = classId
	var objectType string = "mo.VersionContext"
	this.ObjectType = objectType
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MoVersionContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MoVersionContext) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MoVersionContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MoVersionContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterestedMos returns the InterestedMos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoVersionContext) GetInterestedMos() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.InterestedMos
}

// GetInterestedMosOk returns a tuple with the InterestedMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoVersionContext) GetInterestedMosOk() (*[]MoMoRef, bool) {
	if o == nil || o.InterestedMos == nil {
		return nil, false
	}
	return &o.InterestedMos, true
}

// HasInterestedMos returns a boolean if a field has been set.
func (o *MoVersionContext) HasInterestedMos() bool {
	if o != nil && o.InterestedMos != nil {
		return true
	}

	return false
}

// SetInterestedMos gets a reference to the given []MoMoRef and assigns it to the InterestedMos field.
func (o *MoVersionContext) SetInterestedMos(v []MoMoRef) {
	o.InterestedMos = v
}

// GetRefMo returns the RefMo field value if set, zero value otherwise.
func (o *MoVersionContext) GetRefMo() MoMoRef {
	if o == nil || o.RefMo == nil {
		var ret MoMoRef
		return ret
	}
	return *o.RefMo
}

// GetRefMoOk returns a tuple with the RefMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetRefMoOk() (*MoMoRef, bool) {
	if o == nil || o.RefMo == nil {
		return nil, false
	}
	return o.RefMo, true
}

// HasRefMo returns a boolean if a field has been set.
func (o *MoVersionContext) HasRefMo() bool {
	if o != nil && o.RefMo != nil {
		return true
	}

	return false
}

// SetRefMo gets a reference to the given MoMoRef and assigns it to the RefMo field.
func (o *MoVersionContext) SetRefMo(v MoMoRef) {
	o.RefMo = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MoVersionContext) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MoVersionContext) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MoVersionContext) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MoVersionContext) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MoVersionContext) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MoVersionContext) SetVersion(v string) {
	o.Version = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *MoVersionContext) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContext) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *MoVersionContext) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *MoVersionContext) SetVersionType(v string) {
	o.VersionType = &v
}

func (o MoVersionContext) MarshalJSON() ([]byte, error) {
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
	if o.InterestedMos != nil {
		toSerialize["InterestedMos"] = o.InterestedMos
	}
	if o.RefMo != nil {
		toSerialize["RefMo"] = o.RefMo
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoVersionContext) UnmarshalJSON(bytes []byte) (err error) {
	type MoVersionContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string    `json:"ObjectType"`
		InterestedMos []MoMoRef `json:"InterestedMos,omitempty"`
		RefMo         *MoMoRef  `json:"RefMo,omitempty"`
		// The time this versioned Managed Object was created.
		Timestamp *time.Time `json:"Timestamp,omitempty"`
		// The version of the Managed Object, e.g. an incrementing number or a hash id.
		Version *string `json:"Version,omitempty"`
		// Specifies type of version. Currently the only supported value is \"Configured\" that is used to keep track of snapshots of policies and profiles that are intended to be configured to target endpoints. * `Modified` - Version created every time an object is modified. * `Configured` - Version created every time an object is configured to the service profile. * `Deployed` - Version created for objects related to a service profile when it is deployed.
		VersionType *string `json:"VersionType,omitempty"`
	}

	varMoVersionContextWithoutEmbeddedStruct := MoVersionContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMoVersionContextWithoutEmbeddedStruct)
	if err == nil {
		varMoVersionContext := _MoVersionContext{}
		varMoVersionContext.ClassId = varMoVersionContextWithoutEmbeddedStruct.ClassId
		varMoVersionContext.ObjectType = varMoVersionContextWithoutEmbeddedStruct.ObjectType
		varMoVersionContext.InterestedMos = varMoVersionContextWithoutEmbeddedStruct.InterestedMos
		varMoVersionContext.RefMo = varMoVersionContextWithoutEmbeddedStruct.RefMo
		varMoVersionContext.Timestamp = varMoVersionContextWithoutEmbeddedStruct.Timestamp
		varMoVersionContext.Version = varMoVersionContextWithoutEmbeddedStruct.Version
		varMoVersionContext.VersionType = varMoVersionContextWithoutEmbeddedStruct.VersionType
		*o = MoVersionContext(varMoVersionContext)
	} else {
		return err
	}

	varMoVersionContext := _MoVersionContext{}

	err = json.Unmarshal(bytes, &varMoVersionContext)
	if err == nil {
		o.MoBaseComplexType = varMoVersionContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterestedMos")
		delete(additionalProperties, "RefMo")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "VersionType")

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

type NullableMoVersionContext struct {
	value *MoVersionContext
	isSet bool
}

func (v NullableMoVersionContext) Get() *MoVersionContext {
	return v.value
}

func (v *NullableMoVersionContext) Set(val *MoVersionContext) {
	v.value = val
	v.isSet = true
}

func (v NullableMoVersionContext) IsSet() bool {
	return v.isSet
}

func (v *NullableMoVersionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoVersionContext(val *MoVersionContext) *NullableMoVersionContext {
	return &NullableMoVersionContext{value: val, isSet: true}
}

func (v NullableMoVersionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoVersionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
