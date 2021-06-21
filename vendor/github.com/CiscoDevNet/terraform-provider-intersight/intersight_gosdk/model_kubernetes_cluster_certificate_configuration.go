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
	"reflect"
	"strings"
)

// KubernetesClusterCertificateConfiguration Certifcate and key configuration for Kubernetes cluster creation.
type KubernetesClusterCertificateConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Certificate for the Kubernetes API server.
	CaCert *string `json:"CaCert,omitempty"`
	// Private Key for the Kubernetes API server.
	CaKey *string `json:"CaKey,omitempty"`
	// Certificate for the etcd cluster.
	EtcdCert          *string  `json:"EtcdCert,omitempty"`
	EtcdEncryptionKey []string `json:"EtcdEncryptionKey,omitempty"`
	// Private key for the etcd cluster.
	EtcdKey *string `json:"EtcdKey,omitempty"`
	// Certificate for the front proxy to support Kubernetes API aggregators.
	FrontProxyCert *string `json:"FrontProxyCert,omitempty"`
	// Private key for the front proxy to support Kubernetes API aggregators.
	FrontProxyKey *string `json:"FrontProxyKey,omitempty"`
	// Service account private key used by Kubernetes Token Controller to sign generated service account tokens.
	SaPrivateKey *string `json:"SaPrivateKey,omitempty"`
	// Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller.
	SaPublicKey          *string `json:"SaPublicKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterCertificateConfiguration KubernetesClusterCertificateConfiguration

// NewKubernetesClusterCertificateConfiguration instantiates a new KubernetesClusterCertificateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterCertificateConfiguration(classId string, objectType string) *KubernetesClusterCertificateConfiguration {
	this := KubernetesClusterCertificateConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterCertificateConfigurationWithDefaults instantiates a new KubernetesClusterCertificateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterCertificateConfigurationWithDefaults() *KubernetesClusterCertificateConfiguration {
	this := KubernetesClusterCertificateConfiguration{}
	var classId string = "kubernetes.ClusterCertificateConfiguration"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterCertificateConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterCertificateConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterCertificateConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterCertificateConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterCertificateConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetCaCert() string {
	if o == nil || o.CaCert == nil {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetCaCertOk() (*string, bool) {
	if o == nil || o.CaCert == nil {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasCaCert() bool {
	if o != nil && o.CaCert != nil {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *KubernetesClusterCertificateConfiguration) SetCaCert(v string) {
	o.CaCert = &v
}

// GetCaKey returns the CaKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetCaKey() string {
	if o == nil || o.CaKey == nil {
		var ret string
		return ret
	}
	return *o.CaKey
}

// GetCaKeyOk returns a tuple with the CaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetCaKeyOk() (*string, bool) {
	if o == nil || o.CaKey == nil {
		return nil, false
	}
	return o.CaKey, true
}

// HasCaKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasCaKey() bool {
	if o != nil && o.CaKey != nil {
		return true
	}

	return false
}

// SetCaKey gets a reference to the given string and assigns it to the CaKey field.
func (o *KubernetesClusterCertificateConfiguration) SetCaKey(v string) {
	o.CaKey = &v
}

// GetEtcdCert returns the EtcdCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdCert() string {
	if o == nil || o.EtcdCert == nil {
		var ret string
		return ret
	}
	return *o.EtcdCert
}

// GetEtcdCertOk returns a tuple with the EtcdCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdCertOk() (*string, bool) {
	if o == nil || o.EtcdCert == nil {
		return nil, false
	}
	return o.EtcdCert, true
}

// HasEtcdCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdCert() bool {
	if o != nil && o.EtcdCert != nil {
		return true
	}

	return false
}

// SetEtcdCert gets a reference to the given string and assigns it to the EtcdCert field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdCert(v string) {
	o.EtcdCert = &v
}

// GetEtcdEncryptionKey returns the EtcdEncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterCertificateConfiguration) GetEtcdEncryptionKey() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EtcdEncryptionKey
}

// GetEtcdEncryptionKeyOk returns a tuple with the EtcdEncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterCertificateConfiguration) GetEtcdEncryptionKeyOk() (*[]string, bool) {
	if o == nil || o.EtcdEncryptionKey == nil {
		return nil, false
	}
	return &o.EtcdEncryptionKey, true
}

// HasEtcdEncryptionKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdEncryptionKey() bool {
	if o != nil && o.EtcdEncryptionKey != nil {
		return true
	}

	return false
}

// SetEtcdEncryptionKey gets a reference to the given []string and assigns it to the EtcdEncryptionKey field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdEncryptionKey(v []string) {
	o.EtcdEncryptionKey = v
}

// GetEtcdKey returns the EtcdKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdKey() string {
	if o == nil || o.EtcdKey == nil {
		var ret string
		return ret
	}
	return *o.EtcdKey
}

// GetEtcdKeyOk returns a tuple with the EtcdKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetEtcdKeyOk() (*string, bool) {
	if o == nil || o.EtcdKey == nil {
		return nil, false
	}
	return o.EtcdKey, true
}

// HasEtcdKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasEtcdKey() bool {
	if o != nil && o.EtcdKey != nil {
		return true
	}

	return false
}

// SetEtcdKey gets a reference to the given string and assigns it to the EtcdKey field.
func (o *KubernetesClusterCertificateConfiguration) SetEtcdKey(v string) {
	o.EtcdKey = &v
}

// GetFrontProxyCert returns the FrontProxyCert field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyCert() string {
	if o == nil || o.FrontProxyCert == nil {
		var ret string
		return ret
	}
	return *o.FrontProxyCert
}

// GetFrontProxyCertOk returns a tuple with the FrontProxyCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyCertOk() (*string, bool) {
	if o == nil || o.FrontProxyCert == nil {
		return nil, false
	}
	return o.FrontProxyCert, true
}

// HasFrontProxyCert returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasFrontProxyCert() bool {
	if o != nil && o.FrontProxyCert != nil {
		return true
	}

	return false
}

// SetFrontProxyCert gets a reference to the given string and assigns it to the FrontProxyCert field.
func (o *KubernetesClusterCertificateConfiguration) SetFrontProxyCert(v string) {
	o.FrontProxyCert = &v
}

// GetFrontProxyKey returns the FrontProxyKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyKey() string {
	if o == nil || o.FrontProxyKey == nil {
		var ret string
		return ret
	}
	return *o.FrontProxyKey
}

// GetFrontProxyKeyOk returns a tuple with the FrontProxyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetFrontProxyKeyOk() (*string, bool) {
	if o == nil || o.FrontProxyKey == nil {
		return nil, false
	}
	return o.FrontProxyKey, true
}

// HasFrontProxyKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasFrontProxyKey() bool {
	if o != nil && o.FrontProxyKey != nil {
		return true
	}

	return false
}

// SetFrontProxyKey gets a reference to the given string and assigns it to the FrontProxyKey field.
func (o *KubernetesClusterCertificateConfiguration) SetFrontProxyKey(v string) {
	o.FrontProxyKey = &v
}

// GetSaPrivateKey returns the SaPrivateKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetSaPrivateKey() string {
	if o == nil || o.SaPrivateKey == nil {
		var ret string
		return ret
	}
	return *o.SaPrivateKey
}

// GetSaPrivateKeyOk returns a tuple with the SaPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetSaPrivateKeyOk() (*string, bool) {
	if o == nil || o.SaPrivateKey == nil {
		return nil, false
	}
	return o.SaPrivateKey, true
}

// HasSaPrivateKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasSaPrivateKey() bool {
	if o != nil && o.SaPrivateKey != nil {
		return true
	}

	return false
}

// SetSaPrivateKey gets a reference to the given string and assigns it to the SaPrivateKey field.
func (o *KubernetesClusterCertificateConfiguration) SetSaPrivateKey(v string) {
	o.SaPrivateKey = &v
}

// GetSaPublicKey returns the SaPublicKey field value if set, zero value otherwise.
func (o *KubernetesClusterCertificateConfiguration) GetSaPublicKey() string {
	if o == nil || o.SaPublicKey == nil {
		var ret string
		return ret
	}
	return *o.SaPublicKey
}

// GetSaPublicKeyOk returns a tuple with the SaPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterCertificateConfiguration) GetSaPublicKeyOk() (*string, bool) {
	if o == nil || o.SaPublicKey == nil {
		return nil, false
	}
	return o.SaPublicKey, true
}

// HasSaPublicKey returns a boolean if a field has been set.
func (o *KubernetesClusterCertificateConfiguration) HasSaPublicKey() bool {
	if o != nil && o.SaPublicKey != nil {
		return true
	}

	return false
}

// SetSaPublicKey gets a reference to the given string and assigns it to the SaPublicKey field.
func (o *KubernetesClusterCertificateConfiguration) SetSaPublicKey(v string) {
	o.SaPublicKey = &v
}

func (o KubernetesClusterCertificateConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.CaCert != nil {
		toSerialize["CaCert"] = o.CaCert
	}
	if o.CaKey != nil {
		toSerialize["CaKey"] = o.CaKey
	}
	if o.EtcdCert != nil {
		toSerialize["EtcdCert"] = o.EtcdCert
	}
	if o.EtcdEncryptionKey != nil {
		toSerialize["EtcdEncryptionKey"] = o.EtcdEncryptionKey
	}
	if o.EtcdKey != nil {
		toSerialize["EtcdKey"] = o.EtcdKey
	}
	if o.FrontProxyCert != nil {
		toSerialize["FrontProxyCert"] = o.FrontProxyCert
	}
	if o.FrontProxyKey != nil {
		toSerialize["FrontProxyKey"] = o.FrontProxyKey
	}
	if o.SaPrivateKey != nil {
		toSerialize["SaPrivateKey"] = o.SaPrivateKey
	}
	if o.SaPublicKey != nil {
		toSerialize["SaPublicKey"] = o.SaPublicKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterCertificateConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesClusterCertificateConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Certificate for the Kubernetes API server.
		CaCert *string `json:"CaCert,omitempty"`
		// Private Key for the Kubernetes API server.
		CaKey *string `json:"CaKey,omitempty"`
		// Certificate for the etcd cluster.
		EtcdCert          *string  `json:"EtcdCert,omitempty"`
		EtcdEncryptionKey []string `json:"EtcdEncryptionKey,omitempty"`
		// Private key for the etcd cluster.
		EtcdKey *string `json:"EtcdKey,omitempty"`
		// Certificate for the front proxy to support Kubernetes API aggregators.
		FrontProxyCert *string `json:"FrontProxyCert,omitempty"`
		// Private key for the front proxy to support Kubernetes API aggregators.
		FrontProxyKey *string `json:"FrontProxyKey,omitempty"`
		// Service account private key used by Kubernetes Token Controller to sign generated service account tokens.
		SaPrivateKey *string `json:"SaPrivateKey,omitempty"`
		// Service account public key used by Kubernetes API Server to validate service account tokens generated by the Token Controller.
		SaPublicKey *string `json:"SaPublicKey,omitempty"`
	}

	varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct := KubernetesClusterCertificateConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesClusterCertificateConfiguration := _KubernetesClusterCertificateConfiguration{}
		varKubernetesClusterCertificateConfiguration.ClassId = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.ClassId
		varKubernetesClusterCertificateConfiguration.ObjectType = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.ObjectType
		varKubernetesClusterCertificateConfiguration.CaCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.CaCert
		varKubernetesClusterCertificateConfiguration.CaKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.CaKey
		varKubernetesClusterCertificateConfiguration.EtcdCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdCert
		varKubernetesClusterCertificateConfiguration.EtcdEncryptionKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdEncryptionKey
		varKubernetesClusterCertificateConfiguration.EtcdKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.EtcdKey
		varKubernetesClusterCertificateConfiguration.FrontProxyCert = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.FrontProxyCert
		varKubernetesClusterCertificateConfiguration.FrontProxyKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.FrontProxyKey
		varKubernetesClusterCertificateConfiguration.SaPrivateKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.SaPrivateKey
		varKubernetesClusterCertificateConfiguration.SaPublicKey = varKubernetesClusterCertificateConfigurationWithoutEmbeddedStruct.SaPublicKey
		*o = KubernetesClusterCertificateConfiguration(varKubernetesClusterCertificateConfiguration)
	} else {
		return err
	}

	varKubernetesClusterCertificateConfiguration := _KubernetesClusterCertificateConfiguration{}

	err = json.Unmarshal(bytes, &varKubernetesClusterCertificateConfiguration)
	if err == nil {
		o.MoBaseComplexType = varKubernetesClusterCertificateConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CaCert")
		delete(additionalProperties, "CaKey")
		delete(additionalProperties, "EtcdCert")
		delete(additionalProperties, "EtcdEncryptionKey")
		delete(additionalProperties, "EtcdKey")
		delete(additionalProperties, "FrontProxyCert")
		delete(additionalProperties, "FrontProxyKey")
		delete(additionalProperties, "SaPrivateKey")
		delete(additionalProperties, "SaPublicKey")

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

type NullableKubernetesClusterCertificateConfiguration struct {
	value *KubernetesClusterCertificateConfiguration
	isSet bool
}

func (v NullableKubernetesClusterCertificateConfiguration) Get() *KubernetesClusterCertificateConfiguration {
	return v.value
}

func (v *NullableKubernetesClusterCertificateConfiguration) Set(val *KubernetesClusterCertificateConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterCertificateConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterCertificateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterCertificateConfiguration(val *KubernetesClusterCertificateConfiguration) *NullableKubernetesClusterCertificateConfiguration {
	return &NullableKubernetesClusterCertificateConfiguration{value: val, isSet: true}
}

func (v NullableKubernetesClusterCertificateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterCertificateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
