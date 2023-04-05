resource "azurerm_logic_app_workflow" "logic_app_teams" {
  name                = "send_alerts_to_teams_thread-${var.environment}"
  provider            = azurerm.runtime
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_logic_app_trigger_http_request" "send_alerts_to_teams_thread" {
  name         = "send-core-alerts-to-teams-trigger"
  provider     = azurerm.runtime
  logic_app_id = azurerm_logic_app_workflow.logic_app_teams.id

  schema = <<SCHEMA
{
  "schemaId": "azureMonitorCommonAlertSchema",
  "data": {
    "essentials": {
      "alertId": "/subscriptions/<subscription ID>/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330",
      "alertRule": "WCUS-R2-Gen2",
      "severity": "Sev3",
      "signalType": "Metric",
      "monitorCondition": "Resolved",
      "monitoringService": "Platform",
      "alertTargetIDs": [
        "/subscriptions/<subscription ID>/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"
      ],
      "configurationItems": [
        "wcus-r2-gen2"
      ],
      "originAlertId": "3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227",
      "firedDateTime": "2019-03-22T13:58:24.3713213Z",
      "resolvedDateTime": "2019-03-22T14:03:16.2246313Z",
      "description": "",
      "essentialsVersion": "1.0",
      "alertContextVersion": "1.0"
    },
    "alertContext": {
      "properties": null,
      "conditionType": "SingleResourceMultipleMetricCriteria",
      "condition": {
        "windowSize": "PT5M",
        "allOf": [
          {
            "metricName": "Percentage CPU",
            "metricNamespace": "Microsoft.Compute/virtualMachines",
            "operator": "GreaterThan",
            "threshold": "25",
            "timeAggregation": "Average",
            "dimensions": [
              {
                "name": "ResourceId",
                "value": "3efad9dc-3d50-4eac-9c87-8b3fd6f97e4e"
              }
            ],
            "metricValue": 7.727
          }
        ]
      }
    }
  }
}
SCHEMA

}

locals {
  arm_file_path_logic_app = "azurerm_logic_app_workflow.logic_app_teams.json"
}

data "template_file" "logic_app_schema" {
  template = file(local.arm_file_path_logic_app)
}

resource "azurerm_resource_group_template_deployment" "logic_app_deployment" {
  provider                      = azurerm.runtime
  depends_on                    = [azurerm_logic_app_workflow.logic_app_teams]
  resource_group_name           = azurerm_resource_group.rg.name
  deployment_mode               = "Incremental"
  name                          = "logic_app_deployment-${var.environment}"
  template_content              = data.template_file.logic_app_schema.template
  parameters_content = jsonencode({
    "coreChannelParentMessage" = {
      value =  var.teams_core_alerts_channel_thread_parent_message_id
    },
    "ictChannelParentMessage" = {
      value =  var.teams_ict_service_alerts_channel_thread_parent_message_id
    },
    "coreChannelId" = {
      value =  var.teams_d_team_core_alerts_channel_id
    },
    "ictChannelId" = {
      value =  var.teams_global_ict_ict_service_alerts_channel_id
    },
    "coreTeamId" = {
      value =  var.teams_d_team_team_id
    },
    "ictTeamId" = {
      value =  var.teams_global_ict_team_id
    },
    "workflows_send_alerts_to_teams_thread_name" = {
      value = "send-alerts-to-teams-thread-${var.environment}"
    },
  })
}
