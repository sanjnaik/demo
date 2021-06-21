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

// ComputeServerSetting Models the configurable properties of a server in Intersight.
type ComputeServerSetting struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured state of the locator LED for the server. * `None` - No operation property for locator led. * `On` - The Locator Led is turned on. * `Off` - The Locator Led is turned off.
	AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`
	// User configured power state of the server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
	AdminPowerState *string `json:"AdminPowerState,omitempty"`
	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// The property used to identify the name of the server it is associated with.
	Name *string `json:"Name,omitempty"`
	// The name of the device chosen by user for configuring One-Time Boot device.
	OneTimeBootDevice             *string                                      `json:"OneTimeBootDevice,omitempty"`
	PersistentMemoryOperation     NullableComputePersistentMemoryOperation     `json:"PersistentMemoryOperation,omitempty"`
	ServerConfig                  NullableComputeServerConfig                  `json:"ServerConfig,omitempty"`
	StorageControllerOperation    NullableComputeStorageControllerOperation    `json:"StorageControllerOperation,omitempty"`
	StoragePhysicalDriveOperation NullableComputeStoragePhysicalDriveOperation `json:"StoragePhysicalDriveOperation,omitempty"`
	StorageVirtualDriveOperation  NullableComputeStorageVirtualDriveOperation  `json:"StorageVirtualDriveOperation,omitempty"`
	LocatorLed                    *EquipmentLocatorLedRelationship             `json:"LocatorLed,omitempty"`
	RegisteredDevice              *AssetDeviceRegistrationRelationship         `json:"RegisteredDevice,omitempty"`
	RunningWorkflow               *WorkflowWorkflowInfoRelationship            `json:"RunningWorkflow,omitempty"`
	Server                        *ComputePhysicalRelationship                 `json:"Server,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _ComputeServerSetting ComputeServerSetting

// NewComputeServerSetting instantiates a new ComputeServerSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerSetting(classId string, objectType string) *ComputeServerSetting {
	this := ComputeServerSetting{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// NewComputeServerSettingWithDefaults instantiates a new ComputeServerSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerSettingWithDefaults() *ComputeServerSetting {
	this := ComputeServerSetting{}
	var classId string = "compute.ServerSetting"
	this.ClassId = classId
	var objectType string = "compute.ServerSetting"
	this.ObjectType = objectType
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerSetting) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerSetting) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerSetting) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerSetting) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminLocatorLedState returns the AdminLocatorLedState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetAdminLocatorLedState() string {
	if o == nil || o.AdminLocatorLedState == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedState
}

// GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetAdminLocatorLedStateOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedState == nil {
		return nil, false
	}
	return o.AdminLocatorLedState, true
}

// HasAdminLocatorLedState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasAdminLocatorLedState() bool {
	if o != nil && o.AdminLocatorLedState != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedState gets a reference to the given string and assigns it to the AdminLocatorLedState field.
func (o *ComputeServerSetting) SetAdminLocatorLedState(v string) {
	o.AdminLocatorLedState = &v
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *ComputeServerSetting) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *ComputeServerSetting) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeServerSetting) SetName(v string) {
	o.Name = &v
}

// GetOneTimeBootDevice returns the OneTimeBootDevice field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetOneTimeBootDevice() string {
	if o == nil || o.OneTimeBootDevice == nil {
		var ret string
		return ret
	}
	return *o.OneTimeBootDevice
}

// GetOneTimeBootDeviceOk returns a tuple with the OneTimeBootDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetOneTimeBootDeviceOk() (*string, bool) {
	if o == nil || o.OneTimeBootDevice == nil {
		return nil, false
	}
	return o.OneTimeBootDevice, true
}

// HasOneTimeBootDevice returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasOneTimeBootDevice() bool {
	if o != nil && o.OneTimeBootDevice != nil {
		return true
	}

	return false
}

// SetOneTimeBootDevice gets a reference to the given string and assigns it to the OneTimeBootDevice field.
func (o *ComputeServerSetting) SetOneTimeBootDevice(v string) {
	o.OneTimeBootDevice = &v
}

// GetPersistentMemoryOperation returns the PersistentMemoryOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSetting) GetPersistentMemoryOperation() ComputePersistentMemoryOperation {
	if o == nil || o.PersistentMemoryOperation.Get() == nil {
		var ret ComputePersistentMemoryOperation
		return ret
	}
	return *o.PersistentMemoryOperation.Get()
}

// GetPersistentMemoryOperationOk returns a tuple with the PersistentMemoryOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSetting) GetPersistentMemoryOperationOk() (*ComputePersistentMemoryOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersistentMemoryOperation.Get(), o.PersistentMemoryOperation.IsSet()
}

