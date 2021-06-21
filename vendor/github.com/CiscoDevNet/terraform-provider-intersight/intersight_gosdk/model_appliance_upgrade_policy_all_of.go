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
	"time"
)

// ApplianceUpgradePolicyAllOf Definition of the list of properties defined in 'appliance.UpgradePolicy', excluding properties defined in parent classes.
type ApplianceUpgradePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored.
	AutoUpgrade *bool `json:"AutoUpgrade,omitempty"`
	// If enabled, allows the user to define a blackout period during which the appliance will not be upgraded.
	BlackoutDatesEnabled *bool `json:"BlackoutDatesEnabled,omitempty"`
	// End date of the black out period.
	BlackoutEndDate *time.Time `json:"BlackoutEndDate,omitempty"`
	// Start date of the black out period. The appliance will not be upgraded during this period.
	BlackoutStartDate *time.Time `json:"BlackoutStartDate,omitempty"`
	// Indicates if the updated metadata files should be synced immediately or at the next upgrade.
	EnableMetaDataSync   *bool                   `json:"EnableMetaDataSync,omitempty"`
	Schedule             NullableOnpremSchedule  `json:"Schedule,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceUpgradePolicyAllOf ApplianceUpgradePolicyAllOf

// NewApplianceUpgradePolicyAllOf instantiates a new ApplianceUpgradePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceUpgradePolicyAllOf(classId string, objectType string) *ApplianceUpgradePolicyAllOf {
	this := ApplianceUpgradePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableMetaDataSync bool = true
	this.EnableMetaDataSync = &enableMetaDataSync
	return &this
}

// NewApplianceUpgradePolicyAllOfWithDefaults instantiates a new ApplianceUpgradePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceUpgradePolicyAllOfWithDefaults() *ApplianceUpgradePolicyAllOf {
	this := ApplianceUpgradePolicyAllOf{}
	var classId string = "appliance.UpgradePolicy"
	this.ClassId = classId
	var objectType string = "appliance.UpgradePolicy"
	this.ObjectType = objectType
	var enableMetaDataSync bool = true
	this.EnableMetaDataSync = &enableMetaDataSync
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceUpgradePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceUpgradePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceUpgradePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceUpgradePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoUpgrade returns the AutoUpgrade field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetAutoUpgrade() bool {
	if o == nil || o.AutoUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpgrade
}

// GetAutoUpgradeOk returns a tuple with the AutoUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetAutoUpgradeOk() (*bool, bool) {
	if o == nil || o.AutoUpgrade == nil {
		return nil, false
	}
	return o.AutoUpgrade, true
}

// HasAutoUpgrade returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasAutoUpgrade() bool {
	if o != nil && o.AutoUpgrade != nil {
		return true
	}

	return false
}

// SetAutoUpgrade gets a reference to the given bool and assigns it to the AutoUpgrade field.
func (o *ApplianceUpgradePolicyAllOf) SetAutoUpgrade(v bool) {
	o.AutoUpgrade = &v
}

// GetBlackoutDatesEnabled returns the BlackoutDatesEnabled field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutDatesEnabled() bool {
	if o == nil || o.BlackoutDatesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BlackoutDatesEnabled
}

// GetBlackoutDatesEnabledOk returns a tuple with the BlackoutDatesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutDatesEnabledOk() (*bool, bool) {
	if o == nil || o.BlackoutDatesEnabled == nil {
		return nil, false
	}
	return o.BlackoutDatesEnabled, true
}

// HasBlackoutDatesEnabled returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasBlackoutDatesEnabled() bool {
	if o != nil && o.BlackoutDatesEnabled != nil {
		return true
	}

	return false
}

// SetBlackoutDatesEnabled gets a reference to the given bool and assigns it to the BlackoutDatesEnabled field.
func (o *ApplianceUpgradePolicyAllOf) SetBlackoutDatesEnabled(v bool) {
	o.BlackoutDatesEnabled = &v
}

// GetBlackoutEndDate returns the BlackoutEndDate field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutEndDate() time.Time {
	if o == nil || o.BlackoutEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BlackoutEndDate
}

// GetBlackoutEndDateOk returns a tuple with the BlackoutEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutEndDateOk() (*time.Time, bool) {
	if o == nil || o.BlackoutEndDate == nil {
		return nil, false
	}
	return o.BlackoutEndDate, true
}

// HasBlackoutEndDate returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasBlackoutEndDate() bool {
	if o != nil && o.BlackoutEndDate != nil {
		return true
	}

	return false
}

// SetBlackoutEndDate gets a reference to the given time.Time and assigns it to the BlackoutEndDate field.
func (o *ApplianceUpgradePolicyAllOf) SetBlackoutEndDate(v time.Time) {
	o.BlackoutEndDate = &v
}

// GetBlackoutStartDate returns the BlackoutStartDate field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutStartDate() time.Time {
	if o == nil || o.BlackoutStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BlackoutStartDate
}

// GetBlackoutStartDateOk returns a tuple with the BlackoutStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetBlackoutStartDateOk() (*time.Time, bool) {
	if o == nil || o.BlackoutStartDate == nil {
		return nil, false
	}
	return o.BlackoutStartDate, true
}

// HasBlackoutStartDate returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasBlackoutStartDate() bool {
	if o != nil && o.BlackoutStartDate != nil {
		return true
	}

	return false
}

// SetBlackoutStartDate gets a reference to the given time.Time and assigns it to the BlackoutStartDate field.
func (o *ApplianceUpgradePolicyAllOf) SetBlackoutStartDate(v time.Time) {
	o.BlackoutStartDate = &v
}

// GetEnableMetaDataSync returns the EnableMetaDataSync field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetEnableMetaDataSync() bool {
	if o == nil || o.EnableMetaDataSync == nil {
		var ret bool
		return ret
	}
	return *o.EnableMetaDataSync
}

// GetEnableMetaDataSyncOk returns a tuple with the EnableMetaDataSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetEnableMetaDataSyncOk() (*bool, bool) {
	if o == nil || o.EnableMetaDataSync == nil {
		return nil, false
	}
	return o.EnableMetaDataSync, true
}

// HasEnableMetaDataSync returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasEnableMetaDataSync() bool {
	if o != nil && o.EnableMetaDataSync != nil {
		return true
	}

	return false
}

// SetEnableMetaDataSync gets a reference to the given bool and assigns it to the EnableMetaDataSync field.
func (o *ApplianceUpgradePolicyAllOf) SetEnableMetaDataSync(v bool) {
	o.EnableMetaDataSync = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgradePolicyAllOf) GetSchedule() OnpremSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret OnpremSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgradePolicyAllOf) GetScheduleOk() (*OnpremSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableOnpremSchedule and assigns it to the Schedule field.
func (o *ApplianceUpgradePolicyAllOf) SetSchedule(v OnpremSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *ApplianceUpgradePolicyAllOf) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *ApplianceUpgradePolicyAllOf) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceUpgradePolicyAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradePolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceUpgradePolicyAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceUpgradePolicyAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceUpgradePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoUpgrade != nil {
		toSerialize["AutoUpgrade"] = o.AutoUpgrade
	}
	if o.BlackoutDatesEnabled != nil {
		toSerialize["BlackoutDatesEnabled"] = o.BlackoutDatesEnabled
	}
	if o.BlackoutEndDate != nil {
		toSerialize["BlackoutEndDate"] = o.BlackoutEndDate
	}
	if o.BlackoutStartDate != nil {
		toSerialize["BlackoutStartDate"] = o.BlackoutStartDate
	}
	if o.EnableMetaDataSync != nil {
		toSerialize["EnableMetaDataSync"] = o.EnableMetaDataSync
	}
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceUpgradePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceUpgradePolicyAllOf := _ApplianceUpgradePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceUpgradePolicyAllOf); err == nil {
		*o = ApplianceUpgradePolicyAllOf(varApplianceUpgradePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoUpgrade")
		delete(additionalProperties, "BlackoutDatesEnabled")
		delete(additionalProperties, "BlackoutEndDate")
		delete(additionalProperties, "BlackoutStartDate")
		delete(additionalProperties, "EnableMetaDataSync")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceUpgradePolicyAllOf struct {
	value *ApplianceUpgradePolicyAllOf
	isSet bool
}

func (v NullableApplianceUpgradePolicyAllOf) Get() *ApplianceUpgradePolicyAllOf {
	return v.value
}

func (v *NullableApplianceUpgradePolicyAllOf) Set(val *ApplianceUpgradePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceUpgradePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceUpgradePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceUpgradePolicyAllOf(val *ApplianceUpgradePolicyAllOf) *NullableApplianceUpgradePolicyAllOf {
	return &NullableApplianceUpgradePolicyAllOf{value: val, isSet: true}
}

func (v NullableApplianceUpgradePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceUpgradePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
