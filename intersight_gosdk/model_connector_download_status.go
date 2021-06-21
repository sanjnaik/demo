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
	"reflect"
	"strings"
)

// ConnectorDownloadStatus The status of a download operation. The download plugin will report the status of a download operation periodically to report status or to report completion/error.
type ConnectorDownloadStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                        `json:"ObjectType"`
	Checksum   NullableConnectorFileChecksum `json:"Checksum,omitempty"`
	// Any error encountered. Set to empty when download is in progress or completed.
	DownloadError *string `json:"DownloadError,omitempty"`
	// The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.
	DownloadProgress *int64 `json:"DownloadProgress,omitempty"`
	// The number of retries the plugin attempted before succeeding or failing the download.
	DownloadRetries *int64 `json:"DownloadRetries,omitempty"`
	// The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file.
	Sha256checksum       *string `json:"Sha256checksum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorDownloadStatus ConnectorDownloadStatus

// NewConnectorDownloadStatus instantiates a new ConnectorDownloadStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorDownloadStatus(classId string, objectType string) *ConnectorDownloadStatus {
	this := ConnectorDownloadStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorDownloadStatusWithDefaults instantiates a new ConnectorDownloadStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorDownloadStatusWithDefaults() *ConnectorDownloadStatus {
	this := ConnectorDownloadStatus{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorDownloadStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorDownloadStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorDownloadStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorDownloadStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorDownloadStatus) GetChecksum() ConnectorFileChecksum {
	if o == nil || o.Checksum.Get() == nil {
		var ret ConnectorFileChecksum
		return ret
	}
	return *o.Checksum.Get()
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorDownloadStatus) GetChecksumOk() (*ConnectorFileChecksum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checksum.Get(), o.Checksum.IsSet()
}

// HasChecksum returns a boolean if a field has been set.
func (o *ConnectorDownloadStatus) HasChecksum() bool {
	if o != nil && o.Checksum.IsSet() {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given NullableConnectorFileChecksum and assigns it to the Checksum field.
func (o *ConnectorDownloadStatus) SetChecksum(v ConnectorFileChecksum) {
	o.Checksum.Set(&v)
}

// SetChecksumNil sets the value for Checksum to be an explicit nil
func (o *ConnectorDownloadStatus) SetChecksumNil() {
	o.Checksum.Set(nil)
}

// UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
func (o *ConnectorDownloadStatus) UnsetChecksum() {
	o.Checksum.Unset()
}

// GetDownloadError returns the DownloadError field value if set, zero value otherwise.
func (o *ConnectorDownloadStatus) GetDownloadError() string {
	if o == nil || o.DownloadError == nil {
		var ret string
		return ret
	}
	return *o.DownloadError
}

// GetDownloadErrorOk returns a tuple with the DownloadError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetDownloadErrorOk() (*string, bool) {
	if o == nil || o.DownloadError == nil {
		return nil, false
	}
	return o.DownloadError, true
}

// HasDownloadError returns a boolean if a field has been set.
func (o *ConnectorDownloadStatus) HasDownloadError() bool {
	if o != nil && o.DownloadError != nil {
		return true
	}

	return false
}

// SetDownloadError gets a reference to the given string and assigns it to the DownloadError field.
func (o *ConnectorDownloadStatus) SetDownloadError(v string) {
	o.DownloadError = &v
}

// GetDownloadProgress returns the DownloadProgress field value if set, zero value otherwise.
func (o *ConnectorDownloadStatus) GetDownloadProgress() int64 {
	if o == nil || o.DownloadProgress == nil {
		var ret int64
		return ret
	}
	return *o.DownloadProgress
}

// GetDownloadProgressOk returns a tuple with the DownloadProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetDownloadProgressOk() (*int64, bool) {
	if o == nil || o.DownloadProgress == nil {
		return nil, false
	}
	return o.DownloadProgress, true
}

// HasDownloadProgress returns a boolean if a field has been set.
func (o *ConnectorDownloadStatus) HasDownloadProgress() bool {
	if o != nil && o.DownloadProgress != nil {
		return true
	}

	return false
}

// SetDownloadProgress gets a reference to the given int64 and assigns it to the DownloadProgress field.
func (o *ConnectorDownloadStatus) SetDownloadProgress(v int64) {
	o.DownloadProgress = &v
}

// GetDownloadRetries returns the DownloadRetries field value if set, zero value otherwise.
func (o *ConnectorDownloadStatus) GetDownloadRetries() int64 {
	if o == nil || o.DownloadRetries == nil {
		var ret int64
		return ret
	}
	return *o.DownloadRetries
}

// GetDownloadRetriesOk returns a tuple with the DownloadRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetDownloadRetriesOk() (*int64, bool) {
	if o == nil || o.DownloadRetries == nil {
		return nil, false
	}
	return o.DownloadRetries, true
}

// HasDownloadRetries returns a boolean if a field has been set.
func (o *ConnectorDownloadStatus) HasDownloadRetries() bool {
	if o != nil && o.DownloadRetries != nil {
		return true
	}

	return false
}

// SetDownloadRetries gets a reference to the given int64 and assigns it to the DownloadRetries field.
func (o *ConnectorDownloadStatus) SetDownloadRetries(v int64) {
	o.DownloadRetries = &v
}

// GetSha256checksum returns the Sha256checksum field value if set, zero value otherwise.
func (o *ConnectorDownloadStatus) GetSha256checksum() string {
	if o == nil || o.Sha256checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha256checksum
}

// GetSha256checksumOk returns a tuple with the Sha256checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatus) GetSha256checksumOk() (*string, bool) {
	if o == nil || o.Sha256checksum == nil {
		return nil, false
	}
	return o.Sha256checksum, true
}

// HasSha256checksum returns a boolean if a field has been set.
func (o *ConnectorDownloadStatus) HasSha256checksum() bool {
	if o != nil && o.Sha256checksum != nil {
		return true
	}

	return false
}

// SetSha256checksum gets a reference to the given string and assigns it to the Sha256checksum field.
func (o *ConnectorDownloadStatus) SetSha256checksum(v string) {
	o.Sha256checksum = &v
}

func (o ConnectorDownloadStatus) MarshalJSON() ([]byte, error) {
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
	if o.Checksum.IsSet() {
		toSerialize["Checksum"] = o.Checksum.Get()
	}
	if o.DownloadError != nil {
		toSerialize["DownloadError"] = o.DownloadError
	}
	if o.DownloadProgress != nil {
		toSerialize["DownloadProgress"] = o.DownloadProgress
	}
	if o.DownloadRetries != nil {
		toSerialize["DownloadRetries"] = o.DownloadRetries
	}
	if o.Sha256checksum != nil {
		toSerialize["Sha256checksum"] = o.Sha256checksum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorDownloadStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorDownloadStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                        `json:"ObjectType"`
		Checksum   NullableConnectorFileChecksum `json:"Checksum,omitempty"`
		// Any error encountered. Set to empty when download is in progress or completed.
		DownloadError *string `json:"DownloadError,omitempty"`
		// The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.
		DownloadProgress *int64 `json:"DownloadProgress,omitempty"`
		// The number of retries the plugin attempted before succeeding or failing the download.
		DownloadRetries *int64 `json:"DownloadRetries,omitempty"`
		// The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file.
		Sha256checksum *string `json:"Sha256checksum,omitempty"`
	}

	varConnectorDownloadStatusWithoutEmbeddedStruct := ConnectorDownloadStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorDownloadStatusWithoutEmbeddedStruct)
	if err == nil {
		varConnectorDownloadStatus := _ConnectorDownloadStatus{}
		varConnectorDownloadStatus.ClassId = varConnectorDownloadStatusWithoutEmbeddedStruct.ClassId
		varConnectorDownloadStatus.ObjectType = varConnectorDownloadStatusWithoutEmbeddedStruct.ObjectType
		varConnectorDownloadStatus.Checksum = varConnectorDownloadStatusWithoutEmbeddedStruct.Checksum
		varConnectorDownloadStatus.DownloadError = varConnectorDownloadStatusWithoutEmbeddedStruct.DownloadError
		varConnectorDownloadStatus.DownloadProgress = varConnectorDownloadStatusWithoutEmbeddedStruct.DownloadProgress
		varConnectorDownloadStatus.DownloadRetries = varConnectorDownloadStatusWithoutEmbeddedStruct.DownloadRetries
		varConnectorDownloadStatus.Sha256checksum = varConnectorDownloadStatusWithoutEmbeddedStruct.Sha256checksum
		*o = ConnectorDownloadStatus(varConnectorDownloadStatus)
	} else {
		return err
	}

	varConnectorDownloadStatus := _ConnectorDownloadStatus{}

	err = json.Unmarshal(bytes, &varConnectorDownloadStatus)
	if err == nil {
		o.MoBaseMo = varConnectorDownloadStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Checksum")
		delete(additionalProperties, "DownloadError")
		delete(additionalProperties, "DownloadProgress")
		delete(additionalProperties, "DownloadRetries")
		delete(additionalProperties, "Sha256checksum")

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

type NullableConnectorDownloadStatus struct {
	value *ConnectorDownloadStatus
	isSet bool
}

func (v NullableConnectorDownloadStatus) Get() *ConnectorDownloadStatus {
	return v.value
}

func (v *NullableConnectorDownloadStatus) Set(val *ConnectorDownloadStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorDownloadStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorDownloadStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorDownloadStatus(val *ConnectorDownloadStatus) *NullableConnectorDownloadStatus {
	return &NullableConnectorDownloadStatus{value: val, isSet: true}
}

func (v NullableConnectorDownloadStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorDownloadStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
