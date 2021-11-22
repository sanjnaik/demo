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
	"time"
)

// CloudAwsNetworkInterface AWS Network Interface object is represented here.It is Virtual interface that can be attached to an instance in a Virtual Private Cloud (VPC).
type CloudAwsNetworkInterface struct {
	CloudBaseNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                                 `json:"ObjectType"`
	InstanceAttachment NullableCloudNetworkInstanceAttachment `json:"InstanceAttachment,omitempty"`
	// Time when this network interface is created.
	NicCreateTime *time.Time `json:"NicCreateTime,omitempty"`
	// The private DNS hostname name assigned to the network interface.
	PrivateDnsName   *string  `json:"PrivateDnsName,omitempty"`
	PrivateIpAddress []string `json:"PrivateIpAddress,omitempty"`
	// The public DNS hostname name assigned to the network interface.
	PublicDnsName   *string  `json:"PublicDnsName,omitempty"`
	PublicIpAddress []string `json:"PublicIpAddress,omitempty"`
	SecurityGroups  []string `json:"SecurityGroups,omitempty"`
	// The status of the network interface. If the network interface is not attached to an instance, the status is available; if a network interface is attached to an instance the status is in-use.
	Status               *string                     `json:"Status,omitempty"`
	AwsSubnet            *CloudAwsSubnetRelationship `json:"AwsSubnet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsNetworkInterface CloudAwsNetworkInterface

// NewCloudAwsNetworkInterface instantiates a new CloudAwsNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsNetworkInterface(classId string, objectType string) *CloudAwsNetworkInterface {
	this := CloudAwsNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterType string = "Unknown"
	this.AdapterType = &adapterType
	return &this
}

// NewCloudAwsNetworkInterfaceWithDefaults instantiates a new CloudAwsNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsNetworkInterfaceWithDefaults() *CloudAwsNetworkInterface {
	this := CloudAwsNetworkInterface{}
	var classId string = "cloud.AwsNetworkInterface"
	this.ClassId = classId
	var objectType string = "cloud.AwsNetworkInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInstanceAttachment returns the InstanceAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsNetworkInterface) GetInstanceAttachment() CloudNetworkInstanceAttachment {
	if o == nil || o.InstanceAttachment.Get() == nil {
		var ret CloudNetworkInstanceAttachment
		return ret
	}
	return *o.InstanceAttachment.Get()
}

// GetInstanceAttachmentOk returns a tuple with the InstanceAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsNetworkInterface) GetInstanceAttachmentOk() (*CloudNetworkInstanceAttachment, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceAttachment.Get(), o.InstanceAttachment.IsSet()
}

// HasInstanceAttachment returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasInstanceAttachment() bool {
	if o != nil && o.InstanceAttachment.IsSet() {
		return true
	}

	return false
}

// SetInstanceAttachment gets a reference to the given NullableCloudNetworkInstanceAttachment and assigns it to the InstanceAttachment field.
func (o *CloudAwsNetworkInterface) SetInstanceAttachment(v CloudNetworkInstanceAttachment) {
	o.InstanceAttachment.Set(&v)
}

// SetInstanceAttachmentNil sets the value for InstanceAttachment to be an explicit nil
func (o *CloudAwsNetworkInterface) SetInstanceAttachmentNil() {
	o.InstanceAttachment.Set(nil)
}

// UnsetInstanceAttachment ensures that no value is present for InstanceAttachment, not even an explicit nil
func (o *CloudAwsNetworkInterface) UnsetInstanceAttachment() {
	o.InstanceAttachment.Unset()
}

// GetNicCreateTime returns the NicCreateTime field value if set, zero value otherwise.
func (o *CloudAwsNetworkInterface) GetNicCreateTime() time.Time {
	if o == nil || o.NicCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NicCreateTime
}

// GetNicCreateTimeOk returns a tuple with the NicCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetNicCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.NicCreateTime == nil {
		return nil, false
	}
	return o.NicCreateTime, true
}

// HasNicCreateTime returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasNicCreateTime() bool {
	if o != nil && o.NicCreateTime != nil {
		return true
	}

	return false
}

// SetNicCreateTime gets a reference to the given time.Time and assigns it to the NicCreateTime field.
func (o *CloudAwsNetworkInterface) SetNicCreateTime(v time.Time) {
	o.NicCreateTime = &v
}

// GetPrivateDnsName returns the PrivateDnsName field value if set, zero value otherwise.
func (o *CloudAwsNetworkInterface) GetPrivateDnsName() string {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret
	}
	return *o.PrivateDnsName
}

// GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetPrivateDnsNameOk() (*string, bool) {
	if o == nil || o.PrivateDnsName == nil {
		return nil, false
	}
	return o.PrivateDnsName, true
}

// HasPrivateDnsName returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasPrivateDnsName() bool {
	if o != nil && o.PrivateDnsName != nil {
		return true
	}

	return false
}

// SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.
func (o *CloudAwsNetworkInterface) SetPrivateDnsName(v string) {
	o.PrivateDnsName = &v
}

// GetPrivateIpAddress returns the PrivateIpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsNetworkInterface) GetPrivateIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PrivateIpAddress
}

// GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsNetworkInterface) GetPrivateIpAddressOk() (*[]string, bool) {
	if o == nil || o.PrivateIpAddress == nil {
		return nil, false
	}
	return &o.PrivateIpAddress, true
}

// HasPrivateIpAddress returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasPrivateIpAddress() bool {
	if o != nil && o.PrivateIpAddress != nil {
		return true
	}

	return false
}

// SetPrivateIpAddress gets a reference to the given []string and assigns it to the PrivateIpAddress field.
func (o *CloudAwsNetworkInterface) SetPrivateIpAddress(v []string) {
	o.PrivateIpAddress = v
}

// GetPublicDnsName returns the PublicDnsName field value if set, zero value otherwise.
func (o *CloudAwsNetworkInterface) GetPublicDnsName() string {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret
	}
	return *o.PublicDnsName
}

// GetPublicDnsNameOk returns a tuple with the PublicDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetPublicDnsNameOk() (*string, bool) {
	if o == nil || o.PublicDnsName == nil {
		return nil, false
	}
	return o.PublicDnsName, true
}

// HasPublicDnsName returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasPublicDnsName() bool {
	if o != nil && o.PublicDnsName != nil {
		return true
	}

	return false
}

// SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.
func (o *CloudAwsNetworkInterface) SetPublicDnsName(v string) {
	o.PublicDnsName = &v
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsNetworkInterface) GetPublicIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsNetworkInterface) GetPublicIpAddressOk() (*[]string, bool) {
	if o == nil || o.PublicIpAddress == nil {
		return nil, false
	}
	return &o.PublicIpAddress, true
}

// HasPublicIpAddress returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasPublicIpAddress() bool {
	if o != nil && o.PublicIpAddress != nil {
		return true
	}

	return false
}

// SetPublicIpAddress gets a reference to the given []string and assigns it to the PublicIpAddress field.
func (o *CloudAwsNetworkInterface) SetPublicIpAddress(v []string) {
	o.PublicIpAddress = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsNetworkInterface) GetSecurityGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsNetworkInterface) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return &o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *CloudAwsNetworkInterface) SetSecurityGroups(v []string) {
	o.SecurityGroups = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloudAwsNetworkInterface) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudAwsNetworkInterface) SetStatus(v string) {
	o.Status = &v
}

// GetAwsSubnet returns the AwsSubnet field value if set, zero value otherwise.
func (o *CloudAwsNetworkInterface) GetAwsSubnet() CloudAwsSubnetRelationship {
	if o == nil || o.AwsSubnet == nil {
		var ret CloudAwsSubnetRelationship
		return ret
	}
	return *o.AwsSubnet
}

// GetAwsSubnetOk returns a tuple with the AwsSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsNetworkInterface) GetAwsSubnetOk() (*CloudAwsSubnetRelationship, bool) {
	if o == nil || o.AwsSubnet == nil {
		return nil, false
	}
	return o.AwsSubnet, true
}

// HasAwsSubnet returns a boolean if a field has been set.
func (o *CloudAwsNetworkInterface) HasAwsSubnet() bool {
	if o != nil && o.AwsSubnet != nil {
		return true
	}

	return false
}

// SetAwsSubnet gets a reference to the given CloudAwsSubnetRelationship and assigns it to the AwsSubnet field.
func (o *CloudAwsNetworkInterface) SetAwsSubnet(v CloudAwsSubnetRelationship) {
	o.AwsSubnet = &v
}

func (o CloudAwsNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseNetworkInterface, errCloudBaseNetworkInterface := json.Marshal(o.CloudBaseNetworkInterface)
	if errCloudBaseNetworkInterface != nil {
		return []byte{}, errCloudBaseNetworkInterface
	}
	errCloudBaseNetworkInterface = json.Unmarshal([]byte(serializedCloudBaseNetworkInterface), &toSerialize)
	if errCloudBaseNetworkInterface != nil {
		return []byte{}, errCloudBaseNetworkInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InstanceAttachment.IsSet() {
		toSerialize["InstanceAttachment"] = o.InstanceAttachment.Get()
	}
	if o.NicCreateTime != nil {
		toSerialize["NicCreateTime"] = o.NicCreateTime
	}
	if o.PrivateDnsName != nil {
		toSerialize["PrivateDnsName"] = o.PrivateDnsName
	}
	if o.PrivateIpAddress != nil {
		toSerialize["PrivateIpAddress"] = o.PrivateIpAddress
	}
	if o.PublicDnsName != nil {
		toSerialize["PublicDnsName"] = o.PublicDnsName
	}
	if o.PublicIpAddress != nil {
		toSerialize["PublicIpAddress"] = o.PublicIpAddress
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.AwsSubnet != nil {
		toSerialize["AwsSubnet"] = o.AwsSubnet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsNetworkInterface) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                 `json:"ObjectType"`
		InstanceAttachment NullableCloudNetworkInstanceAttachment `json:"InstanceAttachment,omitempty"`
		// Time when this network interface is created.
		NicCreateTime *time.Time `json:"NicCreateTime,omitempty"`
		// The private DNS hostname name assigned to the network interface.
		PrivateDnsName   *string  `json:"PrivateDnsName,omitempty"`
		PrivateIpAddress []string `json:"PrivateIpAddress,omitempty"`
		// The public DNS hostname name assigned to the network interface.
		PublicDnsName   *string  `json:"PublicDnsName,omitempty"`
		PublicIpAddress []string `json:"PublicIpAddress,omitempty"`
		SecurityGroups  []string `json:"SecurityGroups,omitempty"`
		// The status of the network interface. If the network interface is not attached to an instance, the status is available; if a network interface is attached to an instance the status is in-use.
		Status    *string                     `json:"Status,omitempty"`
		AwsSubnet *CloudAwsSubnetRelationship `json:"AwsSubnet,omitempty"`
	}

	varCloudAwsNetworkInterfaceWithoutEmbeddedStruct := CloudAwsNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsNetworkInterface := _CloudAwsNetworkInterface{}
		varCloudAwsNetworkInterface.ClassId = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varCloudAwsNetworkInterface.ObjectType = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varCloudAwsNetworkInterface.InstanceAttachment = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.InstanceAttachment
		varCloudAwsNetworkInterface.NicCreateTime = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.NicCreateTime
		varCloudAwsNetworkInterface.PrivateDnsName = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.PrivateDnsName
		varCloudAwsNetworkInterface.PrivateIpAddress = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.PrivateIpAddress
		varCloudAwsNetworkInterface.PublicDnsName = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.PublicDnsName
		varCloudAwsNetworkInterface.PublicIpAddress = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.PublicIpAddress
		varCloudAwsNetworkInterface.SecurityGroups = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.SecurityGroups
		varCloudAwsNetworkInterface.Status = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.Status
		varCloudAwsNetworkInterface.AwsSubnet = varCloudAwsNetworkInterfaceWithoutEmbeddedStruct.AwsSubnet
		*o = CloudAwsNetworkInterface(varCloudAwsNetworkInterface)
	} else {
		return err
	}

	varCloudAwsNetworkInterface := _CloudAwsNetworkInterface{}

	err = json.Unmarshal(bytes, &varCloudAwsNetworkInterface)
	if err == nil {
		o.CloudBaseNetworkInterface = varCloudAwsNetworkInterface.CloudBaseNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InstanceAttachment")
		delete(additionalProperties, "NicCreateTime")
		delete(additionalProperties, "PrivateDnsName")
		delete(additionalProperties, "PrivateIpAddress")
		delete(additionalProperties, "PublicDnsName")
		delete(additionalProperties, "PublicIpAddress")
		delete(additionalProperties, "SecurityGroups")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "AwsSubnet")

		// remove fields from embedded structs
		reflectCloudBaseNetworkInterface := reflect.ValueOf(o.CloudBaseNetworkInterface)
		for i := 0; i < reflectCloudBaseNetworkInterface.Type().NumField(); i++ {
			t := reflectCloudBaseNetworkInterface.Type().Field(i)

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

type NullableCloudAwsNetworkInterface struct {
	value *CloudAwsNetworkInterface
	isSet bool
}

func (v NullableCloudAwsNetworkInterface) Get() *CloudAwsNetworkInterface {
	return v.value
}

func (v *NullableCloudAwsNetworkInterface) Set(val *CloudAwsNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsNetworkInterface(val *CloudAwsNetworkInterface) *NullableCloudAwsNetworkInterface {
	return &NullableCloudAwsNetworkInterface{value: val, isSet: true}
}

func (v NullableCloudAwsNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
