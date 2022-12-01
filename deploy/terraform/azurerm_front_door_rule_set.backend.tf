resource "azurerm_cdn_frontdoor_rule_set" "backend" {
  provider                 = azurerm.runtime
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
  name                     = "rules"
}