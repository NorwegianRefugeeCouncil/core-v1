resource "azurerm_cdn_frontdoor_route" "backend" {
  provider                      = azurerm.runtime
  name                          = "backend"
  cdn_frontdoor_endpoint_id     = azurerm_cdn_frontdoor_endpoint.backend.id
  cdn_frontdoor_origin_group_id = azurerm_cdn_frontdoor_origin_group.backend.id
  cdn_frontdoor_origin_ids      = [azurerm_cdn_frontdoor_origin.backend.id]
  // cdn_frontdoor_rule_set_ids    = [azurerm_cdn_frontdoor_rule_set.fd.id]
  enabled = true

  forwarding_protocol    = "HttpsOnly"
  https_redirect_enabled = true
  patterns_to_match      = ["/*"]
  supported_protocols    = ["Http", "Https"]

  cdn_frontdoor_custom_domain_ids = [
    azurerm_cdn_frontdoor_custom_domain.backend.id,
  ]
  link_to_default_domain = false
}
