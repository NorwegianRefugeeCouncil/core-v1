resource "azurerm_monitor_diagnostic_setting" "frontdoor" {
  provider                   = azurerm.runtime
  name                       = "diag-frontdoor-${azurerm_cdn_frontdoor_profile.fd.name}"
  target_resource_id         = azurerm_cdn_frontdoor_profile.fd.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.law.id

  enabled_log {
    category = "FrontDoorAccessLog"
  }
  enabled_log {
    category = "FrontDoorHealthProbeLog"
  }
  enabled_log {
    category = "FrontDoorWebApplicationFirewallLog"
  }
  metric {
    category = "AllMetrics"
  }
}
