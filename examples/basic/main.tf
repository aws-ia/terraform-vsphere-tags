module "vsphere_tags" {
  source = "../.."

  tag_category_name             = var.tag_category_name
  tag_category_description      = var.tag_category_description
  tag_category_associable_types = var.tag_category_associable_types
  tag_category_cardinality      = var.tag_category_cardinality
  create_tag_category           = var.create_tag_category
  tags                          = var.tags
  create_tags                   = var.create_tags
}
