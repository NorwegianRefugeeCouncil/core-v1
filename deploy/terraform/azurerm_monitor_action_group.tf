resource "azurerm_monitor_action_group" "ag_teams" {
  name                = "send-notification-to-teams-${var.environment}"
  provider            = azurerm.runtime
  resource_group_name = azurerm_resource_group.rg.name
  short_name          = "notify-teams"

  webhook_receiver {
    name                    = "send_alert_to_okta_workflows"
    service_uri             = "https://nrc.workflows.oktapreview.com/api/flo/bb459c9dcfe01856889a5882639cdea2/invoke?clientToken=fdbde6450a07740538547a2749139ea1703bd5f84d44bcee5b9ca33885f9f530"
    use_common_alert_schema = true
  }
}
