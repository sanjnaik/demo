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
	"fmt"
)

// VirtualizationVmwareVirtualDiskResponse - The response body of a HTTP GET request for the 'virtualization.VmwareVirtualDisk' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'virtualization.VmwareVirtualDisk' resources.
type VirtualizationVmwareVirtualDiskResponse struct {
	MoAggregateTransform                *MoAggregateTransform
	MoDocumentCount                     *MoDocumentCount
	MoTagSummary                        *MoTagSummary
	VirtualizationVmwareVirtualDiskList *VirtualizationVmwareVirtualDiskList
}

// MoAggregateTransformAsVirtualizationVmwareVirtualDiskResponse is a convenience function that returns MoAggregateTransform wrapped in VirtualizationVmwareVirtualDiskResponse
func MoAggregateTransformAsVirtualizationVmwareVirtualDiskResponse(v *MoAggregateTransform) VirtualizationVmwareVirtualDiskResponse {
	return VirtualizationVmwareVirtualDiskResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsVirtualizationVmwareVirtualDiskResponse is a convenience function that returns MoDocumentCount wrapped in VirtualizationVmwareVirtualDiskResponse
func MoDocumentCountAsVirtualizationVmwareVirtualDiskResponse(v *MoDocumentCount) VirtualizationVmwareVirtualDiskResponse {
	return VirtualizationVmwareVirtualDiskResponse{MoDocumentCount: v}
}

// MoTagSummaryAsVirtualizationVmwareVirtualDiskResponse is a convenience function that returns MoTagSummary wrapped in VirtualizationVmwareVirtualDiskResponse
func MoTagSummaryAsVirtualizationVmwareVirtualDiskResponse(v *MoTagSummary) VirtualizationVmwareVirtualDiskResponse {
	return VirtualizationVmwareVirtualDiskResponse{MoTagSummary: v}
}

// VirtualizationVmwareVirtualDiskListAsVirtualizationVmwareVirtualDiskResponse is a convenience function that returns VirtualizationVmwareVirtualDiskList wrapped in VirtualizationVmwareVirtualDiskResponse
func VirtualizationVmwareVirtualDiskListAsVirtualizationVmwareVirtualDiskResponse(v *VirtualizationVmwareVirtualDiskList) VirtualizationVmwareVirtualDiskResponse {
	return VirtualizationVmwareVirtualDiskResponse{VirtualizationVmwareVirtualDiskList: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VirtualizationVmwareVirtualDiskResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal VirtualizationVmwareVirtualDiskResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal VirtualizationVmwareVirtualDiskResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal VirtualizationVmwareVirtualDiskResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'virtualization.VmwareVirtualDisk.List'
	if jsonDict["ObjectType"] == "virtualization.VmwareVirtualDisk.List" {
		// try to unmarshal JSON data into VirtualizationVmwareVirtualDiskList
		err = json.Unmarshal(data, &dst.VirtualizationVmwareVirtualDiskList)
		if err == nil {
			return nil // data stored in dst.VirtualizationVmwareVirtualDiskList, return on the first match
		} else {
			dst.VirtualizationVmwareVirtualDiskList = nil
			return fmt.Errorf("Failed to unmarshal VirtualizationVmwareVirtualDiskResponse as VirtualizationVmwareVirtualDiskList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VirtualizationVmwareVirtualDiskResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.VirtualizationVmwareVirtualDiskList != nil {
		return json.Marshal(&src.VirtualizationVmwareVirtualDiskList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VirtualizationVmwareVirtualDiskResponse) GetActualInstance() interface{} {
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.VirtualizationVmwareVirtualDiskList != nil {
		return obj.VirtualizationVmwareVirtualDiskList
	}

	// all schemas are nil
	return nil
}

type NullableVirtualizationVmwareVirtualDiskResponse struct {
	value *VirtualizationVmwareVirtualDiskResponse
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualDiskResponse) Get() *VirtualizationVmwareVirtualDiskResponse {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualDiskResponse) Set(val *VirtualizationVmwareVirtualDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualDiskResponse(val *VirtualizationVmwareVirtualDiskResponse) *NullableVirtualizationVmwareVirtualDiskResponse {
	return &NullableVirtualizationVmwareVirtualDiskResponse{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
