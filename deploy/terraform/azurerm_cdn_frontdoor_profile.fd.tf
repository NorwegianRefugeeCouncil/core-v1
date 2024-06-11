resource "azurerm_cdn_frontdoor_profile" "fd" {
  provider            = azurerm.runtime
  name                = "fd-${var.app_name}-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  sku_name            = var.frontdoor_sku_name

  response_timeout_seconds = 240
}
