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
)

// SoftwarerepositoryAuthorizationAllOf Definition of the list of properties defined in 'softwarerepository.Authorization', excluding properties defined in parent classes.
type SoftwarerepositoryAuthorizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Indicates whether the value of the 'userId' property has been set.
	IsUserIdSet *bool `json:"IsUserIdSet,omitempty"`
	// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	Password *string `json:"Password,omitempty"`
	// The external repository for which this authorization has been provided. The only supported repository today is cisco.com. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
	RepositoryType *string `json:"RepositoryType,omitempty"`
	// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	UserId               *string                 `json:"UserId,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryAuthorizationAllOf SoftwarerepositoryAuthorizationAllOf

// NewSoftwarerepositoryAuthorizationAllOf instantiates a new SoftwarerepositoryAuthorizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryAuthorizationAllOf(classId string, objectType string) *SoftwarerepositoryAuthorizationAllOf {
	this := SoftwarerepositoryAuthorizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var isUserIdSet bool = false
	this.IsUserIdSet = &isUserIdSet
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// NewSoftwarerepositoryAuthorizationAllOfWithDefaults instantiates a new SoftwarerepositoryAuthorizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryAuthorizationAllOfWithDefaults() *SoftwarerepositoryAuthorizationAllOf {
	this := SoftwarerepositoryAuthorizationAllOf{}
	var classId string = "softwarerepository.Authorization"
	this.ClassId = classId
	var objectType string = "softwarerepository.Authorization"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var isUserIdSet bool = false
	this.IsUserIdSet = &isUserIdSet
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryAuthorizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryAuthorizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryAuthorizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryAuthorizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsUserIdSet returns the IsUserIdSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetIsUserIdSet() bool {
	if o == nil || o.IsUserIdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsUserIdSet
}

// GetIsUserIdSetOk returns a tuple with the IsUserIdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetIsUserIdSetOk() (*bool, bool) {
	if o == nil || o.IsUserIdSet == nil {
		return nil, false
	}
	return o.IsUserIdSet, true
}

// HasIsUserIdSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasIsUserIdSet() bool {
	if o != nil && o.IsUserIdSet != nil {
		return true
	}

	return false
}

// SetIsUserIdSet gets a reference to the given bool and assigns it to the IsUserIdSet field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetIsUserIdSet(v bool) {
	o.IsUserIdSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetRepositoryType returns the RepositoryType field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetRepositoryType() string {
	if o == nil || o.RepositoryType == nil {
		var ret string
		return ret
	}
	return *o.RepositoryType
}

// GetRepositoryTypeOk returns a tuple with the RepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetRepositoryTypeOk() (*string, bool) {
	if o == nil || o.RepositoryType == nil {
		return nil, false
	}
	return o.RepositoryType, true
}

// HasRepositoryType returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasRepositoryType() bool {
	if o != nil && o.RepositoryType != nil {
		return true
	}

	return false
}

// SetRepositoryType gets a reference to the given string and assigns it to the RepositoryType field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetRepositoryType(v string) {
	o.RepositoryType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetUserId(v string) {
	o.UserId = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorizationAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorizationAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *SoftwarerepositoryAuthorizationAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o SoftwarerepositoryAuthorizationAllOf) MarshalJSON() ([]byte, error) {
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
	if o.IsUserIdSet != nil {
		toSerialize["IsUserIdSet"] = o.IsUserIdSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RepositoryType != nil {
		toSerialize["RepositoryType"] = o.RepositoryType
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryAuthorizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryAuthorizationAllOf := _SoftwarerepositoryAuthorizationAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryAuthorizationAllOf); err == nil {
		*o = SoftwarerepositoryAuthorizationAllOf(varSoftwarerepositoryAuthorizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "IsUserIdSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "RepositoryType")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryAuthorizationAllOf struct {
	value *SoftwarerepositoryAuthorizationAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryAuthorizationAllOf) Get() *SoftwarerepositoryAuthorizationAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryAuthorizationAllOf) Set(val *SoftwarerepositoryAuthorizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryAuthorizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryAuthorizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryAuthorizationAllOf(val *SoftwarerepositoryAuthorizationAllOf) *NullableSoftwarerepositoryAuthorizationAllOf {
	return &NullableSoftwarerepositoryAuthorizationAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryAuthorizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryAuthorizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
