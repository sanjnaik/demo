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

// FabricPcOperation PcOperation objects allows the user to alter the state of the port channel.
type FabricPcOperation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured state to disable the port channel. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	AdminState *string `json:"AdminState,omitempty"`
	// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Port Channel Identifier for the collection of ports.
	PcId                 *int64                      `json:"PcId,omitempty"`
	NetworkElement       *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPcOperation FabricPcOperation

// NewFabricPcOperation instantiates a new FabricPcOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPcOperation(classId string, objectType string) *FabricPcOperation {
	this := FabricPcOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState string = "Enabled"
	this.AdminState = &adminState
	var configState string = "None"
	this.ConfigState = &configState
	return &this
}

// NewFabricPcOperationWithDefaults instantiates a new FabricPcOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPcOperationWithDefaults() *FabricPcOperation {
	this := FabricPcOperation{}
	var classId string = "fabric.PcOperation"
	this.ClassId = classId
	var objectType string = "fabric.PcOperation"
	this.ObjectType = objectType
	var adminState string = "Enabled"
	this.AdminState = &adminState
	var configState string = "None"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPcOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPcOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPcOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPcOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricPcOperation) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricPcOperation) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricPcOperation) SetAdminState(v string) {
	o.AdminState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *FabricPcOperation) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *FabricPcOperation) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *FabricPcOperation) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPcOperation) GetPcId() int64 {
	if o == nil || o.PcId == nil {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetPcIdOk() (*int64, bool) {
	if o == nil || o.PcId == nil {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPcOperation) HasPcId() bool {
	if o != nil && o.PcId != nil {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPcOperation) SetPcId(v int64) {
	o.PcId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricPcOperation) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPcOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricPcOperation) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricPcOperation) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o FabricPcOperation) MarshalJSON() ([]byte, error) {
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
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.PcId != nil {
		toSerialize["PcId"] = o.PcId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPcOperation) UnmarshalJSON(bytes []byte) (err error) {
	type FabricPcOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured state to disable the port channel. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
		AdminState *string `json:"AdminState,omitempty"`
		// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState *string `json:"ConfigState,omitempty"`
		// Port Channel Identifier for the collection of ports.
		PcId           *int64                      `json:"PcId,omitempty"`
		NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	}

	varFabricPcOperationWithoutEmbeddedStruct := FabricPcOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricPcOperationWithoutEmbeddedStruct)
	if err == nil {
		varFabricPcOperation := _FabricPcOperation{}
		varFabricPcOperation.ClassId = varFabricPcOperationWithoutEmbeddedStruct.ClassId
		varFabricPcOperation.ObjectType = varFabricPcOperationWithoutEmbeddedStruct.ObjectType
		varFabricPcOperation.AdminState = varFabricPcOperationWithoutEmbeddedStruct.AdminState
		varFabricPcOperation.ConfigState = varFabricPcOperationWithoutEmbeddedStruct.ConfigState
		varFabricPcOperation.PcId = varFabricPcOperationWithoutEmbeddedStruct.PcId
		varFabricPcOperation.NetworkElement = varFabricPcOperationWithoutEmbeddedStruct.NetworkElement
		*o = FabricPcOperation(varFabricPcOperation)
	} else {
		return err
	}

	varFabricPcOperation := _FabricPcOperation{}

	err = json.Unmarshal(bytes, &varFabricPcOperation)
	if err == nil {
		o.MoBaseMo = varFabricPcOperation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "PcId")
		delete(additionalProperties, "NetworkElement")

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

type NullableFabricPcOperation struct {
	value *FabricPcOperation
	isSet bool
}

func (v NullableFabricPcOperation) Get() *FabricPcOperation {
	return v.value
}

func (v *NullableFabricPcOperation) Set(val *FabricPcOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPcOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPcOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPcOperation(val *FabricPcOperation) *NullableFabricPcOperation {
	return &NullableFabricPcOperation{value: val, isSet: true}
}

func (v NullableFabricPcOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPcOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
