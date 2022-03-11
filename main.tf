resource "vsphere_tag_category" "category" {
  count = var.create_vsphere_tag_category && length(var.vsphere_tag_category_name) > 0 ? 1 : 0

  name             = var.vsphere_tag_category_name
  description      = var.vsphere_tag_category_description
  cardinality      = var.vsphere_tag_category_cardinality
  associable_types = var.vsphere_tag_category_associable_types
}

resource "vsphere_tag" "tags" {
  count = var.create_vsphere_tags ? length(var.vsphere_tags) : 0

  name        = var.vsphere_tags[count.index]["name"]
  description = var.vsphere_tags[count.index]["description"]
  category_id = var.create_vsphere_tag_category && length(var.vsphere_tag_category_name) > 0 ? vsphere_tag_category.category[0].id : data.vsphere_tag_category.category[0].id
}
