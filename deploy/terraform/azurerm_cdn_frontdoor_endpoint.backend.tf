resource "azurerm_cdn_frontdoor_endpoint" "backend" {
  provider                 = azurerm.runtime
  name                     = "backend"
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
}
