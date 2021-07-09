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

// AssetUsernamePasswordCredential A credential which performs authentication based on a username and password.
type AssetUsernamePasswordCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password used to authenticate with a managed target.
	Password *string `json:"Password,omitempty"`
	// The username used to authenticate with a managed target.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetUsernamePasswordCredential AssetUsernamePasswordCredential

// NewAssetUsernamePasswordCredential instantiates a new AssetUsernamePasswordCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetUsernamePasswordCredential(classId string, objectType string) *AssetUsernamePasswordCredential {
	this := AssetUsernamePasswordCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewAssetUsernamePasswordCredentialWithDefaults instantiates a new AssetUsernamePasswordCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetUsernamePasswordCredentialWithDefaults() *AssetUsernamePasswordCredential {
	this := AssetUsernamePasswordCredential{}
	var classId string = "asset.UsernamePasswordCredential"
	this.ClassId = classId
	var objectType string = "asset.UsernamePasswordCredential"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetUsernamePasswordCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredential) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetUsernamePasswordCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetUsernamePasswordCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetUsernamePasswordCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredential) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredential) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredential) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *AssetUsernamePasswordCredential) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredential) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredential) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredential) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AssetUsernamePasswordCredential) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredential) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredential) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredential) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AssetUsernamePasswordCredential) SetUsername(v string) {
	o.Username = &v
}

func (o AssetUsernamePasswordCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetUsernamePasswordCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetUsernamePasswordCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password used to authenticate with a managed target.
		Password *string `json:"Password,omitempty"`
		// The username used to authenticate with a managed target.
		Username *string `json:"Username,omitempty"`
	}

	varAssetUsernamePasswordCredentialWithoutEmbeddedStruct := AssetUsernamePasswordCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetUsernamePasswordCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetUsernamePasswordCredential := _AssetUsernamePasswordCredential{}
		varAssetUsernamePasswordCredential.ClassId = varAssetUsernamePasswordCredentialWithoutEmbeddedStruct.ClassId
		varAssetUsernamePasswordCredential.ObjectType = varAssetUsernamePasswordCredentialWithoutEmbeddedStruct.ObjectType
		varAssetUsernamePasswordCredential.IsPasswordSet = varAssetUsernamePasswordCredentialWithoutEmbeddedStruct.IsPasswordSet
		varAssetUsernamePasswordCredential.Password = varAssetUsernamePasswordCredentialWithoutEmbeddedStruct.Password
		varAssetUsernamePasswordCredential.Username = varAssetUsernamePasswordCredentialWithoutEmbeddedStruct.Username
		*o = AssetUsernamePasswordCredential(varAssetUsernamePasswordCredential)
	} else {
		return err
	}

	varAssetUsernamePasswordCredential := _AssetUsernamePasswordCredential{}

	err = json.Unmarshal(bytes, &varAssetUsernamePasswordCredential)
	if err == nil {
		o.AssetCredential = varAssetUsernamePasswordCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetUsernamePasswordCredential struct {
	value *AssetUsernamePasswordCredential
	isSet bool
}

func (v NullableAssetUsernamePasswordCredential) Get() *AssetUsernamePasswordCredential {
	return v.value
}

func (v *NullableAssetUsernamePasswordCredential) Set(val *AssetUsernamePasswordCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetUsernamePasswordCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetUsernamePasswordCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetUsernamePasswordCredential(val *AssetUsernamePasswordCredential) *NullableAssetUsernamePasswordCredential {
	return &NullableAssetUsernamePasswordCredential{value: val, isSet: true}
}

func (v NullableAssetUsernamePasswordCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetUsernamePasswordCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
