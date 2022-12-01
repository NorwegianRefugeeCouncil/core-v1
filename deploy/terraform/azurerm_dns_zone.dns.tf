resource "azurerm_dns_zone" "dns" {
  provider            = azurerm.runtime
  name                = var.dns_zone_name
  resource_group_name = azurerm_resource_group.rg.name
}
