data "azurerm_dns_zone" "dns" {
  provider            = azurerm.infra
  name                = var.infra_dns_zone_name
  resource_group_name = var.infra_resource_group_name
}