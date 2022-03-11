output "vsphere_tag_category" {
  description = "The vSphere tag category."
  value       = module.vsphere_tags.vsphere_tag_category
}

output "vsphere_tags" {
  description = "The list of vSphere tags."
  value       = module.vsphere_tags.vsphere_tags
}
