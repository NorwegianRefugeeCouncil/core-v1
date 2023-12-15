resource "azurerm_private_dns_zone_virtual_network_link" "storage_vnet_link" {
  provider              = azurerm.runtime
  name                  = azurerm_private_dns_zone.storage_dns.name
  private_dns_zone_name = azurerm_private_dns_zone.storage_dns.name
  virtual_network_id    = azurerm_virtual_network.vnet.id
  resource_group_name   = azurerm_resource_group.rg.name
}