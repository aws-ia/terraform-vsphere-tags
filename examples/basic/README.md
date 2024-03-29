<!-- BEGIN_TF_DOCS -->
# Basic example

If deployed with the default values, this example will create two [tags](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-16422FF7-235B-4A44-92E2-532F6AED0923.html): 1/&nbsp;`terraform` and 2/&nbsp;`project` grouped in a [tag category](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-16422FF7-235B-4A44-92E2-532F6AED0923.html) named `example-category` in your [VMware Cloud on AWS](https://aws.amazon.com/vmware/) or [VMware vSphere](https://docs.vmware.com/en/VMware-vSphere/index.html) on&#8209;premises environment.

## Usage

### Configure the provider

One way to configure the VMware vSphere provider is with [environment variables](https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs#argument-reference), for example:

* Required
  * `VSPHERE_USER`: Username for vSphere API operations.
  * `VSPHERE_PASSWORD`: Password for vSphere API operations.
  * `VSPHERE_SERVER`: vCenter Server fully qualified domain name (FQDN) or IP address for vSphere API operations.
* Optional
  * `VSPHERE_ALLOW_UNVERIFIED_SSL`: Boolean that can be set to `true` to disable TLS certificate verification.
    This should be used with care as it could allow an attacker to intercept your auth token.
    If omitted, the default value is `false`.

### Configure the input variables

All of the variables in this example have default values, but if you would like to override any of these, one way to do so is to create a [`terraform.tfvars` variable definition file](https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files) in this directory.

#### Example `terraform.tfvars`

```hcl
tags = {
  override-example = "This example overrides the default value for the tags input variable defined in the variables.tf file in this directory and will create a tag named 'override-example'."
}
```

### Deploy

To deploy this example, execute the following: 1/&nbsp;[`terraform init`](https://www.terraform.io/cli/commands/init), 2/&nbsp;[`terraform plan`](https://www.terraform.io/cli/commands/plan), and 3/&nbsp;[`terraform apply`](https://www.terraform.io/cli/commands/apply).

### Clean-up

When you want to remove the resources, execute the following: [`terraform destroy`](https://www.terraform.io/cli/commands/destroy).

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
| <a name="module_vsphere_tags"></a> [vsphere\_tags](#module\_vsphere\_tags) | aws-ia/tags/vsphere | >= 0.0.1 |

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
| <a name="input_tags"></a> [tags](#input\_tags) | Map of strings defining vSphere tag names and descriptions. | `map(string)` | <pre>{<br>  "project": "terraform-vsphere-tags",<br>  "terraform": "Managed by Terraform"<br>}</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tag_category"></a> [tag\_category](#output\_tag\_category) | The vSphere tag category. |
| <a name="output_tags"></a> [tags](#output\_tags) | The list of vSphere tags. |
<!-- END_TF_DOCS -->