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

// NiatelemetryMsoContractDetailsAllOf Definition of the list of properties defined in 'niatelemetry.MsoContractDetails', excluding properties defined in parent classes.
type NiatelemetryMsoContractDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of contract in Multi-Site Orchestrator.
	ContractName *string `json:"ContractName,omitempty"`
	// Site Ids to which this contract is deployed to.
	DeployedSites *string `json:"DeployedSites,omitempty"`
	// Is the contract local to the Multi-Site Orchestrator.
	IsLocal *string `json:"IsLocal,omitempty"`
	// Unique reference for the contract in the Multi-Site Orchestrator.
	Reference *string `json:"Reference,omitempty"`
	// Schema ID in Multi-Site Orchestrator.
	SchemaId *string `json:"SchemaId,omitempty"`
	// Schema name this contract belongs to in Multi-Site Orchestrator.
	SchemaName *string `json:"SchemaName,omitempty"`
	// Template name this contract belongs to in Multi-Site Orchestrator.
	TemplateName         *string                              `json:"TemplateName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoContractDetailsAllOf NiatelemetryMsoContractDetailsAllOf

// NewNiatelemetryMsoContractDetailsAllOf instantiates a new NiatelemetryMsoContractDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoContractDetailsAllOf(classId string, objectType string) *NiatelemetryMsoContractDetailsAllOf {
	this := NiatelemetryMsoContractDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoContractDetailsAllOfWithDefaults instantiates a new NiatelemetryMsoContractDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoContractDetailsAllOfWithDefaults() *NiatelemetryMsoContractDetailsAllOf {
	this := NiatelemetryMsoContractDetailsAllOf{}
	var classId string = "niatelemetry.MsoContractDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoContractDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoContractDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoContractDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoContractDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoContractDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetContractName(v string) {
	o.ContractName = &v
}

// GetDeployedSites returns the DeployedSites field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetDeployedSites() string {
	if o == nil || o.DeployedSites == nil {
		var ret string
		return ret
	}
	return *o.DeployedSites
}

// GetDeployedSitesOk returns a tuple with the DeployedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetDeployedSitesOk() (*string, bool) {
	if o == nil || o.DeployedSites == nil {
		return nil, false
	}
	return o.DeployedSites, true
}

// HasDeployedSites returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasDeployedSites() bool {
	if o != nil && o.DeployedSites != nil {
		return true
	}

	return false
}

// SetDeployedSites gets a reference to the given string and assigns it to the DeployedSites field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetDeployedSites(v string) {
	o.DeployedSites = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetIsLocal() string {
	if o == nil || o.IsLocal == nil {
		var ret string
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetIsLocalOk() (*string, bool) {
	if o == nil || o.IsLocal == nil {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasIsLocal() bool {
	if o != nil && o.IsLocal != nil {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given string and assigns it to the IsLocal field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetIsLocal(v string) {
	o.IsLocal = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetReference(v string) {
	o.Reference = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasSchemaName() bool {
	if o != nil && o.SchemaName != nil {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoContractDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoContractDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoContractDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoContractDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.DeployedSites != nil {
		toSerialize["DeployedSites"] = o.DeployedSites
	}
	if o.IsLocal != nil {
		toSerialize["IsLocal"] = o.IsLocal
	}
	if o.Reference != nil {
		toSerialize["Reference"] = o.Reference
	}
	if o.SchemaId != nil {
		toSerialize["SchemaId"] = o.SchemaId
	}
	if o.SchemaName != nil {
		toSerialize["SchemaName"] = o.SchemaName
	}
	if o.TemplateName != nil {
		toSerialize["TemplateName"] = o.TemplateName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMsoContractDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryMsoContractDetailsAllOf := _NiatelemetryMsoContractDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryMsoContractDetailsAllOf); err == nil {
		*o = NiatelemetryMsoContractDetailsAllOf(varNiatelemetryMsoContractDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "DeployedSites")
		delete(additionalProperties, "IsLocal")
		delete(additionalProperties, "Reference")
		delete(additionalProperties, "SchemaId")
		delete(additionalProperties, "SchemaName")
		delete(additionalProperties, "TemplateName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryMsoContractDetailsAllOf struct {
	value *NiatelemetryMsoContractDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryMsoContractDetailsAllOf) Get() *NiatelemetryMsoContractDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryMsoContractDetailsAllOf) Set(val *NiatelemetryMsoContractDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoContractDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoContractDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoContractDetailsAllOf(val *NiatelemetryMsoContractDetailsAllOf) *NullableNiatelemetryMsoContractDetailsAllOf {
	return &NullableNiatelemetryMsoContractDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoContractDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoContractDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
