resource "azurerm_monitor_diagnostic_setting" "backend" {
  provider                   = azurerm.runtime
  name                       = "diag-app-service-${azurerm_linux_web_app.app.name}"
  target_resource_id         = azurerm_linux_web_app.app.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.law.id

  enabled_log {
    category = "AppServiceAntivirusScanAuditLogs"
  }
  enabled_log {
    category = "AppServiceHTTPLogs"
  }
  enabled_log {
    category = "AppServiceConsoleLogs"
  }
  enabled_log {
    category = "AppServiceAppLogs"
  }
  enabled_log {
    category = "AppServiceFileAuditLogs"
  }
  enabled_log {
    category = "AppServiceAuditLogs"
  }
  enabled_log {
    category = "AppServiceIPSecAuditLogs"
  }
  enabled_log {
    category = "AppServicePlatformLogs"
  }
  metric {
    category = "AllMetrics"
  }
}
