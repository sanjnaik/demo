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

// BulkMoCloner The MO Cloner interface facilitates making n number of shallow copies of any resource instance which supports the CREATE operation. The copy of the child objects of the resource is not supported currently, so the child references will not be copied in the clone. The MO to be cloned should be specified as an MoRef object in the \"Sources\". The \"Targets\" array should contain n JSON documents each compliant to RFC 7386.  For each target mo to be created, the user can specify the following - - new values for the identity properties, if applicable - new values for specific properties or references of the source MO which need to be overridden in the cloned object.
type BulkMoCloner struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	Responses            []BulkRestResult                      `json:"Responses,omitempty"`
	Sources              []MoBaseMo                            `json:"Sources,omitempty"`
	Targets              []MoBaseMo                            `json:"Targets,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkMoCloner BulkMoCloner

// NewBulkMoCloner instantiates a new BulkMoCloner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMoCloner(classId string, objectType string) *BulkMoCloner {
	this := BulkMoCloner{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkMoClonerWithDefaults instantiates a new BulkMoCloner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMoClonerWithDefaults() *BulkMoCloner {
	this := BulkMoCloner{}
	var classId string = "bulk.MoCloner"
	this.ClassId = classId
	var objectType string = "bulk.MoCloner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkMoCloner) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkMoCloner) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkMoCloner) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkMoCloner) SetObjectType(v string) {
	o.ObjectType = v
}

// GetResponses returns the Responses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetResponses() []BulkRestResult {
	if o == nil {
		var ret []BulkRestResult
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetResponsesOk() (*[]BulkRestResult, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return &o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *BulkMoCloner) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []BulkRestResult and assigns it to the Responses field.
func (o *BulkMoCloner) SetResponses(v []BulkRestResult) {
	o.Responses = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetSources() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetSourcesOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *BulkMoCloner) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []MoBaseMo and assigns it to the Sources field.
func (o *BulkMoCloner) SetSources(v []MoBaseMo) {
	o.Sources = v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkMoCloner) GetTargets() []MoBaseMo {
	if o == nil {
		var ret []MoBaseMo
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkMoCloner) GetTargetsOk() (*[]MoBaseMo, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *BulkMoCloner) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MoBaseMo and assigns it to the Targets field.
func (o *BulkMoCloner) SetTargets(v []MoBaseMo) {
	o.Targets = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkMoCloner) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkMoCloner) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkMoCloner) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkMoCloner) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o BulkMoCloner) MarshalJSON() ([]byte, error) {
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
	if o.Responses != nil {
		toSerialize["Responses"] = o.Responses
	}
	if o.Sources != nil {
		toSerialize["Sources"] = o.Sources
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkMoCloner) UnmarshalJSON(bytes []byte) (err error) {
	type BulkMoClonerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                `json:"ObjectType"`
		Responses    []BulkRestResult                      `json:"Responses,omitempty"`
		Sources      []MoBaseMo                            `json:"Sources,omitempty"`
		Targets      []MoBaseMo                            `json:"Targets,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varBulkMoClonerWithoutEmbeddedStruct := BulkMoClonerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkMoClonerWithoutEmbeddedStruct)
	if err == nil {
		varBulkMoCloner := _BulkMoCloner{}
		varBulkMoCloner.ClassId = varBulkMoClonerWithoutEmbeddedStruct.ClassId
		varBulkMoCloner.ObjectType = varBulkMoClonerWithoutEmbeddedStruct.ObjectType
		varBulkMoCloner.Responses = varBulkMoClonerWithoutEmbeddedStruct.Responses
		varBulkMoCloner.Sources = varBulkMoClonerWithoutEmbeddedStruct.Sources
		varBulkMoCloner.Targets = varBulkMoClonerWithoutEmbeddedStruct.Targets
		varBulkMoCloner.Organization = varBulkMoClonerWithoutEmbeddedStruct.Organization
		*o = BulkMoCloner(varBulkMoCloner)
	} else {
		return err
	}

	varBulkMoCloner := _BulkMoCloner{}

	err = json.Unmarshal(bytes, &varBulkMoCloner)
	if err == nil {
		o.MoBaseMo = varBulkMoCloner.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Responses")
		delete(additionalProperties, "Sources")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "Organization")

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

type NullableBulkMoCloner struct {
	value *BulkMoCloner
	isSet bool
}

func (v NullableBulkMoCloner) Get() *BulkMoCloner {
	return v.value
}

func (v *NullableBulkMoCloner) Set(val *BulkMoCloner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMoCloner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMoCloner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMoCloner(val *BulkMoCloner) *NullableBulkMoCloner {
	return &NullableBulkMoCloner{value: val, isSet: true}
}

func (v NullableBulkMoCloner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMoCloner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
