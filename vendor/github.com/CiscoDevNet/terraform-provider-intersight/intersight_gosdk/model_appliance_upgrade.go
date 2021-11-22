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

// ApplianceUpgrade Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's upgrade service automatically creates an Upgrade managed object when there is a pending software upgrade. User may modify an active Upgrade managed object to reset the software upgrade start time. However, user may not be able to postpone an upgrade beyond the limit enforced by the upgrade grace period set in the software manifest.
type ApplianceUpgrade struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if the software upgrade is active or not.
	Active *bool `json:"Active,omitempty"`
	// Indicates that the request was automatically created by the system.
	AutoCreated     *bool                      `json:"AutoCreated,omitempty"`
	CompletedPhases []OnpremUpgradePhase       `json:"CompletedPhases,omitempty"`
	CurrentPhase    NullableOnpremUpgradePhase `json:"CurrentPhase,omitempty"`
	// Description of the software upgrade.
	Description *string `json:"Description,omitempty"`
	// Elapsed time in seconds during the software upgrade.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
	// End date of the software upgrade.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.
	ErrorCode *int64 `json:"ErrorCode,omitempty"`
	// Software upgrade manifest's fingerprint.
	Fingerprint *string `json:"Fingerprint,omitempty"`
	// Track if software upgrade is upgrading or rolling back.
	IsRollingBack *bool `json:"IsRollingBack,omitempty"`
	// Indicates if the upgrade is triggered by user or due to schedule.
	IsUserTriggered *bool    `json:"IsUserTriggered,omitempty"`
	Messages        []string `json:"Messages,omitempty"`
	// Track if rollback is needed.
	RollbackNeeded *bool                `json:"RollbackNeeded,omitempty"`
	RollbackPhases []OnpremUpgradePhase `json:"RollbackPhases,omitempty"`
	// Status of the Intersight Appliance's software rollback status.
	RollbackStatus *string  `json:"RollbackStatus,omitempty"`
	Services       []string `json:"Services,omitempty"`
	// Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Status of the Intersight Appliance's software upgrade.
	Status *string `json:"Status,omitempty"`
	// TotalPhase represents the total number of the upgradePhases for one upgrade.
	TotalPhases *int64   `json:"TotalPhases,omitempty"`
	UiPackages  []string `json:"UiPackages,omitempty"`
	// Software upgrade manifest's version.
	Version              *string                           `json:"Version,omitempty"`
	Account              *IamAccountRelationship           `json:"Account,omitempty"`
	ImageBundle          *ApplianceImageBundleRelationship `json:"ImageBundle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceUpgrade ApplianceUpgrade

// NewApplianceUpgrade instantiates a new ApplianceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceUpgrade(classId string, objectType string) *ApplianceUpgrade {
	this := ApplianceUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var rollbackNeeded bool = false
	this.RollbackNeeded = &rollbackNeeded
	return &this
}

// NewApplianceUpgradeWithDefaults instantiates a new ApplianceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceUpgradeWithDefaults() *ApplianceUpgrade {
	this := ApplianceUpgrade{}
	var classId string = "appliance.Upgrade"
	this.ClassId = classId
	var objectType string = "appliance.Upgrade"
	this.ObjectType = objectType
	var rollbackNeeded bool = false
	this.RollbackNeeded = &rollbackNeeded
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApplianceUpgrade) SetActive(v bool) {
	o.Active = &v
}

// GetAutoCreated returns the AutoCreated field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetAutoCreated() bool {
	if o == nil || o.AutoCreated == nil {
		var ret bool
		return ret
	}
	return *o.AutoCreated
}

// GetAutoCreatedOk returns a tuple with the AutoCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetAutoCreatedOk() (*bool, bool) {
	if o == nil || o.AutoCreated == nil {
		return nil, false
	}
	return o.AutoCreated, true
}

// HasAutoCreated returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasAutoCreated() bool {
	if o != nil && o.AutoCreated != nil {
		return true
	}

	return false
}

// SetAutoCreated gets a reference to the given bool and assigns it to the AutoCreated field.
func (o *ApplianceUpgrade) SetAutoCreated(v bool) {
	o.AutoCreated = &v
}

// GetCompletedPhases returns the CompletedPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetCompletedPhases() []OnpremUpgradePhase {
	if o == nil {
		var ret []OnpremUpgradePhase
		return ret
	}
	return o.CompletedPhases
}

// GetCompletedPhasesOk returns a tuple with the CompletedPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetCompletedPhasesOk() (*[]OnpremUpgradePhase, bool) {
	if o == nil || o.CompletedPhases == nil {
		return nil, false
	}
	return &o.CompletedPhases, true
}

// HasCompletedPhases returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasCompletedPhases() bool {
	if o != nil && o.CompletedPhases != nil {
		return true
	}

	return false
}

// SetCompletedPhases gets a reference to the given []OnpremUpgradePhase and assigns it to the CompletedPhases field.
func (o *ApplianceUpgrade) SetCompletedPhases(v []OnpremUpgradePhase) {
	o.CompletedPhases = v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetCurrentPhase() OnpremUpgradePhase {
	if o == nil || o.CurrentPhase.Get() == nil {
		var ret OnpremUpgradePhase
		return ret
	}
	return *o.CurrentPhase.Get()
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetCurrentPhaseOk() (*OnpremUpgradePhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPhase.Get(), o.CurrentPhase.IsSet()
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasCurrentPhase() bool {
	if o != nil && o.CurrentPhase.IsSet() {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given NullableOnpremUpgradePhase and assigns it to the CurrentPhase field.
func (o *ApplianceUpgrade) SetCurrentPhase(v OnpremUpgradePhase) {
	o.CurrentPhase.Set(&v)
}

// SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil
func (o *ApplianceUpgrade) SetCurrentPhaseNil() {
	o.CurrentPhase.Set(nil)
}

// UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
func (o *ApplianceUpgrade) UnsetCurrentPhase() {
	o.CurrentPhase.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceUpgrade) SetDescription(v string) {
	o.Description = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *ApplianceUpgrade) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceUpgrade) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *ApplianceUpgrade) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *ApplianceUpgrade) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetIsRollingBack returns the IsRollingBack field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetIsRollingBack() bool {
	if o == nil || o.IsRollingBack == nil {
		var ret bool
		return ret
	}
	return *o.IsRollingBack
}

// GetIsRollingBackOk returns a tuple with the IsRollingBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetIsRollingBackOk() (*bool, bool) {
	if o == nil || o.IsRollingBack == nil {
		return nil, false
	}
	return o.IsRollingBack, true
}

// HasIsRollingBack returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasIsRollingBack() bool {
	if o != nil && o.IsRollingBack != nil {
		return true
	}

	return false
}

// SetIsRollingBack gets a reference to the given bool and assigns it to the IsRollingBack field.
func (o *ApplianceUpgrade) SetIsRollingBack(v bool) {
	o.IsRollingBack = &v
}

// GetIsUserTriggered returns the IsUserTriggered field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetIsUserTriggered() bool {
	if o == nil || o.IsUserTriggered == nil {
		var ret bool
		return ret
	}
	return *o.IsUserTriggered
}

// GetIsUserTriggeredOk returns a tuple with the IsUserTriggered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetIsUserTriggeredOk() (*bool, bool) {
	if o == nil || o.IsUserTriggered == nil {
		return nil, false
	}
	return o.IsUserTriggered, true
}

// HasIsUserTriggered returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasIsUserTriggered() bool {
	if o != nil && o.IsUserTriggered != nil {
		return true
	}

	return false
}

// SetIsUserTriggered gets a reference to the given bool and assigns it to the IsUserTriggered field.
func (o *ApplianceUpgrade) SetIsUserTriggered(v bool) {
	o.IsUserTriggered = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return &o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceUpgrade) SetMessages(v []string) {
	o.Messages = v
}

// GetRollbackNeeded returns the RollbackNeeded field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetRollbackNeeded() bool {
	if o == nil || o.RollbackNeeded == nil {
		var ret bool
		return ret
	}
	return *o.RollbackNeeded
}

// GetRollbackNeededOk returns a tuple with the RollbackNeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetRollbackNeededOk() (*bool, bool) {
	if o == nil || o.RollbackNeeded == nil {
		return nil, false
	}
	return o.RollbackNeeded, true
}

// HasRollbackNeeded returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasRollbackNeeded() bool {
	if o != nil && o.RollbackNeeded != nil {
		return true
	}

	return false
}

// SetRollbackNeeded gets a reference to the given bool and assigns it to the RollbackNeeded field.
func (o *ApplianceUpgrade) SetRollbackNeeded(v bool) {
	o.RollbackNeeded = &v
}

// GetRollbackPhases returns the RollbackPhases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetRollbackPhases() []OnpremUpgradePhase {
	if o == nil {
		var ret []OnpremUpgradePhase
		return ret
	}
	return o.RollbackPhases
}

// GetRollbackPhasesOk returns a tuple with the RollbackPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetRollbackPhasesOk() (*[]OnpremUpgradePhase, bool) {
	if o == nil || o.RollbackPhases == nil {
		return nil, false
	}
	return &o.RollbackPhases, true
}

// HasRollbackPhases returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasRollbackPhases() bool {
	if o != nil && o.RollbackPhases != nil {
		return true
	}

	return false
}

// SetRollbackPhases gets a reference to the given []OnpremUpgradePhase and assigns it to the RollbackPhases field.
func (o *ApplianceUpgrade) SetRollbackPhases(v []OnpremUpgradePhase) {
	o.RollbackPhases = v
}

// GetRollbackStatus returns the RollbackStatus field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetRollbackStatus() string {
	if o == nil || o.RollbackStatus == nil {
		var ret string
		return ret
	}
	return *o.RollbackStatus
}

// GetRollbackStatusOk returns a tuple with the RollbackStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetRollbackStatusOk() (*string, bool) {
	if o == nil || o.RollbackStatus == nil {
		return nil, false
	}
	return o.RollbackStatus, true
}

// HasRollbackStatus returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasRollbackStatus() bool {
	if o != nil && o.RollbackStatus != nil {
		return true
	}

	return false
}

// SetRollbackStatus gets a reference to the given string and assigns it to the RollbackStatus field.
func (o *ApplianceUpgrade) SetRollbackStatus(v string) {
	o.RollbackStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ApplianceUpgrade) SetServices(v []string) {
	o.Services = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceUpgrade) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceUpgrade) SetStatus(v string) {
	o.Status = &v
}

// GetTotalPhases returns the TotalPhases field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetTotalPhases() int64 {
	if o == nil || o.TotalPhases == nil {
		var ret int64
		return ret
	}
	return *o.TotalPhases
}

// GetTotalPhasesOk returns a tuple with the TotalPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetTotalPhasesOk() (*int64, bool) {
	if o == nil || o.TotalPhases == nil {
		return nil, false
	}
	return o.TotalPhases, true
}

// HasTotalPhases returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasTotalPhases() bool {
	if o != nil && o.TotalPhases != nil {
		return true
	}

	return false
}

// SetTotalPhases gets a reference to the given int64 and assigns it to the TotalPhases field.
func (o *ApplianceUpgrade) SetTotalPhases(v int64) {
	o.TotalPhases = &v
}

// GetUiPackages returns the UiPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceUpgrade) GetUiPackages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UiPackages
}

// GetUiPackagesOk returns a tuple with the UiPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceUpgrade) GetUiPackagesOk() (*[]string, bool) {
	if o == nil || o.UiPackages == nil {
		return nil, false
	}
	return &o.UiPackages, true
}

// HasUiPackages returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasUiPackages() bool {
	if o != nil && o.UiPackages != nil {
		return true
	}

	return false
}

// SetUiPackages gets a reference to the given []string and assigns it to the UiPackages field.
func (o *ApplianceUpgrade) SetUiPackages(v []string) {
	o.UiPackages = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceUpgrade) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceUpgrade) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetImageBundle returns the ImageBundle field value if set, zero value otherwise.
func (o *ApplianceUpgrade) GetImageBundle() ApplianceImageBundleRelationship {
	if o == nil || o.ImageBundle == nil {
		var ret ApplianceImageBundleRelationship
		return ret
	}
	return *o.ImageBundle
}

// GetImageBundleOk returns a tuple with the ImageBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgrade) GetImageBundleOk() (*ApplianceImageBundleRelationship, bool) {
	if o == nil || o.ImageBundle == nil {
		return nil, false
	}
	return o.ImageBundle, true
}

// HasImageBundle returns a boolean if a field has been set.
func (o *ApplianceUpgrade) HasImageBundle() bool {
	if o != nil && o.ImageBundle != nil {
		return true
	}

	return false
}

// SetImageBundle gets a reference to the given ApplianceImageBundleRelationship and assigns it to the ImageBundle field.
func (o *ApplianceUpgrade) SetImageBundle(v ApplianceImageBundleRelationship) {
	o.ImageBundle = &v
}

func (o ApplianceUpgrade) MarshalJSON() ([]byte, error) {
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
	if o.Active != nil {
		toSerialize["Active"] = o.Active
	}
	if o.AutoCreated != nil {
		toSerialize["AutoCreated"] = o.AutoCreated
	}
	if o.CompletedPhases != nil {
		toSerialize["CompletedPhases"] = o.CompletedPhases
	}
	if o.CurrentPhase.IsSet() {
		toSerialize["CurrentPhase"] = o.CurrentPhase.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Fingerprint != nil {
		toSerialize["Fingerprint"] = o.Fingerprint
	}
	if o.IsRollingBack != nil {
		toSerialize["IsRollingBack"] = o.IsRollingBack
	}
	if o.IsUserTriggered != nil {
		toSerialize["IsUserTriggered"] = o.IsUserTriggered
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.RollbackNeeded != nil {
		toSerialize["RollbackNeeded"] = o.RollbackNeeded
	}
	if o.RollbackPhases != nil {
		toSerialize["RollbackPhases"] = o.RollbackPhases
	}
	if o.RollbackStatus != nil {
		toSerialize["RollbackStatus"] = o.RollbackStatus
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TotalPhases != nil {
		toSerialize["TotalPhases"] = o.TotalPhases
	}
	if o.UiPackages != nil {
		toSerialize["UiPackages"] = o.UiPackages
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ImageBundle != nil {
		toSerialize["ImageBundle"] = o.ImageBundle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates if the software upgrade is active or not.
		Active *bool `json:"Active,omitempty"`
		// Indicates that the request was automatically created by the system.
		AutoCreated     *bool                      `json:"AutoCreated,omitempty"`
		CompletedPhases []OnpremUpgradePhase       `json:"CompletedPhases,omitempty"`
		CurrentPhase    NullableOnpremUpgradePhase `json:"CurrentPhase,omitempty"`
		// Description of the software upgrade.
		Description *string `json:"Description,omitempty"`
		// Elapsed time in seconds during the software upgrade.
		ElapsedTime *int64 `json:"ElapsedTime,omitempty"`
		// End date of the software upgrade.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.
		ErrorCode *int64 `json:"ErrorCode,omitempty"`
		// Software upgrade manifest's fingerprint.
		Fingerprint *string `json:"Fingerprint,omitempty"`
		// Track if software upgrade is upgrading or rolling back.
		IsRollingBack *bool `json:"IsRollingBack,omitempty"`
		// Indicates if the upgrade is triggered by user or due to schedule.
		IsUserTriggered *bool    `json:"IsUserTriggered,omitempty"`
		Messages        []string `json:"Messages,omitempty"`
		// Track if rollback is needed.
		RollbackNeeded *bool                `json:"RollbackNeeded,omitempty"`
		RollbackPhases []OnpremUpgradePhase `json:"RollbackPhases,omitempty"`
		// Status of the Intersight Appliance's software rollback status.
		RollbackStatus *string  `json:"RollbackStatus,omitempty"`
		Services       []string `json:"Services,omitempty"`
		// Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// Status of the Intersight Appliance's software upgrade.
		Status *string `json:"Status,omitempty"`
		// TotalPhase represents the total number of the upgradePhases for one upgrade.
		TotalPhases *int64   `json:"TotalPhases,omitempty"`
		UiPackages  []string `json:"UiPackages,omitempty"`
		// Software upgrade manifest's version.
		Version     *string                           `json:"Version,omitempty"`
		Account     *IamAccountRelationship           `json:"Account,omitempty"`
		ImageBundle *ApplianceImageBundleRelationship `json:"ImageBundle,omitempty"`
	}

	varApplianceUpgradeWithoutEmbeddedStruct := ApplianceUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varApplianceUpgrade := _ApplianceUpgrade{}
		varApplianceUpgrade.ClassId = varApplianceUpgradeWithoutEmbeddedStruct.ClassId
		varApplianceUpgrade.ObjectType = varApplianceUpgradeWithoutEmbeddedStruct.ObjectType
		varApplianceUpgrade.Active = varApplianceUpgradeWithoutEmbeddedStruct.Active
		varApplianceUpgrade.AutoCreated = varApplianceUpgradeWithoutEmbeddedStruct.AutoCreated
		varApplianceUpgrade.CompletedPhases = varApplianceUpgradeWithoutEmbeddedStruct.CompletedPhases
		varApplianceUpgrade.CurrentPhase = varApplianceUpgradeWithoutEmbeddedStruct.CurrentPhase
		varApplianceUpgrade.Description = varApplianceUpgradeWithoutEmbeddedStruct.Description
		varApplianceUpgrade.ElapsedTime = varApplianceUpgradeWithoutEmbeddedStruct.ElapsedTime
		varApplianceUpgrade.EndTime = varApplianceUpgradeWithoutEmbeddedStruct.EndTime
		varApplianceUpgrade.ErrorCode = varApplianceUpgradeWithoutEmbeddedStruct.ErrorCode
		varApplianceUpgrade.Fingerprint = varApplianceUpgradeWithoutEmbeddedStruct.Fingerprint
		varApplianceUpgrade.IsRollingBack = varApplianceUpgradeWithoutEmbeddedStruct.IsRollingBack
		varApplianceUpgrade.IsUserTriggered = varApplianceUpgradeWithoutEmbeddedStruct.IsUserTriggered
		varApplianceUpgrade.Messages = varApplianceUpgradeWithoutEmbeddedStruct.Messages
		varApplianceUpgrade.RollbackNeeded = varApplianceUpgradeWithoutEmbeddedStruct.RollbackNeeded
		varApplianceUpgrade.RollbackPhases = varApplianceUpgradeWithoutEmbeddedStruct.RollbackPhases
		varApplianceUpgrade.RollbackStatus = varApplianceUpgradeWithoutEmbeddedStruct.RollbackStatus
		varApplianceUpgrade.Services = varApplianceUpgradeWithoutEmbeddedStruct.Services
		varApplianceUpgrade.StartTime = varApplianceUpgradeWithoutEmbeddedStruct.StartTime
		varApplianceUpgrade.Status = varApplianceUpgradeWithoutEmbeddedStruct.Status
		varApplianceUpgrade.TotalPhases = varApplianceUpgradeWithoutEmbeddedStruct.TotalPhases
		varApplianceUpgrade.UiPackages = varApplianceUpgradeWithoutEmbeddedStruct.UiPackages
		varApplianceUpgrade.Version = varApplianceUpgradeWithoutEmbeddedStruct.Version
		varApplianceUpgrade.Account = varApplianceUpgradeWithoutEmbeddedStruct.Account
		varApplianceUpgrade.ImageBundle = varApplianceUpgradeWithoutEmbeddedStruct.ImageBundle
		*o = ApplianceUpgrade(varApplianceUpgrade)
	} else {
		return err
	}

	varApplianceUpgrade := _ApplianceUpgrade{}

	err = json.Unmarshal(bytes, &varApplianceUpgrade)
	if err == nil {
		o.MoBaseMo = varApplianceUpgrade.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Active")
		delete(additionalProperties, "AutoCreated")
		delete(additionalProperties, "CompletedPhases")
		delete(additionalProperties, "CurrentPhase")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ElapsedTime")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Fingerprint")
		delete(additionalProperties, "IsRollingBack")
		delete(additionalProperties, "IsUserTriggered")
		delete(additionalProperties, "Messages")
		delete(additionalProperties, "RollbackNeeded")
		delete(additionalProperties, "RollbackPhases")
		delete(additionalProperties, "RollbackStatus")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TotalPhases")
		delete(additionalProperties, "UiPackages")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ImageBundle")

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

type NullableApplianceUpgrade struct {
	value *ApplianceUpgrade
	isSet bool
}

func (v NullableApplianceUpgrade) Get() *ApplianceUpgrade {
	return v.value
}

func (v *NullableApplianceUpgrade) Set(val *ApplianceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceUpgrade(val *ApplianceUpgrade) *NullableApplianceUpgrade {
	return &NullableApplianceUpgrade{value: val, isSet: true}
}

func (v NullableApplianceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
