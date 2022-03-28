data "vsphere_tag_category" "category" {
  count = (var.create_tag_category) ? 0 : 1

  name = var.tag_category_name
}

data "vsphere_tag" "tags" {
  for_each = (length(data.vsphere_tag_category.category) == 1 && !var.create_tags) ? var.tags : {}

  name        = each.key
  category_id = data.vsphere_tag_category.category[0].id
}
