resource "vsphere_tag_category" "category" {
  count = (var.create_tag_category && length(var.tag_category_name) > 0) ? 1 : 0

  name             = var.tag_category_name
  description      = var.tag_category_description
  cardinality      = var.tag_category_cardinality
  associable_types = var.tag_category_associable_types
}

resource "vsphere_tag" "tags" {
  count = (var.create_tags) ? length(var.tags) : 0

  name        = var.tags[count.index]["name"]
  description = var.tags[count.index]["description"]
  category_id = (var.create_tag_category && length(var.tag_category_name) > 0) ? vsphere_tag_category.category[0].id : data.vsphere_tag_category.category[0].id
}
