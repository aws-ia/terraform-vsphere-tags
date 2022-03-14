<!-- BEGIN_TF_DOCS -->
# vSphere tags Terraform module

The vSphere tags Terraform module either creates or imports a [tag category][category] as well as a list of [tags][tags] in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on-premises environment, and then outputs the configuration of each for use in other Terraform projects.

[category]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
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
| <a name="input_vsphere_tag_category_name"></a> [vsphere\_tag\_category\_name](#input\_vsphere\_tag\_category\_name) | The name of the vSphere tag category. | `string` | n/a | yes |
| <a name="input_vsphere_tags"></a> [vsphere\_tags](#input\_vsphere\_tags) | List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty. | `list(map(string))` | n/a | yes |
| <a name="input_create_vsphere_tag_category"></a> [create\_vsphere\_tag\_category](#input\_create\_vsphere\_tag\_category) | If true, a new vSphere tag category will be created. | `bool` | `false` | no |
| <a name="input_create_vsphere_tags"></a> [create\_vsphere\_tags](#input\_create\_vsphere\_tags) | If true, new vSphere tags will be created for each entry. | `bool` | `false` | no |
| <a name="input_vsphere_tag_category_associable_types"></a> [vsphere\_tag\_category\_associable\_types](#input\_vsphere\_tag\_category\_associable\_types) | A list object types that this category to which this category can be assigned (https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category#associable-object-types). | `list(string)` | <pre>[<br>  "Folder",<br>  "ClusterComputeResource",<br>  "Datacenter",<br>  "Datastore",<br>  "StoragePod",<br>  "DistributedVirtualPortgroup",<br>  "DistributedVirtualSwitch",<br>  "VmwareDistributedVirtualSwitch",<br>  "HostSystem",<br>  "com.vmware.content.Library",<br>  "com.vmware.content.library.Item",<br>  "HostNetwork",<br>  "Network",<br>  "OpaqueNetwork",<br>  "ResourcePool",<br>  "VirtualApp",<br>  "VirtualMachine"<br>]</pre> | no |
| <a name="input_vsphere_tag_category_cardinality"></a> [vsphere\_tag\_category\_cardinality](#input\_vsphere\_tag\_category\_cardinality) | The number of tags that can be assigned from this category to a single object at once. | `string` | `"MULTIPLE"` | no |
| <a name="input_vsphere_tag_category_description"></a> [vsphere\_tag\_category\_description](#input\_vsphere\_tag\_category\_description) | The description of the vSphere tag category. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tag_category"></a> [tag\_category](#output\_tag\_category) | The vSphere tag category. |
| <a name="output_tags"></a> [tags](#output\_tags) | The list of vSphere tags. |
<!-- END_TF_DOCS -->