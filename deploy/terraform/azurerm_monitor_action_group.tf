resource "azurerm_monitor_action_group" "ag-teams" {
  name                = "send-notification-to-teams-${var.environment}"
  provider            = azurerm.runtime
  resource_group_name = azurerm_resource_group.rg.name
  short_name          = "notify-teams"

  logic_app_receiver {
    name                    = azurerm_logic_app_workflow.logic-app-teams.name
    use_common_alert_schema = true
    resource_id             = azurerm_logic_app_workflow.logic-app-teams.id
    callback_url            = azurerm_logic_app_workflow.logic-app-teams.access_endpoint
  }
}
