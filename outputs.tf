output "tag_category" {
  description = "The vSphere tag category."
  value       = var.create_vsphere_tag_category ? vsphere_tag_category.category : data.vsphere_tag_category.category
}

output "tags" {
  description = "The list of vSphere tags."
  value       = var.create_vsphere_tags ? vsphere_tag.tags : data.vsphere_tag.tags
}
