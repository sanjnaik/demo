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

// UuidpoolUuidLeaseAllOf Definition of the list of properties defined in 'uuidpool.UuidLease', excluding properties defined in parent classes.
type UuidpoolUuidLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID Prefix+Suffix numbers.
	Uuid                 *string                         `json:"Uuid,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship           `json:"AssignedToEntity,omitempty"`
	Pool                 *UuidpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *UuidpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *UuidpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolUuidLeaseAllOf UuidpoolUuidLeaseAllOf

// NewUuidpoolUuidLeaseAllOf instantiates a new UuidpoolUuidLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolUuidLeaseAllOf(classId string, objectType string) *UuidpoolUuidLeaseAllOf {
	this := UuidpoolUuidLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUuidpoolUuidLeaseAllOfWithDefaults instantiates a new UuidpoolUuidLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolUuidLeaseAllOfWithDefaults() *UuidpoolUuidLeaseAllOf {
	this := UuidpoolUuidLeaseAllOf{}
	var classId string = "uuidpool.UuidLease"
	this.ClassId = classId
	var objectType string = "uuidpool.UuidLease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolUuidLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolUuidLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolUuidLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolUuidLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolUuidLeaseAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolUuidLeaseAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolUuidLeaseAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *UuidpoolUuidLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolUuidLeaseAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolUuidLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolUuidLeaseAllOf) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolUuidLeaseAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolUuidLeaseAllOf) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *UuidpoolUuidLeaseAllOf) GetPoolMember() UuidpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret UuidpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *UuidpoolUuidLeaseAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given UuidpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *UuidpoolUuidLeaseAllOf) SetPoolMember(v UuidpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *UuidpoolUuidLeaseAllOf) GetUniverse() UuidpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret UuidpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLeaseAllOf) GetUniverseOk() (*UuidpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *UuidpoolUuidLeaseAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given UuidpoolUniverseRelationship and assigns it to the Universe field.
func (o *UuidpoolUuidLeaseAllOf) SetUniverse(v UuidpoolUniverseRelationship) {
	o.Universe = &v
}

func (o UuidpoolUuidLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolUuidLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUuidpoolUuidLeaseAllOf := _UuidpoolUuidLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varUuidpoolUuidLeaseAllOf); err == nil {
		*o = UuidpoolUuidLeaseAllOf(varUuidpoolUuidLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUuidpoolUuidLeaseAllOf struct {
	value *UuidpoolUuidLeaseAllOf
	isSet bool
}

func (v NullableUuidpoolUuidLeaseAllOf) Get() *UuidpoolUuidLeaseAllOf {
	return v.value
}

func (v *NullableUuidpoolUuidLeaseAllOf) Set(val *UuidpoolUuidLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolUuidLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolUuidLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolUuidLeaseAllOf(val *UuidpoolUuidLeaseAllOf) *NullableUuidpoolUuidLeaseAllOf {
	return &NullableUuidpoolUuidLeaseAllOf{value: val, isSet: true}
}

func (v NullableUuidpoolUuidLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolUuidLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
