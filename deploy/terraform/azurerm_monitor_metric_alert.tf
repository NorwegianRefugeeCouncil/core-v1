# metric_namespaces https://learn.microsoft.com/en-us/azure/azure-monitor/alerts/alerts-metric-near-real-time#metrics-and-dimensions-supported
# severity levels https://learn.microsoft.com/en-us/azure/azure-monitor/best-practices-alerts#alert-severity
# frequency s https://tc39.es/proposal-temporal/docs/duration.html#:~:text=Briefly%2C%20the%20ISO%208601%20notation,suffix%20that%20indicates%20the%20unit

##################
# monitor database
# metric_names https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported#microsoftdbforpostgresqlflexibleservers
##################
resource "azurerm_monitor_metric_alert" "postgres_cpu_over_threshold" {
  provider            = azurerm.runtime
  name                = "postgres-cpu-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_postgresql_flexible_server.postgres.id]
  description         = "${var.environment} - Postgres server: CPU percentage average is greater than 95 in the last 5 minutes."
  frequency           = "PT5M"
  severity            = 3
  window_size         = "PT5M"

  criteria {
    metric_namespace = "Microsoft.DBforPostgreSQL/flexibleServers"
    metric_name      = "cpu_percent"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 95
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

resource "azurerm_monitor_metric_alert" "postgres_memory_usage" {
  provider            = azurerm.runtime
  name                = "postgres-memory-usage-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_postgresql_flexible_server.postgres.id]
  description         = "${var.environment} - Postgres server: Memory usage average is greater than 70% within the last hour."
  severity            = 3
  window_size         = "PT1H"
  frequency           = "PT30M"

  criteria {
    metric_namespace = "Microsoft.DBforPostgreSQL/flexibleServers"
    metric_name      = "memory_percent"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 70
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

##########################
# monitor app service plan
##########################


resource "azurerm_monitor_metric_alert" "asp_cpu_over_threshold" {
  provider            = azurerm.runtime
  name                = "asp-cpu-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_service_plan.plan.id]
  description         = "${var.environment} - App Service Plan: CPU percentage average is greater than 70%."
  severity            = 2
  frequency           = "PT1H"
  window_size         = "PT1H"

  criteria {
    metric_namespace = "Microsoft.Web/serverfarms"
    metric_name      = "CpuPercentage"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 70
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

resource "azurerm_monitor_metric_alert" "asp_memory_over_threshold" {
  provider            = azurerm.runtime
  name                = "asp-memory-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_service_plan.plan.id]
  description         = "${var.environment} - App Service Plan: Memory percentage average is greater than 70%."
  severity            = 2
  frequency           = "PT1H"
  window_size         = "PT1H"

  criteria {
    metric_namespace = "Microsoft.Web/serverfarms"
    metric_name      = "MemoryPercentage"
    aggregation      = "Average"
    operator         = "GreaterThan"
    threshold        = 70
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

#################
# monitor web app
# metric_names https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/metrics-supported#microsoftwebsites
#################
resource "azurerm_monitor_metric_alert" "app_health_check" {
  provider            = azurerm.runtime
  name                = "health-check-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "${var.environment} - Web App: HealthCheckStatus is less than 100% okay."
  severity            = 1
  frequency           = "PT1M"
  enabled             = true

  criteria {
    threshold         = 1
    operator          = "LessThan"
    aggregation       = "Average"
    metric_namespace  = "Microsoft.Web/sites"
    metric_name       = "HealthCheckStatus"
    skip_metric_validation = true
  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

resource "azurerm_monitor_metric_alert" "app_response_time" {
  provider            = azurerm.runtime
  name                = "app-response-time-over-threshold-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "${var.environment} - Web App: Response time average is greater than 5s."
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
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

resource "azurerm_monitor_metric_alert" "app_4xx_status_codes" {
  provider            = azurerm.runtime
  name                = "app-4xx-status-codes-${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "${var.environment} - Web App: more than 30 4xx Errors are returned per hour."
  severity            = 3
  frequency           = "PT1H"
  window_size         = "PT1H"

  criteria {
    metric_namespace = "Microsoft.Web/sites"
    metric_name      = "Http4xx"
    aggregation      = "Count"
    operator         = "GreaterThan"
    threshold        = 30
    skip_metric_validation = true

  }

  action {
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}

resource "azurerm_monitor_metric_alert" "app_5xx_status_codes" {
  provider            = azurerm.runtime
  name                = "app-5xx-status-codes${var.environment}"
  resource_group_name = azurerm_resource_group.rg.name
  scopes              = [azurerm_linux_web_app.app.id]
  description         = "${var.environment} - Web App: more than 10 5xx Errors are returned per hour."
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
    action_group_id = azurerm_monitor_action_group.ag_teams.id
  }
}
