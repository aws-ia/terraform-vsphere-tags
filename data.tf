data "vsphere_tag_category" "category" {
  count = (!var.create_tag_category && length(var.tag_category_name) > 0) ? 1 : 0

  name = var.tag_category_name
}

data "vsphere_tag" "tags" {
  count = (length(data.vsphere_tag_category.category) == 1 && !var.create_tags) ? length(var.tags) : 0

  name        = var.tags[count.index]["name"]
  category_id = data.vsphere_tag_category.category[0].id
}
