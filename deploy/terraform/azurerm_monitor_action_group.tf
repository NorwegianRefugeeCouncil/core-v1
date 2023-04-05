resource "azurerm_monitor_action_group" "ag_teams" {
  name                = "send-notification-to-teams-${var.environment}"
  provider            = azurerm.runtime
  resource_group_name = azurerm_resource_group.rg.name
  short_name          = "notify-teams"

  logic_app_receiver {
    name                    = azurerm_logic_app_workflow.logic_app_teams.name
    use_common_alert_schema = true
    resource_id             = azurerm_logic_app_workflow.logic_app_teams.id
    callback_url            = azurerm_logic_app_workflow.logic_app_teams.access_endpoint
  }
}
