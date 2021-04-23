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
)

// LicenseLicenseReservationOp Customer operation object to request reservation code.
type LicenseLicenseReservationOp struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Revervation code used to install the license.
	AuthCode *string `json:"AuthCode,omitempty"`
	// Flag to indicate whether authorization code is installed.
	AuthCodeInstalled *bool `json:"AuthCodeInstalled,omitempty"`
	// Confirm code used to complete the license update on smart license account.
	ConfirmCode *string `json:"ConfirmCode,omitempty"`
	// Trigger the generation of request code for specific license reservation.
	GenerateRequestCode *bool `json:"GenerateRequestCode,omitempty"`
	// Trigger the generation of return code for specific license reservation.
	GenerateReturnCode *bool `json:"GenerateReturnCode,omitempty"`
	// Revervation code used to generate authorization code from CSSM.
	RequestCode *string `json:"RequestCode,omitempty"`
	// Return code used to return the reserved license to smart license account.
	ReturnCode           *string                 `json:"ReturnCode,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseLicenseReservationOp LicenseLicenseReservationOp

// NewLicenseLicenseReservationOp instantiates a new LicenseLicenseReservationOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseLicenseReservationOp(classId string, objectType string) *LicenseLicenseReservationOp {
	this := LicenseLicenseReservationOp{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseLicenseReservationOpWithDefaults instantiates a new LicenseLicenseReservationOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseLicenseReservationOpWithDefaults() *LicenseLicenseReservationOp {
	this := LicenseLicenseReservationOp{}
	var classId string = "license.LicenseReservationOp"
	this.ClassId = classId
	var objectType string = "license.LicenseReservationOp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseLicenseReservationOp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseLicenseReservationOp) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseLicenseReservationOp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseLicenseReservationOp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetAuthCode() string {
	if o == nil || o.AuthCode == nil {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetAuthCodeOk() (*string, bool) {
	if o == nil || o.AuthCode == nil {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasAuthCode() bool {
	if o != nil && o.AuthCode != nil {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *LicenseLicenseReservationOp) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetAuthCodeInstalled returns the AuthCodeInstalled field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetAuthCodeInstalled() bool {
	if o == nil || o.AuthCodeInstalled == nil {
		var ret bool
		return ret
	}
	return *o.AuthCodeInstalled
}

// GetAuthCodeInstalledOk returns a tuple with the AuthCodeInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetAuthCodeInstalledOk() (*bool, bool) {
	if o == nil || o.AuthCodeInstalled == nil {
		return nil, false
	}
	return o.AuthCodeInstalled, true
}

// HasAuthCodeInstalled returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasAuthCodeInstalled() bool {
	if o != nil && o.AuthCodeInstalled != nil {
		return true
	}

	return false
}

// SetAuthCodeInstalled gets a reference to the given bool and assigns it to the AuthCodeInstalled field.
func (o *LicenseLicenseReservationOp) SetAuthCodeInstalled(v bool) {
	o.AuthCodeInstalled = &v
}

// GetConfirmCode returns the ConfirmCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetConfirmCode() string {
	if o == nil || o.ConfirmCode == nil {
		var ret string
		return ret
	}
	return *o.ConfirmCode
}

// GetConfirmCodeOk returns a tuple with the ConfirmCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetConfirmCodeOk() (*string, bool) {
	if o == nil || o.ConfirmCode == nil {
		return nil, false
	}
	return o.ConfirmCode, true
}

// HasConfirmCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasConfirmCode() bool {
	if o != nil && o.ConfirmCode != nil {
		return true
	}

	return false
}

// SetConfirmCode gets a reference to the given string and assigns it to the ConfirmCode field.
func (o *LicenseLicenseReservationOp) SetConfirmCode(v string) {
	o.ConfirmCode = &v
}

// GetGenerateRequestCode returns the GenerateRequestCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetGenerateRequestCode() bool {
	if o == nil || o.GenerateRequestCode == nil {
		var ret bool
		return ret
	}
	return *o.GenerateRequestCode
}

// GetGenerateRequestCodeOk returns a tuple with the GenerateRequestCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetGenerateRequestCodeOk() (*bool, bool) {
	if o == nil || o.GenerateRequestCode == nil {
		return nil, false
	}
	return o.GenerateRequestCode, true
}

// HasGenerateRequestCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasGenerateRequestCode() bool {
	if o != nil && o.GenerateRequestCode != nil {
		return true
	}

	return false
}

// SetGenerateRequestCode gets a reference to the given bool and assigns it to the GenerateRequestCode field.
func (o *LicenseLicenseReservationOp) SetGenerateRequestCode(v bool) {
	o.GenerateRequestCode = &v
}

// GetGenerateReturnCode returns the GenerateReturnCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetGenerateReturnCode() bool {
	if o == nil || o.GenerateReturnCode == nil {
		var ret bool
		return ret
	}
	return *o.GenerateReturnCode
}

// GetGenerateReturnCodeOk returns a tuple with the GenerateReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetGenerateReturnCodeOk() (*bool, bool) {
	if o == nil || o.GenerateReturnCode == nil {
		return nil, false
	}
	return o.GenerateReturnCode, true
}

// HasGenerateReturnCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasGenerateReturnCode() bool {
	if o != nil && o.GenerateReturnCode != nil {
		return true
	}

	return false
}

// SetGenerateReturnCode gets a reference to the given bool and assigns it to the GenerateReturnCode field.
func (o *LicenseLicenseReservationOp) SetGenerateReturnCode(v bool) {
	o.GenerateReturnCode = &v
}

// GetRequestCode returns the RequestCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetRequestCode() string {
	if o == nil || o.RequestCode == nil {
		var ret string
		return ret
	}
	return *o.RequestCode
}

// GetRequestCodeOk returns a tuple with the RequestCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetRequestCodeOk() (*string, bool) {
	if o == nil || o.RequestCode == nil {
		return nil, false
	}
	return o.RequestCode, true
}

// HasRequestCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasRequestCode() bool {
	if o != nil && o.RequestCode != nil {
		return true
	}

	return false
}

// SetRequestCode gets a reference to the given string and assigns it to the RequestCode field.
func (o *LicenseLicenseReservationOp) SetRequestCode(v string) {
	o.RequestCode = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetReturnCode() string {
	if o == nil || o.ReturnCode == nil {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetReturnCodeOk() (*string, bool) {
	if o == nil || o.ReturnCode == nil {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasReturnCode() bool {
	if o != nil && o.ReturnCode != nil {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *LicenseLicenseReservationOp) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *LicenseLicenseReservationOp) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseReservationOp) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *LicenseLicenseReservationOp) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *LicenseLicenseReservationOp) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o LicenseLicenseReservationOp) MarshalJSON() ([]byte, error) {
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
	if o.AuthCode != nil {
		toSerialize["AuthCode"] = o.AuthCode
	}
	if o.AuthCodeInstalled != nil {
		toSerialize["AuthCodeInstalled"] = o.AuthCodeInstalled
	}
	if o.ConfirmCode != nil {
		toSerialize["ConfirmCode"] = o.ConfirmCode
	}
	if o.GenerateRequestCode != nil {
		toSerialize["GenerateRequestCode"] = o.GenerateRequestCode
	}
	if o.GenerateReturnCode != nil {
		toSerialize["GenerateReturnCode"] = o.GenerateReturnCode
	}
	if o.RequestCode != nil {
		toSerialize["RequestCode"] = o.RequestCode
	}
	if o.ReturnCode != nil {
		toSerialize["ReturnCode"] = o.ReturnCode
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseLicenseReservationOp) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseLicenseReservationOpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Revervation code used to install the license.
		AuthCode *string `json:"AuthCode,omitempty"`
		// Flag to indicate whether authorization code is installed.
		AuthCodeInstalled *bool `json:"AuthCodeInstalled,omitempty"`
		// Confirm code used to complete the license update on smart license account.
		ConfirmCode *string `json:"ConfirmCode,omitempty"`
		// Trigger the generation of request code for specific license reservation.
		GenerateRequestCode *bool `json:"GenerateRequestCode,omitempty"`
		// Trigger the generation of return code for specific license reservation.
		GenerateReturnCode *bool `json:"GenerateReturnCode,omitempty"`
		// Revervation code used to generate authorization code from CSSM.
		RequestCode *string `json:"RequestCode,omitempty"`
		// Return code used to return the reserved license to smart license account.
		ReturnCode *string                 `json:"ReturnCode,omitempty"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varLicenseLicenseReservationOpWithoutEmbeddedStruct := LicenseLicenseReservationOpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseLicenseReservationOpWithoutEmbeddedStruct)
	if err == nil {
		varLicenseLicenseReservationOp := _LicenseLicenseReservationOp{}
		varLicenseLicenseReservationOp.ClassId = varLicenseLicenseReservationOpWithoutEmbeddedStruct.ClassId
		varLicenseLicenseReservationOp.ObjectType = varLicenseLicenseReservationOpWithoutEmbeddedStruct.ObjectType
		varLicenseLicenseReservationOp.AuthCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.AuthCode
		varLicenseLicenseReservationOp.AuthCodeInstalled = varLicenseLicenseReservationOpWithoutEmbeddedStruct.AuthCodeInstalled
		varLicenseLicenseReservationOp.ConfirmCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.ConfirmCode
		varLicenseLicenseReservationOp.GenerateRequestCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.GenerateRequestCode
		varLicenseLicenseReservationOp.GenerateReturnCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.GenerateReturnCode
		varLicenseLicenseReservationOp.RequestCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.RequestCode
		varLicenseLicenseReservationOp.ReturnCode = varLicenseLicenseReservationOpWithoutEmbeddedStruct.ReturnCode
		varLicenseLicenseReservationOp.Account = varLicenseLicenseReservationOpWithoutEmbeddedStruct.Account
		*o = LicenseLicenseReservationOp(varLicenseLicenseReservationOp)
	} else {
		return err
	}

	varLicenseLicenseReservationOp := _LicenseLicenseReservationOp{}

	err = json.Unmarshal(bytes, &varLicenseLicenseReservationOp)
	if err == nil {
		o.MoBaseMo = varLicenseLicenseReservationOp.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthCode")
		delete(additionalProperties, "AuthCodeInstalled")
		delete(additionalProperties, "ConfirmCode")
		delete(additionalProperties, "GenerateRequestCode")
		delete(additionalProperties, "GenerateReturnCode")
		delete(additionalProperties, "RequestCode")
		delete(additionalProperties, "ReturnCode")
		delete(additionalProperties, "Account")

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

type NullableLicenseLicenseReservationOp struct {
	value *LicenseLicenseReservationOp
	isSet bool
}

func (v NullableLicenseLicenseReservationOp) Get() *LicenseLicenseReservationOp {
	return v.value
}

func (v *NullableLicenseLicenseReservationOp) Set(val *LicenseLicenseReservationOp) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseLicenseReservationOp) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseLicenseReservationOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseLicenseReservationOp(val *LicenseLicenseReservationOp) *NullableLicenseLicenseReservationOp {
	return &NullableLicenseLicenseReservationOp{value: val, isSet: true}
}

func (v NullableLicenseLicenseReservationOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseLicenseReservationOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
