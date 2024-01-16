resource "azurerm_private_dns_zone" "storage_dns" {
  provider            = azurerm.runtime
  name                = "privatelink.blob.core.windows.net"
  resource_group_name = azurerm_resource_group.rg.name
}