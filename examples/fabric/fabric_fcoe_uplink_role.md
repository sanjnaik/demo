### Resource Creation

```hcl
resource "intersight_fabric_fcoe_uplink_role" "fabric_fcoe_uplink_role1" {
  pcid              = 100
  admin_speed       = "Auto"
  ud_ld_admin_state = true
}
```