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

// RecommendationAbstractItemAllOf Definition of the list of properties defined in 'recommendation.AbstractItem', excluding properties defined in parent classes.
type RecommendationAbstractItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the physical device recommended.
	Name *string `json:"Name,omitempty"`
	// The type of physical device recommended. * `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk. * `Node` - The Enum value Node represents that the item type recommended is a Storage Node. * `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationAbstractItemAllOf RecommendationAbstractItemAllOf

// NewRecommendationAbstractItemAllOf instantiates a new RecommendationAbstractItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationAbstractItemAllOf(classId string, objectType string) *RecommendationAbstractItemAllOf {
	this := RecommendationAbstractItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationAbstractItemAllOfWithDefaults instantiates a new RecommendationAbstractItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationAbstractItemAllOfWithDefaults() *RecommendationAbstractItemAllOf {
	this := RecommendationAbstractItemAllOf{}
	var classId string = "recommendation.PhysicalItem"
	this.ClassId = classId
	var objectType string = "recommendation.PhysicalItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationAbstractItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationAbstractItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationAbstractItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationAbstractItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecommendationAbstractItemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecommendationAbstractItemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecommendationAbstractItemAllOf) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RecommendationAbstractItemAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationAbstractItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RecommendationAbstractItemAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RecommendationAbstractItemAllOf) SetType(v string) {
	o.Type = &v
}

func (o RecommendationAbstractItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationAbstractItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationAbstractItemAllOf := _RecommendationAbstractItemAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationAbstractItemAllOf); err == nil {
		*o = RecommendationAbstractItemAllOf(varRecommendationAbstractItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationAbstractItemAllOf struct {
	value *RecommendationAbstractItemAllOf
	isSet bool
}

func (v NullableRecommendationAbstractItemAllOf) Get() *RecommendationAbstractItemAllOf {
	return v.value
}

func (v *NullableRecommendationAbstractItemAllOf) Set(val *RecommendationAbstractItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationAbstractItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationAbstractItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationAbstractItemAllOf(val *RecommendationAbstractItemAllOf) *NullableRecommendationAbstractItemAllOf {
	return &NullableRecommendationAbstractItemAllOf{value: val, isSet: true}
}

func (v NullableRecommendationAbstractItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationAbstractItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
