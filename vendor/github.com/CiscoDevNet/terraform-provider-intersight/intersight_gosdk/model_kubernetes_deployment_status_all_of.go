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

// KubernetesDeploymentStatusAllOf Definition of the list of properties defined in 'kubernetes.DeploymentStatus', excluding properties defined in parent classes.
type KubernetesDeploymentStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
	// The generation observed by the deployment controller.
	ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
	// Total number of ready pods targeted by this deployment.
	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty"`
	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Replicas *int64 `json:"Replicas,omitempty"`
	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	UpdatedReplicas      *int64 `json:"UpdatedReplicas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesDeploymentStatusAllOf KubernetesDeploymentStatusAllOf

// NewKubernetesDeploymentStatusAllOf instantiates a new KubernetesDeploymentStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesDeploymentStatusAllOf(classId string, objectType string) *KubernetesDeploymentStatusAllOf {
	this := KubernetesDeploymentStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var availableReplicas int64 = 0
	this.AvailableReplicas = &availableReplicas
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	var readyReplicas int64 = 0
	this.ReadyReplicas = &readyReplicas
	var replicas int64 = 0
	this.Replicas = &replicas
	var updatedReplicas int64 = 0
	this.UpdatedReplicas = &updatedReplicas
	return &this
}

// NewKubernetesDeploymentStatusAllOfWithDefaults instantiates a new KubernetesDeploymentStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesDeploymentStatusAllOfWithDefaults() *KubernetesDeploymentStatusAllOf {
	this := KubernetesDeploymentStatusAllOf{}
	var classId string = "kubernetes.DeploymentStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.DeploymentStatus"
	this.ObjectType = objectType
	var availableReplicas int64 = 0
	this.AvailableReplicas = &availableReplicas
	var observedGeneration int64 = 0
	this.ObservedGeneration = &observedGeneration
	var readyReplicas int64 = 0
	this.ReadyReplicas = &readyReplicas
	var replicas int64 = 0
	this.Replicas = &replicas
	var updatedReplicas int64 = 0
	this.UpdatedReplicas = &updatedReplicas
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesDeploymentStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesDeploymentStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesDeploymentStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesDeploymentStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvailableReplicas returns the AvailableReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatusAllOf) GetAvailableReplicas() int64 {
	if o == nil || o.AvailableReplicas == nil {
		var ret int64
		return ret
	}
	return *o.AvailableReplicas
}

// GetAvailableReplicasOk returns a tuple with the AvailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetAvailableReplicasOk() (*int64, bool) {
	if o == nil || o.AvailableReplicas == nil {
		return nil, false
	}
	return o.AvailableReplicas, true
}

// HasAvailableReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatusAllOf) HasAvailableReplicas() bool {
	if o != nil && o.AvailableReplicas != nil {
		return true
	}

	return false
}

// SetAvailableReplicas gets a reference to the given int64 and assigns it to the AvailableReplicas field.
func (o *KubernetesDeploymentStatusAllOf) SetAvailableReplicas(v int64) {
	o.AvailableReplicas = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatusAllOf) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatusAllOf) HasObservedGeneration() bool {
	if o != nil && o.ObservedGeneration != nil {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubernetesDeploymentStatusAllOf) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetReadyReplicas returns the ReadyReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatusAllOf) GetReadyReplicas() int64 {
	if o == nil || o.ReadyReplicas == nil {
		var ret int64
		return ret
	}
	return *o.ReadyReplicas
}

// GetReadyReplicasOk returns a tuple with the ReadyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetReadyReplicasOk() (*int64, bool) {
	if o == nil || o.ReadyReplicas == nil {
		return nil, false
	}
	return o.ReadyReplicas, true
}

// HasReadyReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatusAllOf) HasReadyReplicas() bool {
	if o != nil && o.ReadyReplicas != nil {
		return true
	}

	return false
}

// SetReadyReplicas gets a reference to the given int64 and assigns it to the ReadyReplicas field.
func (o *KubernetesDeploymentStatusAllOf) SetReadyReplicas(v int64) {
	o.ReadyReplicas = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatusAllOf) GetReplicas() int64 {
	if o == nil || o.Replicas == nil {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetReplicasOk() (*int64, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatusAllOf) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *KubernetesDeploymentStatusAllOf) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetUpdatedReplicas returns the UpdatedReplicas field value if set, zero value otherwise.
func (o *KubernetesDeploymentStatusAllOf) GetUpdatedReplicas() int64 {
	if o == nil || o.UpdatedReplicas == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedReplicas
}

// GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesDeploymentStatusAllOf) GetUpdatedReplicasOk() (*int64, bool) {
	if o == nil || o.UpdatedReplicas == nil {
		return nil, false
	}
	return o.UpdatedReplicas, true
}

// HasUpdatedReplicas returns a boolean if a field has been set.
func (o *KubernetesDeploymentStatusAllOf) HasUpdatedReplicas() bool {
	if o != nil && o.UpdatedReplicas != nil {
		return true
	}

	return false
}

// SetUpdatedReplicas gets a reference to the given int64 and assigns it to the UpdatedReplicas field.
func (o *KubernetesDeploymentStatusAllOf) SetUpdatedReplicas(v int64) {
	o.UpdatedReplicas = &v
}

func (o KubernetesDeploymentStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvailableReplicas != nil {
		toSerialize["AvailableReplicas"] = o.AvailableReplicas
	}
	if o.ObservedGeneration != nil {
		toSerialize["ObservedGeneration"] = o.ObservedGeneration
	}
	if o.ReadyReplicas != nil {
		toSerialize["ReadyReplicas"] = o.ReadyReplicas
	}
	if o.Replicas != nil {
		toSerialize["Replicas"] = o.Replicas
	}
	if o.UpdatedReplicas != nil {
		toSerialize["UpdatedReplicas"] = o.UpdatedReplicas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesDeploymentStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesDeploymentStatusAllOf := _KubernetesDeploymentStatusAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesDeploymentStatusAllOf); err == nil {
		*o = KubernetesDeploymentStatusAllOf(varKubernetesDeploymentStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailableReplicas")
		delete(additionalProperties, "ObservedGeneration")
		delete(additionalProperties, "ReadyReplicas")
		delete(additionalProperties, "Replicas")
		delete(additionalProperties, "UpdatedReplicas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesDeploymentStatusAllOf struct {
	value *KubernetesDeploymentStatusAllOf
	isSet bool
}

func (v NullableKubernetesDeploymentStatusAllOf) Get() *KubernetesDeploymentStatusAllOf {
	return v.value
}

func (v *NullableKubernetesDeploymentStatusAllOf) Set(val *KubernetesDeploymentStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesDeploymentStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesDeploymentStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesDeploymentStatusAllOf(val *KubernetesDeploymentStatusAllOf) *NullableKubernetesDeploymentStatusAllOf {
	return &NullableKubernetesDeploymentStatusAllOf{value: val, isSet: true}
}

func (v NullableKubernetesDeploymentStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesDeploymentStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
