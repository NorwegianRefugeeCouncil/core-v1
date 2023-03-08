data "azurerm_managed_api" "ama" {
  name     = "servicebus"
  location = azurerm_resource_group.rg.location
}

resource "azurerm_servicebus_namespace" "asn" {
  name                = "acctestsbn-conn"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  sku                 = "Basic"
}

resource "azurerm_api_connection" "teams" {
  name                = "teams-connection"
  resource_group_name = azurerm_resource_group.rg.name
  managed_api_id      = data.azurerm_managed_api.ama.id
  display_name        = "Microsoft Service Alerts"
  parameter_values = {
    connectionString = azurerm_servicebus_namespace.asn.default_primary_connection_string
  }

  lifecycle {
    # NOTE: since the connectionString is a secure value it's not returned from the API
    ignore_changes = [parameter_values]
  }
}