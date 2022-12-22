resource "azurerm_cdn_frontdoor_rule" "backend_disable_download_compression" {
  provider = azurerm.runtime
  # Required as per terraform documentation
  depends_on = [
    azurerm_cdn_frontdoor_origin_group.backend,
    azurerm_cdn_frontdoor_origin.backend,
  ]

  name                      = "disableDownloadCompression"
  cdn_frontdoor_rule_set_id = azurerm_cdn_frontdoor_rule_set.backend.id
  order                     = 2
  behavior_on_match         = "Continue"

  actions {
    route_configuration_override_action {
      compression_enabled = false
    }
  }

  conditions {
    request_uri_condition {
      operator     = "EndsWith"
      match_values = ["/download"]
    }
  }
}