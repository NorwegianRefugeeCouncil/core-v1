resource "azurerm_application_insights" "aisd" {
  name                = "example-appinsights"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  application_type    = "web"
}

resource "azurerm_monitor_action_group" "teams" {
  name                = "notify-teams"
  resource_group_name = azurerm_resource_group.rg.name
  short_name          = "notify-teams"

  logic_app_receiver {
    name                    = azurerm_logic_app_workflow.law.name
    use_common_alert_schema = true
    resource_id             = azurerm_logic_app_workflow.law.id
    callback_url            = azurerm_logic_app_workflow.law.access_endpoint
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "dep-latency" {
  name                = "Dependency Latency Degradation"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "P1D"
  detector_type       = "DependencyPerformanceDegradationDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "exceptions" {
  name                = "Exception Anomalies detected"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "P1D"
  detector_type       = "ExceptionVolumeChangedDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "failures" {
  name                = "Failure Anomalies"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "PT1M"
  detector_type       = "FailureAnomaliesDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "mem-leak" {
  name                = "Potential Memory Leak"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "P1D"
  detector_type       = "MemoryLeakDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "resp-latency" {
  name                = "Response Latency Degradation"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "P1D"
  detector_type       = "RequestPerformanceDegradationDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}

resource "azurerm_monitor_smart_detector_alert_rule" "trace-severity" {
  name                = "Trace Severity Degradation"
  resource_group_name = azurerm_resource_group.rg.name
  severity            = "Sev3"
  scope_resource_ids  = [azurerm_application_insights.aisd.id]
  frequency           = "P1D"
  detector_type       = "TraceSeverityDetector"

  action_group {
    ids = [azurerm_monitor_action_group.teams.id]
  }
}