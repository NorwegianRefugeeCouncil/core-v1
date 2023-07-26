data "azurerm_storage_account" "download_storage" {
  name                = "${var.download_storage_name}"
  resource_group_name = "${var.resource_group_name}"
  location            = "${var.location}"
  account_tier        = "Standard"
}

resource "azurerm_storage_container" "download_storage_container" {
  name                  = "${var.download_storage_container_name}"
  storage_account_name  = "${data.azurerm_storage_account.download_storage.name}"
  container_access_type = "private"
}

resource "azurerm_storage_management_policy" "delete_download_files" {
  storage_account_id = azurerm_storage_account.download_storage.id

  rule {
    name    = "delete_download_files"
    enabled = true
    filters {
      prefix_match = ["${var.download_storage_container_name}/"]
      blob_types   = ["blockBlob"]
      match_blob_index_tag {
        name      = "tag1"
        operation = "=="
        value     = "val1"
      }
    }
    actions {
      base_blob {
        delete_after_days_since_creation_greater_than = 1
      }
      snapshot {
        delete_after_days_since_creation_greater_than = 1
      }
    }
  }
}

