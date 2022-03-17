<!-- BEGIN_TF_DOCS -->
# Basic example

If deployed with the default values, this example will create a [tag category][category] named `example-category` and two [tags][tags]: 1/ `terraform` and 2/ `project` in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on-premises environment.

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

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_vsphere_tags"></a> [vsphere\_tags](#module\_vsphere\_tags) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_create_tag_category"></a> [create\_tag\_category](#input\_create\_tag\_category) | If true, a new vSphere tag category will be created. | `bool` | `true` | no |
| <a name="input_create_tags"></a> [create\_tags](#input\_create\_tags) | If true, new vSphere tags will be created for each entry. | `bool` | `true` | no |
| <a name="input_tag_category_associable_types"></a> [tag\_category\_associable\_types](#input\_tag\_category\_associable\_types) | A list object types that this category to which this category can be assigned (https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category#associable-object-types). | `list(string)` | <pre>[<br>  "Folder",<br>  "ClusterComputeResource",<br>  "Datacenter",<br>  "Datastore",<br>  "StoragePod",<br>  "DistributedVirtualPortgroup",<br>  "DistributedVirtualSwitch",<br>  "VmwareDistributedVirtualSwitch",<br>  "HostSystem",<br>  "com.vmware.content.Library",<br>  "com.vmware.content.library.Item",<br>  "HostNetwork",<br>  "Network",<br>  "OpaqueNetwork",<br>  "ResourcePool",<br>  "VirtualApp",<br>  "VirtualMachine"<br>]</pre> | no |
| <a name="input_tag_category_cardinality"></a> [tag\_category\_cardinality](#input\_tag\_category\_cardinality) | The number of tags that can be assigned from this category to a single object at once. | `string` | `"MULTIPLE"` | no |
| <a name="input_tag_category_description"></a> [tag\_category\_description](#input\_tag\_category\_description) | The description of the vSphere tag category. | `string` | `null` | no |
| <a name="input_tag_category_name"></a> [tag\_category\_name](#input\_tag\_category\_name) | The name of the vSphere tag category. | `string` | `"example-category"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty. | `list(map(string))` | <pre>[<br>  {<br>    "description": "Managed by Terraform",<br>    "name": "terraform"<br>  },<br>  {<br>    "description": "terraform-vsphere-tags",<br>    "name": "project"<br>  }<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tag_category"></a> [tag\_category](#output\_tag\_category) | The vSphere tag category. |
| <a name="output_tags"></a> [tags](#output\_tags) | The list of vSphere tags. |
<!-- END_TF_DOCS -->