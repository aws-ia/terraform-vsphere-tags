<!-- BEGIN_TF_DOCS -->
# vSphere Tags Terraform Module

Terraform module that either creates or imports a list of [tags][tags] grouped in a [tag category][category] in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on&#8209;premises environment for labeling your inventory objects with metadata to make it easier to sort and search for these objects.

Of note, tags and categories can span multiple [vCenter Server][vsphere] instances.
When you use [Hybrid Linked Mode][hybrid], tags and tag categories are maintained across your linked domain, which means the on&#8209;premises and VMware Cloud on AWS software&#8209;defined data centers (SDDCs) share tags and tag attributes.
If multiple on&#8209;premises vCenter Server instances are configured to use [Enhanced Linked Mode][enhanced], tags and categories are replicated across the vCenter Server instances.

## Usage

### Create new tags in a new tag category

```hcl
module "vsphere_tags" {
  source = "aws-ia/tags/vsphere"
  # The AWS Integration & Automation team recommends pinning every module to a specific version
  # version = "x.x.x"

  tag_category_name        = "example-category"
  tag_category_description = "Example tag category."
  tag_category_cardinality = "MULTIPLE"
  create_tag_category      = true
  create_tags              = true

  tags = [
    {
      name        = "terraform"
      description = "Managed by Terraform"
    },
    {
      name        = "project"
      description = "terraform-vsphere-tags"
    },
  ]
}
```

### Create new tags in an existing tag category

```hcl
module "vsphere_tags" {
  source = "aws-ia/tags/vsphere"
  # The AWS Integration & Automation team recommends pinning every module to a specific version
  # version = "x.x.x"

  tag_category_name        = "example-category"
  create_tag_category      = false
  create_tags              = true

  tags = [
    {
      name        = "terraform"
      description = "Managed by Terraform"
    },
    {
      name        = "project"
      description = "terraform-vsphere-tags"
    },
  ]
}
```

### Import existing tags in an existing tag category

```hcl
module "vsphere_tags" {
  source = "aws-ia/tags/vsphere"
  # The AWS Integration & Automation team recommends pinning every module to a specific version
  # version = "x.x.x"

  tag_category_name        = "example-category"
  create_tag_category      = false
  create_tags              = false

  tags = [
    {
      name        = "terraform"
      description = "Managed by Terraform"
    },
    {
      name        = "project"
      description = "terraform-vsphere-tags"
    },
  ]
}
```

[category]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
[enhanced]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-6ADB06EF-E342-457E-A17B-1EA31C0F6D4B.html
[hybrid]: https://docs.vmware.com/en/VMware-Cloud-on-AWS/services/com.vmware.vsphere.vmc-aws-manage-data-center-vms.doc/GUID-91C57891-4D61-4F4C-B580-74F3000B831D.html
[tags]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-2FF21224-B6BC-499B-AD8B-D2C4309AD9DC.html
[vsphere]: https://docs.vmware.com/en/VMware-vSphere/index.html
[vmconaws]: https://aws.amazon.com/vmware/

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
| <a name="input_tags"></a> [tags](#input\_tags) | List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty. | `list(map(string))` | n/a | yes |
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