// HasPersistentMemoryOperation returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasPersistentMemoryOperation() bool {
	if o != nil && o.PersistentMemoryOperation.IsSet() {
		return true
	}

	return false
}

// SetPersistentMemoryOperation gets a reference to the given NullableComputePersistentMemoryOperation and assigns it to the PersistentMemoryOperation field.
func (o *ComputeServerSetting) SetPersistentMemoryOperation(v ComputePersistentMemoryOperation) {
	o.PersistentMemoryOperation.Set(&v)
}

// SetPersistentMemoryOperationNil sets the value for PersistentMemoryOperation to be an explicit nil
func (o *ComputeServerSetting) SetPersistentMemoryOperationNil() {
	o.PersistentMemoryOperation.Set(nil)
}

// UnsetPersistentMemoryOperation ensures that no value is present for PersistentMemoryOperation, not even an explicit nil
func (o *ComputeServerSetting) UnsetPersistentMemoryOperation() {
	o.PersistentMemoryOperation.Unset()
}

// GetServerConfig returns the ServerConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSetting) GetServerConfig() ComputeServerConfig {
	if o == nil || o.ServerConfig.Get() == nil {
		var ret ComputeServerConfig
		return ret
	}
	return *o.ServerConfig.Get()
}

// GetServerConfigOk returns a tuple with the ServerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSetting) GetServerConfigOk() (*ComputeServerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerConfig.Get(), o.ServerConfig.IsSet()
}

// HasServerConfig returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasServerConfig() bool {
	if o != nil && o.ServerConfig.IsSet() {
		return true
	}

	return false
}

// SetServerConfig gets a reference to the given NullableComputeServerConfig and assigns it to the ServerConfig field.
func (o *ComputeServerSetting) SetServerConfig(v ComputeServerConfig) {
	o.ServerConfig.Set(&v)
}

// SetServerConfigNil sets the value for ServerConfig to be an explicit nil
func (o *ComputeServerSetting) SetServerConfigNil() {
	o.ServerConfig.Set(nil)
}

// UnsetServerConfig ensures that no value is present for ServerConfig, not even an explicit nil
func (o *ComputeServerSetting) UnsetServerConfig() {
	o.ServerConfig.Unset()
}

// GetStorageControllerOperation returns the StorageControllerOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSetting) GetStorageControllerOperation() ComputeStorageControllerOperation {
	if o == nil || o.StorageControllerOperation.Get() == nil {
		var ret ComputeStorageControllerOperation
		return ret
	}
	return *o.StorageControllerOperation.Get()
}

// GetStorageControllerOperationOk returns a tuple with the StorageControllerOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSetting) GetStorageControllerOperationOk() (*ComputeStorageControllerOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageControllerOperation.Get(), o.StorageControllerOperation.IsSet()
}

// HasStorageControllerOperation returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasStorageControllerOperation() bool {
	if o != nil && o.StorageControllerOperation.IsSet() {
		return true
	}

	return false
}

// SetStorageControllerOperation gets a reference to the given NullableComputeStorageControllerOperation and assigns it to the StorageControllerOperation field.
func (o *ComputeServerSetting) SetStorageControllerOperation(v ComputeStorageControllerOperation) {
	o.StorageControllerOperation.Set(&v)
}

// SetStorageControllerOperationNil sets the value for StorageControllerOperation to be an explicit nil
func (o *ComputeServerSetting) SetStorageControllerOperationNil() {
	o.StorageControllerOperation.Set(nil)
}

// UnsetStorageControllerOperation ensures that no value is present for StorageControllerOperation, not even an explicit nil
func (o *ComputeServerSetting) UnsetStorageControllerOperation() {
	o.StorageControllerOperation.Unset()
}

// GetStoragePhysicalDriveOperation returns the StoragePhysicalDriveOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSetting) GetStoragePhysicalDriveOperation() ComputeStoragePhysicalDriveOperation {
	if o == nil || o.StoragePhysicalDriveOperation.Get() == nil {
		var ret ComputeStoragePhysicalDriveOperation
		return ret
	}
	return *o.StoragePhysicalDriveOperation.Get()
}

// GetStoragePhysicalDriveOperationOk returns a tuple with the StoragePhysicalDriveOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSetting) GetStoragePhysicalDriveOperationOk() (*ComputeStoragePhysicalDriveOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoragePhysicalDriveOperation.Get(), o.StoragePhysicalDriveOperation.IsSet()
}

// HasStoragePhysicalDriveOperation returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasStoragePhysicalDriveOperation() bool {
	if o != nil && o.StoragePhysicalDriveOperation.IsSet() {
		return true
	}

	return false
}

// SetStoragePhysicalDriveOperation gets a reference to the given NullableComputeStoragePhysicalDriveOperation and assigns it to the StoragePhysicalDriveOperation field.
func (o *ComputeServerSetting) SetStoragePhysicalDriveOperation(v ComputeStoragePhysicalDriveOperation) {
	o.StoragePhysicalDriveOperation.Set(&v)
}

