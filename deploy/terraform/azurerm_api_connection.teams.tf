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
  name                = "teams"
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

locals {
  arm_file_path_api_connection = "azurerm_api_connection.teams.json"
}

data "template_file" "teams-connection-schema" {
  template = file(local.arm_file_path_api_connection)
}

resource "azurerm_resource_group_template_deployment" "teams-connection-deployment" {
  depends_on = [azurerm_api_connection.teams]
  resource_group_name = azurerm_resource_group.rg.name
  deployment_mode = "Incremental"
  name = "teams-connection-deployment"
  template_content = data.template_file.teams-connection-schema.rendered
}
