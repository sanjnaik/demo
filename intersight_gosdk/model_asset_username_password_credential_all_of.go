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

// AssetUsernamePasswordCredentialAllOf Definition of the list of properties defined in 'asset.UsernamePasswordCredential', excluding properties defined in parent classes.
type AssetUsernamePasswordCredentialAllOf struct {
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

type _AssetUsernamePasswordCredentialAllOf AssetUsernamePasswordCredentialAllOf

// NewAssetUsernamePasswordCredentialAllOf instantiates a new AssetUsernamePasswordCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetUsernamePasswordCredentialAllOf(classId string, objectType string) *AssetUsernamePasswordCredentialAllOf {
	this := AssetUsernamePasswordCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewAssetUsernamePasswordCredentialAllOfWithDefaults instantiates a new AssetUsernamePasswordCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetUsernamePasswordCredentialAllOfWithDefaults() *AssetUsernamePasswordCredentialAllOf {
	this := AssetUsernamePasswordCredentialAllOf{}
	var classId string = "asset.UsernamePasswordCredential"
	this.ClassId = classId
	var objectType string = "asset.UsernamePasswordCredential"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetUsernamePasswordCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetUsernamePasswordCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetUsernamePasswordCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetUsernamePasswordCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredentialAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredentialAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredentialAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *AssetUsernamePasswordCredentialAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredentialAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredentialAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredentialAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AssetUsernamePasswordCredentialAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AssetUsernamePasswordCredentialAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetUsernamePasswordCredentialAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AssetUsernamePasswordCredentialAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AssetUsernamePasswordCredentialAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o AssetUsernamePasswordCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *AssetUsernamePasswordCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetUsernamePasswordCredentialAllOf := _AssetUsernamePasswordCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varAssetUsernamePasswordCredentialAllOf); err == nil {
		*o = AssetUsernamePasswordCredentialAllOf(varAssetUsernamePasswordCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetUsernamePasswordCredentialAllOf struct {
	value *AssetUsernamePasswordCredentialAllOf
	isSet bool
}

func (v NullableAssetUsernamePasswordCredentialAllOf) Get() *AssetUsernamePasswordCredentialAllOf {
	return v.value
}

func (v *NullableAssetUsernamePasswordCredentialAllOf) Set(val *AssetUsernamePasswordCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetUsernamePasswordCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetUsernamePasswordCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetUsernamePasswordCredentialAllOf(val *AssetUsernamePasswordCredentialAllOf) *NullableAssetUsernamePasswordCredentialAllOf {
	return &NullableAssetUsernamePasswordCredentialAllOf{value: val, isSet: true}
}

func (v NullableAssetUsernamePasswordCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetUsernamePasswordCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
