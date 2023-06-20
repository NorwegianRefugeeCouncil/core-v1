resource "azurerm_monitor_activity_log_alert" "postgres_health" {
  provider            = azurerm.runtime
  name                = "postgres-health-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_postgresql_flexible_server.postgres.id]
  description         = "${var.environment} - Postgres server: health check"

  criteria {
    resource_id    = azurerm_postgresql_flexible_server.postgres.id
    operation_name = "Microsoft.Resourcehealth/healthevent/Activated/action"
    category       = "ResourceHealth"
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}
