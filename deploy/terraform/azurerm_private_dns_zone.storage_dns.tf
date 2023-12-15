resource "azurerm_private_dns_zone" "storage_dns" {
  provider            = azurerm.runtime
  name                = "${var.app_name}-${var.environment}.storage.azure.com"
  resource_group_name = azurerm_resource_group.rg.name
}