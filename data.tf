data "vsphere_tag_category" "category" {
  count = !var.create_vsphere_tag_category && length(var.vsphere_tag_category_name) > 0 ? 1 : 0

  name = var.vsphere_tag_category_name
}

data "vsphere_tag" "tags" {
  count = !var.create_vsphere_tag_category && !var.create_vsphere_tags ? length(var.vsphere_tags) : 0

  name        = var.vsphere_tags[count.index]["name"]
  category_id = var.create_vsphere_tag_category && length(var.vsphere_tag_category_name) > 0 ? vsphere_tag_category.category[0].id : data.vsphere_tag_category.category[0].id
}
