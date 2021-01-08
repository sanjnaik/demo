---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_appliance_pc_role"
description: |-
  Configuration object sent by user to create an appliance port channel.
---

# Resource: intersight_fabric_appliance_pc_role
Configuration object sent by user to create an appliance port channel.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `admin_speed`:(string) Admin configured speed for the port channel.* `Auto` - Admin configurable speed AUTO ( default ).* `1Gbps` - Admin configurable speed 1Gbps.* `10Gbps` - Admin configurable speed 10Gbps.* `25Gbps` - Admin configurable speed 25Gbps.* `40Gbps` - Admin configurable speed 40Gbps.* `100Gbps` - Admin configurable speed 100Gbps. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `eth_network_control_policy`:(Array with Maximum of one item) - A reference to a fabricEthNetworkControlPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `eth_network_group_policy`:(Array with Maximum of one item) - A reference to a fabricEthNetworkGroupPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mode`:(string) Port mode to be set on the appliance port-channel.* `trunk` - Trunk Mode Switch Port Type.* `access` - Access Mode Switch Port Type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pc_id`:(int) Unique Identifier of the port-channel, local to this switch. 
* `port_policy`:(Array with Maximum of one item) - A reference to a fabricPortPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `ports`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `port_id`:(int) Port Identifier of the Switch/FEX/Chassis Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
  + `slot_id`:(int) Slot Identifier of the Switch/FEX/Chassis Interface. 
* `priority`:(string) The 'name' of the System QoS Class.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_fabric_appliance_pc_role` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_appliance_pc_role.example 1234567890987654321abcde
```