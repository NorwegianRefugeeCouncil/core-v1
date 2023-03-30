# metric_namespaces https://learn.microsoft.com/en-us/azure/azure-monitor/alerts/alerts-metric-near-real-time#metrics-and-dimensions-supported
# severity levels https://learn.microsoft.com/en-us/azure/azure-monitor/best-practices-alerts#alert-severity
# frequency s https://tc39.es/proposal-temporal/docs/duration.html#:~:text=Briefly%2C%20the%20ISO%208601%20notation,suffix%20that%20indicates%20the%20unit

##################
# monitor database
# metric_names https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported#microsoftdbforpostgresqlflexibleservers
##################
resource "azurerm_monitor_metric_alert" "postgresCpuOverThreshold" {
  provider            = azurerm.runtime
  name                = "postgres-cpu-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_postgresql_flexible_server.postgres.id]
  description         = "Action will be triggered when the CPU percentage average is greater than 80."
  frequency           = "PT1M"
  severity            = 3

  criteria {
    metric_namespace = "Microsoft.DBforPostgreSQL/servers"
    metric_name      = "cpu_percent"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 80
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "postgresMemoryUsage" {
  provider            = azurerm.runtime
  name                = "postgres-memory-usage-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_postgresql_flexible_server.postgres.id]
  description         = "Action will be triggered when the memory usage average is greater than 70%."
  severity            = 3
  window_size         = "PT1H"
  frequency           = "PT30M"


  criteria {
    metric_namespace = "Microsoft.DBforPostgreSQL/servers"
    metric_name      = "memory_percent"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 70
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

#################
# monitor web app
# metric_names https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported#microsoftwebsites
#################
resource "azurerm_monitor_metric_alert" "appHealthCheck" {
  provider            = azurerm.runtime
  name                = "health-check-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when the HealthCheckStatus is less than 100% okay."
  severity            = 1
  frequency           = "PT1M"
  enabled             = false

  criteria {
    threshold         = 1
    operator          = "LessThan"
    aggregation       = "Average"
    metric_name       = "Microsoft.Web/sites"
    metric_namespace  = "HealthCheckStatus"
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "appCpuOverThreshold" {
  provider            = azurerm.runtime
  name                = "app-cpu-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when the CPU percentage average is greater than 80."
  severity            = 2
  frequency           = "PT1M"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "CpuTime"
    aggregation      = "Maximum"
    operator         = "GreaterThan"
    threshold        = 0.4
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "appMemoryOverThreshold" {
  provider            = azurerm.runtime
  name                = "app-memory-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when the Memory working set average is greater than 300MB."
  severity            = 2
  frequency           = "PT1M"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "MemoryWorkingSet"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 300000000
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "appResponseTime" {
  provider            = azurerm.runtime
  name                = "app-response-time-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when the Response time average is greater than 5s."
  severity            = 3
  frequency           = "PT5M"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "HttpResponseTime"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 5
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "app4xxStatusCodes" {
  provider            = azurerm.runtime
  name                = "app-4xx-status-codes-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when more than 10 4xx Errors are returned per hour."
  severity            = 3
  frequency           = "PT1H"
  window_size         = "PT1H"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "Http4xx"
    aggregation      = "Count"
    operator         = "GreaterThan"
    threshold        = 10
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}

resource "azurerm_monitor_metric_alert" "app5xxStatusCodes" {
  provider            = azurerm.runtime
  name                = "app-5xx-status-codes${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "Action will be triggered when more than 10 5xx Errors are returned per hour."
  severity            = 1
  frequency           = "PT1H"
  window_size         = "PT1H"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "Http5xx"
    aggregation      = "Count"
    operator         = "GreaterThan"
    threshold        = 10
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag-teams.id
  }
}
