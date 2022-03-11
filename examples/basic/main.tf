module "vsphere_tags" {
  source = "../.."

  vsphere_tag_category_name             = var.vsphere_tag_category_name
  vsphere_tag_category_description      = var.vsphere_tag_category_description
  vsphere_tag_category_associable_types = var.vsphere_tag_category_associable_types
  vsphere_tag_category_cardinality      = var.vsphere_tag_category_cardinality
  create_vsphere_tag_category           = var.create_vsphere_tag_category
  vsphere_tags                          = var.vsphere_tags
  create_vsphere_tags                   = var.create_vsphere_tags
}
