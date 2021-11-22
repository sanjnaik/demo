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
)

// AssetNewRelicCredentialAllOf Definition of the list of properties defined in 'asset.NewRelicCredential', excluding properties defined in parent classes.
type AssetNewRelicCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                 `json:"ObjectType"`
	GraphQlCredential    *AssetApiKeyCredential `json:"GraphQlCredential,omitempty"`
	RestApiCredential    *AssetApiKeyCredential `json:"RestApiCredential,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetNewRelicCredentialAllOf AssetNewRelicCredentialAllOf

// NewAssetNewRelicCredentialAllOf instantiates a new AssetNewRelicCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetNewRelicCredentialAllOf(classId string, objectType string) *AssetNewRelicCredentialAllOf {
	this := AssetNewRelicCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetNewRelicCredentialAllOfWithDefaults instantiates a new AssetNewRelicCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetNewRelicCredentialAllOfWithDefaults() *AssetNewRelicCredentialAllOf {
	this := AssetNewRelicCredentialAllOf{}
	var classId string = "asset.NewRelicCredential"
	this.ClassId = classId
	var objectType string = "asset.NewRelicCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetNewRelicCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetNewRelicCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetNewRelicCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetNewRelicCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetNewRelicCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetNewRelicCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGraphQlCredential returns the GraphQlCredential field value if set, zero value otherwise.
func (o *AssetNewRelicCredentialAllOf) GetGraphQlCredential() AssetApiKeyCredential {
	if o == nil || o.GraphQlCredential == nil {
		var ret AssetApiKeyCredential
		return ret
	}
	return *o.GraphQlCredential
}

// GetGraphQlCredentialOk returns a tuple with the GraphQlCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetNewRelicCredentialAllOf) GetGraphQlCredentialOk() (*AssetApiKeyCredential, bool) {
	if o == nil || o.GraphQlCredential == nil {
		return nil, false
	}
	return o.GraphQlCredential, true
}

// HasGraphQlCredential returns a boolean if a field has been set.
func (o *AssetNewRelicCredentialAllOf) HasGraphQlCredential() bool {
	if o != nil && o.GraphQlCredential != nil {
		return true
	}

	return false
}

// SetGraphQlCredential gets a reference to the given AssetApiKeyCredential and assigns it to the GraphQlCredential field.
func (o *AssetNewRelicCredentialAllOf) SetGraphQlCredential(v AssetApiKeyCredential) {
	o.GraphQlCredential = &v
}

// GetRestApiCredential returns the RestApiCredential field value if set, zero value otherwise.
func (o *AssetNewRelicCredentialAllOf) GetRestApiCredential() AssetApiKeyCredential {
	if o == nil || o.RestApiCredential == nil {
		var ret AssetApiKeyCredential
		return ret
	}
	return *o.RestApiCredential
}

// GetRestApiCredentialOk returns a tuple with the RestApiCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetNewRelicCredentialAllOf) GetRestApiCredentialOk() (*AssetApiKeyCredential, bool) {
	if o == nil || o.RestApiCredential == nil {
		return nil, false
	}
	return o.RestApiCredential, true
}

// HasRestApiCredential returns a boolean if a field has been set.
func (o *AssetNewRelicCredentialAllOf) HasRestApiCredential() bool {
	if o != nil && o.RestApiCredential != nil {
		return true
	}

	return false
}

// SetRestApiCredential gets a reference to the given AssetApiKeyCredential and assigns it to the RestApiCredential field.
func (o *AssetNewRelicCredentialAllOf) SetRestApiCredential(v AssetApiKeyCredential) {
	o.RestApiCredential = &v
}

func (o AssetNewRelicCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.GraphQlCredential != nil {
		toSerialize["GraphQlCredential"] = o.GraphQlCredential
	}
	if o.RestApiCredential != nil {
		toSerialize["RestApiCredential"] = o.RestApiCredential
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetNewRelicCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetNewRelicCredentialAllOf := _AssetNewRelicCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varAssetNewRelicCredentialAllOf); err == nil {
		*o = AssetNewRelicCredentialAllOf(varAssetNewRelicCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GraphQlCredential")
		delete(additionalProperties, "RestApiCredential")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetNewRelicCredentialAllOf struct {
	value *AssetNewRelicCredentialAllOf
	isSet bool
}

func (v NullableAssetNewRelicCredentialAllOf) Get() *AssetNewRelicCredentialAllOf {
	return v.value
}

func (v *NullableAssetNewRelicCredentialAllOf) Set(val *AssetNewRelicCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetNewRelicCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetNewRelicCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetNewRelicCredentialAllOf(val *AssetNewRelicCredentialAllOf) *NullableAssetNewRelicCredentialAllOf {
	return &NullableAssetNewRelicCredentialAllOf{value: val, isSet: true}
}

func (v NullableAssetNewRelicCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetNewRelicCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
