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
	"reflect"
	"strings"
)

// StorageRemoteKeySetting Models the remote key configurarion required for disks encryptions. KMIP is the only remote key protocol supported in the policy.
type StorageRemoteKeySetting struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The password for the KMIP server login.
	Password *string `json:"Password,omitempty"`
	// The port to which the KMIP client should connect.
	Port *int64 `json:"Port,omitempty"`
	// The IP address of the primary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
	PrimaryServer *string `json:"PrimaryServer,omitempty"`
	// The IP address of the secondary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
	SecondaryServer *string `json:"SecondaryServer,omitempty"`
	// The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server.
	ServerCertificate *string `json:"ServerCertificate,omitempty"`
	// The user name for the KMIP server login.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageRemoteKeySetting StorageRemoteKeySetting

// NewStorageRemoteKeySetting instantiates a new StorageRemoteKeySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRemoteKeySetting(classId string, objectType string) *StorageRemoteKeySetting {
	this := StorageRemoteKeySetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	var port int64 = 5696
	this.Port = &port
	return &this
}

// NewStorageRemoteKeySettingWithDefaults instantiates a new StorageRemoteKeySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRemoteKeySettingWithDefaults() *StorageRemoteKeySetting {
	this := StorageRemoteKeySetting{}
	var classId string = "storage.RemoteKeySetting"
	this.ClassId = classId
	var objectType string = "storage.RemoteKeySetting"
	this.ObjectType = objectType
	var port int64 = 5696
	this.Port = &port
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageRemoteKeySetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageRemoteKeySetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageRemoteKeySetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageRemoteKeySetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *StorageRemoteKeySetting) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StorageRemoteKeySetting) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *StorageRemoteKeySetting) SetPort(v int64) {
	o.Port = &v
}

// GetPrimaryServer returns the PrimaryServer field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetPrimaryServer() string {
	if o == nil || o.PrimaryServer == nil {
		var ret string
		return ret
	}
	return *o.PrimaryServer
}

// GetPrimaryServerOk returns a tuple with the PrimaryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetPrimaryServerOk() (*string, bool) {
	if o == nil || o.PrimaryServer == nil {
		return nil, false
	}
	return o.PrimaryServer, true
}

// HasPrimaryServer returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasPrimaryServer() bool {
	if o != nil && o.PrimaryServer != nil {
		return true
	}

	return false
}

// SetPrimaryServer gets a reference to the given string and assigns it to the PrimaryServer field.
func (o *StorageRemoteKeySetting) SetPrimaryServer(v string) {
	o.PrimaryServer = &v
}

// GetSecondaryServer returns the SecondaryServer field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetSecondaryServer() string {
	if o == nil || o.SecondaryServer == nil {
		var ret string
		return ret
	}
	return *o.SecondaryServer
}

// GetSecondaryServerOk returns a tuple with the SecondaryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetSecondaryServerOk() (*string, bool) {
	if o == nil || o.SecondaryServer == nil {
		return nil, false
	}
	return o.SecondaryServer, true
}

// HasSecondaryServer returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasSecondaryServer() bool {
	if o != nil && o.SecondaryServer != nil {
		return true
	}

	return false
}

// SetSecondaryServer gets a reference to the given string and assigns it to the SecondaryServer field.
func (o *StorageRemoteKeySetting) SetSecondaryServer(v string) {
	o.SecondaryServer = &v
}

// GetServerCertificate returns the ServerCertificate field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetServerCertificate() string {
	if o == nil || o.ServerCertificate == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificate
}

// GetServerCertificateOk returns a tuple with the ServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetServerCertificateOk() (*string, bool) {
	if o == nil || o.ServerCertificate == nil {
		return nil, false
	}
	return o.ServerCertificate, true
}

// HasServerCertificate returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasServerCertificate() bool {
	if o != nil && o.ServerCertificate != nil {
		return true
	}

	return false
}

// SetServerCertificate gets a reference to the given string and assigns it to the ServerCertificate field.
func (o *StorageRemoteKeySetting) SetServerCertificate(v string) {
	o.ServerCertificate = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *StorageRemoteKeySetting) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageRemoteKeySetting) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StorageRemoteKeySetting) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StorageRemoteKeySetting) SetUsername(v string) {
	o.Username = &v
}

