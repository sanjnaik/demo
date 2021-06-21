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

// WorkflowFileDownloadOp File operation to download a given file from Intersight storage services such as AWS or Minio bucket to a specified path on one or more Intersight connected devices.
type WorkflowFileDownloadOp struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Path on the Intersight connected device to which file needs to be downloaded.
	DestinationPath *string `json:"DestinationPath,omitempty"`
	// Source bucket name hosting the file.
	SourceBucket *string `json:"SourceBucket,omitempty"`
	// Name of the file to be downloaded from bucket to endpoint devices.
	SourceFile           *string `json:"SourceFile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowFileDownloadOp WorkflowFileDownloadOp

// NewWorkflowFileDownloadOp instantiates a new WorkflowFileDownloadOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFileDownloadOp(classId string, objectType string) *WorkflowFileDownloadOp {
	this := WorkflowFileDownloadOp{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowFileDownloadOpWithDefaults instantiates a new WorkflowFileDownloadOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFileDownloadOpWithDefaults() *WorkflowFileDownloadOp {
	this := WorkflowFileDownloadOp{}
	var classId string = "workflow.FileDownloadOp"
	this.ClassId = classId
	var objectType string = "workflow.FileDownloadOp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowFileDownloadOp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileDownloadOp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowFileDownloadOp) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowFileDownloadOp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileDownloadOp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowFileDownloadOp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDestinationPath returns the DestinationPath field value if set, zero value otherwise.
func (o *WorkflowFileDownloadOp) GetDestinationPath() string {
	if o == nil || o.DestinationPath == nil {
		var ret string
		return ret
	}
	return *o.DestinationPath
}

// GetDestinationPathOk returns a tuple with the DestinationPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileDownloadOp) GetDestinationPathOk() (*string, bool) {
	if o == nil || o.DestinationPath == nil {
		return nil, false
	}
	return o.DestinationPath, true
}

// HasDestinationPath returns a boolean if a field has been set.
func (o *WorkflowFileDownloadOp) HasDestinationPath() bool {
	if o != nil && o.DestinationPath != nil {
		return true
	}

	return false
}

// SetDestinationPath gets a reference to the given string and assigns it to the DestinationPath field.
func (o *WorkflowFileDownloadOp) SetDestinationPath(v string) {
	o.DestinationPath = &v
}

// GetSourceBucket returns the SourceBucket field value if set, zero value otherwise.
func (o *WorkflowFileDownloadOp) GetSourceBucket() string {
	if o == nil || o.SourceBucket == nil {
		var ret string
		return ret
	}
	return *o.SourceBucket
}

// GetSourceBucketOk returns a tuple with the SourceBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileDownloadOp) GetSourceBucketOk() (*string, bool) {
	if o == nil || o.SourceBucket == nil {
		return nil, false
	}
	return o.SourceBucket, true
}

// HasSourceBucket returns a boolean if a field has been set.
func (o *WorkflowFileDownloadOp) HasSourceBucket() bool {
	if o != nil && o.SourceBucket != nil {
		return true
	}

	return false
}

// SetSourceBucket gets a reference to the given string and assigns it to the SourceBucket field.
func (o *WorkflowFileDownloadOp) SetSourceBucket(v string) {
	o.SourceBucket = &v
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise.
func (o *WorkflowFileDownloadOp) GetSourceFile() string {
	if o == nil || o.SourceFile == nil {
		var ret string
		return ret
	}
	return *o.SourceFile
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileDownloadOp) GetSourceFileOk() (*string, bool) {
	if o == nil || o.SourceFile == nil {
		return nil, false
	}
	return o.SourceFile, true
}

// HasSourceFile returns a boolean if a field has been set.
func (o *WorkflowFileDownloadOp) HasSourceFile() bool {
	if o != nil && o.SourceFile != nil {
		return true
	}

	return false
}

// SetSourceFile gets a reference to the given string and assigns it to the SourceFile field.
func (o *WorkflowFileDownloadOp) SetSourceFile(v string) {
	o.SourceFile = &v
}

func (o WorkflowFileDownloadOp) MarshalJSON() ([]byte, error) {
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
	if o.DestinationPath != nil {
		toSerialize["DestinationPath"] = o.DestinationPath
	}
	if o.SourceBucket != nil {
		toSerialize["SourceBucket"] = o.SourceBucket
	}
	if o.SourceFile != nil {
		toSerialize["SourceFile"] = o.SourceFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowFileDownloadOp) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowFileDownloadOpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Path on the Intersight connected device to which file needs to be downloaded.
		DestinationPath *string `json:"DestinationPath,omitempty"`
		// Source bucket name hosting the file.
		SourceBucket *string `json:"SourceBucket,omitempty"`
		// Name of the file to be downloaded from bucket to endpoint devices.
		SourceFile *string `json:"SourceFile,omitempty"`
	}

	varWorkflowFileDownloadOpWithoutEmbeddedStruct := WorkflowFileDownloadOpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowFileDownloadOpWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowFileDownloadOp := _WorkflowFileDownloadOp{}
		varWorkflowFileDownloadOp.ClassId = varWorkflowFileDownloadOpWithoutEmbeddedStruct.ClassId
		varWorkflowFileDownloadOp.ObjectType = varWorkflowFileDownloadOpWithoutEmbeddedStruct.ObjectType
		varWorkflowFileDownloadOp.DestinationPath = varWorkflowFileDownloadOpWithoutEmbeddedStruct.DestinationPath
		varWorkflowFileDownloadOp.SourceBucket = varWorkflowFileDownloadOpWithoutEmbeddedStruct.SourceBucket
		varWorkflowFileDownloadOp.SourceFile = varWorkflowFileDownloadOpWithoutEmbeddedStruct.SourceFile
		*o = WorkflowFileDownloadOp(varWorkflowFileDownloadOp)
	} else {
		return err
	}

	varWorkflowFileDownloadOp := _WorkflowFileDownloadOp{}

	err = json.Unmarshal(bytes, &varWorkflowFileDownloadOp)
	if err == nil {
		o.MoBaseComplexType = varWorkflowFileDownloadOp.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestinationPath")
		delete(additionalProperties, "SourceBucket")
		delete(additionalProperties, "SourceFile")

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

type NullableWorkflowFileDownloadOp struct {
	value *WorkflowFileDownloadOp
	isSet bool
}

func (v NullableWorkflowFileDownloadOp) Get() *WorkflowFileDownloadOp {
	return v.value
}

func (v *NullableWorkflowFileDownloadOp) Set(val *WorkflowFileDownloadOp) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFileDownloadOp) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFileDownloadOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFileDownloadOp(val *WorkflowFileDownloadOp) *NullableWorkflowFileDownloadOp {
	return &NullableWorkflowFileDownloadOp{value: val, isSet: true}
}

func (v NullableWorkflowFileDownloadOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFileDownloadOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
