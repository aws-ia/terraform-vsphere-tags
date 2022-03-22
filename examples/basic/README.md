<!-- BEGIN_TF_DOCS -->
# Basic example

If deployed with the default values, this example will create two [tags][tags]: 1/&nbsp;`terraform` and 2/&nbsp;`project` grouped in a [tag category][category] named `example-category` in your [VMware Cloud on AWS][vmconaws] or [VMware vSphere][vsphere] on&#8209;premises environment.

## Usage

### Configure the provider

One way to configure the VMware vSphere provider is with [environment variables][env\_vars], for example:

* Required
  * `VSPHERE_USER`: Username for vSphere API operations.
  * `VSPHERE_PASSWORD`: Password for vSphere API operations.
  * `VSPHERE_SERVER`: vCenter Server fully qualified domain name (FQDN) or IP address for vSphere API operations.
* Optional
  * `VSPHERE_ALLOW_UNVERIFIED_SSL`: Boolean that can be set to `true` to disable TLS certificate verification.
    This should be used with care as it could allow an attacker to intercept your auth token.
    If omitted, the default value is `false`.

### Configure the input variables

All of the variables in this example have default values, but if you would like to override any of these, one way to do so is to create a [`terraform.tfvars` variable definition file][tfvars] in this directory.

#### Example `terraform.tfvars`

```hcl
tags = [{
  "name"        = "override-example"
  "description" = "This example overrides the default value for the tags input variable defined in the variables.tf file in this directory and will create a tag named 'override-example'."
}]
```

### Deploy

To deploy this example, execute the following: 1/&nbsp;[`terraform init`][tf\_init], 2/&nbsp;[`terraform plan`][tf\_plan], and 3/&nbsp;[`terraform apply`][tf\_apply].

### Clean-up

When you want to remove the resources, execute the following: [`terraform destroy`][tf\_destroy].

[category]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
[env\_vars]: https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs#argument-reference
[tags]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vcenterhost.doc/GUID-2FF21224-B6BC-499B-AD8B-D2C4309AD9DC.html
[tf\_apply]: https://www.terraform.io/cli/commands/apply
[tf\_destroy]: https://www.terraform.io/cli/commands/destroy
[tf\_init]: https://www.terraform.io/cli/commands/init
[tf\_plan]: https://www.terraform.io/cli/commands/plan
[tfvars]: https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files
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
| <a name="module_vsphere_tags"></a> [vsphere\_tags](#module\_vsphere\_tags) | aws-ia/vsphere-tags/vsphere | >= 0.0.1 |

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