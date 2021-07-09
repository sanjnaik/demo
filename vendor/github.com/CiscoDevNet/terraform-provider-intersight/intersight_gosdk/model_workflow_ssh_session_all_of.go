/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowSshSessionAllOf Definition of the list of properties defined in 'workflow.SshSession', excluding properties defined in parent classes.
type WorkflowSshSessionAllOf struct {
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

type _WorkflowSshSessionAllOf WorkflowSshSessionAllOf

// NewWorkflowSshSessionAllOf instantiates a new WorkflowSshSessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSshSessionAllOf(classId string, objectType string) *WorkflowSshSessionAllOf {
	this := WorkflowSshSessionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var messageType string = "ExecuteCommand"
	this.MessageType = &messageType
	return &this
}

// NewWorkflowSshSessionAllOfWithDefaults instantiates a new WorkflowSshSessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSshSessionAllOfWithDefaults() *WorkflowSshSessionAllOf {
	this := WorkflowSshSessionAllOf{}
	var classId string = "workflow.SshSession"
	this.ClassId = classId
	var objectType string = "workflow.SshSession"
	this.ObjectType = objectType
	var messageType string = "ExecuteCommand"
	this.MessageType = &messageType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSshSessionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSshSessionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSshSessionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSshSessionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileTransferToRemote returns the FileTransferToRemote field value if set, zero value otherwise.
func (o *WorkflowSshSessionAllOf) GetFileTransferToRemote() WorkflowFileTransfer {
	if o == nil || o.FileTransferToRemote == nil {
		var ret WorkflowFileTransfer
		return ret
	}
	return *o.FileTransferToRemote
}

// GetFileTransferToRemoteOk returns a tuple with the FileTransferToRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetFileTransferToRemoteOk() (*WorkflowFileTransfer, bool) {
	if o == nil || o.FileTransferToRemote == nil {
		return nil, false
	}
	return o.FileTransferToRemote, true
}

// HasFileTransferToRemote returns a boolean if a field has been set.
func (o *WorkflowSshSessionAllOf) HasFileTransferToRemote() bool {
	if o != nil && o.FileTransferToRemote != nil {
		return true
	}

	return false
}

// SetFileTransferToRemote gets a reference to the given WorkflowFileTransfer and assigns it to the FileTransferToRemote field.
func (o *WorkflowSshSessionAllOf) SetFileTransferToRemote(v WorkflowFileTransfer) {
	o.FileTransferToRemote = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *WorkflowSshSessionAllOf) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *WorkflowSshSessionAllOf) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *WorkflowSshSessionAllOf) SetMessageType(v string) {
	o.MessageType = &v
}

// GetSshCommand returns the SshCommand field value if set, zero value otherwise.
func (o *WorkflowSshSessionAllOf) GetSshCommand() WorkflowSshCmd {
	if o == nil || o.SshCommand == nil {
		var ret WorkflowSshCmd
		return ret
	}
	return *o.SshCommand
}

// GetSshCommandOk returns a tuple with the SshCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetSshCommandOk() (*WorkflowSshCmd, bool) {
	if o == nil || o.SshCommand == nil {
		return nil, false
	}
	return o.SshCommand, true
}

// HasSshCommand returns a boolean if a field has been set.
func (o *WorkflowSshSessionAllOf) HasSshCommand() bool {
	if o != nil && o.SshCommand != nil {
		return true
	}

	return false
}

// SetSshCommand gets a reference to the given WorkflowSshCmd and assigns it to the SshCommand field.
func (o *WorkflowSshSessionAllOf) SetSshCommand(v WorkflowSshCmd) {
	o.SshCommand = &v
}

// GetSshConfiguration returns the SshConfiguration field value if set, zero value otherwise.
func (o *WorkflowSshSessionAllOf) GetSshConfiguration() WorkflowSshConfig {
	if o == nil || o.SshConfiguration == nil {
		var ret WorkflowSshConfig
		return ret
	}
	return *o.SshConfiguration
}

// GetSshConfigurationOk returns a tuple with the SshConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshSessionAllOf) GetSshConfigurationOk() (*WorkflowSshConfig, bool) {
	if o == nil || o.SshConfiguration == nil {
		return nil, false
	}
	return o.SshConfiguration, true
}

// HasSshConfiguration returns a boolean if a field has been set.
func (o *WorkflowSshSessionAllOf) HasSshConfiguration() bool {
	if o != nil && o.SshConfiguration != nil {
		return true
	}

	return false
}

// SetSshConfiguration gets a reference to the given WorkflowSshConfig and assigns it to the SshConfiguration field.
func (o *WorkflowSshSessionAllOf) SetSshConfiguration(v WorkflowSshConfig) {
	o.SshConfiguration = &v
}

func (o WorkflowSshSessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *WorkflowSshSessionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSshSessionAllOf := _WorkflowSshSessionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSshSessionAllOf); err == nil {
		*o = WorkflowSshSessionAllOf(varWorkflowSshSessionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileTransferToRemote")
		delete(additionalProperties, "MessageType")
		delete(additionalProperties, "SshCommand")
		delete(additionalProperties, "SshConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSshSessionAllOf struct {
	value *WorkflowSshSessionAllOf
	isSet bool
}

func (v NullableWorkflowSshSessionAllOf) Get() *WorkflowSshSessionAllOf {
	return v.value
}

func (v *NullableWorkflowSshSessionAllOf) Set(val *WorkflowSshSessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSshSessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSshSessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSshSessionAllOf(val *WorkflowSshSessionAllOf) *NullableWorkflowSshSessionAllOf {
	return &NullableWorkflowSshSessionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSshSessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSshSessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
