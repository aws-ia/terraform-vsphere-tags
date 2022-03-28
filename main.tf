resource "vsphere_tag_category" "category" {
  count = (var.create_tag_category) ? 1 : 0

  name             = var.tag_category_name
  description      = var.tag_category_description
  cardinality      = var.tag_category_cardinality
  associable_types = var.tag_category_associable_types
}

resource "vsphere_tag" "tags" {
  for_each = (var.create_tags) ? var.tags : {}

  name        = each.key
  description = each.value
  category_id = (var.create_tag_category) ? vsphere_tag_category.category[0].id : data.vsphere_tag_category.category[0].id
}
