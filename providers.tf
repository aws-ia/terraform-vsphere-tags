terraform {
  required_version = ">= 1.1.0"
  required_providers {
    vsphere = {
      source  = "hashicorp/vsphere"
      version = ">= 2.1.0"
    }
  }
}
