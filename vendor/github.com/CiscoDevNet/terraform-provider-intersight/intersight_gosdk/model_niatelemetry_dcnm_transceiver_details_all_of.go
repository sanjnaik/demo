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

// NiatelemetryDcnmTransceiverDetailsAllOf Definition of the list of properties defined in 'niatelemetry.DcnmTransceiverDetails', excluding properties defined in parent classes.
type NiatelemetryDcnmTransceiverDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Part number of the transceiver in the fabric inventory.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Product Id of the transceiver in the fabric inventory.
	ProductId *string `json:"ProductId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the transceiver in the fabric inventory.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Vendor Id of the transceiver in the fabric inventory.
	VendorId             *string                              `json:"VendorId,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDcnmTransceiverDetailsAllOf NiatelemetryDcnmTransceiverDetailsAllOf

// NewNiatelemetryDcnmTransceiverDetailsAllOf instantiates a new NiatelemetryDcnmTransceiverDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDcnmTransceiverDetailsAllOf(classId string, objectType string) *NiatelemetryDcnmTransceiverDetailsAllOf {
	this := NiatelemetryDcnmTransceiverDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDcnmTransceiverDetailsAllOfWithDefaults instantiates a new NiatelemetryDcnmTransceiverDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDcnmTransceiverDetailsAllOfWithDefaults() *NiatelemetryDcnmTransceiverDetailsAllOf {
	this := NiatelemetryDcnmTransceiverDetailsAllOf{}
	var classId string = "niatelemetry.DcnmTransceiverDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.DcnmTransceiverDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetProductId(v string) {
	o.ProductId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetVendorId() string {
	if o == nil || o.VendorId == nil {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetVendorIdOk() (*string, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetVendorId(v string) {
	o.VendorId = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryDcnmTransceiverDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryDcnmTransceiverDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.ProductId != nil {
		toSerialize["ProductId"] = o.ProductId
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryDcnmTransceiverDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryDcnmTransceiverDetailsAllOf := _NiatelemetryDcnmTransceiverDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryDcnmTransceiverDetailsAllOf); err == nil {
		*o = NiatelemetryDcnmTransceiverDetailsAllOf(varNiatelemetryDcnmTransceiverDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "ProductId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "VendorId")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryDcnmTransceiverDetailsAllOf struct {
	value *NiatelemetryDcnmTransceiverDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryDcnmTransceiverDetailsAllOf) Get() *NiatelemetryDcnmTransceiverDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsAllOf) Set(val *NiatelemetryDcnmTransceiverDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDcnmTransceiverDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDcnmTransceiverDetailsAllOf(val *NiatelemetryDcnmTransceiverDetailsAllOf) *NullableNiatelemetryDcnmTransceiverDetailsAllOf {
	return &NullableNiatelemetryDcnmTransceiverDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryDcnmTransceiverDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDcnmTransceiverDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