func (o StorageRemoteKeySetting) MarshalJSON() ([]byte, error) {
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
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.PrimaryServer != nil {
		toSerialize["PrimaryServer"] = o.PrimaryServer
	}
	if o.SecondaryServer != nil {
		toSerialize["SecondaryServer"] = o.SecondaryServer
	}
	if o.ServerCertificate != nil {
		toSerialize["ServerCertificate"] = o.ServerCertificate
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageRemoteKeySetting) UnmarshalJSON(bytes []byte) (err error) {
	type StorageRemoteKeySettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// The password for the KMIP server login.
		Password *string `json:"Password,omitempty"`
		// The port to which the KMIP client should connect.
		Port *int64 `json:"Port,omitempty"`
		// The IP address of the primary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
		PrimaryServer *string `json:"PrimaryServer,omitempty"`
		// The IP address of the secondary KMIP server. It could be an IPv4 address, an IPv6 address, or a hostname. Hostnames are valid only when Inband is configured for the CIMC address.
		SecondaryServer *string `json:"SecondaryServer,omitempty"`
		// The certificate/ public key of the KMIP server. It is required for initiating secure communication with the server.
		ServerCertificate *string `json:"ServerCertificate,omitempty"`
		// The user name for the KMIP server login.
		Username *string `json:"Username,omitempty"`
	}

	varStorageRemoteKeySettingWithoutEmbeddedStruct := StorageRemoteKeySettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageRemoteKeySettingWithoutEmbeddedStruct)
	if err == nil {
		varStorageRemoteKeySetting := _StorageRemoteKeySetting{}
		varStorageRemoteKeySetting.ClassId = varStorageRemoteKeySettingWithoutEmbeddedStruct.ClassId
		varStorageRemoteKeySetting.ObjectType = varStorageRemoteKeySettingWithoutEmbeddedStruct.ObjectType
		varStorageRemoteKeySetting.IsPasswordSet = varStorageRemoteKeySettingWithoutEmbeddedStruct.IsPasswordSet
		varStorageRemoteKeySetting.Password = varStorageRemoteKeySettingWithoutEmbeddedStruct.Password
		varStorageRemoteKeySetting.Port = varStorageRemoteKeySettingWithoutEmbeddedStruct.Port
		varStorageRemoteKeySetting.PrimaryServer = varStorageRemoteKeySettingWithoutEmbeddedStruct.PrimaryServer
		varStorageRemoteKeySetting.SecondaryServer = varStorageRemoteKeySettingWithoutEmbeddedStruct.SecondaryServer
		varStorageRemoteKeySetting.ServerCertificate = varStorageRemoteKeySettingWithoutEmbeddedStruct.ServerCertificate
		varStorageRemoteKeySetting.Username = varStorageRemoteKeySettingWithoutEmbeddedStruct.Username
		*o = StorageRemoteKeySetting(varStorageRemoteKeySetting)
	} else {
		return err
	}

	varStorageRemoteKeySetting := _StorageRemoteKeySetting{}

	err = json.Unmarshal(bytes, &varStorageRemoteKeySetting)
	if err == nil {
		o.MoBaseComplexType = varStorageRemoteKeySetting.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "PrimaryServer")
		delete(additionalProperties, "SecondaryServer")
		delete(additionalProperties, "ServerCertificate")
		delete(additionalProperties, "Username")

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

type NullableStorageRemoteKeySetting struct {
	value *StorageRemoteKeySetting
	isSet bool
}

func (v NullableStorageRemoteKeySetting) Get() *StorageRemoteKeySetting {
	return v.value
}

func (v *NullableStorageRemoteKeySetting) Set(val *StorageRemoteKeySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRemoteKeySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRemoteKeySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRemoteKeySetting(val *StorageRemoteKeySetting) *NullableStorageRemoteKeySetting {
	return &NullableStorageRemoteKeySetting{value: val, isSet: true}
}

func (v NullableStorageRemoteKeySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRemoteKeySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
