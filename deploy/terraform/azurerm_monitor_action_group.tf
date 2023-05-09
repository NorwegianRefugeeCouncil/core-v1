resource "azurerm_monitor_action_group" "ag_teams" {
  name                = "send-notification-to-teams-${var.environment}"
  provider            = azurerm.runtime
  resource_group_name = azurerm_resource_group.rg.name
  short_name          = "notify-teams"

  webhook_receiver {
    name                    = "send_alert_to_okta_workflows"
    service_uri             = var.action_group_webhook_url
    use_common_alert_schema = true
  }
}
