output "vsphere_tag_category" {
  value = {
    name             = var.create_vsphere_tag_category ? vsphere_tag_category.category[0].name : data.vsphere_tag_category.category[0].name
    description      = var.create_vsphere_tag_category ? vsphere_tag_category.category[0].description : data.vsphere_tag_category.category[0].description
    cardinality      = var.create_vsphere_tag_category ? vsphere_tag_category.category[0].cardinality : data.vsphere_tag_category.category[0].cardinality
    associable_types = var.create_vsphere_tag_category ? vsphere_tag_category.category[0].associable_types : data.vsphere_tag_category.category[0].associable_types
    id               = var.create_vsphere_tag_category ? vsphere_tag_category.category[0].id : data.vsphere_tag_category.category[0].id
  }
}

output "vsphere_tags" {
  value = var.create_vsphere_tags ? vsphere_tag.tags : data.vsphere_tag.tags
}