// SetStoragePhysicalDriveOperationNil sets the value for StoragePhysicalDriveOperation to be an explicit nil
func (o *ComputeServerSetting) SetStoragePhysicalDriveOperationNil() {
	o.StoragePhysicalDriveOperation.Set(nil)
}

// UnsetStoragePhysicalDriveOperation ensures that no value is present for StoragePhysicalDriveOperation, not even an explicit nil
func (o *ComputeServerSetting) UnsetStoragePhysicalDriveOperation() {
	o.StoragePhysicalDriveOperation.Unset()
}

// GetStorageVirtualDriveOperation returns the StorageVirtualDriveOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSetting) GetStorageVirtualDriveOperation() ComputeStorageVirtualDriveOperation {
	if o == nil || o.StorageVirtualDriveOperation.Get() == nil {
		var ret ComputeStorageVirtualDriveOperation
		return ret
	}
	return *o.StorageVirtualDriveOperation.Get()
}

// GetStorageVirtualDriveOperationOk returns a tuple with the StorageVirtualDriveOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSetting) GetStorageVirtualDriveOperationOk() (*ComputeStorageVirtualDriveOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageVirtualDriveOperation.Get(), o.StorageVirtualDriveOperation.IsSet()
}

// HasStorageVirtualDriveOperation returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasStorageVirtualDriveOperation() bool {
	if o != nil && o.StorageVirtualDriveOperation.IsSet() {
		return true
	}

	return false
}

// SetStorageVirtualDriveOperation gets a reference to the given NullableComputeStorageVirtualDriveOperation and assigns it to the StorageVirtualDriveOperation field.
func (o *ComputeServerSetting) SetStorageVirtualDriveOperation(v ComputeStorageVirtualDriveOperation) {
	o.StorageVirtualDriveOperation.Set(&v)
}

// SetStorageVirtualDriveOperationNil sets the value for StorageVirtualDriveOperation to be an explicit nil
func (o *ComputeServerSetting) SetStorageVirtualDriveOperationNil() {
	o.StorageVirtualDriveOperation.Set(nil)
}

// UnsetStorageVirtualDriveOperation ensures that no value is present for StorageVirtualDriveOperation, not even an explicit nil
func (o *ComputeServerSetting) UnsetStorageVirtualDriveOperation() {
	o.StorageVirtualDriveOperation.Unset()
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeServerSetting) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeServerSetting) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningWorkflow returns the RunningWorkflow field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetRunningWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.RunningWorkflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.RunningWorkflow
}

// GetRunningWorkflowOk returns a tuple with the RunningWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetRunningWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflow == nil {
		return nil, false
	}
	return o.RunningWorkflow, true
}

// HasRunningWorkflow returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasRunningWorkflow() bool {
	if o != nil && o.RunningWorkflow != nil {
		return true
	}

	return false
}

// SetRunningWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflow field.
func (o *ComputeServerSetting) SetRunningWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflow = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *ComputeServerSetting) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o ComputeServerSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminLocatorLedState != nil {
		toSerialize["AdminLocatorLedState"] = o.AdminLocatorLedState
	}
	if o.AdminPowerState != nil {
		toSerialize["AdminPowerState"] = o.AdminPowerState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OneTimeBootDevice != nil {
		toSerialize["OneTimeBootDevice"] = o.OneTimeBootDevice
	}
	if o.PersistentMemoryOperation.IsSet() {
		toSerialize["PersistentMemoryOperation"] = o.PersistentMemoryOperation.Get()
	}
	if o.ServerConfig.IsSet() {
		toSerialize["ServerConfig"] = o.ServerConfig.Get()
	}
	if o.StorageControllerOperation.IsSet() {
		toSerialize["StorageControllerOperation"] = o.StorageControllerOperation.Get()
	}
	if o.StoragePhysicalDriveOperation.IsSet() {
		toSerialize["StoragePhysicalDriveOperation"] = o.StoragePhysicalDriveOperation.Get()
	}
	if o.StorageVirtualDriveOperation.IsSet() {
		toSerialize["StorageVirtualDriveOperation"] = o.StorageVirtualDriveOperation.Get()
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningWorkflow != nil {
		toSerialize["RunningWorkflow"] = o.RunningWorkflow
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerSetting) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeServerSettingWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// User configured state of the locator LED for the server. * `None` - No operation property for locator led. * `On` - The Locator Led is turned on. * `Off` - The Locator Led is turned off.
		AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`
		// User configured power state of the server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
		AdminPowerState *string `json:"AdminPowerState,omitempty"`
		// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState *string `json:"ConfigState,omitempty"`
		// The property used to identify the name of the server it is associated with.
		Name *string `json:"Name,omitempty"`
		// The name of the device chosen by user for configuring One-Time Boot device.
		OneTimeBootDevice             *string                                      `json:"OneTimeBootDevice,omitempty"`
		PersistentMemoryOperation     NullableComputePersistentMemoryOperation     `json:"PersistentMemoryOperation,omitempty"`
		ServerConfig                  NullableComputeServerConfig                  `json:"ServerConfig,omitempty"`
		StorageControllerOperation    NullableComputeStorageControllerOperation    `json:"StorageControllerOperation,omitempty"`
		StoragePhysicalDriveOperation NullableComputeStoragePhysicalDriveOperation `json:"StoragePhysicalDriveOperation,omitempty"`
		StorageVirtualDriveOperation  NullableComputeStorageVirtualDriveOperation  `json:"StorageVirtualDriveOperation,omitempty"`
		LocatorLed                    *EquipmentLocatorLedRelationship             `json:"LocatorLed,omitempty"`
		RegisteredDevice              *AssetDeviceRegistrationRelationship         `json:"RegisteredDevice,omitempty"`
		RunningWorkflow               *WorkflowWorkflowInfoRelationship            `json:"RunningWorkflow,omitempty"`
		Server                        *ComputePhysicalRelationship                 `json:"Server,omitempty"`
	}

	varComputeServerSettingWithoutEmbeddedStruct := ComputeServerSettingWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeServerSettingWithoutEmbeddedStruct)
	if err == nil {
		varComputeServerSetting := _ComputeServerSetting{}
		varComputeServerSetting.ClassId = varComputeServerSettingWithoutEmbeddedStruct.ClassId
		varComputeServerSetting.ObjectType = varComputeServerSettingWithoutEmbeddedStruct.ObjectType
		varComputeServerSetting.AdminLocatorLedState = varComputeServerSettingWithoutEmbeddedStruct.AdminLocatorLedState
		varComputeServerSetting.AdminPowerState = varComputeServerSettingWithoutEmbeddedStruct.AdminPowerState
		varComputeServerSetting.ConfigState = varComputeServerSettingWithoutEmbeddedStruct.ConfigState
		varComputeServerSetting.Name = varComputeServerSettingWithoutEmbeddedStruct.Name
		varComputeServerSetting.OneTimeBootDevice = varComputeServerSettingWithoutEmbeddedStruct.OneTimeBootDevice
		varComputeServerSetting.PersistentMemoryOperation = varComputeServerSettingWithoutEmbeddedStruct.PersistentMemoryOperation
		varComputeServerSetting.ServerConfig = varComputeServerSettingWithoutEmbeddedStruct.ServerConfig
		varComputeServerSetting.StorageControllerOperation = varComputeServerSettingWithoutEmbeddedStruct.StorageControllerOperation
		varComputeServerSetting.StoragePhysicalDriveOperation = varComputeServerSettingWithoutEmbeddedStruct.StoragePhysicalDriveOperation
		varComputeServerSetting.StorageVirtualDriveOperation = varComputeServerSettingWithoutEmbeddedStruct.StorageVirtualDriveOperation
		varComputeServerSetting.LocatorLed = varComputeServerSettingWithoutEmbeddedStruct.LocatorLed
		varComputeServerSetting.RegisteredDevice = varComputeServerSettingWithoutEmbeddedStruct.RegisteredDevice
		varComputeServerSetting.RunningWorkflow = varComputeServerSettingWithoutEmbeddedStruct.RunningWorkflow
		varComputeServerSetting.Server = varComputeServerSettingWithoutEmbeddedStruct.Server
		*o = ComputeServerSetting(varComputeServerSetting)
	} else {
		return err
	}

	varComputeServerSetting := _ComputeServerSetting{}

	err = json.Unmarshal(bytes, &varComputeServerSetting)
	if err == nil {
		o.InventoryBase = varComputeServerSetting.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminLocatorLedState")
		delete(additionalProperties, "AdminPowerState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OneTimeBootDevice")
		delete(additionalProperties, "PersistentMemoryOperation")
		delete(additionalProperties, "ServerConfig")
		delete(additionalProperties, "StorageControllerOperation")
		delete(additionalProperties, "StoragePhysicalDriveOperation")
		delete(additionalProperties, "StorageVirtualDriveOperation")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningWorkflow")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableComputeServerSetting struct {
	value *ComputeServerSetting
	isSet bool
}

func (v NullableComputeServerSetting) Get() *ComputeServerSetting {
	return v.value
}

func (v *NullableComputeServerSetting) Set(val *ComputeServerSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerSetting(val *ComputeServerSetting) *NullableComputeServerSetting {
	return &NullableComputeServerSetting{value: val, isSet: true}
}

func (v NullableComputeServerSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
