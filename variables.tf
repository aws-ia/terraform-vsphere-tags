variable "tag_category_name" {
  type        = string
  description = "The name of the vSphere tag category."
  nullable    = false

  validation {
    condition     = length(var.tag_category_name) > 0
    error_message = "Must be a string of one or more characters."
  }
}

variable "tag_category_description" {
  type        = string
  description = "The description of the vSphere tag category."
  default     = null
  nullable    = true
}

variable "tag_category_cardinality" {
  type        = string
  description = "The number of tags that can be assigned from this category to a single object at once."
  default     = "MULTIPLE"
  nullable    = false

  validation {
    condition     = contains(["SINGLE", "MULTIPLE"], var.tag_category_cardinality)
    error_message = "Accepted values: 'SINGLE', 'MULTIPLE'."
  }
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
  nullable = false

  validation {
    condition     = length(var.tag_category_associable_types) > 0
    error_message = "Must be a list of one or more strings."
  }

  validation {
    condition = alltrue([
      for t in var.tag_category_associable_types : contains([
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
      ], t)
    ])
    error_message = "Accepted values: 'Folder', 'ClusterComputeResource', 'Datacenter', 'Datastore', 'StoragePod', 'DistributedVirtualPortgroup', 'DistributedVirtualSwitch', 'VmwareDistributedVirtualSwitch', 'HostSystem', 'com.vmware.content.Library', 'com.vmware.content.library.Item', 'HostNetwork', 'Network', 'OpaqueNetwork', 'ResourcePool', 'VirtualApp', 'VirtualMachine'."
  }
}

variable "create_tag_category" {
  type        = bool
  description = "If true, a new vSphere tag category will be created."
  default     = false
  nullable    = false
}

variable "tags" {
  type        = list(map(string))
  description = "List of one or more maps of strings defining vSphere tags. Each map must only have 'name' & 'description' keys, and the value for 'name' cannot be empty."
  nullable    = false

  /*
  example = [
    {
      name        = "terraform"
      description = "Managed by Terraform"
    },
    {
      name        = "project"
      description = "terraform-vsphere-tags"
    },
  ]
  */

  validation {
    condition     = length(var.tags) > 0
    error_message = "Must be a list of one or more maps of strings."
  }

  validation {
    condition     = alltrue([for t in var.tags : length(keys(t)) == 2])
    error_message = "Must be a list of one or more maps of strings with exactly 2 keys."
  }

  validation {
    condition     = alltrue([for t in var.tags : alltrue([for k, v in t : (k == "name" && length(v) > 0) || k == "description"])])
    error_message = "Must be a list of one or more maps of strings with only 'name' & 'description' keys, and the value for 'name' cannot be empty."
  }
}

variable "create_tags" {
  type        = bool
  description = "If true, new vSphere tags will be created for each entry."
  default     = false
  nullable    = false
}
