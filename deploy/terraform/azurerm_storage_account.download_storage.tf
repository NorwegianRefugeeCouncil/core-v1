data "azurerm_storage_account" "download_storage" {
  name                      = "${var.download_storage_account_name}"
  resource_group_name       = azurerm_resource_group.rg.name
  location                  = "${var.location}"
  account_tier              = "Standard"
  account_replication_type  = "LRS"

  network_rules {
    default_action             = "Deny"
    virtual_network_subnet_ids = [azurerm_subnet.storage_subnet.id]
  }
}

resource "azurerm_storage_container" "download_storage_container" {
  name                  = "${var.download_storage_container_name}"
  storage_account_name  = data.azurerm_storage_account.download_storage.name
  container_access_type = "private"
}

resource "azurerm_storage_management_policy" "delete_download_files" {
  storage_account_id = data.azurerm_storage_account.download_storage.id

  rule {
    name    = "delete_download_files"
    enabled = true
    filters {
      prefix_match = ["${var.download_storage_container_name}/"]
      blob_types   = ["blockBlob"]
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

