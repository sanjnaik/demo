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

// StorageNetAppLunMapAllOf Definition of the list of properties defined in 'storage.NetAppLunMap', excluding properties defined in parent classes.
type StorageNetAppLunMapAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Universally unique identifier of the LUN.
	Uuid *string `json:"Uuid,omitempty"`
	// An array of relationships to storageNetAppInitiatorGroup resources.
	Host                 []StorageNetAppInitiatorGroupRelationship `json:"Host,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship       `json:"Tenant,omitempty"`
	Volume               *StorageNetAppLunRelationship             `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLunMapAllOf StorageNetAppLunMapAllOf

// NewStorageNetAppLunMapAllOf instantiates a new StorageNetAppLunMapAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLunMapAllOf(classId string, objectType string) *StorageNetAppLunMapAllOf {
	this := StorageNetAppLunMapAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunMapAllOfWithDefaults instantiates a new StorageNetAppLunMapAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunMapAllOfWithDefaults() *StorageNetAppLunMapAllOf {
	this := StorageNetAppLunMapAllOf{}
	var classId string = "storage.NetAppLunMap"
	this.ClassId = classId
	var objectType string = "storage.NetAppLunMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLunMapAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMapAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLunMapAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLunMapAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMapAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLunMapAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppLunMapAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMapAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppLunMapAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppLunMapAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLunMapAllOf) GetHost() []StorageNetAppInitiatorGroupRelationship {
	if o == nil {
		var ret []StorageNetAppInitiatorGroupRelationship
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLunMapAllOf) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return &o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *StorageNetAppLunMapAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given []StorageNetAppInitiatorGroupRelationship and assigns it to the Host field.
func (o *StorageNetAppLunMapAllOf) SetHost(v []StorageNetAppInitiatorGroupRelationship) {
	o.Host = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppLunMapAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMapAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppLunMapAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppLunMapAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageNetAppLunMapAllOf) GetVolume() StorageNetAppLunRelationship {
	if o == nil || o.Volume == nil {
		var ret StorageNetAppLunRelationship
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunMapAllOf) GetVolumeOk() (*StorageNetAppLunRelationship, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageNetAppLunMapAllOf) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given StorageNetAppLunRelationship and assigns it to the Volume field.
func (o *StorageNetAppLunMapAllOf) SetVolume(v StorageNetAppLunRelationship) {
	o.Volume = &v
}

func (o StorageNetAppLunMapAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppLunMapAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppLunMapAllOf := _StorageNetAppLunMapAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppLunMapAllOf); err == nil {
		*o = StorageNetAppLunMapAllOf(varStorageNetAppLunMapAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Tenant")
		delete(additionalProperties, "Volume")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppLunMapAllOf struct {
	value *StorageNetAppLunMapAllOf
	isSet bool
}

func (v NullableStorageNetAppLunMapAllOf) Get() *StorageNetAppLunMapAllOf {
	return v.value
}

func (v *NullableStorageNetAppLunMapAllOf) Set(val *StorageNetAppLunMapAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLunMapAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLunMapAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLunMapAllOf(val *StorageNetAppLunMapAllOf) *NullableStorageNetAppLunMapAllOf {
	return &NullableStorageNetAppLunMapAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppLunMapAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLunMapAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
