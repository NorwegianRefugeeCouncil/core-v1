resource "azurerm_role_assignment" "storage_data_contributor" {
  provider             = azurerm.runtime
  principal_id         = azurerm_user_assigned_identity.app.principal_id
  role_definition_name = "Storage Blob Data Contributor"
  scope                = azurerm_storage_account.download_storage.id
}