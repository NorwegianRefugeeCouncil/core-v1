resource "azurerm_cdn_frontdoor_custom_domain_association" "backend" {
  provider                       = azurerm.runtime
  cdn_frontdoor_custom_domain_id = azurerm_cdn_frontdoor_custom_domain.backend.id
  cdn_frontdoor_route_ids        = [azurerm_cdn_frontdoor_route.backend.id]
}
