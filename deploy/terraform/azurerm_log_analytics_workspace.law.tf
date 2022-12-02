resource "azurerm_log_analytics_workspace" "law" {
  provider            = azurerm.runtime
  name                = "law-${var.app_name}-${var.environment}"
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
  retention_in_days   = 30
}
