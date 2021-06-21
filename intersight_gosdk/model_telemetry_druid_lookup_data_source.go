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
)

// TelemetryDruidLookupDataSource Lookup datasources correspond to Druid's key-value lookup objects. In Druid SQL, they reside in the the lookup schema. They are preloaded in memory on all servers, so they can be accessed rapidly. They can be joined onto regular tables using the join operator. Lookup datasources are key-value oriented and always have exactly two columns, k (the key) and v (the value), and both are always strings. Lookups can be joined with a base table either using an explicit join, or by using the SQL LOOKUP function. However, the join operator must evaluate the condition on each row, whereas the LOOKUP function can defer evaluation until after an aggregation phase. This means that the LOOKUP function is usually faster than joining to a lookup datasource.
type TelemetryDruidLookupDataSource struct {
	// The type of data source.
	Type string `json:"type"`
	// the name of the lookup object.
	Lookup               string `json:"lookup"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidLookupDataSource TelemetryDruidLookupDataSource

// NewTelemetryDruidLookupDataSource instantiates a new TelemetryDruidLookupDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidLookupDataSource(type_ string, lookup string) *TelemetryDruidLookupDataSource {
	this := TelemetryDruidLookupDataSource{}
	this.Type = type_
	this.Lookup = lookup
	return &this
}

// NewTelemetryDruidLookupDataSourceWithDefaults instantiates a new TelemetryDruidLookupDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidLookupDataSourceWithDefaults() *TelemetryDruidLookupDataSource {
	this := TelemetryDruidLookupDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidLookupDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidLookupDataSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidLookupDataSource) SetType(v string) {
	o.Type = v
}

// GetLookup returns the Lookup field value
func (o *TelemetryDruidLookupDataSource) GetLookup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidLookupDataSource) GetLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lookup, true
}

// SetLookup sets field value
func (o *TelemetryDruidLookupDataSource) SetLookup(v string) {
	o.Lookup = v
}

func (o TelemetryDruidLookupDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["lookup"] = o.Lookup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidLookupDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidLookupDataSource := _TelemetryDruidLookupDataSource{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidLookupDataSource); err == nil {
		*o = TelemetryDruidLookupDataSource(varTelemetryDruidLookupDataSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "lookup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidLookupDataSource struct {
	value *TelemetryDruidLookupDataSource
	isSet bool
}

func (v NullableTelemetryDruidLookupDataSource) Get() *TelemetryDruidLookupDataSource {
	return v.value
}

func (v *NullableTelemetryDruidLookupDataSource) Set(val *TelemetryDruidLookupDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidLookupDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidLookupDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidLookupDataSource(val *TelemetryDruidLookupDataSource) *NullableTelemetryDruidLookupDataSource {
	return &NullableTelemetryDruidLookupDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidLookupDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidLookupDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
