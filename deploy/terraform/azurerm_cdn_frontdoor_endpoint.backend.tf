resource "azurerm_cdn_frontdoor_endpoint" "backend" {
  provider                 = azurerm.runtime
  name                     = "backend-${random_id.app_id.hex}"
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
}
