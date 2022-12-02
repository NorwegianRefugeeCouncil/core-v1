resource "azurerm_cdn_frontdoor_rule" "backend_disable_auth_cache" {
  provider = azurerm.runtime
  # Required as per terraform documentation
  depends_on = [
    azurerm_cdn_frontdoor_origin_group.backend,
    azurerm_cdn_frontdoor_origin.backend,
  ]

  name                      = "disableAuthCache"
  cdn_frontdoor_rule_set_id = azurerm_cdn_frontdoor_rule_set.backend.id
  order                     = 1
  behavior_on_match         = "Continue"

  actions {
    route_configuration_override_action {
      cache_behavior                = "Disabled"
    }
  }

  conditions {
    request_uri_condition {
      operator = "BeginsWith"
      match_values = ["/.auth"]
    }
  }
}