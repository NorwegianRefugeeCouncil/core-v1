resource "azurerm_subnet" "storage_subnet" {
  provider             = azurerm.runtime
  name                 = "storage"
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = [var.storage_subnet_address_space]
  service_endpoints = ["Microsoft.Storage"]
}