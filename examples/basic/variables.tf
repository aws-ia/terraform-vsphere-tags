variable "tag_category_name" {
  type        = string
  description = "The name of the vSphere tag category."
  default     = "example-category"
}

variable "tag_category_description" {
  type        = string
  description = "The description of the vSphere tag category."
  default     = null
}

variable "tag_category_cardinality" {
  type        = string
  description = "The number of tags that can be assigned from this category to a single object at once."
  default     = "MULTIPLE"
}

variable "tag_category_associable_types" {
  type        = list(string)
  description = "A list object types that this category to which this category can be assigned (https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs/resources/tag_category#associable-object-types)."
  default = [
    "Folder",
    "ClusterComputeResource",
    "Datacenter",
    "Datastore",
    "StoragePod",
    "DistributedVirtualPortgroup",
    "DistributedVirtualSwitch",
    "VmwareDistributedVirtualSwitch",
    "HostSystem",
    "com.vmware.content.Library",
    "com.vmware.content.library.Item",
    "HostNetwork",
    "Network",
    "OpaqueNetwork",
    "ResourcePool",
    "VirtualApp",
    "VirtualMachine",
  ]
}

variable "create_tag_category" {
  type        = bool
  description = "If true, a new vSphere tag category will be created."
  default     = true
}

variable "tags" {
  type        = list(map(string))
  description = "List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty."
  default = [
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

variable "create_tags" {
  type        = bool
  description = "If true, new vSphere tags will be created for each entry."
  default     = true
}
