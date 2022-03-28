output "tag_category" {
  description = "The vSphere tag category."
  value       = module.vsphere_tags.tag_category
}

output "tags" {
  description = "The vSphere tags."
  value       = module.vsphere_tags.tags
}
