resource "azurerm_linux_web_app" "app" {
  provider                  = azurerm.runtime
  location                  = var.location
  name                      = "${var.app_name}-${var.environment}-${random_id.app_id.hex}"
  resource_group_name       = azurerm_resource_group.rg.name
  service_plan_id           = azurerm_service_plan.plan.id
  https_only                = true
  virtual_network_subnet_id = azurerm_subnet.runtime_subnet.id
  identity {
    type         = "SystemAssigned, UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.app.id]
  }
  connection_string {
    name  = "db"
    type  = "PostgreSQL"
    value = "postgres://${random_pet.postgres_admin_username.id}:${random_password.postgres_admin_password.result}@${azurerm_postgresql_flexible_server.postgres.fqdn}/${azurerm_postgresql_flexible_server_database.db.name}?sslmode=require"
  }
  site_config {
    ftps_state                                    = "Disabled"
    container_registry_use_managed_identity       = true
    container_registry_managed_identity_client_id = azurerm_user_assigned_identity.app.client_id
    application_stack {
      docker_image     = "ealen/echo-server"
      docker_image_tag = "latest"
    }
  }
  lifecycle {
    ignore_changes = [
      site_config.0.application_stack.0.docker_image,
      site_config.0.application_stack.0.docker_image_tag,
      app_settings["DOCKER_CUSTOM_IMAGE_NAME"],
    ]
  }
}

