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

// WorkflowSshSession This models a single SSH session from Intersight connected endpoint to a remote server. Multiple SSH operations can be run sequentially over a single SSH session.
type WorkflowSshSession struct {
	WorkflowApi
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                `json:"ObjectType"`
	FileTransferToRemote *WorkflowFileTransfer `json:"FileTransferToRemote,omitempty"`
	// The type of SSH message to send to the remote server. * `ExecuteCommand` - Execute a SSH command on the remote server. * `NewSession` - Open a new SSH connection to the remote server. * `FileTransfer` - Transfer a file from Intersight connected device to the remote server. * `CloseSession` - Close the SSH connection to the remote server.
	MessageType          *string            `json:"MessageType,omitempty"`
	SshCommand           *WorkflowSshCmd    `json:"SshCommand,omitempty"`
	SshConfiguration     *WorkflowSshConfig `json:"SshConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSshSession WorkflowSshSession

// NewWorkflowSshSession instantiates a new WorkflowSshSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSshSession(classId string, objectType string) *WorkflowSshSession {
	this := WorkflowSshSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	var messageType string = "ExecuteCommand"
	this.MessageType = &messageType
	return &this
}

// NewWorkflowSshSessionWithDefaults instantiates a new WorkflowSshSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSshSessionWithDefaults() *WorkflowSshSession {
	this := WorkflowSshSession{}
	var classId string = "workflow.SshSession"
	this.ClassId = classId
	var objectType string = "workflow.SshSession"
	this.ObjectType = objectType
	var messageType string = "ExecuteCommand"
	this.MessageType = &messageType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSshSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSshSession) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSshSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSshSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileTransferToRemote returns the FileTransferToRemote field value if set, zero value otherwise.
func (o *WorkflowSshSession) GetFileTransferToRemote() WorkflowFileTransfer {
	if o == nil || o.FileTransferToRemote == nil {
		var ret WorkflowFileTransfer
		return ret
	}
	return *o.FileTransferToRemote
}

// GetFileTransferToRemoteOk returns a tuple with the FileTransferToRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetFileTransferToRemoteOk() (*WorkflowFileTransfer, bool) {
	if o == nil || o.FileTransferToRemote == nil {
		return nil, false
	}
	return o.FileTransferToRemote, true
}

// HasFileTransferToRemote returns a boolean if a field has been set.
func (o *WorkflowSshSession) HasFileTransferToRemote() bool {
	if o != nil && o.FileTransferToRemote != nil {
		return true
	}

	return false
}

// SetFileTransferToRemote gets a reference to the given WorkflowFileTransfer and assigns it to the FileTransferToRemote field.
func (o *WorkflowSshSession) SetFileTransferToRemote(v WorkflowFileTransfer) {
	o.FileTransferToRemote = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *WorkflowSshSession) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *WorkflowSshSession) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *WorkflowSshSession) SetMessageType(v string) {
	o.MessageType = &v
}

// GetSshCommand returns the SshCommand field value if set, zero value otherwise.
func (o *WorkflowSshSession) GetSshCommand() WorkflowSshCmd {
	if o == nil || o.SshCommand == nil {
		var ret WorkflowSshCmd
		return ret
	}
	return *o.SshCommand
}

// GetSshCommandOk returns a tuple with the SshCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetSshCommandOk() (*WorkflowSshCmd, bool) {
	if o == nil || o.SshCommand == nil {
		return nil, false
	}
	return o.SshCommand, true
}

// HasSshCommand returns a boolean if a field has been set.
func (o *WorkflowSshSession) HasSshCommand() bool {
	if o != nil && o.SshCommand != nil {
		return true
	}

	return false
}

// SetSshCommand gets a reference to the given WorkflowSshCmd and assigns it to the SshCommand field.
func (o *WorkflowSshSession) SetSshCommand(v WorkflowSshCmd) {
	o.SshCommand = &v
}

// GetSshConfiguration returns the SshConfiguration field value if set, zero value otherwise.
func (o *WorkflowSshSession) GetSshConfiguration() WorkflowSshConfig {
	if o == nil || o.SshConfiguration == nil {
		var ret WorkflowSshConfig
		return ret
	}
	return *o.SshConfiguration
}

// GetSshConfigurationOk returns a tuple with the SshConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSession) GetSshConfigurationOk() (*WorkflowSshConfig, bool) {
	if o == nil || o.SshConfiguration == nil {
		return nil, false
	}
	return o.SshConfiguration, true
}

// HasSshConfiguration returns a boolean if a field has been set.
func (o *WorkflowSshSession) HasSshConfiguration() bool {
	if o != nil && o.SshConfiguration != nil {
		return true
	}

	return false
}

// SetSshConfiguration gets a reference to the given WorkflowSshConfig and assigns it to the SshConfiguration field.
func (o *WorkflowSshSession) SetSshConfiguration(v WorkflowSshConfig) {
	o.SshConfiguration = &v
}

func (o WorkflowSshSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileTransferToRemote != nil {
		toSerialize["FileTransferToRemote"] = o.FileTransferToRemote
	}
	if o.MessageType != nil {
		toSerialize["MessageType"] = o.MessageType
	}
	if o.SshCommand != nil {
		toSerialize["SshCommand"] = o.SshCommand
	}
	if o.SshConfiguration != nil {
		toSerialize["SshConfiguration"] = o.SshConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSshSession) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowSshSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                `json:"ObjectType"`
		FileTransferToRemote *WorkflowFileTransfer `json:"FileTransferToRemote,omitempty"`
		// The type of SSH message to send to the remote server. * `ExecuteCommand` - Execute a SSH command on the remote server. * `NewSession` - Open a new SSH connection to the remote server. * `FileTransfer` - Transfer a file from Intersight connected device to the remote server. * `CloseSession` - Close the SSH connection to the remote server.
		MessageType      *string            `json:"MessageType,omitempty"`
		SshCommand       *WorkflowSshCmd    `json:"SshCommand,omitempty"`
		SshConfiguration *WorkflowSshConfig `json:"SshConfiguration,omitempty"`
	}

	varWorkflowSshSessionWithoutEmbeddedStruct := WorkflowSshSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowSshSessionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSshSession := _WorkflowSshSession{}
		varWorkflowSshSession.ClassId = varWorkflowSshSessionWithoutEmbeddedStruct.ClassId
		varWorkflowSshSession.ObjectType = varWorkflowSshSessionWithoutEmbeddedStruct.ObjectType
		varWorkflowSshSession.FileTransferToRemote = varWorkflowSshSessionWithoutEmbeddedStruct.FileTransferToRemote
		varWorkflowSshSession.MessageType = varWorkflowSshSessionWithoutEmbeddedStruct.MessageType
		varWorkflowSshSession.SshCommand = varWorkflowSshSessionWithoutEmbeddedStruct.SshCommand
		varWorkflowSshSession.SshConfiguration = varWorkflowSshSessionWithoutEmbeddedStruct.SshConfiguration
		*o = WorkflowSshSession(varWorkflowSshSession)
	} else {
		return err
	}

	varWorkflowSshSession := _WorkflowSshSession{}

	err = json.Unmarshal(bytes, &varWorkflowSshSession)
	if err == nil {
		o.WorkflowApi = varWorkflowSshSession.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileTransferToRemote")
		delete(additionalProperties, "MessageType")
		delete(additionalProperties, "SshCommand")
		delete(additionalProperties, "SshConfiguration")

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowSshSession struct {
	value *WorkflowSshSession
	isSet bool
}

func (v NullableWorkflowSshSession) Get() *WorkflowSshSession {
	return v.value
}

func (v *NullableWorkflowSshSession) Set(val *WorkflowSshSession) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSshSession) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSshSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSshSession(val *WorkflowSshSession) *NullableWorkflowSshSession {
	return &NullableWorkflowSshSession{value: val, isSet: true}
}

func (v NullableWorkflowSshSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSshSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
