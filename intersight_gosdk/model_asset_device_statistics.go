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

// AssetDeviceStatistics Type for saving device statistics related information for HyperFlex and UCS devices. It is used in asset.SubscriptionDeviceContractInformation object to save the current device statistics.
type AssetDeviceStatistics struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the cluster. It is specified only for HyperFlex based devices.
	ClusterName *string `json:"ClusterName,omitempty"`
	// The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected.
	Connected *int64 `json:"Connected,omitempty"`
	// Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices.
	MembershipRatio      *float32            `json:"MembershipRatio,omitempty"`
	VmHost               NullableAssetVmHost `json:"VmHost,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceStatistics AssetDeviceStatistics

// NewAssetDeviceStatistics instantiates a new AssetDeviceStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceStatistics(classId string, objectType string) *AssetDeviceStatistics {
	this := AssetDeviceStatistics{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceStatisticsWithDefaults instantiates a new AssetDeviceStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceStatisticsWithDefaults() *AssetDeviceStatistics {
	this := AssetDeviceStatistics{}
	var classId string = "asset.DeviceStatistics"
	this.ClassId = classId
	var objectType string = "asset.DeviceStatistics"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceStatistics) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceStatistics) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceStatistics) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceStatistics) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AssetDeviceStatistics) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetConnected() int64 {
	if o == nil || o.Connected == nil {
		var ret int64
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetConnectedOk() (*int64, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasConnected() bool {
	if o != nil && o.Connected != nil {
		return true
	}

	return false
}

// SetConnected gets a reference to the given int64 and assigns it to the Connected field.
func (o *AssetDeviceStatistics) SetConnected(v int64) {
	o.Connected = &v
}

// GetMembershipRatio returns the MembershipRatio field value if set, zero value otherwise.
func (o *AssetDeviceStatistics) GetMembershipRatio() float32 {
	if o == nil || o.MembershipRatio == nil {
		var ret float32
		return ret
	}
	return *o.MembershipRatio
}

// GetMembershipRatioOk returns a tuple with the MembershipRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceStatistics) GetMembershipRatioOk() (*float32, bool) {
	if o == nil || o.MembershipRatio == nil {
		return nil, false
	}
	return o.MembershipRatio, true
}

// HasMembershipRatio returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasMembershipRatio() bool {
	if o != nil && o.MembershipRatio != nil {
		return true
	}

	return false
}

// SetMembershipRatio gets a reference to the given float32 and assigns it to the MembershipRatio field.
func (o *AssetDeviceStatistics) SetMembershipRatio(v float32) {
	o.MembershipRatio = &v
}

// GetVmHost returns the VmHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceStatistics) GetVmHost() AssetVmHost {
	if o == nil || o.VmHost.Get() == nil {
		var ret AssetVmHost
		return ret
	}
	return *o.VmHost.Get()
}

// GetVmHostOk returns a tuple with the VmHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceStatistics) GetVmHostOk() (*AssetVmHost, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmHost.Get(), o.VmHost.IsSet()
}

// HasVmHost returns a boolean if a field has been set.
func (o *AssetDeviceStatistics) HasVmHost() bool {
	if o != nil && o.VmHost.IsSet() {
		return true
	}

	return false
}

// SetVmHost gets a reference to the given NullableAssetVmHost and assigns it to the VmHost field.
func (o *AssetDeviceStatistics) SetVmHost(v AssetVmHost) {
	o.VmHost.Set(&v)
}

// SetVmHostNil sets the value for VmHost to be an explicit nil
func (o *AssetDeviceStatistics) SetVmHostNil() {
	o.VmHost.Set(nil)
}

// UnsetVmHost ensures that no value is present for VmHost, not even an explicit nil
func (o *AssetDeviceStatistics) UnsetVmHost() {
	o.VmHost.Unset()
}

func (o AssetDeviceStatistics) MarshalJSON() ([]byte, error) {
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
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.Connected != nil {
		toSerialize["Connected"] = o.Connected
	}
	if o.MembershipRatio != nil {
		toSerialize["MembershipRatio"] = o.MembershipRatio
	}
	if o.VmHost.IsSet() {
		toSerialize["VmHost"] = o.VmHost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceStatistics) UnmarshalJSON(bytes []byte) (err error) {
	type AssetDeviceStatisticsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the cluster. It is specified only for HyperFlex based devices.
		ClusterName *string `json:"ClusterName,omitempty"`
		// The status of the persistent connection between the device connector and Intersight, for HyperFlex or UCS device. 1 represents being connected and 0 represents being disconnected.
		Connected *int64 `json:"Connected,omitempty"`
		// Defines the average proportion of resources used by the device within the cluster. example in a cluster having 3 nodes, the membershipRatio of each node is 1/3 or 0.33. It is specified only for HyperFlex based devices.
		MembershipRatio *float32            `json:"MembershipRatio,omitempty"`
		VmHost          NullableAssetVmHost `json:"VmHost,omitempty"`
	}

	varAssetDeviceStatisticsWithoutEmbeddedStruct := AssetDeviceStatisticsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetDeviceStatisticsWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceStatistics := _AssetDeviceStatistics{}
		varAssetDeviceStatistics.ClassId = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClassId
		varAssetDeviceStatistics.ObjectType = varAssetDeviceStatisticsWithoutEmbeddedStruct.ObjectType
		varAssetDeviceStatistics.ClusterName = varAssetDeviceStatisticsWithoutEmbeddedStruct.ClusterName
		varAssetDeviceStatistics.Connected = varAssetDeviceStatisticsWithoutEmbeddedStruct.Connected
		varAssetDeviceStatistics.MembershipRatio = varAssetDeviceStatisticsWithoutEmbeddedStruct.MembershipRatio
		varAssetDeviceStatistics.VmHost = varAssetDeviceStatisticsWithoutEmbeddedStruct.VmHost
		*o = AssetDeviceStatistics(varAssetDeviceStatistics)
	} else {
		return err
	}

	varAssetDeviceStatistics := _AssetDeviceStatistics{}

	err = json.Unmarshal(bytes, &varAssetDeviceStatistics)
	if err == nil {
		o.MoBaseComplexType = varAssetDeviceStatistics.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "Connected")
		delete(additionalProperties, "MembershipRatio")
		delete(additionalProperties, "VmHost")

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

type NullableAssetDeviceStatistics struct {
	value *AssetDeviceStatistics
	isSet bool
}

func (v NullableAssetDeviceStatistics) Get() *AssetDeviceStatistics {
	return v.value
}

func (v *NullableAssetDeviceStatistics) Set(val *AssetDeviceStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceStatistics(val *AssetDeviceStatistics) *NullableAssetDeviceStatistics {
	return &NullableAssetDeviceStatistics{value: val, isSet: true}
}

func (v NullableAssetDeviceStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
