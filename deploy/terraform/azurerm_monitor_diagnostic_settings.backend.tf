resource "azurerm_monitor_diagnostic_setting" "backend" {
  provider                   = azurerm.runtime
  name                       = "diag-app-service-${azurerm_linux_web_app.app.name}"
  target_resource_id         = azurerm_linux_web_app.app.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.law.id

  log {
    category = "AppServiceAntivirusScanAuditLogs"
  }
  log {
    category = "AppServiceHTTPLogs"
  }
  log {
    category = "AppServiceConsoleLogs"
  }
  log {
    category = "AppServiceAppLogs"
  }
  log {
    category = "AppServiceFileAuditLogs"
  }
  log {
    category = "AppServiceAuditLogs"
  }
  log {
    category = "AppServiceIPSecAuditLogs"
  }
  log {
    category = "AppServicePlatformLogs"
  }
  metric {
    category = "AllMetrics"
  }
}
