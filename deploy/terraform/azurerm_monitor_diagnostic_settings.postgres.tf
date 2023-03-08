
resource "azurerm_monitor_diagnostic_setting" "postgres" {
  provider                   = azurerm.runtime
  name                       = "diag-postgres-${azurerm_postgresql_flexible_server.postgres.name}"
  target_resource_id         = azurerm_postgresql_flexible_server.postgres.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.law.id

  enabled_log {
    category = "PostgreSQLLogs"
  }
  metric {
    category = "AllMetrics"
  }
}
