variable "vsphere_tag_category_name" {
  type    = string
  default = "example-category"
}

variable "vsphere_tag_category_description" {
  type    = string
  default = null
}

variable "vsphere_tag_category_cardinality" {
  type    = string
  default = "MULTIPLE"
}

variable "vsphere_tag_category_associable_types" {
  type = list(string)
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

variable "create_vsphere_tag_category" {
  type    = bool
  default = true
}

variable "vsphere_tags" {
  type = list(map(string))
  default = [
    {
      name        = "Terraform"
      description = "Managed by Terraform"
    },
    {
      name        = "project"
      description = "terraform-vsphere-tags"
    },
  ]
}

variable "create_vsphere_tags" {
  type    = bool
  default = true
}
