<!-- BEGIN_TF_DOCS -->
# vSphere Tags Terraform Module

This Terraform module either creates or imports a list of [tags](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-16422FF7-235B-4A44-92E2-532F6AED0923.html) grouped in a [tag category](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-16422FF7-235B-4A44-92E2-532F6AED0923.html) in your [VMware Cloud on AWS](https://aws.amazon.com/vmware/) or [VMware vSphere](https://docs.vmware.com/en/VMware-vSphere/index.html) on&#8209;premises environment.
You use these tags to label your inventory objects with metadata to make it easier to sort and search for these objects.

Tags and categories can span multiple [vCenter Server](https://docs.vmware.com/en/VMware-vSphere/index.html) instances.
When you use [Hybrid Linked Mode](https://docs.vmware.com/en/VMware-Cloud-on-AWS/services/com.vmware.vsphere.vmc-aws-manage-data-center-vms.doc/GUID-91C57891-4D61-4F4C-B580-74F3000B831D.html), tags and tag categories are maintained across your linked domain. So in this mode, the on&#8209;premises and VMware Cloud on AWS software&#8209;defined data centers (SDDCs) share tags and tag attributes.
If multiple on&#8209;premises vCenter Server instances are configured to use [Enhanced Linked Mode](https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-6ADB06EF-E342-457E-A17B-1EA31C0F6D4B.html), tags and categories are replicated across the vCenter Server instances.

## Usage

### Create new tags in a new tag category

```hcl
module "vsphere_tags" {
  source  = "aws-ia/tags/vsphere"
  version = ">= 0.0.1"

  tag_category_name        = "example-category"
  tag_category_description = "Example tag category."
  tag_category_cardinality = "MULTIPLE"
  create_tag_category      = true
  create_tags              = true

  tags = {
    terraform = "Managed by Terraform"
    project   = "terraform-vsphere-tags"
  }
}
```

### Create new tags in an existing tag category

```hcl
module "vsphere_tags" {
  source  = "aws-ia/tags/vsphere"
  version = ">= 0.0.1"

  tag_category_name   = "example-category"
  create_tag_category = false
  create_tags         = true

  tags = {
    terraform = "Managed by Terraform"
    project   = "terraform-vsphere-tags"
  }
}
```

### Import existing tags in an existing tag category

```hcl
module "vsphere_tags" {
  source  = "aws-ia/tags/vsphere"
  version = ">= 0.0.1"

  tag_category_name   = "example-category"
  create_tag_category = false
  create_tags         = false

  tags = {
    terraform = "Managed by Terraform"
    project   = "terraform-vsphere-tags"
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.1.0 |
| <a name="requirement_vsphere"></a> [vsphere](#requirement\_vsphere) | >= 2.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_vsphere"></a> [vsphere](#provider\_vsphere) | >= 2.1.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [vsphere_tag.tags](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag) | resource |
| [vsphere_tag_category.category](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category) | resource |
| [vsphere_tag.tags](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/tag) | data source |
| [vsphere_tag_category.category](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/data-sources/tag_category) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_tag_category_name"></a> [tag\_category\_name](#input\_tag\_category\_name) | The name of the vSphere tag category. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of strings defining vSphere tag names and descriptions. | `map(string)` | n/a | yes |
| <a name="input_create_tag_category"></a> [create\_tag\_category](#input\_create\_tag\_category) | If true, a new vSphere tag category will be created. | `bool` | `true` | no |
| <a name="input_create_tags"></a> [create\_tags](#input\_create\_tags) | If true, new vSphere tags will be created for each entry. | `bool` | `true` | no |
| <a name="input_tag_category_associable_types"></a> [tag\_category\_associable\_types](#input\_tag\_category\_associable\_types) | A list object types that this category to which this category can be assigned (https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category#associable-object-types). | `list(string)` | <pre>[<br>  "Folder",<br>  "ClusterComputeResource",<br>  "Datacenter",<br>  "Datastore",<br>  "StoragePod",<br>  "DistributedVirtualPortgroup",<br>  "DistributedVirtualSwitch",<br>  "VmwareDistributedVirtualSwitch",<br>  "HostSystem",<br>  "com.vmware.content.Library",<br>  "com.vmware.content.library.Item",<br>  "HostNetwork",<br>  "Network",<br>  "OpaqueNetwork",<br>  "ResourcePool",<br>  "VirtualApp",<br>  "VirtualMachine"<br>]</pre> | no |
| <a name="input_tag_category_cardinality"></a> [tag\_category\_cardinality](#input\_tag\_category\_cardinality) | The number of tags that can be assigned from this category to a single object at once. | `string` | `"MULTIPLE"` | no |
| <a name="input_tag_category_description"></a> [tag\_category\_description](#input\_tag\_category\_description) | The description of the vSphere tag category. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tag_category"></a> [tag\_category](#output\_tag\_category) | The vSphere tag category. |
| <a name="output_tags"></a> [tags](#output\_tags) | The list of vSphere tags. |
<!-- END_TF_DOCS -->