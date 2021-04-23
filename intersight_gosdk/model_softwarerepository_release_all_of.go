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
	"time"
)

// SoftwarerepositoryReleaseAllOf Definition of the list of properties defined in 'softwarerepository.Release', excluding properties defined in parent classes.
type SoftwarerepositoryReleaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The date when the file was released or distributed by its vendor.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
	// The URL for the release notes of this image.
	ReleaseNotesUrl *string  `json:"ReleaseNotesUrl,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The platform type for which the images are released. This can be a Fabric Interconnect or compute server hardware. * `FabricSwitch` - The images in a release that correspond to Fabric Interconnect switches. * `ComputeSystem` - The images in a release that correspond to servers.
	Type *string `json:"Type,omitempty"`
	// Cisco provided release version.
	Version              *string                                `json:"Version,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryReleaseAllOf SoftwarerepositoryReleaseAllOf

// NewSoftwarerepositoryReleaseAllOf instantiates a new SoftwarerepositoryReleaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryReleaseAllOf(classId string, objectType string) *SoftwarerepositoryReleaseAllOf {
	this := SoftwarerepositoryReleaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// NewSoftwarerepositoryReleaseAllOfWithDefaults instantiates a new SoftwarerepositoryReleaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryReleaseAllOfWithDefaults() *SoftwarerepositoryReleaseAllOf {
	this := SoftwarerepositoryReleaseAllOf{}
	var classId string = "softwarerepository.Release"
	this.ClassId = classId
	var objectType string = "softwarerepository.Release"
	this.ObjectType = objectType
	var type_ string = "FabricSwitch"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryReleaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryReleaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryReleaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryReleaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SoftwarerepositoryReleaseAllOf) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *SoftwarerepositoryReleaseAllOf) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetReleaseNotesUrl returns the ReleaseNotesUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryReleaseAllOf) GetReleaseNotesUrl() string {
	if o == nil || o.ReleaseNotesUrl == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNotesUrl
}

// GetReleaseNotesUrlOk returns a tuple with the ReleaseNotesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetReleaseNotesUrlOk() (*string, bool) {
	if o == nil || o.ReleaseNotesUrl == nil {
		return nil, false
	}
	return o.ReleaseNotesUrl, true
}

// HasReleaseNotesUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasReleaseNotesUrl() bool {
	if o != nil && o.ReleaseNotesUrl != nil {
		return true
	}

	return false
}

// SetReleaseNotesUrl gets a reference to the given string and assigns it to the ReleaseNotesUrl field.
func (o *SoftwarerepositoryReleaseAllOf) SetReleaseNotesUrl(v string) {
	o.ReleaseNotesUrl = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryReleaseAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryReleaseAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryReleaseAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SoftwarerepositoryReleaseAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SoftwarerepositoryReleaseAllOf) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryReleaseAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryReleaseAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwarerepositoryReleaseAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryReleaseAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwarerepositoryReleaseAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwarerepositoryReleaseAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwarerepositoryReleaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ReleaseDate != nil {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if o.ReleaseNotesUrl != nil {
		toSerialize["ReleaseNotesUrl"] = o.ReleaseNotesUrl
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryReleaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryReleaseAllOf := _SoftwarerepositoryReleaseAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryReleaseAllOf); err == nil {
		*o = SoftwarerepositoryReleaseAllOf(varSoftwarerepositoryReleaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ReleaseDate")
		delete(additionalProperties, "ReleaseNotesUrl")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryReleaseAllOf struct {
	value *SoftwarerepositoryReleaseAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryReleaseAllOf) Get() *SoftwarerepositoryReleaseAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryReleaseAllOf) Set(val *SoftwarerepositoryReleaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryReleaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryReleaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryReleaseAllOf(val *SoftwarerepositoryReleaseAllOf) *NullableSoftwarerepositoryReleaseAllOf {
	return &NullableSoftwarerepositoryReleaseAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryReleaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryReleaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
